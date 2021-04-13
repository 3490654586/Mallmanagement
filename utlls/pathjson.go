package utlls

import (
	"bufio"
	"encoding/json"
	"os"
)

var _cof *Conf
/**
 *项目配置
 */
type Conf struct {
	App      string `json:"app"`
	AppModel string `json:"app_model"`
	AppHost  string `json:"app_host"`
	AppPost  string `json:"app_post"`
	Mysqls   Mysql  `json:"mysql"`
}


/**
 *Mysql配置
 */
type Mysql struct {
	Databases string `json:"databases"`
	User      string `json:"user"`
	Passwad   string  `json:"passwad"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Db_Name   string `json:"db_name"`
}

func OpenConf(path string) *Conf{
	file,err := os.Open(path)

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_cof)

	if err!=nil {
		return nil
	}
	return _cof
}

func GetConf () *Conf {
	return _cof
}
