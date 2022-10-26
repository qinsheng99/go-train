package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qinsheng99/go-train/internal/model"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bys() []byte {
	body := `
		{
			"configurations":{
                "nodes":[
                    {
                        "operator":"OR",
                        "cpe":[
                            {
                                "cpe23Uri":"cpe:2.3:a:microsoft:ie:6:windows_server_2003_sp1:*:*:*:*:*:*",
                                "cpeMatchString":"cpe:/a:microsoft:ie:6:windows_server_2003_sp1",
                                "vulnerable":"true"
                            }
                        ]
                    },
					{
                        "operator":"OR",
                        "cpe":[
                            {
                                "cpe23Uri":"cpe:2.3:a:microsoft:ie:6:windows_server_2003_sp1:*:*:*:*:*:*",
                                "cpeMatchString":"cpe:/a:microsoft:ie:6:windows_server_2003_sp1",
                                "vulnerable":"true"
                            }
                        ]
                    }
                ]
            }
		}
		`

	var o origin
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		log.Fatalln(err)
	}

	n := o.Configurations.Nodes

	marshal, _ := json.Marshal(n)
	return marshal
}
func TestPostgreSqlCreate(t *testing.T) {

	u, _ := uuid.NewRandom()
	if err := d.Exec(d.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&model.Boy{
			Name: "xiaohu14",
			//Informations: postgres.Jsonb{RawMessage: []byte(string(bys()))},
			Arr:  []int64{14, 28, 42},
			UUid: u,
		})
	})).Error; err != nil {
		t.Fatal(err)
	}
}

func TestPostgreSqlUpdate(t *testing.T) {
	b := &model.Boy{
		Informations: postgres.Jsonb{RawMessage: bys()},
	}
	t.Log(d.Exec(d.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(b).Where("id = 12").Updates(b)
	})).Error)
}

type origin struct {
	Configurations CveConfigurations `json:"configurations"`
}

type ConfNodes struct {
	Operator string    `json:"operator"`
	Cpe      []NodeCpe `json:"cpe"`
}

type CveConfigurations struct {
	Nodes []ConfNodes `json:"nodes"`
}

type NodeCpe struct {
	Cpe23Uri       string `json:"cpe23Uri"`
	CpeMatchString string `json:"cpeMatchString"`
	Vulnerable     string `json:"vulnerable"`
}

func TestPostgreSqlGet(t *testing.T) {
	b := &model.Boy{}
	err := d.Model(b).Where("id = ?", 13).First(b).Error
	if err != nil {
		t.Fatal(err)
	}

	var n []ConfNodes
	err = json.Unmarshal(b.Informations.RawMessage, &n)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(n))
}

var d *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "postgres", "root", "postgres", 5432)
	db, err := gorm.Open(p.New(p.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return
	}
	d = db
	m.Run()
}
