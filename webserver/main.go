package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	http.ListenAndServe(":8888", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml, _ := os.Open("index.html")
	defer indexHtml.Close()
	indexData, _ := io.ReadAll(indexHtml)

	fmt.Fprint(w, string(indexData))
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	mType, _, _ := conn.ReadMessage()
	fmt.Printf("type: %d\n", mType)
	conn.WriteMessage(mType, []byte("message from backend"))
}
