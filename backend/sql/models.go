// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sql

import (
	"database/sql"
	"time"
)

type Product struct {
	ID             int32
	Brand          string
	ProductName    string
	ProductImgUrl  string
	ProductUrl     string
	CurrencyCode   sql.NullString
	RetailPrice    float64
	SalePrice      float64
	KorBrand       sql.NullString
	KorProductName sql.NullString
	ProductID      sql.NullString
	Gender         sql.NullString
	Color          sql.NullString
	Category       sql.NullString
	CategorySpec   sql.NullString
	StoreName      sql.NullString
	MadeIn         sql.NullString
	IsSale         bool
	SoldOut        bool
	UpdatedAt      time.Time
}

type Store struct {
	StoreName               string
	StoreUrl                string
	Country                 string
	Currency                string
	TaxReduction            float64
	IntlShippingFee         ShippingFee
	IntlFreeShippingMin     sql.NullString
	DomesticShippingFee     float64
	DomesticFreeShippingMin sql.NullString
	ShippingFeeCumulation   bool
	DeliveryAgency          string
	BrokerFee               bool
	Ddp                     bool
	UpdatedAt               time.Time
	StoreNameKor            sql.NullString
	TaxReductionManually    sql.NullBool
}
