<<<<<<< HEAD
// Package testgrp contains all the dbtest handlers.
=======
// Package testgrp contains all the test handlers.
>>>>>>> 7-Middleware
package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

<<<<<<< HEAD
	webv1 "github.com/dmitryovchinnikov/service/business/web/v1"
=======
	v1web "github.com/dmitryovchinnikov/service/business/web/v1"
>>>>>>> 7-Middleware
	"github.com/dmitryovchinnikov/service/foundation/web"
	"go.uber.org/zap"
)

<<<<<<< HEAD
// Handlers manages the set of check enpoints.
=======
// Handlers manages the set of check endpoints.
>>>>>>> 7-Middleware
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development.
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
<<<<<<< HEAD
		return webv1.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
=======
		//return errors.New("untrusted error")
		return v1web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		//return web.NewShutdownError("restart service")
		//panic("testing panic")
>>>>>>> 7-Middleware
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

<<<<<<< HEAD
=======
	//statusCode := http.StatusOK
	//h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

>>>>>>> 7-Middleware
	return web.Respond(ctx, w, status, http.StatusOK)
}
