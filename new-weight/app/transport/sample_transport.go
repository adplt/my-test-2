package transport

import (
	"project-name/app/shared/pkg/log"
	"project-name/app/shared/pkg/response"
	"project-name/app/usecase/sample"

	"github.com/gin-gonic/gin"

	structs "project-name/app/shared/structs"
)

type sampleTransport struct {
	sampleUC sample.UseCase
	res      response.ServerResponse
}

func NewSampleTransport(useCase sample.UseCase, serverResponse response.ServerResponse) *sampleTransport {
	return &sampleTransport{
		sampleUC: useCase,
		res:      serverResponse,
	}
}

func (tp *sampleTransport) GetWeightById(c *gin.Context) {
	var body structs.GetWeight

	weightRecordID := c.Param("weightRecordId")
	body.WeightRecordId = &weightRecordID

	res, httpStatusCode, code, err := tp.sampleUC.GetWeightById(body)
	if err != nil {
		log.Error(c, err)
		tp.res.ResponseWithCode(c, httpStatusCode, false, code, response.StatusOwner, err)
		return
	}

	tp.res.SuccessOK(c, res)
}

func (tp *sampleTransport) GetWeight(c *gin.Context) {
	body := c.MustGet("body").(structs.GetWeight)

	res, httpStatusCode, code, err := tp.sampleUC.GetWeight(body)
	if err != nil {
		log.Error(c, err)
		tp.res.ResponseWithCode(c, httpStatusCode, false, code, response.StatusOwner, err)
		return
	}

	tp.res.SuccessOK(c, res)
}

func (tp *sampleTransport) AddWeight(c *gin.Context) {
	body := c.MustGet("body").(structs.WeightData)

	res, httpStatusCode, code, err := tp.sampleUC.AddWeight(body)
	if err != nil {
		log.Error(c, err)
		tp.res.ResponseWithCode(c, httpStatusCode, false, code, response.StatusOwner, err)
		return
	}

	tp.res.SuccessOK(c, res)
}

func (tp *sampleTransport) UpdateWeight(c *gin.Context) {
	body := c.MustGet("body").(structs.WeightData)

	weightRecordID := c.Param("weightRecordId")
	body.WeightRecordId = &weightRecordID

	res, httpStatusCode, code, err := tp.sampleUC.UpdateWeight(body)
	if err != nil {
		log.Error(c, err)
		tp.res.ResponseWithCode(c, httpStatusCode, false, code, response.StatusOwner, err)
		return
	}

	tp.res.SuccessOK(c, res)
}
