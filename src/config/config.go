package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	config_data           sync.Map
	config_branch_name    = ""
	config_path_file_name = ""
)

func init() {
	// 配置文件 绝对位置
	if runtime.GOOS == "windows" {
		config_path_file_name = "F:/Golang/src/autoTransToken/config/config.ini" // windows
	} else {
		config_path_file_name = "./config/config.ini" // linux
	}
}

// 配置文件名 和 分支名称
func InitConfig(branch_name string) {
	config_branch_name = branch_name

	getNameConfigDatas(branch_name)
}

func GetConfigString(key string) string {
	if temp, ok := config_data.Load(key); ok {
		if val, ok := temp.(string); ok {
			return val
		}
	}
	return ""
}

func GetConfigBool(key string) (bool, error) {
	n := GetConfigString(key)
	if n == "" {
		return false, fmt.Errorf("Not Found!")
	}

	return strconv.ParseBool(n)
}

func GetConfigInt(key string) (int, error) {
	n := GetConfigString(key)
	if n == "" {
		return 0, fmt.Errorf("Not Found!")
	}

	return strconv.Atoi(n)
}

func ReLoadConfig() {
	getNameConfigDatas(config_branch_name)
}

func readFile() string {
	fi, err := os.Open(config_path_file_name)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func getNameConfigDatas(name string) {
	file_str := readFile()
	strs := strings.Split(file_str, "[")

	temp_str := ""
	l := len(name)
	for _, v := range strs {
		if len(v) > l &&
			strings.Compare(name, v[:len(name)]) == 0 {
			temp_str = v[l+1:]
			break
		}
	}
	for _, v := range strings.Split(temp_str, "\n") {
		if len(v) < 1 || v[0] == '#' {
			continue
		}

		// 删除注释
		n := strings.IndexByte(v, '#')
		if n != -1 {
			v = v[:n-1]
		}

		n = strings.IndexByte(v, '=')
		if n == -1 {
			continue
		}
		k := strings.Replace(v[:n], "\t", "", -1)
		val := strings.Replace(v[n+1:], "\t", "", -1)
		k = strings.Replace(k, " ", "", -1)
		val = strings.Replace(val, " ", "", -1)
		config_data.Store(k, val)
	}
}
