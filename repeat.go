package slackbot

import (
	"context"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

type EventStore interface {
	ReadEvents() (map[string]slack.RTMEvent, error)
	WriteEvents(map[string]slack.RTMEvent) error
}

// The InMemoryEventStore type is an adapter to allow the use of ordinary map[string]slack.RTMEvent as EventStores.
type InMemoryEventStore map[string]slack.RTMEvent

// ReadEvents is used to satisfy the EventStore interface.
func (s InMemoryEventStore) ReadEvents() (map[string]slack.RTMEvent, error) {
	return s, nil
}

// WriteEvents is used to satisfy the EventStore interface.
func (s InMemoryEventStore) WriteEvents(events map[string]slack.RTMEvent) error {
	s = events
	return nil
}

// todo: NewRepeatBehavior(store, func(m slack.MessageEvent) bool { return strings.HasPrefix(m.Text, "iqvbot ") })

// NewRepeatBehavior creates a behavior that will track the last message event in each channel
func NewRepeatBehavior(store EventStore, shouldTrack func(slack.MessageEvent) bool) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		if !shouldTrack(*m) {
			return nil
		}

		events, err := store.ReadEvents()
		if err != nil {
			return err
		}

		events[m.Channel] = e
		if err := store.WriteEvents(events); err != nil {
			return err
		}

		return nil
	}
}

// NewRepeatCommand returns a cli.Command that repeats the last event sent on the specified slack channel.
// It is recommended you pass in your *slack.RTM.IncomingEvents channel to send repeating events.
func NewRepeatCommand(store EventStore, channelID string, ch chan<- slack.RTMEvent, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:  "repeat",
		Usage: "repeat the last command sent on this channel",
		Action: func(c *cli.Context) error {
			events, err := store.ReadEvents()
			if err != nil {
				return err
			}

			e, ok := events[channelID]
			if !ok {
				return fmt.Errorf("Could not find latest event on this channel")
			}

			ch <- e
			return nil
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
