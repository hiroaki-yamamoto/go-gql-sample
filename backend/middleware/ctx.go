package middleware

import (
  "context"
  "net/http"
)

type ctxKey struct {
  name string
}

// RequestKey represents the key of the request in context.
var RequestKey = &ctxKey{"request"}
// ResponseKey represents the key of the response in context.
var ResponseKey = &ctxKey{"response"}

// ContextReqRespMiddleware puts response writer and request into context
func ContextReqRespMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
    ctx := context.WithValue(req.Context(), RequestKey, req)
    ctx = context.WithValue(ctx, ResponseKey, &resp)
    next.ServeHTTP(resp, req.WithContext(ctx))
  })
}
