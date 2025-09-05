package endpoint

import (
	"github.com/charmingruby/pack/pkg/delivery/http/rest"
	"github.com/gin-gonic/gin"
)

const TIMEOUT_IN_SECONDS = 10

func makeReadiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		rest.SendOKResponse(c, "", nil)
	}
}
