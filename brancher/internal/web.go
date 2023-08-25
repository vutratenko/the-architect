package internal

import (
    "net/http" 
    "github.com/labstack/echo/v4"
)

func saveUser(c echo.Context) error{
    return c.String(http.StatusOK, "Saved!")
}

func getUser(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}

func updateUser(c echo.Context) error{
    id := c.Param("id")
    return c.String(http.StatusOK, "Updated! " + id )
}

func deleteUser(c echo.Context) error{
    id := c.Param("id")
    return c.String(http.StatusOK, "Deleted! " + id)
}

// RunServer - runs Echo web server
func RunServer()  {
	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.POST("/users", saveUser)
    e.GET("/users/:id", getUser)
    e.PUT("/users/:id", updateUser)
    e.DELETE("/users/:id", deleteUser)
    e.GET("/livez", liveness)
    e.GET("/readyz", readiness)

    e.Logger.Fatal(e.Start(":1323"))
}