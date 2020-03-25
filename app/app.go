package app

<<<<<<< HEAD
func StartApp() {

=======
import (
	"net/http"

	"github.com/jadahbakar/golang-microservices/controller"
)

// StartApp function
func StartApp() {
	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
>>>>>>> 81d3414fcf7a59290281d2bd1fccb62f72dd4f29
}
