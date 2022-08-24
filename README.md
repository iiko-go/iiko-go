### Note
Library is still under construction. Most of the methods are not implemented.

#### Example usage 👋
```golang
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iiko-go/iiko-go"
)

func main() {
	client, err := iiko.NewClient("API_LOGIN")
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can set a custom http.Client for making iikoCloud API request.
	// By default, http.DefaultClient is used.
	client.SetHTTPClient(&http.Client{})

	// You can set a custom timeout for making iikoCloud API request.
	// By default, 15 seconds is used.
	client.SetTimeout(5 * time.Second)

	// You can set a custom refreshTokenTimeout for iiko.Client.
	// refreshTokenTimeout is the interval at which the iiko.Client wil try to update the iikoCloud API token.
	// by default, 45 minutes is used.
	client.SetRefreshTokenTimeout(30 * time.Minute)

	stoplistRequest := &iiko.StopListsRequest{
		OrganizationIDs: []uuid.UUID{
			uuid.MustParse("18C40D75-BA2E-4AFA-9DDE-28C46E7A7CEE"),
			uuid.MustParse("80DE4E4A-E2E1-45E6-928B-93628D35F8C2"),
		},
	}

	// iikoCloud API: /api/1/stop_lists
	res, err := client.StopLists(stoplistRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can alswo use custom options.
	// For example, iiko.WithCustomTimeout() will set a custom timeout for one API request.
	res, err = client.StopLists(stoplistRequest, iiko.WithCustomTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Now we can work with IIKO response.

	fmt.Println(res.CorrelationID) // print OperationID.

	for _, terminalGroup := range res.TerminalGroupStopLists {
		fmt.Println(terminalGroup.OrganizationID) // print OrganizationID.

		for _, product := range terminalGroup.Items {
			fmt.Println(product.ProductID) // print ProductID in stoplist.
		}
	}
}
```
