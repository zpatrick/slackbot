package slackbot

import (
	"context"
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

// todo: test that fn is honored
// todo: test that the store gets populated
func TestRepeatBehavior(t *testing.T) {
	var called bool
	fn := func(slack.MessageEvent) bool {
		called = true
		return true
	}

	store := InMemoryEventStore{}
	b := NewRepeatBehavior(store, fn)
	e := NewMessageChannelRTMEvent("cid", "")
	if err := b(context.Background(), e); err != nil {
		t.Fatal(err)
	}

	assert.True(t, called)

	expected := InMemoryEventStore{
		"cid": e,
	}

	assert.Equal(t, expected, store)
}

func TestRepeatCommand(t *testing.T) {
	t.Skip("TODO")
}

func TestRepeatCommandUserInputErrors(t *testing.T) {
	t.Skip("TODO")
}
