package user

import "github.com/AndroDeMohawk/link-cutter/pkg/db"

type Repository struct {
	Database *db.Db
}

type RepositoryDeps struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (r *Repository) Create(user *User) (*User, error) {
	result := r.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
