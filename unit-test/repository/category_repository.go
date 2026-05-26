package repository

import "unit-test/entity"

// Agar kode kita mudah dites, buat kontrak agar berupa interface
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
