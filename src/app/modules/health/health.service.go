package health

import "github.com/lautarok/hexa/src/app/modules/health/dto"

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (service *HealthService) GetHealth() (*dto.GetHealthDto, error) {
	return &dto.GetHealthDto{
		StatusCode:       200,
		Message:          "Alive! :D",
		ProcessingTimeMs: 0,
	}, nil
}
