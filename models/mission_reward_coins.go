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

// MissionRewardCoin is an object representing the database table.
type MissionRewardCoin struct {
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ミッションID
	MissionID int64 `boil:"mission_id" json:"mission_id" toml:"mission_id" yaml:"mission_id"`
	// 付与するコイン数
	CoinCount int64     `boil:"coin_count" json:"coin_count" toml:"coin_count" yaml:"coin_count"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *missionRewardCoinR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L missionRewardCoinL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MissionRewardCoinColumns = struct {
	ID        string
	MissionID string
	CoinCount string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	MissionID: "mission_id",
	CoinCount: "coin_count",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

var MissionRewardCoinTableColumns = struct {
	ID        string
	MissionID string
	CoinCount string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "mission_reward_coins.id",
	MissionID: "mission_reward_coins.mission_id",
	CoinCount: "mission_reward_coins.coin_count",
	UpdatedAt: "mission_reward_coins.updated_at",
	CreatedAt: "mission_reward_coins.created_at",
}

// Generated where

var MissionRewardCoinWhere = struct {
	ID        whereHelperint64
	MissionID whereHelperint64
	CoinCount whereHelperint64
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"mission_reward_coins\".\"id\""},
	MissionID: whereHelperint64{field: "\"mission_reward_coins\".\"mission_id\""},
	CoinCount: whereHelperint64{field: "\"mission_reward_coins\".\"coin_count\""},
	UpdatedAt: whereHelpertime_Time{field: "\"mission_reward_coins\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"mission_reward_coins\".\"created_at\""},
}

// MissionRewardCoinRels is where relationship names are stored.
var MissionRewardCoinRels = struct {
	Mission string
}{
	Mission: "Mission",
}

// missionRewardCoinR is where relationships are stored.
type missionRewardCoinR struct {
	Mission *Mission `boil:"Mission" json:"Mission" toml:"Mission" yaml:"Mission"`
}

// NewStruct creates a new relationship struct
func (*missionRewardCoinR) NewStruct() *missionRewardCoinR {
	return &missionRewardCoinR{}
}

func (r *missionRewardCoinR) GetMission() *Mission {
	if r == nil {
		return nil
	}
	return r.Mission
}

// missionRewardCoinL is where Load methods for each relationship are stored.
type missionRewardCoinL struct{}

var (
	missionRewardCoinAllColumns            = []string{"id", "mission_id", "coin_count", "updated_at", "created_at"}
	missionRewardCoinColumnsWithoutDefault = []string{"mission_id", "coin_count"}
	missionRewardCoinColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	missionRewardCoinPrimaryKeyColumns     = []string{"id"}
	missionRewardCoinGeneratedColumns      = []string{}
)

type (
	// MissionRewardCoinSlice is an alias for a slice of pointers to MissionRewardCoin.
	// This should almost always be used instead of []MissionRewardCoin.
	MissionRewardCoinSlice []*MissionRewardCoin
	// MissionRewardCoinHook is the signature for custom MissionRewardCoin hook methods
	MissionRewardCoinHook func(context.Context, boil.ContextExecutor, *MissionRewardCoin) error

	missionRewardCoinQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	missionRewardCoinType                 = reflect.TypeOf(&MissionRewardCoin{})
	missionRewardCoinMapping              = queries.MakeStructMapping(missionRewardCoinType)
	missionRewardCoinPrimaryKeyMapping, _ = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, missionRewardCoinPrimaryKeyColumns)
	missionRewardCoinInsertCacheMut       sync.RWMutex
	missionRewardCoinInsertCache          = make(map[string]insertCache)
	missionRewardCoinUpdateCacheMut       sync.RWMutex
	missionRewardCoinUpdateCache          = make(map[string]updateCache)
	missionRewardCoinUpsertCacheMut       sync.RWMutex
	missionRewardCoinUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var missionRewardCoinAfterSelectHooks []MissionRewardCoinHook

var missionRewardCoinBeforeInsertHooks []MissionRewardCoinHook
var missionRewardCoinAfterInsertHooks []MissionRewardCoinHook

var missionRewardCoinBeforeUpdateHooks []MissionRewardCoinHook
var missionRewardCoinAfterUpdateHooks []MissionRewardCoinHook

var missionRewardCoinBeforeDeleteHooks []MissionRewardCoinHook
var missionRewardCoinAfterDeleteHooks []MissionRewardCoinHook

var missionRewardCoinBeforeUpsertHooks []MissionRewardCoinHook
var missionRewardCoinAfterUpsertHooks []MissionRewardCoinHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MissionRewardCoin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MissionRewardCoin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MissionRewardCoin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MissionRewardCoin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MissionRewardCoin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MissionRewardCoin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MissionRewardCoin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MissionRewardCoin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MissionRewardCoin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionRewardCoinAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMissionRewardCoinHook registers your hook function for all future operations.
func AddMissionRewardCoinHook(hookPoint boil.HookPoint, missionRewardCoinHook MissionRewardCoinHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		missionRewardCoinAfterSelectHooks = append(missionRewardCoinAfterSelectHooks, missionRewardCoinHook)
	case boil.BeforeInsertHook:
		missionRewardCoinBeforeInsertHooks = append(missionRewardCoinBeforeInsertHooks, missionRewardCoinHook)
	case boil.AfterInsertHook:
		missionRewardCoinAfterInsertHooks = append(missionRewardCoinAfterInsertHooks, missionRewardCoinHook)
	case boil.BeforeUpdateHook:
		missionRewardCoinBeforeUpdateHooks = append(missionRewardCoinBeforeUpdateHooks, missionRewardCoinHook)
	case boil.AfterUpdateHook:
		missionRewardCoinAfterUpdateHooks = append(missionRewardCoinAfterUpdateHooks, missionRewardCoinHook)
	case boil.BeforeDeleteHook:
		missionRewardCoinBeforeDeleteHooks = append(missionRewardCoinBeforeDeleteHooks, missionRewardCoinHook)
	case boil.AfterDeleteHook:
		missionRewardCoinAfterDeleteHooks = append(missionRewardCoinAfterDeleteHooks, missionRewardCoinHook)
	case boil.BeforeUpsertHook:
		missionRewardCoinBeforeUpsertHooks = append(missionRewardCoinBeforeUpsertHooks, missionRewardCoinHook)
	case boil.AfterUpsertHook:
		missionRewardCoinAfterUpsertHooks = append(missionRewardCoinAfterUpsertHooks, missionRewardCoinHook)
	}
}

// One returns a single missionRewardCoin record from the query.
func (q missionRewardCoinQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MissionRewardCoin, error) {
	o := &MissionRewardCoin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mission_reward_coins")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MissionRewardCoin records from the query.
func (q missionRewardCoinQuery) All(ctx context.Context, exec boil.ContextExecutor) (MissionRewardCoinSlice, error) {
	var o []*MissionRewardCoin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MissionRewardCoin slice")
	}

	if len(missionRewardCoinAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MissionRewardCoin records in the query.
func (q missionRewardCoinQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mission_reward_coins rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q missionRewardCoinQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mission_reward_coins exists")
	}

	return count > 0, nil
}

// Mission pointed to by the foreign key.
func (o *MissionRewardCoin) Mission(mods ...qm.QueryMod) missionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MissionID),
	}

	queryMods = append(queryMods, mods...)

	return Missions(queryMods...)
}

// LoadMission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (missionRewardCoinL) LoadMission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMissionRewardCoin interface{}, mods queries.Applicator) error {
	var slice []*MissionRewardCoin
	var object *MissionRewardCoin

	if singular {
		var ok bool
		object, ok = maybeMissionRewardCoin.(*MissionRewardCoin)
		if !ok {
			object = new(MissionRewardCoin)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMissionRewardCoin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMissionRewardCoin))
			}
		}
	} else {
		s, ok := maybeMissionRewardCoin.(*[]*MissionRewardCoin)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMissionRewardCoin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMissionRewardCoin))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &missionRewardCoinR{}
		}
		args = append(args, object.MissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &missionRewardCoinR{}
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
		foreign.R.MissionRewardCoins = append(foreign.R.MissionRewardCoins, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MissionID == foreign.ID {
				local.R.Mission = foreign
				if foreign.R == nil {
					foreign.R = &missionR{}
				}
				foreign.R.MissionRewardCoins = append(foreign.R.MissionRewardCoins, local)
				break
			}
		}
	}

	return nil
}

// SetMission of the missionRewardCoin to the related item.
// Sets o.R.Mission to related.
// Adds o to related.R.MissionRewardCoins.
func (o *MissionRewardCoin) SetMission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Mission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mission_reward_coins\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"mission_id"}),
		strmangle.WhereClause("\"", "\"", 2, missionRewardCoinPrimaryKeyColumns),
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
		o.R = &missionRewardCoinR{
			Mission: related,
		}
	} else {
		o.R.Mission = related
	}

	if related.R == nil {
		related.R = &missionR{
			MissionRewardCoins: MissionRewardCoinSlice{o},
		}
	} else {
		related.R.MissionRewardCoins = append(related.R.MissionRewardCoins, o)
	}

	return nil
}

// MissionRewardCoins retrieves all the records using an executor.
func MissionRewardCoins(mods ...qm.QueryMod) missionRewardCoinQuery {
	mods = append(mods, qm.From("\"mission_reward_coins\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mission_reward_coins\".*"})
	}

	return missionRewardCoinQuery{q}
}

// FindMissionRewardCoin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMissionRewardCoin(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MissionRewardCoin, error) {
	missionRewardCoinObj := &MissionRewardCoin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mission_reward_coins\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, missionRewardCoinObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mission_reward_coins")
	}

	if err = missionRewardCoinObj.doAfterSelectHooks(ctx, exec); err != nil {
		return missionRewardCoinObj, err
	}

	return missionRewardCoinObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MissionRewardCoin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mission_reward_coins provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(missionRewardCoinColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	missionRewardCoinInsertCacheMut.RLock()
	cache, cached := missionRewardCoinInsertCache[key]
	missionRewardCoinInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			missionRewardCoinAllColumns,
			missionRewardCoinColumnsWithDefault,
			missionRewardCoinColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mission_reward_coins\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mission_reward_coins\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mission_reward_coins")
	}

	if !cached {
		missionRewardCoinInsertCacheMut.Lock()
		missionRewardCoinInsertCache[key] = cache
		missionRewardCoinInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MissionRewardCoin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MissionRewardCoin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	missionRewardCoinUpdateCacheMut.RLock()
	cache, cached := missionRewardCoinUpdateCache[key]
	missionRewardCoinUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			missionRewardCoinAllColumns,
			missionRewardCoinPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mission_reward_coins, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mission_reward_coins\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, missionRewardCoinPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, append(wl, missionRewardCoinPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mission_reward_coins row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mission_reward_coins")
	}

	if !cached {
		missionRewardCoinUpdateCacheMut.Lock()
		missionRewardCoinUpdateCache[key] = cache
		missionRewardCoinUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q missionRewardCoinQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mission_reward_coins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mission_reward_coins")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MissionRewardCoinSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionRewardCoinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mission_reward_coins\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, missionRewardCoinPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in missionRewardCoin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all missionRewardCoin")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MissionRewardCoin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mission_reward_coins provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(missionRewardCoinColumnsWithDefault, o)

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

	missionRewardCoinUpsertCacheMut.RLock()
	cache, cached := missionRewardCoinUpsertCache[key]
	missionRewardCoinUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			missionRewardCoinAllColumns,
			missionRewardCoinColumnsWithDefault,
			missionRewardCoinColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			missionRewardCoinAllColumns,
			missionRewardCoinPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mission_reward_coins, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(missionRewardCoinPrimaryKeyColumns))
			copy(conflict, missionRewardCoinPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mission_reward_coins\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(missionRewardCoinType, missionRewardCoinMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mission_reward_coins")
	}

	if !cached {
		missionRewardCoinUpsertCacheMut.Lock()
		missionRewardCoinUpsertCache[key] = cache
		missionRewardCoinUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MissionRewardCoin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MissionRewardCoin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MissionRewardCoin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), missionRewardCoinPrimaryKeyMapping)
	sql := "DELETE FROM \"mission_reward_coins\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mission_reward_coins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mission_reward_coins")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q missionRewardCoinQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no missionRewardCoinQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mission_reward_coins")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mission_reward_coins")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MissionRewardCoinSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(missionRewardCoinBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionRewardCoinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mission_reward_coins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, missionRewardCoinPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from missionRewardCoin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mission_reward_coins")
	}

	if len(missionRewardCoinAfterDeleteHooks) != 0 {
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
func (o *MissionRewardCoin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMissionRewardCoin(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MissionRewardCoinSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MissionRewardCoinSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionRewardCoinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mission_reward_coins\".* FROM \"mission_reward_coins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, missionRewardCoinPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MissionRewardCoinSlice")
	}

	*o = slice

	return nil
}

// MissionRewardCoinExists checks if the MissionRewardCoin row exists.
func MissionRewardCoinExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mission_reward_coins\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mission_reward_coins exists")
	}

	return exists, nil
}

// Exists checks if the MissionRewardCoin row exists.
func (o *MissionRewardCoin) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MissionRewardCoinExists(ctx, exec, o.ID)
}
