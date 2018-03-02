package benchmark

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repository interface {
	GetRandomId() (payment Payment)
	UpdateRow(payment Payment)
}

type memoryRepository struct {
	Conn *gorm.DB
}

func (s *memoryRepository) GetRandomId() (payment Payment) {
	s.Conn.
		Debug().
		Table("payments").
		Select("id").
		Where("id BETWEEN 10000 AND 1000000").
		Order("random()").
		First(&payment)

	return payment
}

func (s *memoryRepository) UpdateRow(payment Payment) {
	payment.CreatedAt = time.Now()
	payment.SynchronizedAt = time.Now()
	s.Conn.
		Debug().
		Save(&payment)
}

func NewRepository(Conn *gorm.DB) Repository {
	return &memoryRepository{Conn}
}
