package vqc

import (
	"testing"
)

func validAngleEmbeddingInput() EmbeddingBuilderInput {
	return EmbeddingBuilderInput{
		Type:      "angle",
		Qubits:    []uint{0, 1, 2},
		Rotation:  "x",
		Normalize: false,
		PadWith:   0.0,
	}
}

func validAmplitudeEmbeddingInput() EmbeddingBuilderInput {
	return EmbeddingBuilderInput{
		Type:      "amplitude",
		Qubits:    []uint{0, 1},
		Rotation:  "z",
		Normalize: true,
		PadWith:   0.0,
	}
}

func validMeasurementInput() MeasurementBuilderInput {
	return MeasurementBuilderInput{
		Type:     "expectation",
		Rotation: "x",
		Qubits:   []uint{0},
	}
}

func validVQCBuilderInputAngle() VQCBuilderInput {
	return VQCBuilderInput{
		NumQubits:   3,
		NumLayers:   2,
		Embedding:   validAngleEmbeddingInput(),
		Measurement: validMeasurementInput(),
	}
}

func validVQCBuilderInputAmplitude() VQCBuilderInput {
	return VQCBuilderInput{
		NumQubits:   2,
		NumLayers:   1,
		Embedding:   validAmplitudeEmbeddingInput(),
		Measurement: validMeasurementInput(),
	}
}

func validGateInput(qubitIndex uint) GateBuilderInput {
	return GateBuilderInput{
		GateType:      "rx",
		Qubit:         qubitIndex,
		ControlQubits: []uint{},
	}
}

func validControlledGateInput(targetQubit, controlQubit uint) GateBuilderInput {
	return GateBuilderInput{
		GateType:      "cnot",
		Qubit:         targetQubit,
		ControlQubits: []uint{controlQubit},
	}
}

func TestNewVQCBuilder(t *testing.T) {
	testCases := []struct {
		name      string
		input     VQCBuilderInput
		expectErr bool
	}{
		{
			name:      "Valid angle embedding builder",
			input:     validVQCBuilderInputAngle(),
			expectErr: false,
		},
		{
			name:      "Valid amplitude embedding builder",
			input:     validVQCBuilderInputAmplitude(),
			expectErr: false,
		},
		{
			name: "Invalid embedding type",
			input: VQCBuilderInput{
				NumQubits: 2,
				NumLayers: 1,
				Embedding: EmbeddingBuilderInput{
					Type:   "invalid_type",
					Qubits: []uint{0, 1},
				},
				Measurement: validMeasurementInput(),
			},
			expectErr: true,
		},
		{
			name: "Invalid qubit index in embedding",
			input: VQCBuilderInput{
				NumQubits: 2,
				NumLayers: 1,
				Embedding: EmbeddingBuilderInput{
					Type:     "angle",
					Qubits:   []uint{0, 5},
					Rotation: "x",
				},
				Measurement: validMeasurementInput(),
			},
			expectErr: true,
		},
		{
			name: "Duplicate qubits in embedding",
			input: VQCBuilderInput{
				NumQubits: 3,
				NumLayers: 1,
				Embedding: EmbeddingBuilderInput{
					Type:     "angle",
					Qubits:   []uint{0, 1, 1},
					Rotation: "x",
				},
				Measurement: validMeasurementInput(),
			},
			expectErr: true,
		},
		{
			name: "Invalid qubit index in measurement",
			input: VQCBuilderInput{
				NumQubits: 2,
				NumLayers: 1,
				Embedding: validAngleEmbeddingInput(),
				Measurement: MeasurementBuilderInput{
					Type:   "expectation",
					Qubits: []uint{5},
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder, err := NewVQCBuilder(tc.input)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewVQCBuilder() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if !tc.expectErr && builder == nil {
				t.Error("NewVQCBuilder() returned nil builder")
			}
		})
	}
}

func TestVQCBuilderWithPreLayer(t *testing.T) {
	builder, _ := NewVQCBuilder(validVQCBuilderInputAngle())

	testCases := []struct {
		name      string
		gates     []GateBuilderInput
		expectErr bool
	}{
		{
			name:      "Valid single gate",
			gates:     []GateBuilderInput{validGateInput(0)},
			expectErr: false,
		},
		{
			name: "Valid multiple gates",
			gates: []GateBuilderInput{
				validGateInput(0),
				validGateInput(1),
				validGateInput(2),
			},
			expectErr: false,
		},
		{
			name:      "Empty gates",
			gates:     []GateBuilderInput{},
			expectErr: false,
		},
		{
			name: "Invalid gate qubit index",
			gates: []GateBuilderInput{
				{GateType: "rx", Qubit: 10},
			},
			expectErr: true,
		},
		{
			name: "Valid controlled gate",
			gates: []GateBuilderInput{
				validControlledGateInput(0, 1),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := builder.WithPreLayer(tc.gates)
			if (err != nil) != tc.expectErr {
				t.Errorf("WithPreLayer() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if !tc.expectErr && result == nil {
				t.Error("WithPreLayer() returned nil")
			}
		})
	}
}

func TestVQCBuilderWithLayer(t *testing.T) {
	builder, _ := NewVQCBuilder(validVQCBuilderInputAngle())

	testCases := []struct {
		name      string
		gates     []GateBuilderInput
		expectErr bool
	}{
		{
			name:      "Valid single gate",
			gates:     []GateBuilderInput{validGateInput(1)},
			expectErr: false,
		},
		{
			name: "Valid multiple gates",
			gates: []GateBuilderInput{
				validGateInput(0),
				validGateInput(1),
				validGateInput(2),
			},
			expectErr: false,
		},
		{
			name:      "Empty gates",
			gates:     []GateBuilderInput{},
			expectErr: false,
		},
		{
			name: "Invalid gate qubit index",
			gates: []GateBuilderInput{
				{GateType: "ry", Qubit: 5},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := builder.WithLayer(tc.gates)
			if (err != nil) != tc.expectErr {
				t.Errorf("WithLayer() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if !tc.expectErr && result == nil {
				t.Error("WithLayer() returned nil")
			}
		})
	}
}

func TestVQCBuilderWithPostLayer(t *testing.T) {
	builder, _ := NewVQCBuilder(validVQCBuilderInputAngle())

	testCases := []struct {
		name      string
		gates     []GateBuilderInput
		expectErr bool
	}{
		{
			name:      "Valid single gate",
			gates:     []GateBuilderInput{validGateInput(2)},
			expectErr: false,
		},
		{
			name: "Valid multiple gates",
			gates: []GateBuilderInput{
				validGateInput(0),
				validGateInput(1),
				validGateInput(2),
			},
			expectErr: false,
		},
		{
			name:      "Empty gates",
			gates:     []GateBuilderInput{},
			expectErr: false,
		},
		{
			name: "Invalid gate qubit index",
			gates: []GateBuilderInput{
				{GateType: "rz", Qubit: 8},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := builder.WithPostLayer(tc.gates)
			if (err != nil) != tc.expectErr {
				t.Errorf("WithPostLayer() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if !tc.expectErr && result == nil {
				t.Error("WithPostLayer() returned nil")
			}
		})
	}
}

func TestVQCBuilderBuild(t *testing.T) {
	testCases := []struct {
		builderSetupFn  func() (*VQCBuilder, error)
		name            string
		expectErr       bool
		shouldReturnVQC bool
	}{
		{
			name: "Build with all layers",
			builderSetupFn: func() (*VQCBuilder, error) {
				builder, err := NewVQCBuilder(validVQCBuilderInputAngle())
				if err != nil {
					return nil, err
				}
				builder.WithPreLayer([]GateBuilderInput{validGateInput(0)})
				builder.WithLayer([]GateBuilderInput{validGateInput(1)})
				builder.WithPostLayer([]GateBuilderInput{validGateInput(2)})
				return builder, nil
			},
			expectErr:       false,
			shouldReturnVQC: true,
		},
		{
			name: "Build without pre layer",
			builderSetupFn: func() (*VQCBuilder, error) {
				builder, err := NewVQCBuilder(validVQCBuilderInputAngle())
				if err != nil {
					return nil, err
				}
				builder.WithLayer([]GateBuilderInput{validGateInput(0)})
				builder.WithPostLayer([]GateBuilderInput{validGateInput(1)})
				return builder, nil
			},
			expectErr:       false,
			shouldReturnVQC: true,
		},
		{
			name: "Build without post layer",
			builderSetupFn: func() (*VQCBuilder, error) {
				builder, err := NewVQCBuilder(validVQCBuilderInputAngle())
				if err != nil {
					return nil, err
				}
				builder.WithPreLayer([]GateBuilderInput{validGateInput(0)})
				builder.WithLayer([]GateBuilderInput{validGateInput(1)})
				return builder, nil
			},
			expectErr:       false,
			shouldReturnVQC: true,
		},
		{
			name: "Build without any layers",
			builderSetupFn: func() (*VQCBuilder, error) {
				return NewVQCBuilder(validVQCBuilderInputAngle())
			},
			expectErr:       false,
			shouldReturnVQC: true,
		},
		{
			name: "Build with amplitude embedding",
			builderSetupFn: func() (*VQCBuilder, error) {
				builder, err := NewVQCBuilder(validVQCBuilderInputAmplitude())
				if err != nil {
					return nil, err
				}
				builder.WithLayer([]GateBuilderInput{validGateInput(0)})
				return builder, nil
			},
			expectErr:       false,
			shouldReturnVQC: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder, setupErr := tc.builderSetupFn()
			if setupErr != nil && !tc.expectErr {
				t.Fatalf("Setup failed: %v", setupErr)
			}

			if builder == nil {
				t.Fatal("Builder setup returned nil")
			}

			vqc, err := builder.Build()
			if (err != nil) != tc.expectErr {
				t.Errorf("Build() error = %v, expectErr = %v", err, tc.expectErr)
			}
			if tc.shouldReturnVQC && vqc == nil {
				t.Error("Build() should return VQC, got nil")
			}
		})
	}
}

func TestVQCBuilderFluentInterface(t *testing.T) {
	builder, err := NewVQCBuilder(validVQCBuilderInputAngle())
	if err != nil {
		t.Fatalf("NewVQCBuilder() failed: %v", err)
	}

	gates := []GateBuilderInput{
		validGateInput(0),
		validGateInput(1),
		validGateInput(2),
	}

	result, err := builder.WithPreLayer(gates)
	if err != nil {
		t.Errorf("WithPreLayer chaining failed: %v", err)
		return
	}

	result, err = result.WithLayer(gates)
	if err != nil {
		t.Errorf("WithLayer chaining failed: %v", err)
		return
	}

	result, err = result.WithPostLayer(gates)
	if err != nil {
		t.Errorf("WithPostLayer chaining failed: %v", err)
		return
	}

	if result == nil {
		t.Error("Fluent interface returned nil")
	}

	vqc, err := result.Build()
	if err != nil {
		t.Errorf("Build after chaining failed: %v", err)
	}
	if vqc == nil {
		t.Error("Build should return VQC, got nil")
	}
}

func TestVQCBuilderErrorPropagation(t *testing.T) {
	testCases := []struct {
		setupFn     func() (*VQCBuilder, error)
		layerFn     func(*VQCBuilder) (*VQCBuilder, error)
		name        string
		description string
		expectErr   bool
	}{
		{
			name: "Invalid gate in pre layer propagates error",
			setupFn: func() (*VQCBuilder, error) {
				return NewVQCBuilder(validVQCBuilderInputAngle())
			},
			layerFn: func(b *VQCBuilder) (*VQCBuilder, error) {
				return b.WithPreLayer([]GateBuilderInput{
					{GateType: "RX", Qubit: 10}, // Out of range
				})
			},
			expectErr:   true,
			description: "Should propagate qubit index error from pre layer",
		},
		{
			name: "Invalid gate in main layer propagates error",
			setupFn: func() (*VQCBuilder, error) {
				return NewVQCBuilder(validVQCBuilderInputAngle())
			},
			layerFn: func(b *VQCBuilder) (*VQCBuilder, error) {
				return b.WithLayer([]GateBuilderInput{
					{GateType: "RY", Qubit: 20}, // Out of range
				})
			},
			expectErr:   true,
			description: "Should propagate qubit index error from main layer",
		},
		{
			name: "Invalid gate in post layer propagates error",
			setupFn: func() (*VQCBuilder, error) {
				return NewVQCBuilder(validVQCBuilderInputAngle())
			},
			layerFn: func(b *VQCBuilder) (*VQCBuilder, error) {
				return b.WithPostLayer([]GateBuilderInput{
					{GateType: "RZ", Qubit: 15}, // Out of range
				})
			},
			expectErr:   true,
			description: "Should propagate qubit index error from post layer",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder, setupErr := tc.setupFn()
			if setupErr != nil && !tc.expectErr {
				t.Fatalf("Setup failed: %v", setupErr)
			}

			_, err := tc.layerFn(builder)
			if (err != nil) != tc.expectErr {
				t.Errorf("%s: error = %v, expectErr = %v", tc.description, err, tc.expectErr)
			}
		})
	}
}

func TestVQCBuilderMultipleCallSequences(t *testing.T) {
	testCases := []struct {
		name           string
		preLayerCalls  int
		mainLayerCalls int
		postLayerCalls int
		expectErr      bool
	}{
		{
			name:           "Single call per layer",
			preLayerCalls:  1,
			mainLayerCalls: 1,
			postLayerCalls: 1,
			expectErr:      false,
		},
		{
			name:           "Multiple calls per layer (last wins)",
			preLayerCalls:  3,
			mainLayerCalls: 2,
			postLayerCalls: 2,
			expectErr:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder, _ := NewVQCBuilder(validVQCBuilderInputAngle())

			for i := 0; i < tc.preLayerCalls; i++ {
				_, err := builder.WithPreLayer([]GateBuilderInput{validGateInput(0)})
				if (err != nil) != tc.expectErr {
					t.Errorf("WithPreLayer call %d failed", i+1)
				}
			}

			for i := 0; i < tc.mainLayerCalls; i++ {
				_, err := builder.WithLayer([]GateBuilderInput{validGateInput(1)})
				if (err != nil) != tc.expectErr {
					t.Errorf("WithLayer call %d failed", i+1)
				}
			}

			for i := 0; i < tc.postLayerCalls; i++ {
				_, err := builder.WithPostLayer([]GateBuilderInput{validGateInput(2)})
				if (err != nil) != tc.expectErr {
					t.Errorf("WithPostLayer call %d failed", i+1)
				}
			}

			_, err := builder.Build()
			if (err != nil) != tc.expectErr {
				t.Errorf("Build failed: %v", err)
			}
		})
	}
}

func TestVQCBuilderEdgeCases(t *testing.T) {
	testCases := []struct {
		name      string
		embedding EmbeddingBuilderInput
		numQubits uint
		numLayers uint
		expectErr bool
	}{
		{
			name:      "Single qubit",
			numQubits: 1,
			numLayers: 1,
			embedding: EmbeddingBuilderInput{
				Type:     "angle",
				Qubits:   []uint{0},
				Rotation: "x",
			},
			expectErr: false,
		},
		{
			name:      "Large number of qubits",
			numQubits: 100,
			numLayers: 10,
			embedding: EmbeddingBuilderInput{
				Type:     "angle",
				Qubits:   []uint{0, 1, 2, 3, 4},
				Rotation: "y",
			},
			expectErr: false,
		},
		{
			name:      "Zero layers",
			numQubits: 3,
			numLayers: 0,
			embedding: validAngleEmbeddingInput(),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := VQCBuilderInput{
				NumQubits: tc.numQubits,
				NumLayers: tc.numLayers,
				Embedding: tc.embedding,
				Measurement: MeasurementBuilderInput{
					Type:     "expectation",
					Rotation: "x",
					Qubits:   []uint{0},
				},
			}

			builder, err := NewVQCBuilder(input)
			if (err != nil) != tc.expectErr {
				t.Errorf("NewVQCBuilder() error = %v, expectErr = %v", err, tc.expectErr)
				return
			}

			if !tc.expectErr && builder == nil {
				t.Error("Expected builder, got nil")
			}
		})
	}
}
