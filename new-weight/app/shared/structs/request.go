package request

// CommonRequest struct
type CommonRequest struct {
	UserId   string `json:"userId" binding:"required,max=40" validate:"required,max=40" swaggertype:"string" example:"Postman"`
	UserName string `json:"userName,omitempty" binding:"max=40" validate:"max=40,omitempty" swaggertype:"string" example:"Postman"`
	Source   string `json:"source,omitempty" binding:"max=40" validate:"max=40,omitempty" swaggertype:"string" example:"Postman"`
}

type WeightData struct {
	WeightRecordId *string `json:"weightRecordId,omitempty" swaggertype:"string" example:"0a7df903-6253-414f-9ff4-2a70bd7fdb04"`
	Date           string  `json:"date" binding:"required" swaggertype:"string" example:"2021-07-24"`
	Max            int     `json:"max" binding:"required" swaggertype:"number" example:"55"`
	Min            int     `json:"min" binding:"required" swaggertype:"number" example:"50"`
	CommonRequest
}

type GetWeight struct {
	Date           string  `form:"date" json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
	WeightRecordId *string `json:"weightRecordId,omitempty" swaggertype:"string" example:"0a7df903-6253-414f-9ff4-2a70bd7fdb04"`
	Limit          int     `json:"limit,omitempty" swaggertype:"number" example:"10"`
	Offset         int     `json:"offset,omitempty" swaggertype:"number" example:"0"`
}
