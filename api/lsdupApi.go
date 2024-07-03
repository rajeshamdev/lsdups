package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LsdupHealth(c *gin.Context) {

	c.String(http.StatusOK, "I am healthy")
}

func LsdupGet(c *gin.Context) {

	c.String(http.StatusOK, "welcome to lsdup server")
}

func LsdupShutdown(c *gin.Context) {

	// sigChan <- syscall.SIGINT
}
