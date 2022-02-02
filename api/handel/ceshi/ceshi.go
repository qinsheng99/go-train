package ceshi

import (
	"encoding/json"
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/entity"
	crequest "github.com/qinsheng99/goWeb/api/entity/ceshi/request"
	drequest "github.com/qinsheng99/goWeb/api/entity/drainage/request"
	"github.com/qinsheng99/goWeb/api/tools/common"
	Err "github.com/qinsheng99/goWeb/err"
	"github.com/qinsheng99/goWeb/internal/dao/idao"
	"github.com/qinsheng99/goWeb/internal/dao/idao/customer"
	"github.com/qinsheng99/goWeb/internal/service/drainage"
	"github.com/qinsheng99/goWeb/library/pool"
	httprequest "github.com/qinsheng99/goWeb/library/request"
	"github.com/qinsheng99/hello"
)

type Handler struct {
	customer customer.CustomerImp
	esImp    idao.EsImp
	drainage drainage.Drainage
}

func NewHandler(customer customer.CustomerImp, esImp idao.EsImp, drainage drainage.Drainage) *Handler {
	return &Handler{
		customer: customer,
		esImp:    esImp,
		drainage: drainage,
	}
}

func (h *Handler) GetEs(c *gin.Context) {
	var q crequest.CeShiGetRequest
	if err := c.ShouldBind(&q); err != nil {
		common.Failure(c, err)
		return
	}
	res, err := h.esImp.GetEsList(q)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, gin.H{
		"res":   res,
		"q":     q,
		"hello": hello.Hello(),
	})

}

func (h *Handler) GetEsList(c *gin.Context) {
	var q crequest.CeShiGetRequest
	if err := c.ShouldBind(&q); err != nil {
		common.Failure(c, err)
		return
	}
	res, err := h.esImp.Get(q)

	if err != nil {
		common.Failure(c, err)
	}

	common.Success(c, res)

}
func (h *Handler) Index(c *gin.Context) {
	//bytedata := make(map[string]interface{})
	//bytedata["uri"] = "123"
	//byte1, err := json.Marshal(bytedata)
	//if err != nil {
	//	common.Failure(c, err)
	//}
	//byte2 := bytes.NewReader(byte1)
	//common.Success(c, hello.Hello())
	//var q crequest.CeShiRequest
	//if err := c.ShouldBindWith(&q, binding.JSON); err != nil && err != io.EOF {
	//	common.Failure(c, err)
	//	return
	//}
	////res, err := h.esImp.PostEsList(q)
	////res, err := h.esImp.Delete(136)
	//res, err := h.esImp.GetEsById(q) // 1148071  1148045
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	p := pool.NewGoPool(pool.WithMaxLimit(2))
	for i := 0; i < 20; i++ {
		p.Submit(func() {
		})
	}
	defer p.Close()
	//defer pool.GoFunc.Wait()
	common.Success(c, gin.H{})
	// var a = []int{1, 2, 3, 4}
	// if err != nil {
	// 	common.Failure(c, err)
	// 	return
	// }

	// common.Success(c, res)
	// h.c.GetByIds()

	// bodyData := make(map[string]interface{})
	// bodyData["uri"] = "drainage2108160938evv76qf"
	// bodyData["type"] = 1
	// byteData, err := json.Marshal(bodyData)
	// if err != nil {
	// 	common.Failure(c, err)
	// 	return
	// }
	// common.Success(c, gin.H{
	// 	"byte": byteData,
	// 	"io":   bytes.NewReader(byteData),
	// })

	// scope := func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("id > ? ", 1)
	// }
	// var customerFollowerUsers []model.CustomerFollowUser
	// query := h.Db.Db.
	// 	Scopes(scope).Model(model.CustomerFollowUser{}).
	// 	Limit(10).
	// 	Find(&customerFollowerUsers)
	// // var data []model.CustomerFollowUserList
	// // query := h.Db.Db.Model(&model.CustomerFollowUser{}).Where("id > ?", 10).Limit(10).Order("id asc").Find(&data)
	// if err := query.Error; err != nil {
	// 	common.Failure(c, err)
	// 	return
	// }
	// common.Success(c, customerFollowerUsers)
	// var d response.ShardFailure
	// if err := c.ShouldBindWith(&d, binding.JSON); err != nil {
	// 	common.QueryFailure(c, err)
	// 	return
	// }
	// common.Success(c, os.Getenv("NIFFLER_ADDRS"))
	// value := c.Request.URL.Query()
	// value.Set("asd", strconv.Itoa(13))
	// value.Set("sa223d", strconv.Itoa(13))
	// value.Set("a321sdsd", strconv.Itoa(13))
}

func (h *Handler) DeleteEs(c *gin.Context) {
	err := h.esImp.DeleteEs()

	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, gin.H{})
}

func (h *Handler) CreateEs(c *gin.Context) {
	err := h.customer.Refresh()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, gin.H{})
}

func (h *Handler) GetList(c *gin.Context) {
	res, err := h.customer.GetList()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handler) GetDrainageList(c *gin.Context) {
	var d drequest.DrainageRequest
	if err := c.ShouldBind(&d); err != nil {
		common.QueryFailure(c, err)
		return
	}
	res, err := h.drainage.GetListForDrainage(d)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handler) Http(c *gin.Context) {
	//url := "http://localhost:111/public/index"
	//bodyData := make(map[string]interface{})
	//bodyData["uri"] = "drainage2108242100ldys2u6"
	//byteData, err := json.Marshal(bodyData)
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	//res, err := http.NewRequest("POST", url, bytes.NewReader(byteData))
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	////res.Header.Set("Content-Type", "application/json")
	////res.Header.Set("Cookie", "cft-userInfo="+"6cA90jBT1sLu6%2BnJyV9C7t2dM%2FUtQOCmOy74RSrgnhQ"+"FSxNLjFQ2AY77cvkwLpAQyK78PIetKRzHlHcS7gKHaF93cRp"+"h6XcLD1mqWmKrKAj6R7K9ogsaVmHjgZSAu9g6um3BZ2%2F0yfb"+"XTF05Kynw5cOUVK1H4P994jVuFUxrhBXcn2DGEAYtWxbxAey"+"Dwpf97Rz0OvZNBek3vviGhzADgy1yEKUcrbT3A4RBHdtsbXDJR%2"+"FL7nfOCcsCh2aZovVqoo1fuYgjoBQ4emDWPljmPABdCvbHF%2BO"+"oG1C%2Bw8Jby6WU%3D")
	//resp, err := client.Do(res)
	//if err != nil || resp == nil {
	//	common.Failure(c, err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//resByte, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	//var data dresponse.Response
	//err = json.Unmarshal(resByte, &data)
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	//common.Success(c, data)
	url := "http://localhost:8000/index"
	bodyData := make(map[string]interface{})
	bodyData["name"] = "17339911151"
	byteData, err := json.Marshal(bodyData)
	if err != nil {
		common.Failure(c, err)
		return
	}
	request, err := httprequest.Get(url, byteData, nil)
	if err != nil {
		common.Failure(c, err)
		return
	}
	var data entity.Response

	err = json.Unmarshal(request, &data)
	if err != nil {
		common.Failure(c, err)
	}
	if data.Code != int64(0) {
		common.Failure(c, Err.Mesage(data.Msg))
		return
	}
	common.Success(c, data)

}

func (h *Handler) LiKou(c *gin.Context) {
	/**
	//todo https://leetcode-cn.com/problems/maximum-gap/
	*/
	// nums := []int{3, 6, 9, 1}
	// if len(nums) < 2 {
	// 	common.Success(c, 0)
	// }
	// var max int = 0
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] > nums[j] {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 		}
	// 	}
	// }
	//
	// for i := 0; i < len(nums)-1; i++ {
	// 	j := i + 1
	// 	num := nums[j] - nums[i]
	// 	fmt.Println(num)
	// 	if num > max {
	// 		max = num
	// 	}
	// }
	// common.Success(c, max)
}
