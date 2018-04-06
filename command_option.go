package slackbot

import "github.com/urfave/cli"

// A CommandOption manipulates a cli.Command
type CommandOption func(cmd cli.Command) cli.Command

// WithName overwrites the command.Name field
func WithName(name string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Name = name
		return cmd
	}
}

// WithShortName overwrites the command.ShortName field
func WithShortName(shortName string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.ShortName = shortName
		return cmd
	}
}

// WithAliases overwrites the command.Aliases field
func WithAliases(aliases []string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Aliases = aliases
		return cmd
	}
}

// WithUsage overwrites the command.Usage field
func WithUsage(usage string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Usage = usage
		return cmd
	}
}

// WithUsageText overwrites the command.UsageText field
func WithUsageText(usageText string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.UsageText = usageText
		return cmd
	}
}

// WithDescription overwrites the command.Description field
func WithDescription(description string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Description = description
		return cmd
	}
}

// WithArgsUsage overwrites the command.ArgsUsage field
func WithArgsUsage(argsUsage string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.ArgsUsage = argsUsage
		return cmd
	}
}

// WithCategory overwrites the command.Category field
func WithCategory(category string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Category = category
		return cmd
	}
}

// WithBashComplete overwrites the command.BashComplete field
func WithBashComplete(bashComplete cli.BashCompleteFunc) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.BashComplete = bashComplete
		return cmd
	}
}

// WithBefore overwrites the command.Before field
func WithBefore(before cli.BeforeFunc) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Before = before
		return cmd
	}
}

// WithAfter overwrites the command.After field
func WithAfter(after cli.AfterFunc) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.After = after
		return cmd
	}
}

// WithAction overwrites the command.Action field
func WithAction(action interface{}) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Action = action
		return cmd
	}
}

// WithOnUsageError overwrites the command.OnUsageError field
func WithOnUsageError(onUsageError cli.OnUsageErrorFunc) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.OnUsageError = onUsageError
		return cmd
	}
}

// WithSubcommands overwrites the command.Subcommands field
func WithSubcommands(subcommands cli.Commands) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Subcommands = subcommands
		return cmd
	}
}

// WithFlags overwrites the command.Flags field
func WithFlags(flags []cli.Flag) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Flags = flags
		return cmd
	}
}

// WithSkipFlagParsing overwrites the command.SkipFlagParsing field
func WithSkipFlagParsing(skipFlagParsing bool) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.SkipFlagParsing = skipFlagParsing
		return cmd
	}
}

// WithSkipArgReorder overwrites the command.SkipArgReorder field
func WithSkipArgReorder(skipArgReorder bool) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.SkipArgReorder = skipArgReorder
		return cmd
	}
}

// WithHideHelp overwrites the command.HideHelp field
func WithHideHelp(hideHelp bool) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.HideHelp = hideHelp
		return cmd
	}
}

// WithHidden overwrites the command.Hidden field
func WithHidden(hidden bool) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Hidden = hidden
		return cmd
	}
}

// WithUseShortOptionHandling overwrites the command.UseShortOptionHandling field
func WithUseShortOptionHandling(useShortOptionHandling bool) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.UseShortOptionHandling = useShortOptionHandling
		return cmd
	}
}

// WithHelpName overwrites the command.HelpName field
func WithHelpName(helpName string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.HelpName = helpName
		return cmd
	}
}

// WithCustomHelpTemplate overwrites the command.CustomHelpTemplate field
func WithCustomHelpTemplate(customHelpTemplate string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.CustomHelpTemplate = customHelpTemplate
		return cmd
	}
}
