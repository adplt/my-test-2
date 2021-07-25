package structs

type AddWeight struct {
	Date string `json:"date" binding:"required" swaggertype:"string" example:"2021-07-24"`
	Max  int    `json:"max" binding:"required" swaggertype:"number" example:"55"`
	Min  int    `json:"min" binding:"required" swaggertype:"number" example:"50"`
	CommonRequest
}

type GetWeight struct {
	Date           string `json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
	WeightRecordId string `json:"weightRecordId,omitempty" swaggertype:"string" example:"0a7df903-6253-414f-9ff4-2a70bd7fdb04"`
}

type UpdateWeight struct {
	Date string `json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
	Max  int    `json:"max,omitempty" swaggertype:"number" example:"55"`
	Min  int    `json:"min,omitempty" swaggertype:"number" example:"50"`
	CommonRequest
}
