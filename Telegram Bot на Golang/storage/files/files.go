package files

import (
	"encoding/gob"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	e "github.com/klevtcov/tg-bot/lib"
	"github.com/klevtcov/tg-bot/storage"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

var ErrNoSavedPages = errors.New("no saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

// функция сохранения полученной страницы
func (s Storage) Save(page *storage.Page) (err error) {
	// определяемм способ обработки ошибок
	defer func() { err = e.WrapIfErr("can't save page: ", err) }()

	// формируем путь до директории, куда будем сохранять данные
	fPath := filepath.Join(s.basePath, page.UserName)

	// создаём нужные директории
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	// формируем имя файла
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	// дописываем имя файла к пути
	fPath = filepath.Join(fPath, fName)

	// создаём файл
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	// записываем в него страницу в нужном формате
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	// определяемм способ обработки ошибок
	defer func() { err = e.WrapIfErr("can't pick random: ", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	// random от 0 до кол-ва файлов. рандом не рандомный, поэтому передаём текущее время
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	// для возврата страницы
	// open decode

}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	// открываем файл, читаем содержимое в переменную f
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode page: ", err)
	}

	// закрываем файл по окончанию чтения
	defer func() {_=f.Close()} ()

	var p storage.Page

	if err :=
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
