package request

type SampleRequest struct {
	Sample int `json:"sample" binding:"required" validate:"required"`
}