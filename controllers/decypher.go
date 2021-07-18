package Controllers

import (
	"io/ioutil"
	"encoding/json"
	"net/http"

	"github.com/armr-dev/cypher-api-go/apiHelpers/cors"
	"github.com/armr-dev/cypher-api-go/resources"
	"github.com/armr-dev/cypher-api-go/services/aes"
	"github.com/armr-dev/cypher-api-go/services/des"
	"github.com/armr-dev/cypher-api-go/services/3des"
	"github.com/armr-dev/cypher-api-go/services/blowfish"
	"github.com/armr-dev/cypher-api-go/services/idea"
)

func DecipherText(w http.ResponseWriter, r *http.Request) {
	Cors.SetupCorsResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Resources.Request

	json.Unmarshal(reqBody, &request)

	var decryptedText string

	switch request.Algorithm {
	case "aes":
		decryptedText, _ = AES.Decrypt(request.Text)

	case "des":
		decryptedText, _ = DES.Decrypt(request.Text)

	case "3des":
		decryptedText, _ = TripleDES.Decrypt(request.Text)

	case "blowfish":
		decryptedText, _ = Blowfish.Decrypt(request.Text)

	case "idea":
		decryptedText, _ = Idea.Decrypt(request.Text)

	default:
		decryptedText = "Error"
	}

	var response Resources.Response

	response.Data.Text = decryptedText

	json.NewEncoder(w).Encode(response)
}