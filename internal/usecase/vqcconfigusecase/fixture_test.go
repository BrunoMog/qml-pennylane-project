package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"uuid"

	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t             *testing.T
	service       *VQCConfigService
	userRepo      *testkit.MockUserRepository
	vqcConfigRepo *testkit.MockVQCConfigRepository
	makeUser      func() *user.User
	makeVQCConfig func(ownerID uuid.UUID) *vqcconfig.VQCConfig
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	vqcConfigRepo := testkit.NewMockVQCConfigRepository()
	service := NewVQCConfigService(vqcConfigRepo, userRepo)
	makeUser := testkit.DefaultUser()
	makeVQCConfig := testkit.DefaultVQCConfig()

	return &testFixture{
		t:             t,
		service:       service,
		userRepo:      userRepo,
		vqcConfigRepo: vqcConfigRepo,
		makeUser:      makeUser,
		makeVQCConfig: makeVQCConfig,
	}
}

func (f *testFixture) createUser(role user.Role) *user.User {
	user := f.makeUser()
	user.SetRole(role)
	err := f.userRepo.Save(user)
	require.NoError(f.t, err)
	return user
}

func (f *testFixture) createVQCConfig(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	vqcConfig := f.makeVQCConfig(ownerID)
	err := f.vqcConfigRepo.Save(vqcConfig)
	require.NoError(f.t, err)
	return vqcConfig
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

func ValidVQCDTO() VQCDTO {
	return VQCDTO{
		NumQubits: 1,
		NumLayers: 1,
		Embedding: EmbeddingDTO{
			EmbeddingType: "angle",
			Qubits:        []uint{0},
			Rotation:      "x",
		},
		Measurement: MeasurementDTO{
			MeasurementType:     "expectation",
			MeasurementRotation: "x",
			Qubits:              []uint{0},
		},
		PreLayer:  LayerDTO{Gates: []QuantumGateDTO{}},
		Layer:     LayerDTO{Gates: []QuantumGateDTO{}},
		PostLayer: LayerDTO{Gates: []QuantumGateDTO{}},
	}
}
