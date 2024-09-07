package model

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProvinceAll struct {
	Name string `json:"name"`
}

type Categories struct {
	CID   int    `json:"category_id"`
	CName string `json:"category_name"`
}

type Products struct {
	PID         int     `json:"product_id"`
	PName       string  `json:"product_name"`
	CID         int     `json:"category_id"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Image_path  string  `json:"image_path"`
}
