package controller

import (
	"coco/service"
	"github.com/gin-gonic/gin"
)

type Index struct {
	Service service.IStartService `inject:""`
}

func (i *Index) GetName(ctx *gin.Context) {
	var message string
	ctx.ShouldBind(&message)
	Json(ctx, &Res{200, i.Service.Say(message)})
}
