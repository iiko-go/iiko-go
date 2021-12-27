### Example usage 👋
```golang
package main

import (
	"fmt"

	"github.com/iiko-go/iiko-go"
)

func main() {
	client, err := iiko.NewClient("API_LOGIN")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// iikoCloud API: /api/1/discounts
	res, err := client.Discounts(&iiko.DiscountsRequest{
		OrganizationIDs: []iiko.UUID{
			"18C40D75-BA2E-4AFA-9DDE-28C46E7A7CEE",
			"80DE4E4A-E2E1-45E6-928B-93628D35F8C2",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, discount := range res.Discounts {
		fmt.Println(discount.OrganizationID)
		for _, cardInfo := range discount.Items {
			fmt.Println(cardInfo.Name, cardInfo.Comment)
		}
	}
}
```
