package apikit

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GenericResponse[T any] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type Router func(api *API)

type API struct {
	Config  *Configuration
	DB      *sql.DB
	App     *echo.Echo
	routers []Router
}

type CustomValidator struct {
	Validate *validator.Validate
}

func NewAPI(options ...Option) *API {
	api := &API{
		App:     echo.New(),
		routers: []Router{},
		Config:  &Configuration{},
	}

	for _, o := range options {
		o(api)
	}

	return api
}

func (api *API) UseRouter(routers ...Router) *API {
	fmt.Println("Setting up routers...")
	api.routers = append(api.routers, routers...)

	return api
}

func (api *API) Start() {
	if isDBConfigValid(api.Config.db) {
		fmt.Printf("\nConnecting to database...\n")
		dbcf := api.Config.db
		ctx, cancel := context.WithTimeout(context.Background(), dbcf.TimeOut)
		defer cancel()

		db, err := ConnectDBContext(ctx, dbcf)
		if err != nil {
			fmt.Printf("\nError connecting to database: %v\n", err.Error())

			return
		}

		fmt.Printf("\nConnected to database %s\n", dbcf.Name)

		fmt.Println(db)

		api.DB = db
	}

	api.App.Validator = NewValidator()

	fmt.Println("Setting up health check endpoint...")

	api.App.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s, OK!", api.Config.api.Label))
	})

	for _, router := range api.routers {
		router(api)
	}

	fmt.Printf("Starting API %s...\n", api.Config.api.Label)

	api.App.Logger.Fatal(api.App.Start(fmt.Sprintf("%s:%d", api.Config.api.Host, api.Config.api.Port)))
}

func isDBConfigValid(dbcf DBConfig) bool {
	return dbcf.Host != "" && dbcf.Name != "" && dbcf.User != "" && dbcf.Password != ""
}
