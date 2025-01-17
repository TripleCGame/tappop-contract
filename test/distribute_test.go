package main

// The public and private keys used for signing are also used by the owner

import (
	"TgJumpGameContract/test/RewardDistributor"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestStr(t *testing.T) {
	amountBig := big.NewInt(10000000000000000)
	amount2n64, _ := strconv.ParseInt("10000000000000000", 10, 64)
	amount2Big := big.NewInt(amount2n64)
	if amountBig.Cmp(amount2Big) != 0 {
		panic("Conversion error:")
	}
}

// Distribution Test
func TestDistribute(t *testing.T) {
	fmt.Printf("serverAddress:%v, contractAddress:%v, ownerPublicKey:%v, ownerPrivateKey:%v", config.ServerAddress, config.ContractAddress, config.OwnerPublicKey, config.OwnerPrivateKey)

	var (
		season       = int64(99999997) // Season
		uniqueId     = int64(99999997) // Global unique ID, can be the ID of the table, or generated by yourself, only used for preventing repeated minting
		tokenAddress = common.HexToAddress("0x681202351a488040Fa4FdCc24188AfB582c9DD62")
		amount       = big.NewInt(399000000) // 0.01 tokens, 6 decimal places
		to           = config.TestUserPublicKey
	)
	userPrivateKeyEcdsa, err := crypto.HexToECDSA(config.TestUserPirvateKey)
	if err != nil {
		panic(err)
	}

	// Connect to the node
	client, _ := ethclient.Dial(config.ServerAddress)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	receiverInfos := []RewardDistributor.RewardDistributorReceiverInfo{
		{
			TokenAddress: tokenAddress,
			Amount:       amount,
		},
	}

	// Server signing: using the private key of the contract manager
	sign, err := MedalSigner(config.OwnerPrivateKey, chainId.Int64(), config.ContractAddress, season, uniqueId, to, receiverInfos)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sign:%v\n", sign)

	// User information
	medal, err := RewardDistributor.NewRewardDistributor(common.HexToAddress(config.ContractAddress), client)
	if err != nil {
		panic(err)
	}

	// Client call claim, verify signature: signature
	{
		signBytes, err := hexutil.Decode(sign)
		if err != nil {
			panic(err)
		}

		// auth
		txOpts, err := Author(userPrivateKeyEcdsa, client, 500000, 0, 0)
		if err != nil {
			panic(err)
		}

		// Set the received token information
		fmt.Println("signBytes,", signBytes)
		tx, err := medal.Distribute(txOpts, big.NewInt(season), receiverInfos, common.HexToAddress(to), big.NewInt(uniqueId), signBytes)
		if err != nil {
			panic(err)
		}
		hash := tx.Hash().Hex()
		fmt.Printf("Distribute Hash:%s \n", hash)
	}

}
