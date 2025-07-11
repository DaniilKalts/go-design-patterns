package person

import "gorm.io/gorm"

// Repository : Target
type Repository interface {
	Create(*Person) error
	GetByID(string) (*Person, error)
	GetAll() ([]*Person, error)
	Update(*Person) error
	Delete(string) error
}

// repo : Adapter
type repo struct {
	db *gorm.DB // db : Adaptee
}

func NewPersonRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(p *Person) error {
	return r.db.Create(p).Error
}

func (r *repo) GetByID(id string) (*Person, error) {
	var p Person

	return &p, r.db.First(&p, "id = ?", id).Error
}

func (r *repo) GetAll() ([]*Person, error) {
	var persons []*Person

	return persons, r.db.Find(&persons).Error
}

func (r *repo) Update(p *Person) error {
	return r.db.Save(p).Error
}

func (r *repo) Delete(id string) error {
	return r.db.Delete(&Person{}, "id = ?", id).Error
}
