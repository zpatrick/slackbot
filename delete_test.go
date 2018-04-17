package slackbot

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nlopes/slack"
	"github.com/zpatrick/slackbot/mock_slack"
)

func TestDeleteCommandWithDefaults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_slack.NewMockSlackClient(ctrl)
	params := &slack.GetConversationHistoryParameters{
		ChannelID: "channel_id",
		Limit:     100,
	}

	history := &slack.GetConversationHistoryResponse{
		Messages: []slack.Message{
			{Msg: slack.Msg{User: "user_id", Timestamp: "t1"}},
			{Msg: slack.Msg{User: "bot_id", Timestamp: "t2"}},
			{Msg: slack.Msg{User: "bot_id", Timestamp: "t3"}},
			{Msg: slack.Msg{User: "user_id", Timestamp: "t4"}},
		},
	}

	client.EXPECT().
		GetConversationHistory(params).
		Return(history, nil)

	client.EXPECT().
		DeleteMessage("channel_id", "t2").
		Return("", "", nil)

	cmd := NewDeleteCommand(client, "bot_id", "channel_id")
	if err := NewTestApp(cmd).Run(strings.Split("slackbot delete", " ")); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteCommandWithLimitFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_slack.NewMockSlackClient(ctrl)
	params := &slack.GetConversationHistoryParameters{
		Limit: 500,
	}

	client.EXPECT().
		GetConversationHistory(params).
		Return(&slack.GetConversationHistoryResponse{}, nil)

	cmd := NewDeleteCommand(client, "", "")
	NewTestApp(cmd).Run(strings.Split("slackbot delete --limit 500", " "))
}
