package demo

import (
	"errors"
)

type VqaRes struct {
	Status    interface{} `json:"status"`
	InferTime float64     `json:"infer_time"`
	Msg       string      `json:"msg"`
	Inference Infer       `json:"inference_result"`
}

func (v *VqaRes) Valication() error {
	status, ok := v.Status.(float64)
	if !ok {
		return nil
	}

	if int(status) == -1 {
		return errors.New(v.Msg)
	}
	return nil
}

type Infer struct {
	Instances string `json:"instances"`
}

type VqaReq struct {
	Image    string `json:"image_path"`
	Question string `json:"question"`
}
