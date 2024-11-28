package main

import (
	"TgJumpGameContract/test/RewardDistributor"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func MedalSigner(
	privateKeyStr string,
	chainId int64,
	contract string,
	season int64,
	uniqueId int64,
	to string,
	receiverInfos []RewardDistributor.RewardDistributorReceiverInfo,
) (string, error) {
	receiverInfoStr := receiverInfosToString(receiverInfos)

	// Sign
	typedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"distribute": {
				{Name: "season", Type: "uint256"},
				{Name: "to", Type: "address"},
				{Name: "uniqueId", Type: "uint256"},
				{Name: "receiverInfos", Type: "string"},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              "RewardDistributor",
			Version:           "1.0.0",
			ChainId:           math.NewHexOrDecimal256(chainId),
			VerifyingContract: contract,
			Salt:              "",
		},
		Message: map[string]interface{}{
			"season":        math.NewHexOrDecimal256(season),
			"to":            to,
			"uniqueId":      math.NewHexOrDecimal256(uniqueId),
			"receiverInfos": receiverInfoStr,
		},
		PrimaryType: "distribute",
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	signature, err := SignWithEip712(privateKey, typedData)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signature), nil
}

// receiverInfosToString function
func receiverInfosToString(receiverInfos []RewardDistributor.RewardDistributorReceiverInfo) string {
	var builder strings.Builder
	builder.WriteString("")

	for _, info := range receiverInfos {
		// Use lowercase encoding for addresses
		addressLower := strings.ToLower(info.TokenAddress.Hex())
		builder.WriteString(fmt.Sprintf("%s%d", addressLower, info.Amount))
	}

	s := builder.String()
	fmt.Println("receiverInfosToString: ", s)
	return s
}

func SignWithEip712(privateKey *ecdsa.PrivateKey, typedData *apitypes.TypedData) ([]byte, error) {
	if privateKey == nil || typedData == nil {
		return nil, errors.New("invalid parameter")
	}

	// 1. Get the Keccak-256 hash of the data to be signed
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}

	typedDataHashEncodedData, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
	if err != nil {
		return nil, err
	}

	fmt.Println("typedDataHashEncodedData: ", typedDataHashEncodedData)

	fmt.Printf("\n")
	fmt.Printf("domainSeparator: %s\n", hexutil.Encode(domainSeparator))
	fmt.Printf("typedDataHash: %s\n", hexutil.Encode(typedDataHash))

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sigHash := crypto.Keccak256(rawData)
	fmt.Printf("rawData: %s\n", hexutil.Encode(rawData))
	fmt.Printf("sigHash: %s\n", hexutil.Encode(sigHash))
	// 2. Sign the hash with private key to get the signature
	signature, err := crypto.Sign(sigHash, privateKey)
	if err != nil {
		return nil, err
	}
	if signature[64] < 27 {
		signature[64] += 27
	}
	return signature, nil
}

func Author(privateKey *ecdsa.PrivateKey, client *ethclient.Client, gasLimit int64, fixGasPrice float64, value int64) (auth *bind.TransactOpts, err error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, err
	}

	var gasPrice *big.Int
	if fixGasPrice > 0 {
		gasPrice = big.NewInt(int64(fixGasPrice * 1000000000))
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	}

	if err != nil {
		return nil, nil
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, nil
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(value))
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
