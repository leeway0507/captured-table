package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Store struct {
	ent.Schema
}
type ShippingFee struct {
	Light float32
	Heavy float32
	Shoes float32
}

// Fields of the Product.
// Fields of the User.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.String("store_name"),
		field.String("url"),
		field.String("country"),
		field.String("currency"),
		field.Float("tax_reduction"),
		field.JSON("intl_shipping_fee", &ShippingFee{}),
		field.Int("intl_free_shipping_price"),
		field.Float("domestic_shipping_fee"),
		field.Float("domestic_free_shipping_price"),
		field.String("delivery_agency"),
		field.Bool("broker_fee"),
		field.Bool("ddp"),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Product.
// Edges of the User.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}