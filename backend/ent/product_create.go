// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/product"
	"backend/ent/store"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetStoreName sets the "store_name" field.
func (pc *ProductCreate) SetStoreName(s string) *ProductCreate {
	pc.mutation.SetStoreName(s)
	return pc
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (pc *ProductCreate) SetNillableStoreName(s *string) *ProductCreate {
	if s != nil {
		pc.SetStoreName(*s)
	}
	return pc
}

// SetBrand sets the "brand" field.
func (pc *ProductCreate) SetBrand(s string) *ProductCreate {
	pc.mutation.SetBrand(s)
	return pc
}

// SetProductName sets the "product_name" field.
func (pc *ProductCreate) SetProductName(s string) *ProductCreate {
	pc.mutation.SetProductName(s)
	return pc
}

// SetProductImgURL sets the "product_img_url" field.
func (pc *ProductCreate) SetProductImgURL(s string) *ProductCreate {
	pc.mutation.SetProductImgURL(s)
	return pc
}

// SetProductURL sets the "product_url" field.
func (pc *ProductCreate) SetProductURL(s string) *ProductCreate {
	pc.mutation.SetProductURL(s)
	return pc
}

// SetPriceCurrency sets the "price_currency" field.
func (pc *ProductCreate) SetPriceCurrency(s string) *ProductCreate {
	pc.mutation.SetPriceCurrency(s)
	return pc
}

// SetRetailPrice sets the "retail_price" field.
func (pc *ProductCreate) SetRetailPrice(f float64) *ProductCreate {
	pc.mutation.SetRetailPrice(f)
	return pc
}

// SetSalePrice sets the "sale_price" field.
func (pc *ProductCreate) SetSalePrice(f float64) *ProductCreate {
	pc.mutation.SetSalePrice(f)
	return pc
}

// SetKorBrand sets the "kor_brand" field.
func (pc *ProductCreate) SetKorBrand(s string) *ProductCreate {
	pc.mutation.SetKorBrand(s)
	return pc
}

// SetNillableKorBrand sets the "kor_brand" field if the given value is not nil.
func (pc *ProductCreate) SetNillableKorBrand(s *string) *ProductCreate {
	if s != nil {
		pc.SetKorBrand(*s)
	}
	return pc
}

// SetKorProductName sets the "kor_product_name" field.
func (pc *ProductCreate) SetKorProductName(s string) *ProductCreate {
	pc.mutation.SetKorProductName(s)
	return pc
}

// SetNillableKorProductName sets the "kor_product_name" field if the given value is not nil.
func (pc *ProductCreate) SetNillableKorProductName(s *string) *ProductCreate {
	if s != nil {
		pc.SetKorProductName(*s)
	}
	return pc
}

// SetProductID sets the "product_id" field.
func (pc *ProductCreate) SetProductID(s string) *ProductCreate {
	pc.mutation.SetProductID(s)
	return pc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pc *ProductCreate) SetNillableProductID(s *string) *ProductCreate {
	if s != nil {
		pc.SetProductID(*s)
	}
	return pc
}

// SetGender sets the "gender" field.
func (pc *ProductCreate) SetGender(s string) *ProductCreate {
	pc.mutation.SetGender(s)
	return pc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (pc *ProductCreate) SetNillableGender(s *string) *ProductCreate {
	if s != nil {
		pc.SetGender(*s)
	}
	return pc
}

// SetColor sets the "color" field.
func (pc *ProductCreate) SetColor(s string) *ProductCreate {
	pc.mutation.SetColor(s)
	return pc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (pc *ProductCreate) SetNillableColor(s *string) *ProductCreate {
	if s != nil {
		pc.SetColor(*s)
	}
	return pc
}

// SetCategory sets the "category" field.
func (pc *ProductCreate) SetCategory(s string) *ProductCreate {
	pc.mutation.SetCategory(s)
	return pc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCategory(s *string) *ProductCreate {
	if s != nil {
		pc.SetCategory(*s)
	}
	return pc
}

// SetCategorySpec sets the "category_spec" field.
func (pc *ProductCreate) SetCategorySpec(s string) *ProductCreate {
	pc.mutation.SetCategorySpec(s)
	return pc
}

// SetNillableCategorySpec sets the "category_spec" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCategorySpec(s *string) *ProductCreate {
	if s != nil {
		pc.SetCategorySpec(*s)
	}
	return pc
}

// SetSoldOut sets the "sold_out" field.
func (pc *ProductCreate) SetSoldOut(b bool) *ProductCreate {
	pc.mutation.SetSoldOut(b)
	return pc
}

// SetNillableSoldOut sets the "sold_out" field if the given value is not nil.
func (pc *ProductCreate) SetNillableSoldOut(b *bool) *ProductCreate {
	if b != nil {
		pc.SetSoldOut(*b)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductCreate) SetUpdatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetStoreID sets the "store" edge to the Store entity by ID.
func (pc *ProductCreate) SetStoreID(id string) *ProductCreate {
	pc.mutation.SetStoreID(id)
	return pc
}

// SetNillableStoreID sets the "store" edge to the Store entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableStoreID(id *string) *ProductCreate {
	if id != nil {
		pc = pc.SetStoreID(*id)
	}
	return pc
}

// SetStore sets the "store" edge to the Store entity.
func (pc *ProductCreate) SetStore(s *Store) *ProductCreate {
	return pc.SetStoreID(s.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.SoldOut(); !ok {
		v := product.DefaultSoldOut
		pc.mutation.SetSoldOut(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := product.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Brand(); !ok {
		return &ValidationError{Name: "brand", err: errors.New(`ent: missing required field "Product.brand"`)}
	}
	if _, ok := pc.mutation.ProductName(); !ok {
		return &ValidationError{Name: "product_name", err: errors.New(`ent: missing required field "Product.product_name"`)}
	}
	if _, ok := pc.mutation.ProductImgURL(); !ok {
		return &ValidationError{Name: "product_img_url", err: errors.New(`ent: missing required field "Product.product_img_url"`)}
	}
	if _, ok := pc.mutation.ProductURL(); !ok {
		return &ValidationError{Name: "product_url", err: errors.New(`ent: missing required field "Product.product_url"`)}
	}
	if _, ok := pc.mutation.PriceCurrency(); !ok {
		return &ValidationError{Name: "price_currency", err: errors.New(`ent: missing required field "Product.price_currency"`)}
	}
	if _, ok := pc.mutation.RetailPrice(); !ok {
		return &ValidationError{Name: "retail_price", err: errors.New(`ent: missing required field "Product.retail_price"`)}
	}
	if _, ok := pc.mutation.SalePrice(); !ok {
		return &ValidationError{Name: "sale_price", err: errors.New(`ent: missing required field "Product.sale_price"`)}
	}
	if _, ok := pc.mutation.SoldOut(); !ok {
		return &ValidationError{Name: "sold_out", err: errors.New(`ent: missing required field "Product.sold_out"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Product.updated_at"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Brand(); ok {
		_spec.SetField(product.FieldBrand, field.TypeString, value)
		_node.Brand = value
	}
	if value, ok := pc.mutation.ProductName(); ok {
		_spec.SetField(product.FieldProductName, field.TypeString, value)
		_node.ProductName = value
	}
	if value, ok := pc.mutation.ProductImgURL(); ok {
		_spec.SetField(product.FieldProductImgURL, field.TypeString, value)
		_node.ProductImgURL = value
	}
	if value, ok := pc.mutation.ProductURL(); ok {
		_spec.SetField(product.FieldProductURL, field.TypeString, value)
		_node.ProductURL = value
	}
	if value, ok := pc.mutation.PriceCurrency(); ok {
		_spec.SetField(product.FieldPriceCurrency, field.TypeString, value)
		_node.PriceCurrency = value
	}
	if value, ok := pc.mutation.RetailPrice(); ok {
		_spec.SetField(product.FieldRetailPrice, field.TypeFloat64, value)
		_node.RetailPrice = value
	}
	if value, ok := pc.mutation.SalePrice(); ok {
		_spec.SetField(product.FieldSalePrice, field.TypeFloat64, value)
		_node.SalePrice = value
	}
	if value, ok := pc.mutation.KorBrand(); ok {
		_spec.SetField(product.FieldKorBrand, field.TypeString, value)
		_node.KorBrand = value
	}
	if value, ok := pc.mutation.KorProductName(); ok {
		_spec.SetField(product.FieldKorProductName, field.TypeString, value)
		_node.KorProductName = value
	}
	if value, ok := pc.mutation.ProductID(); ok {
		_spec.SetField(product.FieldProductID, field.TypeString, value)
		_node.ProductID = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(product.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := pc.mutation.Color(); ok {
		_spec.SetField(product.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := pc.mutation.Category(); ok {
		_spec.SetField(product.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := pc.mutation.CategorySpec(); ok {
		_spec.SetField(product.FieldCategorySpec, field.TypeString, value)
		_node.CategorySpec = value
	}
	if value, ok := pc.mutation.SoldOut(); ok {
		_spec.SetField(product.FieldSoldOut, field.TypeBool, value)
		_node.SoldOut = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.StoreTable,
			Columns: []string{product.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StoreName = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	err      error
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
