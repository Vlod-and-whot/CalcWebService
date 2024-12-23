package Server

import (
	"CalculationWebService/Packages/Calculation"
	"encoding/json"
	"fmt"
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
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid"})
		fmt.Println(err)
		return
	}

	result, err := Calculation.Calc(req.Expression)
	if err != nil {
		switch err.Error() {
		case "недопустимый символ", "пустое выражение", "несоответствующая скобка", "ошибка в выражении", "ошибка - недостаточно данных для вычисления", "ошибка - деление на ноль", "ошибка - неизвестная операция":
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid"})
			fmt.Println(err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Error: "Internal server error"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseSuccess{Result: floatToString(result)})
}

func Server() {
	http.HandleFunc("/api/v1/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ResponseError{Error: "Method not allowed", Code: 405})
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid", Code: 422})
			fmt.Println(err)
			return
		}

		result, err := Calculation.Calc(req.Expression)
		if err != nil {
			switch err.Error() {
			case "invalid expression", "division by zero":
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid", Code: 405})
				fmt.Println(err)
			default:
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ResponseError{Error: "Internal server error", Code: 500})
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseSuccess{Result: floatToString(result), Code: 200})
	})
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
