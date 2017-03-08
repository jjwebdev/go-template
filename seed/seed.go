package seed

import (
	"github.com/jjwebdev/go-template/models"
)

type SeedService struct {
	models.UserService
}

func NewSeedService(us models.UserService) *SeedService {
	return &SeedService{
		us,
	}
}
