package test

import (
	"fmt"
	"goGINLearn/model/result"
	"net/http"
	"reflect"
)

func Index() string {
	res := "hello gin"
	return res
}

func Post(params map[string]string) *result.Result {

	if len(params) == 0 {
		return result.ReturnNoData(http.StatusOK, "no data")
	}
	return result.ReturnData(http.StatusOK, len(params), "success", params)
}

func Get(params map[string]string) *result.Result {
	if params["id"] == "0" {
		return result.ReturnNoData(http.StatusOK, "no data")
	}
	return result.ReturnData(http.StatusOK, 1, "success", params)
}

func QueryPost(params map[string]string) *result.Result {
	if params["id"] == "0" {
		return result.ReturnNoData(http.StatusOK, "no data")
	}
	return result.ReturnData(http.StatusOK, 1, "success", params)
}

func Map(name, age string, ids map[string]string) *result.Result {
	params := make(map[string]interface{})
	params["name"] = name
	params["age"] = age
	params["ids"] = ids
	fmt.Println(reflect.TypeOf(ids))
	return result.ReturnData(http.StatusOK, 1, "success", params)
}
