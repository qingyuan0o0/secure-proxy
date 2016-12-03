package sample

import (
	"io"
	"net/http"
	"regexp"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
}

func copyHeaders(dst, src http.Header) {
	for k, _ := range dst {
		dst.Del(k)
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// validate
	if !regexp.MustCompile(`^\/[a-zA-Z0-9\.]+\.apple\.com\/`).MatchString(r.URL.Path) {
		http.Error(w, r.URL.Path + " invalid url.", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("GET", "http:/" + r.URL.Path, nil)
	reqRange := r.Header.Get("Range")
	if reqRange != "" {
		req.Header.Add("Range", reqRange)
	}

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
}
