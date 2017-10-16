/**
 * liukelin
 *
 */
package main

import (
	"flag"
	"fmt"
	// "log"
	"os"
	// "github.com/go-redis/redis"
	lib "goJobTask/lib"
	// "strings"
)

var Debug *bool = flag.Bool("debug", true, "Are you debug?")
var Server *string = flag.String("server", "", "服务类型 (无/http/cli) ?")
var Port *string = flag.String("port", "8888", "http服务端口，8888")
var Process *string = flag.String("process", "1", "process member ?")
var Redishost *string = flag.String("redishost", "localhost:6379", "redis ip:端口 ?")
var Redispass *string = flag.String("redispass", "", "redis密码 ?")
var Redisdb *string = flag.String("redisdb", "0", "redis db ?")

var Args = os.Args //获取用户输入的所有参数
var Params = make(map[string]string)

/**
 * main
 * @return {[type]} [description]
 * -debug xx -server xx -port xx -process xx
 * -debug
 * -help
 */
func main() {
	flag.Parse()

	fmt.Println("Run:", *Server, "\n")

	Params["server"] = *Server
	Params["port"] = *Port
	Params["process"] = *Process
	Params["redishost"] = *Redishost
	Params["redispass"] = ""
	Params["redisdb"] = "0"

	fmt.Println(Params, "\n")

	switch *Server {

	case "http":
		/**
		 * 启动http server
		 */
		lib.Server_http(Params)

	case "cli":
		/**
		 * 启动脚本服务
		 */
		lib.Server_cli(Params)

	default:
		/**
		 * 启动所有服务
		 */

		// path, err_ := lib.Get_current_path()
		path := Args[0]
		// fmt.Println(path, err_, "\n")
		// fmt.Println(Args, "\n")

		Params["server"] = ""
		Params = lib.Load_json_conf(Params)

		run_http := fmt.Sprintf("%s %s -server=http", path, Params["paramsStr"])
		run_cli := fmt.Sprintf("%s %s -server=cli", path, Params["paramsStr"])

		out_c1 := make(chan string)
		out_c2 := make(chan string)
		err_c1 := make(chan error)
		err_c2 := make(chan error)

		go func() {
			out, err := lib.Run_shell(run_http)
			out_c1 <- out
			err_c1 <- err
		}()

		go func() {
			out, err := lib.Run_shell(run_cli)
			out_c2 <- out
			err_c2 <- err
		}()

		fmt.Println("Run http:", run_http, "\n", <-out_c1, <-err_c1, ".\n")
		fmt.Println("Run cli:", run_cli, "\n", <-out_c2, <-err_c2, ".\n")
	}

	if *Debug == true {
		fmt.Println("your debug is:", *Debug, "\n")
		// fmt.Println("redis:", rClient, "\n")
		// out, err := lib.Run_shell("ls")
		// fmt.Println(out, "\n", err, "\n")
	}
}
