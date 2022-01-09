package response

import pb "github.com/qinsheng99/example/grpc-example/route"

type GrpcRE struct {
	Feature []*pb.Feature `json:"feature"`
}

type ForthRequest struct {
	Mode      int32 `json:"mode" form:"mode" binding:"required"`
	Latitude  int32 `json:"latitude,omitempty" form:"latitude"`
	Longitude int32 `json:"longitude,omitempty" form:"longitude"`
}
