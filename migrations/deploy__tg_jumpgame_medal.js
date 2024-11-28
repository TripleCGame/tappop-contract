const tgJumpGameMedal = artifacts.require("RewardDistributor"); // 这里要使用sol文件中合约的名字

module.exports = function(_deployer,_network,_account) {
  // Use deployer to state migration tasks.
  console.info(_account);
  if(_network=="dev"){
    _deployer.deploy(tgJumpGameMedal,"TG.JumpGame","MEDAL",{overwrite:true});
  }else if (_network=="bsctest"){
    _deployer.deploy(tgJumpGameMedal,"TG.JumpGame","MEDAL",{overwrite:true});
  }else if (_network=="bsc"){
    _deployer.deploy(tgJumpGameMedal,"TG.JumpGame","MEDAL",{overwrite:true});
  }else if (_network=="x1test"){
    _deployer.deploy(tgJumpGameMedal,"TG.JumpGame","MEDAL",{overwrite:true});
  }
};
