package ping

import (
	"net/http"

	"github.com/gin-gonic/gin" //  usa para controlador(http, get)
)

type Controller struct {
}

func NewControllerPing() *Controller {
	return &Controller{}
}

//@Summary get ping pong
//@Tags  ping pong
//@Produce json
//success 200

func (c *Controller) HandlerPing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	}
}
