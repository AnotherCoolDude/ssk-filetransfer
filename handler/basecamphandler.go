package handler

import (
	"fmt"
	"net/http"

	"github.com/AnotherCoolDude/ssk-filetransfer/basecampclient"
	"github.com/gin-gonic/gin"
)

// BCLoginhandler redirects to the bascamp auth url
func BCLoginhandler(c *gin.Context) {
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
	valid := basecampclient.Client.TokenValid()
	c.JSON(http.StatusOK, gin.H{"tokenValid": valid})
}

// BCCallbackhandler handles the callack from basecamp auth
func BCCallbackhandler(c *gin.Context) {
	basecampclient.Client.HandleCallback(c.Request)
	basecampclient.Client.ReceiveID()
	fmt.Printf("client after callback %+v\n", basecampclient.Client)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200/basecamptable/HOVE")
}

// BCGetProjects returns all projects
func BCGetProjects(c *gin.Context) {
	resp, err := basecampclient.Client.Do("GET", "/projects.json", http.NoBody)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Printf("projects requesting client: %+v\n", basecampclient.Client)
	defer resp.Body.Close()
	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/json", resp.Body, map[string]string{})
}

// BCGetContentsByLink returns the content, that is fetchable by the provided link
func BCGetContentsByLink(c *gin.Context) {
	link := c.Query("link")
	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing link parameter in request"})
		return
	}
	resp, err := basecampclient.Client.Do("GET", link, http.NoBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// proadclient.PrettyPrintResponse(resp)
	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/json", resp.Body, map[string]string{})
}
