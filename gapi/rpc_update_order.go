package gapi

import (
	"context"

	tp "github.com/ebukacodes21/peerbill-trader-api/pb"
	"github.com/ebukacodes21/peerbill-trader-api/validate"
	"github.com/ebukacodes21/peerbill-user-api/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (s *GServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	violations := validateUpdateOrderRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	conn, err := grpc.NewClient("0.0.0.0:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to connect to server %s", err)
	}

	client := tp.NewPeerbillTraderClient(conn)
	var bankName *string
	if bank := req.GetBankName(); bank != "" {
		bankName = &bank
	}
	var accountNumber *string
	if account := req.GetAccountNumber(); account != "" {
		accountNumber = &account
	}
	var accountHolder *string
	if holder := req.GetAccountHolder(); holder != "" {
		accountHolder = &holder
	}

	args := tp.UpdateOrderRequest{
		Id:            req.GetId(),
		Username:      req.GetUsername(),
		OrderType:     req.GetOrderType(),
		BankName:      bankName,
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
	}

	_, err = client.UpdateOrder(ctx, &args)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to update order %s", err)
	}

	resp := &pb.UpdateOrderResponse{
		Message: "update successful",
	}

	return resp, nil
}

func validateUpdateOrderRequest(req *pb.UpdateOrderRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := validate.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validate.ValidateType(req.GetOrderType()); err != nil {
		violations = append(violations, fieldViolation("order_type", err))
	}

	if err := validate.ValidateFirstname(req.GetAccountHolder()); err != nil {
		violations = append(violations, fieldViolation("account_holder", err))
	}

	if err := validate.ValidateFirstname(req.GetBankName()); err != nil {
		violations = append(violations, fieldViolation("bank_name", err))
	}

	if err := validate.ValidateFirstname(req.GetAccountNumber()); err != nil {
		violations = append(violations, fieldViolation("account_number", err))
	}

	return violations
}
