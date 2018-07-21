package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_Input_Id_1_GetProduct_Add_Should_Be_Json_Product(t *testing.T) {

	expected := Product{
		Id:    1,
		Name:  "ลำโพง",
		Price: 100.00,
	}

	req := httptest.NewRequest("GET", "http://localhost:3000/add?id=1", nil)
	w := httptest.NewRecorder()

	GetProduct(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var respJson Product
	json.Unmarshal(body, &respJson)

	if expected != respJson {
		t.Errorf("expected '%v' but got '%v'", expected, respJson)
	}
}
