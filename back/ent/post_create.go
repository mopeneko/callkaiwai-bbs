// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/ipaddresslog"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/post"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PostCreate) SetName(s string) *PostCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetGender sets the "gender" field.
func (pc *PostCreate) SetGender(i int) *PostCreate {
	pc.mutation.SetGender(i)
	return pc
}

// SetIntroduction sets the "introduction" field.
func (pc *PostCreate) SetIntroduction(s string) *PostCreate {
	pc.mutation.SetIntroduction(s)
	return pc
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (pc *PostCreate) SetNillableIntroduction(s *string) *PostCreate {
	if s != nil {
		pc.SetIntroduction(*s)
	}
	return pc
}

// SetTweetURL sets the "tweet_url" field.
func (pc *PostCreate) SetTweetURL(s string) *PostCreate {
	pc.mutation.SetTweetURL(s)
	return pc
}

// SetNillableTweetURL sets the "tweet_url" field if the given value is not nil.
func (pc *PostCreate) SetNillableTweetURL(s *string) *PostCreate {
	if s != nil {
		pc.SetTweetURL(*s)
	}
	return pc
}

// SetTiktokURL sets the "tiktok_url" field.
func (pc *PostCreate) SetTiktokURL(s string) *PostCreate {
	pc.mutation.SetTiktokURL(s)
	return pc
}

// SetNillableTiktokURL sets the "tiktok_url" field if the given value is not nil.
func (pc *PostCreate) SetNillableTiktokURL(s *string) *PostCreate {
	if s != nil {
		pc.SetTiktokURL(*s)
	}
	return pc
}

// SetContactURL sets the "contact_url" field.
func (pc *PostCreate) SetContactURL(s string) *PostCreate {
	pc.mutation.SetContactURL(s)
	return pc
}

// SetNillableContactURL sets the "contact_url" field if the given value is not nil.
func (pc *PostCreate) SetNillableContactURL(s *string) *PostCreate {
	if s != nil {
		pc.SetContactURL(*s)
	}
	return pc
}

// SetContactID sets the "contact_id" field.
func (pc *PostCreate) SetContactID(s string) *PostCreate {
	pc.mutation.SetContactID(s)
	return pc
}

// SetNillableContactID sets the "contact_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableContactID(s *string) *PostCreate {
	if s != nil {
		pc.SetContactID(*s)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(s string) *PostCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PostCreate) SetNillableID(s *string) *PostCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// AddIPAddressLogIDs adds the "ip_address_log" edge to the IPAddressLog entity by IDs.
func (pc *PostCreate) AddIPAddressLogIDs(ids ...string) *PostCreate {
	pc.mutation.AddIPAddressLogIDs(ids...)
	return pc
}

// AddIPAddressLog adds the "ip_address_log" edges to the IPAddressLog entity.
func (pc *PostCreate) AddIPAddressLog(i ...*IPAddressLog) *PostCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddIPAddressLogIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks[*Post, PostMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := post.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Post.name"`)}
	}
	if _, ok := pc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Post.gender"`)}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := post.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Post.gender": %w`, err)}
		}
	}
	if v, ok := pc.mutation.TweetURL(); ok {
		if err := post.TweetURLValidator(v); err != nil {
			return &ValidationError{Name: "tweet_url", err: fmt.Errorf(`ent: validator failed for field "Post.tweet_url": %w`, err)}
		}
	}
	if v, ok := pc.mutation.TiktokURL(); ok {
		if err := post.TiktokURLValidator(v); err != nil {
			return &ValidationError{Name: "tiktok_url", err: fmt.Errorf(`ent: validator failed for field "Post.tiktok_url": %w`, err)}
		}
	}
	if v, ok := pc.mutation.ContactURL(); ok {
		if err := post.ContactURLValidator(v); err != nil {
			return &ValidationError{Name: "contact_url", err: fmt.Errorf(`ent: validator failed for field "Post.contact_url": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
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
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Post.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(post.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(post.FieldGender, field.TypeInt, value)
		_node.Gender = value
	}
	if value, ok := pc.mutation.Introduction(); ok {
		_spec.SetField(post.FieldIntroduction, field.TypeString, value)
		_node.Introduction = value
	}
	if value, ok := pc.mutation.TweetURL(); ok {
		_spec.SetField(post.FieldTweetURL, field.TypeString, value)
		_node.TweetURL = value
	}
	if value, ok := pc.mutation.TiktokURL(); ok {
		_spec.SetField(post.FieldTiktokURL, field.TypeString, value)
		_node.TiktokURL = value
	}
	if value, ok := pc.mutation.ContactURL(); ok {
		_spec.SetField(post.FieldContactURL, field.TypeString, value)
		_node.ContactURL = value
	}
	if value, ok := pc.mutation.ContactID(); ok {
		_spec.SetField(post.FieldContactID, field.TypeString, value)
		_node.ContactID = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.IPAddressLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.IPAddressLogTable,
			Columns: []string{post.IPAddressLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ipaddresslog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
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
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
