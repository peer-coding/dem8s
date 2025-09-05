package platform

import (
	"github.com/charmingruby/pack/internal/platform/http/endpoint"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	endpoint.New(r)
}
