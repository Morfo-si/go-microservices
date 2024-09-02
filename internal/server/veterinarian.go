package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllVeterinarians(ctx echo.Context) error {
	veterinarians, err := s.DB.GetAllVeterinarians(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, veterinarians)
}

func (s *EchoServer) AddVeterinarian(ctx echo.Context) error {
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind(veterinarian); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	veterinarian, err := s.DB.AddVeterinarian(ctx.Request().Context(), veterinarian)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, veterinarian)
}

func (s *EchoServer) GetVeterinarianById(ctx echo.Context) error {
	ID := ctx.Param("id")
	veterinarian, err := s.DB.GetVeterinarianById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, veterinarian)
}

func (s *EchoServer) UpdateVeterinarian(ctx echo.Context) error {
	ID := ctx.Param("id")
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind(veterinarian); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != veterinarian.VeterinarianID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}
	veterinarian, err := s.DB.UpdateVeterinarian(ctx.Request().Context(), veterinarian)
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
	return ctx.JSON(http.StatusOK, veterinarian)
}

func (s *EchoServer) DeleteVeterinarian(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteVeterinarian(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
