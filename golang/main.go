package main

import (
	"boostify/src/infra/websocket"
	loader "boostify/src/main"
	"boostify/src/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow/store/sqlstore"
)


func main() {
	fmt.Println("Initializing...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	err = websocket.ConnectWebSocket("localhost:8080")
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer websocket.Conn.Close()
	sessionContainer, err := sqlstore.New("postgres", utils.GetDSN(true), nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("⏳ Starting clients...")
	loader.StartConnection(
		"",
		sessionContainer,
		"CODE",
	)
	fmt.Println("🚀 Clients started!")
	websocket.SendMessage("🚀 Clients started!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Exiting...")
}
