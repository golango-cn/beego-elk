package elk

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestElk(t *testing.T) {

	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger("esLogger", `{"dsn":"http://192.168.31.230:9200/","level":7, "index":"logs-elk"}`)

	b, _ := json.Marshal(map[string]interface{}{
		"requ":   "Test for Request",
		"resp":   "Test for Response",
		"others": 12345,
	})

	//If JSON format, it is automatically converted to columns in ELK
	logs.Info(string(b))
	logs.Error(string(b))

	logs.Info("Hello world")

}
