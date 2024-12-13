package gapi

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	db "github.com/ebukacodes21/peerbill-user-api/db/sqlc"
	"github.com/ebukacodes21/peerbill-user-api/pb"
	"github.com/ebukacodes21/peerbill-user-api/validation"

	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (s *GServer) GenWallet(ctx context.Context, req *pb.GenWalletRequest) (*pb.GenWalletResponse, error) {
	violations := validateGenWalletRquest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type of public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	createWalletParams := db.CreateWalletParams{
		PrivateKey: privateKeyHex,
		PublicKey:  address.String(),
	}

	wallet, err := s.repository.CreateWallet(ctx, createWalletParams)
	if err != nil {
		log.Fatal("failed to create wallet in database: %w ", err)
	}

	result := &pb.GenWalletResponse{
		Address: wallet.PublicKey,
	}
	return result, nil
}

func validateGenWalletRquest(req *pb.GenWalletRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateCrypto(req.GetCrypto()); err != nil {
		violations = append(violations, fieldViolation("crypto", err))
	}

	return violations
}
