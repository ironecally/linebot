package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bluele/slack"
	"github.com/julienschmidt/httprouter"
)

type slackParam struct {
	webhookURL string
	apiToken   string
}

var slackConfig slackParam
var slackBotAPI *slack.Slack

func initSlackBot() {
	slackConfig = slackParam{
		webhookURL: "https://hooks.slack.com/services/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",
		apiToken:   "xoxb-266936225671-d9x0eK7np7lPODSoGubksXux",
	}
	slackBotAPI = slack.New(slackConfig.apiToken)
	err := slackBotAPI.ChatPostMessage("jerry-err", "ready to serve", nil)
	if err != nil {
		panic("fail to init slackbot")
	}

	err = slackBotAPI.JoinChannel("jerry-err")
	if err != nil {
		panic("fail to join jerry-err")
	}
}

func sendToSlackViaWebhook(text string) error {

	payload := map[string]interface{}{
		"text": text,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Post(slackConfig.webhookURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("[panics] error on capturing error : %s \n", err.Error())
		return err
	}

	if resp.StatusCode >= 300 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[panics] error on capturing error : %s \n", err)
			return err
		}
		log.Printf("[panics] error on capturing error : %s \n", string(b))
	}
	return nil
}

func slackTestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := slackBotAPI.ChatPostMessage("jerry-squad", "Hi @here, i'm jerry bot, your personal jerry utility assistant,\ncurrently im still dumb, but believe me, i'm trying :sad:", nil)
	if err != nil {
		responseAPI(w, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	responseAPI(w, map[string]interface{}{"success": true})
	return
}
