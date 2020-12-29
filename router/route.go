package router

import (
	"eatwhat/controller"
	"github.com/gin-gonic/gin"
)

/*
  @description route
  @author angels lose their hair
  @date 2020/11/24
  @version 0.1
*/
type Route struct {

}
var (
	menu *controller.Menu //menu controller
)
func (route *Route) RegisterRoute(t *gin.Engine){
	//menu page
	v1 := t.Group("/v1")
	v1.GET("/menus",menu.GetMenuList)
	v1.POST("/menu",menu.AddMenu)
	v1.PUT("/menu/:id/name/:name",menu.UpdateMenu)
	v1.DELETE("/menu/:id",menu.DeleteMenu)
}