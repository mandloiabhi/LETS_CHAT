package main

import (
	"fmt"
	
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
	_"github.com/lib/pq"
)
type all_Managers struct {
	manager_map map[int]*Manager
}
var count=0
func main(){
    
    myMap := make(map[int]*Manager)

	// Create an instance of all_Managers and initialize the map
	Manager_obj := &all_Managers{
		manager_map: myMap,
	}

	

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		//AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
    router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

    v1Router.Get("/newRoom", Manager_obj.handlerReadiness)
	v1Router.HandleFunc("/ws",Manager_obj.addClientIntoRoom)

	fmt.Println("server is listening on 8080")
	//log.Fatal(srv.ListenAndServe())
	//fmt.Println(("servers is staerd"))
    

	var err1 = srv.ListenAndServe()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(("servers is staerd"))


}