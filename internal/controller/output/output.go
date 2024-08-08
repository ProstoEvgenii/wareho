package output

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type out struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func withHeaders(ctx *fasthttp.RequestCtx, contentType, allowHeaders string, code int) {
	ctx.Response.Header.Set("Content-Type", contentType)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", allowHeaders)
	ctx.Response.Header.SetStatusCode(code)
	ctx.Response.Header.Set("Connection", "close")
}

func withDefaultHeaders(ctx *fasthttp.RequestCtx, code int) {
	withHeaders(ctx, "application/json", "Authorization", code)
}

func CORSOptions(ctx *fasthttp.RequestCtx) {
	withHeaders(ctx, "text/html", "*", 200)
}

func JsonMessageResult(ctx *fasthttp.RequestCtx, code int, r string) {
	jsonResult, err := json.Marshal(out{code, r})
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "internal server error")
		return
	}

	if _, err := ctx.Write(jsonResult); err != nil {
		log.Error().Err(err).Send()
	}

	withDefaultHeaders(ctx, code)
}

func JsonNoIndent(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "internal server error")
		return
	}

	if _, err := ctx.Write(jsonResult); err != nil {
		log.Error().Err(err).Send()
	}

	withDefaultHeaders(ctx, code)
}
