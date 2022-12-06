// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MangoSteen0903/go-blog-graphql/ent/like"
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/predicate"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks    []Hook
	mutation *LikeMutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LikeUpdate) SetCreatedAt(t time.Time) *LikeUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableCreatedAt(t *time.Time) *LikeUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// AddPostIDs adds the "Posts" edge to the Post entity by IDs.
func (lu *LikeUpdate) AddPostIDs(ids ...int) *LikeUpdate {
	lu.mutation.AddPostIDs(ids...)
	return lu
}

// AddPosts adds the "Posts" edges to the Post entity.
func (lu *LikeUpdate) AddPosts(p ...*Post) *LikeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.AddPostIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (lu *LikeUpdate) AddOwnerIDs(ids ...int) *LikeUpdate {
	lu.mutation.AddOwnerIDs(ids...)
	return lu
}

// AddOwner adds the "owner" edges to the User entity.
func (lu *LikeUpdate) AddOwner(u ...*User) *LikeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.AddOwnerIDs(ids...)
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// ClearPosts clears all "Posts" edges to the Post entity.
func (lu *LikeUpdate) ClearPosts() *LikeUpdate {
	lu.mutation.ClearPosts()
	return lu
}

// RemovePostIDs removes the "Posts" edge to Post entities by IDs.
func (lu *LikeUpdate) RemovePostIDs(ids ...int) *LikeUpdate {
	lu.mutation.RemovePostIDs(ids...)
	return lu
}

// RemovePosts removes "Posts" edges to Post entities.
func (lu *LikeUpdate) RemovePosts(p ...*Post) *LikeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.RemovePostIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (lu *LikeUpdate) ClearOwner() *LikeUpdate {
	lu.mutation.ClearOwner()
	return lu
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (lu *LikeUpdate) RemoveOwnerIDs(ids ...int) *LikeUpdate {
	lu.mutation.RemoveOwnerIDs(ids...)
	return lu
}

// RemoveOwner removes "owner" edges to User entities.
func (lu *LikeUpdate) RemoveOwner(u ...*User) *LikeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: like.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(like.FieldCreatedAt, field.TypeTime, value)
	}
	if lu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !lu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !lu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeMutation
}

// SetCreatedAt sets the "created_at" field.
func (luo *LikeUpdateOne) SetCreatedAt(t time.Time) *LikeUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableCreatedAt(t *time.Time) *LikeUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// AddPostIDs adds the "Posts" edge to the Post entity by IDs.
func (luo *LikeUpdateOne) AddPostIDs(ids ...int) *LikeUpdateOne {
	luo.mutation.AddPostIDs(ids...)
	return luo
}

// AddPosts adds the "Posts" edges to the Post entity.
func (luo *LikeUpdateOne) AddPosts(p ...*Post) *LikeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.AddPostIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (luo *LikeUpdateOne) AddOwnerIDs(ids ...int) *LikeUpdateOne {
	luo.mutation.AddOwnerIDs(ids...)
	return luo
}

// AddOwner adds the "owner" edges to the User entity.
func (luo *LikeUpdateOne) AddOwner(u ...*User) *LikeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.AddOwnerIDs(ids...)
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// ClearPosts clears all "Posts" edges to the Post entity.
func (luo *LikeUpdateOne) ClearPosts() *LikeUpdateOne {
	luo.mutation.ClearPosts()
	return luo
}

// RemovePostIDs removes the "Posts" edge to Post entities by IDs.
func (luo *LikeUpdateOne) RemovePostIDs(ids ...int) *LikeUpdateOne {
	luo.mutation.RemovePostIDs(ids...)
	return luo
}

// RemovePosts removes "Posts" edges to Post entities.
func (luo *LikeUpdateOne) RemovePosts(p ...*Post) *LikeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.RemovePostIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (luo *LikeUpdateOne) ClearOwner() *LikeUpdateOne {
	luo.mutation.ClearOwner()
	return luo
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (luo *LikeUpdateOne) RemoveOwnerIDs(ids ...int) *LikeUpdateOne {
	luo.mutation.RemoveOwnerIDs(ids...)
	return luo
}

// RemoveOwner removes "owner" edges to User entities.
func (luo *LikeUpdateOne) RemoveOwner(u ...*User) *LikeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.RemoveOwnerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	var (
		err  error
		node *Like
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Like)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LikeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: like.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Like.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for _, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(like.FieldCreatedAt, field.TypeTime, value)
	}
	if luo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !luo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.PostsTable,
			Columns: like.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !luo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   like.OwnerTable,
			Columns: like.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
