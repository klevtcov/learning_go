package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/klevtcov/tg-bot/clients/telegram"
	event_consumer "github.com/klevtcov/tg-bot/consumer/event-consumer"
	"github.com/klevtcov/tg-bot/events/telegram"
	// "github.com/klevtcov/tg-bot/storage/files"
	"github.com/klevtcov/tg-bot/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "storage"
	sqliteStorage = "data/sqlite/storage.db"
	batchSize     = 100
)

func main() {
	// token = flags.Get(token)

	// tgClient := telegram.New(tgBotHost, mustToken())
	// s := files.New(storagePath)
	s, err := sqlite.New(sqliteStorage)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProccessor := telegram.New(tgClient.New(tgBotHost, mustToken()), s)

	log.Print("service started")

	consumer := event_consumer.New(eventsProccessor, eventsProccessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

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
		"tg-bot-token",
		"",
		"token for access to telegramm bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
