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

// MissionRelease is an object representing the database table.
type MissionRelease struct {
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// 達成必要なミッションID
	CompleteMissionID int64 `boil:"complete_mission_id" json:"complete_mission_id" toml:"complete_mission_id" yaml:"complete_mission_id"`
	// 解放されるミッションID
	ReleaseMissionID int64     `boil:"release_mission_id" json:"release_mission_id" toml:"release_mission_id" yaml:"release_mission_id"`
	UpdatedAt        time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *missionReleaseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L missionReleaseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MissionReleaseColumns = struct {
	ID                string
	CompleteMissionID string
	ReleaseMissionID  string
	UpdatedAt         string
	CreatedAt         string
}{
	ID:                "id",
	CompleteMissionID: "complete_mission_id",
	ReleaseMissionID:  "release_mission_id",
	UpdatedAt:         "updated_at",
	CreatedAt:         "created_at",
}

var MissionReleaseTableColumns = struct {
	ID                string
	CompleteMissionID string
	ReleaseMissionID  string
	UpdatedAt         string
	CreatedAt         string
}{
	ID:                "mission_releases.id",
	CompleteMissionID: "mission_releases.complete_mission_id",
	ReleaseMissionID:  "mission_releases.release_mission_id",
	UpdatedAt:         "mission_releases.updated_at",
	CreatedAt:         "mission_releases.created_at",
}

// Generated where

var MissionReleaseWhere = struct {
	ID                whereHelperint64
	CompleteMissionID whereHelperint64
	ReleaseMissionID  whereHelperint64
	UpdatedAt         whereHelpertime_Time
	CreatedAt         whereHelpertime_Time
}{
	ID:                whereHelperint64{field: "\"mission_releases\".\"id\""},
	CompleteMissionID: whereHelperint64{field: "\"mission_releases\".\"complete_mission_id\""},
	ReleaseMissionID:  whereHelperint64{field: "\"mission_releases\".\"release_mission_id\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"mission_releases\".\"updated_at\""},
	CreatedAt:         whereHelpertime_Time{field: "\"mission_releases\".\"created_at\""},
}

// MissionReleaseRels is where relationship names are stored.
var MissionReleaseRels = struct {
	CompleteMission string
	ReleaseMission  string
}{
	CompleteMission: "CompleteMission",
	ReleaseMission:  "ReleaseMission",
}

// missionReleaseR is where relationships are stored.
type missionReleaseR struct {
	CompleteMission *Mission `boil:"CompleteMission" json:"CompleteMission" toml:"CompleteMission" yaml:"CompleteMission"`
	ReleaseMission  *Mission `boil:"ReleaseMission" json:"ReleaseMission" toml:"ReleaseMission" yaml:"ReleaseMission"`
}

// NewStruct creates a new relationship struct
func (*missionReleaseR) NewStruct() *missionReleaseR {
	return &missionReleaseR{}
}

func (r *missionReleaseR) GetCompleteMission() *Mission {
	if r == nil {
		return nil
	}
	return r.CompleteMission
}

func (r *missionReleaseR) GetReleaseMission() *Mission {
	if r == nil {
		return nil
	}
	return r.ReleaseMission
}

// missionReleaseL is where Load methods for each relationship are stored.
type missionReleaseL struct{}

var (
	missionReleaseAllColumns            = []string{"id", "complete_mission_id", "release_mission_id", "updated_at", "created_at"}
	missionReleaseColumnsWithoutDefault = []string{"complete_mission_id", "release_mission_id"}
	missionReleaseColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	missionReleasePrimaryKeyColumns     = []string{"id"}
	missionReleaseGeneratedColumns      = []string{}
)

type (
	// MissionReleaseSlice is an alias for a slice of pointers to MissionRelease.
	// This should almost always be used instead of []MissionRelease.
	MissionReleaseSlice []*MissionRelease
	// MissionReleaseHook is the signature for custom MissionRelease hook methods
	MissionReleaseHook func(context.Context, boil.ContextExecutor, *MissionRelease) error

	missionReleaseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	missionReleaseType                 = reflect.TypeOf(&MissionRelease{})
	missionReleaseMapping              = queries.MakeStructMapping(missionReleaseType)
	missionReleasePrimaryKeyMapping, _ = queries.BindMapping(missionReleaseType, missionReleaseMapping, missionReleasePrimaryKeyColumns)
	missionReleaseInsertCacheMut       sync.RWMutex
	missionReleaseInsertCache          = make(map[string]insertCache)
	missionReleaseUpdateCacheMut       sync.RWMutex
	missionReleaseUpdateCache          = make(map[string]updateCache)
	missionReleaseUpsertCacheMut       sync.RWMutex
	missionReleaseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var missionReleaseAfterSelectHooks []MissionReleaseHook

var missionReleaseBeforeInsertHooks []MissionReleaseHook
var missionReleaseAfterInsertHooks []MissionReleaseHook

var missionReleaseBeforeUpdateHooks []MissionReleaseHook
var missionReleaseAfterUpdateHooks []MissionReleaseHook

var missionReleaseBeforeDeleteHooks []MissionReleaseHook
var missionReleaseAfterDeleteHooks []MissionReleaseHook

var missionReleaseBeforeUpsertHooks []MissionReleaseHook
var missionReleaseAfterUpsertHooks []MissionReleaseHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MissionRelease) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MissionRelease) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MissionRelease) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MissionRelease) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MissionRelease) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MissionRelease) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MissionRelease) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MissionRelease) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MissionRelease) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range missionReleaseAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMissionReleaseHook registers your hook function for all future operations.
func AddMissionReleaseHook(hookPoint boil.HookPoint, missionReleaseHook MissionReleaseHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		missionReleaseAfterSelectHooks = append(missionReleaseAfterSelectHooks, missionReleaseHook)
	case boil.BeforeInsertHook:
		missionReleaseBeforeInsertHooks = append(missionReleaseBeforeInsertHooks, missionReleaseHook)
	case boil.AfterInsertHook:
		missionReleaseAfterInsertHooks = append(missionReleaseAfterInsertHooks, missionReleaseHook)
	case boil.BeforeUpdateHook:
		missionReleaseBeforeUpdateHooks = append(missionReleaseBeforeUpdateHooks, missionReleaseHook)
	case boil.AfterUpdateHook:
		missionReleaseAfterUpdateHooks = append(missionReleaseAfterUpdateHooks, missionReleaseHook)
	case boil.BeforeDeleteHook:
		missionReleaseBeforeDeleteHooks = append(missionReleaseBeforeDeleteHooks, missionReleaseHook)
	case boil.AfterDeleteHook:
		missionReleaseAfterDeleteHooks = append(missionReleaseAfterDeleteHooks, missionReleaseHook)
	case boil.BeforeUpsertHook:
		missionReleaseBeforeUpsertHooks = append(missionReleaseBeforeUpsertHooks, missionReleaseHook)
	case boil.AfterUpsertHook:
		missionReleaseAfterUpsertHooks = append(missionReleaseAfterUpsertHooks, missionReleaseHook)
	}
}

// One returns a single missionRelease record from the query.
func (q missionReleaseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MissionRelease, error) {
	o := &MissionRelease{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mission_releases")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MissionRelease records from the query.
func (q missionReleaseQuery) All(ctx context.Context, exec boil.ContextExecutor) (MissionReleaseSlice, error) {
	var o []*MissionRelease

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MissionRelease slice")
	}

	if len(missionReleaseAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MissionRelease records in the query.
func (q missionReleaseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mission_releases rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q missionReleaseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mission_releases exists")
	}

	return count > 0, nil
}

// CompleteMission pointed to by the foreign key.
func (o *MissionRelease) CompleteMission(mods ...qm.QueryMod) missionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CompleteMissionID),
	}

	queryMods = append(queryMods, mods...)

	return Missions(queryMods...)
}

// ReleaseMission pointed to by the foreign key.
func (o *MissionRelease) ReleaseMission(mods ...qm.QueryMod) missionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ReleaseMissionID),
	}

	queryMods = append(queryMods, mods...)

	return Missions(queryMods...)
}

// LoadCompleteMission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (missionReleaseL) LoadCompleteMission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMissionRelease interface{}, mods queries.Applicator) error {
	var slice []*MissionRelease
	var object *MissionRelease

	if singular {
		var ok bool
		object, ok = maybeMissionRelease.(*MissionRelease)
		if !ok {
			object = new(MissionRelease)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMissionRelease)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMissionRelease))
			}
		}
	} else {
		s, ok := maybeMissionRelease.(*[]*MissionRelease)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMissionRelease)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMissionRelease))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &missionReleaseR{}
		}
		args = append(args, object.CompleteMissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &missionReleaseR{}
			}

			for _, a := range args {
				if a == obj.CompleteMissionID {
					continue Outer
				}
			}

			args = append(args, obj.CompleteMissionID)

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
		object.R.CompleteMission = foreign
		if foreign.R == nil {
			foreign.R = &missionR{}
		}
		foreign.R.CompleteMissionMissionReleases = append(foreign.R.CompleteMissionMissionReleases, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CompleteMissionID == foreign.ID {
				local.R.CompleteMission = foreign
				if foreign.R == nil {
					foreign.R = &missionR{}
				}
				foreign.R.CompleteMissionMissionReleases = append(foreign.R.CompleteMissionMissionReleases, local)
				break
			}
		}
	}

	return nil
}

// LoadReleaseMission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (missionReleaseL) LoadReleaseMission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMissionRelease interface{}, mods queries.Applicator) error {
	var slice []*MissionRelease
	var object *MissionRelease

	if singular {
		var ok bool
		object, ok = maybeMissionRelease.(*MissionRelease)
		if !ok {
			object = new(MissionRelease)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMissionRelease)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMissionRelease))
			}
		}
	} else {
		s, ok := maybeMissionRelease.(*[]*MissionRelease)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMissionRelease)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMissionRelease))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &missionReleaseR{}
		}
		args = append(args, object.ReleaseMissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &missionReleaseR{}
			}

			for _, a := range args {
				if a == obj.ReleaseMissionID {
					continue Outer
				}
			}

			args = append(args, obj.ReleaseMissionID)

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
		object.R.ReleaseMission = foreign
		if foreign.R == nil {
			foreign.R = &missionR{}
		}
		foreign.R.ReleaseMissionMissionReleases = append(foreign.R.ReleaseMissionMissionReleases, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ReleaseMissionID == foreign.ID {
				local.R.ReleaseMission = foreign
				if foreign.R == nil {
					foreign.R = &missionR{}
				}
				foreign.R.ReleaseMissionMissionReleases = append(foreign.R.ReleaseMissionMissionReleases, local)
				break
			}
		}
	}

	return nil
}

// SetCompleteMission of the missionRelease to the related item.
// Sets o.R.CompleteMission to related.
// Adds o to related.R.CompleteMissionMissionReleases.
func (o *MissionRelease) SetCompleteMission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Mission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mission_releases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"complete_mission_id"}),
		strmangle.WhereClause("\"", "\"", 2, missionReleasePrimaryKeyColumns),
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

	o.CompleteMissionID = related.ID
	if o.R == nil {
		o.R = &missionReleaseR{
			CompleteMission: related,
		}
	} else {
		o.R.CompleteMission = related
	}

	if related.R == nil {
		related.R = &missionR{
			CompleteMissionMissionReleases: MissionReleaseSlice{o},
		}
	} else {
		related.R.CompleteMissionMissionReleases = append(related.R.CompleteMissionMissionReleases, o)
	}

	return nil
}

// SetReleaseMission of the missionRelease to the related item.
// Sets o.R.ReleaseMission to related.
// Adds o to related.R.ReleaseMissionMissionReleases.
func (o *MissionRelease) SetReleaseMission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Mission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mission_releases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"release_mission_id"}),
		strmangle.WhereClause("\"", "\"", 2, missionReleasePrimaryKeyColumns),
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

	o.ReleaseMissionID = related.ID
	if o.R == nil {
		o.R = &missionReleaseR{
			ReleaseMission: related,
		}
	} else {
		o.R.ReleaseMission = related
	}

	if related.R == nil {
		related.R = &missionR{
			ReleaseMissionMissionReleases: MissionReleaseSlice{o},
		}
	} else {
		related.R.ReleaseMissionMissionReleases = append(related.R.ReleaseMissionMissionReleases, o)
	}

	return nil
}

// MissionReleases retrieves all the records using an executor.
func MissionReleases(mods ...qm.QueryMod) missionReleaseQuery {
	mods = append(mods, qm.From("\"mission_releases\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mission_releases\".*"})
	}

	return missionReleaseQuery{q}
}

// FindMissionRelease retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMissionRelease(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MissionRelease, error) {
	missionReleaseObj := &MissionRelease{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mission_releases\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, missionReleaseObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mission_releases")
	}

	if err = missionReleaseObj.doAfterSelectHooks(ctx, exec); err != nil {
		return missionReleaseObj, err
	}

	return missionReleaseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MissionRelease) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mission_releases provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(missionReleaseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	missionReleaseInsertCacheMut.RLock()
	cache, cached := missionReleaseInsertCache[key]
	missionReleaseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			missionReleaseAllColumns,
			missionReleaseColumnsWithDefault,
			missionReleaseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(missionReleaseType, missionReleaseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(missionReleaseType, missionReleaseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mission_releases\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mission_releases\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mission_releases")
	}

	if !cached {
		missionReleaseInsertCacheMut.Lock()
		missionReleaseInsertCache[key] = cache
		missionReleaseInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MissionRelease.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MissionRelease) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	missionReleaseUpdateCacheMut.RLock()
	cache, cached := missionReleaseUpdateCache[key]
	missionReleaseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			missionReleaseAllColumns,
			missionReleasePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mission_releases, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mission_releases\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, missionReleasePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(missionReleaseType, missionReleaseMapping, append(wl, missionReleasePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mission_releases row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mission_releases")
	}

	if !cached {
		missionReleaseUpdateCacheMut.Lock()
		missionReleaseUpdateCache[key] = cache
		missionReleaseUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q missionReleaseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mission_releases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mission_releases")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MissionReleaseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionReleasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mission_releases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, missionReleasePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in missionRelease slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all missionRelease")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MissionRelease) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mission_releases provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(missionReleaseColumnsWithDefault, o)

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

	missionReleaseUpsertCacheMut.RLock()
	cache, cached := missionReleaseUpsertCache[key]
	missionReleaseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			missionReleaseAllColumns,
			missionReleaseColumnsWithDefault,
			missionReleaseColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			missionReleaseAllColumns,
			missionReleasePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mission_releases, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(missionReleasePrimaryKeyColumns))
			copy(conflict, missionReleasePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mission_releases\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(missionReleaseType, missionReleaseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(missionReleaseType, missionReleaseMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mission_releases")
	}

	if !cached {
		missionReleaseUpsertCacheMut.Lock()
		missionReleaseUpsertCache[key] = cache
		missionReleaseUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MissionRelease record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MissionRelease) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MissionRelease provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), missionReleasePrimaryKeyMapping)
	sql := "DELETE FROM \"mission_releases\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mission_releases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mission_releases")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q missionReleaseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no missionReleaseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mission_releases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mission_releases")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MissionReleaseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(missionReleaseBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionReleasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mission_releases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, missionReleasePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from missionRelease slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mission_releases")
	}

	if len(missionReleaseAfterDeleteHooks) != 0 {
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
func (o *MissionRelease) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMissionRelease(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MissionReleaseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MissionReleaseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), missionReleasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mission_releases\".* FROM \"mission_releases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, missionReleasePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MissionReleaseSlice")
	}

	*o = slice

	return nil
}

// MissionReleaseExists checks if the MissionRelease row exists.
func MissionReleaseExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mission_releases\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mission_releases exists")
	}

	return exists, nil
}

// Exists checks if the MissionRelease row exists.
func (o *MissionRelease) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MissionReleaseExists(ctx, exec, o.ID)
}
