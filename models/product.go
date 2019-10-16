package models

import "time"

type Product struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Sku         string    `db:"sku" json:"sku"`
	Price       string    `db:"price" json:"price"`
	Discount    string    `db:"discount" json:"discount"`
	RewardPoint int       `db:"reward_point" json:"reward_point"`
	RedeemPoint int       `db:"redeem_point" json:"redeem_point"`
	RedeemStamp int       `db:"redeem_stamp" json:"redeem_stamp"`
	Status      string    `db:"status" json:"status"`
	MerchantID  int       `db:"merchant_id" json:"merchant_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type ProductInput struct {
	Name        string    `json:"name" validate:"required"`
	Sku         string    `json:"sku" validate:"required"`
	Price       string    `json:"price" validate:"required"`
	Discount    string    `json:"discount" validate:"required"`
	RewardPoint int       `json:"reward_point" validate:"required"`
	RedeemPoint int       `json:"redeem_point" validate:"required"`
	RedeemStamp int       `json:"redeem_stamp" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	MerchantID  int       `json:"merchant_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at""`
}
