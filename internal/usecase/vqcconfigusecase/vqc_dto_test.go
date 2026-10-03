package vqcconfigusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVQCFromDTO(t *testing.T) {
	vqcDTO := validVQCDTO()
	vqc := validVQC(t)
	vqcBuilded, err := buildVQCFromDTO(vqcDTO)
	assert.NoError(t, err)
	assert.Equal(t, vqc, vqcBuilded)
}

func TestBuildVQCDTOFromVQC(t *testing.T) {
	vqcOutputDTO := validVQCDTO()
	vqc := validVQC(t)
	vqcOutputDTOBuilded := buildVQCDTOFromVQC(vqc)
	assert.Equal(t, vqcOutputDTO, vqcOutputDTOBuilded)
}
