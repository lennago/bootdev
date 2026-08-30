package main

import (
	"net/http"

	"github.com/lennago/bootdotdev_chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefreshJWT(w http.ResponseWriter, r *http.Request) {
	type rsp struct {
		Token string `json:"token"`
	}
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}
	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't fetch user for refresh token", err)
		return
	}
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, jwtExpirationTime)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}
	respondWithJSON(w, http.StatusOK, rsp{
		Token: token,
	})
}

func (cfg *apiConfig) handlerRevokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}
	if _, err := cfg.db.RevokeRefreshToken(r.Context(), refreshToken); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke session", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
