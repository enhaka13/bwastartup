package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
)


type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository Repository
	campaignRepository campaign.Repository
	paymentService payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	} 

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not authorized as campaign's owner")
	}
	
	transactions, err := s.repository.GetByCampaignID(input.ID)
	 if err != nil {
		return []Transaction{}, err
	 }

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignID: input.CampaignID,
		Amount: input.Amount,
		UserID: input.User.ID,
		Status: "pending",

	}
	
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return Transaction{}, err
	}

	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return Transaction{}, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return Transaction{}, err
	}

	return newTransaction, nil
}