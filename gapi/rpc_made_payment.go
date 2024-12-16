package gapi

import (
	"context"
	"fmt"
	"math/big"

	"time"

	"github.com/ebukacodes21/peerbill-user-api/pb"
	"github.com/ebukacodes21/peerbill-user-api/validation"

	tp "github.com/ebukacodes21/peerbill-trader-api/pb"
	tu "github.com/ebukacodes21/peerbill-trader-api/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (s *GServer) MadePayment(ctx context.Context, req *pb.MadePaymentRequest) (*pb.MadePaymentResponse, error) {
	violations := validateMadePaymentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	conn, err := grpc.NewClient("0.0.0.0:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create grpc connection: %s", err)
	}

	client := tp.NewPeerbillTraderClient(conn)
	args := &tp.GetOrderRequest{
		Id:        req.GetId(),
		OrderType: req.GetOrderType(),
	}

	resp, err := client.GetOrder(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get order: %s", err)
	}

	wallet, err := s.repository.GetWallet(ctx, resp.Order.EscrowAddress)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no wallet found: %v", err)
	}

	balance := tu.CheckBalance(ctx, resp.Order.Crypto, wallet.PublicKey)
	balanceBigInt, success := new(big.Int).SetString(balance, 10)
	if !success {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse balance string: %v", balance)
	}

	orderAmountBigInt := float64ToBigInt(float64(resp.Order.CryptoAmount), 18)
	if balanceBigInt.Cmp(orderAmountBigInt) < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "insufficient balance: have %s, need %s", balanceBigInt.String(), orderAmountBigInt.String())
	}

	// If the order is already received, just move the crypto to the user
	if resp.Order.IsReceived {
		err := transferCrypto(resp.Order.Crypto, wallet.PrivateKey, req.GetUserAddress(), *orderAmountBigInt, req.GetEscrowAddress())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to send crypto: %v", err)
		}

		// Update the order status to indicate that the order is completed and not expired
		updateArgs := &tp.UpdateOrderRequest{
			Id:        req.GetId(),
			Username:  req.GetUsername(),
			OrderType: req.GetOrderType(),
			IsExpired: false,
		}

		result, err := client.UpdateOrder(ctx, updateArgs)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update order after receiving payment: %v", err)
		}

		// Return the success response
		resp := &pb.MadePaymentResponse{
			Message: result.Message,
		}
		return resp, nil
	}

	// If the order is not received, check if time has passed and balance is available
	if time.Now().After(resp.Order.Duration.AsTime()) {
		// If there is enough balance in the wallet, transfer the crypto to the escrow (or user) >=
		if balanceBigInt.Cmp(orderAmountBigInt) < 0 {
			// Transfer the crypto from the wallet to the escrow address (or user address if applicable)
			err := transferCrypto(resp.Order.Crypto, wallet.PrivateKey, req.GetUserAddress(), *orderAmountBigInt, req.GetEscrowAddress())
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to send crypto to escrow: %v", err)
			}

			// Update the order status to completed and not expired
			updateArgs := &tp.UpdateOrderRequest{
				Id:        req.GetId(),
				Username:  req.GetUsername(),
				OrderType: req.GetOrderType(),
				IsExpired: false,
			}

			result, err := client.UpdateOrder(ctx, updateArgs)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update order after successful payment: %v", err)
			}

			// Return the success response
			resp := &pb.MadePaymentResponse{
				Message: result.Message,
			}
			return resp, nil
		} else {
			// If the balance is insufficient, mark the order as expired and completed
			updateArgs := &tp.UpdateOrderRequest{
				Id:        req.GetId(),
				Username:  req.GetUsername(),
				OrderType: req.GetOrderType(),
				IsExpired: true,
			}

			result, err := client.UpdateOrder(ctx, updateArgs)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update order after insufficient balance: %v", err)
			}

			// Return the failure message
			resp := &pb.MadePaymentResponse{
				Message: result.Message,
			}
			return resp, nil
		}
	}

	return nil, status.Errorf(codes.FailedPrecondition, "order is not completed and 30 minutes have not passed since the order's duration")
}

func validateMadePaymentRequest(req *pb.MadePaymentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := validation.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validation.ValidateType(req.GetOrderType()); err != nil {
		violations = append(violations, fieldViolation("order_type", err))
	}

	if err := validation.ValidateWalletAddress(req.GetUserAddress()); err != nil {
		violations = append(violations, fieldViolation("user_address", err))
	}

	if err := validation.ValidateWalletAddress(req.GetEscrowAddress()); err != nil {
		violations = append(violations, fieldViolation("escrow_address", err))
	}

	return violations
}

func transferCrypto(cryptoType, privateKey, userAddress string, amount big.Int, escrowAddress string) error {
	switch cryptoType {
	case "ETH":
		tu.SendNative("https://eth-mainnet.g.alchemy.com/v2/xtngPwzqpjqcWBvKoVKLiBwYo1kTbWxe", privateKey, userAddress, amount)
	case "BNB":
		tu.SendNative("https://bsc-dataseed.binance.org/", privateKey, userAddress, amount)
	case "USDT":
		tu.SendNative("https://bsc-dataseed.binance.org/", "admin-wallet-private-key", escrowAddress, amount)
		time.Sleep(3 * time.Second)
		tu.SendTokens("https://bsc-dataseed.binance.org/", privateKey, userAddress, "0x55d398326f99059fF775485246999027B3197955", amount)
	case "USDC":
		tu.SendNative("https://bsc-dataseed.binance.org/", "admin-wallet-private-key", escrowAddress, amount)
		time.Sleep(3 * time.Second)
		tu.SendTokens("https://bsc-dataseed.binance.org/", privateKey, userAddress, "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", amount)
	default:
		return fmt.Errorf("unsupported crypto type: %s", cryptoType)
	}
	return nil
}

func float64ToBigInt(amount float64, decimals int) *big.Int {
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	value := new(big.Float).SetFloat64(amount)
	value.Mul(value, new(big.Float).SetInt(multiplier))
	intValue, _ := value.Int(nil)
	return intValue
}
