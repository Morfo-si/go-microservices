package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *GinServer) GetAllAppointments(ctx *gin.Context) {
	appointmentId := ctx.Param("appointmentId")

	appointments, err := s.DB.GetAllAppointments(ctx.Request.Context(), appointmentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, appointments)
}

func (s *GinServer) AddAppointment(ctx *gin.Context) {
	appointment := new(models.Appointment)
	if err := ctx.Bind(appointment); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	appointment, err := s.DB.AddAppointment(ctx.Request.Context(), appointment)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			ctx.JSON(http.StatusConflict, err)
			return
		default:
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.IndentedJSON(http.StatusCreated, appointment)
}

func (s *GinServer) GetAppointmentById(ctx *gin.Context) {
	ID := ctx.Param("id")
	appointment, err := s.DB.GetAppointmentById(ctx.Request.Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			ctx.JSON(http.StatusNotFound, err)
			return
		default:
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.IndentedJSON(http.StatusOK, appointment)
}

func (s *GinServer) UpdateAppointment(ctx *gin.Context) {
	ID := ctx.Param("id")
	appointment := new(models.Appointment)
	if err := ctx.Bind(appointment); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	if ID != appointment.AppointmentID {
		ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
		return
	}
	appointment, err := s.DB.UpdateAppointment(ctx.Request.Context(), appointment)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			ctx.JSON(http.StatusNotFound, err)
			return
		case *dberrors.ConflictError:
			ctx.JSON(http.StatusConflict, err)
			return
		default:
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.IndentedJSON(http.StatusOK, appointment)
}

func (s *GinServer) DeleteAppointment(ctx *gin.Context) {
	ID := ctx.Param("id")
	err := s.DB.DeleteAppointment(ctx.Request.Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusResetContent, nil)
}
