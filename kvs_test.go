package slackbot

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValCommandList(t *testing.T) {
	store := InMemoryKeyValStore{
		"key0": "val0",
		"key1": "val1",
	}

	w := bytes.NewBuffer(nil)
	cmd := NewKVSCommand(store, w)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot kvs ls", " ")); err != nil {
		t.Fatal(err)
	}

	output := w.String()
	for key, val := range store {
		assert.Contains(t, output, key)
		assert.Contains(t, output, val)
	}
}

func TestKeyValCommandRemove(t *testing.T) {
	store := InMemoryKeyValStore{
		"key0": "val0",
		"key1": "val1",
	}

	cmd := NewKVSCommand(store, ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot kvs rm key0", " ")); err != nil {
		t.Fatal(err)
	}

	expected := InMemoryKeyValStore{
		"key1": "val1",
	}

	assert.Equal(t, expected, store)
}

func TestKeyValCommandRemoveUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing KEY argument": strings.Split("slackbot kvs rm", " "),
		"kvs does not exist":   strings.Split("slackbot kvs rm key", " "),
	}

	cmd := NewKVSCommand(InMemoryKeyValStore{}, ioutil.Discard)
	app := NewTestApp(cmd)
	for key, args := range cases {
		t.Run(key, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}

func TestKeyValCommandAdd(t *testing.T) {
	store := InMemoryKeyValStore{
		"key0": "val0",
	}

	cases := map[string][]string{
		"add new entry":            strings.Split("slackbot kvs add key1 val1", " "),
		"overwrite existing entry": strings.Split("slackbot kvs add --force key0 updated", " "),
	}

	app := NewTestApp(NewKVSCommand(store, ioutil.Discard))
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

func TestKeyValCommandAddUserInputErrors(t *testing.T) {
	store := InMemoryKeyValStore{
		"key": "val",
	}

	cases := map[string][]string{
		"missing KEY argument":      strings.Split("slackbot kvs add", " "),
		"missing VAL argument":      strings.Split("slackbot kvs add key", " "),
		"duplicate KEY w/o --force": strings.Split("slackbot kvs add key val", " "),
	}

	cmd := NewKVSCommand(store, ioutil.Discard)
	app := NewTestApp(cmd)
	for key, args := range cases {
		t.Run(key, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
