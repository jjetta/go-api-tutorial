package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jjetta/go-api-tutorial/api"
	"github.com/jjetta/go-api-tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(writer http.ResponseWriter, req *http.Request){
    var params = api.CoinBalanceParams{}
    var decoder *schema.Decoder = schema.NewDecoder()
    var err error

    err = decoder.Decode(&params, req.URL.Query())

    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(writer)
        return
    }

    var db *tools.DatabaseInterface
    db, err = tools.NewDatabase()
    if err != nil {
        api.InternalErrorHandler(writer)
        return 
    }

    var tokenDetails *tools.CoinDetails
    tokenDetails = (*db).GetUserCoins(params.Username)
    if tokenDetails == nil {
        log.Error(err)
        api.InternalErrorHandler(writer)
        return 
    }

    var res = api.CoinBalanceResponse{
        Balance: (*tokenDetails).Coins,
        Code:    http.StatusOK,
    }

    writer.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(writer).Encode(res)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(writer)
        return 
    }
}   
