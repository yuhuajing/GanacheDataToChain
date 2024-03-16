package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
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

func AddOrUpdateUserData(number,
	workamount,
	persion,
	workmethod,
	worktime,
	remarks string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}

	err, auth := gentx(80000)
	if auth == nil {
		return ""
	}
	tx, err := instance.AddorupdateData(auth, number,
		workamount,
		persion,
		workmethod,
		worktime,
		remarks)
	if err != nil {
		fmt.Println("error creating instance:%s", err)
		return ""
		//log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
