package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t             *testing.T
	service       *VQCConfigService
	userRepo      *testkit.MockUserRepository
	vqcConfigRepo *testkit.MockVQCConfigRepository
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	vqcConfigRepo := testkit.NewMockVQCConfigRepository()
	service := NewVQCConfigService(vqcConfigRepo, userRepo)

	return &testFixture{
		t:             t,
		service:       service,
		userRepo:      userRepo,
		vqcConfigRepo: vqcConfigRepo,
	}
}

func (f *testFixture) createUser(role user.Role) *user.User {
	makeUser := testkit.DefaultUser()
	user := makeUser()
	user.SetRole(role)
	err := f.userRepo.Save(user)
	require.NoError(f.t, err)
	return user
}

func (f *testFixture) createVQCConfig(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	makeVQCConfig := testkit.DefaultVQCConfig()
	vqcConfig := makeVQCConfig(ownerID)
	err := f.vqcConfigRepo.Save(vqcConfig)
	require.NoError(f.t, err)
	return vqcConfig
}

func validVQCInput() VQCInputDTO {
	return VQCInputDTO{
		NumQubits: 1,
		NumLayers: 1,
		Embedding: EmbeddingInputDTO{
			EmbeddingType: "angle",
			Qubits:        []uint{0},
			Rotation:      "x",
		},
		Measurement: MeasurementInputDTO{
			MeasurementType:     "expectation",
			MeasurementRotation: "x",
			Qubits:              []uint{0},
		},
	}
}

func validVQC() vqc.VQC {
	qubitZero, err := vqc.NewQubit(0, 1)
	if err != nil {
		panic(err)
	}
	qubits := []vqc.Qubit{qubitZero}
	embedding, err := vqc.NewAngleEmbedding(qubits, vqc.XRotation)
	if err != nil {
		panic(err)
	}
	measurement, err := vqc.NewMeasurement(qubits, vqc.ExpectationMeasurement, vqc.XMeasurementRotation)
	if err != nil {
		panic(err)
	}
	input := vqc.VQCBaseInput{
		NumQubits:   1,
		NumLayers:   1,
		Embedding:   embedding,
		Measurement: measurement,
	}
	vqc, err := vqc.NewVQC(input)
	if err != nil {
		panic(err)
	}
	return vqc
}

func ValidVQCOutputDTO() VQCOutputDTO {
	return VQCOutputDTO{
		NumQubits: 1,
		NumLayers: 1,
		Embedding: EmbeddingInputDTO{
			EmbeddingType: "angle",
			Qubits:        []uint{0},
			Rotation:      "x",
		},
		Measurement: MeasurementInputDTO{
			MeasurementType:     "expectation",
			MeasurementRotation: "x",
			Qubits:              []uint{0},
		},
		PreLayer:  LayerInputDTO{Gates: []QuantumGateInputDTO{}},
		Layer:     LayerInputDTO{Gates: []QuantumGateInputDTO{}},
		PostLayer: LayerInputDTO{Gates: []QuantumGateInputDTO{}},
	}
}
