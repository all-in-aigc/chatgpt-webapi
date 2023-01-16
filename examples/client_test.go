package examples

import (
	"fmt"
	"net/http"
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
	token := `copy-from-cookies`
	cfValue := "copy-from-cookies"
	userAgent := "copy-your-user-agent"

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
	}

	cli = chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
		chatgpt.WithUserAgent(userAgent),
	)
}
