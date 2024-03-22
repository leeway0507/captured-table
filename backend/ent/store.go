// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/schema"
	"backend/ent/store"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Store is the model entity for the Store schema.
type Store struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// KorID holds the value of the "kor_id" field.
	KorID string `json:"kor_id,omitempty"`
	// StoreURL holds the value of the "store_url" field.
	StoreURL string `json:"store_url,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// TaxReduction holds the value of the "tax_reduction" field.
	TaxReduction float64 `json:"tax_reduction,omitempty"`
	// TaxReductionManually holds the value of the "tax_reduction_manually" field.
	TaxReductionManually bool `json:"tax_reduction_manually,omitempty"`
	// IntlShippingFee holds the value of the "intl_shipping_fee" field.
	IntlShippingFee *schema.ShippingFee `json:"intl_shipping_fee,omitempty"`
	// IntlFreeShippingMin holds the value of the "intl_free_shipping_min" field.
	IntlFreeShippingMin int `json:"intl_free_shipping_min,omitempty"`
	// DomesticShippingFee holds the value of the "domestic_shipping_fee" field.
	DomesticShippingFee float64 `json:"domestic_shipping_fee,omitempty"`
	// DomesticFreeShippingMin holds the value of the "domestic_free_shipping_min" field.
	DomesticFreeShippingMin float64 `json:"domestic_free_shipping_min,omitempty"`
	// ShippingFeeCumulation holds the value of the "shipping_fee_cumulation" field.
	ShippingFeeCumulation bool `json:"shipping_fee_cumulation,omitempty"`
	// DeliveryAgency holds the value of the "delivery_agency" field.
	DeliveryAgency string `json:"delivery_agency,omitempty"`
	// BrokerFee holds the value of the "broker_fee" field.
	BrokerFee bool `json:"broker_fee,omitempty"`
	// Ddp holds the value of the "ddp" field.
	Ddp bool `json:"ddp,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StoreQuery when eager-loading is set.
	Edges        StoreEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StoreEdges holds the relations/edges for other nodes in the graph.
type StoreEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e StoreEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Store) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case store.FieldIntlShippingFee:
			values[i] = new([]byte)
		case store.FieldTaxReductionManually, store.FieldShippingFeeCumulation, store.FieldBrokerFee, store.FieldDdp:
			values[i] = new(sql.NullBool)
		case store.FieldTaxReduction, store.FieldDomesticShippingFee, store.FieldDomesticFreeShippingMin:
			values[i] = new(sql.NullFloat64)
		case store.FieldIntlFreeShippingMin:
			values[i] = new(sql.NullInt64)
		case store.FieldID, store.FieldKorID, store.FieldStoreURL, store.FieldCountry, store.FieldCurrency, store.FieldDeliveryAgency:
			values[i] = new(sql.NullString)
		case store.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Store fields.
func (s *Store) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case store.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case store.FieldKorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kor_id", values[i])
			} else if value.Valid {
				s.KorID = value.String
			}
		case store.FieldStoreURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_url", values[i])
			} else if value.Valid {
				s.StoreURL = value.String
			}
		case store.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				s.Country = value.String
			}
		case store.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				s.Currency = value.String
			}
		case store.FieldTaxReduction:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax_reduction", values[i])
			} else if value.Valid {
				s.TaxReduction = value.Float64
			}
		case store.FieldTaxReductionManually:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tax_reduction_manually", values[i])
			} else if value.Valid {
				s.TaxReductionManually = value.Bool
			}
		case store.FieldIntlShippingFee:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field intl_shipping_fee", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.IntlShippingFee); err != nil {
					return fmt.Errorf("unmarshal field intl_shipping_fee: %w", err)
				}
			}
		case store.FieldIntlFreeShippingMin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field intl_free_shipping_min", values[i])
			} else if value.Valid {
				s.IntlFreeShippingMin = int(value.Int64)
			}
		case store.FieldDomesticShippingFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field domestic_shipping_fee", values[i])
			} else if value.Valid {
				s.DomesticShippingFee = value.Float64
			}
		case store.FieldDomesticFreeShippingMin:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field domestic_free_shipping_min", values[i])
			} else if value.Valid {
				s.DomesticFreeShippingMin = value.Float64
			}
		case store.FieldShippingFeeCumulation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field shipping_fee_cumulation", values[i])
			} else if value.Valid {
				s.ShippingFeeCumulation = value.Bool
			}
		case store.FieldDeliveryAgency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_agency", values[i])
			} else if value.Valid {
				s.DeliveryAgency = value.String
			}
		case store.FieldBrokerFee:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field broker_fee", values[i])
			} else if value.Valid {
				s.BrokerFee = value.Bool
			}
		case store.FieldDdp:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ddp", values[i])
			} else if value.Valid {
				s.Ddp = value.Bool
			}
		case store.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Store.
// This includes values selected through modifiers, order, etc.
func (s *Store) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the Store entity.
func (s *Store) QueryProduct() *ProductQuery {
	return NewStoreClient(s.config).QueryProduct(s)
}

// Update returns a builder for updating this Store.
// Note that you need to call Store.Unwrap() before calling this method if this Store
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Store) Update() *StoreUpdateOne {
	return NewStoreClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Store entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Store) Unwrap() *Store {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Store is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Store) String() string {
	var builder strings.Builder
	builder.WriteString("Store(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("kor_id=")
	builder.WriteString(s.KorID)
	builder.WriteString(", ")
	builder.WriteString("store_url=")
	builder.WriteString(s.StoreURL)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(s.Country)
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(s.Currency)
	builder.WriteString(", ")
	builder.WriteString("tax_reduction=")
	builder.WriteString(fmt.Sprintf("%v", s.TaxReduction))
	builder.WriteString(", ")
	builder.WriteString("tax_reduction_manually=")
	builder.WriteString(fmt.Sprintf("%v", s.TaxReductionManually))
	builder.WriteString(", ")
	builder.WriteString("intl_shipping_fee=")
	builder.WriteString(fmt.Sprintf("%v", s.IntlShippingFee))
	builder.WriteString(", ")
	builder.WriteString("intl_free_shipping_min=")
	builder.WriteString(fmt.Sprintf("%v", s.IntlFreeShippingMin))
	builder.WriteString(", ")
	builder.WriteString("domestic_shipping_fee=")
	builder.WriteString(fmt.Sprintf("%v", s.DomesticShippingFee))
	builder.WriteString(", ")
	builder.WriteString("domestic_free_shipping_min=")
	builder.WriteString(fmt.Sprintf("%v", s.DomesticFreeShippingMin))
	builder.WriteString(", ")
	builder.WriteString("shipping_fee_cumulation=")
	builder.WriteString(fmt.Sprintf("%v", s.ShippingFeeCumulation))
	builder.WriteString(", ")
	builder.WriteString("delivery_agency=")
	builder.WriteString(s.DeliveryAgency)
	builder.WriteString(", ")
	builder.WriteString("broker_fee=")
	builder.WriteString(fmt.Sprintf("%v", s.BrokerFee))
	builder.WriteString(", ")
	builder.WriteString("ddp=")
	builder.WriteString(fmt.Sprintf("%v", s.Ddp))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Stores is a parsable slice of Store.
type Stores []*Store
