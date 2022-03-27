package container

import (
	"project-name/app/adapter/database"
	"project-name/app/adapter/rest_api"
	weightRepo "project-name/app/repository/weight"
	"project-name/app/shared/pkg/response"
	"project-name/app/shared/utils"
	useCaseSample "project-name/app/usecase/sample"
)

type Container struct {
	SampleUsecase  useCaseSample.UseCase
	ResponseParser response.ServerResponse
}

func SetupContainer() Container {
	// setup database
	mysql := database.SetupMySQL()
	// mongodb := database.SetupMongoDB()

	// init utils
	http := utils.NewUtilsHttp()
	httpRequest := utils.NewUtilHttpRequest()

	// init repo
	restApi := rest_api.NewRestAPI(http, httpRequest)
	weight := weightRepo.NewWeightRepo(mysql)

	// init use case
	sampleUseCase := useCaseSample.NewSampleUseCase(restApi, weight)

	// init response server parser
	responseParser := response.New(mysql.DB.DB())

	return Container{
		SampleUsecase:  sampleUseCase,
		ResponseParser: responseParser,
	}
}
