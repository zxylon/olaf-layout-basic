package handler

import (
	"github.com/zxylon/olaf-layout-basic/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
