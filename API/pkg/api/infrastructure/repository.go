package infrastructure

import (
	"hospitalApi/pkg/repository"

	"gorm.io/gorm"
)

type Repository struct {
	IPatientsRepository repository.IPatientsRepository
	IStaffsRepository   repository.IStaffsRepository
}

func CreateRepository(db *gorm.DB, helper Helper) (result Repository) {
	return Repository{
		IPatientsRepository: repository.MakeIPatientsRepository(db),
		IStaffsRepository:   repository.MakeIStaffsRepository(db, helper.ICommon),
	}
}
