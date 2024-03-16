package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"main/common/config"
	"main/trace"
	"math/big"
)

func gentx(gaslimit uint64) (error, *bind.TransactOpts) {
	publicKey := config.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("私钥解析失败！！")
		return fmt.Errorf("%s", "Private parse error"), nil
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	chianid, err := config.Client.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, chianid)
	if err != nil {
		fmt.Println("连接区块链失败，请重启区块链后端！！")
		return fmt.Errorf("%s", err.Error()), nil
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return nil, auth
}

func Deploy() (error, string) {
	err, auth := gentx(3000000)
	if err != nil {
		fmt.Print("GenTX err: ", err)
		return fmt.Errorf("%s", err.Error()), ""
	}

	address, _, _, err := trace.DeployTrace(auth, config.Client)
	if err != nil {
		fmt.Println("链拥堵，稍后再试！！")
		return fmt.Errorf("%s", err.Error()), ""
	}
	return nil, address.Hex()
}
