package Handler

import (
	"CalculationWebService/Packages/Calculation"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Request struct {
	Expression string `json:"expression"`
}

type ResponseSuccess struct {
	Result string `json:"result"`
	Code   int    `json:"code"`
}

type ResponseError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseError{Error: "Method not allowed", Code: 405})
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid", Code: 422})
		log.Println("Error decoding request body:", err)
		return
	}

	result, err := Calculation.Calc(req.Expression)
	if err != nil {
		switch err.Error() {
		case "invalid expression", "division by zero":
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid", Code: 422})
			log.Println("Calculation error:", err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Error: "Internal server error", Code: 500})
			log.Println("Internal server error:", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseSuccess{Result: floatToString(result), Code: 200})
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
