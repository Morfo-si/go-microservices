package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morfo-si/go-microservices/internal/database"
	"github.com/morfo-si/go-microservices/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx *gin.Context)
	Liveness(ctx *gin.Context)

	GetAllOwners(ctx *gin.Context)
	AddOwner(ctx *gin.Context)
	GetOwnerById(ctx *gin.Context)
	UpdateOwner(ctx *gin.Context)
	DeleteOwner(ctx *gin.Context)

	GetAllAppointments(ctx *gin.Context)
	AddAppointment(ctx *gin.Context)
	GetAppointmentById(ctx *gin.Context)
	UpdateAppointment(ctx *gin.Context)
	DeleteAppointment(ctx *gin.Context)

	GetAllPets(ctx *gin.Context)
	AddPet(ctx *gin.Context)
	GetPetById(ctx *gin.Context)
	UpdatePet(ctx *gin.Context)
	DeletePet(ctx *gin.Context)

	GetAllVeterinarians(ctx *gin.Context)
	AddVeterinarian(ctx *gin.Context)
	GetVeterinarianById(ctx *gin.Context)
	UpdateVeterinarian(ctx *gin.Context)
	DeleteVeterinarian(ctx *gin.Context)
}

type GinServer struct {
	echo *gin.Engine
	DB   database.DatabaseClient
}

func NewGinServer(db database.DatabaseClient) Server {
	server := &GinServer{
		echo: gin.Default(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (g *GinServer) Start() error {
	if err := g.echo.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (g *GinServer) registerRoutes() {
	g.echo.GET("/readiness", g.Readiness)
	g.echo.GET("/liveness", g.Liveness)

	og := g.echo.Group("/owners")
	og.GET("", g.GetAllOwners)
	og.GET("/:id", g.GetOwnerById)
	og.POST("", g.AddOwner)
	og.PUT("/:id", g.UpdateOwner)
	og.DELETE("/:id", g.DeleteOwner)

	ag := g.echo.Group("/appointments")
	ag.GET("", g.GetAllAppointments)
	ag.GET("/:id", g.GetAppointmentById)
	ag.POST("", g.AddAppointment)
	ag.PUT("/:id", g.UpdateAppointment)
	ag.DELETE("/:id", g.DeleteAppointment)

	pg := g.echo.Group("/pets")
	pg.GET("", g.GetAllPets)
	pg.GET("/:id", g.GetPetById)
	pg.POST("", g.AddPet)
	pg.PUT("/:id", g.UpdatePet)
	pg.DELETE("/:id", g.DeletePet)

	vg := g.echo.Group("/veterinarians")
	vg.GET("", g.GetAllVeterinarians)
	vg.GET("/:id", g.GetVeterinarianById)
	vg.POST("", g.AddVeterinarian)
	vg.PUT("/:id", g.UpdateVeterinarian)
	vg.DELETE("/:id", g.DeleteVeterinarian)
}

func (g *GinServer) Readiness(ctx *gin.Context) {
	ready := g.DB.Ready()
	if ready {
		ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
	}
}

func (g *GinServer) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
