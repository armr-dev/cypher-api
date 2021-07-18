package Controllers

import (
	"io/ioutil"
	"encoding/json"
	"net/http"

	"github.com/armr-dev/cypher-api-go/resources"
	"github.com/armr-dev/cypher-api-go/apiHelpers/cors"
	"github.com/armr-dev/cypher-api-go/services/aes"
	"github.com/armr-dev/cypher-api-go/services/des"
	"github.com/armr-dev/cypher-api-go/services/3des"
	"github.com/armr-dev/cypher-api-go/services/blowfish"
	"github.com/armr-dev/cypher-api-go/services/idea"
)

func CypherText(w http.ResponseWriter, r *http.Request) {
	Cors.SetupCorsResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Resources.Request

	json.Unmarshal(reqBody, &request)

	var encryptedText string

	switch request.Algorithm {
	case "aes":
		encryptedText, _ = AES.Encrypt(request.Text)

	case "des":
		encryptedText, _ = DES.Encrypt(request.Text)

	case "3des":
		encryptedText, _ = TripleDES.Encrypt(request.Text)

	case "blowfish":
		encryptedText = Blowfish.Encrypt(request.Text)

	case "idea":
		encryptedText, _ = Idea.Encrypt(request.Text)

	default:
		encryptedText = "Error"
	}

	var response Resources.Response

	response.Data.Text = encryptedText

	json.NewEncoder(w).Encode(response)
}
