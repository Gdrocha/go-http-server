package main

import (
	"fmt"
	"go-http-server/src/routes"
	"net/http"
)

func main() {
	port := 8081

	routes.RegisterRoutes()

	fmt.Printf("Server is running on http://localhost:%d\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
