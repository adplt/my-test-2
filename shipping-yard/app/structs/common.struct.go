package structs

// CommonRequest struct
type CommonRequest struct {
	UserId   string `json:"userId" binding:"required,max=40" validate:"required,max=40" swaggertype:"string" example:"Postman"`
	UserName string `json:"userName,omitempty" binding:"max=40" validate:"max=40,omitempty" swaggertype:"string" example:"Postman"`
	Source   string `json:"source,omitempty" binding:"max=40" validate:"max=40,omitempty" swaggertype:"string" example:"Postman"`
}

// ResponseCommonArray struct
type ResponseCommonArray struct {
	Data    []map[string]interface{} `json:"data"`
	Error   string                   `json:"error" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
	Message string                   `json:"message" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
}

// ResponseCommonObject struct
type ResponseCommonObject struct {
	Data    map[string]interface{} `json:"data"`
	Error   string                 `json:"error" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
	Message string                 `json:"message" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
}

// ResponseCommonObject struct
type ResponseCommonString struct {
	Data    string `json:"data"`
	Error   string `json:"error" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
	Message string `json:"message" binding:"required,max=40" validate:"max=40,regexp=^[0-9a-zA-Z !@#$%^&*()<>?.:]*$,nonzero,nonnil"`
}

// ResponseError400 struct
type ResponseError400 struct {
	Data    interface{} `json:"data" swaggertype:"object" example:""`
	Message string      `json:"message" swaggertype:"string" example:"Key: 'GetRestaurantTrx.RestName' Error:Field validation for 'RestName' failed on the 'required' tag"`
	Status  string      `json:"status" swaggertype:"string" example:"FAILED"`
}

// ResponseError500 struct
type ResponseError500 struct {
	Data    interface{} `json:"data" swaggertype:"object" example:""`
	Message string      `json:"message" swaggertype:"string" example:"Restaurant not found"`
	Status  string      `json:"status" swaggertype:"string" example:"FAILED"`
}

// LogJSON struct
type LogJSON struct {
	ClientIP   string  `json:"clientip"`
	TimeStamp  string  `json:"timestamp"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	StatusCode int     `json:"statuscode"`
	Latency    float64 `json:"latency"`
	Request    string  `json:"request"`
	Response   string  `json:"response"`
}
