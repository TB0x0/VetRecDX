package auth

import (
	"encoding/json"
	"net/http"
	"os"
)

func CreateAuthConfig() (AuthConfig, error) {
	secretKey := os.Getenv("SECRET_KEY")
	issuer := os.Getenv("ISSUER")

	return AuthConfig{
		SecretKey: []byte(secretKey),
		Issuer:    issuer,
	}, nil
}

func ConstructJWT(authConfig AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userInfo User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userInfo); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		encodeResult, err := encodeJWT(userInfo, authConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{"token": encodeResult}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeconstructJWT(authConfig AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputToken := r.PathValue("token")

		decodeResult, err := decodeJWT(inputToken, authConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := map[string]interface{}{
			"claims": decodeResult,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
