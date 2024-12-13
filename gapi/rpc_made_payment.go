package gapi

import (
	"context"
	"fmt"

	"time"

	"github.com/ebukacodes21/peerbill-user-api/pb"
	"github.com/ebukacodes21/peerbill-user-api/utils"
	"github.com/ebukacodes21/peerbill-user-api/validation"

	tp "github.com/ebukacodes21/peerbill-trader-api/pb"
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
		return nil, status.Errorf(codes.Internal, "unable to run grpc conn %s", err)
	}
	client := tp.NewPeerbillTraderClient(conn)

	args := &tp.GetOrderRequest{
		Id:        req.GetId(),
		OrderType: req.GetOrderType(),
	}

	resp, err := client.GetOrder(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get order %s", err)
	}

	wallet, err := s.repository.GetWallet(ctx, resp.Order.EscrowAddress)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no wallet found: %v", err)
	}

	if resp.Order.IsReceived || time.Now().After(resp.Order.Duration.AsTime()) {
		// move crypto from escrow to user address
		err := transferCrypto(resp.Order.Crypto, wallet.PrivateKey, req.GetUserAddress(), float64(resp.Order.CryptoAmount), req.GetEscrowAddress())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to send crypto: %v", err)
		}

		// post to update
		args := &tp.UpdateOrderRequest{
			Id:        req.GetId(),
			Username:  req.GetUsername(),
			OrderType: req.GetOrderType(),
		}

		result, err := client.UpdateOrder(ctx, args)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update order: %v", err)
		}

		resp := &pb.MadePaymentResponse{
			Message: result.Message,
		}

		return resp, nil

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

func transferCrypto(cryptoType, privateKey, userAddress string, amount float64, escrowAddress string) error {
	switch cryptoType {
	case "ETH":
		return utils.SendNative("https://eth-mainnet.g.alchemy.com/v2/xtngPwzqpjqcWBvKoVKLiBwYo1kTbWxe", privateKey, userAddress, amount)
	case "BNB":
		return utils.SendNative("https://bsc-dataseed.binance.org/", privateKey, userAddress, amount)
	case "USDT":
		if err := utils.SendNative("https://bsc-dataseed.binance.org/", "admin-wallet-private-key", escrowAddress, amount); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
		return utils.SendTokens("https://bsc-dataseed.binance.org/", privateKey, userAddress, "0x55d398326f99059fF775485246999027B3197955", amount, 6)
	case "USDC":
		if err := utils.SendNative("https://bsc-dataseed.binance.org/", "admin-wallet-private-key", escrowAddress, amount); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
		return utils.SendTokens("https://bsc-dataseed.binance.org/", privateKey, userAddress, "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", amount, 6)
	default:
		return fmt.Errorf("unsupported crypto type: %s", cryptoType)
	}
}
