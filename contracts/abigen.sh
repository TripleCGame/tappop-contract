
# build
solc --abi --bin -o ../build --overwrite  RewardDistributor.sol

abigen --bin=../build/RewardDistributor.bin --abi=../build/RewardDistributor.abi --pkg=RewardDistributor --out=../script/RewardDistributor/RewardDistributor.go
