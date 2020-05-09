package controllers

import (
	"fmt"
	"go_simple_rest/models"
	"net/http"

	"github.com/labstack/echo"
)

// BuildResponse - creates a response object out of data given
func BuildResponse(status string, error string, data interface{}) echo.Map {
	response := echo.Map{
		"status": status,
	}
	if error != "" {
		response["error"] = error
		return response
	}

	response["data"] = data
	return response
}

// GetStatus provides rest endpoint to check the status of the rest server
func GetStatus(c echo.Context) error {
	response := BuildResponse("success", "", "Server is UP and Running")
	return c.JSON(http.StatusOK, response)
}

func TestGet(c echo.Context) error {
	data := make(map[string]string)

	test_model := models.TestStruct{}

	if err := c.Bind(&test_model); err != nil {
		fmt.Println("Error in getting json params from the request")
	}

	test_query_param := c.QueryParam("test_query_param")

	data["message"] = "This is a test GET"
	data["test_arg"] = test_model.Test_param
	data["test_query_param"] = test_query_param

	response := BuildResponse("success", "", data)
	return c.JSON(http.StatusOK, response)
}
