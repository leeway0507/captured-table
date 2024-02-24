// Code generated by ent, DO NOT EDIT.

package product

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStoreID holds the string denoting the store_id field in the database.
	FieldStoreID = "store_id"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldProductName holds the string denoting the product_name field in the database.
	FieldProductName = "product_name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldKorBrand holds the string denoting the kor_brand field in the database.
	FieldKorBrand = "kor_brand"
	// FieldKorProductName holds the string denoting the kor_product_name field in the database.
	FieldKorProductName = "kor_product_name"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCategorySpec holds the string denoting the category_spec field in the database.
	FieldCategorySpec = "category_spec"
	// FieldSoldOut holds the string denoting the sold_out field in the database.
	FieldSoldOut = "sold_out"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the product in the database.
	Table = "products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldStoreID,
	FieldBrand,
	FieldProductName,
	FieldPrice,
	FieldKorBrand,
	FieldKorProductName,
	FieldProductID,
	FieldGender,
	FieldColor,
	FieldCategory,
	FieldCategorySpec,
	FieldSoldOut,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"delivery_agency_products",
	"store_products",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSoldOut holds the default value on creation for the "sold_out" field.
	DefaultSoldOut bool
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderW Gender = "w"
	GenderM Gender = "m"
	GenderB Gender = "b"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderW, GenderM, GenderB:
		return nil
	default:
		return fmt.Errorf("product: invalid enum value for gender field: %q", ge)
	}
}

// OrderOption defines the ordering options for the Product queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStoreID orders the results by the store_id field.
func ByStoreID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreID, opts...).ToFunc()
}

// ByBrand orders the results by the brand field.
func ByBrand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrand, opts...).ToFunc()
}

// ByProductName orders the results by the product_name field.
func ByProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByKorBrand orders the results by the kor_brand field.
func ByKorBrand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKorBrand, opts...).ToFunc()
}

// ByKorProductName orders the results by the kor_product_name field.
func ByKorProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKorProductName, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByCategorySpec orders the results by the category_spec field.
func ByCategorySpec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategorySpec, opts...).ToFunc()
}

// BySoldOut orders the results by the sold_out field.
func BySoldOut(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSoldOut, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
