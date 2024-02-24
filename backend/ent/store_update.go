// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/product"
	"backend/ent/schema"
	"backend/ent/store"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreUpdate is the builder for updating Store entities.
type StoreUpdate struct {
	config
	hooks    []Hook
	mutation *StoreMutation
}

// Where appends a list predicates to the StoreUpdate builder.
func (su *StoreUpdate) Where(ps ...predicate.Store) *StoreUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStoreName sets the "store_name" field.
func (su *StoreUpdate) SetStoreName(s string) *StoreUpdate {
	su.mutation.SetStoreName(s)
	return su
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStoreName(s *string) *StoreUpdate {
	if s != nil {
		su.SetStoreName(*s)
	}
	return su
}

// SetURL sets the "url" field.
func (su *StoreUpdate) SetURL(s string) *StoreUpdate {
	su.mutation.SetURL(s)
	return su
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (su *StoreUpdate) SetNillableURL(s *string) *StoreUpdate {
	if s != nil {
		su.SetURL(*s)
	}
	return su
}

// SetCountry sets the "country" field.
func (su *StoreUpdate) SetCountry(s string) *StoreUpdate {
	su.mutation.SetCountry(s)
	return su
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (su *StoreUpdate) SetNillableCountry(s *string) *StoreUpdate {
	if s != nil {
		su.SetCountry(*s)
	}
	return su
}

// SetCurrency sets the "currency" field.
func (su *StoreUpdate) SetCurrency(s string) *StoreUpdate {
	su.mutation.SetCurrency(s)
	return su
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (su *StoreUpdate) SetNillableCurrency(s *string) *StoreUpdate {
	if s != nil {
		su.SetCurrency(*s)
	}
	return su
}

// SetTaxReduction sets the "tax_reduction" field.
func (su *StoreUpdate) SetTaxReduction(f float64) *StoreUpdate {
	su.mutation.ResetTaxReduction()
	su.mutation.SetTaxReduction(f)
	return su
}

// SetNillableTaxReduction sets the "tax_reduction" field if the given value is not nil.
func (su *StoreUpdate) SetNillableTaxReduction(f *float64) *StoreUpdate {
	if f != nil {
		su.SetTaxReduction(*f)
	}
	return su
}

// AddTaxReduction adds f to the "tax_reduction" field.
func (su *StoreUpdate) AddTaxReduction(f float64) *StoreUpdate {
	su.mutation.AddTaxReduction(f)
	return su
}

// SetIntlShippingFee sets the "intl_shipping_fee" field.
func (su *StoreUpdate) SetIntlShippingFee(sf *schema.ShippingFee) *StoreUpdate {
	su.mutation.SetIntlShippingFee(sf)
	return su
}

// SetIntlFreeShippingPrice sets the "intl_free_shipping_price" field.
func (su *StoreUpdate) SetIntlFreeShippingPrice(i int) *StoreUpdate {
	su.mutation.ResetIntlFreeShippingPrice()
	su.mutation.SetIntlFreeShippingPrice(i)
	return su
}

// SetNillableIntlFreeShippingPrice sets the "intl_free_shipping_price" field if the given value is not nil.
func (su *StoreUpdate) SetNillableIntlFreeShippingPrice(i *int) *StoreUpdate {
	if i != nil {
		su.SetIntlFreeShippingPrice(*i)
	}
	return su
}

// AddIntlFreeShippingPrice adds i to the "intl_free_shipping_price" field.
func (su *StoreUpdate) AddIntlFreeShippingPrice(i int) *StoreUpdate {
	su.mutation.AddIntlFreeShippingPrice(i)
	return su
}

// SetDomesticShippingFee sets the "domestic_shipping_fee" field.
func (su *StoreUpdate) SetDomesticShippingFee(f float64) *StoreUpdate {
	su.mutation.ResetDomesticShippingFee()
	su.mutation.SetDomesticShippingFee(f)
	return su
}

// SetNillableDomesticShippingFee sets the "domestic_shipping_fee" field if the given value is not nil.
func (su *StoreUpdate) SetNillableDomesticShippingFee(f *float64) *StoreUpdate {
	if f != nil {
		su.SetDomesticShippingFee(*f)
	}
	return su
}

// AddDomesticShippingFee adds f to the "domestic_shipping_fee" field.
func (su *StoreUpdate) AddDomesticShippingFee(f float64) *StoreUpdate {
	su.mutation.AddDomesticShippingFee(f)
	return su
}

// SetDomesticFreeShippingPrice sets the "domestic_free_shipping_price" field.
func (su *StoreUpdate) SetDomesticFreeShippingPrice(f float64) *StoreUpdate {
	su.mutation.ResetDomesticFreeShippingPrice()
	su.mutation.SetDomesticFreeShippingPrice(f)
	return su
}

// SetNillableDomesticFreeShippingPrice sets the "domestic_free_shipping_price" field if the given value is not nil.
func (su *StoreUpdate) SetNillableDomesticFreeShippingPrice(f *float64) *StoreUpdate {
	if f != nil {
		su.SetDomesticFreeShippingPrice(*f)
	}
	return su
}

// AddDomesticFreeShippingPrice adds f to the "domestic_free_shipping_price" field.
func (su *StoreUpdate) AddDomesticFreeShippingPrice(f float64) *StoreUpdate {
	su.mutation.AddDomesticFreeShippingPrice(f)
	return su
}

// SetDeliveryAgency sets the "delivery_agency" field.
func (su *StoreUpdate) SetDeliveryAgency(s string) *StoreUpdate {
	su.mutation.SetDeliveryAgency(s)
	return su
}

// SetNillableDeliveryAgency sets the "delivery_agency" field if the given value is not nil.
func (su *StoreUpdate) SetNillableDeliveryAgency(s *string) *StoreUpdate {
	if s != nil {
		su.SetDeliveryAgency(*s)
	}
	return su
}

// SetBrokerFee sets the "broker_fee" field.
func (su *StoreUpdate) SetBrokerFee(b bool) *StoreUpdate {
	su.mutation.SetBrokerFee(b)
	return su
}

// SetNillableBrokerFee sets the "broker_fee" field if the given value is not nil.
func (su *StoreUpdate) SetNillableBrokerFee(b *bool) *StoreUpdate {
	if b != nil {
		su.SetBrokerFee(*b)
	}
	return su
}

// SetDdp sets the "ddp" field.
func (su *StoreUpdate) SetDdp(b bool) *StoreUpdate {
	su.mutation.SetDdp(b)
	return su
}

// SetNillableDdp sets the "ddp" field if the given value is not nil.
func (su *StoreUpdate) SetNillableDdp(b *bool) *StoreUpdate {
	if b != nil {
		su.SetDdp(*b)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StoreUpdate) SetUpdatedAt(t time.Time) *StoreUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *StoreUpdate) SetNillableUpdatedAt(t *time.Time) *StoreUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (su *StoreUpdate) AddProductIDs(ids ...int) *StoreUpdate {
	su.mutation.AddProductIDs(ids...)
	return su
}

// AddProducts adds the "products" edges to the Product entity.
func (su *StoreUpdate) AddProducts(p ...*Product) *StoreUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddProductIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (su *StoreUpdate) Mutation() *StoreMutation {
	return su.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (su *StoreUpdate) ClearProducts() *StoreUpdate {
	su.mutation.ClearProducts()
	return su
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (su *StoreUpdate) RemoveProductIDs(ids ...int) *StoreUpdate {
	su.mutation.RemoveProductIDs(ids...)
	return su
}

// RemoveProducts removes "products" edges to Product entities.
func (su *StoreUpdate) RemoveProducts(p ...*Product) *StoreUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StoreUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StoreUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StoreUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(store.Table, store.Columns, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.StoreName(); ok {
		_spec.SetField(store.FieldStoreName, field.TypeString, value)
	}
	if value, ok := su.mutation.URL(); ok {
		_spec.SetField(store.FieldURL, field.TypeString, value)
	}
	if value, ok := su.mutation.Country(); ok {
		_spec.SetField(store.FieldCountry, field.TypeString, value)
	}
	if value, ok := su.mutation.Currency(); ok {
		_spec.SetField(store.FieldCurrency, field.TypeString, value)
	}
	if value, ok := su.mutation.TaxReduction(); ok {
		_spec.SetField(store.FieldTaxReduction, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedTaxReduction(); ok {
		_spec.AddField(store.FieldTaxReduction, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.IntlShippingFee(); ok {
		_spec.SetField(store.FieldIntlShippingFee, field.TypeJSON, value)
	}
	if value, ok := su.mutation.IntlFreeShippingPrice(); ok {
		_spec.SetField(store.FieldIntlFreeShippingPrice, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedIntlFreeShippingPrice(); ok {
		_spec.AddField(store.FieldIntlFreeShippingPrice, field.TypeInt, value)
	}
	if value, ok := su.mutation.DomesticShippingFee(); ok {
		_spec.SetField(store.FieldDomesticShippingFee, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedDomesticShippingFee(); ok {
		_spec.AddField(store.FieldDomesticShippingFee, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.DomesticFreeShippingPrice(); ok {
		_spec.SetField(store.FieldDomesticFreeShippingPrice, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedDomesticFreeShippingPrice(); ok {
		_spec.AddField(store.FieldDomesticFreeShippingPrice, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.DeliveryAgency(); ok {
		_spec.SetField(store.FieldDeliveryAgency, field.TypeString, value)
	}
	if value, ok := su.mutation.BrokerFee(); ok {
		_spec.SetField(store.FieldBrokerFee, field.TypeBool, value)
	}
	if value, ok := su.mutation.Ddp(); ok {
		_spec.SetField(store.FieldDdp, field.TypeBool, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(store.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedProductsIDs(); len(nodes) > 0 && !su.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StoreUpdateOne is the builder for updating a single Store entity.
type StoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreMutation
}

// SetStoreName sets the "store_name" field.
func (suo *StoreUpdateOne) SetStoreName(s string) *StoreUpdateOne {
	suo.mutation.SetStoreName(s)
	return suo
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStoreName(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStoreName(*s)
	}
	return suo
}

// SetURL sets the "url" field.
func (suo *StoreUpdateOne) SetURL(s string) *StoreUpdateOne {
	suo.mutation.SetURL(s)
	return suo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableURL(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetURL(*s)
	}
	return suo
}

// SetCountry sets the "country" field.
func (suo *StoreUpdateOne) SetCountry(s string) *StoreUpdateOne {
	suo.mutation.SetCountry(s)
	return suo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableCountry(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetCountry(*s)
	}
	return suo
}

// SetCurrency sets the "currency" field.
func (suo *StoreUpdateOne) SetCurrency(s string) *StoreUpdateOne {
	suo.mutation.SetCurrency(s)
	return suo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableCurrency(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetCurrency(*s)
	}
	return suo
}

// SetTaxReduction sets the "tax_reduction" field.
func (suo *StoreUpdateOne) SetTaxReduction(f float64) *StoreUpdateOne {
	suo.mutation.ResetTaxReduction()
	suo.mutation.SetTaxReduction(f)
	return suo
}

// SetNillableTaxReduction sets the "tax_reduction" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableTaxReduction(f *float64) *StoreUpdateOne {
	if f != nil {
		suo.SetTaxReduction(*f)
	}
	return suo
}

// AddTaxReduction adds f to the "tax_reduction" field.
func (suo *StoreUpdateOne) AddTaxReduction(f float64) *StoreUpdateOne {
	suo.mutation.AddTaxReduction(f)
	return suo
}

// SetIntlShippingFee sets the "intl_shipping_fee" field.
func (suo *StoreUpdateOne) SetIntlShippingFee(sf *schema.ShippingFee) *StoreUpdateOne {
	suo.mutation.SetIntlShippingFee(sf)
	return suo
}

// SetIntlFreeShippingPrice sets the "intl_free_shipping_price" field.
func (suo *StoreUpdateOne) SetIntlFreeShippingPrice(i int) *StoreUpdateOne {
	suo.mutation.ResetIntlFreeShippingPrice()
	suo.mutation.SetIntlFreeShippingPrice(i)
	return suo
}

// SetNillableIntlFreeShippingPrice sets the "intl_free_shipping_price" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableIntlFreeShippingPrice(i *int) *StoreUpdateOne {
	if i != nil {
		suo.SetIntlFreeShippingPrice(*i)
	}
	return suo
}

// AddIntlFreeShippingPrice adds i to the "intl_free_shipping_price" field.
func (suo *StoreUpdateOne) AddIntlFreeShippingPrice(i int) *StoreUpdateOne {
	suo.mutation.AddIntlFreeShippingPrice(i)
	return suo
}

// SetDomesticShippingFee sets the "domestic_shipping_fee" field.
func (suo *StoreUpdateOne) SetDomesticShippingFee(f float64) *StoreUpdateOne {
	suo.mutation.ResetDomesticShippingFee()
	suo.mutation.SetDomesticShippingFee(f)
	return suo
}

// SetNillableDomesticShippingFee sets the "domestic_shipping_fee" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableDomesticShippingFee(f *float64) *StoreUpdateOne {
	if f != nil {
		suo.SetDomesticShippingFee(*f)
	}
	return suo
}

// AddDomesticShippingFee adds f to the "domestic_shipping_fee" field.
func (suo *StoreUpdateOne) AddDomesticShippingFee(f float64) *StoreUpdateOne {
	suo.mutation.AddDomesticShippingFee(f)
	return suo
}

// SetDomesticFreeShippingPrice sets the "domestic_free_shipping_price" field.
func (suo *StoreUpdateOne) SetDomesticFreeShippingPrice(f float64) *StoreUpdateOne {
	suo.mutation.ResetDomesticFreeShippingPrice()
	suo.mutation.SetDomesticFreeShippingPrice(f)
	return suo
}

// SetNillableDomesticFreeShippingPrice sets the "domestic_free_shipping_price" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableDomesticFreeShippingPrice(f *float64) *StoreUpdateOne {
	if f != nil {
		suo.SetDomesticFreeShippingPrice(*f)
	}
	return suo
}

// AddDomesticFreeShippingPrice adds f to the "domestic_free_shipping_price" field.
func (suo *StoreUpdateOne) AddDomesticFreeShippingPrice(f float64) *StoreUpdateOne {
	suo.mutation.AddDomesticFreeShippingPrice(f)
	return suo
}

// SetDeliveryAgency sets the "delivery_agency" field.
func (suo *StoreUpdateOne) SetDeliveryAgency(s string) *StoreUpdateOne {
	suo.mutation.SetDeliveryAgency(s)
	return suo
}

// SetNillableDeliveryAgency sets the "delivery_agency" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableDeliveryAgency(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetDeliveryAgency(*s)
	}
	return suo
}

// SetBrokerFee sets the "broker_fee" field.
func (suo *StoreUpdateOne) SetBrokerFee(b bool) *StoreUpdateOne {
	suo.mutation.SetBrokerFee(b)
	return suo
}

// SetNillableBrokerFee sets the "broker_fee" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableBrokerFee(b *bool) *StoreUpdateOne {
	if b != nil {
		suo.SetBrokerFee(*b)
	}
	return suo
}

// SetDdp sets the "ddp" field.
func (suo *StoreUpdateOne) SetDdp(b bool) *StoreUpdateOne {
	suo.mutation.SetDdp(b)
	return suo
}

// SetNillableDdp sets the "ddp" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableDdp(b *bool) *StoreUpdateOne {
	if b != nil {
		suo.SetDdp(*b)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StoreUpdateOne) SetUpdatedAt(t time.Time) *StoreUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableUpdatedAt(t *time.Time) *StoreUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (suo *StoreUpdateOne) AddProductIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.AddProductIDs(ids...)
	return suo
}

// AddProducts adds the "products" edges to the Product entity.
func (suo *StoreUpdateOne) AddProducts(p ...*Product) *StoreUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddProductIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (suo *StoreUpdateOne) Mutation() *StoreMutation {
	return suo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (suo *StoreUpdateOne) ClearProducts() *StoreUpdateOne {
	suo.mutation.ClearProducts()
	return suo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (suo *StoreUpdateOne) RemoveProductIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.RemoveProductIDs(ids...)
	return suo
}

// RemoveProducts removes "products" edges to Product entities.
func (suo *StoreUpdateOne) RemoveProducts(p ...*Product) *StoreUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the StoreUpdate builder.
func (suo *StoreUpdateOne) Where(ps ...predicate.Store) *StoreUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StoreUpdateOne) Select(field string, fields ...string) *StoreUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Store entity.
func (suo *StoreUpdateOne) Save(ctx context.Context) (*Store, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StoreUpdateOne) SaveX(ctx context.Context) *Store {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StoreUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StoreUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StoreUpdateOne) sqlSave(ctx context.Context) (_node *Store, err error) {
	_spec := sqlgraph.NewUpdateSpec(store.Table, store.Columns, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Store.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, store.FieldID)
		for _, f := range fields {
			if !store.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != store.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.StoreName(); ok {
		_spec.SetField(store.FieldStoreName, field.TypeString, value)
	}
	if value, ok := suo.mutation.URL(); ok {
		_spec.SetField(store.FieldURL, field.TypeString, value)
	}
	if value, ok := suo.mutation.Country(); ok {
		_spec.SetField(store.FieldCountry, field.TypeString, value)
	}
	if value, ok := suo.mutation.Currency(); ok {
		_spec.SetField(store.FieldCurrency, field.TypeString, value)
	}
	if value, ok := suo.mutation.TaxReduction(); ok {
		_spec.SetField(store.FieldTaxReduction, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedTaxReduction(); ok {
		_spec.AddField(store.FieldTaxReduction, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.IntlShippingFee(); ok {
		_spec.SetField(store.FieldIntlShippingFee, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.IntlFreeShippingPrice(); ok {
		_spec.SetField(store.FieldIntlFreeShippingPrice, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedIntlFreeShippingPrice(); ok {
		_spec.AddField(store.FieldIntlFreeShippingPrice, field.TypeInt, value)
	}
	if value, ok := suo.mutation.DomesticShippingFee(); ok {
		_spec.SetField(store.FieldDomesticShippingFee, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedDomesticShippingFee(); ok {
		_spec.AddField(store.FieldDomesticShippingFee, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.DomesticFreeShippingPrice(); ok {
		_spec.SetField(store.FieldDomesticFreeShippingPrice, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedDomesticFreeShippingPrice(); ok {
		_spec.AddField(store.FieldDomesticFreeShippingPrice, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.DeliveryAgency(); ok {
		_spec.SetField(store.FieldDeliveryAgency, field.TypeString, value)
	}
	if value, ok := suo.mutation.BrokerFee(); ok {
		_spec.SetField(store.FieldBrokerFee, field.TypeBool, value)
	}
	if value, ok := suo.mutation.Ddp(); ok {
		_spec.SetField(store.FieldDdp, field.TypeBool, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(store.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !suo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ProductsTable,
			Columns: []string{store.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Store{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}