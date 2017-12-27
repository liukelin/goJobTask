/**
 * redis
 * type list
 * message queue
 *
 * 必要函数
 * NewClient 	创建连接
 * Push_data	push数据
 * Pop_data		消费数据
 */
package redisClass

import (
	"fmt"
	"github.com/go-redis/redis"
	// "reflect"
	// "goJobTask/lib"
	"strconv"
	"time"
)

/**
 * 接口方法
 */
type MqClass struct {
	// redis
	Redis_host string
	Redis_port int
	Redis_db   int
	Redis_pass string
	// RedisClient *redis.Client
	// Connect()
	// push
	// Push_data(string, string) (bool, error)
	// pop
	// Pop_data(string, interface{})
	// 回调函数

}

// 回调函数
// type CallbackFunc func(string) bool

var RedisClient *redis.Client

/**
 * [NewClient 创建连接]
 * @param {[type]} ) (*redis.Client, error [description]
 */
func (MqClass *MqClass) Connect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     MqClass.Redis_host + ":" + strconv.Itoa(MqClass.Redis_port),
		Password: MqClass.Redis_pass, // no password set
		DB:       MqClass.Redis_db,   // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println("redis conn:", pong, err)
	// Output: PONG <nil>
	if err != nil {
		failOnError(err, "redis connect error.")
	}
	// return err
}

/**
 * [pop_data push数据]
 * data json string
 * @return {[type]} [description]
 */
func (MqClass *MqClass) Push_data(queueName string, data string) (bool, error) {
	err := RedisClient.RPush(queueName, data).Err()
	MqClass.Check_redis_conn(err)
	return true, err
}

/**
 * [pop_data 监听消费数据，回调]
 * @return {[type]} [description]
 */
func (MqClass *MqClass) Pop_data(queueName string, Callback func(string) bool) {

	MqClass.Connect()
	for {
		d, err := RedisClient.LPop(queueName).Result()

		MqClass.Check_redis_conn(err)

		if err == nil {
			// 回调执行
			ack := Callback(d)
			if !ack {
				fmt.Println("receiver 数据处理失败，将要重试")
			}
		} else {
			time.Sleep(1e9) // sleep one second
		}
		fmt.Println("end.\r\n")
	}
}

/**
 * 判断重连
 */
func (MqClass *MqClass) Check_redis_conn(err error) {
	if err != nil {
		if err != redis.Nil {
			MqClass.Connect()
		}
	}
	// return true
}

func failOnError(err error, msg string) {
	if err != nil {
		// log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
