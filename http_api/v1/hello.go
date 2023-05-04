package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ticket-reservation/app"
	customError "ticket-reservation/custom_error"
	"ticket-reservation/http_api/response"
	"ticket-reservation/http_api/routes"
)

var HelloRoutes = routes.Routes{
	routes.Route{
		Name:        "Hello",
		Path:        "/hello",
		Method:      "POST",
		HandlerFunc: Hello,
	},
}

func init() {
	RouteDefinitions = append(RouteDefinitions, routes.RouteDefinition{
		Routes: HelloRoutes,
		Prefix: "",
	})
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b *Book) ToString() string {
	return b.Title + " " + b.Author
}

// FIXME
func Hello(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	logger := ctx.Logger

	books := Book{Title: "Title Name", Author: "Jezy"}
	logger.Infof(books.ToString())

	var input app.RegisterParams

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return &customError.UserError{
			Code:           customError.InvalidJSONString,
			Message:        "Invalid JSON string",
			HTTPStatusCode: http.StatusBadRequest,
		}
	}

	logger.Infof("%+v", string(body))

	resData, err := ctx.Register(input)
	if err != nil {
		return err
	}

	logger.Infof("%+v", resData)
	logger.Infof("INPUT:")
	logger.Infof("%+v", input)

	data, err := json.Marshal(&response.Response{
		Code:    0,
		Message: "",
		Data:    books,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
