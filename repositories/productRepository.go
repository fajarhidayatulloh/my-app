package repositories

import (
	"database/sql"
	"time"

	"github.com/my-app/infrastructures"
	"github.com/my-app/models"
	log "github.com/sirupsen/logrus"
)

// IProductRepository init
type IProductRepository interface {
	StoreProduct(data models.Product) (models.Product, error)
	ProductList() ([]models.Product, error)
	GetProductByID(int) (models.Product, error)
}

// ProductRepository behaviour
type ProductRepository struct {
	DB infrastructures.ISQLConnection
}

// StoreProduct is
func (r *ProductRepository) StoreProduct(data models.Product) (models.Product, error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO product 
							(product.name, product.sku, product.price, product.discount, 
							product.reward_point, product.redeem_point, product.redeem_stamp, 
							product.status, product.merchant_id, product.created_at) 
							VALUES (?,?,?,?,?,?,?,?,?,?)`,
	)

	if err != nil {
		return data, err
	}

	res, err := stmt.Exec(data.Name, data.Sku, data.Price, data.Discount,
		data.RewardPoint, data.RedeemPoint, data.RedeemStamp,
		data.Status, data.MerchantID, data.CreatedAt.Format(time.RFC3339),
	)

	if err != nil {
		return data, err
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.WithFields(log.Fields{
			"event": "StoreProduct",
		}).Error(err)
	}

	return data, err
}

// ProductList is
func (r *ProductRepository) ProductList() (products []models.Product, err error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	rows, err := db.Query(`SELECT id,name, sku, price, discount, reward_point, 
							redeem_point, redeem_stamp, status, merchant_id, created_at
							FROM product`)

	if err == sql.ErrNoRows {
		err = nil
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Sku,
			&product.Price,
			&product.Discount,
			&product.RewardPoint,
			&product.RedeemPoint,
			&product.RedeemStamp,
			&product.Status,
			&product.MerchantID,
			&product.CreatedAt,
		); err != nil {
			log.WithFields(log.Fields{
				"event": "get_product",
			}).Error(err)
		}

		products = append(products, product)
	}
	return
}

// GetProductByID is
func (r *ProductRepository) GetProductByID(ID int) (product models.Product, err error) {
	db := r.DB.GetPlayerWriteDb()
	defer db.Close()

	rows := db.QueryRow(`SELECT id,name, sku, price, discount, reward_point, 
							redeem_point, redeem_stamp, status, merchant_id, created_at
							FROM product where id = ?`, ID)

	err = rows.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Sku,
		&product.Discount,
		&product.RewardPoint,
		&product.RedeemPoint,
		&product.RedeemStamp,
		&product.Status,
		&product.MerchantID,
		&product.CreatedAt,
	)
	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		log.WithFields(log.Fields{
			"event": "get_user",
			"id":    ID,
		}).Error(err)
	}

	return product, err
}
