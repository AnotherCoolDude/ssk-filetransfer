package handler

import (
	"errors"
	"fmt"
	"github.com/AnotherCoolDude/ssk-filetransfer/basecampclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	clientsMap map[string]int
)

func init() {
	clientsMap = map[string]int{}
}

func handleClient(c *gin.Context) error {
	shortname := c.Query("shortname")
	if shortname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing shortname query parameter"})
		return errors.New("missing query parameter: shortname")
	}

	if len(clientsMap) == 0 {
		clientsMap[shortname] = 0
		return nil
	}

	if _, ok := clientsMap[shortname]; ok {
		basecampclient.ChangeClient(clientsMap[shortname])
		return nil
	}
	clientsMap[shortname] = basecampclient.AddClient()
	basecampclient.ChangeClient(clientsMap[shortname])
	return nil
}

// BCLoginhandler redirects to the bascamp auth url
func BCLoginhandler(c *gin.Context) {
	err := handleClient(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if basecampclient.Client.TokenValid() {
		fmt.Println("token still valid")
		c.JSON(http.StatusAccepted, gin.H{"token": "still valid"})
		return
	}
	accessurl := basecampclient.Client.AuthCodeURL()
	c.JSON(http.StatusAccepted, gin.H{"redirectURL": accessurl})
}

// BCTokenValidHandler validates wether the token in client is valid
func BCTokenValidHandler(c *gin.Context) {
	err := handleClient(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	valid := basecampclient.Client.TokenValid()
	c.JSON(http.StatusOK, gin.H{"tokenValid": valid})
}

// BCCallbackhandler handles the callack from basecamp auth
func BCCallbackhandler(c *gin.Context) {
	basecampclient.Client.HandleCallback(c.Request)
	basecampclient.Client.ReceiveID()
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200/")
}

// BCGetProjects returns all projects
func BCGetProjects(c *gin.Context) {
	resp, err := basecampclient.Client.Do("GET", "/projects.json", http.NoBody)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp.Body)
}
