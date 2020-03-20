package main

import (
	"bytes"
	"encoding/json"
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
	// Setup
	e := echo.New()
	db := core.NewMockDB()
	ctrls := InitializeMockControllers(db)
	handler := newHandler(e, ctrls)
	s := httptest.NewServer(handler)
	defer s.Close()

	const BaseRoot = "api"

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
			name: "/ping [get]",
			args: args{
				method:    http.MethodGet,
				pathParam: BaseRoot + "/ping",
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"OK"}`},
		},
		{
			name: "/register [post] Success",
			args: args{
				method:    http.MethodPost,
				pathParam: BaseRoot + "/register",
				reqBody:   bytes.NewBufferString(`{"name":"user","email":"dummy@email.com","password":"test1234","password_confirmation":"test1234"}`),
			},
			want: want{statusCode: http.StatusCreated,
				respBody: `{"message":"User successfully created!","user":{"id":1,"name":"user","email":"dummy@email.com"}}`},
		},
		{
			name: "/register [post] Fail",
			args: args{
				method:    http.MethodPost,
				pathParam: BaseRoot + "/register",
				reqBody:   bytes.NewBufferString(`{"name":"user","email":"dummy@email.com","password":"test1234","password_confirmation":"testtest"}`),
			},
			want: want{statusCode: http.StatusBadRequest,
				respBody: `{"message":"password does not match"}`},
		},
		{
			name: "/login [post] Success",
			args: args{
				method:    http.MethodPost,
				pathParam: BaseRoot + "/login",
				reqBody:   bytes.NewBufferString(`{"email":"dummy@email.com","password":"test1234"}`),
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"Successfully logged in!", "user":{"id":1,"name":"user","email":"dummy@email.com"}}`},
		},
		{
			name: "/login [post] Fail: Not registered email",
			args: args{
				method:    http.MethodPost,
				pathParam: BaseRoot + "/login",
				reqBody:   bytes.NewBufferString(`{"email":"invalid@email.com","password":"test1234"}`),
			},
			want: want{statusCode: http.StatusNotFound,
				respBody: `{"message":"User does not exist"}`},
		},
		{
			name: "/login [post] Fail: Invalid password",
			args: args{
				method:    http.MethodPost,
				pathParam: BaseRoot + "/login",
				reqBody:   bytes.NewBufferString(`{"email":"dummy@email.com","password":"invalid_password"}`),
			},
			want: want{statusCode: http.StatusUnauthorized,
				respBody: `{"message":"Password is invalid"}`},
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
			//got := string(body)
			var got controller.Response
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("json.Unmarshal failed: %s", err)
			}
			var want controller.Response
			if err := json.Unmarshal([]byte(tt.want.respBody), &want); err != nil {
				t.Fatalf("json.Unmarshal failed: %s", err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("request = /%v, got %v, want %v\n", tt.args.pathParam, got, want)
			}
		})
	}
}
