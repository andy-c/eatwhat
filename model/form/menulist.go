/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/26
  @version:
**/
package form

type MenuList struct {
	Name string `form:"name" binding:"required"`
	Pic string `form:"pic" binding:"required"`
}
