package servertype

import (
	"ginserver/bin"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config     bin.Config
	Router     *gin.Engine
	AuthRoutes gin.IRoutes
	Testmode   bool
}

var AuthorizationHeaderKey = "Authorization"
var AuthorizationTypeBearer = "bearer"
var AuthorizationPayloadKey = "authorization_payload"
