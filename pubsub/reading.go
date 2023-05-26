package pubsub

import (
	"encoding/base64"
	"encoding/json"
	"github.com/solution-labs/tools/toolserror"
	"io"
	"net/http"
	"strings"
	"time"
)

type PubSubMessage struct {
	Attributes   map[string]string
	Data         string
	MessageID    string
	PublishTime  time.Time
	Subscription string
}

// ReadMessageFromPost takes the HTTP request feed and returns PubSubMessage struct
// converts time string to time object
func ReadMessageFromPost(r *http.Request) (message PubSubMessage, err error) {

	body, _ := io.ReadAll(r.Body)

	if len(body) == 0 {
		return message, toolserror.Wrap("pubsub:ReadMessageFromPost:[Missing Body]", err)
	}

	return ReadMessageFromByte(body)

}

// ReadMessageFromString takes string feed and returns PubSubMessage struct
func ReadMessageFromString(body string) (message PubSubMessage, err error) {
	return ReadMessageFromByte([]byte(body))
}

// ReadMessageFromByte takes []byte feed and returns PubSubMessage struct
func ReadMessageFromByte(body []byte) (message PubSubMessage, err error) {

	var msg = struct {
		Message struct {
			Attributes  map[string]string `json:"attributes"`
			Data        string            `json:"data"`
			MessageID   string            `json:"message_id"`
			PublishTime string            `json:"publish_time"`
		} `json:"message"`
		Subscription string `json:"subscription"`
	}{}

	err = json.Unmarshal(body, &msg)

	if err != nil {
		return message, toolserror.Wrap("pubsub:ReadMessage:", err)
	}

	b4d, err := base64.StdEncoding.DecodeString(msg.Message.Data)
	if err != nil {
		return message, toolserror.Wrap("pubsub:ReadMessage:decode", err)
	}
	message.Data = string(b4d)
	message.Attributes = msg.Message.Attributes
	message.MessageID = msg.Message.MessageID
	message.Subscription = msg.Subscription

	//2020-04-27T10:59:01.995Z = 2020-04-27 10:59:01
	msg.Message.PublishTime = strings.Replace(msg.Message.PublishTime, "T", " ", -1)[:19]

	message.PublishTime, err = time.Parse("2006-01-02 15:04:05", msg.Message.PublishTime)
	if err != nil {
		toolserror.Warning("pubsub:ReadMessage:PublishTime", err)
	}

	return message, nil

}
