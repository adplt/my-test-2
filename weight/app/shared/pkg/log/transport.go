package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextKeyRequestStart godoc.
var ContextKeyRequestStart = &contextKey{"RequestStart"}

type transport struct {
	ctx *gin.Context
	req *http.Request
	res *http.Response
}

type contextKey struct {
	name string
}

// RoundTrip godoc.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var err error
	trace := &httptrace.ClientTrace{GotConn: t.GotConn}

	ctx := context.WithValue(req.Context(), ContextKeyRequestStart, time.Now())
	req = req.WithContext(httptrace.WithClientTrace(ctx, trace))
	t.req = req

	t.ReqInfo()

	t.res, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return t.res, err
	}

	t.ResInfo()

	return t.res, err
}

func (t *transport) GotConn(gotConnInfo httptrace.GotConnInfo) {
	// Info(t.ctx, fmt.Sprintf("%+v", gotConnInfo))
}

// ReqInfo godoc.
func (t *transport) ReqInfo() {
	msg := "Request  "
	if t.req != nil {

		msg += fmt.Sprintf("%v %v ", t.req.Method, t.req.URL.String())

		bodyBytes := []byte("")
		if t.req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(t.req.Body)
		}
		t.req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var body interface{}
		_ = json.Unmarshal(bodyBytes, &body)
		msg += fmt.Sprintf("%v", body)
	}

	Info(t.ctx, msg)
}

// ResInfo godoc.
func (t *transport) ResInfo() {
	msg := "Response "
	if t.res == nil {
		msg += fmt.Sprintf("%v no response", t.req.URL)
	} else {
		msg += fmt.Sprintf("%v %v %v ", t.req.Method, t.res.Request.URL.String(), t.res.StatusCode)
		ctx := t.res.Request.Context()
		if start, ok := ctx.Value(ContextKeyRequestStart).(time.Time); ok {
			msg += fmt.Sprintf("%v", time.Since(start))
		}
		if t.res.Body != nil {
			bodyBytes := []byte("")
			if t.res.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(t.res.Body)
			}
			t.res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			var body interface{}
			_ = json.Unmarshal(bodyBytes, &body)
			msg += fmt.Sprintf(" %v", body)
		}
	}
	Info(t.ctx, msg)
}

// Transport godoc.
func Transport(ctx *gin.Context) *transport {
	return &transport{ctx: ctx}
}
