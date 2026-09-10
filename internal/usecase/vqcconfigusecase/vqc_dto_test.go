package vqcconfigusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVQCFromDTO(t *testing.T) {
	vqcDTO := ValidVQCDTO()
	vqc := validVQC()
	vqcBuilded, err := buildVQCFromDTO(vqcDTO)
	assert.NoError(t, err)
	assert.Equal(t, vqc, vqcBuilded)
}

func TestBuildVQCDTOFromVQC(t *testing.T) {
	vqcOutputDTO := ValidVQCDTO()
	vqc := validVQC()
	vqcOutputDTOBuilded := buildVQCDTOFromVQC(vqc)
	assert.Equal(t, vqcOutputDTO, vqcOutputDTOBuilded)
}
