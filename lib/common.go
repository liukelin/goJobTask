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
	"reflect"
	"strings"
	// "runtime"
	"strconv"
)

/**
 * [Fmt_message msg输出]
 * @param {[type]} msg ...interface{} [description]
 */
func Fmt_message(msg ...interface{}) {
	fmt.Println(msg)
}

// conf
type Configuration struct {
	http_port int

	// task最大数
	task_process int
	debug        bool

	// 监测频率
	sentinel_time int

	// 任务脚本路径
	task_script_file string

	// MQ类型 redis/rabbitmq
	mq string

	// redis
	redis_host string
	redis_port int
	redis_db   int

	// mysql
	mysql_host string
	mysql_port int
	mysql_user string
	mysql_pass string

	shell_cli   string
	php_cli     string
	java_cli    string
	python2_cli string
	python3_cli string

	// 其他配置
	conf_file string
	path      string
	server    string
	paramsStr string
	signKey   string
}

/**
 * [load_json_conf 加载获取配置]
 * @return {[type]} [description]
 */
func Load_conf(config *Configuration) (*Configuration, error) {

	file, _ := os.Open(config.conf_file)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		// fmt.Println("Error:", err)
		return config, err
	}
	// fmt.Println(conf.Path)

	paramsStr := ""
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

	value := reflect.ValueOf(conf)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		paramsStr = fmt.Sprintf("%s -%s=%s", paramsStr, i, value.Field(i))
	}

	config.paramsStr = paramsStr
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
 * [consu_data 消费数据]
 * @param  {[type]} d string)       (ack bool [description]
 * @return {[type]}      [description]
 */
func Consu_data(d string) bool {

	shell := ""
	// base64_decode
	decodeBytes, errBase := base64.StdEncoding.DecodeString(d)
	if errBase != nil {
		shell = d
	} else {
		shell = string(decodeBytes)
	}

	dMaps := loads_json(shell)
	v, ok := dMaps["shell"]
	if ok {
		fmt.Println("Run shell:", v, "\n")
		// 执行shell
		go func() {
			out, err := Run_shell(v)

			fmt.Println(out, "\n")
			if err != nil {
				fmt.Println(err, ".\n")
			} else {
				fmt.Println("success.\n")
			}
		}()
		// 异步处理(test)

		return true
	}

	return false
}

/**
 * 解析json
 */
// type d_json struct {
//  Shell string
//  Time  string
// }

func loads_json(jsonStr string) (maps map[string]string) {
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
			//  fmt.Println(i, u)
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
