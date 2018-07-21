package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_Input_Id_1_GetProductHandler_Add_Should_Be_Json_Product(t *testing.T) {

	expected := Product{
		Id:    1,
		Name:  "ลำโพง",
		Price: 100.00,
	}

	request := httptest.NewRequest("GET", "http://localhost:3000/add?id=1", nil)
	responseRecorder := httptest.NewRecorder()

	GetProductHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	var respJson Product
	json.Unmarshal(body, &respJson)

	if expected != respJson {
		t.Errorf("expected '%v' but got '%v'", expected, respJson)
	}
}
