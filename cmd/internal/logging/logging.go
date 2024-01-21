package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/aidk/gocatfacts/cmd/internal/api/service"
	"github.com/aidk/gocatfacts/cmd/internal/api/types"
)

// LoggingService is a simple implementation of Service that logs requests and responses.
type LoggingService struct {
	svc service.Service
}

// NewLoggingService returns a new LoggingService.
func NewLoggingService(svc service.Service) service.Service {
	return &LoggingService{
		svc: svc,
	}
}

// GetFact logs the request and response of the GetFact method.
func (s *LoggingService) GetFact(ctx context.Context) (fact *types.Fact, err error) {

	// we defer the logging of the request and response
	defer func(start time.Time) {
		fmt.Printf("======== LOGS ========\n%v\n", time.Now())
		fmt.Printf("fact: %s\nerr: %v\ntime: %v\n", fact.Message, err, time.Since(start))
	}(time.Now()) // we pass the current time to calculate the time taken by the request

	return s.svc.GetFact(ctx)
}
