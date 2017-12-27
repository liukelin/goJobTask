/**
 * common
 *
 */
package lib

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	// "reflect"
	"strings"
	// "runtime"
	"bytes"
	"strconv"
	// "goJobTask/rabbitClass"
	// "goJobTask/redisClass"
)

// type CallbackFunc func(string) bool

/**
 * 队列类实现方法
 */
type MqFunc interface {
	// 连接
	Connect()
	// push
	Push_data(string, string) (bool, error)
	// pop
	Pop_data(string, func(string) bool)
}

/**
 * 公共接口方法
 */
type CommonFunc interface {
	// 消费消息回调方法
	Consu_data(string) bool
}

// conf
type Configuration struct {
	Http_port int

	// task最大数
	Task_process int
	Debug        bool

	// 监测频率
	Sentinel_time int

	// 任务脚本路径
	Task_script_file string

	// MQ类型 redis/rabbitmq
	Mq string

	// rabbit amqp
	Amqp string

	// redis
	Redis_host string
	Redis_port int
	Redis_db   int
	Redis_pass string

	// mysql
	Mysql_host string
	Mysql_port int
	Mysql_user string
	Mysql_pass string

	Shell_cli   string
	Php_cli     string
	Java_cli    string
	Python2_cli string
	Python3_cli string

	// 其他配置
	Conf_file string
	Path      string
	Server    string
	ParamsStr string
}

// 传输信息
type Parameter struct {
	// 其他
	Shell string
	// 参数 json
	Data string
	// 时间
	Time string
	// 路径
	Path string
	// 脚本类型（php/python/js）
	Language string
}

func init() {
	// MqFuncMap = make(map[string]MqFunc)
	// MqFuncMap["redis"] = &redis
}

/**
 * [load_json_conf 加载获取配置]
 * @return {[type]} [description]
 */
func Load_conf(config *Configuration) (*Configuration, error) {
	fmt.Println(config.Conf_file)

	file, _ := os.Open(config.Conf_file)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
		return config, err
	}
	// fmt.Println(conf.Path)

	// paramsStr := ""
	// for k, v := range config {
	// 	if v != "" {
	// 		paramsStr = fmt.Sprintf("%s -%s=%s", paramsStr, k, v)
	// 	}
	// }
	// t := reflect.TypeOf(conf)
	// v := reflect.ValueOf(conf)
	// for k := 0; k < t.NumFiled(); k++ {
	// 	// fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	// 	paramsStr = fmt.Sprintf("%s -%s=%s", paramsStr, t.Filed(k).Name, v.Field(k).Interface())
	// }

	// value := reflect.ValueOf(conf)
	// for i := 0; i < value.NumField(); i++ {
	// 	// fmt.Printf("Field %d: %v\n", i, value.Field(i))
	// 	paramsStr = fmt.Sprintf("%s -%d=%s", paramsStr, value.Filed(i).Name, value.Field(i))
	// }
	// config.ParamsStr = paramsStr

	config.Http_port = conf.Http_port
	config.Task_process = conf.Task_process
	config.Debug = conf.Debug
	config.Sentinel_time = conf.Sentinel_time
	config.Task_script_file = conf.Task_script_file
	config.Mq = conf.Mq

	config.Redis_host = conf.Redis_host
	config.Redis_port = conf.Redis_port
	config.Redis_db = conf.Redis_db
	config.Redis_pass = conf.Redis_pass

	config.Mysql_host = conf.Mysql_host
	config.Mysql_port = conf.Mysql_port
	config.Mysql_user = conf.Mysql_user
	config.Mysql_pass = conf.Mysql_pass

	config.Shell_cli = conf.Shell_cli
	config.Php_cli = conf.Php_cli
	config.Java_cli = conf.Java_cli
	config.Python2_cli = conf.Python2_cli
	config.Python3_cli = conf.Python3_cli
	return config, nil
}

/**
 * [Get_current_path 获取项目路径]
 * @param {[type]} ) (string, error [description]
 */
func Get_current_path() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

/**
 * [Check_err description]
 * @param {[type]} err error [description]
 */
func Check_err(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * [Fmt_message msg输出]
 * @param {[type]} msg ...interface{} [description]
 */
func Fmt_message(msg ...interface{}) {
	// fmt.Println(msg)
	fmt.Sprintf("%s", msg)
}

/**
 * [consu_data 消费数据（回调函数）]
 * 处理队列数据
 * @param  {[type]} d string json)
 * @return {[type]} (ack bool [description]
 */
func Consu_data(d string) bool {
	fmt.Println(d, ".\n")
	shell := ""
	// base64_decode
	decodeBytes, errBase := base64.StdEncoding.DecodeString(d)
	if errBase != nil {
		shell = d
	} else {
		shell = string(decodeBytes)
	}

	dMaps := Loads_json(shell)
	v, ok := dMaps["shell"]
	if ok {
		fmt.Println("Run shell:", v, "\n")
		// 执行shell
		// go func() {
		out, err := Run_shell(v)

		fmt.Println(out, "\n")
		if err != nil {
			fmt.Println(err, ".\n")
		} else {
			fmt.Println("success.\n")
		}
		// }()
		// 异步处理(test)

		return true
	}

	return false
}

/**
 f = map[string]interface{}{
    "Name": "Wednesday",
    "Age":  6,
    "Parents": []interface{}{
        "Gomez",
        "Morticia",
    },
    "Parents": {"sss"}interface{}{
        "Gomez",
        "Morticia",
    },
}
*/
func Loads_json(jsonStr string) (maps map[string]string) {
	// var s d_json
	// jsonStr := `{"shell":"ls -la","time":"2017-09-21"}`
	// json.Unmarshal([]byte(jsonStr), &s)

	maps_ := make(map[string]string)

	var s interface{}
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		return maps_
	}

	m := s.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			// fmt.Println(k, "is string", vv)
			maps_[k] = vv

		case int:
			// fmt.Println(k, "is int", vv)
			maps_[k] = strconv.Itoa(vv)

		case float64:
			// fmt.Println(k, "is float64", vv)
			maps_[k] = strconv.FormatFloat(vv, 'E', -1, 32)

		case []interface{}:
			// 数组字典
			// fmt.Println(k, "is an array:")
			// for i, u := range vv {
			// 	fmt.Println(i, u)
			// }
		case map[string]interface{}:
			// 字典字典

		default:
			// fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	// fmt.Println(maps_)
	return maps_
}

/**
 * byte 转 string
 */
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
