package controller

import (
	"github.com/fasthttp/router"
	"warehouse/config"
	"warehouse/internal/controller/handlers"
	"warehouse/internal/controller/output"
	"warehouse/internal/materials"
)

func NewRouter(cfg *config.Config, e materials.UseCase) *router.Router {
	r := router.New()

	// material handlers
	materialsHandler := handlers.NewMaterialsHandler(e)

	// Создание материала
	r.POST("/materials", materialsHandler.CreateMaterial)
	// Получение по UUID материала
	//r.GET("/materials/{uuid}", materialsHandler.GetMaterialByID)
	//// Обновление материала
	//r.PUT("/materials/{uuid}", materialsHandler.UpdateMaterial)
	////Получение всех материалов
	//r.GET("/materials", materialsHandler.GetAllMaterials)

	r.OPTIONS("/materials", output.CORSOptions)
	r.OPTIONS("/materials/{uuid}", output.CORSOptions)
	return r
}
