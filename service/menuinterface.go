/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package service

import (
	"eatwhat/model/answer"
	"eatwhat/model/form"
)

type Menu interface {
	GetMenuList(menuform *form.MenuList) []answer.Menu
	AddMenu(menuform *form.MenuList) ( int64 )
	DeleteMenu(id string ,name string) int64
	UpdateMenu(id string) int64
}