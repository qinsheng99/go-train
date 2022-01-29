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

// SelectSort 选择排序
func (s *SortHandler) SelectSort(c *gin.Context)  {
	var arr = []int{6,3,1,7,2,4,5,8,9}
	sort.SelectSort(arr)
	common.Success(c, arr)
}

// BubblingSort 冒泡排序
func (s *SortHandler) BubblingSort(c *gin.Context)  {
	var arr = []int{6,3,1,7,2,4,5,8,9}
	sort.BubblingSort(arr)
	common.Success(c, arr)
}

// InsertSort 插入排序
func (s *SortHandler) InsertSort(c *gin.Context)  {
	var arr = []int{6,3,1,7,2,4,5,8,9}
	sort.InsertSort(arr)
	common.Success(c, arr)
}

// ShellSort 希尔排序
func (s *SortHandler) ShellSort(c *gin.Context)  {
	var arr = []int{9,6,11,3,5,12,8,7,10,15,14,4,1,13,2}
	sort.ShellSort(arr)
	common.Success(c, arr)
}

// MergeSort 归并排序
func (s *SortHandler) MergeSort(c *gin.Context)  {
	var arr = []int{9,6,11,3,5,12,8,7,10,15,14,4,1,13,2}
	sort.MergeSort(arr, 0, len(arr) - 1)
	common.Success(c, arr)
}

// QuickSort 快速排序
func (s *SortHandler) QuickSort(c *gin.Context)  {
	var arr = []int{9,6,11,3,5,12,8,7,10,15,14,4,1,13,2}
	sort.QuickSort(arr, 0, len(arr) - 1)
	common.Success(c, arr)
}

// CountSort 计数排序
func (s *SortHandler) CountSort(c *gin.Context)  {
	var arr = []int{2,4,2,3,4,1,1,0,0,5,6,9,8,5,7,4,0,9}
	common.Success(c, sort.CountSort(arr))
}

// RadixSort 基数排序
func (s *SortHandler) RadixSort(c *gin.Context)  {
	var arr = []int{421, 240, 1115, 532, 2305, 430, 4124}
	common.Success(c, sort.RadixSort(arr))
}
