package main

import (
	"flag"
	"log"
	"time"

	hbot "github.com/whyrusleeping/hellabot"
	"github.com/whyrusleeping/hellabot/examples/commands/command"
	"github.com/whyrusleeping/hellabot/examples/commands/config"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

// Flags for passing arguments to the program
var configFile = flag.String("config", "config.toml", "path to config file")

// core holds the command environment (bot connection and db)
var core *command.Core

// cmdList holds our command list, which tells the bot what to respond to.
var cmdList *command.List

// Main method
func main() {
	// Parse flags, this is needed for the flag package to work.
	// See https://godoc.org/flag
	flag.Parse()
	// Read the TOML Config
	conf := config.FromFile(*configFile)
	// Validate the config to see it's not missing anything vital.
	config.ValidateConfig(conf)

	// Setup our options anonymous function.. This gets called on the hbot.Bot object internally, applying the options inside.
	options := func(bot *hbot.Bot) {
		bot.SSL = conf.SSL
		if conf.ServerPassword != "" {
			bot.Password = conf.ServerPassword
		}
		bot.Channels = conf.Channels
		bot.PingTimeout = 8760 * time.Hour
	}
	// Create a new instance of hbot.Bot
	bot, err := hbot.NewBot(conf.Server, conf.Nick, options)
	if err != nil {
		log.Fatal(err)
	}
	// Setup the command environment
	core = &command.Core{bot, &conf}
	// Add the command trigger (this is what triggers all command handling)
	bot.AddTrigger(CommandTrigger)
	// Set the default bot logger to stdout
	bot.Logger.SetHandler(log15.StdoutHandler)
	// Initialize the command list
	cmdList = &command.List{
		Prefix:   "!",
		Commands: make(map[string]command.Command),
	}
	// Add commands to handle
	cmdList.AddCommand(command.Command{
		Name:        "manifestgraph",
		Description: "It provides the manifest graph of a specific OpenShift release image. See the below link for more information.\nhttps://github.com/openshift/enhancements/blob/master/dev-guide/cluster-version-operator/user/reconciliation.md#manifest-graph",
		Usage:       "!manifestgraph 4.10.10",
		Run:         core.Manifestgraph,
	})

	// Start up bot (blocks until disconnect)
	bot.Run()
	log.Println("Bot shutting down.")
}

// CommandTrigger passes all incoming messages to the commandList parser.
var CommandTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	func(bot *hbot.Bot, m *hbot.Message) bool {
		cmdList.Process(bot, m)
		return false
	},
}
