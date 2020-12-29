/**
  @description:
  @author: angels lose their hair
  @date: 2020/11/24
  @version:
**/
package log

import (
	"eatwhat/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"time"
)

type Log struct{
}

//init log
func (log *Log) InitLog() {
	for _,channel:= range conf.LogChannel {
		channelLog,err:= os.Create(conf.LOG_PATH+channel+".log")
		if err!=nil {
			fmt.Println( "can not create"+channel+"log,errinfo"+err.Error())
		}
		//set log writer
		if strings.Compare("info",channel) == 0 {
			gin.DefaultWriter = io.MultiWriter(os.Stdout,channelLog)
		}else if strings.Compare("error",channel) == 0 {
			gin.DefaultErrorWriter = io.Writer(channelLog)
		}else{
			// only two channels ,for ohter channel ,needs set up
			gin.DefaultWriter = io.MultiWriter(os.Stdout,channelLog)
		}
	}
}

//set log formate
func (log *Log) SetLogFormate( param gin.LogFormatterParams) string{
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
