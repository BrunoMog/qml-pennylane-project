package vqc_config_usecase

type VQCConfigService struct {
	vqcConfigRepository VQCConfigRepository
	userRepository      UserRepository
}

func NewVQCConfigService(vqcConfigRepository VQCConfigRepository, userRepository UserRepository) *VQCConfigService {
	return &VQCConfigService{
		vqcConfigRepository: vqcConfigRepository,
		userRepository:      userRepository,
	}
}
