package helloWorldController

import (
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	DB
}

func (hw HelloWorldController) List(ctx *gin.Context) {
	var users []*entities.HelloWorldEntity
	hw.DB.Find(&users)

	payload := make([]presenters.ApiReturnModel, len(users))

	for _, entity := range users {
		payload = append(payload, entity.CastToApiReturnModel())
	}

	response := hw.JsonPresenter.EnvelopeList(payload)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	jsonPresenter *presenters.JsonPresenter,
) {
	helloWorldController := HelloWorldController{
		DB:            db,
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/hello-world", helloWorldController.List)
}
