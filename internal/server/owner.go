package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morfo-si/go-microservices/internal/dberrors"
	"github.com/morfo-si/go-microservices/internal/models"
)

func (s *GinServer) GetAllOwners(ctx *gin.Context) {
	emailAddress := ctx.Param("emailAddress")

	owners, err := s.DB.GetAllOwners(ctx.Request.Context(), emailAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, owners)
}

func (s *GinServer) AddOwner(ctx *gin.Context) {
	owner := new(models.Owner)
	if err := ctx.Bind(owner); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	owner, err := s.DB.AddOwner(ctx.Request.Context(), owner)
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
	ctx.IndentedJSON(http.StatusCreated, owner)
}

func (s *GinServer) GetOwnerById(ctx *gin.Context) {
	ID := ctx.Param("id")
	owner, err := s.DB.GetOwnerById(ctx.Request.Context(), ID)
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
	ctx.IndentedJSON(http.StatusOK, owner)
}

func (s *GinServer) UpdateOwner(ctx *gin.Context) {
	ID := ctx.Param("id")
	owner := new(models.Owner)
	if err := ctx.Bind(owner); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
		return
	}
	if ID != owner.OwnerID {
		ctx.JSON(http.StatusBadRequest, "id on path does not match id on body")
		return
	}
	owner, err := s.DB.UpdateOwner(ctx.Request.Context(), owner)
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
	ctx.IndentedJSON(http.StatusOK, owner)
}

func (s *GinServer) DeleteOwner(ctx *gin.Context) {
	ID := ctx.Param("id")
	err := s.DB.DeleteOwner(ctx.Request.Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusResetContent, nil)
}
