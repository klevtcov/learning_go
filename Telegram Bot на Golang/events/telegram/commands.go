package telegram

import (
	"errors"
	"log"
	"net/url"
	"strings"

	e "github.com/klevtcov/tg-bot/lib"
	"github.com/klevtcov/tg-bot/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	// убираем пробелы из полученного текста
	text = strings.TrimSpace(text)

	log.Printf("got new command `%s` from `%s`", text, username)

	// сохраниение страницы add page: http://...
	// rnd page: /rnd
	// help: /help
	// start: /start: hi + help

	if isAddCmd(text) {
		// TODO: addPage()
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:
	}

}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	// в возврате ошибка именованная, т.к. ожидается в разных местах
	// сразу задаём отложенную обработку ошибок через defer
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	// проверяем, существует ли у нас данная страница в сохранённых
	isExists, err := p.storage.IsExist(page)
	if err != nil {
		return err
	}

	//  если существует, возвращаем в чат сообщение
	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	// пытаемся сохранить страницу
	if err := p.storage.Save(page); err != nil {
		return err
	}

	// если сохранилась корректно, отправляем сообщение пользователю
	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()

	// пытаемся получитьслучайную ссылку по имени пользователя
	// если нет сохранённых страниц, выводим ошибку
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}

	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}
}

func isAddCmd(text string) bool {
	return isUrl(text)
}

func isUrl(text string) bool {
	// не распарсит ссылки без протокола, т.е. вида ya.ru
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
