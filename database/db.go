/**
  @description:db
  @author: angels lose their hair
  @date: 2020/11/24
  @version:v1
**/
package database

import (
	"eatwhat/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

type DB struct{
}

var dbConn *gorm.DB

//init db pool
func (db *DB) GetDB() *gorm.DB {
	master_dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&collation=%s",
		  conf.DB_MASTER_USERNAME,
		  conf.DB_MASTER_PASSWORD,
		  conf.DB_MASTER_HOST,
		  conf.DB_MASTER_port,
		  conf.DB_MASTER_DATABASE,
		  conf.DB_MASTER_CHARSET,
		  conf.DB_MASTER_COLLATION,
		)
	slave_dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&collation=%s",
		conf.DB_SLAVE_USERNAME,
		conf.DB_SLAVE_PASSWORD,
		conf.DB_SLAVE_HOST,
		conf.DB_SLAVE_port,
		conf.DB_SLAVE_DATABASE,
		conf.DB_SLAVE_CHARSET,
		conf.DB_SLAVE_COLLATION,
	)
	sqllog,err :=os.Create(conf.LOG_PATH+conf.SQLChannel+".log")
	if err!=nil {
		fmt.Println("create sql log error ,error info is "+err.Error())
	}
	dbConn,err:= gorm.Open(mysql.Open(master_dsn),&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: conf.DB_MASTER_PREFIX,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating:true,
		Logger: logger.New(
			log.New(sqllog, "\r\n",log.LstdFlags),
			logger.Config{
				SlowThreshold: 2*time.Second,
				Colorful: false,
				LogLevel: logger.Error,
			},
		),
	})
	if err != nil {
		log.Fatal("connect error:"+err.Error())
	}
	dbConn.Use( dbresolver.Register(dbresolver.Config{
		        Replicas: []gorm.Dialector{mysql.Open(slave_dsn)},
	         }).
		        SetConnMaxLifetime(30*time.Minute).
		        SetMaxIdleConns(20).
	            SetMaxOpenConns(50),
	)
	return dbConn
}
