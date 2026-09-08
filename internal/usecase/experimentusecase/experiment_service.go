package experimentusecase

type ExperimentService struct {
	experimentRepository  ExperimentRepository
	vqcConfigRepository   VQCConfigRepository
	trainConfigRepository TrainConfigRepository
	userRepository        UserRepository
}

func NewExperimentService(
	experimentRepository ExperimentRepository,
	vqcConfigRepository VQCConfigRepository,
	trainConfigRepository TrainConfigRepository,
	userRepository UserRepository,
) *ExperimentService {
	return &ExperimentService{
		experimentRepository:  experimentRepository,
		vqcConfigRepository:   vqcConfigRepository,
		trainConfigRepository: trainConfigRepository,
		userRepository:        userRepository,
	}
}
