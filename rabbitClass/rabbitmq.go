/**
 * rabbitmq
 * type Work queues
 * message queue
 *
 * 必要函数
 * NewClient    创建连接
 * Push_data    push数据
 * Pop_data     消费数据
 */
/**
 * https://segmentfault.com/a/1190000010516906
 */

package rabbitClass

import (
	// "bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	// "time"
	"bytes"
	// "goJobTask/lib"
)

/**
 * 方法
 */
type MqClass struct {
	Amqp string
	// Connect(string)
	// // push
	// Push_data(string, string) (bool, error)
	// // pop
	// Pop_data(string, interface{})
}

// 回调函数
// type CallbackFunc func(string) bool

var conn *amqp.Connection
var channel *amqp.Channel

// var queueDeclare *amqp.QueueDeclare

/**
 * [NewClient 创建连接]
 * 方法实现
 * @param {[type]} ) (*redis.Client, error [description]
 */
func (MqClass *MqClass) Connect() {
	var err error

	conn, err = amqp.Dial(MqClass.Amqp)

	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()
	/**
		queueDeclare, err = channel.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
	    **/
	failOnError(err, "Failed to declare a queue")
}

/**
 * [pop_data push数据]
 * data json string
 * @return {[type]} [description]
 */
func (MqClass *MqClass) Push_data(queueName string, body string) (bool, error) {

	if channel == nil {
		MqClass.Connect()
	}
	err := channel.Publish(
		"",        // exchange
		queueName, // queueDeclare.Name, // routing key
		false,     // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	return true, err
}

/**
 * [pop_data 消费数据]
 * @return {[type]} [description]
 */
func (MqClass *MqClass) Pop_data(queueName string, Callback func(string) bool) {

	if channel == nil {
		MqClass.Connect()
	}

	err := channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := channel.Consume(
		queueName, //q.Name, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {

		// 使用callback消费数据
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)

			body := BytesToString(&(msg.Body))

			// 当接收者消息处理失败的时候，
			// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
			// 通过重试可以成功的操作，那么这个时候是需要重试的
			// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
			/**
			            for !CommonFunc.Consu_data(body) {
							// log.Warnf("receiver 数据处理失败，将要重试")
							time.Sleep(1 * time.Second)
						}
			            **/

			// 不重试
			ret := Callback(*body)
			if ret {

			}

			// 确认收到本条消息, multiple必须为false
			msg.Ack(false)
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

/**
 * byte 转 string
 */
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
