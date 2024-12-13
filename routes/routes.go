package routes

import (
	"belajar-api/config"
	"belajar-api/handlers"
	"belajar-api/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	organizationService := &services.MasterOrganizationService{DB: db}
	natureService := &services.MasterNatureService{DB: db}
	priorityService := &services.MasterPriorityService{DB: db}
	userBelajarService := &services.UserBelajarService{DB: db}

	userBelajarHandler := &handlers.UserBelajarHandler{UserBelajarService: userBelajarService}
	organizationHandler := &handlers.MasterOrganizationHandler{MasterOrganizationService: organizationService}
	natureHandler := &handlers.MasterNatureHandler{MasterNatureService: natureService}
	priorityHandler := &handlers.MasterPriorityHandler{MasterPriorityService: priorityService}

	api := app.Group("/api")
	api.Static("/public", "/public")

	userBelajarRoutes := api.Group("/userBelajar")
	userBelajarRoutes.Get("/", userBelajarHandler.GetAllUserBelajar)

	//protected := api.Group("/v1", middleware.JWTMiddleware())
	protected := api.Group("/v1")

	organizationRoutes := protected.Group("/organizations")
	organizationRoutes.Post("/", organizationHandler.CreateMasterOrganization)
	organizationRoutes.Get("/", organizationHandler.GetAllMasterOrganizations)
	organizationRoutes.Get("/:uid", organizationHandler.GetMasterOrganizationByUID)
	organizationRoutes.Put("/:uid", organizationHandler.UpdateMasterOrganization)
	organizationRoutes.Delete("/:uid", organizationHandler.DeleteMasterOrganization)

	natureRoutes := protected.Group("/natures")
	natureRoutes.Post("/", natureHandler.CreateMasterNature)
	natureRoutes.Get("/", natureHandler.GetAllMasterNatures)
	natureRoutes.Get("/:uid", natureHandler.GetMasterNatureByUID)
	natureRoutes.Put("/:uid", natureHandler.UpdateMasterNature)
	natureRoutes.Delete("/:uid", natureHandler.DeleteMasterNature)

	priorityRoutes := protected.Group("/priorities")
	priorityRoutes.Post("/", priorityHandler.CreateMasterPriority)
	priorityRoutes.Get("/", priorityHandler.GetAllMasterPriorities)
	priorityRoutes.Get("/:uid", priorityHandler.GetMasterPriorityByUID)
	priorityRoutes.Put("/:uid", priorityHandler.UpdateMasterPriority)
	priorityRoutes.Delete("/:uid", priorityHandler.DeleteMasterPriority)
}
