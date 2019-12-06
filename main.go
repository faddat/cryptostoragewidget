package main

import (
	"fmt"
	"github.com/virtualeconomy/go-v-sdk/vsys"
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
	b := []byte("Lets test right on mainnet")
	tx := acc.BuildPayment("ARFvzdhQKJ5GX5haE6UjQHKCaPrL9kCnfiR", 1e8, b)
	resp, err := vsys.SendPaymentTx(tx)
	fmt.Println(resp)
}
