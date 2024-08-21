package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/2Nemanja/RSSAggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	log.Println("pokusaj kreiranja")
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
		respondWithError(w, 400, fmt.Sprintf("Coultn't create a new user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFollowToFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFolloewedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	followed_feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
		respondWithError(w, 400, fmt.Sprintf("Coultn't create a new user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(followed_feeds))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID_str := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowID_str)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed_id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteaFeedFollow(r.Context(), database.DeleteaFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
