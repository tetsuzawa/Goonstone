package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_newHandler(t *testing.T) {
	// Setup
	e := newEcho()
	db := newDB()
	handler := newHandler(e, db)
	s := httptest.NewServer(handler)
	defer s.Close()

	type args struct {
		param string
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
			name: "/",
			args: args{param: ""},
			want: want{statusCode: http.StatusNotFound, respBody: "{\"message\":\"Not Found\"}\n"},
		},
		{
			name: "apple",
			args: args{param: "apple"},
			want: want{statusCode: http.StatusOK, respBody: "apple"},
		},
		{
			name: "orange123",
			args: args{param: "orange123"},
			want: want{statusCode: http.StatusOK, respBody: "orange123"},
		},
		{
			name: "Grape",
			args: args{param: "Grape"},
			want: want{statusCode: http.StatusOK, respBody: "Grape"},
		},
		{
			name: "BANANA",
			args: args{param: "BANANA"},
			want: want{statusCode: http.StatusOK, respBody: "BANANA"},
		},
		{
			name: "3",
			args: args{param: "3"},
			want: want{statusCode: http.StatusOK, respBody: "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Request and Assertions
			res, err := http.Get(s.URL + "/" + tt.args.param)
			if err != nil {
				t.Fatalf("http.Get failed: %s", err)
			}
			if !reflect.DeepEqual(res.StatusCode, tt.want.statusCode) {
				t.Fatalf("res.StatusCode: got: %d, want: %d", res.StatusCode, tt.want.statusCode)
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatalf("ioutil.ReadAll failed: %s", err)
			}
			got := string(body)

			if !reflect.DeepEqual(got, tt.want.respBody) {
				t.Errorf("request = /%v, got %v, want %v", tt.args.param, got, tt.want.respBody)
			}
		})
	}
}
