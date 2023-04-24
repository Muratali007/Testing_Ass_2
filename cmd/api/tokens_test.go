package main

import (
	"net/http/httptest"
	"testing"
)

func TestCreateAuthenticationTokenHandler(t *testing.T) {

	app := newTestApplication(t)

	r := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()
	app.createAuthenticationTokenHandler(w, r)
}
