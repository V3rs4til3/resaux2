package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", gererConnection)
	http.HandleFunc("/ws", ws)

	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe(":8080", nil)
}

func gererConnection(w http.ResponseWriter, r *http.Request) {
	vueRaw, _ := os.ReadFile("./views/index.html")
	vue := string(vueRaw)
	vue = strings.Replace(vue, "###Titre###", "Blearg", 1)
	w.Header().Set("Content-Type", "text/html")
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
