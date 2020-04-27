package solutionlabs_tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// {"message":{"attributes":{"source":"parkon"}, "data":"eyJzaXRlaWQiOiI4Mjc4IiwidnJtIjoiVEVTVDQiLCJmcm9tIjoiMjAyMC0wNC0yNyAxMjowMCIsInRvIjoiMjAyMC0wNC0yNyAxMjowNSIsImxhYmVsIjoiIiwiYWN0aW9uIjoiQUREIiwiY2xpZW50aWQiOiIxNzAxOCIsInJlbW90ZV90a24iOiIxMzIiLCJ1c2VyaWQiOiIxIn0=", "messageId":"1149481834963349","message_id":"1149481834963349","publishTime":"2020-04-27T10:59:01.995Z","publish_time":"2020-04-27T10:59:01.995Z"},"subscription":"projects/paymypcn-backoffice/subscriptions/vrmlist_2_whitelist_api"}

type PubSubMessage struct {
	Attributes   map[string]string
	Data         string
	MessageID    string
	PublishTime  time.Time
	Subscription string
}

func ReadMessageFromPost(r *http.Request) (message PubSubMessage, err error) {

	var msg = struct {
		Message struct {
			Attributes  map[string]string `json:"attributes"`
			Data        string            `json:"data"`
			MessageID   string            `json:"message_id"`
			PublishTime string            `json:"publish_time"`
		} `json:"message"`
		Subscription string `json:"subscription"`
	}{}

	body, _ := ioutil.ReadAll(r.Body)

	if len(body) == 0 {
		return message, errors.New("missing body")
	}

	err = json.Unmarshal(body, &msg)

	if err != nil {
		return message, fmt.Errorf("ReadMessageFromPost:", err)
	}

	message.Data, err = Base64ToString(msg.Message.Data)
	if err != nil {
		log.Println("ReadMessageFromPost:0x01:", err)
	}
	message.Attributes = msg.Message.Attributes
	message.MessageID = msg.Message.MessageID
	message.Subscription = msg.Subscription

	//2020-04-27T10:59:01.995Z = 2020-04-27 10:59:01
	msg.Message.PublishTime = strings.Replace(msg.Message.PublishTime, "T", " ", -1)[:19]

	message.PublishTime, err = time.Parse("2001-01-02 15:04:05", msg.Message.PublishTime)
	if err != nil {
		log.Println("ReadMessageFromPost:0x02:", err)
	}

	return message, nil

}
