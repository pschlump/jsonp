package jsonp

import (
	"net/http"
	"testing"
)

type TestAAAResponseWriter struct {
	hdr    http.Header
	buf    []byte
	status int
}

func (ho TestAAAResponseWriter) Header() http.Header {
	return http.Header(ho.hdr)
}

func (ho TestAAAResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (ho TestAAAResponseWriter) WriteHeader(code int) {
	ho.status = code
}

func Test301Redirect(t *testing.T) {

}

func Test_Exists(t *testing.T) {
	/*
		if Exists("./no-such-file.do-not-create") {
			t.Errorf("Expected false, file should not exist\n")
		}
		if !Exists("./test_data") {
			t.Errorf("Expected true, check of directory\n")
		}
		if !Exists("./test_data/file-exists") {
			t.Errorf("Expected true, check of directory\n")
		}
	*/
	SetJsonPrefix(")]}',\n")
	if JSON_Prefix != ")]}',\n" {
		t.Errorf("Expected prefix to be set, not set correctly\n")
	}

	var bob = "{\"bob\":\"bob\"}\n"

	if PrependPrefix(bob) != ")]}',\n{\"bob\":\"bob\"}\n" {
		t.Errorf("Expected prefix to be used, not used correctly\n")
	}

	SetJsonPrefix("")

	res := TestAAAResponseWriter{
		hdr:    make(http.Header),
		buf:    make([]byte, 0, 200),
		status: 0,
	}

	s := JsonP_Param(bob, res, "callback000")
	// fmt.Printf("s >%s<\n", s)
	if s != "callback000({\"bob\":\"bob\"}\n);" {
		t.Errorf("Expected JSONp callback, got >%s< \n", s)
	}

	res1 := TestAAAResponseWriter{
		hdr:    make(http.Header),
		buf:    make([]byte, 0, 200),
		status: 0,
	}

	req := http.Request{
		Method:     "GET",
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		RemoteAddr: "[::1]:64770",
		RequestURI: "/api/status?callback=callback12323123123",
		Host:       "localhost:8204",
		Header:     make(http.Header),
	}

	u := JsonP(`{"status":"success"}`, res1, &req)
	// Check return value
	// fmt.Printf("s= ->%s<-\n", u)
	if u != `callback12323123123({"status":"success"});` {
		t.Errorf("Expected JSONp callback, got >%s< \n", u)
	}
	// check that headers are set corectly
	// fmt.Printf("hdr= ->%s<-\n", res1.hdr)
	if _, ok := res1.hdr["Content-Type"]; !ok {
		t.Errorf("Expected to have a Content-Type header set , but did not find it.")
	} else {
		if v := res1.hdr["Content-Type"]; len(v) > 0 && v[0] != "application/javascript" {
			t.Errorf("Expected to have a Content-Type header set to appliation/javascript but got %s instead\n", v)
		}
	}

}
