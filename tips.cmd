@REM MetaMask
@REM https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn?hl=en-US
@REM https://faucet.rinkeby.io/

@REM Node JS
@REM https://nodejs.org/en/download/

@REM Truffle
@REM npm install -g truffle
@REM npm install @openzeppelin/contracts
@REM npm install @truffle/hdwallet-provider
@REM npm install truffle-privatekey-provider
@REM npm install dotenv

@REM https://trufflesuite.com/ganache/
@REM npm install -g ganache-cli

@REM Infura
@REM https://infura.io/

@REM truffle compile

@REM Collect imports for validate on etherscan.io by single file:
@REM npm run build-contracts

@REM Run new console window and network
@REM ganache-cli --networkId 4447

@REM Deploy and Run local tests:

@REM truffle migrate --f 1 --to 1 --network development

@REM This project tests:
@REM truffle test ./test/testing.js --compile-none --migrations_directory migrations_null

@REM Call solidity contract from Golang
@REM go get -u github.com/ethereum/go-ethereum
@REM https://geth.ethereum.org/downloads/

@REM solcjs --abi contracts/TrainingContract.sol -o build

@REM abigen --abi=./build/contracts_TrainingContract_sol_TrainingContract.abi --pkg=TrainingContract --out=TrainingContract.go
