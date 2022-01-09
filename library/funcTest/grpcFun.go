package funcTest

import (
	"context"
	"gin/api/entity/grpc/response"
	"io"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/qinsheng99/example/grpc-example/route"
)

var sy sync.WaitGroup

func RunFirst(client pb.RouteGuideClient) (*pb.Feature, error) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		return nil, err
	}
	return feature, nil
}

func RunSec(client pb.RouteGuideClient) (a response.GrpcRE, err error) {
	steam, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{Latitude: 313374060, Longitude: 121358540},
		Hi: &pb.Point{Latitude: 311034130, Longitude: 121598790},
	})
	if err != nil {
		return response.GrpcRE{}, err
	}
	for {
		feature, err := steam.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return response.GrpcRE{}, err
		}
		a.Feature = append(a.Feature, feature)
	}
	return a, nil
}

func RunThird(client pb.RouteGuideClient) (*pb.RouteSummary, error) {
	steam, err := client.RecordRoute(context.Background())

	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}
	if err != nil {
		return nil, err
	}
	for _, point := range points {
		err := steam.Send(point)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond)
	}

	recv, err := steam.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	return recv, nil
}

func RunForth(client pb.RouteGuideClient, c *gin.Context) (a response.GrpcRE, err error) {
	steam, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		feature, err := steam.Recv()
		if err != nil {
			log.Fatalln(err)
			return
		}
		a.Feature = append(a.Feature, feature)
	}()
	var r response.ForthRequest
	request := pb.RecommendRequest{
		Point: new(pb.Point),
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		return response.GrpcRE{}, err
	}
	request.Mode = pb.RecommendationMode(r.Mode)
	request.Point.Latitude = r.Latitude
	request.Point.Longitude = r.Longitude

	err = steam.Send(&request)
	if err != nil {
		return response.GrpcRE{}, err
	}
	return
}
