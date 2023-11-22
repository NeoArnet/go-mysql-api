package model

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProvinceAll struct {
	Name string `json:"name"`
}
