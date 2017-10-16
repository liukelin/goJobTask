/**
 * run shell
 */
package lib

import (
	// "fmt"
	// "io/ioutil"
	// "encoding/base64"
	// "encoding/json"
	"os/exec"
	// "reflect"
	// "runtime"
	// "strconv"
	// "time"
)

/**
 * 执行shell命令
 */
func Run_shell(shell string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", shell)
	out_, err := cmd.Output()
	// out_, err := cmd.CombinedOutput()
	if err != nil {
		return string(out_), err
		// panic(err.Error())
	}
	if err := cmd.Start(); err != nil {
		return string(out_), err
		// panic(err.Error())
	}
	if err := cmd.Wait(); err != nil {
		// panic(err.Error())
	}
	return string(out_), err
}
