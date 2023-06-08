package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var rates []*rate

type rate struct {
	Sell    float32
	Buy     float32
	currncy string
}

func (rate *rate) String() string {
	return fmt.Sprintf("%s TO GEL: %.4f GEL TO %s: %.4f", rate.currncy, rate.Buy, rate.currncy, rate.Sell)
}

func main() {
	Client := &http.Client{}
	currensySlice := []string{"USD", "EUR", "RUB", "GBP"}
	for _, currncy := range currensySlice {

		getRate(Client, currncy)
	}
	for _, doneRate := range rates {
		fmt.Println(doneRate)
	}
}

func getRate(Client *http.Client, currency string) {
	getRq, err := http.NewRequest(http.MethodGet, "https://api.businessonline.ge/api/rates/commercial/"+currency, nil)
	if err != nil {
		log.Println(err)
	}

	ens, err := Client.Do(getRq)
	if err != nil {
		log.Println(err)
	}

	if ens != nil && ens.Body != nil {
		readenbody, err := io.ReadAll(ens.Body)
		defer ens.Body.Close()

		if err != nil {
			log.Println(err)
		}
		var rate rate
		json.Unmarshal(readenbody, &rate)
		rate.currncy = currency
		rates = append(rates, &rate)
	}
}
