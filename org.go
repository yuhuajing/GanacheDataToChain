package main

import (
	"fmt"
	"main/common/config"
	"main/core/deploy"
	"main/explorer"
	"os"
)

func main() {
	if config.TraceSC == "" {
		err, address := deploy.Deploy()
		if err == nil {
			fmt.Println(address)
			config.TraceSC = address
		} else {
			fmt.Println(err)
			fmt.Println("合约部署失败！！")
			os.Exit(0)
		}
	}
	explorer.Explorer()
}
