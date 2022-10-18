package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

// {
// 	"ip": "197.210.76.155",
// 	"city": "Abuja",
// 	"region": "FCT",
// 	"country": "NG",
// 	"loc": "9.0579,7.4951",
// 	"org": "AS29465 MTN NIGERIA Communication limited",
// 	"timezone": "Africa/Lagos"
// 	}

type Ip struct {
	Ip       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Timezone string `json::"timezone"`
	Postal   string `json::"postal"`
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Traces your ip address",
	Long:  `traces your ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Printf("tracking ip %v\n", ip)
				showData(ip)
			}
		} else {
			fmt.Println("please enter an IP address to be traced")
		}
	},
}

func showData(ip string) {
	var url = "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)
	ipInfo := Ip{}
	err := json.Unmarshal(responseByte, &ipInfo)
	if err != nil {
		log.Println("cannot unmarshal data")
	}

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("DATA found! =>")
	fmt.Printf(" IP:%s\n CITY: %s COUNTRY: %s\n REGION: %s\n LOCATION: %s\n TIMEZONE: %s\n POSTAL: %s\n ",
		ipInfo.Ip, ipInfo.City, ipInfo.Country, ipInfo.Region, ipInfo.Loc, ipInfo.Timezone, ipInfo.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("could not get requested URL")
		return nil
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("cannot read bytes ")
	}
	return responseByte
}
