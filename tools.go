package tools

import (
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
