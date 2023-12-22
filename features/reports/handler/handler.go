package handler

import (
	"net/http"
	"wanderer/features/reports"

	echo "github.com/labstack/echo/v4"
)

func NewReportHandler(reportService reports.Service) reports.Handler {
	return &reportHandler{
		reportService: reportService,
	}
}

type reportHandler struct {
	reportService reports.Service
}

func (hdl *reportHandler) Dashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		result, err := hdl.reportService.Dashboard(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		if result != nil {
			var data ReportResponse
			data.FromEntity(*result)

			response["data"] = data
		}
		response["message"] = "get report success"
		return c.JSON(http.StatusOK, response)
	}
}
