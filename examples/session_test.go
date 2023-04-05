package examples

import (
	"fmt"
	"log"
)

func ExampleClient_AuthSession() {
	cli := getClient()

	res, err := cli.AuthSession()

	if err != nil {
		log.Fatalf("auth session failed: %v\n", err)
	}

	log.Printf("session: %s\n", res)

	fmt.Println("xxx")
	// Output: xxx
}
