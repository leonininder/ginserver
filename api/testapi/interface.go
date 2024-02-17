package testapi

import (
	"ginserver/api/servertype"
)

var server servertype.Server

func BuildInterface(getserver servertype.Server) {

	server = getserver
	//server.AuthRoutes.GET("/user")
	//server.AuthRoutes.POST("/jsontest", ClientCreateEvent)

	server.Router.POST("/jsontest", ClientCreateEvent)
}
