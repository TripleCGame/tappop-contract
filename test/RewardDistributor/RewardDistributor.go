// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RewardDistributor

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RewardDistributorReceiverInfo is an auto generated low-level Go binding around an user-defined struct.
type RewardDistributorReceiverInfo struct {
	TokenAddress common.Address
	Amount       *big.Int
}

// RewardDistributorMetaData contains all meta data concerning the RewardDistributor contract.
var RewardDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebugLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"season\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRewardDistributor.ReceiverInfo[]\",\"name\":\"receiverInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"uniqueId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRewardDistributor.ReceiverInfo[]\",\"name\":\"receiverInfos\",\"type\":\"tuple[]\"}],\"name\":\"receiverInfosToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSinger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5033600660016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506124cd8061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630677a84e14610064578063238ac933146100945780638129fc1c146100b2578063856d18f5146100bc578063f25259f6146100d8578063f851a440146100f4575b5f80fd5b61007e6004803603810190610079919061144f565b610112565b60405161008b919061150a565b60405180910390f35b61009c6101d5565b6040516100a99190611569565b60405180910390f35b6100ba6101fa565b005b6100d660048036038101906100d19190611634565b610369565b005b6100f260048036038101906100ed91906116eb565b610816565b005b6100fc6108e9565b6040516101099190611569565b60405180910390f35b60605f60405180602001604052805f81525090505f8484905090505f5b818110156101c9575f61016a87878481811061014e5761014d611716565b5b9050604002015f01602081019061016591906116eb565b61090f565b90505f61019288888581811061018357610182611716565b5b90506040020160200135610c0a565b90508482826040516020016101a99392919061177d565b60405160208183030381529060405294505050808060010191505061012f565b50819250505092915050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102819061181d565b60405180910390fd5b60065f9054906101000a900460ff16156102d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d090611885565b60405180910390fd5b61034d6040518060400160405280601181526020017f5265776172644469737472696275746f720000000000000000000000000000008152506040518060400160405280600581526020017f312e302e30000000000000000000000000000000000000000000000000000000815250610d88565b600160065f6101000a81548160ff021916908315150217905550565b5f878560405160200161037d929190611908565b604051602081830303815290604052905060085f8581526020019081526020015f205f9054906101000a900460ff16156103ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e39061197d565b60405180910390fd5b6009816040516103fc919061199b565b90815260200160405180910390205f9054906101000a900460ff1615610457576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044e906119fb565b60405180910390fd5b61046e6104678989898989610e63565b8484610f20565b6104ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a490611a63565b60405180910390fd5b600160085f8681526020019081526020015f205f6101000a81548160ff02191690831515021790555060016009826040516104e8919061199b565b90815260200160405180910390205f6101000a81548160ff0219169083151502179055505f5b878790508110156106b2575f88888381811061052d5761052c611716565b5b9050604002015f01602081019061054491906116eb565b90505f89898481811061055a57610559611716565b5b9050604002016020013590505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cb90611acb565b60405180910390fd5b5f8111610616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060d90611b33565b60405180910390fd5b5f61062083610fc6565b90507f205f2aa8b8e5c1833d51be8c38a914476253c19ae627270a65c00878a095afef838a83856040516106579493929190611b60565b60405180910390a1818110156106a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069990611bed565b60405180910390fd5b505050808060010191505061050e565b505f5b8787905081101561080b575f8888838181106106d4576106d3611716565b5b9050604002015f0160208101906106eb91906116eb565b90505f89898481811061070157610700611716565b5b9050604002016020013590505f8290505f8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8b856040518363ffffffff1660e01b815260040161074d929190611c0b565b6020604051808303815f875af1158015610769573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078d9190611c67565b90508061079986610c0a565b6040516020016107a99190611d02565b604051602081830303815290604052906107f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f0919061150a565b60405180910390fd5b505050505080806001019150506106b5565b505050505050505050565b600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089d9061181d565b60405180910390fd5b8060075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f8273ffffffffffffffffffffffffffffffffffffffff165f1b90505f6040518060400160405280601081526020017f303132333435363738396162636465660000000000000000000000000000000081525090505f602a67ffffffffffffffff81111561098257610981611d23565b5b6040519080825280601f01601f1916602001820160405280156109b45781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106109eb576109ea611716565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610a4e57610a4d611716565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b6014811015610bfe5782600485600c84610a989190611d7d565b60208110610aa957610aa8611716565b5b1a60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c60f81c60ff1681518110610ae857610ae7611716565b5b602001015160f81c60f81b82600283610b019190611db0565b6002610b0d9190611d7d565b81518110610b1e57610b1d611716565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535082600f60f81b85600c84610b609190611d7d565b60208110610b7157610b70611716565b5b1a60f81b1660f81c60ff1681518110610b8d57610b8c611716565b5b602001015160f81c60f81b82600283610ba69190611db0565b6003610bb29190611d7d565b81518110610bc357610bc2611716565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053508080600101915050610a7e565b50809350505050919050565b60605f8203610c50576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d83565b5f8290505f5b5f8214610c7f578080610c6890611df1565b915050600a82610c789190611e65565b9150610c56565b5f8167ffffffffffffffff811115610c9a57610c99611d23565b5b6040519080825280601f01601f191660200182016040528015610ccc5781602001600182028036833780820191505090505b5090505f8290505b5f8614610d7b57600181610ce89190611e95565b90505f600a8088610cf99190611e65565b610d039190611db0565b87610d0e9190611e95565b6030610d1a9190611ed4565b90505f8160f81b905080848481518110610d3757610d36611716565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610d729190611e65565b97505050610cd4565b819450505050505b919050565b5f8054906101000a900460ff1615610dd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcc90611f78565b60405180910390fd5b5f828051906020012090505f828051906020012090505f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90508260038190555081600481905550610e2561113c565b600281905550610e36818484611143565b6001819055508060058190555060015f806101000a81548160ff0219169083151502179055505050505050565b5f808585905011610ea9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea090611fe0565b60405180910390fd5b5f610eb48686610112565b9050610f147fc697b3a3c419e7cc8c65645114138804ea3d4b56344df25385cfbb66f9ba02798886868580519060200120604051602001610ef9959493929190612016565b60405160208183030381529060405280519060200120611183565b91505095945050505050565b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610fa68585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050506111bb565b73ffffffffffffffffffffffffffffffffffffffff161490509392505050565b5f806040518060400160405280601281526020017f62616c616e63654f66286164647265737329000000000000000000000000000081525080519060200120306040516024016110169190611569565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f808473ffffffffffffffffffffffffffffffffffffffff168360405161109c91906120ab565b5f60405180830381855afa9150503d805f81146110d4576040519150601f19603f3d011682016040523d82523d5f602084013e6110d9565b606091505b50915091508161111e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111159061210b565b60405180910390fd5b80806020019051810190611132919061213d565b9350505050919050565b5f46905090565b5f83838361114f61113c565b30604051602001611164959493929190612168565b6040516020818303038152906040528051906020012090509392505050565b5f61118c611231565b8260405160200161119e929190612223565b604051602081830303815290604052805190602001209050919050565b5f6041825114611200576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f7906122a3565b60405180910390fd5b5f805f602085015192506040850151915060608501515f1a905061122686828585611263565b935050505092915050565b5f60025461123d61113c565b0361124c576001549050611260565b61125d600554600354600454611143565b90505b90565b5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0825f1c11156112c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c090612331565b60405180910390fd5b601b8460ff1614806112de5750601c8460ff16145b61131d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611314906123bf565b60405180910390fd5b5f6001868686866040515f815260200160405260405161134094939291906123ec565b6020604051602081039080840390855afa158015611360573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036113da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d190612479565b60405180910390fd5b80915050949350505050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261140f5761140e6113ee565b5b8235905067ffffffffffffffff81111561142c5761142b6113f2565b5b602083019150836040820283011115611448576114476113f6565b5b9250929050565b5f8060208385031215611465576114646113e6565b5b5f83013567ffffffffffffffff811115611482576114816113ea565b5b61148e858286016113fa565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6114dc8261149a565b6114e681856114a4565b93506114f68185602086016114b4565b6114ff816114c2565b840191505092915050565b5f6020820190508181035f83015261152281846114d2565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6115538261152a565b9050919050565b61156381611549565b82525050565b5f60208201905061157c5f83018461155a565b92915050565b5f819050919050565b61159481611582565b811461159e575f80fd5b50565b5f813590506115af8161158b565b92915050565b6115be81611549565b81146115c8575f80fd5b50565b5f813590506115d9816115b5565b92915050565b5f8083601f8401126115f4576115f36113ee565b5b8235905067ffffffffffffffff811115611611576116106113f2565b5b60208301915083600182028301111561162d5761162c6113f6565b5b9250929050565b5f805f805f805f60a0888a03121561164f5761164e6113e6565b5b5f61165c8a828b016115a1565b975050602088013567ffffffffffffffff81111561167d5761167c6113ea565b5b6116898a828b016113fa565b9650965050604061169c8a828b016115cb565b94505060606116ad8a828b016115a1565b935050608088013567ffffffffffffffff8111156116ce576116cd6113ea565b5b6116da8a828b016115df565b925092505092959891949750929550565b5f60208284031215611700576116ff6113e6565b5b5f61170d848285016115cb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81905092915050565b5f6117578261149a565b6117618185611743565b93506117718185602086016114b4565b80840191505092915050565b5f611788828661174d565b9150611794828561174d565b91506117a0828461174d565b9150819050949350505050565b7f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f5f8201527f6e2e000000000000000000000000000000000000000000000000000000000000602082015250565b5f6118076022836114a4565b9150611812826117ad565b604082019050919050565b5f6020820190508181035f830152611834816117fb565b9050919050565b7f696e697469616c697a653a20416c726561647920696e697469616c697a6564215f82015250565b5f61186f6020836114a4565b915061187a8261183b565b602082019050919050565b5f6020820190508181035f83015261189c81611863565b9050919050565b5f819050919050565b6118bd6118b882611582565b6118a3565b82525050565b5f8160601b9050919050565b5f6118d9826118c3565b9050919050565b5f6118ea826118cf565b9050919050565b6119026118fd82611549565b6118e0565b82525050565b5f61191382856118ac565b60208201915061192382846118f1565b6014820191508190509392505050565b7f556e69717565496420416c7265616479204469737472696275746564210000005f82015250565b5f611967601d836114a4565b915061197282611933565b602082019050919050565b5f6020820190508181035f8301526119948161195b565b9050919050565b5f6119a6828461174d565b915081905092915050565b7f536561736f6e20546f20416c72656164792044697374726962757465642100005f82015250565b5f6119e5601e836114a4565b91506119f0826119b1565b602082019050919050565b5f6020820190508181035f830152611a12816119d9565b9050919050565b7f496e76616c6964207369676e61747572652100000000000000000000000000005f82015250565b5f611a4d6012836114a4565b9150611a5882611a19565b602082019050919050565b5f6020820190508181035f830152611a7a81611a41565b9050919050565b7f496e76616c696420746f6b656e20616464726573732e000000000000000000005f82015250565b5f611ab56016836114a4565b9150611ac082611a81565b602082019050919050565b5f6020820190508181035f830152611ae281611aa9565b9050919050565b7f496e76616c696420616d6f756e742e00000000000000000000000000000000005f82015250565b5f611b1d600f836114a4565b9150611b2882611ae9565b602082019050919050565b5f6020820190508181035f830152611b4a81611b11565b9050919050565b611b5a81611582565b82525050565b5f608082019050611b735f83018761155a565b611b80602083018661155a565b611b8d6040830185611b51565b611b9a6060830184611b51565b95945050505050565b7f496e73756666696369656e742062616c616e63652e00000000000000000000005f82015250565b5f611bd76015836114a4565b9150611be282611ba3565b602082019050919050565b5f6020820190508181035f830152611c0481611bcb565b9050919050565b5f604082019050611c1e5f83018561155a565b611c2b6020830184611b51565b9392505050565b5f8115159050919050565b611c4681611c32565b8114611c50575f80fd5b50565b5f81519050611c6181611c3d565b92915050565b5f60208284031215611c7c57611c7b6113e6565b5b5f611c8984828501611c53565b91505092915050565b7f546f6b656e207472616e73666572206661696c656420666f7220746f6b656e205f8201527f696e646578200000000000000000000000000000000000000000000000000000602082015250565b5f611cec602683611743565b9150611cf782611c92565b602682019050919050565b5f611d0c82611ce0565b9150611d18828461174d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d8782611582565b9150611d9283611582565b9250828201905080821115611daa57611da9611d50565b5b92915050565b5f611dba82611582565b9150611dc583611582565b9250828202611dd381611582565b91508282048414831517611dea57611de9611d50565b5b5092915050565b5f611dfb82611582565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e2d57611e2c611d50565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611e6f82611582565b9150611e7a83611582565b925082611e8a57611e89611e38565b5b828204905092915050565b5f611e9f82611582565b9150611eaa83611582565b9250828203905081811115611ec257611ec1611d50565b5b92915050565b5f60ff82169050919050565b5f611ede82611ec8565b9150611ee983611ec8565b9250828201905060ff811115611f0257611f01611d50565b5b92915050565b7f656970496e697469616c697a65643a20416c726561647920696e697469616c695f8201527f7a65642100000000000000000000000000000000000000000000000000000000602082015250565b5f611f626024836114a4565b9150611f6d82611f08565b604082019050919050565b5f6020820190508181035f830152611f8f81611f56565b9050919050565b7f696e76616c6964207265636569766572496e666f7321000000000000000000005f82015250565b5f611fca6016836114a4565b9150611fd582611f96565b602082019050919050565b5f6020820190508181035f830152611ff781611fbe565b9050919050565b5f819050919050565b61201081611ffe565b82525050565b5f60a0820190506120295f830188612007565b6120366020830187611b51565b612043604083018661155a565b6120506060830185611b51565b61205d6080830184612007565b9695505050505050565b5f81519050919050565b5f81905092915050565b5f61208582612067565b61208f8185612071565b935061209f8185602086016114b4565b80840191505092915050565b5f6120b6828461207b565b915081905092915050565b7f45787465726e616c2063616c6c206661696c65640000000000000000000000005f82015250565b5f6120f56014836114a4565b9150612100826120c1565b602082019050919050565b5f6020820190508181035f830152612122816120e9565b9050919050565b5f815190506121378161158b565b92915050565b5f60208284031215612152576121516113e6565b5b5f61215f84828501612129565b91505092915050565b5f60a08201905061217b5f830188612007565b6121886020830187612007565b6121956040830186612007565b6121a26060830185611b51565b6121af608083018461155a565b9695505050505050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f6121ed600283611743565b91506121f8826121b9565b600282019050919050565b5f819050919050565b61221d61221882611ffe565b612203565b82525050565b5f61222d826121e1565b9150612239828561220c565b602082019150612249828461220c565b6020820191508190509392505050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e677468005f82015250565b5f61228d601f836114a4565b915061229882612259565b602082019050919050565b5f6020820190508181035f8301526122ba81612281565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f61231b6022836114a4565b9150612326826122c1565b604082019050919050565b5f6020820190508181035f8301526123488161230f565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f6123a96022836114a4565b91506123b48261234f565b604082019050919050565b5f6020820190508181035f8301526123d68161239d565b9050919050565b6123e681611ec8565b82525050565b5f6080820190506123ff5f830187612007565b61240c60208301866123dd565b6124196040830185612007565b6124266060830184612007565b95945050505050565b7f45434453413a20696e76616c6964207369676e617475726500000000000000005f82015250565b5f6124636018836114a4565b915061246e8261242f565b602082019050919050565b5f6020820190508181035f83015261249081612457565b905091905056fea2646970667358221220d192fd88685bfa1493bfc48b4493139ec09c20d3949e258dacbbd62bfa472ee664736f6c63430008190033",
}

// RewardDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardDistributorMetaData.ABI instead.
var RewardDistributorABI = RewardDistributorMetaData.ABI

// RewardDistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardDistributorMetaData.Bin instead.
var RewardDistributorBin = RewardDistributorMetaData.Bin

// DeployRewardDistributor deploys a new Ethereum contract, binding an instance of RewardDistributor to it.
func DeployRewardDistributor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RewardDistributor, error) {
	parsed, err := RewardDistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardDistributorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardDistributor{RewardDistributorCaller: RewardDistributorCaller{contract: contract}, RewardDistributorTransactor: RewardDistributorTransactor{contract: contract}, RewardDistributorFilterer: RewardDistributorFilterer{contract: contract}}, nil
}

// RewardDistributor is an auto generated Go binding around an Ethereum contract.
type RewardDistributor struct {
	RewardDistributorCaller     // Read-only binding to the contract
	RewardDistributorTransactor // Write-only binding to the contract
	RewardDistributorFilterer   // Log filterer for contract events
}

// RewardDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardDistributorSession struct {
	Contract     *RewardDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardDistributorCallerSession struct {
	Contract *RewardDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RewardDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardDistributorTransactorSession struct {
	Contract     *RewardDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RewardDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardDistributorRaw struct {
	Contract *RewardDistributor // Generic contract binding to access the raw methods on
}

// RewardDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardDistributorCallerRaw struct {
	Contract *RewardDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// RewardDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardDistributorTransactorRaw struct {
	Contract *RewardDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardDistributor creates a new instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributor(address common.Address, backend bind.ContractBackend) (*RewardDistributor, error) {
	contract, err := bindRewardDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardDistributor{RewardDistributorCaller: RewardDistributorCaller{contract: contract}, RewardDistributorTransactor: RewardDistributorTransactor{contract: contract}, RewardDistributorFilterer: RewardDistributorFilterer{contract: contract}}, nil
}

// NewRewardDistributorCaller creates a new read-only instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorCaller(address common.Address, caller bind.ContractCaller) (*RewardDistributorCaller, error) {
	contract, err := bindRewardDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorCaller{contract: contract}, nil
}

// NewRewardDistributorTransactor creates a new write-only instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardDistributorTransactor, error) {
	contract, err := bindRewardDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorTransactor{contract: contract}, nil
}

// NewRewardDistributorFilterer creates a new log filterer instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardDistributorFilterer, error) {
	contract, err := bindRewardDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorFilterer{contract: contract}, nil
}

// bindRewardDistributor binds a generic wrapper to an already deployed contract.
func bindRewardDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributor *RewardDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributor.Contract.RewardDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributor *RewardDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.Contract.RewardDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributor *RewardDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributor.Contract.RewardDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributor *RewardDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributor *RewardDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributor *RewardDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributor.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorSession) Admin() (common.Address, error) {
	return _RewardDistributor.Contract.Admin(&_RewardDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) Admin() (common.Address, error) {
	return _RewardDistributor.Contract.Admin(&_RewardDistributor.CallOpts)
}

// ReceiverInfosToString is a free data retrieval call binding the contract method 0x0677a84e.
//
// Solidity: function receiverInfosToString((address,uint256)[] receiverInfos) pure returns(string)
func (_RewardDistributor *RewardDistributorCaller) ReceiverInfosToString(opts *bind.CallOpts, receiverInfos []RewardDistributorReceiverInfo) (string, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "receiverInfosToString", receiverInfos)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReceiverInfosToString is a free data retrieval call binding the contract method 0x0677a84e.
//
// Solidity: function receiverInfosToString((address,uint256)[] receiverInfos) pure returns(string)
func (_RewardDistributor *RewardDistributorSession) ReceiverInfosToString(receiverInfos []RewardDistributorReceiverInfo) (string, error) {
	return _RewardDistributor.Contract.ReceiverInfosToString(&_RewardDistributor.CallOpts, receiverInfos)
}

// ReceiverInfosToString is a free data retrieval call binding the contract method 0x0677a84e.
//
// Solidity: function receiverInfosToString((address,uint256)[] receiverInfos) pure returns(string)
func (_RewardDistributor *RewardDistributorCallerSession) ReceiverInfosToString(receiverInfos []RewardDistributorReceiverInfo) (string, error) {
	return _RewardDistributor.Contract.ReceiverInfosToString(&_RewardDistributor.CallOpts, receiverInfos)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RewardDistributor *RewardDistributorSession) Signer() (common.Address, error) {
	return _RewardDistributor.Contract.Signer(&_RewardDistributor.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) Signer() (common.Address, error) {
	return _RewardDistributor.Contract.Signer(&_RewardDistributor.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0x856d18f5.
//
// Solidity: function distribute(uint256 season, (address,uint256)[] receiverInfos, address to, uint256 uniqueId, bytes signature) returns()
func (_RewardDistributor *RewardDistributorTransactor) Distribute(opts *bind.TransactOpts, season *big.Int, receiverInfos []RewardDistributorReceiverInfo, to common.Address, uniqueId *big.Int, signature []byte) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "distribute", season, receiverInfos, to, uniqueId, signature)
}

// Distribute is a paid mutator transaction binding the contract method 0x856d18f5.
//
// Solidity: function distribute(uint256 season, (address,uint256)[] receiverInfos, address to, uint256 uniqueId, bytes signature) returns()
func (_RewardDistributor *RewardDistributorSession) Distribute(season *big.Int, receiverInfos []RewardDistributorReceiverInfo, to common.Address, uniqueId *big.Int, signature []byte) (*types.Transaction, error) {
	return _RewardDistributor.Contract.Distribute(&_RewardDistributor.TransactOpts, season, receiverInfos, to, uniqueId, signature)
}

// Distribute is a paid mutator transaction binding the contract method 0x856d18f5.
//
// Solidity: function distribute(uint256 season, (address,uint256)[] receiverInfos, address to, uint256 uniqueId, bytes signature) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) Distribute(season *big.Int, receiverInfos []RewardDistributorReceiverInfo, to common.Address, uniqueId *big.Int, signature []byte) (*types.Transaction, error) {
	return _RewardDistributor.Contract.Distribute(&_RewardDistributor.TransactOpts, season, receiverInfos, to, uniqueId, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardDistributor *RewardDistributorTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardDistributor *RewardDistributorSession) Initialize() (*types.Transaction, error) {
	return _RewardDistributor.Contract.Initialize(&_RewardDistributor.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RewardDistributor *RewardDistributorTransactorSession) Initialize() (*types.Transaction, error) {
	return _RewardDistributor.Contract.Initialize(&_RewardDistributor.TransactOpts)
}

// SetSinger is a paid mutator transaction binding the contract method 0xf25259f6.
//
// Solidity: function setSinger(address _signer) returns()
func (_RewardDistributor *RewardDistributorTransactor) SetSinger(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "setSinger", _signer)
}

// SetSinger is a paid mutator transaction binding the contract method 0xf25259f6.
//
// Solidity: function setSinger(address _signer) returns()
func (_RewardDistributor *RewardDistributorSession) SetSinger(_signer common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetSinger(&_RewardDistributor.TransactOpts, _signer)
}

// SetSinger is a paid mutator transaction binding the contract method 0xf25259f6.
//
// Solidity: function setSinger(address _signer) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) SetSinger(_signer common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetSinger(&_RewardDistributor.TransactOpts, _signer)
}

// RewardDistributorDebugLogIterator is returned from FilterDebugLog and is used to iterate over the raw logs and unpacked data for DebugLog events raised by the RewardDistributor contract.
type RewardDistributorDebugLogIterator struct {
	Event *RewardDistributorDebugLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardDistributorDebugLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributorDebugLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardDistributorDebugLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardDistributorDebugLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributorDebugLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributorDebugLog represents a DebugLog event raised by the RewardDistributor contract.
type RewardDistributorDebugLog struct {
	TokenAddress common.Address
	To           common.Address
	Balance      *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDebugLog is a free log retrieval operation binding the contract event 0x205f2aa8b8e5c1833d51be8c38a914476253c19ae627270a65c00878a095afef.
//
// Solidity: event DebugLog(address tokenAddress, address to, uint256 balance, uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) FilterDebugLog(opts *bind.FilterOpts) (*RewardDistributorDebugLogIterator, error) {

	logs, sub, err := _RewardDistributor.contract.FilterLogs(opts, "DebugLog")
	if err != nil {
		return nil, err
	}
	return &RewardDistributorDebugLogIterator{contract: _RewardDistributor.contract, event: "DebugLog", logs: logs, sub: sub}, nil
}

// WatchDebugLog is a free log subscription operation binding the contract event 0x205f2aa8b8e5c1833d51be8c38a914476253c19ae627270a65c00878a095afef.
//
// Solidity: event DebugLog(address tokenAddress, address to, uint256 balance, uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) WatchDebugLog(opts *bind.WatchOpts, sink chan<- *RewardDistributorDebugLog) (event.Subscription, error) {

	logs, sub, err := _RewardDistributor.contract.WatchLogs(opts, "DebugLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributorDebugLog)
				if err := _RewardDistributor.contract.UnpackLog(event, "DebugLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebugLog is a log parse operation binding the contract event 0x205f2aa8b8e5c1833d51be8c38a914476253c19ae627270a65c00878a095afef.
//
// Solidity: event DebugLog(address tokenAddress, address to, uint256 balance, uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) ParseDebugLog(log types.Log) (*RewardDistributorDebugLog, error) {
	event := new(RewardDistributorDebugLog)
	if err := _RewardDistributor.contract.UnpackLog(event, "DebugLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
