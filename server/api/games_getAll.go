package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"net/http"
)

// @Tags Games
// @Router /api/games [get]
///
// @Produce json
// @Success 200 {object} GetAllGamesResponse
// @Failure 400 {object} ErrorResponse
func getAllGames(w http.ResponseWriter, _ *http.Request) {
	var games map[string]*types.GameDefinition
	var err error

	if games, err = database.GetAllGames(); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, GetAllGamesResponse{
		Games: games,
	})
	return
}
