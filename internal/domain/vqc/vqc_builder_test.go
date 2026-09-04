package vqc

import "testing"

func validVQCBuilderInput() VQCBuilderInput {
	return VQCBuilderInput{
		NumQubits: 3,
		NumLayers: 2,
		Embedding: EmbeddingBuilderInput{
			Type:     EmbeddingTypeAngle,
			Rotation: XRotation,
			Qubits:   []Qubit{0, 1, 2},
		},
		Measurement: MeasurementBuilderInput{
			Type:     ExpectationMeasurement,
			Rotation: XMeasurementRotation,
			Qubits:   []Qubit{0},
		},
	}
}

func newTestGate(t *testing.T, gateType GateType, qubit Qubit, controls []Qubit) QuantumGate {
	t.Helper()
	gate, err := NewQuantumGate(gateType, qubit, controls)
	if err != nil {
		t.Fatalf("NewQuantumGate() returned an unexpected error: %v", err)
	}
	return *gate
}

func TestNewVQCBuilder(t *testing.T) {
	testCases := []struct {
		name      string
		input     VQCBuilderInput
		expectErr bool
	}{
		{name: "angle embedding", input: validVQCBuilderInput()},
		{
			name: "amplitude embedding",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.NumQubits = 2
				input.Embedding = EmbeddingBuilderInput{
					Type:      EmbeddingTypeAmplitude,
					Qubits:    []Qubit{0, 1},
					Normalize: true,
				}
				return input
			}(),
		},
		{
			name: "invalid embedding type",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Embedding.Type = EmbeddingType("invalid")
				return input
			}(),
			expectErr: true,
		},
		{
			name: "invalid embedding rotation",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Embedding.Rotation = EmbeddingRotation("invalid")
				return input
			}(),
			expectErr: true,
		},
		{
			name: "invalid measurement type",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Measurement.Type = MeasurementType("invalid")
				return input
			}(),
			expectErr: true,
		},
		{
			name: "invalid measurement rotation",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Measurement.Rotation = MeasurementRotation("invalid")
				return input
			}(),
			expectErr: true,
		},
		{
			name: "zero embedding qubits",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Embedding.Qubits = nil
				return input
			}(),
			expectErr: true,
		},
		{
			name: "zero measurement qubits",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.Measurement.Qubits = nil
				return input
			}(),
			expectErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder, err := NewVQCBuilder(testCase.input)
			if (err != nil) != testCase.expectErr {
				t.Fatalf("NewVQCBuilder() error = %v, expectErr = %v", err, testCase.expectErr)
			}
			if !testCase.expectErr && builder == nil {
				t.Fatal("NewVQCBuilder() returned nil builder")
			}
		})
	}
}

func TestVQCBuilderLayerMethods(t *testing.T) {
	testCases := []struct {
		name      string
		configure func(*VQCBuilder, []QuantumGate) (*VQCBuilder, error)
		layer     func(*VQCBuilder) Layer
	}{
		{name: "pre layer", configure: func(b *VQCBuilder, g []QuantumGate) (*VQCBuilder, error) { return b.WithPreLayer(g) }, layer: func(b *VQCBuilder) Layer { return b.preLayer }},
		{name: "main layer", configure: func(b *VQCBuilder, g []QuantumGate) (*VQCBuilder, error) { return b.WithLayer(g) }, layer: func(b *VQCBuilder) Layer { return b.layer }},
		{name: "post layer", configure: func(b *VQCBuilder, g []QuantumGate) (*VQCBuilder, error) { return b.WithPostLayer(g) }, layer: func(b *VQCBuilder) Layer { return b.postLayer }},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder, err := NewVQCBuilder(validVQCBuilderInput())
			if err != nil {
				t.Fatalf("NewVQCBuilder() returned an unexpected error: %v", err)
			}
			gates := []QuantumGate{newTestGate(t, RXGate, 0, nil), newTestGate(t, CNOTGate, 1, []Qubit{2})}

			result, err := testCase.configure(builder, gates)
			if err != nil || result != builder {
				t.Fatalf("layer method returned builder=%v, error=%v", result == builder, err)
			}
			stored := testCase.layer(builder).Gates()
			if len(stored) != len(gates) || !stored[0].Equal(gates[0]) || !stored[1].Equal(gates[1]) {
				t.Fatalf("layer did not store the supplied gates")
			}

			if _, err := testCase.configure(builder, nil); err != nil {
				t.Fatalf("empty layer returned an unexpected error: %v", err)
			}
			if len(testCase.layer(builder).Gates()) != 0 {
				t.Error("empty layer did not replace the previous gates")
			}
		})
	}
}

func TestVQCBuilderBuild(t *testing.T) {
	testCases := []struct {
		name      string
		input     VQCBuilderInput
		configure func(*VQCBuilder) error
		expectErr bool
	}{
		{name: "without layers", input: validVQCBuilderInput()},
		{
			name:  "with all layers",
			input: validVQCBuilderInput(),
			configure: func(builder *VQCBuilder) error {
				gates := []QuantumGate{newTestGate(t, RXGate, 0, nil)}
				if _, err := builder.WithPreLayer(gates); err != nil {
					return err
				}
				if _, err := builder.WithLayer(gates); err != nil {
					return err
				}
				_, err := builder.WithPostLayer(gates)
				return err
			},
		},
		{
			name: "zero qubits fails at build",
			input: func() VQCBuilderInput {
				input := validVQCBuilderInput()
				input.NumQubits = 0
				return input
			}(),
			expectErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder, err := NewVQCBuilder(testCase.input)
			if err != nil {
				t.Fatalf("NewVQCBuilder() returned an unexpected error: %v", err)
			}
			if testCase.configure != nil {
				if err := testCase.configure(builder); err != nil {
					t.Fatalf("builder configuration returned an unexpected error: %v", err)
				}
			}

			result, err := builder.Build()
			if (err != nil) != testCase.expectErr {
				t.Fatalf("Build() error = %v, expectErr = %v", err, testCase.expectErr)
			}
			if !testCase.expectErr {
				if result == nil {
					t.Fatal("Build() returned nil VQC")
				}
				if result.NumQubits() != testCase.input.NumQubits || result.NumLayers() != testCase.input.NumLayers {
					t.Error("Build() did not preserve the VQC dimensions")
				}
			}
		})
	}
}

func TestVQCBuilderFluentInterface(t *testing.T) {
	builder, err := NewVQCBuilder(validVQCBuilderInput())
	if err != nil {
		t.Fatalf("NewVQCBuilder() returned an unexpected error: %v", err)
	}
	gates := []QuantumGate{newTestGate(t, RXGate, 0, nil)}

	result, err := builder.WithPreLayer(gates)
	if err != nil {
		t.Fatalf("WithPreLayer() returned an unexpected error: %v", err)
	}
	result, err = result.WithLayer(gates)
	if err != nil {
		t.Fatalf("WithLayer() returned an unexpected error: %v", err)
	}
	result, err = result.WithPostLayer(gates)
	if err != nil {
		t.Fatalf("WithPostLayer() returned an unexpected error: %v", err)
	}
	if result != builder {
		t.Error("layer methods did not preserve fluent builder identity")
	}
	if _, err := result.Build(); err != nil {
		t.Fatalf("Build() after fluent configuration returned an unexpected error: %v", err)
	}
}

func TestVQCBuilderLastLayerConfigurationWins(t *testing.T) {
	builder, err := NewVQCBuilder(validVQCBuilderInput())
	if err != nil {
		t.Fatalf("NewVQCBuilder() returned an unexpected error: %v", err)
	}
	firstGate := newTestGate(t, RXGate, 0, nil)
	secondGate := newTestGate(t, RYGate, 1, nil)
	if _, err := builder.WithLayer([]QuantumGate{firstGate}); err != nil {
		t.Fatalf("first WithLayer() returned an unexpected error: %v", err)
	}
	if _, err := builder.WithLayer([]QuantumGate{secondGate}); err != nil {
		t.Fatalf("second WithLayer() returned an unexpected error: %v", err)
	}

	stored := builder.layer.Gates()
	if len(stored) != 1 || !stored[0].Equal(secondGate) {
		t.Errorf("last WithLayer() did not replace the previous layer")
	}
}
