package gapi

import (
	"bytes"
	"context"
	"encoding/json"

	"net/http"
	"time"

	"github.com/ebukacodes21/peerbill-user-api/pb"
	"github.com/ebukacodes21/peerbill-user-api/utils"
	"github.com/ebukacodes21/peerbill-user-api/validation"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderRequest struct {
	Id        int64  `json:"id"`
	OrderType string `json:"order_type"`
}

type OrderResponse struct {
	ID            int64      `json:"id"`
	EscrowAddress string     `json:"escrow_address"`
	Crypto        string     `json:"crypto"`
	Fiat          string     `json:"fiat"`
	FiatAmount    float64    `json:"fiat_amount"`
	CryptoAmount  float64    `json:"crypto_amount"`
	Username      string     `json:"username"`
	Rate          float64    `json:"rate"`
	IsAccepted    bool       `json:"is_accepted"`
	IsCompleted   bool       `json:"is_completed"`
	IsRejected    bool       `json:"is_rejected"`
	IsReceived    bool       `json:"is_received"`
	CreatedAt     *time.Time `json:"created_at"`
	Duration      *time.Time `json:"duration"`
	UserAddress   string     `json:"user_address"`
	OrderType     string     `json:"order_type"`
	BankName      *string    `json:"bank_name,omitempty"`
	AccountNumber *string    `json:"account_number,omitempty"`
	AccountHolder *string    `json:"account_holder,omitempty"`
}

func (s *GServer) MadePayment(ctx context.Context, req *pb.MadePaymentRequest) (*pb.MadePaymentResponse, error) {
	violations := validateMadePaymentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	reqData := OrderRequest{
		Id:        req.GetId(),
		OrderType: req.GetOrderType(),
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal request data: %v", err)
	}

	resp, err := http.Post("http://localhost:8002/api/get-order", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch order details: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "failed to fetch order details: %v", resp.Status)
	}

	var order OrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to decode order response: %v", err)
	}

	wallet, err := s.repository.GetWallet(ctx, order.EscrowAddress)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no wallet found: %v", err)
	}

	if order.IsReceived || time.Now().After(*order.Duration) {
		// move crypto from escrow to user address
		if order.Crypto == "ETH" {
			utils.SendNative("https://eth-mainnet.g.alchemy.com/v2/xtngPwzqpjqcWBvKoVKLiBwYo1kTbWxe", wallet.PrivateKey, req.GetUserAddress(), order.CryptoAmount)
		}

		if order.Crypto == "BNB" {
			// utils.SendNative("")
		}

		// post to update
		resp := &pb.MadePaymentResponse{
			Message: "crypto has been sent to your wallet",
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
