package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllAppointments(ctx fiber.Ctx) error {
	appointmentId := ctx.Params("appointmentId")

	appointments, err := s.DB.GetAllAppointments(ctx.Context(), appointmentId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(appointments)
}

func (s *EchoServer) AddAppointment(ctx fiber.Ctx) error {
	appointment := new(models.Appointment)
	if err := ctx.Bind().Body(appointment); err != nil {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(err)
	}
	appointment, err := s.DB.AddAppointment(ctx.Context(), appointment)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.Status(fiber.StatusConflict).JSON(err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(appointment)
}

func (s *EchoServer) GetAppointmentById(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	appointment, err := s.DB.GetAppointmentById(ctx.Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(appointment)
}

func (s *EchoServer) UpdateAppointment(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	appointment := new(models.Appointment)
	if err := ctx.Bind().Body(appointment); err != nil {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(err)
	}
	if ID != appointment.AppointmentID {
		return ctx.Status(fiber.StatusBadRequest).JSON("id on path doesn't match id on body")
	}
	appointment, err := s.DB.UpdateAppointment(ctx.Context(), appointment)
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
	return ctx.Status(fiber.StatusOK).JSON(appointment)
}

func (s *EchoServer) DeleteAppointment(ctx fiber.Ctx) error {
	ID := ctx.Params("id")
	err := s.DB.DeleteAppointment(ctx.Context(), ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusResetContent).JSON(nil)
}
