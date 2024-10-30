package middleware

import(
    "errors"
    "net/http"

    "github.com/jjetta/go-api-tutorial/api"
    "github.com/jjetta/go-api-tutorial/internal/tools"
    log "github.com/sirupsen/logrus"
)


var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

        var username string = req.URL.Query().Get("username")
        var token = req.Header.Get("Authorization")
        var err error 

        if username == "" || token == "" {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(writer, UnAuthorizedError)
            return
        }
    
        var db *tools.DatabaseInterface
        db, err = tools.NewDatabase()
        
        if err != nil {
            api.InternalErrorHandler(writer)
            return  
        }

        var loginDetails *tools.LoginDetails
        loginDetails = (*db).GetUserLoginDetails(username)

        if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(writer, UnAuthorizedError)
            return  
        }
        
        next.ServeHTTP(writer, req)
    })
}
