package gameasure

// import (
// 	"net/http"
// )

// func (mw *Stats) Handler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		beginning, recorder := mw.Begin(w)

// 		h.ServeHTTP(recorder, r)

// 		mw.End(beginning, recorder)
// 	})
// }
