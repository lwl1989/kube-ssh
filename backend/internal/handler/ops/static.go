package ops

import (
	"github.com/go-libraries/kube-manager/backend/internal/svc"
	"net/http"
	"os"
)

func DirHandler(serverCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filedir := serverCtx.Config.StaticDir + r.URL.Path
		if r.URL.Path == "/" {
			filedir = serverCtx.Config.StaticDir + "/index.html"
		}

		_, err := os.Stat(filedir)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Path == "/index.html" {
			r.URL.Path = "/"
		}
		handler := http.StripPrefix("", http.FileServer(http.Dir(serverCtx.Config.StaticDir)))
		handler.ServeHTTP(w, r)
	}
}
