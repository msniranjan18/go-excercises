package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cmdm CMDManager) ReadLines() ([]string, error) {
	var prices []string
	var price string
	fmt.Println("Eneter price list:")
	for {
		fmt.Scan(&price)
		prices = append(prices, price)
		if price == "0" {
			break
		}
	}
	return prices, nil
}

func (cmdm CMDManager) WriteJson(data any) error {
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(data)
	//fmt.Println(data)
	if err != nil {
		return fmt.Errorf("failed to encode data into the JSON")
	}
	return nil
}
