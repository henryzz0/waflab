// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

type Testset struct {
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Desc      string   `xorm:"varchar(100)" json:"desc"`
	TargetUrl string   `xorm:"varchar(100)" json:"targetUrl"`
	Testcases []string `xorm:"mediumtext" json:"testcases"`
}

func GetTestsets() []*Testset {
	testsets := []*Testset{}
	err := ormManager.engine.Desc("created_time").Find(&testsets)
	if err != nil {
		panic(err)
	}

	return testsets
}

func GetTestset(id string) *Testset {
	s := Testset{Name: id}
	existed, err := ormManager.engine.Get(&s)
	if err != nil {
		panic(err)
	}

	if existed {
		return &s
	} else {
		return nil
	}
}

func UpdateTestset(id string, testset *Testset) bool {
	if GetTestset(id) == nil {
		return false
	}

	_, err := ormManager.engine.Id(id).AllCols().Update(testset)
	if err != nil {
		panic(err)
	}

	//return affected != 0
	return true
}

func AddTestset(testset *Testset) bool {
	affected, err := ormManager.engine.Insert(testset)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteTestset(testset *Testset) bool {
	affected, err := ormManager.engine.Id(testset.Name).Delete(&Testset{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
