package main

// The public and private keys used for signing are also the owner's public and private keys

import (
	"fmt"
	"testing"

	"TgJumpGameContract/test/RewardDistributor"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestSetSinger(t *testing.T) {
	client, _ := ethclient.Dial(config.ServerAddress)
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

	// Set signer, to the public key of signPrev
	tx, err := medal.SetSinger(auth, common.HexToAddress(config.OwnerPublicKey))
	if err != nil {
		panic("error")
	}
	if tx == nil {
		panic("error")
	}
	hash := tx.Hash().Hex()
	fmt.Printf("setsigner hash:%v", hash)
}
