package main

import (
	"fmt"
	"github.com/virtualeconomy/go-v-sdk/vsys"
)

func main() {



vsys.InitApi("https://wallet.v.systems/api", vsys.Mainnet)
acc := vsys.InitAccount(Mainnet)



fmt.Println("Chicken")


}
