package db

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) listProvinces(ctx *gin.Context) {
	provinces, err := server.store.ListProvinces(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, provinces)
}

type ListProvinceDistrictRequest struct {
	ID int32 `uri:"id" binding:"required,min=10,max=96"`
}

func (server *Server) listProvinceDistricts(ctx *gin.Context) {

	req := ListProvinceDistrictRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	districts, err := server.store.ListProvincesDistrict(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, districts)
}

type ListDistrictTambonRequest struct {
	ID int32 `uri:"id" binding:"required,min=1001,max=9613"`
}

func (server *Server) listDistrictTambon(ctx *gin.Context) {

	req := ListDistrictTambonRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	districts, err := server.store.ListDistrictTambon(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, districts)
}

type ListTambonPostcodeRequest struct {
	ID int32 `uri:"id" binding:"required,min=100101,max=961303"`
}

func (server *Server) listTambonPostcode(ctx *gin.Context) {
	req := ListTambonPostcodeRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postcode, err := server.store.ListTambonPostcode(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, postcode)

}
