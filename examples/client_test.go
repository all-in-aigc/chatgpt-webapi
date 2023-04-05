package examples

import (
	"fmt"
	"net/http"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

var (
	debug        bool
	accessToken  string
	sessionToken string
	cfValue      string
	puid         string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *chatgpt.Client
}

func getClient() *chatgpt.Client {
	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: sessionToken,
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

	cli := chatgpt.NewClient(
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithDebug(debug),
		chatgpt.WithAccessToken(accessToken),
		chatgpt.WithCookies(cookies),
	)

	return cli
}
