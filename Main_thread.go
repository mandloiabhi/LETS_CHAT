package main


import(
	
	"net/http"
	_"encoding/json"
	"fmt"
	"strconv"
	_"github.com/gorilla/websocket"

)

func (All_Manager *all_Managers) addClientIntoRoom(w http.ResponseWriter, r *http.Request){

      // read room id from the request 
	  // then find the manager of that room id from the map
	  // after it send  r htttp request to the channel of that room 
	type parameters struct {
		Room_id string
	}
	// decoder := json.NewDecoder(r.Body)
	// var params  parameters
	// // err := decoder.Decode(&params)
	// // fmt.Println(params)
	// // if err != nil {
	// // 	// respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	// // 	fmt.Println("error in decoding the request", err)
	// // 	return ;
	// // 	// handle here properly

	// // } 
    // // t:=params.Room_id
	t:="4"
	roomid,err := strconv.Atoi(t)
	if err !=nil {
		fmt.Println("error in decoding the request")
		return ;
	}
    
	mk:=(All_Manager.manager_map)[roomid]

	// pt:= new(rtformat)
	// pt.rq=r
	// pt.wq=w

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Upgrade error:", err)
        return
    }
    // defer conn.Close()
	count=count+1
	fmt.Println(count)

	mk.manager_channel<-conn

    fmt.Println("send the connection in channel for client ",count)
}