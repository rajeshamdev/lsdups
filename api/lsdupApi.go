package api

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rajeshamdev/lsdups/lsdups"
)

var (
	SigChan chan os.Signal
)

func LsdupHealth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "up",
	})
}

type Data struct {
	Dups map[string][]string `json:"dups"`
}

func LsdupGet(c *gin.Context) {

	dir := c.Query("dir")

	if dir == "" {
		dir = "."
	}

	dupsMap := lsdups.Lsdups(dir)

	data := Data{
		Dups: dupsMap,
	}

	c.JSON(http.StatusOK, data)
}

func LsdupShutdown(c *gin.Context) {

	SigChan <- syscall.SIGINT

	c.JSON(http.StatusOK, gin.H{
		"status": "shutdown",
	})
}
