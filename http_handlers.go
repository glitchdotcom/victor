package victor

import (
	"fmt"
	"github.com/brettbuddin/victor/pkg/httpserver"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func handlers(bot *Robot) {
	router := mux.NewRouter()

	router.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		text, _ := ioutil.ReadAll(r.Body)
		bot.Receive(&textMessage{string(text)})
	}).Methods("POST")

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, httpserver.Message("ok", bot.Store().All()))
	}).Methods("GET")

	router.HandleFunc("/data/{key}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		if val, ok := bot.Store().Get(vars["key"]); ok {
			fmt.Fprintf(w, httpserver.Message("ok", val))
		} else {
			fmt.Fprintf(w, httpserver.Message("error", "key doesn't exist"))
		}
	}).Methods("GET")

	router.HandleFunc("/data/{key}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		body, _ := ioutil.ReadAll(r.Body)
		bot.Store().Set(vars["key"], string(body))
		fmt.Fprintf(w, httpserver.Message("ok", "key set"))
	}).Methods("POST", "PUT")

	router.HandleFunc("/data/{key}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		bot.Store().Delete(vars["key"])
		fmt.Fprintf(w, httpserver.Message("ok", "key deleted"))
	}).Methods("DELETE")

	bot.HTTP().Handle("/", router)
}

type textMessage struct {
	text string
}

func (m *textMessage) UserId() string {
	return ""
}

func (m *textMessage) UserName() string {
	return ""
}

func (m *textMessage) ChannelId() string {
	return ""
}

func (m *textMessage) ChannelName() string {
	return ""
}

func (m *textMessage) Text() string {
	return m.text
}
