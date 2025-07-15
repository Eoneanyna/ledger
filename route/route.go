package route

import (
	"github.com/gin-gonic/gin"
)

var PublicPrefix = "/api/v1"

func RegisterRoute(r *gin.Engine) {
	rootRoute := r.Group(PublicPrefix)
	RegisterLegerGroup(rootRoute)
	RegisterUserGroup(rootRoute)
}
