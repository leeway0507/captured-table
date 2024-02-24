// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/product"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int `json:"store_id,omitempty"`
	// Brand holds the value of the "brand" field.
	Brand string `json:"brand,omitempty"`
	// ProductName holds the value of the "product_name" field.
	ProductName string `json:"product_name,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// KorBrand holds the value of the "kor_brand" field.
	KorBrand string `json:"kor_brand,omitempty"`
	// KorProductName holds the value of the "kor_product_name" field.
	KorProductName string `json:"kor_product_name,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID string `json:"product_id,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender product.Gender `json:"gender,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// CategorySpec holds the value of the "category_spec" field.
	CategorySpec string `json:"category_spec,omitempty"`
	// SoldOut holds the value of the "sold_out" field.
	SoldOut bool `json:"sold_out,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt                time.Time `json:"updated_at,omitempty"`
	delivery_agency_products *int
	store_products           *int
	selectValues             sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldSoldOut:
			values[i] = new(sql.NullBool)
		case product.FieldID, product.FieldStoreID, product.FieldPrice:
			values[i] = new(sql.NullInt64)
		case product.FieldBrand, product.FieldProductName, product.FieldKorBrand, product.FieldKorProductName, product.FieldProductID, product.FieldGender, product.FieldColor, product.FieldCategory, product.FieldCategorySpec:
			values[i] = new(sql.NullString)
		case product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case product.ForeignKeys[0]: // delivery_agency_products
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[1]: // store_products
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldStoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				pr.StoreID = int(value.Int64)
			}
		case product.FieldBrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field brand", values[i])
			} else if value.Valid {
				pr.Brand = value.String
			}
		case product.FieldProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_name", values[i])
			} else if value.Valid {
				pr.ProductName = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = int(value.Int64)
			}
		case product.FieldKorBrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kor_brand", values[i])
			} else if value.Valid {
				pr.KorBrand = value.String
			}
		case product.FieldKorProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kor_product_name", values[i])
			} else if value.Valid {
				pr.KorProductName = value.String
			}
		case product.FieldProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				pr.ProductID = value.String
			}
		case product.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pr.Gender = product.Gender(value.String)
			}
		case product.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				pr.Color = value.String
			}
		case product.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				pr.Category = value.String
			}
		case product.FieldCategorySpec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category_spec", values[i])
			} else if value.Valid {
				pr.CategorySpec = value.String
			}
		case product.FieldSoldOut:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sold_out", values[i])
			} else if value.Valid {
				pr.SoldOut = value.Bool
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field delivery_agency_products", value)
			} else if value.Valid {
				pr.delivery_agency_products = new(int)
				*pr.delivery_agency_products = int(value.Int64)
			}
		case product.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field store_products", value)
			} else if value.Valid {
				pr.store_products = new(int)
				*pr.store_products = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.StoreID))
	builder.WriteString(", ")
	builder.WriteString("brand=")
	builder.WriteString(pr.Brand)
	builder.WriteString(", ")
	builder.WriteString("product_name=")
	builder.WriteString(pr.ProductName)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("kor_brand=")
	builder.WriteString(pr.KorBrand)
	builder.WriteString(", ")
	builder.WriteString("kor_product_name=")
	builder.WriteString(pr.KorProductName)
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(pr.ProductID)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", pr.Gender))
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(pr.Color)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(pr.Category)
	builder.WriteString(", ")
	builder.WriteString("category_spec=")
	builder.WriteString(pr.CategorySpec)
	builder.WriteString(", ")
	builder.WriteString("sold_out=")
	builder.WriteString(fmt.Sprintf("%v", pr.SoldOut))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product
