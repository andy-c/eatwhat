/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package read

import (
	"eatwhat/database"
	"eatwhat/model/answer"
	"eatwhat/model/form"
)

var (
	db *database.DB
)

type MenuList struct {
}
//fetch menu list
func (menu *MenuList) FetchMenuList(menuform *form.MenuList) []answer.Menu {
	menulist:=[]answer.Menu{}
	db.GetDB().Model(&menulist).Where("deleted = 0").Find(&menulist)
	return menulist
}

