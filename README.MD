## Development environment
1. Install nodejs
2. Vscode Install solidity plugin和SCStudio plugin，press F5 to fast compilation
3. Install solc
```
npm install -g solc
```   
4. Install truffle  
```
npm install -g truffle
```
5. Install ganache for debug, websit:https://trufflesuite.com/ganache, then run ganache
6. Install truffle-flattener 
```
npm install -g truffle-flattener
```

## Preparation work
1. Place the wallet private key in the.env file
2. Modify truffle-config.js, write the private key into the variable priviteKeys, and configure the network environment
3. Run script truffle create migration develop_triplec_medal,automatically create deploy scripts and modify them

## truffle compile
truffle compile

## Deploying
1. Deploy contract to the local network
```
truffle migrate --network dev
```
2. Deploy contract to a public chain
```
sudo truffle migrate --network bsctest
```
3. Create full sol file
```
truffle-flattener ./contracts/medal/RewardDistributor.sol --output ./deployment/RewardDistributor.full.sol 
```
Then delete the redundant //SPDX-License-Identifier: MIT in fullsol, and finally upload the full.sol file to the contract source code through the blockchain browser

## Testing
1. Generate go resource

```
bash ./deployment/abigen.sh
```

2. Run script or test

in ./script path


## License

TripleC Contracts is released under the [MIT License](LICENSE).

