package location

import (
	"github.com/gin-gonic/gin"
	"github.com/limadlp/estados-cidades-api/internal/app/handlers/locations/dto"
	repositories "github.com/limadlp/estados-cidades-api/internal/infrastructure/repositories/location"
)

type Location struct {
	locationRepository *repositories.LocationRepository
}

func NewLocationHandler(locationRepository *repositories.LocationRepository) *Location {
	return &Location{
		locationRepository: repositories.NewLocationRepository(),
	}
}

func (l *Location) GetAllStates(c *gin.Context) {
	states, err := l.locationRepository.GetStates()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var statesResponse []dto.StateResponse

	for _, s := range states {
		statesResponse = append(statesResponse, dto.StateResponse{
			Acronym: s.Acronym,
			Name:    s.Name,
		})

	}
	c.JSON(200, statesResponse)

}
