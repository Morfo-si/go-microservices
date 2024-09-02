package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morfo-si/go-microservices/internal/database"
	"github.com/morfo-si/go-microservices/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllOwners(ctx echo.Context) error
	AddOwner(ctx echo.Context) error
	GetOwnerById(ctx echo.Context) error
	UpdateOwner(ctx echo.Context) error
	DeleteOwner(ctx echo.Context) error

	GetAllAppointments(ctx echo.Context) error
	AddAppointment(ctx echo.Context) error
	GetAppointmentById(ctx echo.Context) error
	UpdateAppointment(ctx echo.Context) error
	DeleteAppointment(ctx echo.Context) error

	GetAllPets(ctx echo.Context) error
	AddPet(ctx echo.Context) error
	GetPetById(ctx echo.Context) error
	UpdatePet(ctx echo.Context) error
	DeletePet(ctx echo.Context) error

	GetAllVeterinarians(ctx echo.Context) error
	AddVeterinarian(ctx echo.Context) error
	GetVeterinarianById(ctx echo.Context) error
	UpdateVeterinarian(ctx echo.Context) error
	DeleteVeterinarian(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	og := s.echo.Group("/owners")
	og.GET("", s.GetAllOwners)
	og.GET("/:id", s.GetOwnerById)
	og.POST("", s.AddOwner)
	og.PUT("/:id", s.UpdateOwner)
	og.DELETE("/:id", s.DeleteOwner)

	ag := s.echo.Group("/appointments")
	ag.GET("", s.GetAllAppointments)
	ag.GET("/:id", s.GetAppointmentById)
	ag.POST("", s.AddAppointment)
	ag.PUT("/:id", s.UpdateAppointment)
	ag.DELETE("/:id", s.DeleteAppointment)

	pg := s.echo.Group("/pets")
	pg.GET("", s.GetAllPets)
	pg.GET("/:id", s.GetPetById)
	pg.POST("", s.AddPet)
	pg.PUT("/:id", s.UpdatePet)
	pg.DELETE("/:id", s.DeletePet)

	vg := s.echo.Group("/veterinarians")
	vg.GET("", s.GetAllVeterinarians)
	vg.GET("/:id", s.GetVeterinarianById)
	vg.POST("", s.AddVeterinarian)
	vg.PUT("/:id", s.UpdateVeterinarian)
	vg.DELETE("/:id", s.DeleteVeterinarian)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
