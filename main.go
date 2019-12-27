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


	//loads the private key from config.toml
	config, _ := toml.LoadFile("config.toml")
	pkstring := config.Get("VSYS.privatekey").(string)
	//b := config.Get("VSYS.prediction").(string)

	//prints private key to terminal for verificaqtion
	//fmt.Println("Your Private Key is: ", pkstring)

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

	for i := 1; i<=5; i++ {
		z := string(i)
		tx := acc.BuildPayment("ARFWV2aphzfZ5VKLk6xgxPEZhumnSgQBU7y", 1e6, z)
		resp, err := vsys.SendPaymentTx(tx)
		if err != nil {
			fmt.Println("There's been an error!")
			fmt.Println(err)
		}
		fmt.Println(resp.Error)
		fmt.Println(resp.Id)
	}
}
