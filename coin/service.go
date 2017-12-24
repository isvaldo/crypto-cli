package coin

import (
	"github.com/isvaldo/go-coinmarketcap-client"
	"github.com/pkg/errors"
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
	"time"
	"strconv"
)

type Service struct {
	Client coinmarket.Interface
}


func (c *Service) TopCoinsTable(limit int) error {
	coins,err :=c.Client.GetTickerWithLimits(limit)
	if err != nil {
		return errors.Wrap(err,"Error while fetch coin data")
	}
	table := tablewriter.NewWriter(os.Stdout)
	var body [][]string
	for _, coin := range coins.TickerList{
		n, _ := strconv.Atoi(coin.LastUpdated)
		s := time.Unix(int64(n), 64).Format("15:04:05")
		body= append(body, []string{
			coin.Rank,
			coin.Name,
			coin.Symbol,
			fmt.Sprintf("$%s",coin.PriceUsd),
			fmt.Sprintf("%s%%",coin.PercentChange1H),
			fmt.Sprintf("%s%%",coin.PercentChange24H),
			fmt.Sprintf("%s%%",coin.PercentChange7D),
			s,
		})
	}

	headers := []string{
		"Rank",
		"Name",
		"Symbol",
		"PriceUsd",
		"PercentChange1H",
		"PercentChange24H",
		"PercentChange7D",
		"LastUpdated"}



	table.SetHeader(headers)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.FgHiGreenColor,tablewriter.Bold},
		)
	for _, v := range body {
		table.Append(v)
	}
	table.Render()
	return nil
}