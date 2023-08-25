package internal

import (
    "net/http" 
    "github.com/labstack/echo/v4"
)


func liveness(c echo.Context) error{
	l := &livenessProbeResponse{
		Code: 200,
		Response: "OK",
	}
	return c.JSON(http.StatusOK, l)
}

func readiness(c echo.Context) error{
	r := readinessProbeResponse{
		Code: 200,
		Response: "OK",
		Postgres: "OK",
		Rabbitmq: "OK",
	}
	return c.JSON(http.StatusOK, r)
}