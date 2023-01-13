package datacookie

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Data interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)

	Save() error
}

func New(w http.ResponseWriter, r *http.Request) Data {
	d := &dataCookie{
		wr: w,
		rd: r,
	}

	d.load()
	return d
}

type dataCookie struct {
	wr http.ResponseWriter
	rd *http.Request

	values map[string]interface{}
}

// Получение значения из хранилища значений
func (d *dataCookie) Get(key string) interface{} {
	val, ok := d.values[key]
	if !ok {
		return nil
	}

	return val
}

// Установка значения в хранилище значений
func (d *dataCookie) Set(key string, value interface{}) {
	d.values[key] = value
}

// Удаление значения из хранилища значений
func (d *dataCookie) Remove(key string) {
	delete(d.values, key)
}

// Сохранения токена с обновленными данными
func (d *dataCookie) Save() error {
	_, token, err := tokenizer.Encode(d.values)
	if err != nil {
		return err
	}

	http.SetCookie(d.wr, &http.Cookie{
		Name:     "data",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(12 * time.Hour),
	})

	return nil
}

func (d *dataCookie) load() {
	c, err := d.rd.Cookie("data")
	if err != nil || c == nil {
		d.values = make(map[string]interface{})
		return
	}

	token, err := tokenizer.Decode(c.Value)
	if err != nil {
		d.values = make(map[string]interface{})
		log.Println("Не удалось декодировать токен", err)
		return
	}

	values, err := token.AsMap(context.Background())
	if err != nil {
		d.values = make(map[string]interface{})
		log.Println("Не удалось извлечь данные в виде карты из токена", err)
		return
	}

	d.values = values
}
