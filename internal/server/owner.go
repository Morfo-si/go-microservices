package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllOwners(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")

	owners, err := s.DB.GetAllOwners(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, owners)
}

func (s *EchoServer) AddOwner(ctx echo.Context) error {
	owner := new(models.Owner)
	if err := ctx.Bind(owner); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	owner, err := s.DB.AddOwner(ctx.Request().Context(), owner)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, owner)
}

func (s *EchoServer) GetOwnerById(ctx echo.Context) error {
	ID := ctx.Param("id")
	owner, err := s.DB.GetOwnerById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, owner)
}

func (s *EchoServer) UpdateOwner(ctx echo.Context) error {
	ID := ctx.Param("id")
	owner := new(models.Owner)
	if err := ctx.Bind(owner); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != owner.OwnerID {
		return ctx.JSON(http.StatusBadRequest, "id on path does not match id on body")
	}
	owner, err := s.DB.UpdateOwner(ctx.Request().Context(), owner)
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
	return ctx.JSON(http.StatusOK, owner)
}

func (s *EchoServer) DeleteOwner(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteOwner(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
