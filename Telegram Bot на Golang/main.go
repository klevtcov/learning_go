package main

import (
	"flag"
	"github.com/klevtcov/tg-bot/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// token = flags.Get(token)

	tgClient = telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New(tgClient) - будет получать данный с тг

	// processor = processor.New(tgClient) будет отправлять данные в тг

	// consumer.Start(fetcher, processor)
}

// must - приставка к функциям, в которых мы не обрабатываем ошибки, и которые критичны
//
//	т.е. при ошибке в этой функции, всЯ программа не запустится
//	и не имеет смысла - т.к. без токена бот работать не будет.
//
// токен мы будем брать в виде параметра программы при запуске
func mustToken() string {
	// bot - tg-bot-token 'my token' - будем указывать токен в виде флага
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegramm bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
