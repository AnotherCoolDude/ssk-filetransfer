package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AnotherCoolDude/ssk-filetransfer/basecampclient"
	"github.com/gin-gonic/gin"
)

var (
	clientsMap map[string]int
)

func init() {
	clientsMap = map[string]int{}
}

// BCVerifyClient verifies the client based on the provided shortname
func BCVerifyClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortname := c.Query("shortname")
		if shortname == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing shortname query parameter"})
			c.Abort()
			return
		}

		if len(clientsMap) == 0 {
			clientsMap[shortname] = 0
		}

		if _, ok := clientsMap[shortname]; ok {
			basecampclient.ChangeClient(clientsMap[shortname])
		} else {
			clientsMap[shortname] = basecampclient.AddClient()
			basecampclient.ChangeClient(clientsMap[shortname])
		}

		fmt.Printf("current client: %+v\n", basecampclient.Client)
		fmt.Println(c.HandlerName())
		if strings.Contains(c.HandlerName(), "BCLoginhandler") {
			fmt.Println("try to connect to login handler, letting through")
			c.Next()
			return
		}

		if !basecampclient.Client.TokenValid() {
			c.Redirect(http.StatusTemporaryRedirect, "/bc/login")
			c.Abort()
			return
		}

		if !basecampclient.Client.IDValid() {
			basecampclient.Client.ReceiveID()
			fmt.Printf("updated client: %+v\n", basecampclient.Client)
		}

		c.Next()
	}
}
