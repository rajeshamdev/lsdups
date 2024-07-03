package api

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	SigChan chan os.Signal
)

func LsdupHealth(c *gin.Context) {

	c.String(http.StatusOK, "I am healthy")
}

func LsdupGet(c *gin.Context) {

	c.String(http.StatusOK, "welcome to lsdup server")
}

func LsdupShutdown(c *gin.Context) {

	SigChan <- syscall.SIGINT
}
