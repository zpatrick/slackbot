deps: 
	go get github.com/golang/mock/mockgen
	go install github.com/golang/mock/mockgen

mocks:
	mockgen -package mock github.com/zpatrick/slackbot SlackClient > mock_slack/mock_slack_client.go

test:
	go test ./... -v

.PHONY: deps mocks build test
