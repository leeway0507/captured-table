// Code generated by ent, DO NOT EDIT.

package deliveryagency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the deliveryagency type in the database.
	Label = "delivery_agency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldVATReductionRate holds the string denoting the vat_reduction_rate field in the database.
	FieldVATReductionRate = "vat_reduction_rate"
	// FieldShippingFee holds the string denoting the shipping_fee field in the database.
	FieldShippingFee = "shipping_fee"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the deliveryagency in the database.
	Table = "delivery_agencies"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "delivery_agency_products"
)

// Columns holds all SQL columns for deliveryagency fields.
var Columns = []string{
	FieldID,
	FieldCountry,
	FieldVATReductionRate,
	FieldShippingFee,
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

// OrderOption defines the ordering options for the DeliveryAgency queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByVATReductionRate orders the results by the VAT_reduction_rate field.
func ByVATReductionRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVATReductionRate, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
	)
}