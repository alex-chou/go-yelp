package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alex-chou/go-yelp/yelp"
)

var (
	// apiKey as provided by Yelp when an app is created. Can be found here:
	// https://www.yelp.com/developers/v3/manage_app
	apiKey string

	// shared client for the example.
	client yelp.Client
)

func init() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("`API_KEY` must be specified")
	}
	client = yelp.New(http.DefaultClient, apiKey)
}

// Example usage: `API_KEY=api_key go run example/main.go`
func main() {
	ctx := context.Background()

	// Uncomment one to try out the API call.
	//RunBusinessSearch(ctx)
	RunGetBusiness(ctx)
}

// RunBusinessSearch makes a Business Search request.
func RunBusinessSearch(ctx context.Context) {
	results, err := client.BusinessSearch(ctx, &yelp.BusinessSearchOptions{
		Location: yelp.StringPointer("New York"),
	})
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(results)
}

// RunGetBusiness makes a Get Business request.
func RunGetBusiness(ctx context.Context) {
	business, err := client.GetBusiness(ctx, &yelp.GetBusinessOptions{
		ID: "nI1UYDCYUTt23TpGxqnLKg",
	})
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(business)
}

// prettyPrint prints the input.
func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
