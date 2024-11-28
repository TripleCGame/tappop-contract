package main

import (
	"TgJumpGameContract/test/RewardDistributor"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMedalNftSigner(t *testing.T) {
	const contractAddress = "0x98f564dd002b3712eb294ddcfd1c23eaa2bdb9bb"                // Contract address
	const signPrev = "cf4b4365e4bfa8b42a8fc95a74f6215dafedff21b701033936d0e3734608c63b" // Private key for signing, should be kept on server only, do not expose
	const chainId = 97

	var (
		season        = int64(1)                                     // Season
		uniqueId      = int64(7)                                     // Global unique ID, can be table ID or self-generated, used to prevent duplicate minting
		to            = "0x131d13bab97e20dab92761d0e91c43645cce90d6" // Address to mint to
		receiverInfos = []RewardDistributor.RewardDistributorReceiverInfo{
			{
				TokenAddress: common.HexToAddress("0xeB3Eb991D39Dac92616da64b7c6D5af5cCFf1627"),
				Amount:       big.NewInt(10000000000000000),
			},
		}
	)

	sign, err := MedalSigner(signPrev, chainId, contractAddress, season, uniqueId, to, receiverInfos)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sign)
}
