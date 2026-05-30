package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

type UserHandler struct {
	client userpb.UserServiceClient
}

func NewUserHandler(client userpb.UserServiceClient) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "invalid user id",
		})
	}

	id := int32(id64)
	resp, err := h.client.GetUser(context.Background(), &userpb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{})
	}

	return c.JSON(http.StatusOK, resp)
}
