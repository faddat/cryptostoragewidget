package main

import (
	"fmt"
	"github.com/walkbean/vsys-sdk-go/vsys"
	"io/ioutil"
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
	fmt.Print(string(privatekey))
	pkstring := string(privatekey)

	//Do actual, interesting things

	//initalize account
	acc := vsys.InitAccount(vsys.Mainnet)
	acc.BuildFromPrivateKey(pkstring)
	//b := []byte("Lets test right on mainnet")
	tx := acc.BuildPayment("ARFWV2aphzfZ5VKLk6xgxPEZhumnSgQBU7y", 1e8, []byte{})
	resp, err := vsys.SendPaymentTx(tx)
	fmt.Println(resp.Error)
	fmt.Println(resp.Id)
}
