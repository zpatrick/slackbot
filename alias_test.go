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

type TestAliasStore map[string]string

func (s TestAliasStore) ReadAliases() (map[string]string, error) {
	return s, nil
}

func (s TestAliasStore) WriteAliases(aliases map[string]string) error {
	s = aliases
	return nil
}

func TestAliasBehavior(t *testing.T) {
	store := TestAliasStore{
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
			Event:  slack.RTMEvent{Data: &slack.MessageEvent{}},
			Assert: func(t *testing.T, e slack.RTMEvent) {},
		},
		"foo replaced with bar": {
			Event: slack.RTMEvent{Data: &slack.MessageEvent{
				Msg: slack.Msg{Text: "foo"},
			}},
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "bar", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"alias with flag": {
			Event: slack.RTMEvent{Data: &slack.MessageEvent{
				Msg: slack.Msg{Text: "cmd arg0"},
			}},
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
	store := TestAliasStore{
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
	store := TestAliasStore{
		"name0": "value0",
		"name1": "value1",
	}

	cmd := NewAliasCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias rm name0", " ")); err != nil {
		t.Fatal(err)
	}

	expected := TestAliasStore{
		"name1": "value1",
	}

	assert.Equal(t, expected, store)
}

func TestAliasCommandRemoveUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing NAME argument": strings.Split("slackbot alias rm", " "),
		"alias does not exist":  strings.Split("slackbot alias rm name", " "),
	}

	cmd := NewAliasCommand(TestAliasStore{}, ioutil.Discard)
	app := NewTestApp(cmd)
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			if err, ok := app.Run(args).(*UserInputError); !ok {
				t.Fatalf("Error %#v is not *UserInputError!", err)
			}
		})
	}
}

func TestAliasCommandSet(t *testing.T) {
	store := TestAliasStore{
		"name0": "value0",
		"name1": "value1",
	}

	cmd := NewAliasCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot alias set name0 updated", " ")); err != nil {
		t.Fatal(err)
	}

	expected := TestAliasStore{
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

	cmd := NewAliasCommand(TestAliasStore{}, ioutil.Discard)
	app := NewTestApp(cmd)
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			if err, ok := app.Run(args).(*UserInputError); !ok {
				t.Fatalf("Error %#v is not *UserInputError!", err)
			}
		})
	}
}
