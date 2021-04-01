package main

import (
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lndaquino/true-tickets-sum-metrics/api/controllers"
	"github.com/lndaquino/true-tickets-sum-metrics/api/routers"
	"github.com/lndaquino/true-tickets-sum-metrics/datastore"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/metric"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/queue"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/worker"
)

func init() {
	godotenv.Load()
}

func main() {
	var e *echo.Echo

	app, worker, err := setupApplication()
	if err != nil {
		panic("Error starting application ==> " + err.Error())
	}

	go worker.Run()

	e = app.MakeControllers()
	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.StartServer(server))
}

func setupApplication() (*routers.SystemRoutes, *worker.Worker, error) {
	metricRepo := datastore.NewInMemoryRepo()
	queue := queue.NewQueue()
	metricUsecase := metric.NewMetricUsecase(metricRepo, queue)
	metricController := controllers.NewMetricController(metricUsecase)
	systemRoutes := routers.NewSystemRoutes(metricController)
	worker := worker.NewWorker(metricRepo, queue)

	return systemRoutes, worker, nil
}
