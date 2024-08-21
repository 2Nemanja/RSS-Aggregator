package main

import (
	"fmt"
	"net/http"

	"github.com/2Nemanja/RSSAggregator/internal/auth"
	"github.com/2Nemanja/RSSAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) //creating a custom type with regular handler fun params but with database.User included so we can hava access to the database in the handler func

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc { //API func that gets handler as an parameter, returns Handler func
	return func(w http.ResponseWriter, r *http.Request) { // beauty of this is now we have all the access to the database in the regular handler func and by this we get the reusable code for our handlers
		ApiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get an api key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), ApiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get user by provided ApiKey: %v", err))
			return
		}

		handler(w, r, user) //func call with the parameter user, tipical example of this middleware implementation is handler_user.go  in the getUserHandler
	}
}
