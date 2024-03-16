package config

import (
	"github.com/ethereum/go-ethereum/crypto"
	"main/common/ethconn"
)

var (
	EthServer       = "http://127.0.0.1:7545"
	PrivateKey, err = crypto.HexToECDSA("fa3720302a59da92f8cfbd1ce96ce3cb045d63d71d3d151c7b1e5fcc8489cfd1")
	Client          = ethconn.ConnBlockchain(EthServer)
	TraceSC         = "0x59043EB527419a9fa400ECe202f9B6f2DeE71ee5"
)

// 浏览器
// https://sepolia.etherscan.io/address/0xC6F48C4a0EEc939b0841c07AAe1836b179EA7a84
