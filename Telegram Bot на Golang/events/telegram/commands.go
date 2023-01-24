package telegram

import (
	"log"
	"strings"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	// убираем пробелы из полученного текста
	text = strings.TrimSpace(text)
	log.Printf("got new command")
}
