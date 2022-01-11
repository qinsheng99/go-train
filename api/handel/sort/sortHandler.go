package sortHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/library/sort"
)

type SortHandler struct {}

func NewSort() *SortHandler {
	return &SortHandler{}
}

func (s *SortHandler) SelectSort(c *gin.Context)  {
	var arr = []int{6,3,1,7,2,4,5,8,9}
	sort.SelectSort(arr)
	common.Success(c, arr)
}

func (s *SortHandler) BubblingSort(c *gin.Context)  {
	var arr = []int{6,3,1,7,2,4,5,8,9}
	sort.BubblingSort(arr)
	common.Success(c, arr)
}
