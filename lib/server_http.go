/**
 * http server
 */
package lib

import (
	"fmt"
	// "html/template"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	"github.com/go-redis/redis"
	// "reflect"
	"strconv"
	// "strings"
	"time"
)

var RConn0 *redis.Client
var Rerr0 error
var Params *Configuration

/**
 * web server 入口
 * @return {[type]} [description]
 */
func Server_http(params *Configuration) {

	Params = params

	http.HandleFunc("/", server_http_action)

	// strconv.Itoa(port)
	portStr := ":" + strconv.Itoa(params.http_port)

	fmt.Println("your portStr is:", portStr, "\n")

	// mux := http.NewServeMux()
	// err := http.ListenAndServe(portStr, mux)
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// go mux.Run()
	fmt.Println("RedisClient connection error.\n")
}

/**
 * [web_server_action 请求业务处理]
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func server_http_action(w http.ResponseWriter, r *http.Request) {
	// 解析参数, 默认是不会解析的
	r.ParseForm()

	// d := r.Form["d"]
	d := r.FormValue("d")
	key := r.FormValue("sign")

	if len(Params.signKey) > 0 && Params.signKey != key {

		fmt.Fprintf(w, "signKey error.")

	} else {

		fmt.Println(time.Now(), "body:", r.Form, "d:", d)
		fmt.Fprintf(w, "success.")
	}

}
