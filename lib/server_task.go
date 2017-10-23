/**
 * server cli
 */
package lib

import (
	"fmt"
	"github.com/go-redis/redis"
	// "io/ioutil"
	// "encoding/base64"
	// "encoding/json"
	// "os/exec"
	// "reflect"
	// "runtime"
	// "strconv"
	"time"
)

var RConn *redis.Client
var Rerr error

// 睡眠等等时间
// var sleep_ int = 2

/**
 * 后台脚本server 入口
 * @return {[type]} [description]
 */
func Server_task(params *Configuration) {

	// 按CPU核数 设置并行数量
	// runtime.GOMAXPROCS(runtime.NumCPU())

	for {

		fmt.Println(time.Now(), ".\n")

		// 阻塞 2s
		// time.Sleep(time.Second * 2)
		// time.Sleep(time.Second)
		// time.Sleep(1000 * time.Millisecond)
		time.Sleep(1e9) // sleep one second

		// 非阻塞
		// time.After(time.Second + 10)
		continue
	}

}
