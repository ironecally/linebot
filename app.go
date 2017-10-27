package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/line/line-bot-sdk-go/linebot"
)

var botClient *linebot.Client

func main() {
	var err error
	botClient, err = linebot.New("", "")
	if err != nil {
		log.Fatalf("Failed to create bot handler")
		return
	}

	router := httprouter.New()
	router.GET("/test", testHandler)
	http.ListenAndServe(":8080", router)
	fmt.Println("httprouter is on! やった！！")
}

func testHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("got : %s!\n", string(body))

	events, err := botClient.ParseRequest(r)
	if err != nil {
		responseAPI(w, err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Printf("user_id: %s, group_id: %s, room_id: %s", event.Source.UserID, event.Source.GroupID, event.Source.RoomID)
		}
	}

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
