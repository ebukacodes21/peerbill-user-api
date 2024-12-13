package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func SendNative(provider string, privateKeyString string, receiverAddress string, amount float64) error {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Print(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Print(err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Print(err)
		return err
	}

	sender := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Print(err)
		return err
	}

	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	receiver := common.HexToAddress(receiverAddress)

	amountInWei := new(big.Int)
	amountInWei.SetString(fmt.Sprintf("%f", amount*math.Pow(10, 18)), 10)

	// unsigned tx ===
	tx := types.NewTransaction(nonce, receiver, amountInWei, gasLimit, gasPrice, nil)
	// sign tx
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Print(err)
		return err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Print(err)
		return err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func SendTokens(provider string, privateKeyString string, receiverAddr string, contractAddr string, tokenAmount float64, decimals int) error {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Print(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Print(err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Print(err)
		return err
	}

	sender := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Print(err)
		return err
	}

	value := big.NewInt(0)
	toAddress := common.HexToAddress(receiverAddr)
	tokenAddress := common.HexToAddress(contractAddr)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	tokenAmountScaled := tokenAmount * math.Pow(10, float64(decimals))
	tokenAmountStr := strconv.FormatFloat(tokenAmountScaled, 'f', -1, 64)

	// Convert string to big.Int
	tokenAmountInt := new(big.Int)
	tokenAmountInt, ok = tokenAmountInt.SetString(tokenAmountStr, 10)
	if !ok {
		log.Print("Failed to convert token amount to big.Int")
		return err
	}

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(tokenAmountInt.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Print(err)
		return err
	}
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Print(err)
		return err
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Print(err)
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Print(err)
		return err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
