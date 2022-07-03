package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/dmitryovchinnikov/service/foundation/web"
	"go.uber.org/zap"
)

// Logger writes some information about the request to the logs in the
// format: TraceID : (200) GET /foo -> IP ADDR (latency)
func Logger(log *zap.SugaredLogger) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, err := web.GetValues(ctx)
			if err != nil {
<<<<<<< HEAD
				return web.NewShutdownError("web value missing from context")
=======
				return err
				//return web.NewShutdownError("web value missing from context")
>>>>>>> 7-Middleware
			}

			log.Infow("request started", "traceid", v.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			// Call the next handler.
			err = handler(ctx, w, r)

			log.Infow("request completed", "traceid", v.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))

<<<<<<< HEAD
			// Return the error so it can be handled further up the chain.
=======
			// Return the error, so it can be handled further up the chain.
>>>>>>> 7-Middleware
			return err
		}

		return h
	}

	return m
}
