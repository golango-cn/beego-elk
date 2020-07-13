package elk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"net/url"
	"strings"
	"time"
)

// NewES return a LoggerInterface
func NewElkES() logs.Logger {
	cw := &elkLogger{
		Level: logs.LevelDebug,
	}
	return cw
}

type elkLogger struct {
	*elasticsearch.Client

	DSN   string `json:"dsn"`
	Level int    `json:"level"`
	Index string `json:"index"`
}

var levelPrefix = []string{"[M]", "[A]", "[C]", "[E]", "[W]", "[N]", "[I]", "[D]"}

func (el *elkLogger) Init(jsonconfig string) error {
	err := json.Unmarshal([]byte(jsonconfig), el)
	if err != nil {
		return err
	}
	if el.DSN == "" {
		return errors.New("empty dsn")
	} else if u, err := url.Parse(el.DSN); err != nil {
		return err
	} else if u.Path == "" {
		return errors.New("missing prefix")
	} else {
		conn, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{el.DSN},
		})
		if err != nil {
			return err
		}
		el.Client = conn
	}
	return nil
}

func (el *elkLogger) WriteMsg(when time.Time, msg string, level int) error {

	if level > el.Level {
		return nil
	}

	msg = el.originMsg(msg, level)

	var values map[string]interface{}
	err := json.Unmarshal([]byte(msg), &values)

	var body []byte
	if err != nil {
		idx := ElkLogDocument{
			Timestamp: when.Format(time.RFC3339),
			Msg:       msg,
			LogLevel: levelPrefix[level],
		}
		b, err := json.Marshal(idx)
		if err != nil {
			return err
		}
		body = b
	} else {
		values["timestamp"] = when.Format(time.RFC3339)
		values["log_level"] = levelPrefix[level]
		b, err := json.Marshal(values)
		if err != nil {
			return err
		}
		body = b
	}

	index := el.Index
	if len(el.Index) == 0 {
		index = fmt.Sprintf("%04d.%02d.%02d", when.Year(), when.Month(), when.Day())
	} else {
		index = fmt.Sprintf("%s.%04d.%02d.%02d", el.Index, when.Year(), when.Month(), when.Day())
	}

	req := esapi.IndexRequest{
		Index:        index,
		DocumentType: "logs",
		Body:         strings.NewReader(string(body)),
	}
	_, err = req.Do(context.Background(), el.Client)
	return err

}

func (e elkLogger) Destroy() {
}

func (e elkLogger) Flush() {
}

func (e elkLogger) originMsg(msg string, level int) string {
	if len(msg) == 0 {
		return msg
	}
	prefix_l := len(levelPrefix[level]) + 2
	if len(msg) < prefix_l {
		return msg
	}
	msg = msg[prefix_l:]
	return msg

}

type ElkLogDocument struct {
	Timestamp string `json:"timestamp"`
	Msg       string `json:"msg"`
	LogLevel  string `json:"log_level"`
}

func init() {
	logs.Register("esLogger", NewElkES)
}
