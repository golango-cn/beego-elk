### beego-elk
Beego Elk extensions

### install
go get github.com/golango-cn/beego-elk

### Usage
```golang

import (
    "encoding/json"
    "github.com/astaxie/beego/logs"
    _ "github.com/golango-cn/beego-elk"
)

func TestBeegoElk() {

    logs.SetLogger(logs.AdapterConsole)
    logs.SetLogger("esLogger", `{"dsn":"http://192.168.31.230:9200/","level":7}`)
    
    b, _ := json.Marshal(map[string]interface{}{
        "requ":   "Test for Request",
        "resp":   "Test for Response",
        "others": 12345,
    })
    
    //If it is in JSON format, it is automatically converted to columns in ELK
    logs.Info(string(b))
    logs.Error(string(b))
    
    logs.Info("Hello world")

}

```

