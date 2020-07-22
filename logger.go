package elk

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/golango-cn/beego-elk"
)

// EsLogger
func NewEsLogger(server, index string) (*logs.BeeLogger, error) {

	logger := logs.NewLogger()
	logger.Async(100)
	if len(server) > 0 {
		logger.SetLogger("esLogger", fmt.Sprintf(`{"dsn":"%s","level":7,"index":"%s"}`, server, index))
		return logger, nil
	} else {
		logger.SetLogger(logs.AdapterConsole)
		return logger, errors.New("未配置ES服务，使用Console记录日志")
	}

}
