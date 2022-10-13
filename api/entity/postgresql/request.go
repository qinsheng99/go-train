package postgresqlRequest

type Boy struct {
	Id           int64                  `json:"id" gorm:"column:id"`
	Name         string                 `json:"name" gorm:"column:name"`
	Informations map[string]interface{} `json:"information"`
	Arr          []int64                `json:"arr"`
}
