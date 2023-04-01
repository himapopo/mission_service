// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GetItemMission is an object representing the database table.
type GetItemMission struct {
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ミッションID
	MissionID int64 `boil:"mission_id" json:"mission_id" toml:"mission_id" yaml:"mission_id"`
	// 獲得必要アイテムID
	ItemID    int64     `boil:"item_id" json:"item_id" toml:"item_id" yaml:"item_id"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *getItemMissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L getItemMissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GetItemMissionColumns = struct {
	ID        string
	MissionID string
	ItemID    string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	MissionID: "mission_id",
	ItemID:    "item_id",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

var GetItemMissionTableColumns = struct {
	ID        string
	MissionID string
	ItemID    string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "get_item_missions.id",
	MissionID: "get_item_missions.mission_id",
	ItemID:    "get_item_missions.item_id",
	UpdatedAt: "get_item_missions.updated_at",
	CreatedAt: "get_item_missions.created_at",
}

// Generated where

var GetItemMissionWhere = struct {
	ID        whereHelperint64
	MissionID whereHelperint64
	ItemID    whereHelperint64
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"get_item_missions\".\"id\""},
	MissionID: whereHelperint64{field: "\"get_item_missions\".\"mission_id\""},
	ItemID:    whereHelperint64{field: "\"get_item_missions\".\"item_id\""},
	UpdatedAt: whereHelpertime_Time{field: "\"get_item_missions\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"get_item_missions\".\"created_at\""},
}

// GetItemMissionRels is where relationship names are stored.
var GetItemMissionRels = struct {
	Item    string
	Mission string
}{
	Item:    "Item",
	Mission: "Mission",
}

// getItemMissionR is where relationships are stored.
type getItemMissionR struct {
	Item    *Item    `boil:"Item" json:"Item" toml:"Item" yaml:"Item"`
	Mission *Mission `boil:"Mission" json:"Mission" toml:"Mission" yaml:"Mission"`
}

// NewStruct creates a new relationship struct
func (*getItemMissionR) NewStruct() *getItemMissionR {
	return &getItemMissionR{}
}

func (r *getItemMissionR) GetItem() *Item {
	if r == nil {
		return nil
	}
	return r.Item
}

func (r *getItemMissionR) GetMission() *Mission {
	if r == nil {
		return nil
	}
	return r.Mission
}

// getItemMissionL is where Load methods for each relationship are stored.
type getItemMissionL struct{}

var (
	getItemMissionAllColumns            = []string{"id", "mission_id", "item_id", "updated_at", "created_at"}
	getItemMissionColumnsWithoutDefault = []string{"mission_id", "item_id"}
	getItemMissionColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	getItemMissionPrimaryKeyColumns     = []string{"id"}
	getItemMissionGeneratedColumns      = []string{}
)

type (
	// GetItemMissionSlice is an alias for a slice of pointers to GetItemMission.
	// This should almost always be used instead of []GetItemMission.
	GetItemMissionSlice []*GetItemMission
	// GetItemMissionHook is the signature for custom GetItemMission hook methods
	GetItemMissionHook func(context.Context, boil.ContextExecutor, *GetItemMission) error

	getItemMissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	getItemMissionType                 = reflect.TypeOf(&GetItemMission{})
	getItemMissionMapping              = queries.MakeStructMapping(getItemMissionType)
	getItemMissionPrimaryKeyMapping, _ = queries.BindMapping(getItemMissionType, getItemMissionMapping, getItemMissionPrimaryKeyColumns)
	getItemMissionInsertCacheMut       sync.RWMutex
	getItemMissionInsertCache          = make(map[string]insertCache)
	getItemMissionUpdateCacheMut       sync.RWMutex
	getItemMissionUpdateCache          = make(map[string]updateCache)
	getItemMissionUpsertCacheMut       sync.RWMutex
	getItemMissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var getItemMissionAfterSelectHooks []GetItemMissionHook

var getItemMissionBeforeInsertHooks []GetItemMissionHook
var getItemMissionAfterInsertHooks []GetItemMissionHook

var getItemMissionBeforeUpdateHooks []GetItemMissionHook
var getItemMissionAfterUpdateHooks []GetItemMissionHook

var getItemMissionBeforeDeleteHooks []GetItemMissionHook
var getItemMissionAfterDeleteHooks []GetItemMissionHook

var getItemMissionBeforeUpsertHooks []GetItemMissionHook
var getItemMissionAfterUpsertHooks []GetItemMissionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GetItemMission) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GetItemMission) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GetItemMission) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GetItemMission) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GetItemMission) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GetItemMission) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GetItemMission) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GetItemMission) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GetItemMission) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range getItemMissionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGetItemMissionHook registers your hook function for all future operations.
func AddGetItemMissionHook(hookPoint boil.HookPoint, getItemMissionHook GetItemMissionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		getItemMissionAfterSelectHooks = append(getItemMissionAfterSelectHooks, getItemMissionHook)
	case boil.BeforeInsertHook:
		getItemMissionBeforeInsertHooks = append(getItemMissionBeforeInsertHooks, getItemMissionHook)
	case boil.AfterInsertHook:
		getItemMissionAfterInsertHooks = append(getItemMissionAfterInsertHooks, getItemMissionHook)
	case boil.BeforeUpdateHook:
		getItemMissionBeforeUpdateHooks = append(getItemMissionBeforeUpdateHooks, getItemMissionHook)
	case boil.AfterUpdateHook:
		getItemMissionAfterUpdateHooks = append(getItemMissionAfterUpdateHooks, getItemMissionHook)
	case boil.BeforeDeleteHook:
		getItemMissionBeforeDeleteHooks = append(getItemMissionBeforeDeleteHooks, getItemMissionHook)
	case boil.AfterDeleteHook:
		getItemMissionAfterDeleteHooks = append(getItemMissionAfterDeleteHooks, getItemMissionHook)
	case boil.BeforeUpsertHook:
		getItemMissionBeforeUpsertHooks = append(getItemMissionBeforeUpsertHooks, getItemMissionHook)
	case boil.AfterUpsertHook:
		getItemMissionAfterUpsertHooks = append(getItemMissionAfterUpsertHooks, getItemMissionHook)
	}
}

// One returns a single getItemMission record from the query.
func (q getItemMissionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GetItemMission, error) {
	o := &GetItemMission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for get_item_missions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GetItemMission records from the query.
func (q getItemMissionQuery) All(ctx context.Context, exec boil.ContextExecutor) (GetItemMissionSlice, error) {
	var o []*GetItemMission

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GetItemMission slice")
	}

	if len(getItemMissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GetItemMission records in the query.
func (q getItemMissionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count get_item_missions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q getItemMissionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if get_item_missions exists")
	}

	return count > 0, nil
}

// Item pointed to by the foreign key.
func (o *GetItemMission) Item(mods ...qm.QueryMod) itemQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ItemID),
	}

	queryMods = append(queryMods, mods...)

	return Items(queryMods...)
}

// Mission pointed to by the foreign key.
func (o *GetItemMission) Mission(mods ...qm.QueryMod) missionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MissionID),
	}

	queryMods = append(queryMods, mods...)

	return Missions(queryMods...)
}

// LoadItem allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (getItemMissionL) LoadItem(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGetItemMission interface{}, mods queries.Applicator) error {
	var slice []*GetItemMission
	var object *GetItemMission

	if singular {
		var ok bool
		object, ok = maybeGetItemMission.(*GetItemMission)
		if !ok {
			object = new(GetItemMission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGetItemMission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGetItemMission))
			}
		}
	} else {
		s, ok := maybeGetItemMission.(*[]*GetItemMission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGetItemMission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGetItemMission))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &getItemMissionR{}
		}
		args = append(args, object.ItemID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &getItemMissionR{}
			}

			for _, a := range args {
				if a == obj.ItemID {
					continue Outer
				}
			}

			args = append(args, obj.ItemID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`items`),
		qm.WhereIn(`items.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Item")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Item")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if len(itemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Item = foreign
		if foreign.R == nil {
			foreign.R = &itemR{}
		}
		foreign.R.GetItemMissions = append(foreign.R.GetItemMissions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ItemID == foreign.ID {
				local.R.Item = foreign
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.GetItemMissions = append(foreign.R.GetItemMissions, local)
				break
			}
		}
	}

	return nil
}

// LoadMission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (getItemMissionL) LoadMission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGetItemMission interface{}, mods queries.Applicator) error {
	var slice []*GetItemMission
	var object *GetItemMission

	if singular {
		var ok bool
		object, ok = maybeGetItemMission.(*GetItemMission)
		if !ok {
			object = new(GetItemMission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGetItemMission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGetItemMission))
			}
		}
	} else {
		s, ok := maybeGetItemMission.(*[]*GetItemMission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGetItemMission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGetItemMission))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &getItemMissionR{}
		}
		args = append(args, object.MissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &getItemMissionR{}
			}

			for _, a := range args {
				if a == obj.MissionID {
					continue Outer
				}
			}

			args = append(args, obj.MissionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`missions`),
		qm.WhereIn(`missions.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Mission")
	}

	var resultSlice []*Mission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Mission")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for missions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for missions")
	}

	if len(missionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Mission = foreign
		if foreign.R == nil {
			foreign.R = &missionR{}
		}
		foreign.R.GetItemMission = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MissionID == foreign.ID {
				local.R.Mission = foreign
				if foreign.R == nil {
					foreign.R = &missionR{}
				}
				foreign.R.GetItemMission = local
				break
			}
		}
	}

	return nil
}

// SetItem of the getItemMission to the related item.
// Sets o.R.Item to related.
// Adds o to related.R.GetItemMissions.
func (o *GetItemMission) SetItem(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Item) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"get_item_missions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"item_id"}),
		strmangle.WhereClause("\"", "\"", 2, getItemMissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ItemID = related.ID
	if o.R == nil {
		o.R = &getItemMissionR{
			Item: related,
		}
	} else {
		o.R.Item = related
	}

	if related.R == nil {
		related.R = &itemR{
			GetItemMissions: GetItemMissionSlice{o},
		}
	} else {
		related.R.GetItemMissions = append(related.R.GetItemMissions, o)
	}

	return nil
}

// SetMission of the getItemMission to the related item.
// Sets o.R.Mission to related.
// Adds o to related.R.GetItemMission.
func (o *GetItemMission) SetMission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Mission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"get_item_missions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"mission_id"}),
		strmangle.WhereClause("\"", "\"", 2, getItemMissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MissionID = related.ID
	if o.R == nil {
		o.R = &getItemMissionR{
			Mission: related,
		}
	} else {
		o.R.Mission = related
	}

	if related.R == nil {
		related.R = &missionR{
			GetItemMission: o,
		}
	} else {
		related.R.GetItemMission = o
	}

	return nil
}

// GetItemMissions retrieves all the records using an executor.
func GetItemMissions(mods ...qm.QueryMod) getItemMissionQuery {
	mods = append(mods, qm.From("\"get_item_missions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"get_item_missions\".*"})
	}

	return getItemMissionQuery{q}
}

// FindGetItemMission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGetItemMission(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*GetItemMission, error) {
	getItemMissionObj := &GetItemMission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"get_item_missions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, getItemMissionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from get_item_missions")
	}

	if err = getItemMissionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return getItemMissionObj, err
	}

	return getItemMissionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GetItemMission) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no get_item_missions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(getItemMissionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	getItemMissionInsertCacheMut.RLock()
	cache, cached := getItemMissionInsertCache[key]
	getItemMissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			getItemMissionAllColumns,
			getItemMissionColumnsWithDefault,
			getItemMissionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(getItemMissionType, getItemMissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(getItemMissionType, getItemMissionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"get_item_missions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"get_item_missions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into get_item_missions")
	}

	if !cached {
		getItemMissionInsertCacheMut.Lock()
		getItemMissionInsertCache[key] = cache
		getItemMissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GetItemMission.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GetItemMission) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	getItemMissionUpdateCacheMut.RLock()
	cache, cached := getItemMissionUpdateCache[key]
	getItemMissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			getItemMissionAllColumns,
			getItemMissionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update get_item_missions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"get_item_missions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, getItemMissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(getItemMissionType, getItemMissionMapping, append(wl, getItemMissionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update get_item_missions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for get_item_missions")
	}

	if !cached {
		getItemMissionUpdateCacheMut.Lock()
		getItemMissionUpdateCache[key] = cache
		getItemMissionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q getItemMissionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for get_item_missions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for get_item_missions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GetItemMissionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), getItemMissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"get_item_missions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, getItemMissionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in getItemMission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all getItemMission")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GetItemMission) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no get_item_missions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(getItemMissionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	getItemMissionUpsertCacheMut.RLock()
	cache, cached := getItemMissionUpsertCache[key]
	getItemMissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			getItemMissionAllColumns,
			getItemMissionColumnsWithDefault,
			getItemMissionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			getItemMissionAllColumns,
			getItemMissionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert get_item_missions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(getItemMissionPrimaryKeyColumns))
			copy(conflict, getItemMissionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"get_item_missions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(getItemMissionType, getItemMissionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(getItemMissionType, getItemMissionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert get_item_missions")
	}

	if !cached {
		getItemMissionUpsertCacheMut.Lock()
		getItemMissionUpsertCache[key] = cache
		getItemMissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GetItemMission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GetItemMission) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GetItemMission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), getItemMissionPrimaryKeyMapping)
	sql := "DELETE FROM \"get_item_missions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from get_item_missions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for get_item_missions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q getItemMissionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no getItemMissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from get_item_missions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for get_item_missions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GetItemMissionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(getItemMissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), getItemMissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"get_item_missions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, getItemMissionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from getItemMission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for get_item_missions")
	}

	if len(getItemMissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GetItemMission) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGetItemMission(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GetItemMissionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GetItemMissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), getItemMissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"get_item_missions\".* FROM \"get_item_missions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, getItemMissionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GetItemMissionSlice")
	}

	*o = slice

	return nil
}

// GetItemMissionExists checks if the GetItemMission row exists.
func GetItemMissionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"get_item_missions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if get_item_missions exists")
	}

	return exists, nil
}

// Exists checks if the GetItemMission row exists.
func (o *GetItemMission) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GetItemMissionExists(ctx, exec, o.ID)
}
