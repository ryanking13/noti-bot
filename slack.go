package main

import "github.com/nlopes/slack"

func post(url string, msg *slack.WebhookMessage) {
	slack.PostWebhook(url, msg)
}
