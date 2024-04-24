package main
import (
	"net/http"
	"go-http-server/routes"
)

func main() {
	http.HandleFunc("/promotions/{id}", routes.Promotions)
	http.ListenAndServe(":8080", nil);
}
