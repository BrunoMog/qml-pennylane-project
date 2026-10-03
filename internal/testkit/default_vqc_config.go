package testkit

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"strconv"

	"testing"
	"uuid"

	"github.com/stretchr/testify/require"
)

func DefaultVQCConfig(t *testing.T) func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	t.Helper()
	count := 0

	return func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
		t.Helper()
		count++

		name, err := vqcconfig.NewName("Test VQC Config " + strconv.Itoa(count))
		require.NoError(t, err)
		description, err := vqcconfig.NewDescription("Test VQC Config Description " + strconv.Itoa(count))
		require.NoError(t, err)
		vqcConfig, err := vqcconfig.NewVQCConfig(
			ownerID,
			name,
			description,
			ValidVQC(uint(count)+2, t),
		)
		require.NoError(t, err)
		return vqcConfig
	}
}

func ValidVQC(count uint, t *testing.T) vqc.VQC {
	t.Helper()

	qubitZero, err := vqc.NewQubit(0, count)
	require.NoError(t, err)
	qubits := []vqc.Qubit{qubitZero}
	embedding, err := vqc.NewAngleEmbedding(qubits, vqc.XRotation)
	require.NoError(t, err)
	measurement, err := vqc.NewMeasurement(qubits, vqc.ExpectationMeasurement, vqc.XMeasurementRotation)
	require.NoError(t, err)
	input := vqc.VQCBaseInput{
		NumQubits:   count,
		NumLayers:   0,
		Embedding:   embedding,
		Measurement: measurement,
	}
	newVQC, err := vqc.NewVQC(input)
	require.NoError(t, err)
	return newVQC
}
