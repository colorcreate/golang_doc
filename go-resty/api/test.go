package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty"
	"github.com/labstack/echo"
)

//Test is
type Test struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

//GetAll is
func GetAll(c echo.Context) error {
	resp, err := resty.R().
		SetResult([]Test{}).
		Get("http://5aaa3b084cf36300144e96d6.mockapi.io/test")
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, resp.Result().(*[]Test))
}

//GetByID is
func GetByID(c echo.Context) error {
	id := c.Param("id")
	resp, err := resty.R().
		SetResult(Test{}).
		Get("http://5aaa3b084cf36300144e96d6.mockapi.io/test/" + id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, resp.Result().(*Test))
}

//Create is
func Create(c echo.Context) error {
	req := c.Request()
	body, _ := ioutil.ReadAll(req.Body)
	var param Test
	err := json.Unmarshal(body, &param)
	if err != nil {
		panic(err)
	}

	resp, err := resty.R().
		SetBody(param).
		SetResult(Test{}).
		Post("http://5aaa3b084cf36300144e96d6.mockapi.io/test")
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, resp.Result().(*Test))
}

//Remove is
func Remove(c echo.Context) error {
	id := c.Param("id")
	_, err := resty.R().
		Delete("http://5aaa3b084cf36300144e96d6.mockapi.io/test/" + id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, id)
}

//Update is
func Update(c echo.Context) error {
	id := c.Param("id")
	req := c.Request()
	body, _ := ioutil.ReadAll(req.Body)

	var param Test
	err := json.Unmarshal(body, &param)
	if err != nil {
		panic(err)
	}

	resp, err := resty.R().
		SetBody(param).
		SetResult(Test{}).
		Put("http://5aaa3b084cf36300144e96d6.mockapi.io/test/" + id)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, resp.Result().(*Test))
}

//Hello is/G
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
