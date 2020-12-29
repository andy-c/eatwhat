/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package wrirte

import (
	"eatwhat/database"
	"eatwhat/model/dbmodel"
	"eatwhat/model/form"
	"time"
)

type Menu struct {

}

var (
	db *database.DB
)
//update
func (menu *Menu) UpdateMenu(id string,name string) (row int64) {
	result:=db.GetDB().Table("menu").Where("id =? AND deleted = 0 ",id).Updates(map[string]interface{}{
		"name":name,
		"updateuser":"test",
		"updatetime":time.Now(),
	})
	return result.RowsAffected
}
//insert
func (menu *Menu) InsertMenu(menuform *form.MenuList) (row int64){
	menuinfo:= dbmodel.Menu{Name:menuform.Name,Pic: menuform.Pic,Createtime: 123,Updatetime: time.Now().Second(),Createuser: "test",Updateuser: "test"}
	result := db.GetDB().Create(&menuinfo)
	return result.RowsAffected
}

//delete
func (menu *Menu) DeleteMenu(id string ) (row int64){
	result:= db.GetDB().Table("menu").Where("id = ?",id).Updates(map[string]interface{}{"updateuser":"test","updatetime":time.Now(),"deleted":1})
	return result.RowsAffected
}