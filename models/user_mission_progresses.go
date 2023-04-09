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

// UserMissionProgress is an object representing the database table.
type UserMissionProgress struct {
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ユーザーID
	UserMissionID int64 `boil:"user_mission_id" json:"user_mission_id" toml:"user_mission_id" yaml:"user_mission_id"`
	// 達成条件に関する現在の値
	ProgressValue int64 `boil:"progress_value" json:"progress_value" toml:"progress_value" yaml:"progress_value"`
	// 最終進捗更新日時
	LastProgressUpdatedAt time.Time `boil:"last_progress_updated_at" json:"last_progress_updated_at" toml:"last_progress_updated_at" yaml:"last_progress_updated_at"`
	UpdatedAt             time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt             time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userMissionProgressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userMissionProgressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserMissionProgressColumns = struct {
	ID                    string
	UserMissionID         string
	ProgressValue         string
	LastProgressUpdatedAt string
	UpdatedAt             string
	CreatedAt             string
}{
	ID:                    "id",
	UserMissionID:         "user_mission_id",
	ProgressValue:         "progress_value",
	LastProgressUpdatedAt: "last_progress_updated_at",
	UpdatedAt:             "updated_at",
	CreatedAt:             "created_at",
}

var UserMissionProgressTableColumns = struct {
	ID                    string
	UserMissionID         string
	ProgressValue         string
	LastProgressUpdatedAt string
	UpdatedAt             string
	CreatedAt             string
}{
	ID:                    "user_mission_progresses.id",
	UserMissionID:         "user_mission_progresses.user_mission_id",
	ProgressValue:         "user_mission_progresses.progress_value",
	LastProgressUpdatedAt: "user_mission_progresses.last_progress_updated_at",
	UpdatedAt:             "user_mission_progresses.updated_at",
	CreatedAt:             "user_mission_progresses.created_at",
}

// Generated where

var UserMissionProgressWhere = struct {
	ID                    whereHelperint64
	UserMissionID         whereHelperint64
	ProgressValue         whereHelperint64
	LastProgressUpdatedAt whereHelpertime_Time
	UpdatedAt             whereHelpertime_Time
	CreatedAt             whereHelpertime_Time
}{
	ID:                    whereHelperint64{field: "\"user_mission_progresses\".\"id\""},
	UserMissionID:         whereHelperint64{field: "\"user_mission_progresses\".\"user_mission_id\""},
	ProgressValue:         whereHelperint64{field: "\"user_mission_progresses\".\"progress_value\""},
	LastProgressUpdatedAt: whereHelpertime_Time{field: "\"user_mission_progresses\".\"last_progress_updated_at\""},
	UpdatedAt:             whereHelpertime_Time{field: "\"user_mission_progresses\".\"updated_at\""},
	CreatedAt:             whereHelpertime_Time{field: "\"user_mission_progresses\".\"created_at\""},
}

// UserMissionProgressRels is where relationship names are stored.
var UserMissionProgressRels = struct {
	UserMission string
}{
	UserMission: "UserMission",
}

// userMissionProgressR is where relationships are stored.
type userMissionProgressR struct {
	UserMission *UserMission `boil:"UserMission" json:"UserMission" toml:"UserMission" yaml:"UserMission"`
}

// NewStruct creates a new relationship struct
func (*userMissionProgressR) NewStruct() *userMissionProgressR {
	return &userMissionProgressR{}
}

func (r *userMissionProgressR) GetUserMission() *UserMission {
	if r == nil {
		return nil
	}
	return r.UserMission
}

// userMissionProgressL is where Load methods for each relationship are stored.
type userMissionProgressL struct{}

var (
	userMissionProgressAllColumns            = []string{"id", "user_mission_id", "progress_value", "last_progress_updated_at", "updated_at", "created_at"}
	userMissionProgressColumnsWithoutDefault = []string{"user_mission_id"}
	userMissionProgressColumnsWithDefault    = []string{"id", "progress_value", "last_progress_updated_at", "updated_at", "created_at"}
	userMissionProgressPrimaryKeyColumns     = []string{"id"}
	userMissionProgressGeneratedColumns      = []string{}
)

type (
	// UserMissionProgressSlice is an alias for a slice of pointers to UserMissionProgress.
	// This should almost always be used instead of []UserMissionProgress.
	UserMissionProgressSlice []*UserMissionProgress
	// UserMissionProgressHook is the signature for custom UserMissionProgress hook methods
	UserMissionProgressHook func(context.Context, boil.ContextExecutor, *UserMissionProgress) error

	userMissionProgressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userMissionProgressType                 = reflect.TypeOf(&UserMissionProgress{})
	userMissionProgressMapping              = queries.MakeStructMapping(userMissionProgressType)
	userMissionProgressPrimaryKeyMapping, _ = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, userMissionProgressPrimaryKeyColumns)
	userMissionProgressInsertCacheMut       sync.RWMutex
	userMissionProgressInsertCache          = make(map[string]insertCache)
	userMissionProgressUpdateCacheMut       sync.RWMutex
	userMissionProgressUpdateCache          = make(map[string]updateCache)
	userMissionProgressUpsertCacheMut       sync.RWMutex
	userMissionProgressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userMissionProgressAfterSelectHooks []UserMissionProgressHook

var userMissionProgressBeforeInsertHooks []UserMissionProgressHook
var userMissionProgressAfterInsertHooks []UserMissionProgressHook

var userMissionProgressBeforeUpdateHooks []UserMissionProgressHook
var userMissionProgressAfterUpdateHooks []UserMissionProgressHook

var userMissionProgressBeforeDeleteHooks []UserMissionProgressHook
var userMissionProgressAfterDeleteHooks []UserMissionProgressHook

var userMissionProgressBeforeUpsertHooks []UserMissionProgressHook
var userMissionProgressAfterUpsertHooks []UserMissionProgressHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserMissionProgress) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserMissionProgress) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserMissionProgress) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserMissionProgress) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserMissionProgress) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserMissionProgress) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserMissionProgress) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserMissionProgress) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserMissionProgress) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userMissionProgressAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserMissionProgressHook registers your hook function for all future operations.
func AddUserMissionProgressHook(hookPoint boil.HookPoint, userMissionProgressHook UserMissionProgressHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userMissionProgressAfterSelectHooks = append(userMissionProgressAfterSelectHooks, userMissionProgressHook)
	case boil.BeforeInsertHook:
		userMissionProgressBeforeInsertHooks = append(userMissionProgressBeforeInsertHooks, userMissionProgressHook)
	case boil.AfterInsertHook:
		userMissionProgressAfterInsertHooks = append(userMissionProgressAfterInsertHooks, userMissionProgressHook)
	case boil.BeforeUpdateHook:
		userMissionProgressBeforeUpdateHooks = append(userMissionProgressBeforeUpdateHooks, userMissionProgressHook)
	case boil.AfterUpdateHook:
		userMissionProgressAfterUpdateHooks = append(userMissionProgressAfterUpdateHooks, userMissionProgressHook)
	case boil.BeforeDeleteHook:
		userMissionProgressBeforeDeleteHooks = append(userMissionProgressBeforeDeleteHooks, userMissionProgressHook)
	case boil.AfterDeleteHook:
		userMissionProgressAfterDeleteHooks = append(userMissionProgressAfterDeleteHooks, userMissionProgressHook)
	case boil.BeforeUpsertHook:
		userMissionProgressBeforeUpsertHooks = append(userMissionProgressBeforeUpsertHooks, userMissionProgressHook)
	case boil.AfterUpsertHook:
		userMissionProgressAfterUpsertHooks = append(userMissionProgressAfterUpsertHooks, userMissionProgressHook)
	}
}

// One returns a single userMissionProgress record from the query.
func (q userMissionProgressQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserMissionProgress, error) {
	o := &UserMissionProgress{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_mission_progresses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserMissionProgress records from the query.
func (q userMissionProgressQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserMissionProgressSlice, error) {
	var o []*UserMissionProgress

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserMissionProgress slice")
	}

	if len(userMissionProgressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserMissionProgress records in the query.
func (q userMissionProgressQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_mission_progresses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userMissionProgressQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_mission_progresses exists")
	}

	return count > 0, nil
}

// UserMission pointed to by the foreign key.
func (o *UserMissionProgress) UserMission(mods ...qm.QueryMod) userMissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserMissionID),
	}

	queryMods = append(queryMods, mods...)

	return UserMissions(queryMods...)
}

// LoadUserMission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userMissionProgressL) LoadUserMission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserMissionProgress interface{}, mods queries.Applicator) error {
	var slice []*UserMissionProgress
	var object *UserMissionProgress

	if singular {
		var ok bool
		object, ok = maybeUserMissionProgress.(*UserMissionProgress)
		if !ok {
			object = new(UserMissionProgress)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserMissionProgress)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserMissionProgress))
			}
		}
	} else {
		s, ok := maybeUserMissionProgress.(*[]*UserMissionProgress)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserMissionProgress)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserMissionProgress))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userMissionProgressR{}
		}
		args = append(args, object.UserMissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userMissionProgressR{}
			}

			for _, a := range args {
				if a == obj.UserMissionID {
					continue Outer
				}
			}

			args = append(args, obj.UserMissionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_missions`),
		qm.WhereIn(`user_missions.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserMission")
	}

	var resultSlice []*UserMission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserMission")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_missions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_missions")
	}

	if len(userMissionAfterSelectHooks) != 0 {
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
		object.R.UserMission = foreign
		if foreign.R == nil {
			foreign.R = &userMissionR{}
		}
		foreign.R.UserMissionProgresses = append(foreign.R.UserMissionProgresses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserMissionID == foreign.ID {
				local.R.UserMission = foreign
				if foreign.R == nil {
					foreign.R = &userMissionR{}
				}
				foreign.R.UserMissionProgresses = append(foreign.R.UserMissionProgresses, local)
				break
			}
		}
	}

	return nil
}

// SetUserMission of the userMissionProgress to the related item.
// Sets o.R.UserMission to related.
// Adds o to related.R.UserMissionProgresses.
func (o *UserMissionProgress) SetUserMission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserMission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_mission_progresses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_mission_id"}),
		strmangle.WhereClause("\"", "\"", 2, userMissionProgressPrimaryKeyColumns),
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

	o.UserMissionID = related.ID
	if o.R == nil {
		o.R = &userMissionProgressR{
			UserMission: related,
		}
	} else {
		o.R.UserMission = related
	}

	if related.R == nil {
		related.R = &userMissionR{
			UserMissionProgresses: UserMissionProgressSlice{o},
		}
	} else {
		related.R.UserMissionProgresses = append(related.R.UserMissionProgresses, o)
	}

	return nil
}

// UserMissionProgresses retrieves all the records using an executor.
func UserMissionProgresses(mods ...qm.QueryMod) userMissionProgressQuery {
	mods = append(mods, qm.From("\"user_mission_progresses\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_mission_progresses\".*"})
	}

	return userMissionProgressQuery{q}
}

// FindUserMissionProgress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserMissionProgress(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*UserMissionProgress, error) {
	userMissionProgressObj := &UserMissionProgress{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_mission_progresses\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userMissionProgressObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_mission_progresses")
	}

	if err = userMissionProgressObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userMissionProgressObj, err
	}

	return userMissionProgressObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserMissionProgress) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_mission_progresses provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userMissionProgressColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userMissionProgressInsertCacheMut.RLock()
	cache, cached := userMissionProgressInsertCache[key]
	userMissionProgressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userMissionProgressAllColumns,
			userMissionProgressColumnsWithDefault,
			userMissionProgressColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_mission_progresses\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_mission_progresses\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_mission_progresses")
	}

	if !cached {
		userMissionProgressInsertCacheMut.Lock()
		userMissionProgressInsertCache[key] = cache
		userMissionProgressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserMissionProgress.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserMissionProgress) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userMissionProgressUpdateCacheMut.RLock()
	cache, cached := userMissionProgressUpdateCache[key]
	userMissionProgressUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userMissionProgressAllColumns,
			userMissionProgressPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_mission_progresses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_mission_progresses\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userMissionProgressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, append(wl, userMissionProgressPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_mission_progresses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_mission_progresses")
	}

	if !cached {
		userMissionProgressUpdateCacheMut.Lock()
		userMissionProgressUpdateCache[key] = cache
		userMissionProgressUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userMissionProgressQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_mission_progresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_mission_progresses")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserMissionProgressSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userMissionProgressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_mission_progresses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userMissionProgressPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userMissionProgress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userMissionProgress")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserMissionProgress) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_mission_progresses provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(userMissionProgressColumnsWithDefault, o)

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

	userMissionProgressUpsertCacheMut.RLock()
	cache, cached := userMissionProgressUpsertCache[key]
	userMissionProgressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userMissionProgressAllColumns,
			userMissionProgressColumnsWithDefault,
			userMissionProgressColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userMissionProgressAllColumns,
			userMissionProgressPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_mission_progresses, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userMissionProgressPrimaryKeyColumns))
			copy(conflict, userMissionProgressPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_mission_progresses\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userMissionProgressType, userMissionProgressMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_mission_progresses")
	}

	if !cached {
		userMissionProgressUpsertCacheMut.Lock()
		userMissionProgressUpsertCache[key] = cache
		userMissionProgressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserMissionProgress record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserMissionProgress) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserMissionProgress provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userMissionProgressPrimaryKeyMapping)
	sql := "DELETE FROM \"user_mission_progresses\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_mission_progresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_mission_progresses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userMissionProgressQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userMissionProgressQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_mission_progresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_mission_progresses")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserMissionProgressSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userMissionProgressBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userMissionProgressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_mission_progresses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userMissionProgressPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userMissionProgress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_mission_progresses")
	}

	if len(userMissionProgressAfterDeleteHooks) != 0 {
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
func (o *UserMissionProgress) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserMissionProgress(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserMissionProgressSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserMissionProgressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userMissionProgressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_mission_progresses\".* FROM \"user_mission_progresses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userMissionProgressPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserMissionProgressSlice")
	}

	*o = slice

	return nil
}

// UserMissionProgressExists checks if the UserMissionProgress row exists.
func UserMissionProgressExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_mission_progresses\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_mission_progresses exists")
	}

	return exists, nil
}

// Exists checks if the UserMissionProgress row exists.
func (o *UserMissionProgress) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserMissionProgressExists(ctx, exec, o.ID)
}
