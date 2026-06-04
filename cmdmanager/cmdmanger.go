package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd *CMDManager) ReadFile() ([]string, error) {
	fmt.Println("Please enter your prices, Confirm every price with Enter:	")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)
		if price == "0" || price == "" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd *CMDManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
