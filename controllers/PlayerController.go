package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-app/helpers"
	"github.com/my-app/infrastructures"
	"github.com/my-app/models"
	"github.com/my-app/repositories"
	"github.com/my-app/services"
)

// InitPlayerController is
func InitPlayerController() *PlayerController {
	playerRepository := new(repositories.PlayerRepository)
	playerRepository.DB = &infrastructures.SQLConnection{}

	playerService := new(services.PlayerService)
	playerService.PlayerRepository = playerRepository

	playerController := new(PlayerController)
	playerController.PlayerService = playerService

	return playerController
}

// PlayerController is
type PlayerController struct {
	PlayerService services.IPlayerService
}

// StorePlayer is
func (p *PlayerController) StorePlayer(res http.ResponseWriter, req *http.Request) {
	var player models.Player
	//Read request data
	body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(body, &player)

	if err != nil {
		helpers.Response(res, http.StatusBadRequest, "Failed read input data")
		return
	}

	result, err := p.PlayerService.StorePlayer(player)

	if err == nil {
		helpers.Response(res, http.StatusOK, result)
	} else {
		helpers.Response(res, http.StatusBadRequest, err)
	}

	return
}
