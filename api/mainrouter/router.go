package mainrouter

import (
	"ginserver/api/servertype"
	"ginserver/api/testapi"
	"ginserver/bin"

	"github.com/gin-gonic/gin"
)

func BuildMasterInterface(
	router *gin.Engine,
	config bin.Config,
	testmode bool,
) {
	server := servertype.Server{
		Router:   router,
		Testmode: testmode,
	}
	testapi.BuildInterface(server)

	//server.AuthRoutes = router.Group("/api").Use(middleware.AuthMiddleware(server.TokenMaker, server.Testmode))
	//user.BuildInterface(server)
}
