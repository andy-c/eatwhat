/**
  @description:
  @author: angels lose their hair
  @date: 2020/12/10
  @version:
**/
package dbmodel

type Menu struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Pic string
	Deleted int
	Updatetime int
	Createtime int
	Createuser string
	Updateuser string
}
