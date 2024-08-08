package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"warehouse/internal/entity"
)

func ctxToCreateMaterial(ctx *fasthttp.RequestCtx) (entity.Material, error) {
	body := ctx.PostBody()
	var material entity.Material

	if err := json.Unmarshal(body, &material); err != nil {
		return material, fmt.Errorf("bad request")
	}

	if material.Type == "" {
		return material, fmt.Errorf("arg 'type' is missing")
	}

	if material.Name == "" {
		return material, fmt.Errorf("arg 'name' is missing")
	}

	if material.Description == "" {
		return material, fmt.Errorf("arg 'description' is missing")
	}

	return material, nil
}
