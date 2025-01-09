package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		Message   string      `json:"message,omitempty"`
		Timestamp int64       `json:"timestamp,omitempty"`
		Body      interface{} `json:"body,omitempty"`
	}

	nowFunc func() time.Time
)

var (
	now nowFunc = time.Now
)

func respond(ctx *gin.Context, code int, response Response, now nowFunc) {
	response.Timestamp = now().Unix()
	ctx.JSONP(code, response)
}

func FormatResponse(ctx *gin.Context, code int, message string, body interface{}) {
	response := Response{}
	if body != nil {
		response.Body = body
	}
	response.Message = message

	respond(ctx, code, response, now)
}
