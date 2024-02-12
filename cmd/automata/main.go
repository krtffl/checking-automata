package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"

	"github.com/krtffl/checking-automata/internal/config"
)

var configPath = flag.String("config", "config/config.yaml", "config file path")

const (
	CHECKING_BUTTON = `//button[contains(., 'Entrar')]`
	REMOTE_BUTTON   = `//button[contains(., 'Teletrabajo')]`
)

func main() {
	flag.Parse()

	cfg := config.Load(viper.New(), *configPath)

	// B R O W S E R    S E T  U P
	// it needs to attach to a currently running app
	// otherwise two factor auth will be required to login
	allowCtx, cancel := chromedp.NewRemoteAllocator(
		context.Background(),
		cfg.Browser.Address,
	)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allowCtx,
		chromedp.WithLogf(log.Printf),
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
	// mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	var msgContent string
	err := chromedp.Run(
		ctx,

		// navigating to checking page
		chromedp.Navigate(cfg.Browser.Page),

		// should already be logged in and with
		// location permissions
		chromedp.WaitVisible(CHECKING_BUTTON, chromedp.BySearch),
		chromedp.Click(CHECKING_BUTTON, chromedp.BySearch),
		chromedp.WaitVisible(REMOTE_BUTTON, chromedp.BySearch),
		chromedp.Click(REMOTE_BUTTON, chromedp.BySearch),
	)
	if err != nil {
		log.Printf("[CheckingAutomata] - Failed to check in. %v", err)
		msgContent = fmt.Sprintf(
			`good morning
            there's been a problem checking in for today: %v.
             
            
            please make sure you handle it yourself.

            checkingautomata,`,
			err,
		)
	} else {
		msgContent = fmt.Sprintf(
			`
            good morning
            everything is in order.


            checkingautomata,`,
		)
	}

	msg := mg.NewMessage(
		cfg.Mailgun.From,
		cfg.Mailgun.Subject,
		msgContent,
		cfg.Mailgun.To,
	)

	m, id, err := mg.Send(ctx, msg)
	if err != nil {
		log.Printf("[CheckingAutomata] - Failed to send email. %v", err)
	}

	log.Printf("[CheckingAutomata] - Successfully checked in and notified. %s. %s", m, id)
}
