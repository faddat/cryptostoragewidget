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
	privatekey, err := ioutil.ReadFile("privatekey")
	check(err)
	pkstring := string(privatekey)
	fmt.Println(pkstring)

	//Do actual, interesting things

	//initalize account
	acc := vsys.InitAccount(vsys.Mainnet)
	acc.BuildFromPrivateKey(pkstring)
	info, err := acc.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info.Address)
	//b := []byte("Lets test right on mainnet")
	tx := acc.BuildPayment("ARFWV2aphzfZ5VKLk6xgxPEZhumnSgQBU7y", 1e8, []byte{})
	resp, err := vsys.SendPaymentTx(tx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Error)
	fmt.Println(resp.Id)
}
