package link

import (
	"github.com/AndroDeMohawk/link-cutter/pkg/db"
	"gorm.io/gorm/clause"
)

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

func (r *Repository) Create(link *Link) (*Link, error) {
	result := r.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *Repository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := r.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *Repository) Update(link *Link) (*Link, error) {
	result := r.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) FindById(id uint) error {
	result := r.Database.DB.First(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
