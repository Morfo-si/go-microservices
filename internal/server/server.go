package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/morfo-si/go-microservices/internal/database"
	"github.com/morfo-si/go-microservices/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx fiber.Ctx) error
	Liveness(ctx fiber.Ctx) error

	GetAllOwners(ctx fiber.Ctx) error
	AddOwner(ctx fiber.Ctx) error
	GetOwnerById(ctx fiber.Ctx) error
	UpdateOwner(ctx fiber.Ctx) error
	DeleteOwner(ctx fiber.Ctx) error

	GetAllAppointments(ctx fiber.Ctx) error
	AddAppointment(ctx fiber.Ctx) error
	GetAppointmentById(ctx fiber.Ctx) error
	UpdateAppointment(ctx fiber.Ctx) error
	DeleteAppointment(ctx fiber.Ctx) error

	GetAllPets(ctx fiber.Ctx) error
	AddPet(ctx fiber.Ctx) error
	GetPetById(ctx fiber.Ctx) error
	UpdatePet(ctx fiber.Ctx) error
	DeletePet(ctx fiber.Ctx) error

	GetAllVeterinarians(ctx fiber.Ctx) error
	AddVeterinarian(ctx fiber.Ctx) error
	GetVeterinarianById(ctx fiber.Ctx) error
	UpdateVeterinarian(ctx fiber.Ctx) error
	DeleteVeterinarian(ctx fiber.Ctx) error
}

type EchoServer struct {
	app *fiber.App
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		app: fiber.New(fiber.Config{
			AppName:       "PetClinic",
			BodyLimit:     fiber.DefaultBodyLimit,
			ServerHeader:  "PetClinic",
			StrictRouting: false,
			ReadTimeout:   1 * time.Second,
			WriteTimeout:  1 * time.Second,
			IdleTimeout:   10 * time.Second,
		}),
		DB: db,
	}

	server.app.Use(logger.New(logger.Config{
		Format:        "${time} [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone:      "UTC",
		Output:        os.Stdout,
		DisableColors: false,
	}))

	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.app.Listen(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.app.Get("/readiness", s.Readiness)
	s.app.Get("/liveness", s.Liveness)

	og := s.app.Group("/owners")
	og.Get("", s.GetAllOwners)
	og.Get("/:id", s.GetOwnerById)
	og.Post("", s.AddOwner)
	og.Put("/:id", s.UpdateOwner)
	og.Delete("/:id", s.DeleteOwner)

	ag := s.app.Group("/appointments")
	ag.Get("", s.GetAllAppointments)
	ag.Get("/:id", s.GetAppointmentById)
	ag.Post("", s.AddAppointment)
	ag.Put("/:id", s.UpdateAppointment)
	ag.Delete("/:id", s.DeleteAppointment)

	pg := s.app.Group("/pets")
	pg.Get("", s.GetAllPets)
	pg.Get("/:id", s.GetPetById)
	pg.Post("", s.AddPet)
	pg.Put("/:id", s.UpdatePet)
	pg.Delete("/:id", s.DeletePet)

	vg := s.app.Group("/veterinarians")
	vg.Get("", s.GetAllVeterinarians)
	vg.Get("/:id", s.GetVeterinarianById)
	vg.Post("", s.AddVeterinarian)
	vg.Put("/:id", s.UpdateVeterinarian)
	vg.Delete("/:id", s.DeleteVeterinarian)
}

func (s *EchoServer) Readiness(ctx fiber.Ctx) error {
	ready := s.DB.Ready()
	if ready {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(models.Health{Status: "OK"})
	}
	ctx.Status(fiber.StatusInternalServerError)
	return ctx.JSON(models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx fiber.Ctx) error {
	ctx.Status(fiber.StatusOK)
	return ctx.JSON(models.Health{Status: "OK"})
}
