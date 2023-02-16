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
	puid := "copy-from-cookies"

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
		{
			Name:  "_puid",
			Value: puid,
		},
	}

	cli = chatgpt.NewClient(
		chatgpt.WithDebug(false),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)
}
