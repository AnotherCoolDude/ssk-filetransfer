package basecampclient

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/rs/xid"
	"golang.org/x/oauth2"
)

var (
	clients []*Basecampclient
	// Client represents a basecamp client
	Client *Basecampclient
	// AccessTypeWebserver adds a query parameter that sets type to web_server
	accessTypeWebserver = oauth2.SetAuthURLParam("type", "web_server")
)

func init() {
	clients = []*Basecampclient{}
	Client = new(
		"34a8b658744a7f190a03765149bbcf2282161add",
		"d0736bffee7d2be41f2f8f277a8a01b944912b4d",
		"christian.hovenbitzer@selinka-schmitz.de",
		"ProadInterface",
		"http://localhost:3000/callback",
	)
}

// Basecampclient wraps all necessary data for accessing basecamp api
type Basecampclient struct {
	email       string
	appName     string
	id          int
	oauthConfig *oauth2.Config
	state       string
	code        string
	token       *oauth2.Token
	httpclient  *http.Client
}

// new returns a new basecamp client
func new(clientID, clientSecret, email, appName, callbackURL string) *Basecampclient {
	c := &Basecampclient{
		appName: appName,
		id:      0,
		state:   xid.New().String(),
		code:    "",
		token:   &oauth2.Token{},
		oauthConfig: &oauth2.Config{
			RedirectURL:  callbackURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthStyle: oauth2.AuthStyleAutoDetect,
				AuthURL:   "https://launchpad.37signals.com/authorization/new",
				TokenURL:  "https://launchpad.37signals.com/authorization/token",
			},
		},
		httpclient: http.DefaultClient,
	}
	clients = append(clients, c)
	return c
}

// AddClient adds a new Client and returns the client's index
func AddClient() int {
	c := new(
		"34a8b658744a7f190a03765149bbcf2282161add",
		"d0736bffee7d2be41f2f8f277a8a01b944912b4d",
		"christian.hovenbitzer@selinka-schmitz.de",
		"ProadInterface",
		"http://localhost:3000/callback",
	)
	clients = append(clients, c)
	return len(clients) - 1
}

// ChangeClient changes the current client
func ChangeClient(index int) {
	if index > len(clients) {
		fmt.Println("[basecampclient/client.go/ChangeClient] index greater than clients array")
		return
	}
	Client = clients[index]
}

// AuthCodeURL returns a url, that asks bascamp for access
func (c *Basecampclient) AuthCodeURL() string {
	return c.oauthConfig.AuthCodeURL(c.state, accessTypeWebserver)
}

// HandleCallback extracts the code from basecamp and exchages it for a token
func (c *Basecampclient) HandleCallback(request *http.Request) {
	code := request.FormValue("code")
	state := request.FormValue("state")
	if state != c.state {
		fmt.Println("[basecampclient/client.go/handleCallback] state doesn't match")
		return
	}
	t, err := c.oauthConfig.Exchange(oauth2.NoContext, code, accessTypeWebserver)
	if err != nil {
		fmt.Println("[basecampclient/client.go/handleCallback] couldn't exchange token: " + err.Error())
		return
	}
	c.token = t
}

// ReceiveID receives the id of current user, which is needed for any other request
func (c *Basecampclient) ReceiveID() {
	if !c.token.Valid() {
		fmt.Println("need valid token to receive ID")
		return
	}
	resp, err := c.Do("GET", "https://launchpad.37signals.com/authorization.json", http.NoBody)
	if err != nil {
		fmt.Println("[basecampclient/client.go/handleCallback] couldn't make request to auth endpoint")
		return
	}
	respbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[basecampclient/client.go/handleCallback] couldn't read response body")
		return
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.Unmarshal(respbytes, &result)
	if err != nil {
		fmt.Println("[basecampclient/client.go/handleCallback] couldn't umarshal response body")
		return
	}
	identity := result["identity"].(map[string]interface{})
	c.id = int(identity["id"].(float64))
}

// TokenValid reports wether the token is non-nil and not expired
func (c *Basecampclient) TokenValid() bool {
	return c.token.Valid()
}

// RefreshToken refreshes the current token ToDo: flesh out function
func (c *Basecampclient) RefreshToken() {
	if c.token.Valid() {
		return
	}
	fmt.Println(c.token.RefreshToken)
}

func (c *Basecampclient) addHeader(request *http.Request) {
	request.Header.Add("Authorization", "Bearer "+c.token.AccessToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", fmt.Sprintf("%s (%s)", c.appName, c.email))
}

func (c *Basecampclient) baseURL() string {
	return fmt.Sprintf("https://3.basecampapi.com/%d/", c.id)
}
