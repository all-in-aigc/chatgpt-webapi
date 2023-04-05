package examples

import (
	"fmt"
	"log"
)

func ExampleClient_GetModels() {
	cli := getClient()

	res, cookies, err := cli.GetModels()

	if err != nil {
		log.Fatalf("get models failed: %v\n", err)
	}

	for _, v := range res.Get("models").Array() {
		log.Printf("model: %s\n", v.String())
	}

	for _, v := range cookies {
		log.Printf("cookie: %s, %s, expires: %v\n", v.Name, v.Value, v.Expires)
	}

	fmt.Println("xxx")
	// Output: xxx
}
