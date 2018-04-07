package slackbot

import (
	"bytes"
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestAliasBehavior(t *testing.T) {
	store := InMemoryAliasStore{
		"foo": "bar",
		"cmd": "cmd --flag",
	}

	cases := map[string]struct {
		Event  slack.RTMEvent
		Assert func(t *testing.T, e slack.RTMEvent)
	}{
		"non-message event": {
			Event:  slack.RTMEvent{},
			Assert: func(t *testing.T, e slack.RTMEvent) {},
		},
		"empty message event": {
			Event: NewMessageRTMEvent(""),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"foo replaced with bar": {
			Event: NewMessageRTMEvent("foo"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "bar", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"alias with flag": {
			Event: NewMessageRTMEvent("cmd arg0"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "cmd --flag arg0", e.Data.(*slack.MessageEvent).Text)
			},
		},
	}

	b := NewAliasBehavior(store)
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := b(context.Background(), c.Event); err != nil {
				t.Fatal(err)
			}

			c.Assert(t, c.Event)
		})
	}
}

func TestAliasCommandList(t *testing.T) {
	store := InMemoryAliasStore{
		"name0": "value0",
		"name1": "value1",
	}

	w := bytes.NewBuffer(nil)
	cmd := NewAliasCommand(store, w)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias ls", " ")); err != nil {
		t.Fatal(err)
	}

	output := w.String()
	for k, v := range store {
		assert.Contains(t, output, k)
		assert.Contains(t, output, v)
	}
}

func TestAliasCommandRemove(t *testing.T) {
	store := InMemoryAliasStore{
		"name0": "value0",
		"name1": "value1",
	}

	cmd := NewAliasCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias rm name0", " ")); err != nil {
		t.Fatal(err)
	}

	expected := InMemoryAliasStore{
		"name1": "value1",
	}

	assert.Equal(t, expected, store)
}

func TestAliasCommandRemoveUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing NAME argument": strings.Split("slackbot alias rm", " "),
		"alias does not exist":  strings.Split("slackbot alias rm name", " "),
	}

	cmd := NewAliasCommand(InMemoryAliasStore{}, ioutil.Discard)
	app := NewTestApp(cmd)
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}

func TestAliasCommandSet(t *testing.T) {
	store := InMemoryAliasStore{
		"name0": "value0",
		"name1": "value1",
	}

	cmd := NewAliasCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias set name0 updated", " ")); err != nil {
		t.Fatal(err)
	}

	expected := InMemoryAliasStore{
		"name0": "updated",
		"name1": "value1",
	}

	assert.Equal(t, expected, store)
}

func TestAliasCommandSetUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing NAME argument":  strings.Split("slackbot alias set", " "),
		"missing VALUE argument": strings.Split("slackbot alias set name", " "),
	}

	cmd := NewAliasCommand(InMemoryAliasStore{}, ioutil.Discard)
	app := NewTestApp(cmd)
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
