// Package testgrp contains all the test handlers.
package testgrp

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development.
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)

	statusCode := http.StatusOK
	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

	//return web.Respond(ctx, w, status, http.StatusOK)
	return nil
}
