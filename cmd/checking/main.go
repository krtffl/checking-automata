package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"

	"github.com/krtffl/checking-automata/internal/config"
)

var configPath = flag.String("config", "config/config.yaml", "config file path")

const (
	CHECKING_BUTTON      = `//button[contains(., 'Entrar')]`
	CHECKING_TYPE_BUTTON = `//button[contains(., 'Teletrabajo')]`
)

func main() {
	flag.Parse()

	cfg := config.Load(viper.New(), *configPath)

	// B R O W S E R    S E T  U P
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(
		ctx,
		time.Duration(cfg.Browser.Timeout)*time.Second,
	)
	defer cancel()

	// M A I L G U N    S E T  U P
	mg := mailgun.NewMailgun(cfg.Mailgun.Domain, cfg.Mailgun.Key)

	// When you have an EU-domain, you must specify the endpoint:
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	var msgContent string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(cfg.Browser.Page),
		chromedp.WaitVisible(CHECKING_BUTTON, chromedp.BySearch),
		chromedp.Click(CHECKING_BUTTON, chromedp.BySearch),
		chromedp.WaitVisible(CHECKING_TYPE_BUTTON, chromedp.BySearch),
		chromedp.Click(CHECKING_TYPE_BUTTON, chromedp.BySearch),
	)
	if err != nil {
		log.Printf("[CheckingAutomata] - Failed to check in. %v", err)
	}

	msg := mg.NewMessage(
		cfg.Mailgun.From,
		cfg.Mailgun.Subject,
		msgContent,
		cfg.Mailgun.To,
	)

	_, _, err = mg.Send(ctx, msg)
	if err != nil {
		log.Printf("[CheckingAutomata] - Failed to send email. %v", err)
	}
}
