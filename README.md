### beego-elk
Beego Elk extensions

### install
go get github.com/golango-cn/beego-elk

### Usage
```golang

logs.SetLogger(logs.AdapterConsole)
logs.SetLogger("esLogger", `{"dsn":"http://192.168.31.230:9200/","level":7, "index":"logs-elk"}`)

b, _ := json.Marshal(map[string]interface{}{
    "requ":   "Request to test",
    "resp":   "Response to test",
    "others": 12345,
})

//If it is in JSON format, it is automatically converted to columns in ELK
logs.Info(string(b))
logs.Error(string(b))

logs.Info("Hello world")
```

