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

var Conf_file *string = flag.String("confile", "conf.json", "配置文件")
var Server *string = flag.String("server", "all", "服务类型 (all/http/task/sentinel) ?")

var Args = os.Args //获取用户输入的所有参数
// var Config = make(map[string]string)
var Config *lib.Configuration

// var Config *lib.Configuration = new(lib.Configuration)

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

	// 创建抽象对象
	// Config = new(lib.Configuration)

	// path, err_ := lib.Get_current_path()
	// path := Args[0]
	// fmt.Println(path, err_, "\n")
	// fmt.Println(Args, "\n")
	Config.path = Args[0]
	Config.conf_file = *Conf_file
	Config.server = *Server

	Config, _ = lib.Load_conf(Config)

	fmt.Println(Config, "\n")

	switch *Server {

	case "http":
		/**
		 * 启动http server任务接收、后台管理web
		 */
		lib.Server_http(Config)

	case "task":
		/**
		 * 启动脚本服务主进程
		 */
		lib.Server_task(Config)

	case "sentinel":
		/**
		 * 用于监听
		 * 任务执行状态 业务 （供查询）
		 * 任务子进程的状态（根据pid）
		 *
		 */
	case "all":
		/**
		 * 启动所有服务
		 */

		// run_http := fmt.Sprintf("%s %s -server=http", path, Params["paramsStr"])
		// run_task := fmt.Sprintf("%s %s -server=task", path, Params["paramsStr"])

		// out_c1 := make(chan string)
		// out_c2 := make(chan string)
		err_c1 := make(chan error)
		err_c2 := make(chan error)

		go func() {
			Config.server = "http"
			lib.Server_http(Config)

			// 使用独立进程
			// _, err := lib.Run_shell(run_http)
			// out_c1 <- out
			err_c1 <- nil
		}()

		go func() {
			Config.server = "task"
			lib.Server_task(Config)

			// 使用独立进程
			// _, err := lib.Run_shell(run_task)
			// out_c2 <- out
			err_c2 <- nil
		}()

		fmt.Println("Run http:", "\n", <-err_c1, ".\n")
		fmt.Println("Run task:", "\n", <-err_c2, ".\n")
	}

	// if *Debug == true {
	// fmt.Println("your debug is:", *Debug, "\n")
	// fmt.Println("redis:", rClient, "\n")
	// out, err := lib.Run_shell("ls")
	// fmt.Println(out, "\n", err, "\n")
	// }
}
