package main

// Owner's public and private keys are used for signing

import (
	"TgJumpGameContract/test/RewardDistributor"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestInitialize(t *testing.T) {

	client, err := ethclient.Dial(config.ServerAddress)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(config.OwnerPrivateKey)
	if err != nil {
		panic(err)
	}
	auth, err := Author(privateKey, client, 500000, 0, 0)
	if err != nil {
		panic(err)
	}
	medal, err := RewardDistributor.NewRewardDistributor(common.HexToAddress(config.ContractAddress), client)
	if err != nil {
		panic(err)
	}

	// Set signer to the public key corresponding to signPrev
	tx, err := medal.Initialize(auth)
	if err != nil {
		panic("error")
	}
	if tx == nil {
		panic("error")
	}
	hash := tx.Hash().Hex()
	fmt.Printf("setsigner hash: %v", hash)
}
