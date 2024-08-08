package handlers

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"warehouse/internal/controller/output"
	"warehouse/internal/materials"
)

type MaterialsHandler struct {
	uc materials.UseCase
}

func NewMaterialsHandler(uc materials.UseCase) *MaterialsHandler { return &MaterialsHandler{uc: uc} }

func (m *MaterialsHandler) CreateMaterial(ctx *fasthttp.RequestCtx) {
	var err error
	material, err := ctxToCreateMaterial(ctx)
	if err != nil {
		log.Err(err).Send()
		output.JsonMessageResult(ctx, 500, err.Error())
		return
	}

	id, err := m.uc.Create(context.Background(), &material)
	if err != nil {
		log.Err(err).Send()
		output.JsonMessageResult(ctx, 500, err.Error())
		return
	}
	output.JsonNoIndent(ctx, 200, map[string]string{"uuid": id})
}

//func (m *MaterialsHandler) GetMaterialByID(ctx *fasthttp.RequestCtx) {
//	//var err error
//
//	uuid := ctx.UserValue("uuid").(string)
//
//}
