package vqcconfigusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVQCDTOToVQC(t *testing.T) {
	vqcDTO := ValidVQCDTO()
	vqc := validVQC()
	vqcBuilded, err := buildVQCDTOToVQC(vqcDTO)
	assert.NoError(t, err)
	assert.Equal(t, vqc, vqcBuilded)
}

func TestBuildVQCToVQCDTO(t *testing.T) {
	vqcOutputDTO := ValidVQCDTO()
	vqc := validVQC()
	vqcOutputDTOBuilded := buildVQCToVQCDTO(vqc)
	assert.Equal(t, vqcOutputDTO, vqcOutputDTOBuilded)
}
