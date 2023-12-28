package handler

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"wanderer/config"
	"wanderer/features/bookings"
	"wanderer/helpers/filters"
	"wanderer/helpers/tokens"
	"wanderer/utils/files"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jung-kurt/gofpdf"
	echo "github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func NewBookingHandler(bookingService bookings.Service, jwtConfig config.JWT, cloud files.Cloud) bookings.Handler {
	return &bookingHandler{
		bookingService: bookingService,
		jwtConfig:      jwtConfig,
		cloud:          cloud,
	}
}

type bookingHandler struct {
	bookingService bookings.Service
	jwtConfig      config.JWT
	cloud          files.Cloud
}

func (hdl *bookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var baseUrl = c.Scheme() + "://" + c.Request().Host

		var pagination = new(filters.Pagination)
		c.Bind(pagination)
		if pagination.Start != 0 && pagination.Limit == 0 {
			pagination.Limit = 5
		}

		var search = new(filters.Search)
		c.Bind(search)

		var sort = new(filters.Sort)
		c.Bind(sort)

		result, totalData, err := hdl.bookingService.GetAll(context.Background(), filters.Filter{Pagination: *pagination, Sort: *sort})
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []BookingResponse
		for _, booking := range result {
			var tmpBooking = new(BookingResponse)
			tmpBooking.FromEntity(booking)

			tmpBooking.User.Image = ""

			data = append(data, *tmpBooking)
		}
		response["data"] = data

		if pagination.Limit != 0 {
			var paginationResponse = make(map[string]any)
			if pagination.Start >= (pagination.Limit) {
				prev := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start-pagination.Limit, pagination.Limit)
				if search.Keyword != "" {
					prev += "&keyword=" + search.Keyword
				}
				paginationResponse["prev"] = prev
			} else {
				paginationResponse["prev"] = nil
			}

			if totalData > pagination.Start+pagination.Limit {
				next := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start+pagination.Limit, pagination.Limit)
				if search.Keyword != "" {
					next += "&keyword=" + search.Keyword
				}
				paginationResponse["next"] = next
			} else {
				paginationResponse["next"] = nil
			}
			response["pagination"] = paginationResponse
		}

		response["message"] = "get all tour success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *bookingHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		bookingCode, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid booking code"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.GetDetail(c.Request().Context(), bookingCode)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		if result != nil {
			var data = new(BookingResponse)
			data.FromEntity(*result)

			response["data"] = data
		}

		response["message"] = "get detail booking success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *bookingHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(BookingCreateRequest)

		token := c.Get("user")
		if token == nil {
			response["message"] = "unauthorized access"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, err := tokens.ExtractToken(hdl.jwtConfig.Secret, token.(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.Create(c.Request().Context(), request.ToEntity(userId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data = new(BookingResponse)
		data.FromEntity(*result)

		response["message"] = "create booking success"
		response["data"] = data
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *bookingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(BookingUpdateRequest)

		bookingCode, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid booking code"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		if request.Bank != "" {
			result, err := hdl.bookingService.ChangePaymentMethod(c.Request().Context(), bookingCode, request.ToEntity().Payment)
			if err != nil {
				c.Logger().Error(err)

				if strings.Contains(err.Error(), "validate: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
					return c.JSON(http.StatusBadRequest, response)
				}

				if strings.Contains(err.Error(), "not found: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
					return c.JSON(http.StatusNotFound, response)
				}

				if strings.Contains(err.Error(), "unprocessable: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "unprocessable: ", "")
					return c.JSON(http.StatusUnprocessableEntity, response)
				}

				response["message"] = "internal server error"
				return c.JSON(http.StatusInternalServerError, response)
			}

			var data = new(BookingResponse)
			data.FromEntity(bookings.Booking{Total: result.BookingTotal, Payment: *result})

			response["message"] = "change payment method success"
			response["data"] = data
		} else if request.Status != "" {
			if err := hdl.bookingService.UpdateBookingStatus(c.Request().Context(), bookingCode, request.Status); err != nil {
				c.Logger().Error(err)

				if strings.Contains(err.Error(), "validate: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
					return c.JSON(http.StatusBadRequest, response)
				}

				if strings.Contains(err.Error(), "not found: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
					return c.JSON(http.StatusNotFound, response)
				}

				if strings.Contains(err.Error(), "unprocessable: ") {
					response["message"] = strings.ReplaceAll(err.Error(), "unprocessable: ", "")
					return c.JSON(http.StatusUnprocessableEntity, response)
				}

				response["message"] = "internal server error"
				return c.JSON(http.StatusInternalServerError, response)
			}

			if request.Status == "cancel" {
				response["message"] = "cancel success"
			} else if request.Status == "refund" {
				response["message"] = "refund requested"
			} else if request.Status == "refunded" {
				response["message"] = "approve refund success"
			}
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *bookingHandler) PaymentNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request = new(PaymentNotificationRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			return c.JSON(http.StatusBadRequest, "bad request")
		}

		code, err := strconv.Atoi(request.Code)
		if err != nil {
			c.Logger().Error(err)

			return c.JSON(http.StatusBadRequest, "bad request")
		}

		if err = hdl.bookingService.UpdatePaymentStatus(c.Request().Context(), code, request.Status); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				return c.JSON(http.StatusBadRequest, strings.ReplaceAll(err.Error(), "validate: ", ""))
			}

			if strings.Contains(err.Error(), "unprocessable: ") {
				return c.JSON(http.StatusBadRequest, strings.ReplaceAll(err.Error(), "unprocessable: ", ""))
			}

			return c.JSON(http.StatusInternalServerError, "internal server error")
		}

		return c.JSON(http.StatusOK, "ok")
	}
}

func (hdl *bookingHandler) ExportFileCsv(data []ExportFileResponse, c echo.Context) error {
	path := "transaction-list.csv"

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Booking Code", "Name", "Tour Package", "Duration", "Price", "Status"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, booking := range data {
		row := []string{
			strconv.FormatInt(int64(booking.Code), 10),
			booking.User.Name,
			booking.Tour.Title,
			strconv.FormatInt(int64(booking.Tour.Duration), 10),
			strconv.FormatInt(int64(booking.Total), 10),
			booking.Status,
		}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = hdl.cloud.Upload(context.Background(), "csv-folder", file)
	if err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=transaction-list.csv")

	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(c.Response().Writer, file)
	if err != nil {
		return err
	}

	return nil
}

func (hdl *bookingHandler) ExportFileExcel(data []ExportFileResponse, c echo.Context) error {
	path := "transaction-list.xlsx"

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	xlsx := excelize.NewFile()

	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	style, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#ffc430"}, Pattern: 1},
	})
	err = xlsx.SetCellStyle("Sheet1", "A1", "F1", style)

	headers := []string{"Booking Code", "Name", "Tour Package", "Duration", "Price", "Status"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col)
		xlsx.SetCellValue(sheetName, cell, header)
	}

	for row, booking := range data {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), strconv.FormatInt(int64(booking.Code), 10))
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), booking.User.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), booking.Tour.Title)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", row+2), booking.Tour.Duration)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", row+2), booking.Total)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", row+2), booking.Status)
	}

	err = xlsx.SaveAs(path)
	if err != nil {
		panic(err)
	}

	_, err = hdl.cloud.Upload(context.Background(), "excel-folder", file)
	if err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/xlsx")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=transaction-list.xlsx")

	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(c.Response().Writer, file)
	if err != nil {
		return err
	}

	return nil
}

func (hdl *bookingHandler) ExportFilePdf(data []ExportFileResponse, c echo.Context) error {
	path := "transaction-list.pdf"

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)
	pdf.SetFontSize(11)
	pdf.SetFillColor(255, 196, 48)

	headers := []string{"Booking Code", "Name", "Tour Package", "Duration", "Price", "Status"}
	for _, header := range headers {
		pdf.CellFormat(30, 10, header, "1", 0, "C", true, 0, "")
	}

	for _, booking := range data {
		pdf.Ln(-1)
		pdf.CellFormat(30, 10, strconv.FormatInt(int64(booking.Code), 10), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, booking.User.Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, booking.Tour.Title, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, strconv.FormatInt(int64(booking.Tour.Duration), 10), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, strconv.FormatInt(int64(booking.Total), 10), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, booking.Status, "1", 0, "C", false, 0, "")
	}

	err = pdf.OutputFileAndClose(path)
	if err != nil {
		panic(err)
	}

	_, err = hdl.cloud.Upload(context.Background(), "pdf-folder", file)
	if err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=transaction-list.pdf")

	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(c.Response().Writer, file)
	if err != nil {
		return err
	}

	return nil
}

func (hdl *bookingHandler) ExportReportTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		result, err := hdl.bookingService.Export(context.Background())
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []ExportFileResponse
		for _, export := range result {
			var tmpExport = new(ExportFileResponse)
			tmpExport.FromEntity(export)

			tmpExport.User.Image = ""

			data = append(data, *tmpExport)
		}

		export := c.QueryParam("type")
		if export == "csv" {
			err = hdl.ExportFileCsv(data, c)
			if err != nil {
				c.Logger().Error(err)
				response["message"] = "Error exporting data"
				return c.JSON(http.StatusInternalServerError, response)
			}

			err = os.Remove("transaction-list.csv")
			if err != nil {
				return err
			}

			return nil
		}

		if export == "excel" {
			err = hdl.ExportFileExcel(data, c)
			if err != nil {
				c.Logger().Error(err)
				response["message"] = "Error exporting data"
				return c.JSON(http.StatusInternalServerError, response)
			}

			err = os.Remove("transaction-list.xlsx")
			if err != nil {
				return err
			}

			return nil
		}

		if export == "pdf" {
			err = hdl.ExportFilePdf(data, c)
			if err != nil {
				c.Logger().Error(err)
				response["message"] = "Error exporting data"
				return c.JSON(http.StatusInternalServerError, response)
			}

			err = os.Remove("transaction-list.pdf")
			if err != nil {
				return err
			}

			return nil
		}

		response["message"] = "export transaction list success"
		return c.JSON(http.StatusOK, response)
	}
}
