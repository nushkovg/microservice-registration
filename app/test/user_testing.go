// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user": user TestHelpers
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-registration/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-registration
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Microkubes/microservice-registration/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// RegisterUserBadRequest runs the method Register of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegisterUserBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.UserPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	registerCtx, __err := app.NewRegisterUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	registerCtx.Payload = payload

	// Perform action
	__err = ctrl.Register(registerCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// RegisterUserCreated runs the method Register of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegisterUserCreated(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.UserPayload) (http.ResponseWriter, *app.Users) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	registerCtx, __err := app.NewRegisterUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	registerCtx.Payload = payload

	// Perform action
	__err = ctrl.Register(registerCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.Users
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.Users)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Users", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// RegisterUserInternalServerError runs the method Register of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegisterUserInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.UserPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	registerCtx, __err := app.NewRegisterUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	registerCtx.Payload = payload

	// Perform action
	__err = ctrl.Register(registerCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ResendVerificationUserBadRequest runs the method ResendVerification of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ResendVerificationUserBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.ResendVerificationPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register/resend-verification"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	resendVerificationCtx, __err := app.NewResendVerificationUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	resendVerificationCtx.Payload = payload

	// Perform action
	__err = ctrl.ResendVerification(resendVerificationCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ResendVerificationUserInternalServerError runs the method ResendVerification of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ResendVerificationUserInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.ResendVerificationPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register/resend-verification"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	resendVerificationCtx, __err := app.NewResendVerificationUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	resendVerificationCtx.Payload = payload

	// Perform action
	__err = ctrl.ResendVerification(resendVerificationCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ResendVerificationUserOK runs the method ResendVerification of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ResendVerificationUserOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, payload *app.ResendVerificationPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/users/register/resend-verification"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	resendVerificationCtx, __err := app.NewResendVerificationUserContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil
	}
	resendVerificationCtx.Payload = payload

	// Perform action
	__err = ctrl.ResendVerification(resendVerificationCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
