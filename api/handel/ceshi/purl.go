package ceshi

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-train/api/tools/common"
	"github.com/qinsheng99/go-train/library/purl"
)

func (h *Handler) Purl(c *gin.Context) {
	var req struct {
		Coordinates []string `json:"coordinates"`
	}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.QueryFailure(c, err)
		return
	}

	p, err := purl.PasePurl(req.Coordinates[0])
	if err != nil {
		common.Failure(c, err)
		return
	}
	p["version"], _ = p.GetVersion()

	common.Success(c, p)
}
