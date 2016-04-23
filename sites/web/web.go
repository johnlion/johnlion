/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月11日 下午2:07
 *********************************************/

package web

import (
	"flag"
	"fmt"
	"github.com/johnlion/sites/config"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

var (
	ip     *string
	port   *int
	addr   string
	target *string //远程服务器
)

//获取外部参数
func Flag() {
	flag.String("b ******************************************** only for web ******************************************** -b", "", "")
	// web服务器IP与端口号
	ip = flag.String("b_ip", "0.0.0.0", "   <Web Server IP>")
	port = flag.Int("b_port", 80, "   <Web Server Port>")
	target = flag.String("target", "https://www.baidu.com", "target URL for reverse proxy")

}

func SetFlag() {
	ip = flag.String("b_ip", config.IP, "   <Web Server IP>")
	port = flag.Int("b_port", config.PORT, "   <Web Server Port>")
	target = flag.String("target", "http://test.jxapple2015.com", "target URL for reverse proxy")

	flag.Parse()
}

func Run() {

	// 设置CPU核心数量
	runtime.GOMAXPROCS(runtime.NumCPU())

	// web服务器地址
	addr = *ip + ":" + strconv.Itoa(*port)
	fmt.Println("[Local Server][", addr, "] start ...")
	fmt.Printf("%s\n", *target)

	// 预绑定路由
	Router()

	log.Fatal(http.ListenAndServe(addr, nil))
}
