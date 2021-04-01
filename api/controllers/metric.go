package controllers

import (
	"math"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/entity"
)

type MetricController struct {
	usecase MetricUsecase
}

type MetricUsecase interface {
	Get(string) int
	Create(entity.Metric)
}

func NewMetricController(metricUsecase MetricUsecase) *MetricController {
	return &MetricController{
		usecase: metricUsecase,
	}
}

func (ctl *MetricController) Get(c echo.Context) error {
	key := c.Param("key")

	value := ctl.usecase.Get(key)
	return c.JSON(http.StatusOK, map[string]int{
		"value": value,
	})
}

func (ctl *MetricController) Create(c echo.Context) error {
	key := c.Param("key")
	requestBody := new(createBodyRequest)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	value := int(math.Round(requestBody.Value))
	metric := entity.Metric{
		Key:   key,
		Value: value,
	}

	ctl.usecase.Create(metric)

	return c.JSON(http.StatusOK, nil)
}

type createBodyRequest struct {
	Value float64
}
