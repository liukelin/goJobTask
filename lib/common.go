/**
 * common
 *
 */
package lib

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	// "reflect"
)

/**
 * [Fmt_message msg输出]
 * @param {[type]} msg ...interface{} [description]
 */
func Fmt_message(msg ...interface{}) {
	fmt.Println(msg)
}

/**
 * [load_json_conf 加载获取配置]
 * @return {[type]} [description]
 */
func Load_json_conf(Params map[string]string) map[string]string {

	paramsStr := ""
	for k, v := range Params {
		if v != "" {
			paramsStr = fmt.Sprintf("%s -%s=%s", paramsStr, k, v)
		}
	}
	Params["paramsStr"] = paramsStr
	return Params
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
