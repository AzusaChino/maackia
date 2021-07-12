package httputil

import (
	"context"
	"net"
	"net/http"
)

type pathParam struct{}

type QueryOrigin struct{}

func ContextWithPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, pathParam{}, path)
}

func ContextFromRequest(ctx context.Context, r *http.Request) context.Context {
	var ip string
	if r.RemoteAddr != "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	var path string
	if v := ctx.Value(pathParam{}); v != nil {
		path = v.(string)
	}
	return context.WithValue(ctx, QueryOrigin{}, map[string]interface{}{
		"httpRequest": map[string]string{
			"clientIP": ip,
			"method":   r.Method,
			"path":     path,
		},
	})
}
