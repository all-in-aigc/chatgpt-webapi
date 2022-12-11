package examples

import (
	"fmt"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

// chatgpt client
var cli *chatgpt.Client

func ExampleNewClient() {
	fmt.Printf("%T", cli)

	// Output: *chatgpt.Client
}

func init() {
	// your __Secure-next-auth.session-token
	token := `YOUR-SESSION-TOKEN`
	cli = chatgpt.NewClient(
		chatgpt.WithToken(token),
		chatgpt.WithTimeout(30*time.Second),
	)
}
