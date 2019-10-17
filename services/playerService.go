package services

import (
	"github.com/my-app/models"
	"github.com/my-app/repositories"
)

// IPlayerService is
type IPlayerService interface {
	StorePlayer(models.Player) (models.Player, error)
}

// PlayerService is
type PlayerService struct {
	PlayerRepository repositories.IPlayerRepository
}

// StorePlayer is
func (p *PlayerService) StorePlayer(data models.Player) (result models.Player, err error) {
	result, err = p.PlayerRepository.StorePlayer(data)
	return result, err
}
