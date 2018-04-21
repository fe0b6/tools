package tools

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
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
