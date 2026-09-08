package vqcconfigusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVQC(t *testing.T) {
	vqcInputDTO := validVQCInput()
	vqc := validVQC()
	vqcBuilded, err := buildVQC(vqcInputDTO)
	assert.NoError(t, err)
	assert.Equal(t, vqc, vqcBuilded)
}

func TestBuildVQCOutput(t *testing.T) {
	vqcOutputDTO := ValidVQCOutputDTO()
	vqc := validVQC()
	vqcOutputDTOBuilded := buildVQCOutput(vqc)
	assert.Equal(t, vqcOutputDTO, vqcOutputDTOBuilded)
}
