/*
 * @Author: your name
 * @Date: 2020-10-01 16:27:12
 * @LastEditTime: 2020-10-01 16:28:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \undefinede:\SunRun-master\main.go
 */
package main

import (
	"flag"
)

var str = flag.String("s", "", "IMECode")

func main(){
	flag.Parse()
	justRun(*str, "2400")
}
