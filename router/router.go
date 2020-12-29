package router

import (
	Log "eatwhat/log"
	"github.com/gin-gonic/gin"
)

/*
  @description router
  @author angellosetheirhair
  @date 2020/11/24
  @version 0.1
*/
type Router struct{

}
func (rt *Router) InitRouter(log *Log.Log) *gin.Engine {
	t:=gin.New()
	//register all  plugin
	t.Use(gin.LoggerWithFormatter(log.SetLogFormate))
	t.Use(gin.Recovery())
	//register all route
	route := new(Route)
    route.RegisterRoute(t)
	return t
}