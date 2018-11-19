package conf

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Init 初始化
func Init(r string) {
	Root = r
}

// Root Root
var Root string

// ConfDir 配置路径
func ConfDir() string {
	return Root + "/conf"
}

// LoadJSON 工具方法，加载配置中的json文件
func LoadJSON(file string, data interface{}) error {
	bs, err := ioutil.ReadFile(ConfDir() + "/" + file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bs, data)
}

func LoadFile(file string) ([]byte, error) {
	if !filepath.IsAbs(file) {
		file = ConfDir() + "/" + file
	}

	return ioutil.ReadFile(file)
}
