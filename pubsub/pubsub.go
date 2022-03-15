package pubsub

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/solution-labs/tools/base64"
	"io/ioutil"
	"log"
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

	body, _ := ioutil.ReadAll(r.Body)

	if len(body) == 0 {
		return message, errors.New("missing body")
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
		return message, fmt.Errorf("ReadMessageFromPost:", err)
	}

	message.Data, err = base64.Base64ToString(msg.Message.Data)
	if err != nil {
		log.Println("ReadMessage:0x01:", err)
	}
	message.Attributes = msg.Message.Attributes
	message.MessageID = msg.Message.MessageID
	message.Subscription = msg.Subscription

	//2020-04-27T10:59:01.995Z = 2020-04-27 10:59:01
	msg.Message.PublishTime = strings.Replace(msg.Message.PublishTime, "T", " ", -1)[:19]

	message.PublishTime, err = time.Parse("2006-01-02 15:04:05", msg.Message.PublishTime)
	if err != nil {
		log.Println("ReadMessage:0x02:", err)
		log.Println("PublishTime:", msg.Message.PublishTime)
	}

	return message, nil

}
