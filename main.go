package main

import (
	"path"

	"github.com/AnotherCoolDude/ssk-filetransfer/handler"
	"github.com/AnotherCoolDude/ssk-filetransfer/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/callback", handler.BCCallbackhandler)

	proad := r.Group("/proad")
	proad.GET("/employees", handler.GetEmployeesHandler)
	proad.GET("/projects/:urno", handler.GetProjectsByEmployeeHandler)
	proad.GET("/projects", handler.GetFilteredProjects)

	files := r.Group("/files")
	files.GET("/project", handler.GetContentForProject)

	basecamp := r.Group("/bc")
	basecamp.Use(middleware.BCVerifyClient())
	basecamp.GET("/valid", handler.BCTokenValidHandler)
	basecamp.GET("/login", handler.BCLoginhandler)
	basecamp.GET("/projects", handler.BCGetProjects)
	basecamp.GET("/link", handler.BCGetContentsByLink)

	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := path.Ext(file)
		if file == "" || ext == "" {
			c.File("./frontend/dist/frontend/index.html")
		} else {
			c.File("./frontend/dist/frontend/" + path.Join(dir, file))
		}
	})

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
