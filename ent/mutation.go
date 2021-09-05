// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/turnon/nervous/ent/event"
	"github.com/turnon/nervous/ent/predicate"
	"github.com/turnon/nervous/ent/tag"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEvent = "Event"
	TypeTag   = "Tag"
)

// EventMutation represents an operation that mutates the Event nodes in the graph.
type EventMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	start_at      *time.Time
	end_at        *time.Time
	clearedFields map[string]struct{}
	tags          map[int]struct{}
	removedtags   map[int]struct{}
	clearedtags   bool
	done          bool
	oldValue      func(context.Context) (*Event, error)
	predicates    []predicate.Event
}

var _ ent.Mutation = (*EventMutation)(nil)

// eventOption allows management of the mutation configuration using functional options.
type eventOption func(*EventMutation)

// newEventMutation creates new mutation for the Event entity.
func newEventMutation(c config, op Op, opts ...eventOption) *EventMutation {
	m := &EventMutation{
		config:        c,
		op:            op,
		typ:           TypeEvent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEventID sets the ID field of the mutation.
func withEventID(id int) eventOption {
	return func(m *EventMutation) {
		var (
			err   error
			once  sync.Once
			value *Event
		)
		m.oldValue = func(ctx context.Context) (*Event, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Event.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEvent sets the old Event of the mutation.
func withEvent(node *Event) eventOption {
	return func(m *EventMutation) {
		m.oldValue = func(context.Context) (*Event, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EventMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EventMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EventMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *EventMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EventMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Event entity.
// If the Event object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EventMutation) ResetName() {
	m.name = nil
}

// SetStartAt sets the "start_at" field.
func (m *EventMutation) SetStartAt(t time.Time) {
	m.start_at = &t
}

// StartAt returns the value of the "start_at" field in the mutation.
func (m *EventMutation) StartAt() (r time.Time, exists bool) {
	v := m.start_at
	if v == nil {
		return
	}
	return *v, true
}

// OldStartAt returns the old "start_at" field's value of the Event entity.
// If the Event object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventMutation) OldStartAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStartAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStartAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartAt: %w", err)
	}
	return oldValue.StartAt, nil
}

// ResetStartAt resets all changes to the "start_at" field.
func (m *EventMutation) ResetStartAt() {
	m.start_at = nil
}

// SetEndAt sets the "end_at" field.
func (m *EventMutation) SetEndAt(t time.Time) {
	m.end_at = &t
}

// EndAt returns the value of the "end_at" field in the mutation.
func (m *EventMutation) EndAt() (r time.Time, exists bool) {
	v := m.end_at
	if v == nil {
		return
	}
	return *v, true
}

// OldEndAt returns the old "end_at" field's value of the Event entity.
// If the Event object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventMutation) OldEndAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEndAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEndAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEndAt: %w", err)
	}
	return oldValue.EndAt, nil
}

// ResetEndAt resets all changes to the "end_at" field.
func (m *EventMutation) ResetEndAt() {
	m.end_at = nil
}

// AddTagIDs adds the "tags" edge to the Tag entity by ids.
func (m *EventMutation) AddTagIDs(ids ...int) {
	if m.tags == nil {
		m.tags = make(map[int]struct{})
	}
	for i := range ids {
		m.tags[ids[i]] = struct{}{}
	}
}

// ClearTags clears the "tags" edge to the Tag entity.
func (m *EventMutation) ClearTags() {
	m.clearedtags = true
}

// TagsCleared reports if the "tags" edge to the Tag entity was cleared.
func (m *EventMutation) TagsCleared() bool {
	return m.clearedtags
}

// RemoveTagIDs removes the "tags" edge to the Tag entity by IDs.
func (m *EventMutation) RemoveTagIDs(ids ...int) {
	if m.removedtags == nil {
		m.removedtags = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.tags, ids[i])
		m.removedtags[ids[i]] = struct{}{}
	}
}

// RemovedTags returns the removed IDs of the "tags" edge to the Tag entity.
func (m *EventMutation) RemovedTagsIDs() (ids []int) {
	for id := range m.removedtags {
		ids = append(ids, id)
	}
	return
}

// TagsIDs returns the "tags" edge IDs in the mutation.
func (m *EventMutation) TagsIDs() (ids []int) {
	for id := range m.tags {
		ids = append(ids, id)
	}
	return
}

// ResetTags resets all changes to the "tags" edge.
func (m *EventMutation) ResetTags() {
	m.tags = nil
	m.clearedtags = false
	m.removedtags = nil
}

// Where appends a list predicates to the EventMutation builder.
func (m *EventMutation) Where(ps ...predicate.Event) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EventMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Event).
func (m *EventMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EventMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, event.FieldName)
	}
	if m.start_at != nil {
		fields = append(fields, event.FieldStartAt)
	}
	if m.end_at != nil {
		fields = append(fields, event.FieldEndAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EventMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case event.FieldName:
		return m.Name()
	case event.FieldStartAt:
		return m.StartAt()
	case event.FieldEndAt:
		return m.EndAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EventMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case event.FieldName:
		return m.OldName(ctx)
	case event.FieldStartAt:
		return m.OldStartAt(ctx)
	case event.FieldEndAt:
		return m.OldEndAt(ctx)
	}
	return nil, fmt.Errorf("unknown Event field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EventMutation) SetField(name string, value ent.Value) error {
	switch name {
	case event.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case event.FieldStartAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartAt(v)
		return nil
	case event.FieldEndAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEndAt(v)
		return nil
	}
	return fmt.Errorf("unknown Event field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EventMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EventMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EventMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Event numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EventMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EventMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EventMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Event nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EventMutation) ResetField(name string) error {
	switch name {
	case event.FieldName:
		m.ResetName()
		return nil
	case event.FieldStartAt:
		m.ResetStartAt()
		return nil
	case event.FieldEndAt:
		m.ResetEndAt()
		return nil
	}
	return fmt.Errorf("unknown Event field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EventMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tags != nil {
		edges = append(edges, event.EdgeTags)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EventMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case event.EdgeTags:
		ids := make([]ent.Value, 0, len(m.tags))
		for id := range m.tags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EventMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtags != nil {
		edges = append(edges, event.EdgeTags)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EventMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case event.EdgeTags:
		ids := make([]ent.Value, 0, len(m.removedtags))
		for id := range m.removedtags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EventMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtags {
		edges = append(edges, event.EdgeTags)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EventMutation) EdgeCleared(name string) bool {
	switch name {
	case event.EdgeTags:
		return m.clearedtags
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EventMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Event unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EventMutation) ResetEdge(name string) error {
	switch name {
	case event.EdgeTags:
		m.ResetTags()
		return nil
	}
	return fmt.Errorf("unknown Event edge %s", name)
}

// TagMutation represents an operation that mutates the Tag nodes in the graph.
type TagMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	events        map[int]struct{}
	removedevents map[int]struct{}
	clearedevents bool
	done          bool
	oldValue      func(context.Context) (*Tag, error)
	predicates    []predicate.Tag
}

var _ ent.Mutation = (*TagMutation)(nil)

// tagOption allows management of the mutation configuration using functional options.
type tagOption func(*TagMutation)

// newTagMutation creates new mutation for the Tag entity.
func newTagMutation(c config, op Op, opts ...tagOption) *TagMutation {
	m := &TagMutation{
		config:        c,
		op:            op,
		typ:           TypeTag,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTagID sets the ID field of the mutation.
func withTagID(id int) tagOption {
	return func(m *TagMutation) {
		var (
			err   error
			once  sync.Once
			value *Tag
		)
		m.oldValue = func(ctx context.Context) (*Tag, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tag.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTag sets the old Tag of the mutation.
func withTag(node *Tag) tagOption {
	return func(m *TagMutation) {
		m.oldValue = func(context.Context) (*Tag, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TagMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TagMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TagMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *TagMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TagMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Tag entity.
// If the Tag object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TagMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TagMutation) ResetName() {
	m.name = nil
}

// AddEventIDs adds the "events" edge to the Event entity by ids.
func (m *TagMutation) AddEventIDs(ids ...int) {
	if m.events == nil {
		m.events = make(map[int]struct{})
	}
	for i := range ids {
		m.events[ids[i]] = struct{}{}
	}
}

// ClearEvents clears the "events" edge to the Event entity.
func (m *TagMutation) ClearEvents() {
	m.clearedevents = true
}

// EventsCleared reports if the "events" edge to the Event entity was cleared.
func (m *TagMutation) EventsCleared() bool {
	return m.clearedevents
}

// RemoveEventIDs removes the "events" edge to the Event entity by IDs.
func (m *TagMutation) RemoveEventIDs(ids ...int) {
	if m.removedevents == nil {
		m.removedevents = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.events, ids[i])
		m.removedevents[ids[i]] = struct{}{}
	}
}

// RemovedEvents returns the removed IDs of the "events" edge to the Event entity.
func (m *TagMutation) RemovedEventsIDs() (ids []int) {
	for id := range m.removedevents {
		ids = append(ids, id)
	}
	return
}

// EventsIDs returns the "events" edge IDs in the mutation.
func (m *TagMutation) EventsIDs() (ids []int) {
	for id := range m.events {
		ids = append(ids, id)
	}
	return
}

// ResetEvents resets all changes to the "events" edge.
func (m *TagMutation) ResetEvents() {
	m.events = nil
	m.clearedevents = false
	m.removedevents = nil
}

// Where appends a list predicates to the TagMutation builder.
func (m *TagMutation) Where(ps ...predicate.Tag) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TagMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Tag).
func (m *TagMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TagMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, tag.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TagMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tag.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TagMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tag.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Tag field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TagMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tag.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TagMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TagMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TagMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TagMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TagMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TagMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tag nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TagMutation) ResetField(name string) error {
	switch name {
	case tag.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TagMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.events != nil {
		edges = append(edges, tag.EdgeEvents)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TagMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgeEvents:
		ids := make([]ent.Value, 0, len(m.events))
		for id := range m.events {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TagMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedevents != nil {
		edges = append(edges, tag.EdgeEvents)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TagMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgeEvents:
		ids := make([]ent.Value, 0, len(m.removedevents))
		for id := range m.removedevents {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TagMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedevents {
		edges = append(edges, tag.EdgeEvents)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TagMutation) EdgeCleared(name string) bool {
	switch name {
	case tag.EdgeEvents:
		return m.clearedevents
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TagMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TagMutation) ResetEdge(name string) error {
	switch name {
	case tag.EdgeEvents:
		m.ResetEvents()
		return nil
	}
	return fmt.Errorf("unknown Tag edge %s", name)
}
