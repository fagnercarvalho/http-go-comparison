package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"path"
	"sync"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		e.Logger.Fatal(e.Start(":6122"))

		defer wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		certFile, keyFile := getCerts()
		e.Logger.Fatal(e.StartTLS(":6123", certFile, keyFile))

		defer wg.Done()
	}(&wg)

	wg.Wait()
}

func hello(c echo.Context) error {
	proxy := c.Request().Header["X-Server-Version"]

	if proxy != nil {
		return c.String(http.StatusOK, fmt.Sprintf("Hello from Echo (nginx %s)!", proxy[0]))
	}

	version := c.Request().Proto
	return c.String(http.StatusOK, fmt.Sprintf("Hello from Echo (%s)!", version))
}

func getCerts() (string, string) {
	return path.Join("/go/src/app", "cert.pem"), path.Join("/go/src/app", "key.pem")
}

