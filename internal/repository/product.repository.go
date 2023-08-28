package repository

import (
	"context"

	"github.com/SaulBenitez/inventory/internal/entity"
)

const (
	qryInsertProduct = `
		INSERT INTO PRODUCTS (name, description, price, created_by)
		VALUES (?, ?, ?, ?)
	`

	qryGetProducts = `
		SELECT 
			id,
			name,
			description,
			price,
			created_by
		FROM PRODUCTS
	`

	qryGetProductByID = `
		SELECT 
			id,
			name,
			description,
			price,
			created_by
		FROM PRODUCTS
		WHERE id = ?
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float64, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, description, price, createdBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}
	
	err := r.db.SelectContext(ctx, &pp, qryGetProducts)
	if err != nil {
		return nil, err
	}
	
	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}
	
	err := r.db.GetContext(ctx, p, qryGetProductByID, id)
	if err != nil {
		return nil, err
	}
	
	return p, nil
}
