package coinapp

import (
	"fmt"
	"reflect"
)

type AssetData struct {
	Id                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

type Asset struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func (a AssetData) PrintCoinappInfo() {
	// ANSI escape codes for different colors
	colors := []string{"\x1b[31m", "\x1b[32m", "\x1b[34m"}
	// ANSI escape code to reset text color to default
	resetColor := "\x1b[0m"

	// Get the type of the struct
	structType := reflect.TypeOf(a)

	// Iterate through the struct fields
	for i := 0; i < structType.NumField(); i++ {
		// Get the corresponding color for this message
		color := colors[i%len(colors)]
		field := structType.Field(i)
		fieldValue := reflect.ValueOf(a).Field(i).Interface()
		fieldName := field.Name
		formattedMessage := fmt.Sprintf("[%s :: %s%s%s]", fieldName, color, fieldValue, resetColor)
		fmt.Println(formattedMessage)
	}

}
