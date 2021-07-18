package Controllers

import (
	"net/http"
	"fmt"

	"github.com/armr-dev/cypher-api-go/apiHelpers/cors"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	Cors.SetupCorsResponse(&w, r)

	fmt.Fprintf(w, "Welcome to the home page!\n\n")
	fmt.Fprintf(w, "This is the cypher-api in Golang!\n")
	fmt.Fprintf(w, "You can check our two endpoints: /cypher and /decipher\n\n")
}