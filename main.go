package main

import (
	"fmt"
	"github.com/walkbean/vsys-sdk-go/vsys"
	"log"
	"github.com/pelletier/go-toml"
)


func main() {

	//Connect to VSYS public API
	vsys.InitApi("https://wallet.v.systems/api", vsys.Mainnet)


	//prints the seed or private key for user verification

	config, _ := toml.LoadFile("config.toml")
	pkstring := config.Get("VSYS.privatekey").(string)

	fmt.Println("Your Private Key is: ", pkstring)

	//initalize account
	acc := vsys.InitAccount(vsys.Mainnet)

	//acc.BuildFromSeed(seedstring, 0)
	acc.BuildFromPrivateKey(pkstring)
	info, err := acc.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	//Print generated address
	fmt.Println("Your address is: ", info.Address)

	b := "One more time just for the hell of it"
	tx := acc.BuildPayment("ARFWV2aphzfZ5VKLk6xgxPEZhumnSgQBU7y", 1e6, b)
	resp, err := vsys.SendPaymentTx(tx)


	if err != nil {
		fmt.Println("There's been an error!")
		fmt.Println(err)
	}

	fmt.Println(resp.Error)
	fmt.Println(resp.Id)
}
