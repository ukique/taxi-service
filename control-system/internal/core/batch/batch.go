package batch

import (
	"context"
	"log"
	"time"

	"github.com/ukique/taxi-service/internal/models"
)

type LocationRepository interface {
	SaveLocationBatch(ctx context.Context, events []models.OrderCoordinateEvent) error
}
type Batch struct {
	batchSize          int           // how many message maximum can be in Batch.
	waitTimeout        time.Duration // how many milliseconds batch waits a new coordinate before writing.
	coordinatesChan    chan models.OrderCoordinateEvent
	locationRepository LocationRepository
}

func NewBatch(batchSize int, waitTimeout time.Duration, coordinatesChan chan models.OrderCoordinateEvent, locationRepository LocationRepository) *Batch {
	return &Batch{
		batchSize:          batchSize,
		waitTimeout:        waitTimeout,
		coordinatesChan:    coordinatesChan,
		locationRepository: locationRepository,
	}
}

func (b *Batch) CoordinatesBatch() {
	var events []models.OrderCoordinateEvent

	timer := time.NewTimer(b.waitTimeout)
	defer timer.Stop()

	flush := func() {
		if len(events) == 0 {
			return
		}
		if err := b.locationRepository.SaveLocationBatch(context.Background(), events); err != nil {
			log.Println("failed to SaveLocationBatch:", err)
		}
		events = nil
	}

	for {
		select {
		case event := <-b.coordinatesChan:
			events = append(events, event)

			if len(events) >= b.batchSize {
				flush()
				// reset timer since we just flushed
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(b.waitTimeout)
			}

		case <-timer.C:
			flush()
			timer.Reset(b.waitTimeout)
		}
	}
}
