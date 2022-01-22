package conf

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tangwh1206/twitter/core"
	"github.com/tangwh1206/twitter/pkg/jsonx"
	"gopkg.in/yaml.v2"
)

var (
	setting *core.Setting
)

func GetSetting() *core.Setting {
	return setting
}

func Init() {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "prod"
	}
	confFilename := "config_" + env + ".yaml"
	confDir := os.Getenv("CONF_DIR")
	filepath := confDir + "/" + confFilename
	log.Printf("config filepath=%v\n", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("os.Open file fail, filename=%v, err=%v\n", filepath, err)
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("ioutil.ReadAll fail, err=%v\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(content, &setting)
	if err != nil {
		log.Printf("yaml.Unmarshal fail, content=%v, err=%v\n", content, err)
		panic(err)
	}
	log.Printf("conf=%v", jsonx.BeautifyDump(setting))
}
