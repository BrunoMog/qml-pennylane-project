package vqc_config

type VQCConfigRepository interface {
	Save(config *VQCConfig) error
	FindByName(name string) (*VQCConfig, error)
	FindAll() ([]*VQCConfig, error)
	Delete(config *VQCConfig) error
}
