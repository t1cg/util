package weblog

import (
	"net/http"
	"reflect"
	"runtime"
	//project library
	"github.com/t1cg/util/logger"
)

// HTTPRequestInfo struct defines the field values and the method to log the http request.
type HTTPRequestInfo struct {
	Handler       string
	Method        string
	ContentLength int64
	Host          string
	Origin        string
	Referer       string
	UserAgent     string
	RemoteAddr    string
	RequestURI    string
}

// Log function ...
func (h *HTTPRequestInfo) Log() {
	logger.L.Info.Printf(
		"handler: %s, httpMethod: %s, h.ContentLength: %v, host: %s, origin: %s, referer: %s, userAgent: %s, remoteAddr: %s, requestUri: %s",
		h.Handler, h.Method, h.ContentLength, h.Host, h.Origin, h.Referer, h.UserAgent, h.RemoteAddr, h.RequestURI)
}

// LogMe function ...
func LogMe(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//get the caller name
		funcName := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()

		//log the http request details
		logInfo := &HTTPRequestInfo{
			Handler:       funcName,
			Method:        r.Method,
			ContentLength: r.ContentLength,
			Host:          r.Host,
			Origin:        r.Header.Get("Origin"),
			Referer:       r.Header.Get("Referer"),
			UserAgent:     r.Header.Get("User-Agent"),
			RemoteAddr:    r.RemoteAddr,
			RequestURI:    r.RequestURI,
		}

		logInfo.Log()
		h(w, r)
	}
}
