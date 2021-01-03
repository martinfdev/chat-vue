package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	//sockets
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("connected:", so.ID())
		log.Println("New User Connected!")
		so.Join("room_chat")
		return nil
	})

	server.OnEvent("/", "chat message", func(so socketio.Conn, msg string) {
		log.Println(msg)
		so.Emit("chat message", msg)
		server.BroadcastToRoom("/", "room_chat", "chat message", msg)
	})

	go server.Serve()
	defer server.Close()

	//http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
