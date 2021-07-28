package controllers

import (
	"github.com/bulters/NerdKit/application"
	"github.com/gin-gonic/gin"
)

// Attachable defines the interface to which all controllers
// need to adhere to.
type Attachable interface {
	Setup(a *application.BaseApplication) error
	Routes(s *gin.RouterGroup) error
}
