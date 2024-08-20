package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zxylon/olaf-layout-basic/internal/service"
	"github.com/zxylon/olaf-layout-basic/pkg/helper/resp"
	"go.uber.org/zap"
)

func NewUserHandler(handler *Handler,
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

type UserHandler struct {
	*Handler
	userService service.UserService
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
