package api

import (
    "encoding/json"
    "net/http"
)

// Coin Balance Params
type CoinBalanceParams struct{
    Username string

}

type CoinBalanceResponse struct{
    // HTTP Status Code: success
    Code int

    // Account Balance
    Balance int32

}

type  Error struct{
    // HTTP Stasus Code: error
    Code int

    // Error message
    Message string

}

func writeError(writer http.ResponseWriter, message string, code int) {
    res := Error{
        Code: code,
        Message: message,
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(code)

    json.NewEncoder(writer).Encode(res)
}

var (
        RequestErrorHandler = func(writer http.ResponseWriter, err error) {
            writeError(writer, err.Error(), http.StatusBadRequest)
        }


        InternalErrorHandler = func(writer http.ResponseWriter) {
            writeError(writer, "An unexpected error occurred.", http.StatusInternalServerError)
        }

)
