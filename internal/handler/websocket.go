package handler

import (
	"github.com/gorilla/websocket"
	"math"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request, consumer func(conn *websocket.Conn)) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	consumer(conn)

	err = conn.Close()
	if err != nil {
		return
	}
}

func SendToWebSocket(conn *websocket.Conn) {
	i := 0
	for {
		data := make([]byte, 2)
		data[0] = (byte)((math.Sin(float64(i)/100.0) + 1) / 2.0 * 256)
		data[1] = (byte)((math.Cos(float64(i)/100.0) + 1) / 2.0 * 256)
		i = i + 1
		err := conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			return
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func WebSocketHandlerWrapper(w http.ResponseWriter, r *http.Request) {
	WebSocketHandler(w, r, SendToWebSocket)
}
