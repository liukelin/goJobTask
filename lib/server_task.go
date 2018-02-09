/**
 * server cli
 */
package lib

import (
	"fmt"
	// "github.com/go-redis/redis"
	// "io/ioutil"
	// "encoding/base64"
	// "encoding/json"
	// "os/exec"
	// "reflect"
	// "runtime"
	// "strconv"
	// "time"
	"goJobTask/rabbitClass"
	"goJobTask/redisClass"
)

// 睡眠等等时间
// var sleep_ int = 2

/**
 * 启动所有服务 server 入口
 * @return {[type]} [description]
 */
func Server_task(config *Configuration) {

	// 按CPU核数 设置并行数量
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// 获取列表 启动所有任务go

	queueName := "test_task"
	ret, _ := Start_task(config, queueName)

	fmt.Println("start Listening queueName.", queueName, ret, "\n")

	// CommonFunc := new(CommonFunc)

}

/**
 * 启动 消费服务
 * @param {[type]} Config *Configuration [description]
 */
func Start_task(Config *Configuration, queueName string) (bool, error) {

	var mqFunc &MqFunc
	// 启动服务
	if Config.Mq == "redis" {
		mqFunc = &redisClass.MqClass{Config.Redis_host, Config.Redis_port, Config.Redis_db, Config.Redis_pass}
	} else if Config.Mq == "rabbitmq" {
		mqFunc = &rabbitClass.MqClass{Config.Amqp}
		
	} else {
		mqFunc = &redisClass.MqClass{Config.Redis_host, Config.Redis_port, Config.Redis_db, Config.Redis_pass}
	}
	mqFunc.Pop_data(queueName, Consu_data)

	return true, nil
}
