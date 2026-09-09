package testkit

import (
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"strconv"

	"uuid"
)

func DefaultVQCConfig() func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	count := 0
	return func(ownerID uuid.UUID) *vqcconfig.VQCConfig {
		count++
		name, err := vqcconfig.NewName("Test VQC Config " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		description, err := vqcconfig.NewDescription("Test VQC Config Description " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		vqcConfig, err := vqcconfig.NewVQCConfig(
			ownerID,
			name,
			description,
			ValidVQC(uint(count)+2),
		)
		if err != nil {
			panic(err)
		}
		return vqcConfig
	}
}

func ValidVQC(count uint) vqc.VQC {
	qubitZero, err := vqc.NewQubit(0, count)
	if err != nil {
		panic(err)
	}
	qubits := []vqc.Qubit{qubitZero}
	embedding, error := vqc.NewAngleEmbedding(qubits, vqc.XRotation)
	if error != nil {
		panic(error)
	}
	measurement, error := vqc.NewMeasurement(qubits, vqc.ExpectationMeasurement, vqc.XMeasurementRotation)
	if error != nil {
		panic(error)
	}
	input := vqc.VQCBaseInput{
		NumQubits:   count,
		NumLayers:   0,
		Embedding:   embedding,
		Measurement: measurement,
	}
	newVQC, err := vqc.NewVQC(input)
	if err != nil {
		panic(err)
	}
	return newVQC
}
