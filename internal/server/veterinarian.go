package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *GinServer) GetAllVeterinarians(ctx *gin.Context) {
	veterinarians, err := s.DB.GetAllVeterinarians(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, veterinarians)
}

func (s *GinServer) AddVeterinarian(ctx *gin.Context) {
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind(veterinarian); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	veterinarian, err := s.DB.AddVeterinarian(ctx.Request.Context(), veterinarian)
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
	ctx.IndentedJSON(http.StatusCreated, veterinarian)
}

func (s *GinServer) GetVeterinarianById(ctx *gin.Context) {
	ID := ctx.Param("id")
	veterinarian, err := s.DB.GetVeterinarianById(ctx.Request.Context(), ID)
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
	ctx.IndentedJSON(http.StatusOK, veterinarian)
}

func (s *GinServer) UpdateVeterinarian(ctx *gin.Context) {
	ID := ctx.Param("id")
	veterinarian := new(models.Veterinarian)
	if err := ctx.Bind(veterinarian); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	if ID != veterinarian.VeterinarianID {
		ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
		return
	}
	veterinarian, err := s.DB.UpdateVeterinarian(ctx.Request.Context(), veterinarian)
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
	ctx.IndentedJSON(http.StatusOK, veterinarian)
}

func (s *GinServer) DeleteVeterinarian(ctx *gin.Context) {
	ID := ctx.Param("id")
	err := s.DB.DeleteVeterinarian(ctx.Request.Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusResetContent, nil)
}
