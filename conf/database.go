/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package conf

const (
	//master conf
	DB_MASTER_HOST = "youraddress"
	DB_MASTER_port = "3306"
	DB_MASTER_CHARSET = "utf8"
	DB_MASTER_COLLATION = "utf8_general_ci"
	DB_MASTER_PREFIX="food_"
	DB_MASTER_USERNAME = "test"
	DB_MASTER_PASSWORD = "test"
	DB_MASTER_DATABASE = "test"
	//slave conf
	DB_SLAVE_HOST = "youraddress"
	DB_SLAVE_port = "3306"
	DB_SLAVE_CHARSET = "utf8"
	DB_SLAVE_COLLATION = "utf8_general_ci"
	DB_SLAVE_USERNAME = "test"
	DB_SLAVE_PASSWORD = "test"
	DB_SLAVE_DATABASE = "test"
)