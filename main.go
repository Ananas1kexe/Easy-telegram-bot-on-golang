package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"

	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	botToken := "Your token here"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer func() {
		_, _ = bot.StopPoll(ctx, nil)
	}()

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)
		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Start"),
				tu.KeyboardButton("Help"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("Support"),
				tu.KeyboardButton("Discord"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("Back"),
			),
		)
		message := tu.Message(
			chatID,
			"hello im BotSnare on go",
		).WithReplyMarkup(keyboard)

		_, _ = bot.SendMessage(ctx, message)
		return nil
	}, th.CommandEqual("start"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		chatID := tu.ID(update.Message.Chat.ID)
		message := tu.Message(
			chatID,

			"Current time:\n"+
				currentTime,
		)
		_, _ = bot.SendMessage(ctx, message)
		return nil
	}, th.CommandEqual("time"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)
		message := tu.Message(
			chatID,

			"command /ping - for answer pong\n"+
				"command /start - for start\n"+
				"command /time - for current time",
		)
		_, _ = bot.SendMessage(ctx, message)
		return nil
	}, th.CommandEqual("help"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)
		message := tu.Message(
			chatID,
			"Pong",
		)
		_, _ = bot.SendMessage(ctx, message)
		return nil
	}, th.CommandEqual("ping"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := tu.ID(update.Message.Chat.ID)
		message := tu.Message(
			chatID,
			"idk",
		)
		_, _ = bot.SendMessage(ctx, message)
		return nil
	}, th.AnyMessage())

	bh.Start()
}
