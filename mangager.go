package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
    "time"
	"github.com/gorilla/websocket"
)

/*
 the code for creating new room or manager is over
 the code for clients is remaining
 adding new client to manager and also sending request from main thread to manager thread via channel is also remaining
 code for starting new thread or go routine for new manager is to be written in new_Manager
 code for client reading
 code for cleint writing  are also remaing
 ping pong code is also remaining

*/

var (
	/**
	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
	*/
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	mmid            int
	mp              map[int]*Client
	manager_channel chan *websocket.Conn
	mu              sync.RWMutex
}
type rtformat struct {
	rq *http.Request
	wq http.ResponseWriter
}

func new_Manager(manager_id int) *Manager {

	new_manager := new(Manager)
	new_manager.mmid = manager_id
	new_manager.manager_channel = make(chan *websocket.Conn)
	new_manager.mp = make(map[int]*Client)
	return new_manager
}

func (mt *Manager) Add_client(client *Client, conn *websocket.Conn) {

	mt.mp[client.client_id] = client
	client.Mg = mt
	
	client.connection = conn
	fmt.Println("client is added in the room", client.client_id)

	//fmt.Println("write thread started")

    go client.writeMessages()
	time.Sleep(3* time.Second) 
	go client.readMessages()

	

	//fmt.Println("read and write threads are started")
	//  w:=http.ResponseWriter
	//  respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok","message":"new client is in the  room"})

}
func (mt *Manager) Receive_new_client() {
	client_count := 0
	for {
		
		new_client_http_request := <-mt.manager_channel

		new_cl := new_Client(client_count)
		fmt.Println("new client receiving")
		mt.Add_client(new_cl, new_client_http_request)
		client_count = client_count + 1

	}
}

func (All_Manager *all_Managers) handlerReadiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AJFAJ")
	// respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok","message":"new room added"})
	newman := new_Manager(4)
	go newman.Receive_new_client()
	(All_Manager.manager_map)[newman.mmid] = newman
	respondWithJSON(w, http.StatusOK, map[string]string{"status": fmt.Sprint(newman.mmid), "message": "new room added"})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
