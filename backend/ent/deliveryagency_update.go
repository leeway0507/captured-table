// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/deliveryagency"
	"backend/ent/predicate"
	"backend/ent/schema"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeliveryAgencyUpdate is the builder for updating DeliveryAgency entities.
type DeliveryAgencyUpdate struct {
	config
	hooks    []Hook
	mutation *DeliveryAgencyMutation
}

// Where appends a list predicates to the DeliveryAgencyUpdate builder.
func (dau *DeliveryAgencyUpdate) Where(ps ...predicate.DeliveryAgency) *DeliveryAgencyUpdate {
	dau.mutation.Where(ps...)
	return dau
}

// SetCountry sets the "country" field.
func (dau *DeliveryAgencyUpdate) SetCountry(s string) *DeliveryAgencyUpdate {
	dau.mutation.SetCountry(s)
	return dau
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (dau *DeliveryAgencyUpdate) SetNillableCountry(s *string) *DeliveryAgencyUpdate {
	if s != nil {
		dau.SetCountry(*s)
	}
	return dau
}

// SetVATReductionRate sets the "VAT_reduction_rate" field.
func (dau *DeliveryAgencyUpdate) SetVATReductionRate(f float64) *DeliveryAgencyUpdate {
	dau.mutation.ResetVATReductionRate()
	dau.mutation.SetVATReductionRate(f)
	return dau
}

// SetNillableVATReductionRate sets the "VAT_reduction_rate" field if the given value is not nil.
func (dau *DeliveryAgencyUpdate) SetNillableVATReductionRate(f *float64) *DeliveryAgencyUpdate {
	if f != nil {
		dau.SetVATReductionRate(*f)
	}
	return dau
}

// AddVATReductionRate adds f to the "VAT_reduction_rate" field.
func (dau *DeliveryAgencyUpdate) AddVATReductionRate(f float64) *DeliveryAgencyUpdate {
	dau.mutation.AddVATReductionRate(f)
	return dau
}

// SetShippingFee sets the "shipping_fee" field.
func (dau *DeliveryAgencyUpdate) SetShippingFee(ssf *schema.AgencyShippingFee) *DeliveryAgencyUpdate {
	dau.mutation.SetShippingFee(ssf)
	return dau
}

// SetUpdatedAt sets the "updated_at" field.
func (dau *DeliveryAgencyUpdate) SetUpdatedAt(t time.Time) *DeliveryAgencyUpdate {
	dau.mutation.SetUpdatedAt(t)
	return dau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dau *DeliveryAgencyUpdate) SetNillableUpdatedAt(t *time.Time) *DeliveryAgencyUpdate {
	if t != nil {
		dau.SetUpdatedAt(*t)
	}
	return dau
}

// Mutation returns the DeliveryAgencyMutation object of the builder.
func (dau *DeliveryAgencyUpdate) Mutation() *DeliveryAgencyMutation {
	return dau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dau *DeliveryAgencyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dau.sqlSave, dau.mutation, dau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dau *DeliveryAgencyUpdate) SaveX(ctx context.Context) int {
	affected, err := dau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dau *DeliveryAgencyUpdate) Exec(ctx context.Context) error {
	_, err := dau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dau *DeliveryAgencyUpdate) ExecX(ctx context.Context) {
	if err := dau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dau *DeliveryAgencyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(deliveryagency.Table, deliveryagency.Columns, sqlgraph.NewFieldSpec(deliveryagency.FieldID, field.TypeInt))
	if ps := dau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dau.mutation.Country(); ok {
		_spec.SetField(deliveryagency.FieldCountry, field.TypeString, value)
	}
	if value, ok := dau.mutation.VATReductionRate(); ok {
		_spec.SetField(deliveryagency.FieldVATReductionRate, field.TypeFloat64, value)
	}
	if value, ok := dau.mutation.AddedVATReductionRate(); ok {
		_spec.AddField(deliveryagency.FieldVATReductionRate, field.TypeFloat64, value)
	}
	if value, ok := dau.mutation.ShippingFee(); ok {
		_spec.SetField(deliveryagency.FieldShippingFee, field.TypeJSON, value)
	}
	if value, ok := dau.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryagency.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryagency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dau.mutation.done = true
	return n, nil
}

// DeliveryAgencyUpdateOne is the builder for updating a single DeliveryAgency entity.
type DeliveryAgencyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeliveryAgencyMutation
}

// SetCountry sets the "country" field.
func (dauo *DeliveryAgencyUpdateOne) SetCountry(s string) *DeliveryAgencyUpdateOne {
	dauo.mutation.SetCountry(s)
	return dauo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (dauo *DeliveryAgencyUpdateOne) SetNillableCountry(s *string) *DeliveryAgencyUpdateOne {
	if s != nil {
		dauo.SetCountry(*s)
	}
	return dauo
}

// SetVATReductionRate sets the "VAT_reduction_rate" field.
func (dauo *DeliveryAgencyUpdateOne) SetVATReductionRate(f float64) *DeliveryAgencyUpdateOne {
	dauo.mutation.ResetVATReductionRate()
	dauo.mutation.SetVATReductionRate(f)
	return dauo
}

// SetNillableVATReductionRate sets the "VAT_reduction_rate" field if the given value is not nil.
func (dauo *DeliveryAgencyUpdateOne) SetNillableVATReductionRate(f *float64) *DeliveryAgencyUpdateOne {
	if f != nil {
		dauo.SetVATReductionRate(*f)
	}
	return dauo
}

// AddVATReductionRate adds f to the "VAT_reduction_rate" field.
func (dauo *DeliveryAgencyUpdateOne) AddVATReductionRate(f float64) *DeliveryAgencyUpdateOne {
	dauo.mutation.AddVATReductionRate(f)
	return dauo
}

// SetShippingFee sets the "shipping_fee" field.
func (dauo *DeliveryAgencyUpdateOne) SetShippingFee(ssf *schema.AgencyShippingFee) *DeliveryAgencyUpdateOne {
	dauo.mutation.SetShippingFee(ssf)
	return dauo
}

// SetUpdatedAt sets the "updated_at" field.
func (dauo *DeliveryAgencyUpdateOne) SetUpdatedAt(t time.Time) *DeliveryAgencyUpdateOne {
	dauo.mutation.SetUpdatedAt(t)
	return dauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dauo *DeliveryAgencyUpdateOne) SetNillableUpdatedAt(t *time.Time) *DeliveryAgencyUpdateOne {
	if t != nil {
		dauo.SetUpdatedAt(*t)
	}
	return dauo
}

// Mutation returns the DeliveryAgencyMutation object of the builder.
func (dauo *DeliveryAgencyUpdateOne) Mutation() *DeliveryAgencyMutation {
	return dauo.mutation
}

// Where appends a list predicates to the DeliveryAgencyUpdate builder.
func (dauo *DeliveryAgencyUpdateOne) Where(ps ...predicate.DeliveryAgency) *DeliveryAgencyUpdateOne {
	dauo.mutation.Where(ps...)
	return dauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dauo *DeliveryAgencyUpdateOne) Select(field string, fields ...string) *DeliveryAgencyUpdateOne {
	dauo.fields = append([]string{field}, fields...)
	return dauo
}

// Save executes the query and returns the updated DeliveryAgency entity.
func (dauo *DeliveryAgencyUpdateOne) Save(ctx context.Context) (*DeliveryAgency, error) {
	return withHooks(ctx, dauo.sqlSave, dauo.mutation, dauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dauo *DeliveryAgencyUpdateOne) SaveX(ctx context.Context) *DeliveryAgency {
	node, err := dauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dauo *DeliveryAgencyUpdateOne) Exec(ctx context.Context) error {
	_, err := dauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dauo *DeliveryAgencyUpdateOne) ExecX(ctx context.Context) {
	if err := dauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dauo *DeliveryAgencyUpdateOne) sqlSave(ctx context.Context) (_node *DeliveryAgency, err error) {
	_spec := sqlgraph.NewUpdateSpec(deliveryagency.Table, deliveryagency.Columns, sqlgraph.NewFieldSpec(deliveryagency.FieldID, field.TypeInt))
	id, ok := dauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeliveryAgency.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deliveryagency.FieldID)
		for _, f := range fields {
			if !deliveryagency.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deliveryagency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dauo.mutation.Country(); ok {
		_spec.SetField(deliveryagency.FieldCountry, field.TypeString, value)
	}
	if value, ok := dauo.mutation.VATReductionRate(); ok {
		_spec.SetField(deliveryagency.FieldVATReductionRate, field.TypeFloat64, value)
	}
	if value, ok := dauo.mutation.AddedVATReductionRate(); ok {
		_spec.AddField(deliveryagency.FieldVATReductionRate, field.TypeFloat64, value)
	}
	if value, ok := dauo.mutation.ShippingFee(); ok {
		_spec.SetField(deliveryagency.FieldShippingFee, field.TypeJSON, value)
	}
	if value, ok := dauo.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryagency.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &DeliveryAgency{config: dauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryagency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dauo.mutation.done = true
	return _node, nil
}
