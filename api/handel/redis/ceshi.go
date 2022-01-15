package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/qinsheng99/goWeb/api/entity/redis/request"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/service/ceshi"
	"github.com/qinsheng99/goWeb/library/funcTest"
	"github.com/qinsheng99/goWeb/library/redisClient"
	timeFun "github.com/qinsheng99/goWeb/library/time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v8"
	pb "github.com/qinsheng99/example/grpc-example/route"
	"google.golang.org/grpc"
)

type Handle struct {
	c  ceshi.CeShiService
	ri redisClient.RedisInterface
}

var ctx = context.Background()

func NewH(c ceshi.CeShiService, ri redisClient.RedisInterface) *Handle {
	return &Handle{
		c:  c,
		ri: ri,
	}
}

func (h *Handle) SetR(c *gin.Context) {
	var re redisClient.RE
	err := c.ShouldBindWith(&re, binding.Query)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.Set(ctx, re.Name, re.Data, time.Minute*5)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) GetR(c *gin.Context) {
	var re redisClient.RE
	err := c.ShouldBindWith(&re, binding.Query)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.Get(ctx, re.Name)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) ExistsR(c *gin.Context) {
	var re redisClient.RE
	err := c.ShouldBindWith(&re, binding.Query)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.Exists(ctx, re.Name)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) DelR(c *gin.Context) {
	var re redisClient.RE
	err := c.ShouldBindWith(&re, binding.Query)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.Del(ctx, re.Name)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) Hset(c *gin.Context) {
	var re request.Hset
	err := c.ShouldBindWith(&re, binding.JSON)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.HSet(ctx, re.Name, re.Field, re.Data)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) HGetOrAll(c *gin.Context) {
	var re request.Hset
	err := c.ShouldBindWith(&re, binding.JSON)
	if err != nil {
		common.QueryFailure(c, err)
		return
	}
	if re.Bo == false {
		all, err := h.ri.HGetAll(ctx, re.Name)
		if err != nil {
			common.Failure(c, err)
			return
		}
		common.Success(c, all)
		return
	}
	get, err := h.ri.HGet(ctx, re.Name, re.Field)

	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, get)

}

func (h *Handle) Grpc(c *gin.Context) {
	coon, err := grpc.Dial("localhost:5001", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		common.Failure(c, err)
	}
	defer coon.Close()

	client := pb.NewRouteGuideClient(coon)
	//feature, err := funcTest.RunFirst(client)
	//if err != nil {
	//	common.Failure(c, err)
	//}
	//common.Success(c, feature)

	//steam, err := funcTest.RunSec(client)
	//
	//if err != nil {
	//	common.Failure(c, err)
	//}
	//common.Success(c, steam)

	//recv, err := funcTest.RunThird(client)
	//
	//if err != nil {
	//	common.Failure(c, err)
	//}
	//
	//common.Success(c, recv)
	common.Success(c, client)
	_, err = funcTest.RunForth(client, c)
	if err != nil {
		common.Failure(c, err)
	}

}

func (h *Handle) Sadd(c *gin.Context) {
	for i := 0; i < 10; i++ {
		now := time.Now().Unix()
		timeString := timeFun.TimeIntToString(now)
		_, err := h.ri.SAdd(ctx, "ceshi", timeString)
		if err != nil {
			common.Failure(c, err)
			return
		}
		time.Sleep(time.Second)
	}

	common.Success(c, "")
}

func (h *Handle) SMembers(c *gin.Context) {
	res, err := h.ri.SMembers(ctx, "ceshi")
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, res)
}

func (h *Handle) SRandMember(c *gin.Context) {
	var count request.RandMember
	if err := c.ShouldBindWith(&count, binding.Query); err != nil {
		common.QueryFailure(c, err)
		return
	}
	res, err := h.ri.SRandMemberN(ctx, "ceshi", count.Count)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, res)
}

func (h *Handle) Llen(c *gin.Context) {
	res, err := h.ri.Llen(ctx, "ceshi")
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, res)
}

func (h *Handle) Lpush(c *gin.Context) {
	var re redisClient.RE
	if err := c.ShouldBindWith(&re, binding.Query); err != nil {
		common.QueryFailure(c, err)
		return
	}
	if re.PushType == "1" {
		res, err := h.ri.Rpush(ctx, "ceshi", re.Data)
		if err != nil {
			common.Failure(c, err)
			return
		}
		common.Success(c, res)
		return
	}
	res, err := h.ri.Lpush(ctx, "ceshi", re.Data)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) LRange(c *gin.Context) {
	var re redisClient.RE
	if err := c.ShouldBindWith(&re, binding.Query); err != nil {
		common.QueryFailure(c, err)
		return
	}
	res, err := h.ri.LRange(ctx, "ceshi", re.Start, re.Stop)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, res)
}

func (h *Handle) Lpop(c *gin.Context) {
	var re redisClient.RE
	if err := c.ShouldBindWith(&re, binding.Query); err != nil {
		common.QueryFailure(c, err)
		return
	}
	if re.PushType == "1" {
		res, err := h.ri.RPop(ctx, "ceshi")
		if err != nil {
			common.Failure(c, err)
			return
		}
		common.Success(c, res)
		return
	}
	res, err := h.ri.LPop(ctx, "ceshi")
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) Dump(c *gin.Context) {
	var re redisClient.RE
	if err := c.ShouldBindWith(&re, binding.Query); err != nil {
		common.QueryFailure(c, err)
		return
	}
	b, err := h.ri.Dump(ctx, re.Name)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, b)
}

func (h *Handle) MSet(c *gin.Context) {
	res, err := h.ri.MSet(ctx, "fruit", "apple", "drink", "beer", "food", "cookies")
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) MGet(c *gin.Context) {
	res, err := h.ri.MGet(ctx, "fruit", "food")
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) Zadd(c *gin.Context) {
	var a = []int{123, 456, 789, 100, 23}
	var b []*redis.Z
	var score = 5000

	for k, v := range a {
		b = append(b, &redis.Z{
			Score:  float64(score - k),
			Member: "店铺号:" + strconv.Itoa(v),
		})
	}
	res, err := h.ri.Zadd(ctx, "score", b...)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) Zrange(c *gin.Context) {
	res, err := h.ri.ZRevrange(ctx, "score", 0, -1)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (h *Handle) Zrank(c *gin.Context) {
	res, err := h.ri.ZRank(ctx, "score", "店铺号:456")
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}
