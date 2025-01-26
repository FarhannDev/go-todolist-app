package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist-app/models"
	"todolist-app/router"
)

func main() {
	models.Todos = append(models.Todos, models.Todo{ID: "1", Task: "Learn Golang", Checked: false})

	r := router.SetupRouter()
	fmt.Println("Server berjalan pada port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
