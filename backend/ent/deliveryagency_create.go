// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/deliveryagency"
	"backend/ent/schema"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeliveryAgencyCreate is the builder for creating a DeliveryAgency entity.
type DeliveryAgencyCreate struct {
	config
	mutation *DeliveryAgencyMutation
	hooks    []Hook
}

// SetCountry sets the "country" field.
func (dac *DeliveryAgencyCreate) SetCountry(s string) *DeliveryAgencyCreate {
	dac.mutation.SetCountry(s)
	return dac
}

// SetVATReductionRate sets the "VAT_reduction_rate" field.
func (dac *DeliveryAgencyCreate) SetVATReductionRate(f float64) *DeliveryAgencyCreate {
	dac.mutation.SetVATReductionRate(f)
	return dac
}

// SetShippingFee sets the "shipping_fee" field.
func (dac *DeliveryAgencyCreate) SetShippingFee(ssf *schema.AgencyShippingFee) *DeliveryAgencyCreate {
	dac.mutation.SetShippingFee(ssf)
	return dac
}

// SetUpdatedAt sets the "updated_at" field.
func (dac *DeliveryAgencyCreate) SetUpdatedAt(t time.Time) *DeliveryAgencyCreate {
	dac.mutation.SetUpdatedAt(t)
	return dac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dac *DeliveryAgencyCreate) SetNillableUpdatedAt(t *time.Time) *DeliveryAgencyCreate {
	if t != nil {
		dac.SetUpdatedAt(*t)
	}
	return dac
}

// Mutation returns the DeliveryAgencyMutation object of the builder.
func (dac *DeliveryAgencyCreate) Mutation() *DeliveryAgencyMutation {
	return dac.mutation
}

// Save creates the DeliveryAgency in the database.
func (dac *DeliveryAgencyCreate) Save(ctx context.Context) (*DeliveryAgency, error) {
	dac.defaults()
	return withHooks(ctx, dac.sqlSave, dac.mutation, dac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DeliveryAgencyCreate) SaveX(ctx context.Context) *DeliveryAgency {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DeliveryAgencyCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DeliveryAgencyCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dac *DeliveryAgencyCreate) defaults() {
	if _, ok := dac.mutation.UpdatedAt(); !ok {
		v := deliveryagency.DefaultUpdatedAt()
		dac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dac *DeliveryAgencyCreate) check() error {
	if _, ok := dac.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "DeliveryAgency.country"`)}
	}
	if _, ok := dac.mutation.VATReductionRate(); !ok {
		return &ValidationError{Name: "VAT_reduction_rate", err: errors.New(`ent: missing required field "DeliveryAgency.VAT_reduction_rate"`)}
	}
	if _, ok := dac.mutation.ShippingFee(); !ok {
		return &ValidationError{Name: "shipping_fee", err: errors.New(`ent: missing required field "DeliveryAgency.shipping_fee"`)}
	}
	if _, ok := dac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeliveryAgency.updated_at"`)}
	}
	return nil
}

func (dac *DeliveryAgencyCreate) sqlSave(ctx context.Context) (*DeliveryAgency, error) {
	if err := dac.check(); err != nil {
		return nil, err
	}
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dac.mutation.id = &_node.ID
	dac.mutation.done = true
	return _node, nil
}

func (dac *DeliveryAgencyCreate) createSpec() (*DeliveryAgency, *sqlgraph.CreateSpec) {
	var (
		_node = &DeliveryAgency{config: dac.config}
		_spec = sqlgraph.NewCreateSpec(deliveryagency.Table, sqlgraph.NewFieldSpec(deliveryagency.FieldID, field.TypeInt))
	)
	if value, ok := dac.mutation.Country(); ok {
		_spec.SetField(deliveryagency.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := dac.mutation.VATReductionRate(); ok {
		_spec.SetField(deliveryagency.FieldVATReductionRate, field.TypeFloat64, value)
		_node.VATReductionRate = value
	}
	if value, ok := dac.mutation.ShippingFee(); ok {
		_spec.SetField(deliveryagency.FieldShippingFee, field.TypeJSON, value)
		_node.ShippingFee = value
	}
	if value, ok := dac.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryagency.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// DeliveryAgencyCreateBulk is the builder for creating many DeliveryAgency entities in bulk.
type DeliveryAgencyCreateBulk struct {
	config
	err      error
	builders []*DeliveryAgencyCreate
}

// Save creates the DeliveryAgency entities in the database.
func (dacb *DeliveryAgencyCreateBulk) Save(ctx context.Context) ([]*DeliveryAgency, error) {
	if dacb.err != nil {
		return nil, dacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DeliveryAgency, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeliveryAgencyMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DeliveryAgencyCreateBulk) SaveX(ctx context.Context) []*DeliveryAgency {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DeliveryAgencyCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DeliveryAgencyCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}
