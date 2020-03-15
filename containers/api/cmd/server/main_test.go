package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/tetsuzawa/Goonstone/containers/api/cmd/server/controller"
	"github.com/tetsuzawa/Goonstone/containers/api/internal/core"
)

func InitializeMockControllers(db *core.MockDB) *controller.Controllers {
	repository := core.NewMockGateway(db)
	provider := core.NewProvider(repository)
	controllerController := controller.NewController(provider)
	controllers := controller.NewControllers(controllerController)
	return controllers
}

func Test_newHandler(t *testing.T) {
	e := echo.New()

	// Setup
	db := core.NewMockDB()
	ctrls := InitializeMockControllers(db)
	handler := newHandler(e, ctrls)
	s := httptest.NewServer(handler)
	defer s.Close()

	type args struct {
		method    string
		pathParam string
		reqBody   *bytes.Buffer
	}
	type want struct {
		statusCode int
		respBody   string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "/ping/ [get]",
			args: args{
				method:    http.MethodGet,
				pathParam: "api/ping/",
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"OK"}
`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Request and Assertions
			client := new(http.Client)
			var req *http.Request
			var err error
			switch tt.args.method {
			case http.MethodGet:
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, nil)
			case http.MethodPost:
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, tt.args.reqBody)
			case http.MethodPatch:
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, tt.args.reqBody)
			case http.MethodDelete:
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, nil)
			default:
				t.Fatalf("method not allowed")
			}
			if err != nil {
				t.Fatalf("http.NewRequest: %v", err)
			}
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("client.Do: %v", err)
			}

			if !reflect.DeepEqual(resp.StatusCode, tt.want.statusCode) {
				t.Fatalf("resp.StatusCode: got: %d, want: %d", resp.StatusCode, tt.want.statusCode)
			}
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("ioutil.ReadAll failed: %s", err)
			}
			got := string(body)

			if !reflect.DeepEqual(got, tt.want.respBody) {
				t.Errorf("request = /%v, got %v, want %v", tt.args.pathParam, got, tt.want.respBody)
			}
		})
	}
}
