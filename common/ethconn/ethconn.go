package ethconn

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnBlockchain(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	if err != nil {
		fmt.Println("Conn Chain err, 稍后在试!")
		return nil
		//log.Fatalf("Eth connect error:%v", err)
	}
	return nclient
}
