// Code generated by ent, DO NOT EDIT.

package store

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the store type in the database.
	Label = "store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "store_name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldTaxReduction holds the string denoting the tax_reduction field in the database.
	FieldTaxReduction = "tax_reduction"
	// FieldIntlShippingFee holds the string denoting the intl_shipping_fee field in the database.
	FieldIntlShippingFee = "intl_shipping_fee"
	// FieldIntlFreeShippingFee holds the string denoting the intl_free_shipping_fee field in the database.
	FieldIntlFreeShippingFee = "intl_free_shipping_fee"
	// FieldDomesticShippingFee holds the string denoting the domestic_shipping_fee field in the database.
	FieldDomesticShippingFee = "domestic_shipping_fee"
	// FieldDomesticFreeShippingFee holds the string denoting the domestic_free_shipping_fee field in the database.
	FieldDomesticFreeShippingFee = "domestic_free_shipping_fee"
	// FieldShippingFeeCumulation holds the string denoting the shipping_fee_cumulation field in the database.
	FieldShippingFeeCumulation = "shipping_fee_cumulation"
	// FieldDeliveryAgency holds the string denoting the delivery_agency field in the database.
	FieldDeliveryAgency = "delivery_agency"
	// FieldBrokerFee holds the string denoting the broker_fee field in the database.
	FieldBrokerFee = "broker_fee"
	// FieldDdp holds the string denoting the ddp field in the database.
	FieldDdp = "ddp"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// ProductFieldID holds the string denoting the ID field of the Product.
	ProductFieldID = "id"
	// Table holds the table name of the store in the database.
	Table = "stores"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "store_name"
)

// Columns holds all SQL columns for store fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldCountry,
	FieldCurrency,
	FieldTaxReduction,
	FieldIntlShippingFee,
	FieldIntlFreeShippingFee,
	FieldDomesticShippingFee,
	FieldDomesticFreeShippingFee,
	FieldShippingFeeCumulation,
	FieldDeliveryAgency,
	FieldBrokerFee,
	FieldDdp,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Store queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByTaxReduction orders the results by the tax_reduction field.
func ByTaxReduction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxReduction, opts...).ToFunc()
}

// ByIntlFreeShippingFee orders the results by the intl_free_shipping_fee field.
func ByIntlFreeShippingFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntlFreeShippingFee, opts...).ToFunc()
}

// ByDomesticShippingFee orders the results by the domestic_shipping_fee field.
func ByDomesticShippingFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDomesticShippingFee, opts...).ToFunc()
}

// ByDomesticFreeShippingFee orders the results by the domestic_free_shipping_fee field.
func ByDomesticFreeShippingFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDomesticFreeShippingFee, opts...).ToFunc()
}

// ByShippingFeeCumulation orders the results by the shipping_fee_cumulation field.
func ByShippingFeeCumulation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingFeeCumulation, opts...).ToFunc()
}

// ByDeliveryAgency orders the results by the delivery_agency field.
func ByDeliveryAgency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeliveryAgency, opts...).ToFunc()
}

// ByBrokerFee orders the results by the broker_fee field.
func ByBrokerFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrokerFee, opts...).ToFunc()
}

// ByDdp orders the results by the ddp field.
func ByDdp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDdp, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, ProductFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}
