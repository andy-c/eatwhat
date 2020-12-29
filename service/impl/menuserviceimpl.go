/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package impl

import (
	"eatwhat/dao/read"
	"eatwhat/dao/wrirte"
	"eatwhat/model/answer"
	"eatwhat/model/form"
)

type Menu struct{
    X int
    Y int
}

var (
	menuRead *read.MenuList
	menuWrite     *wrirte.Menu
)
//get menu list
func (menu *Menu) GetMenuList(menuform *form.MenuList) []answer.Menu{
    menulist := menuRead.FetchMenuList(menuform)
	return menulist
}

//add menu
func (menu *Menu) AddMenu(menuform *form.MenuList) int64{
    rows:=menuWrite.InsertMenu(menuform)
    return rows
}
//delete menu
func (menu *Menu) DeleteMenu(id string) int64{
   row:=menuWrite.DeleteMenu(id)
   return row
}

//update menu
func (menu *Menu) UpdateMenu( id string,name string) int64{
	row:=menuWrite.UpdateMenu(id,name)
	return row
}