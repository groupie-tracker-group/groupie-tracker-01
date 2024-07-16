package groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
)



// THIS IS THE API URL

var API = "https://groupietrackers.herokuapp.com/api"

// THIS FUNCTION WILL FETCH THE DATA FROM THE API AND RETURN IT AS A BYTE ARRAY
func GetData(category string) ([]byte, error) {
	// THE GET FUNCTION WILL MAKE A GET REQUEST TO THE API AND RETURN THE RESPONSE
	resp, err := http.Get(API + "/" + category)
	if err != nil {
		return nil, fmt.Errorf("error fetching %s: %v", API+"/"+category, err)
	}

	defer resp.Body.Close()

	var data interface{}
	// THE DECODER FUNCTION WILL DECODE THE JSON OBJECT INTO A GO OBJECT AND STORE IT IN THE DATA VARIABLE WHICH IS AN INTERFACE
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}
	// THE MARSHALLINDENT FUNCTION WILL RETURN A JSON OBJECT
	jsonData, err := json.MarshalIndent(data, "", "       ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	return jsonData, nil
}
