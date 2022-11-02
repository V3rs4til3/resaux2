package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", gererConnection)
	http.HandleFunc("/calcul", loadResult)
	http.HandleFunc("/ws", ws)

	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe(":8080", nil)
}

func gererConnection(w http.ResponseWriter, r *http.Request) {
	vueRaw, _ := os.ReadFile("./views/index.html")
	vue := string(vueRaw)
	vue = strings.Replace(vue, "###Titre###", "Home", 1)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, vue)

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		content, _ := os.ReadFile("./html/home.html")
		io.WriteString(w, string(content))
	case "POST":
		var num = r.FormValue("num")
		
		var resultCookie = http.Cookie{Name: "result", Value: strconv.Itoa(result)}
		http.SetCookie(w, &resultCookie)
		http.Redirect(w, r, "/result", http.StatusTemporaryRedirect)
	}

}

func loadResult(w http.ResponseWriter, _ *http.Request) {
	vueRaw, _ := os.ReadFile("./views/calcul.html")
	vue := string(vueRaw)
	io.WriteString(w, vue)
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer c.Close()

	for {
		mt, message, _ := c.ReadMessage()
		log.Printf("Rec: %s", message)
		c.WriteMessage(mt, message)
	}
}
