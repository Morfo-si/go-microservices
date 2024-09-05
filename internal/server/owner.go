package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllOwners(ctx fiber.Ctx) error {
	emailAddress := ctx.Params("emailAddress")

	owners, err := s.DB.GetAllOwners(ctx.Context(), emailAddress)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(owners)
}

func (s *EchoServer) AddOwner(ctx fiber.Ctx) error {
	owner := new(models.Owner)
	if err := ctx.Bind().Body(owner); err != nil {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(err)
	}
	owner, err := s.DB.AddOwner(ctx.Context(), owner)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.Status(fiber.StatusConflict).JSON(err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(fiber.StatusCreated).JSON(owner)
}

func (s *EchoServer) GetOwnerById(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	owner, err := s.DB.GetOwnerById(ctx.Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(owner)
}

func (s *EchoServer) UpdateOwner(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	owner := new(models.Owner)
	if err := ctx.Bind().Body(owner); err != nil {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(err)
	}
	if ID != owner.OwnerID {
		return ctx.Status(fiber.StatusBadRequest).JSON("id on path does not match id on body")
	}
	owner, err := s.DB.UpdateOwner(ctx.Context(), owner)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		case *dberrors.ConflictError:
			return ctx.Status(fiber.StatusConflict).JSON(err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(owner)
}

func (s *EchoServer) DeleteOwner(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	err := s.DB.DeleteOwner(ctx.Context(), ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusResetContent).JSON(nil)
}
