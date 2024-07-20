package repository

import (
	"Project_Niko/internal/domain"
	"errors"
)

type LoanRepositoryInterface interface {
	LoanSaver
	LoanUpdater
	LoanDeleter
}

type LoanSaver interface {
	SaveLoan(loan *domain.Loan) error 
}

type LoanUpdater interface {
	UpdateLoan(loan *domain.Loan) error
}

type LoanDeleter interface {
	DeleteLoan(loadID int) error
}

type LoanRepository struct {
	loans map[int]domain.Loan
}

func NewLoanRepository() LoanRepositoryInterface {
	return &LoanRepository{
		loans: map[int]domain.Loan{},
	}
}

func (repo *LoanRepository) SaveLoan(loan *domain.Loan) error {
	if _, exists := repo.loans[loan.ID]; exists {
		return errors.New("loan already exist")
	}
	
	repo.loans[loan.ID] = *loan
	return nil
}

func (repo *LoanRepository) UpdateLoan(loan *domain.Loan) error {
	if _, exists := repo.loans[loan.ID]; exists{
		repo.loans[loan.ID] = *loan
		return nil
	}

	return errors.New("id loan not found")
}

func (repo *LoanRepository) DeleteLoan(loanID int) error {
	if _,exists := repo.loans[loanID]; !exists {
		return errors.New("id loan not found")
	}

	delete(repo.loans, loanID)
	return nil
}