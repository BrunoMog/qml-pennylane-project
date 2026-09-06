package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadVQCConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) LoadVQCConfigInput
		testName      string
	}{
		{
			testName: "load VQCConfig by ID successfully",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				vqcConfigID := vqcConfig.VQCConfigID()
				return LoadVQCConfigInput{
					CallerID:      user.ID(),
					VQCConfigID:   &vqcConfigID,
					VQCConfigName: nil,
				}
			},
			expectedError: nil,
		},
		{
			testName: "load VQCConfig by name successfully",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				vqcConfigName := vqcConfig.Name()
				return LoadVQCConfigInput{
					CallerID:      user.ID(),
					VQCConfigID:   nil,
					VQCConfigName: &vqcConfigName,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				vqcConfigID := vqcConfig.VQCConfigID()
				return LoadVQCConfigInput{
					CallerID:      uuid.New(),
					VQCConfigID:   &vqcConfigID,
					VQCConfigName: nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent VQCConfig by ID",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfigID := uuid.New()
				return LoadVQCConfigInput{
					CallerID:      user.ID(),
					VQCConfigID:   &vqcConfigID,
					VQCConfigName: nil,
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "inexistent VQCConfig by name",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfigName := "nonexistent-config"
				return LoadVQCConfigInput{
					CallerID:      user.ID(),
					VQCConfigID:   nil,
					VQCConfigName: &vqcConfigName,
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) LoadVQCConfigInput {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				vqcConfigID := vqcConfig.VQCConfigID()
				return LoadVQCConfigInput{
					CallerID:      unauthorizedUser.ID(),
					VQCConfigID:   &vqcConfigID,
					VQCConfigName: nil,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "both VQCConfigID and VQCConfigName are nil",
			setup: func(f *testFixture) LoadVQCConfigInput {
				user := f.createUser(user.RoleUser)
				return LoadVQCConfigInput{
					CallerID:      user.ID(),
					VQCConfigID:   nil,
					VQCConfigName: nil,
				}
			},
			expectedError: &InvalidInputError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			output, err := fixture.service.LoadVQCConfig(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, input.CallerID, output.OwnerID)
				if input.VQCConfigID != nil {
					assert.Equal(t, input.VQCConfigID, &output.VQCConfigID)
				}
				if input.VQCConfigName != nil {
					assert.Equal(t, *input.VQCConfigName, output.Name)
				}
			}
		})
	}
}

func TestLoadAllVQCConfigs(t *testing.T) {
	tests := []struct {
		setup              func(f *testFixture) LoadAllVQCConfigsInput
		testName           string
		expectedSizeOutput int
	}{
		{
			testName: "load all VQCConfigs successfully",
			setup: func(f *testFixture) LoadAllVQCConfigsInput {
				user := f.createUser(user.RoleUser)
				f.createVQCConfig(user.ID())
				f.createVQCConfig(user.ID())
				return LoadAllVQCConfigsInput{
					CallerID: user.ID(),
				}
			},
			expectedSizeOutput: 2,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) LoadAllVQCConfigsInput {
				return LoadAllVQCConfigsInput{
					CallerID: uuid.New(),
				}
			},
			expectedSizeOutput: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			output, err := fixture.service.LoadAllVQCConfigs(input)
			require.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, tt.expectedSizeOutput, len(output.VQCConfigs))
		})
	}
}
