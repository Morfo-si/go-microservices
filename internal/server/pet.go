package server

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *AppServer) GetAllPets(ctx fiber.Ctx) error {
	pets, err := s.DB.GetAllPets(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(pets)
}

func (s *AppServer) AddPet(ctx fiber.Ctx) error {
	pet := new(models.Pet)
	if err := ctx.Bind().Body(pet); err != nil {
		return ctx.Status(http.StatusUnsupportedMediaType).JSON(err)
	}
	pet, err := s.DB.AddPet(ctx.Context(), pet)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.Status(http.StatusConflict).JSON(err)
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(http.StatusCreated).JSON(pet)
}

func (s *AppServer) GetPetById(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	pet, err := s.DB.GetPetById(ctx.Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(http.StatusNotFound).JSON(err)
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(http.StatusOK).JSON(pet)
}

func (s *AppServer) UpdatePet(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	pet := new(models.Pet)
	if err := ctx.Bind().Body(pet); err != nil {
		return ctx.Status(http.StatusUnsupportedMediaType).JSON(err)
	}
	if ID != pet.PetID {
		return ctx.Status(http.StatusBadRequest).JSON("id on path doesn't match id on body")
	}
	pet, err := s.DB.UpdatePet(ctx.Context(), pet)
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
	return ctx.Status(http.StatusOK).JSON(pet)
}

func (s *AppServer) DeletePet(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	err := s.DB.DeletePet(ctx.Context(), ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusResetContent).JSON(nil)
}
