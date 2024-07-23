package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

)


//reusable function to parse the json into required structure
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing Request Body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

//this function is basically for formatting the error depending on the status
func WriteJSON (w http.ResponseWriter,status int, v any) error{
	w.Header().Add("content-type","application/json");
	w.WriteHeader(status);

	return json.NewEncoder(w).Encode(v);

}

//this function is to convert the error into key value pairs and writing the error.
func WriteError(w http.ResponseWriter,status int, err error) {
	WriteJSON(w,status,map[string]string{"error": err.Error()})
}
