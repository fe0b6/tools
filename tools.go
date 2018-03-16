package tools

import "strings"

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
func AppendSet(m, s string) string {
	nm := strings.TrimLeft(strings.TrimRight(m, "}"), "{")

	arr := strings.Split(nm, ",")
	narr := []string{}
	for _, v := range arr {
		if v == s {
			return m
		} else if v != "" {
			narr = append(narr, v)
		}
	}

	narr = append(narr, s)
	m = strings.Join(narr, ",")

	return "{" + m + "}"
}
