package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllPets(ctx echo.Context) error {
	pets, err := s.DB.GetAllPets(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, pets)
}

func (s *EchoServer) AddPet(ctx echo.Context) error {
	pet := new(models.Pet)
	if err := ctx.Bind(pet); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	pet, err := s.DB.AddPet(ctx.Request().Context(), pet)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, pet)
}

func (s *EchoServer) GetPetById(ctx echo.Context) error {
	ID := ctx.Param("id")
	pet, err := s.DB.GetPetById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (s *EchoServer) UpdatePet(ctx echo.Context) error {
	ID := ctx.Param("id")
	pet := new(models.Pet)
	if err := ctx.Bind(pet); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != pet.PetID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}
	pet, err := s.DB.UpdatePet(ctx.Request().Context(), pet)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (s *EchoServer) DeletePet(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeletePet(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
