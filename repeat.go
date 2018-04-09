package slackbot

import (
	"context"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// TODO: Seems like using the same method as Delete would work here. 
// I just incorporate shouldTrack func(slack.Message) bool and go through each message that way 






// The EventStore interface is used to read/write slack.RTMEvents to persistent storage
type EventStore interface {
	ReadEvents() (events map[string]slack.RTMEvent, err error)
	WriteEvents(events map[string]slack.RTMEvent) error
}

// The InMemoryEventStore type is an adapter to allow the use of a map[string]slack.RTMEvent as an EventStore.
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

// NewRepeatBehavior creates a behavior that stores the last message event from each channel.
// Events are only stored if shouldTrack returns true.
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

// NewRepeatCommand returns a cli.Command that sends the last event sent on the specified slack channel to ch.
// It is recommended you pass in a *slack.RTM.IncomingEvents channel as ch.
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
				return fmt.Errorf("Could not find latest command on this channel")
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
