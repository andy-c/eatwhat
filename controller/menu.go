package controller

import (
	"eatwhat/model/form"
	"eatwhat/service/impl"
	"github.com/gin-gonic/gin"
)

/*
  @description home controller
  @author angels lose their hair
  @date 2020/11/24
  @version 0.1
*/

type Menu struct{
	Base
}
var (
	menuService *impl.Menu //menu service
)
//choose menu
func (menu *Menu) GetMenuList(c *gin.Context) {
	var menuform *form.MenuList
	if err := c.ShouldBindQuery(menuform);err!=nil{
		c.JSON(200,gin.H{"code":50001,"message":"menu name is empty","data":""})
		return
	}
    menulist:=menuService.GetMenuList(menuform)
	c.JSON(200,gin.H{"code":200,"message":"","data":menulist})
}

//insert menu
func (menu *Menu) AddMenu( c *gin.Context ){
	menuform:=new(form.MenuList)
    if err:=c.ShouldBind(menuform);err!=nil{
		c.JSON(200,gin.H{"code":50001,"message":"params are empty pls check","data":""})
		return
	}
	row:=menuService.AddMenu(menuform)
	c.JSON(200,gin.H{"row":row})
}

//update menu
func (menu *Menu) UpdateMenu(c *gin.Context){
	id:=c.Param("id")
	name:=c.Param("name")
	row:=menuService.UpdateMenu(id,name)
	c.JSON(200,gin.H{"row":row})
}

//delete menu
func (menu *Menu) DeleteMenu(c *gin.Context){
    id:= c.Param("id")
	row:=menuService.DeleteMenu(id)
	c.JSON(200,gin.H{"row":row})
}