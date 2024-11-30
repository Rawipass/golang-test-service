package repository

import (
	"context"
	"github.com/Rawipass/golang-test-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// Product
func (r *ProductRepository) CreateProduct(name string, quantity int, price float64) (*models.Product, error) {
	var product models.Product

	query := `INSERT INTO products (name, quantity, price) VALUES ($1, $2, $3) RETURNING id, name, quantity, price`
	err := r.db.QueryRow(context.Background(), query, name, quantity, price).Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	query := `SELECT id, name, quantity, price FROM products`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, name, quantity, price FROM products WHERE id = $1`
	err := r.db.QueryRow(context.Background(), query, id).Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Commission
func (r *ProductRepository) GetCommissionByID(id string) (*models.Commission, error) {
	var commission models.Commission
	query := `SELECT id, order_id, affiliate_id, amount FROM commissions WHERE id = $1`
	err := r.db.QueryRow(context.Background(), query, id).Scan(&commission.ID, &commission.OrderID, &commission.AffiliateID, &commission.Amount)
	if err != nil {
		return nil, err
	}
	return &commission, nil
}

func (r *ProductRepository) GetAllCommissions() ([]models.Commission, error) {
	var commissions []models.Commission

	query := `SELECT id, order_id, affiliate_id, amount FROM commissions`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commission models.Commission
		if err := rows.Scan(&commission.ID, &commission.OrderID, &commission.AffiliateID, &commission.Amount); err != nil {
			return nil, err
		}
		commissions = append(commissions, commission)
	}

	return commissions, nil
}

func (r ProductRepository) CreateCommission(commission models.Commission) error {
    query := `INSERT INTO commissions (id, order_id, affiliate_id, amount) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(context.Background(), query, commission.ID, commission.OrderID, commission.AffiliateID, commission.Amount)
    return err
}

// Affiliate
func (r *ProductRepository) CreateAffiliate(name string, masterID int) (*models.Affiliate, error) {
	var affiliate models.Affiliate

	query := `INSERT INTO affiliates (name, master_affiliate) VALUES ($1, $2) RETURNING id, name, master_affiliate`
	err := r.db.QueryRow(context.Background(), query, name, masterID).Scan(&affiliate.ID, &affiliate.Name, &affiliate.MasterAffiliate)
	if err != nil {
		return nil, err
	}

	return &affiliate, nil
}

func (r *ProductRepository) GetAllAffiliates() ([]models.Affiliate, error) {
	var affiliates []models.Affiliate

	query := `SELECT id, name, master_affiliate FROM affiliates`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var affiliate models.Affiliate
		if err := rows.Scan(&affiliate.ID, &affiliate.Name, &affiliate.MasterAffiliate); err != nil {
			return nil, err
		}
		affiliates = append(affiliates, affiliate)
	}

	return affiliates, nil
}

func (r *ProductRepository) GetAffiliateByID(id string) (*models.Affiliate, error) {
	query := `
        SELECT id, name, master_affiliate, balance
        FROM affiliates
        WHERE id = $1
    `

	var affiliate models.Affiliate
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&affiliate.ID,
		&affiliate.Name,
		&affiliate.MasterAffiliate,
		&affiliate.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &affiliate, nil
}

//order 
func (r *ProductRepository) CreateOrder(order models.Order) error {
    query := `INSERT INTO orders (id, affiliate_id, product_id, total_amount) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(context.Background(), query, order.ID, order.AffiliateID, order.ProductID, order.TotalAmount)
    return err
}
