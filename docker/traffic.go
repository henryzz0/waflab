package object

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Traffic struct {
	TrafficId     string  `xorm:"text notnull pk"`
	TestTitle     string  `xorm:"text"`
	Meta          string  `xorm:"text"`
	File          string  `xorm:"text"`
	Input         string  `xorm:"text"`
	Output        string  `xorm:"text"`
	Request       string  `xorm:"blob"`
	RawRequest    string  `xorm:"blob"`
	RawResponse   string  `xorm:"blob"`
	RawLog        string  `xorm:"text"`
	TestingResult string  `xorm:"text"`
	DurationTime  float64 `xorm:"real"`
}

func newOrmManager(path string) *xorm.Engine {
	ormManager, err := xorm.NewEngine("sqlite3", path)
	if err != nil {
		panic(err)
	}

	return ormManager
}

func getTraffics(ormManager *xorm.Engine) []*Traffic {
	traffics := []*Traffic{}
	err := ormManager.Find(&traffics)
	if err != nil {
		panic(err)
	}

	return traffics
}

func addTraffic(ormManager *xorm.Engine, traffic *Traffic) bool {
	affected, err := ormManager.Insert(traffic)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
