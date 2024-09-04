package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *GinServer) GetAllPets(ctx *gin.Context) {
	pets, err := s.DB.GetAllPets(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, pets)
}

func (s *GinServer) AddPet(ctx *gin.Context) {
	pet := new(models.Pet)
	if err := ctx.BindJSON(pet); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	pet, err := s.DB.AddPet(ctx.Request.Context(), pet)
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
	ctx.IndentedJSON(http.StatusCreated, pet)
}

func (s *GinServer) GetPetById(ctx *gin.Context) {
	ID := ctx.Param("id")
	pet, err := s.DB.GetPetById(ctx.Request.Context(), ID)
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
	ctx.IndentedJSON(http.StatusOK, pet)
}

func (s *GinServer) UpdatePet(ctx *gin.Context) {
	ID := ctx.Param("id")
	pet := new(models.Pet)
	if err := ctx.Bind(pet); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	if ID != pet.PetID {
		ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
		return
	}
	pet, err := s.DB.UpdatePet(ctx.Request.Context(), pet)
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
	ctx.IndentedJSON(http.StatusOK, pet)
}

func (s *GinServer) DeletePet(ctx *gin.Context) {
	ID := ctx.Param("id")
	err := s.DB.DeletePet(ctx.Request.Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusResetContent, nil)
}
