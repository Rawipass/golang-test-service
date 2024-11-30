package usecase

import (
	"github.com/Rawipass/golang-test-service/internal/product/repository"
	"github.com/Rawipass/golang-test-service/models"
	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: *repo}
}

// Product
func (u *ProductUseCase) CreateProduct(name string, quantity int, price float64) (*models.Product, error) {
	return u.repo.CreateProduct(name, quantity, price)
}

func (u *ProductUseCase) ListProducts() ([]models.Product, error) {
	return u.repo.GetAllProducts()
}

func (u *ProductUseCase) GetProductDetail(id string) (*models.Product, error) {
	return u.repo.GetProductByID(id)
}

// Commission
func (u *ProductUseCase) GetCommissionDetail(id string) (*models.Commission, error) {
	return u.repo.GetCommissionByID(id)
}

func (u *ProductUseCase) ListCommissions() ([]models.Commission, error) {
	return u.repo.GetAllCommissions()
}

// Affiliate
func (u *ProductUseCase) CreateAffiliate(name string, masterID int) (*models.Affiliate, error) {
	return u.repo.CreateAffiliate(name, masterID)
}

func (u *ProductUseCase) ListAffiliates() ([]models.Affiliate, error) {
	return u.repo.GetAllAffiliates()
}

func (u *ProductUseCase) GetAffiliateDetail(id string) (*models.Affiliate, error) {
	return u.repo.GetAffiliateByID(id)
}

// sale
func (u *ProductUseCase) CreateSale(affiliateID uuid.UUID, productID uuid.UUID, productPrice float64) ([]models.Commission, error) {

	orderID := uuid.New()
	totalAmount := productPrice

	order := models.Order{
		ID:          orderID,
		AffiliateID: affiliateID,
		ProductID:   productID,
		TotalAmount: totalAmount,
	}

	err := u.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	affiliates, err := u.repo.GetAllAffiliates()
	if err != nil {
		return nil, err
	}

	commissions, err := u.calculateCommissions(affiliateID, affiliates, productPrice)
	if err != nil {
		return nil, err
	}

	for _, commission := range commissions {
		err := u.repo.CreateCommission(commission)
		if err != nil {
			return nil, err
		}
	}

	return commissions, nil
}

func (u *ProductUseCase) calculateCommissions(affiliateID uuid.UUID, affiliates []models.Affiliate, productPrice float64) ([]models.Commission, error) {
	var commissions []models.Commission
	currentAffiliate := affiliateID

	for level := 0; currentAffiliate != uuid.Nil && level < len(affiliates); level++ {
		affiliate := findAffiliateByID(affiliates, currentAffiliate)
		if affiliate == nil {
			break
		}

		commissionPercentage := 0.0
		switch level {
		case 0: // L4
			commissionPercentage = 0.05 // 5%
		case 1: // L3
			commissionPercentage = 0.10 // 10%
		case 2: // L2
			commissionPercentage = 0.15 // 15%
		case 3: // L1
			commissionPercentage = 0.20 // 20%
		}

		commissionAmount := commissionPercentage * productPrice

		commission := models.Commission{
			ID:          uuid.New(),
			OrderID:     uuid.New(),
			AffiliateID: affiliate.ID,
			Amount:      commissionAmount,
		}

		commissions = append(commissions, commission)

		currentAffiliate = affiliate.MasterAffiliate
	}

	return commissions, nil
}

func findAffiliateByID(affiliates []models.Affiliate, affiliateID uuid.UUID) *models.Affiliate {
	for _, affiliate := range affiliates {
		if affiliate.ID == affiliateID {
			return &affiliate
		}
	}
	return nil
}
