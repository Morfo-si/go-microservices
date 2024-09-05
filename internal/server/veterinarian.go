package server

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllVeterinarians(ctx fiber.Ctx) error {
	veterinarians, err := s.DB.GetAllVeterinarians(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(veterinarians)
}

func (s *EchoServer) AddVeterinarian(ctx fiber.Ctx) error {
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind().Body(veterinarian); err != nil {
		return ctx.Status(http.StatusUnsupportedMediaType).JSON(err)
	}
	veterinarian, err := s.DB.AddVeterinarian(ctx.Context(), veterinarian)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.Status(http.StatusConflict).JSON(err)
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(http.StatusCreated).JSON(veterinarian)
}

func (s *EchoServer) GetVeterinarianById(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	veterinarian, err := s.DB.GetVeterinarianById(ctx.Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(http.StatusNotFound).JSON(err)
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(http.StatusOK).JSON(veterinarian)
}

func (s *EchoServer) UpdateVeterinarian(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind().Body(veterinarian); err != nil {
		return ctx.Status(http.StatusUnsupportedMediaType).JSON(err)
	}
	if ID != veterinarian.VeterinarianID {
		return ctx.Status(http.StatusBadRequest).JSON("id on path doesn't match id on body")
	}
	veterinarian, err := s.DB.UpdateVeterinarian(ctx.Context(), veterinarian)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(http.StatusNotFound).JSON(err)
		case *dberrors.ConflictError:
			return ctx.Status(http.StatusConflict).JSON(err)
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(http.StatusOK).JSON(veterinarian)
}

func (s *EchoServer) DeleteVeterinarian(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	err := s.DB.DeleteVeterinarian(ctx.Context(), ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusResetContent).JSON(nil)
}
