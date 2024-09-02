package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *EchoServer) GetAllAppointments(ctx echo.Context) error {
	appointmentId := ctx.QueryParam("appointmentId")

	appointments, err := s.DB.GetAllAppointments(ctx.Request().Context(), appointmentId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, appointments)
}

func (s *EchoServer) AddAppointment(ctx echo.Context) error {
	appointment := new(models.Appointment)
	if err := ctx.Bind(appointment); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	appointment, err := s.DB.AddAppointment(ctx.Request().Context(), appointment)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, appointment)
}

func (s *EchoServer) GetAppointmentById(ctx echo.Context) error {
	ID := ctx.Param("id")
	appointment, err := s.DB.GetAppointmentById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, appointment)
}

func (s *EchoServer) UpdateAppointment(ctx echo.Context) error {
	ID := ctx.Param("id")
	appointment := new(models.Appointment)
	if err := ctx.Bind(appointment); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != appointment.AppointmentID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}
	appointment, err := s.DB.UpdateAppointment(ctx.Request().Context(), appointment)
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
	return ctx.JSON(http.StatusOK, appointment)
}

func (s *EchoServer) DeleteAppointment(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteAppointment(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
