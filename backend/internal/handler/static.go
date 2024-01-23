package handler

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/ops"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterStaticHandler(server *rest.Server, serverCtx *svc.ServiceContext) {
	pathPrefix := "/terminal/"
	staticFileHandler := http.FileServer(
		&assetfs.AssetFS{Asset: ops.Asset, AssetDir: ops.AssetDir, Prefix: "static"},
	)
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    pathPrefix + "js/:name",
		Handler: http.StripPrefix(pathPrefix, staticFileHandler).ServeHTTP,
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    pathPrefix + "css/:name",
		Handler: http.StripPrefix(pathPrefix, staticFileHandler).ServeHTTP,
	})
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   pathPrefix + "config.js",
		Handler: func(w http.ResponseWriter, request *http.Request) {
			w.Header().Set("Content-Type", "application/javascript")
			_, _ = w.Write([]byte("var gotty_term = 'xterm';"))
		},
	})
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   pathPrefix + "auth_token.js",
		Handler: func(w http.ResponseWriter, request *http.Request) {
			w.Header().Set("Content-Type", "application/javascript")
			_, _ = w.Write([]byte("var gotty_auth_token = '';"))
		},
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    pathPrefix + "favicon.png",
		Handler: http.StripPrefix(pathPrefix, staticFileHandler).ServeHTTP,
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/js/:name",
		Handler: ops.DirHandler(serverCtx),
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/css/:name",
		Handler: ops.DirHandler(serverCtx),
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/fonts/:name",
		Handler: ops.DirHandler(serverCtx),
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/img/:name",
		Handler: ops.DirHandler(serverCtx),
	})
}
