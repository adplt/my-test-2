package structs

// ResponseAddWeight struct
type ResponseAddWeight struct {
	Data    TxWeightRecord `json:"data"`
	Message string         `json:"message" swaggertype:"string" example:"Add weight successfully"`
	Status  string         `json:"status" swaggertype:"string" example:"SUCCESS"`
}

// ResponseUpdateWeight struct
type ResponseUpdateWeight struct {
	Data    TxWeightRecord `json:"data"`
	Message string         `json:"message" swaggertype:"string" example:"Update weight successfully"`
	Status  string         `json:"status" swaggertype:"string" example:"SUCCESS"`
}

// ResponseGetWeight struct
type ResponseGetWeight struct {
	Data    []TxWeightRecord `json:"data"`
	Message string           `json:"message" swaggertype:"string" example:"Get weight successfully"`
	Status  string           `json:"status" swaggertype:"string" example:"SUCCESS"`
}
