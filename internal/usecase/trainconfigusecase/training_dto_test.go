package trainconfigusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTrainingFromDTO(t *testing.T) {
	dto := validTrainDTO()
	training, err := buildTrainingFromDTO(dto)
	assert.NoError(t, err)
	expectedTraining := validTrain()
	assert.True(t, expectedTraining.Equals(training))
}

func TestBuildTrainingDTOFromTraining(t *testing.T) {
	training := validTrain()
	dto := buildTrainingDTOFromTraining(training)
	expectedDTO := validTrainDTO()
	assert.Equal(t, expectedDTO, dto)
}
