package response

import "github.com/qinsheng99/go-train/internal/model"

type DrainageResponse struct {
	Name      string           `json:"name"`
	BackType  int              `json:"backType"`
	BackImage string           `json:"backImage"`
	BrandName string           `json:"brandName"`
	BrandLogo string           `json:"brandLogo"`
	Users     []*model.QyUser1 `json:"users"`
	Tags      []*model.QyTag1  `json:"tags"`
	QrCode    string           `json:"qrCode"`
}

type Response struct {
	Code    int                    `json:"code"`
	Msg     string                 `json:"msg"`
	NowTime int64                  `json:"nowTime"`
	Data    map[string]interface{} `json:"data"`
}

type DrainageResponse1 struct {
	Name       string                   `json:"name"`
	BackType   int                      `json:"backType"`
	BackImage  string                   `json:"backImage"`
	BrandName  string                   `json:"brandName"`
	BrandLogo  string                   `json:"brandLogo"`
	Users      []*model.QyUser1         `json:"users"`
	Tags       []*model.QyTag1          `json:"tags"`
	QrCode     string                   `json:"qrCode"`
	WelcomeMsg []map[string]interface{} `json:"welcomeMsg"`
}
