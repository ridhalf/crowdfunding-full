package helper

import (
	"crowdfunding/model/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Ok(message string, data interface{}) web.WebResponse {
	response := web.APIResponse(message, http.StatusOK, "OK", data)
	return response
}

func BadRequest(message string, data interface{}) web.WebResponse {
	response := web.APIResponse(message, http.StatusBadRequest, "Bad Request", data)
	return response
}
func NotFound(message string) web.WebResponse {
	response := web.APIResponse(message, http.StatusNotFound, "Not Found", nil)
	return response
}
func InternalServerError(message string) web.WebResponse {
	response := web.APIResponse(message, http.StatusInternalServerError, "Internal Server Error", nil)
	return response
}
func UnprocessableEntity(message string, err interface{}) web.WebResponse {

	errors := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	errorMessage := gin.H{"errors": errors}
	response := web.APIResponse(message, http.StatusUnprocessableEntity, "Unprocessable Entity", errorMessage)
	return response
}
func UnprocessableEntityString(message string, err string) web.WebResponse {

	errorMessage := gin.H{"errors": err}
	response := web.APIResponse(message, http.StatusUnprocessableEntity, "Unprocessable Entity", errorMessage)
	return response
}
func UnAuthorized(message string) web.WebResponse {
	response := web.APIResponse(message, http.StatusUnauthorized, "Unauthorized", nil)
	return response
}
func Forbidden(message string, data interface{}) web.WebResponse {
	response := web.APIResponse(message, http.StatusForbidden, "Forbidden", data)
	return response
}
