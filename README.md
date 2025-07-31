# nova-api-client-go

A simple client for the [Wandelbots Nova](https://www.wandelbots.com/) API.

## basic usage

This Example will list all available controllers and print their joint configurations.
```golang
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	
	v2 "github.com/wandelbotsgmbh/nova-api-client-go/v25/pkg/nova/v2"
)

func main() {
	// replace with actual token or load from env (only needed when developing locally against nova cloud instances)
	token := ""
	// replace with your instance url (keep the /api/v2 at the end)
	host := "https://id.instance.wandelbots.io/api/v2"
	// the default cell
	cell := "cell"

	client, err := v2.NewClientWithResponses(host)
	if token != "" {
		client, err = v2.NewClientWithResponses(host, withAuthToken(token))
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	resp, err := client.ListRobotControllersWithResponse(context.TODO(), cell)
	if err != nil {
		log.Fatalf("Failed to list robot controllers: %v", err)
	}

	controllers := resp.JSON200

	if controllers == nil {
		log.Fatalf("No controllers found or response was empty: %s", string(resp.Body))
	}

	for _, controller := range *controllers {
		printMotionGroupPositions(client, cell, controller)
	}
}

func printMotionGroupPositions(client *v2.ClientWithResponses, cell, controller string) {
	resp, err := client.GetControllerDescriptionWithResponse(context.TODO(), cell, controller)
	if err != nil {
		log.Printf("Failed to get robot controller %s: %v", controller, err)
		return
	}

	if resp.JSON200 == nil {
		log.Printf("No data found for robot controller %s", controller)
		return
	}

	motionGroups := resp.JSON200.ConnectedMotionGroups
	for _, mg := range motionGroups {
		printMotionGroupPosition(client, cell, controller, mg)
	}
}

func printMotionGroupPosition(client *v2.ClientWithResponses, cell, controller, motionGroup string) {
	resp, err := client.GetMotionGroupStateWithResponse(context.TODO(), cell, controller, motionGroup)
	if err != nil {
		log.Printf("Failed to get motion group %s state: %v", motionGroup, err)
		return
	}
	if resp.JSON200 == nil {
		log.Printf("No data found for motion group %s on controller %s", motionGroup, controller)
		return
	}
	fmt.Println("controller:", controller, "motionGroup:", motionGroup, "joint positions:", resp.JSON200.Positions)
}

func withAuthToken(token string) v2.ClientOption {
	return v2.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
}
```
