package status

import (
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/appleboy/gofight.v1"
	"net/http"
	"testing"
)

func httpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/api/status", StatusHandler)

	return r
}

func TestStatusHandler(t *testing.T) {

	r := gofight.New()

	r.GET("/api/status").
		Run(httpRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "go_version")

			assert.NotEmpty(t, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
