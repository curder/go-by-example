package structures

import (
	"encoding/json"
	"testing"
)

type JsonPerson struct {
	Name string `json:"name"` // 反射键
	Age int `json:"age"`
}

func TestStructureAndJson(t *testing.T)  {
	p1:= JsonPerson{
		Name: "curder",
		Age: 28,
	}

	p2, error := json.Marshal(p1) // 序列化结构体
 	if error != nil {
		t.Logf("mashal error %v", error)
		return
	}

	t.Logf("\n序列化数据：%v\n", string(p2))

	// 反序列化
	var p3 JsonPerson
	v1 := `{"name": "luo", "age": 28}`
	error = json.Unmarshal([]byte(v1), &p3)
	if error != nil {
		t.Logf("unmashl error %v", error)
		return
	}
	t.Logf("\n反序列化数据：%#v\n", p3)

}
