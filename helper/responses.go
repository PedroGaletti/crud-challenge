package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	internalServerError = "Something is wrong in your request, err: %s"
	badRequestBody      = "Something is wrong in your request body, err: %s"
	noContent           = "No records found in the database, err: %s"
	ok                  = "Your request was processed with success"
)

func InternalServerErrorResponse(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, DefaultReponse{Message: fmt.Sprintf(internalServerError, err.Error())})
}

func BadRequestBodyReponse(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, DefaultReponse{Message: fmt.Sprintf(badRequestBody, err.Error())})
}

func NoContentResponse(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusNoContent, DefaultReponse{Message: fmt.Sprintf(noContent, err.Error())})
}

func OkDataResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func OkResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DefaultReponse{Message: ok})
}
