package usecases

import "flavioltonon/hmv/domain/entity"

type EmergencyUsecase interface {
	CreateEmergency(pacientID string) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error)
}
