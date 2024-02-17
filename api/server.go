package api

import (
	"ginserver/api/mainrouter"
	"ginserver/bin"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config   bin.Config
	Router   *gin.Engine
	Testmode bool
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewServer(config bin.Config, testmode bool) (*Server, error) {
	server := &Server{
		Config:   config,
		Testmode: testmode,
	}

	router := gin.Default()
	//if !testmode {
	//	router.Use(loggerwrite.LoggerToFile())
	//}

	router.Use(CORSMiddleware())
	server.Router = router

	//登入驗證
	mainrouter.BuildMasterInterface(
		server.Router,
		server.Config,
		testmode,
	)

	server.Router = router
	return server, nil

}

func (server *Server) Start(address string) error {

	return server.Router.Run(address)
}
