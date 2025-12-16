package httpserver

import (
	"net/http"

	"github.com/HosseinForouzan/user-management/service/userservice"
	"github.com/labstack/echo/v4"
)

func (s Server) UserRegister(c echo.Context) error {
	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.UserSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
