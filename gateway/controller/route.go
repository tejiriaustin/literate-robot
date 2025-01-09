package controller

import (
	"context"
	"errors"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/literate-robot/core/config"
	coreResponse "github.com/tejiriaustin/literate-robot/core/response"
)

func BindRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	conf *config.Environment,
) {
	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		coreResponse.FormatResponse(c, http.StatusOK, "Gateway is alive and well", nil)
	})

	r.GET("/service-info", func(c *gin.Context) {
		commitHash := getCommitFromBuildInfo()
		ipAddress, er := getOutboundIP()
		if er != nil {
			ipAddress = "**NOT_FOUND**"
		}

		payload := map[string]interface{}{
			"apiVersion":  "v1",
			"commit":      commitHash,
			"shortCommit": commitHash[:7],
			"env":         conf.GetAsString("ENVIRONMENT"),
			"outboundIP":  ipAddress,
		}
		coreResponse.FormatResponse(c, http.StatusOK, "Successful", payload)
	})
}

func getCommitFromBuildInfo() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}

func getOutboundIP() (string, error) {
	c, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("failure retrieving Outbound IP")
	}
	defer func() {
		err := c.Close()
		if err != nil {
			//Log error
		}
	}()

	a := c.LocalAddr().String()
	i := strings.LastIndex(a, ":")

	return a[0:i], nil
}
