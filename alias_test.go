package slackbot

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestAliasBehavior(t *testing.T) {
	store := InMemoryKeyValStore{
		"foo": "bar",
		"gif": "gif --random",
		"x":   "delete",
	}

	cases := map[string]struct {
		B      Behavior
		Event  slack.RTMEvent
		Assert func(t *testing.T, e slack.RTMEvent)
	}{
		"non-message event": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event:  slack.RTMEvent{},
			Assert: func(t *testing.T, e slack.RTMEvent) {},
		},
		"empty message event": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event: NewMessageRTMEvent(""),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"foo replaced with bar": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event: NewMessageRTMEvent("slackbot foo"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot bar", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"gif replaced with gif --random": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event: NewMessageRTMEvent("slackbot gif cars"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot gif --random cars", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"do not alias when shouldProcess returns false": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return false
			}),
			Event: NewMessageRTMEvent("slackbot foo"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot foo", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"do not alias when there is less than 2 args": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event: NewMessageRTMEvent("foo"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "foo", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"only alias the 2nd arg": {
			B: NewAliasBehavior(store, func(m *slack.MessageEvent) bool {
				return true
			}),
			Event: NewMessageRTMEvent("slackbot x example x y z"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot delete example x y z", e.Data.(*slack.MessageEvent).Text)
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := c.B(c.Event); err != nil {
				t.Fatal(err)
			}

			c.Assert(t, c.Event)
		})
	}
}

func TestAliasCommandList(t *testing.T) {
	store := InMemoryKeyValStore{
		"key0": "val0",
		"key1": "val1",
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
	store := InMemoryKeyValStore{
		"key0": "val0",
		"key1": "val1",
	}

	cmd := NewAliasCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias rm key0", " ")); err != nil {
		t.Fatal(err)
	}

	expected := InMemoryKeyValStore{
		"key1": "val1",
	}

	assert.Equal(t, expected, store)
}

func TestAliasCommandRemoveUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing KEY argument": strings.Split("slackbot alias rm", " "),
		"alias does not exist": strings.Split("slackbot alias rm key", " "),
	}

	app := NewTestApp(NewAliasCommand(InMemoryKeyValStore{}, ioutil.Discard))
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}

func TestAliasCommandAdd(t *testing.T) {
	store := InMemoryKeyValStore{
		"key0": "val0",
	}

	cases := map[string][]string{
		"add new entry":            strings.Split("slackbot alias add key1 val1", " "),
		"overwrite existing entry": strings.Split("slackbot alias add --force key0 updated", " "),
	}

	app := NewTestApp(NewAliasCommand(store, ioutil.Discard))
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			if err := app.Run(args); err != nil {
				t.Fatal(err)
			}
		})
	}

	expected := InMemoryKeyValStore{
		"key0": "updated",
		"key1": "val1",
	}

	assert.Equal(t, expected, store)
}

func TestAliasCommandAddUserInputErrors(t *testing.T) {
	store := InMemoryKeyValStore{
		"key": "val",
	}

	cases := map[string][]string{
		"missing KEY argument":      strings.Split("slackbot alias add", " "),
		"missing VAL argument":      strings.Split("slackbot alias add key", " "),
		"duplicate KEY w/o --force": strings.Split("slackbot alias add key val", " "),
	}

	app := NewTestApp(NewAliasCommand(store, ioutil.Discard))
	for key, args := range cases {
		t.Run(key, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
