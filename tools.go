package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	processPidReg *regexp.Regexp
)

func init() {
	processPidReg = regexp.MustCompile("^\\S+\\s+(\\d+)")
}

// ChunkSliceString - Разбиваем массив строк на несколько
func ChunkSliceString(arr []string, size int) (ans [][]string) {
	msize := len(arr) / size
	if len(arr)%size != 0 {
		msize++
	}
	ans = make([][]string, msize)

	l := len(arr)
	now := 0
	i := 0
	for {
		next := now + size

		if now+size > l {
			next = l
		}

		ans[i] = arr[now:next]
		i++

		if next == l {
			break
		}
		now = next
	}

	return
}

// ChunkSliceInt - Разбиваем массив int на несколько
func ChunkSliceInt(arr []int, size int) (ans [][]int) {
	msize := len(arr) / size
	if len(arr)%size != 0 {
		msize++
	}
	ans = make([][]int, msize)

	l := len(arr)
	now := 0
	i := 0
	for {
		next := now + size

		if now+size > l {
			next = l
		}

		ans[i] = arr[now:next]
		i++

		if next == l {
			break
		}
		now = next
	}

	return
}

// ChunkSliceInterface - Разбиваем массив interface{} на несколько
func ChunkSliceInterface(arr []interface{}, size int) (ans [][]interface{}) {
	msize := len(arr) / size
	if len(arr)%size != 0 {
		msize++
	}
	ans = make([][]interface{}, msize)

	l := len(arr)
	now := 0
	i := 0
	for {
		next := now + size

		if now+size > l {
			next = l
		}

		ans[i] = arr[now:next]
		i++

		if next == l {
			break
		}
		now = next
	}

	return
}

// ChunkSliceMapStrInt - Разбиваем массив map[string]int на несколько
func ChunkSliceMapStrInt(arr []map[string]int, size int) (ans [][]map[string]int) {
	msize := len(arr) / size
	if len(arr)%size != 0 {
		msize++
	}
	ans = make([][]map[string]int, msize)

	l := len(arr)
	now := 0
	i := 0
	for {
		next := now + size

		if now+size > l {
			next = l
		}

		ans[i] = arr[now:next]
		i++

		if next == l {
			break
		}
		now = next
	}

	return
}

// ChunkSliceMapStrInterface - Разбиваем массив map[string]interface{} на несколько
func ChunkSliceMapStrInterface(arr []map[string]interface{}, size int) (ans [][]map[string]interface{}) {
	msize := len(arr) / size
	if len(arr)%size != 0 {
		msize++
	}
	ans = make([][]map[string]interface{}, msize)

	l := len(arr)
	now := 0
	i := 0
	for {
		next := now + size

		if now+size > l {
			next = l
		}

		ans[i] = arr[now:next]
		i++

		if next == l {
			break
		}
		now = next
	}

	return
}

// InArray - проверяем содержит ли массив строку
func InArray(arr []string, str string) (ok bool) {
	for i := range arr {
		if arr[i] == str {
			ok = true
			break
		}
	}

	return
}

// HasArray - проверяем содержитат ли массивы пересечение
func HasArray(arr1, arr2 []string) (ok bool) {
	if len(arr1) > len(arr2) {
		h := map[string]bool{}
		for _, a := range arr2 {
			h[a] = true
		}

		for _, a := range arr1 {
			if h[a] {
				ok = true
				break
			}
		}
	} else {
		h := map[string]bool{}
		for _, a := range arr1 {
			h[a] = true
		}

		for _, a := range arr2 {
			if h[a] {
				ok = true
				break
			}
		}
	}

	return
}

// GetPlaceholders - Возвращает строку placeholders нужной длины
func GetPlaceholders(l int) string {
	return strings.TrimRight(strings.Repeat("?,", l), ",")
}

// AppendSet - добавляем элемент в массив
func AppendSet(m []string, s string) []string {

	h := map[string]bool{s: true}
	for _, v := range m {
		if v == s {
			return m
		}
		h[v] = true
	}

	narr := []string{}
	for k := range h {
		narr = append(narr, k)
	}

	return narr
}

// RemoveSet - удаляем элемент в массив
func RemoveSet(m []string, s string) []string {

	narr := []string{}
	for _, v := range m {
		if v == s {
			continue
		} else if v != "" {
			narr = append(narr, v)
		}
	}

	return narr
}

// CheckSet - проверяем есть ли элемент в массиве
func CheckSet(m []string, s string) bool {
	return InArray(m, s)
}

// FloatTrunc - обрезаем float64 до нужной длины
func FloatTrunc(num, precision float64) float64 {
	output := math.Pow(10, precision)
	return float64(int(num*output+math.Copysign(0.5, num*output))) / output
}

// GetProcessID - получаем id процесса
func GetProcessID(qs []string) (pid int) {
	cmd := exec.Command("ps", "axuw")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("[error]", err)
		return
	}

	for _, str := range strings.Split(string(out), "\n") {
		var ok int
		for _, q := range qs {
			if strings.Contains(strings.ToLower(str), strings.ToLower(q)) {
				ok++
			}
		}

		if ok != len(qs) {
			continue
		}

		found := processPidReg.FindStringSubmatch(str)
		if len(found) > 0 {
			p, _ := strconv.ParseInt(found[1], 10, 64)
			pid = int(p)
			return
		}
	}

	return
}

// GetTimezones - получаем список временых зон
func GetTimezones() (tz []Timezone) {

	tz = []Timezone{
		Timezone{Name: "UTC−12:00 — Линия перемены даты", Value: -12},
		Timezone{Name: "UTC−11:00 — Американское Самоа", Value: -11},
		Timezone{Name: "UTC−10:00 — Гавайи", Value: -10},
		Timezone{Name: "UTC−09:00 — Аляска", Value: -9},
		Timezone{Name: "UTC−08:00 — Тихоокеанское время (США и Канада)", Value: -8},
		Timezone{Name: "UTC−07:00 — Аризона, Чиуауа, Ла-Пас, Масатлан", Value: -7},
		Timezone{Name: "UTC−06:00 — Центральное время (США и Канада)", Value: -6},
		Timezone{Name: "UTC−05:00 — Восточное время (США и Канада)", Value: -5},
		Timezone{Name: "UTC−04:30 — Каракас", Value: -4.5},
		Timezone{Name: "UTC−04:00 — Атлантическое время (Канада)", Value: -4},
		Timezone{Name: "UTC−03:30 — Ньюфаундленд", Value: -3.5},
		Timezone{Name: "UTC−03:00 — Бразилиа, Буэнос-Айрес, Джорджтаун", Value: -3},
		Timezone{Name: "UTC−02:00 — Среднеатлантическое время", Value: -2},
		Timezone{Name: "UTC−01:00 — Азорские острова, Кабо-Верде", Value: -1},
		Timezone{Name: "UTC+00:00 — Западноевропейское время (Лондон)", Value: 0},
		Timezone{Name: "UTC+01:00 — Центральноевропейское время (Берлин, Париж)", Value: 1},
		Timezone{Name: "UTC+02:00 — Киев, Калининград, Египет, Израиль", Value: 2},
		Timezone{Name: "UTC+03:00 — Московское время", Value: 3},
		Timezone{Name: "UTC+03:30 — Тегеранское время", Value: 3.5},
		Timezone{Name: "UTC+04:00 — Самарское время", Value: 4},
		Timezone{Name: "UTC+04:30 — Афганистан", Value: 4.5},
		Timezone{Name: "UTC+05:00 — Екатеринбургское время", Value: 5},
		Timezone{Name: "UTC+05:30 — Индия, Шри-Ланка", Value: 5.5},
		Timezone{Name: "UTC+05:45 — Непал", Value: 5.75},
		Timezone{Name: "UTC+06:00 — Омское время, Новосибирск", Value: 6},
		Timezone{Name: "UTC+06:30 — Мьянма", Value: 6.5},
		Timezone{Name: "UTC+07:00 — Красноярское время", Value: 7},
		Timezone{Name: "UTC+08:00 — Иркутское время, Гонконг, Китай", Value: 8},
		Timezone{Name: "UTC+08:45 — пять городов Австралии", Value: 8.75},
		Timezone{Name: "UTC+09:00 — Якутское время, Корея, Япония", Value: 9},
		Timezone{Name: "UTC+09:30 — Центральноавстралийское время (Аделаида, Дарвин)", Value: 9.5},
		Timezone{Name: "UTC+10:00 — Владивостокское время", Value: 10},
		Timezone{Name: "UTC+10:30 — часть Австралии", Value: 10.5},
		Timezone{Name: "UTC+11:00 — Среднеколымское время", Value: 11},
		Timezone{Name: "UTC+11:30 — остров Норфолк (Австралия)", Value: 11.5},
		Timezone{Name: "UTC+12:00 — Камчатское время, Новая Зеландия", Value: 12},
		Timezone{Name: "UTC+12:45 — архипелаг Чатем (Новая Зеландия)", Value: 12.75},
		Timezone{Name: "UTC+13:00 — Самоа, Тонга", Value: 13},
		Timezone{Name: "UTC+13:45 — летнее время на архипелаге Чатем (Новая Зеландия)", Value: 13.75},
		Timezone{Name: "UTC+14:00 — Острова Лайн", Value: 14},
	}

	for i := range tz {
		tz[i].Value *= 60
	}

	return
}

// FromGob - преобразуем gob в объект
func FromGob(i interface{}, b []byte) {
	var s bytes.Buffer
	s.Write(b)
	gr := gob.NewDecoder(&s)
	gr.Decode(i)
}

// ToGob - преобразуем объект в gob
func ToGob(i interface{}) []byte {
	var s bytes.Buffer
	gr := gob.NewEncoder(&s)
	gr.Encode(i)
	return s.Bytes()
}

// ToJSON - преобразуем объект в json
func ToJSON(i interface{}) (b []byte) {
	b, err := json.Marshal(i)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}

// GetMonthName - Возвращаем название месяца
func GetMonthName(m int) string {
	switch m {
	case 1:
		return "Январь"
	case 2:
		return "Февраль"
	case 3:
		return "Март"
	case 4:
		return "Апрель"
	case 5:
		return "Май"
	case 6:
		return "Июнь"
	case 7:
		return "Июль"
	case 8:
		return "Август"
	case 9:
		return "Сентябрь"
	case 10:
		return "Октябрь"
	case 11:
		return "Ноябрь"
	case 12:
		return "Декабрь"
	}

	// 0 тоже декабрь
	return "Декабрь"
}

// ArrToInterface - приведение любого среза к типу интерфейс
func ArrToInterface(a interface{}) []interface{} {
	s := reflect.ValueOf(a)
	if s.Kind() != reflect.Slice {
		return []interface{}{a}
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// IsClosedChan - Проверяем закрыт ли канал или нет
func IsClosedChan(c chan struct{}) (ok bool) {
	select {
	case <-c:
		ok = true
	default:
	}

	return
}

// AESEncrypt - Шифруем данные с помощью AES
func AESEncrypt(key, data []byte) (cipherData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherData = make([]byte, aes.BlockSize+len(data))
	iv := cipherData[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherData[aes.BlockSize:], data)
	return
}

// AESDecrypt - Расшифровываем данные с помощью AES
func AESDecrypt(key, cipherData []byte) (data []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherData) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherData, cipherData)

	data = cipherData
	return
}

// WaitTo - функция ожидания конкретного времени
func WaitTo(t map[string]int) {
	tn := time.Now()

	var nt time.Time
	s, m, h := t["s"], t["m"], t["h"]

	if _, ok := t["h"]; ok { // Если указаны часы
		nt = tn.Truncate(time.Hour).Add(time.Duration(m)*time.Minute + time.Duration(s)*time.Second)

		if tn.Hour() > h || (tn.Hour() == h && tn.Minute() >= m) {
			nt = nt.AddDate(0, 0, 1).Add(time.Duration(h-tn.Hour()) * time.Hour)
		} else {
			nt = nt.Add(time.Duration(h-tn.Hour()) * time.Hour)
		}

	} else if _, ok := t["m"]; ok { // Если указаны минуты
		nt = tn.Truncate(time.Hour).Add(time.Duration(m)*time.Minute + time.Duration(s)*time.Second)

		if tn.Minute() >= m {
			nt = nt.Add(time.Hour)
		}
	} else if _, ok := t["s"]; ok { // Только секунды
		nt = tn.Truncate(time.Minute).Add(time.Duration(s) * time.Second)

		if tn.Second() >= s {
			nt = nt.Add(time.Minute)
		}
	}

	time.Sleep(nt.Sub(time.Now()))
}

// IsNil - Проверяем пустая структура или нет
func IsNil(v interface{}) bool {
	return isNil(reflect.ValueOf(v))
}
func isNil(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool() == false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0

	case reflect.Ptr, reflect.Interface:
		return isNil(v.Elem())

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isNil(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.String:
		return v.Len() == 0

	case reflect.Slice, reflect.Map:
		return v.IsNil()

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isNil(v.Field(i)) {
				return false
			}
		}
		return true
	// reflect.Chan, reflect.UnsafePointer, reflect.Func
	default:
		return v.IsNil()
	}
}

// InterfaceToMapStrInt - преобразуем интерфейс в map[string]int
func InterfaceToMapStrInt(i interface{}) (h map[string]int) {
	h = map[string]int{}

	if th, ok := i.(map[string]interface{}); ok {
		for k, v := range th {
			if vint, ok := v.(float64); ok {
				h[k] = int(vint)
			}
		}
	}

	return
}
