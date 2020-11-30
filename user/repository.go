package user

import (
	"gorm.io/gorm"
)

// Repository for inserting data to database
type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository function received main
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
