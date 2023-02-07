package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&transactions).Error; err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}