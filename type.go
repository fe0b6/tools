package tools

// SelectObj - объект для формирования выпадающих списков
type SelectObj struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Timezone - объект временой зоны
type Timezone struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
