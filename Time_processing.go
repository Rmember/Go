package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix() //时间戳
	F_Time := time.Now().Format("2006-01-02 15:04:05") //格式化时间
	str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")//时间戳转str格式化时间
	the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	unix_time := the_time.Unix() //str格式化时间转时间戳

	fmt.Println(timestamp)
	fmt.Println(F_Time)
	fmt.Println(str_time)
	fmt.Println(unix_time)
}
