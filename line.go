package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/line/line-bot-sdk-go/linebot"
)

var botClient *linebot.Client
var accessToken string
var channelSecret string

func initLine() {
	log.Println("initializing line")
	accessToken = os.Getenv("line_accessToken")
	if accessToken == "" {
		log.Fatalf("line_accessToken not found!")
	}
	channelSecret = os.Getenv("line_channelSecret")
	if channelSecret == "" {
		log.Fatalf("line_channelSecret not found!")
	}

	var err error
	botClient, err = linebot.New(channelSecret, accessToken)
	if err != nil {
		log.Fatalf("Failed to create bot handler")
		return
	}
}
func testLineHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseAPI(w, err.Error())
		return
	}

	if len(body) > 1 {
		log.Printf("got : %s!\n", string(body))
	}

	events, err := botClient.ParseRequest(r)
	if err != nil {
		responseAPI(w, err.Error())
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Printf("user_id: %s, group_id: %s, room_id: %s", event.Source.UserID, event.Source.GroupID, event.Source.RoomID)
		}
	}

	responseAPI(w, map[string]interface{}{"success": true})
	return
}

func responseAPI(w http.ResponseWriter, v interface{}) {
	data := map[string]interface{}{
		"data": v,
	}
	jsonVal, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte{})
	}
	w.Write(jsonVal)
	return
}
