package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateVQCConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) CreateVQCConfigInput
		testName      string
	}{
		{
			testName: "create VQCConfig successfully",
			setup: func(f *testFixture) CreateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQC:         &vqc.VQC{},
				}
			},
			expectedError: nil,
		},
		{
			testName: "fail to create VQCConfig due to non-existent user",
			setup: func(f *testFixture) CreateVQCConfigInput {
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    uuid.New(),
					VQC:         &vqc.VQC{},
				}
			},
			expectedError: &UserNotFoundError{},
		},
		{
			testName: "fail to create VQCConfig due to duplicate name",
			setup: func(f *testFixture) CreateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcconfig := f.createVQCConfig(user.ID())
				vqcconfig.SetName("Test Config")
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQC:         &vqc.VQC{},
				}
			},
			expectedError: &VQCConfigNameAlreadyExistsError{},
		},
		{
			testName: "nil reference to VQC",
			setup: func(f *testFixture) CreateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQC:         nil,
				}
			},
			expectedError: &vqcconfig.VQCConfigMissingVQCError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)

			output, err := f.service.CreateVQCConfig(input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, input.Name, output.Name)
				assert.Equal(t, input.Description, output.Description)
				assert.NotZero(t, output.VQCId)
				assert.NotZero(t, output.CreatedAt)
			}
		})
	}

}
