package main

import (
	"net/http"

	"gihub.com/longth97/gomod/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
