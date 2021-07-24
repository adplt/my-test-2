package structs

type AddWeight struct {
	Date string `json:"date" binding:"required" swaggertype:"string" example:"2021-07-24"`
	Max  int    `json:"max" binding:"required" swaggertype:"number" example:"55"`
	Min  int    `json:"min" binding:"required" swaggertype:"number" example:"50"`
	CommonRequest
}

type GetWeight struct {
	Date string `json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
}

type UpdateWeight struct {
	Date string `json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
	Max  int    `json:"max,omitempty" swaggertype:"number" example:"55"`
	Min  int    `json:"min,omitempty" swaggertype:"number" example:"50"`
	CommonRequest
}
