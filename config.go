package main

import "os"

var GITHUB_ACCESS_TOKEN string = os.Getenv("GITHUB_ACCESS_TOKEN")
var GITHUB_REPO_OWNER string = os.Getenv("GITHUB_REPO_OWNER")
var GITHUB_REPO_NAME string = os.Getenv("GITHUB_REPO_NAME")
var GITHUB_ISSUE_ID string = os.Getenv("GITHUB_ISSUE_ID")
var SLACK_WEBHOOK_URL string = os.Getenv("SLACK_WEBHOOK_URL")
