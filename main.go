package main

import (
	"github.com/isvaldo/go-coinmarketcap-client"
	"github.com/isvaldo/crypto-cli/coin"
	"os"
	"os/exec"
	"time"
	"fmt"
)

func main() {

	service := coin.Service{
		Client: coinmarket.New("https://api.coinmarketcap.com"),
	}

	lock := time.Tick(time.Duration(time.Second*20))
	service.TopCoinsTable(30)
	for tick := range lock{
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		service.TopCoinsTable(30)
		fmt.Println(fmt.Sprintf("last refresh %s",tick.Format("15:04:05")))
	}

}
