package main

import (
	"fmt"
	"github.com/walkbean/vsys-sdk-go/vsys"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Connect to VSYS public API
	vsys.InitApi("https://wallet.v.systems/api", vsys.Mainnet)

	//read private key from file
	seed, err := ioutil.ReadFile("seed")
	check(err)
	seedstring := string(seed)
	//prints the seed for user verification
	fmt.Println(seedstring)


	//initalize account
	acc := vsys.InitAccount(vsys.Mainnet)
	acc.BuildFromSeed(seedstring, 0)
	info, err := acc.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	//Print generated address
	fmt.Println("Your address is:")
	fmt.Println(info.Address)

	//b := []byte("Lets test right on mainnet")
	tx := acc.BuildPayment("ARFWV2aphzfZ5VKLk6xgxPEZhumnSgQBU7y", 1e8, []byte{})
	resp, err := vsys.SendPaymentTx(tx)


	if err != nil {
		fmt.Println("There's been an error!")
		fmt.Println(err)
	}

	fmt.Println(resp.Error)
	fmt.Println(resp.Id)
}
