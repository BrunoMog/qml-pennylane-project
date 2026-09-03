package vqcconfigusecase

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
