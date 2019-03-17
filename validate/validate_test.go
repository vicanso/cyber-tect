package validate

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		s := &struct {
			Method string `json:"method" valid:"xMethods"`
			Type   int    `json:"type" valid:"xIntIn(1|3|5)"`
			Size   int    `json:"size" valid:"xIntRange(1|2)"`
		}{}
		err := Do(s, []byte(`{
			"method": "GET",
			"type": 3,
			"size": 10
		}`))
		if err == nil {
			t.Fatalf("validate should be fail")
		}
		err = Do(s, map[string]interface{}{
			"method": "GET",
			"type":   3,
			"size":   1,
		})
		if err != nil {
			t.Fatalf("validate fail, %v", err)
		}
	})
}
