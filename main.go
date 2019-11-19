package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/nlopes/slack"
	"github.com/thoas/go-funk"
)

type Target struct {
	code       string
	noticeType string
	price      int
}

func checkEnv() {
	if GITHUB_ACCESS_TOKEN == "" ||
		GITHUB_REPO_OWNER == "" ||
		GITHUB_REPO_NAME == "" ||
		GITHUB_ISSUE_ID == "" ||
		SLACK_WEBHOOK_URL == "" {
		panic("Parameters not given")
	}
}

func getTargets() []Target {
	id, err := strconv.Atoi(GITHUB_ISSUE_ID)
	if err != nil {
		panic(err)
	}
	targetStr, err := getIssue(GITHUB_REPO_OWNER, GITHUB_REPO_NAME, id, GITHUB_ACCESS_TOKEN)
	if err != nil {
		panic(err)
	}

	targets := strings.Fields(*targetStr)
	targetsFiltered := []Target{}
	for _, target := range targets {
		target = strings.TrimSpace(target)
		// ignore comment which starts with `!`
		if strings.HasPrefix(target, "!") {
			continue
		}
		targetInfo := strings.Split(target, "|")
		code := targetInfo[0]
		noticeType := targetInfo[1]
		price, _ := strconv.Atoi(targetInfo[2])
		targetsFiltered = append(targetsFiltered, Target{
			code:       code,
			noticeType: noticeType,
			price:      price,
		})
	}

	return targetsFiltered
}

func notice(msg string) {
	slackMsg := &slack.WebhookMessage{
		Text: msg,
	}

	post(SLACK_WEBHOOK_URL, slackMsg)
}

func checkTarget(target Target, infos []StockInfo) {
	for _, info := range infos {
		if info.code != target.code {
			continue
		}

		if target.noticeType == "UP" && target.price <= info.currentValue {
			notice(fmt.Sprintf(":arrow_up: %s %d [%f (%f%%)]",
				info.name, info.currentValue, info.changeAmount, info.changeRate))
		} else if target.noticeType == "DOWN" && target.price >= info.currentValue {
			notice(fmt.Sprintf(":arrow_down: %s %d [%f (%f%%)]",
				info.name, info.currentValue, info.changeAmount, info.changeRate))
		}
	}
}

func main() {
	checkEnv()
	targets := getTargets()
	targetCodes := funk.Map(targets, func(t Target) string {
		return t.code
	}).([]string)

	stockInfos, err := poll(targetCodes)
	if err != nil {
		panic(err)
	}

	var wait sync.WaitGroup
	wait.Add(len(targets))
	for _, target := range targets {
		go func() {
			defer wait.Done()
			checkTarget(target, stockInfos)
		}()
	}
	wait.Wait()
}
