package request

type MBOnBoardingRequest struct {
	Sample string `json:"sample" binding:"required" validate:"required"`
}
