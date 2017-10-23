/**
 * message queue
 */
package lib

import (
	// "fmt"
	"github.com/go-redis/redis"
	// "reflect"
)

func Client(addr string, password string, db int) (*redis.Client, error) {
	rClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := rClient.Ping().Result()
	// fmt.Println(pong, reflect.TypeOf(pong), err, reflect.TypeOf(err))
	// Output: PONG <nil>
	return rClient, err
}

/**
 * [pop_data 创建数据]
 * data json string
 * @return {[type]} [description]
 */
func Push_data(queue string, data string) (bool, error) {
	return false, nil
}

/**
 * [pop_data 消费数据]
 * @return {[type]} [description]
 */
func Pop_data(queue string) (string, error) {
	return "false", nil
}
