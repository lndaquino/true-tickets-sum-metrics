package routers

import (
	"github.com/labstack/echo"
	"github.com/lndaquino/true-tickets-sum-metrics/api/controllers"
)

// SystemRoutes struct models a system level router
type SystemRoutes struct {
	metricController *controllers.MetricController
}

// NewSystemRoutes returns a SystemRoutes instance
func NewSystemRoutes(c *controllers.MetricController) *SystemRoutes {
	return &SystemRoutes{
		metricController: c,
	}
}

// MakeControllers setups the app routes
func (routes *SystemRoutes) MakeControllers() *echo.Echo {
	e := echo.New()

	e.GET("/metric/:key/sum", routes.metricController.Get)
	e.POST("/metric/:key", routes.metricController.Create)

	return e
}
