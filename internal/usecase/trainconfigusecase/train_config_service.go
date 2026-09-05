package trainconfigusecase

type TrainConfigService struct {
	trainConfigRepository TrainConfigRepository
	userRepository        UserRepository
}

func NewTrainConfigService(trainConfigRepository TrainConfigRepository, userRepository UserRepository) *TrainConfigService {
	return &TrainConfigService{
		trainConfigRepository: trainConfigRepository,
		userRepository:        userRepository,
	}
}
