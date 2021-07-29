package model

import (
	"fmt"
	"testing"
	"encoding/json"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NestedMap(t *testing.T) {
	Convey("Test_Dao_GetAllEntrances", t, func() {
		m1 := make(map[string]interface{})
		d1 := make(map[string]interface{})
		mb1 := make(map[string]interface{})
		mb2 := make(map[string]interface{})
		mb1["b1"] = map[string]string{"c1": "11", "2": "22"}
		mb1["b2"] = map[string]string{"c2": "12", "2": "22"}
		mb2["b3"] = map[string]string{"c3": "13", "2": "22"}
		mb2["br"] = map[string]string{"c4": "14", "2": "22"}
		m1["a1"] = mb1
		m1["a2"] = mb2

		bytes, _ := json.Marshal(m1)
		fmt.Printf("data: %s\n", string(bytes))

		err := json.Unmarshal([]byte(`{"List":{"user":{"@order_by":"mtime desc","@select":"id, name, age","@where":"age = 20"},"archive{}获取单个数据":{"id=":"555s5","user_id@":"user/id"}}}`), &d1)
		if err != nil {
			fmt.Printf("error: %+v", err)
		}
		fmt.Printf("map: %+v", d1)

	})
}
