package slackbot

import (
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
	"github.com/zpatrick/slackbot/mock_slack"
)

func TestRepeatCommandWithDefaults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_slack.NewMockSlackClient(ctrl)
	params := &slack.GetConversationHistoryParameters{
		ChannelID: "channel_id",
		Limit:     100,
	}

	history := &slack.GetConversationHistoryResponse{
		Messages: []slack.Message{
			{Msg: slack.Msg{User: "user_id", Text: "foo"}},
			{Msg: slack.Msg{User: "bot_id", Text: "bar"}},
			{Msg: slack.Msg{User: "user_id", Text: "baz"}},
		},
	}

	client.EXPECT().
		GetConversationHistory(params).
		Return(history, nil)

	eventc := make(chan slack.RTMEvent)
	go func() {
		select {
		case e := <-eventc:
			assert.Equal(t, "message", e.Type)
			assert.IsType(t, &slack.MessageEvent{}, e.Data)
			assert.Equal(t, "channel_id", e.Data.(*slack.MessageEvent).Channel)
			assert.Equal(t, "user_id", e.Data.(*slack.MessageEvent).User)
			assert.Equal(t, "foo", e.Data.(*slack.MessageEvent).Text)
		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}()

	isCommand := func(m slack.Message) bool {
		return m.User != "bot_id"
	}

	cmd := NewRepeatCommand(client, "channel_id", eventc, isCommand)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot repeat", " ")); err != nil {
		t.Fatal(err)
	}
}

func TestRepeatCommandWithLimitFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_slack.NewMockSlackClient(ctrl)
	params := &slack.GetConversationHistoryParameters{
		Limit: 500,
	}

	client.EXPECT().
		GetConversationHistory(params).
		Return(&slack.GetConversationHistoryResponse{}, nil)

	cmd := NewRepeatCommand(client, "", nil, nil)
	NewTestApp(cmd).Run(strings.Split("slackbot repeat --limit 500", " "))
}
