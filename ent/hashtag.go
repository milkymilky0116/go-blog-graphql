// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/MangoSteen0903/go-blog-graphql/ent/hashtag"
)

// Hashtag is the model entity for the Hashtag schema.
type Hashtag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hashtag holds the value of the "hashtag" field.
	Hashtag string `json:"hashtag,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HashtagQuery when eager-loading is set.
	Edges HashtagEdges `json:"edges"`
}

// HashtagEdges holds the relations/edges for other nodes in the graph.
type HashtagEdges struct {
	// Posts holds the value of the Posts edge.
	Posts []*Post `json:"Posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedPosts map[string][]*Post
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e HashtagEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "Posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hashtag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hashtag.FieldID:
			values[i] = new(sql.NullInt64)
		case hashtag.FieldHashtag:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Hashtag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hashtag fields.
func (h *Hashtag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hashtag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hashtag.FieldHashtag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashtag", values[i])
			} else if value.Valid {
				h.Hashtag = value.String
			}
		}
	}
	return nil
}

// QueryPosts queries the "Posts" edge of the Hashtag entity.
func (h *Hashtag) QueryPosts() *PostQuery {
	return (&HashtagClient{config: h.config}).QueryPosts(h)
}

// Update returns a builder for updating this Hashtag.
// Note that you need to call Hashtag.Unwrap() before calling this method if this Hashtag
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hashtag) Update() *HashtagUpdateOne {
	return (&HashtagClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Hashtag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hashtag) Unwrap() *Hashtag {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hashtag is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hashtag) String() string {
	var builder strings.Builder
	builder.WriteString("Hashtag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("hashtag=")
	builder.WriteString(h.Hashtag)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPosts returns the Posts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (h *Hashtag) NamedPosts(name string) ([]*Post, error) {
	if h.Edges.namedPosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := h.Edges.namedPosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (h *Hashtag) appendNamedPosts(name string, edges ...*Post) {
	if h.Edges.namedPosts == nil {
		h.Edges.namedPosts = make(map[string][]*Post)
	}
	if len(edges) == 0 {
		h.Edges.namedPosts[name] = []*Post{}
	} else {
		h.Edges.namedPosts[name] = append(h.Edges.namedPosts[name], edges...)
	}
}

// Hashtags is a parsable slice of Hashtag.
type Hashtags []*Hashtag

func (h Hashtags) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}