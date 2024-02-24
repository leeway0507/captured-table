// Code generated by ent, DO NOT EDIT.

package product

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// StoreID applies equality check predicate on the "store_id" field. It's identical to StoreIDEQ.
func StoreID(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStoreID, v))
}

// Brand applies equality check predicate on the "brand" field. It's identical to BrandEQ.
func Brand(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldBrand, v))
}

// ProductName applies equality check predicate on the "product_name" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// KorBrand applies equality check predicate on the "kor_brand" field. It's identical to KorBrandEQ.
func KorBrand(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldKorBrand, v))
}

// KorProductName applies equality check predicate on the "kor_product_name" field. It's identical to KorProductNameEQ.
func KorProductName(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldKorProductName, v))
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductID, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldColor, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategory, v))
}

// CategorySpec applies equality check predicate on the "category_spec" field. It's identical to CategorySpecEQ.
func CategorySpec(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategorySpec, v))
}

// SoldOut applies equality check predicate on the "sold_out" field. It's identical to SoldOutEQ.
func SoldOut(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSoldOut, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// StoreIDEQ applies the EQ predicate on the "store_id" field.
func StoreIDEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStoreID, v))
}

// StoreIDNEQ applies the NEQ predicate on the "store_id" field.
func StoreIDNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStoreID, v))
}

// StoreIDIn applies the In predicate on the "store_id" field.
func StoreIDIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldStoreID, vs...))
}

// StoreIDNotIn applies the NotIn predicate on the "store_id" field.
func StoreIDNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldStoreID, vs...))
}

// StoreIDGT applies the GT predicate on the "store_id" field.
func StoreIDGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldStoreID, v))
}

// StoreIDGTE applies the GTE predicate on the "store_id" field.
func StoreIDGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldStoreID, v))
}

// StoreIDLT applies the LT predicate on the "store_id" field.
func StoreIDLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldStoreID, v))
}

// StoreIDLTE applies the LTE predicate on the "store_id" field.
func StoreIDLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldStoreID, v))
}

// BrandEQ applies the EQ predicate on the "brand" field.
func BrandEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldBrand, v))
}

// BrandNEQ applies the NEQ predicate on the "brand" field.
func BrandNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldBrand, v))
}

// BrandIn applies the In predicate on the "brand" field.
func BrandIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldBrand, vs...))
}

// BrandNotIn applies the NotIn predicate on the "brand" field.
func BrandNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldBrand, vs...))
}

// BrandGT applies the GT predicate on the "brand" field.
func BrandGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldBrand, v))
}

// BrandGTE applies the GTE predicate on the "brand" field.
func BrandGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldBrand, v))
}

// BrandLT applies the LT predicate on the "brand" field.
func BrandLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldBrand, v))
}

// BrandLTE applies the LTE predicate on the "brand" field.
func BrandLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldBrand, v))
}

// BrandContains applies the Contains predicate on the "brand" field.
func BrandContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldBrand, v))
}

// BrandHasPrefix applies the HasPrefix predicate on the "brand" field.
func BrandHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldBrand, v))
}

// BrandHasSuffix applies the HasSuffix predicate on the "brand" field.
func BrandHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldBrand, v))
}

// BrandEqualFold applies the EqualFold predicate on the "brand" field.
func BrandEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldBrand, v))
}

// BrandContainsFold applies the ContainsFold predicate on the "brand" field.
func BrandContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldBrand, v))
}

// ProductNameEQ applies the EQ predicate on the "product_name" field.
func ProductNameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// ProductNameNEQ applies the NEQ predicate on the "product_name" field.
func ProductNameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldProductName, v))
}

// ProductNameIn applies the In predicate on the "product_name" field.
func ProductNameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldProductName, vs...))
}

// ProductNameNotIn applies the NotIn predicate on the "product_name" field.
func ProductNameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldProductName, vs...))
}

// ProductNameGT applies the GT predicate on the "product_name" field.
func ProductNameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldProductName, v))
}

// ProductNameGTE applies the GTE predicate on the "product_name" field.
func ProductNameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldProductName, v))
}

// ProductNameLT applies the LT predicate on the "product_name" field.
func ProductNameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldProductName, v))
}

// ProductNameLTE applies the LTE predicate on the "product_name" field.
func ProductNameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldProductName, v))
}

// ProductNameContains applies the Contains predicate on the "product_name" field.
func ProductNameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldProductName, v))
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "product_name" field.
func ProductNameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldProductName, v))
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "product_name" field.
func ProductNameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldProductName, v))
}

// ProductNameEqualFold applies the EqualFold predicate on the "product_name" field.
func ProductNameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldProductName, v))
}

// ProductNameContainsFold applies the ContainsFold predicate on the "product_name" field.
func ProductNameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldProductName, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// KorBrandEQ applies the EQ predicate on the "kor_brand" field.
func KorBrandEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldKorBrand, v))
}

// KorBrandNEQ applies the NEQ predicate on the "kor_brand" field.
func KorBrandNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldKorBrand, v))
}

// KorBrandIn applies the In predicate on the "kor_brand" field.
func KorBrandIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldKorBrand, vs...))
}

// KorBrandNotIn applies the NotIn predicate on the "kor_brand" field.
func KorBrandNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldKorBrand, vs...))
}

// KorBrandGT applies the GT predicate on the "kor_brand" field.
func KorBrandGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldKorBrand, v))
}

// KorBrandGTE applies the GTE predicate on the "kor_brand" field.
func KorBrandGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldKorBrand, v))
}

// KorBrandLT applies the LT predicate on the "kor_brand" field.
func KorBrandLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldKorBrand, v))
}

// KorBrandLTE applies the LTE predicate on the "kor_brand" field.
func KorBrandLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldKorBrand, v))
}

// KorBrandContains applies the Contains predicate on the "kor_brand" field.
func KorBrandContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldKorBrand, v))
}

// KorBrandHasPrefix applies the HasPrefix predicate on the "kor_brand" field.
func KorBrandHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldKorBrand, v))
}

// KorBrandHasSuffix applies the HasSuffix predicate on the "kor_brand" field.
func KorBrandHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldKorBrand, v))
}

// KorBrandIsNil applies the IsNil predicate on the "kor_brand" field.
func KorBrandIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldKorBrand))
}

// KorBrandNotNil applies the NotNil predicate on the "kor_brand" field.
func KorBrandNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldKorBrand))
}

// KorBrandEqualFold applies the EqualFold predicate on the "kor_brand" field.
func KorBrandEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldKorBrand, v))
}

// KorBrandContainsFold applies the ContainsFold predicate on the "kor_brand" field.
func KorBrandContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldKorBrand, v))
}

// KorProductNameEQ applies the EQ predicate on the "kor_product_name" field.
func KorProductNameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldKorProductName, v))
}

// KorProductNameNEQ applies the NEQ predicate on the "kor_product_name" field.
func KorProductNameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldKorProductName, v))
}

// KorProductNameIn applies the In predicate on the "kor_product_name" field.
func KorProductNameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldKorProductName, vs...))
}

// KorProductNameNotIn applies the NotIn predicate on the "kor_product_name" field.
func KorProductNameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldKorProductName, vs...))
}

// KorProductNameGT applies the GT predicate on the "kor_product_name" field.
func KorProductNameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldKorProductName, v))
}

// KorProductNameGTE applies the GTE predicate on the "kor_product_name" field.
func KorProductNameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldKorProductName, v))
}

// KorProductNameLT applies the LT predicate on the "kor_product_name" field.
func KorProductNameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldKorProductName, v))
}

// KorProductNameLTE applies the LTE predicate on the "kor_product_name" field.
func KorProductNameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldKorProductName, v))
}

// KorProductNameContains applies the Contains predicate on the "kor_product_name" field.
func KorProductNameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldKorProductName, v))
}

// KorProductNameHasPrefix applies the HasPrefix predicate on the "kor_product_name" field.
func KorProductNameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldKorProductName, v))
}

// KorProductNameHasSuffix applies the HasSuffix predicate on the "kor_product_name" field.
func KorProductNameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldKorProductName, v))
}

// KorProductNameIsNil applies the IsNil predicate on the "kor_product_name" field.
func KorProductNameIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldKorProductName))
}

// KorProductNameNotNil applies the NotNil predicate on the "kor_product_name" field.
func KorProductNameNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldKorProductName))
}

// KorProductNameEqualFold applies the EqualFold predicate on the "kor_product_name" field.
func KorProductNameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldKorProductName, v))
}

// KorProductNameContainsFold applies the ContainsFold predicate on the "kor_product_name" field.
func KorProductNameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldKorProductName, v))
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldProductID, vs...))
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldProductID, v))
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldProductID, v))
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldProductID, v))
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldProductID, v))
}

// ProductIDContains applies the Contains predicate on the "product_id" field.
func ProductIDContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldProductID, v))
}

// ProductIDHasPrefix applies the HasPrefix predicate on the "product_id" field.
func ProductIDHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldProductID, v))
}

// ProductIDHasSuffix applies the HasSuffix predicate on the "product_id" field.
func ProductIDHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldProductID, v))
}

// ProductIDIsNil applies the IsNil predicate on the "product_id" field.
func ProductIDIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldProductID))
}

// ProductIDNotNil applies the NotNil predicate on the "product_id" field.
func ProductIDNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldProductID))
}

// ProductIDEqualFold applies the EqualFold predicate on the "product_id" field.
func ProductIDEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldProductID, v))
}

// ProductIDContainsFold applies the ContainsFold predicate on the "product_id" field.
func ProductIDContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldProductID, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldGender, vs...))
}

// GenderIsNil applies the IsNil predicate on the "gender" field.
func GenderIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldGender))
}

// GenderNotNil applies the NotNil predicate on the "gender" field.
func GenderNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldGender))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldColor, v))
}

// ColorIsNil applies the IsNil predicate on the "color" field.
func ColorIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldColor))
}

// ColorNotNil applies the NotNil predicate on the "color" field.
func ColorNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldColor))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldColor, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldCategory))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldCategory, v))
}

// CategorySpecEQ applies the EQ predicate on the "category_spec" field.
func CategorySpecEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategorySpec, v))
}

// CategorySpecNEQ applies the NEQ predicate on the "category_spec" field.
func CategorySpecNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCategorySpec, v))
}

// CategorySpecIn applies the In predicate on the "category_spec" field.
func CategorySpecIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCategorySpec, vs...))
}

// CategorySpecNotIn applies the NotIn predicate on the "category_spec" field.
func CategorySpecNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCategorySpec, vs...))
}

// CategorySpecGT applies the GT predicate on the "category_spec" field.
func CategorySpecGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCategorySpec, v))
}

// CategorySpecGTE applies the GTE predicate on the "category_spec" field.
func CategorySpecGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCategorySpec, v))
}

// CategorySpecLT applies the LT predicate on the "category_spec" field.
func CategorySpecLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCategorySpec, v))
}

// CategorySpecLTE applies the LTE predicate on the "category_spec" field.
func CategorySpecLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCategorySpec, v))
}

// CategorySpecContains applies the Contains predicate on the "category_spec" field.
func CategorySpecContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldCategorySpec, v))
}

// CategorySpecHasPrefix applies the HasPrefix predicate on the "category_spec" field.
func CategorySpecHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldCategorySpec, v))
}

// CategorySpecHasSuffix applies the HasSuffix predicate on the "category_spec" field.
func CategorySpecHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldCategorySpec, v))
}

// CategorySpecIsNil applies the IsNil predicate on the "category_spec" field.
func CategorySpecIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldCategorySpec))
}

// CategorySpecNotNil applies the NotNil predicate on the "category_spec" field.
func CategorySpecNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldCategorySpec))
}

// CategorySpecEqualFold applies the EqualFold predicate on the "category_spec" field.
func CategorySpecEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldCategorySpec, v))
}

// CategorySpecContainsFold applies the ContainsFold predicate on the "category_spec" field.
func CategorySpecContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldCategorySpec, v))
}

// SoldOutEQ applies the EQ predicate on the "sold_out" field.
func SoldOutEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSoldOut, v))
}

// SoldOutNEQ applies the NEQ predicate on the "sold_out" field.
func SoldOutNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldSoldOut, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
