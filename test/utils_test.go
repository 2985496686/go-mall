package test

import (
	"encoding/json"
	"fmt"
	"mall/utils"
	"testing"
)

type TestUser struct {
	Name string `toMap:"name"`
	Age  int    `toMap:"age"`
	Sex  bool   `toMap:"sex"`
}

func TestOtoMap(t *testing.T) {
	u := TestUser{Name: "张三", Age: 19, Sex: true}
	toMap, _ := utils.StructToMap(u)
	fmt.Println(toMap)
}

func TestJson(t *testing.T) {
	u := TestUser{Name: "张三", Age: 19, Sex: true}
	marshal, _ := json.Marshal(u)
	fmt.Println(string(marshal))
	user := TestUser{}
	json.Unmarshal(marshal, user)
}
