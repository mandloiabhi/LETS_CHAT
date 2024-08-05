package main

import(
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	
)
type Client struct{
	client_id int
	client_channel chan []byte
	connection *websocket.Conn
	Mg *Manager
}

func new_Client(client_id int) *Client{
	new_client:= new(Client)
	new_client.client_id=client_id
	// new_manager.mp = make(map[int]*Client)
	new_client.client_channel=make(chan []byte)
	return new_client
}

func (c *Client) readMessages() {
	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}
		log.Println("MessageType: ", messageType)
		log.Println("Payload: ", string(payload))

		
		for _, client := range c.Mg.mp {
			fmt.Println(payload)
			select {
				
			case (client.client_channel) <- payload:
				log.Println("Sent to client with ID:", client.client_id)
			default:
				log.Println("Failed to send message to client with ID:", client.client_id)
			}
		}
		
	}
}


func (c *Client) writeMessages() {
	
	for {log.Println("i am waiting for broadcast msg")

		select {
		case message, ok := <-c.client_channel:
			// Ok will be false Incase the egress channel is closed
			if !ok {
				// Manager has closed this connection channel, so communicate that to frontend
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					// Log that the connection is closed and the reason
					log.Println("connection closed: ", err)
				}
				// Return to close the goroutine
				return
			}
			// Write a Regular text message to the connection
			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
			}
			log.Println("sent message")
		}

	}
}