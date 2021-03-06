// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Allergy is an object representing the database table.
type Allergy struct {
	ID       int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Category string    `boil:"category" json:"category" toml:"category" yaml:"category"`
	Menu     string    `boil:"menu" json:"menu" toml:"menu" yaml:"menu"`
	Shrimp   null.Bool `boil:"shrimp" json:"shrimp,omitempty" toml:"shrimp" yaml:"shrimp,omitempty"`
	Crab     null.Bool `boil:"crab" json:"crab,omitempty" toml:"crab" yaml:"crab,omitempty"`
	Wheat    null.Bool `boil:"wheat" json:"wheat,omitempty" toml:"wheat" yaml:"wheat,omitempty"`
	Soba     null.Bool `boil:"soba" json:"soba,omitempty" toml:"soba" yaml:"soba,omitempty"`
	Eggs     null.Bool `boil:"eggs" json:"eggs,omitempty" toml:"eggs" yaml:"eggs,omitempty"`
	Milk     null.Bool `boil:"milk" json:"milk,omitempty" toml:"milk" yaml:"milk,omitempty"`
	Peanuts  null.Bool `boil:"peanuts" json:"peanuts,omitempty" toml:"peanuts" yaml:"peanuts,omitempty"`
	Walnuts  null.Bool `boil:"walnuts" json:"walnuts,omitempty" toml:"walnuts" yaml:"walnuts,omitempty"`

	R *allergyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L allergyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AllergyColumns = struct {
	ID       string
	Category string
	Menu     string
	Shrimp   string
	Crab     string
	Wheat    string
	Soba     string
	Eggs     string
	Milk     string
	Peanuts  string
	Walnuts  string
}{
	ID:       "id",
	Category: "category",
	Menu:     "menu",
	Shrimp:   "shrimp",
	Crab:     "crab",
	Wheat:    "wheat",
	Soba:     "soba",
	Eggs:     "eggs",
	Milk:     "milk",
	Peanuts:  "peanuts",
	Walnuts:  "walnuts",
}

var AllergyTableColumns = struct {
	ID       string
	Category string
	Menu     string
	Shrimp   string
	Crab     string
	Wheat    string
	Soba     string
	Eggs     string
	Milk     string
	Peanuts  string
	Walnuts  string
}{
	ID:       "allergies.id",
	Category: "allergies.category",
	Menu:     "allergies.menu",
	Shrimp:   "allergies.shrimp",
	Crab:     "allergies.crab",
	Wheat:    "allergies.wheat",
	Soba:     "allergies.soba",
	Eggs:     "allergies.eggs",
	Milk:     "allergies.milk",
	Peanuts:  "allergies.peanuts",
	Walnuts:  "allergies.walnuts",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AllergyWhere = struct {
	ID       whereHelperint
	Category whereHelperstring
	Menu     whereHelperstring
	Shrimp   whereHelpernull_Bool
	Crab     whereHelpernull_Bool
	Wheat    whereHelpernull_Bool
	Soba     whereHelpernull_Bool
	Eggs     whereHelpernull_Bool
	Milk     whereHelpernull_Bool
	Peanuts  whereHelpernull_Bool
	Walnuts  whereHelpernull_Bool
}{
	ID:       whereHelperint{field: "\"allergies\".\"id\""},
	Category: whereHelperstring{field: "\"allergies\".\"category\""},
	Menu:     whereHelperstring{field: "\"allergies\".\"menu\""},
	Shrimp:   whereHelpernull_Bool{field: "\"allergies\".\"shrimp\""},
	Crab:     whereHelpernull_Bool{field: "\"allergies\".\"crab\""},
	Wheat:    whereHelpernull_Bool{field: "\"allergies\".\"wheat\""},
	Soba:     whereHelpernull_Bool{field: "\"allergies\".\"soba\""},
	Eggs:     whereHelpernull_Bool{field: "\"allergies\".\"eggs\""},
	Milk:     whereHelpernull_Bool{field: "\"allergies\".\"milk\""},
	Peanuts:  whereHelpernull_Bool{field: "\"allergies\".\"peanuts\""},
	Walnuts:  whereHelpernull_Bool{field: "\"allergies\".\"walnuts\""},
}

// AllergyRels is where relationship names are stored.
var AllergyRels = struct {
}{}

// allergyR is where relationships are stored.
type allergyR struct {
}

// NewStruct creates a new relationship struct
func (*allergyR) NewStruct() *allergyR {
	return &allergyR{}
}

// allergyL is where Load methods for each relationship are stored.
type allergyL struct{}

var (
	allergyAllColumns            = []string{"id", "category", "menu", "shrimp", "crab", "wheat", "soba", "eggs", "milk", "peanuts", "walnuts"}
	allergyColumnsWithoutDefault = []string{"category", "menu", "shrimp", "crab", "wheat", "soba", "eggs", "milk", "peanuts", "walnuts"}
	allergyColumnsWithDefault    = []string{"id"}
	allergyPrimaryKeyColumns     = []string{"id"}
)

type (
	// AllergySlice is an alias for a slice of pointers to Allergy.
	// This should almost always be used instead of []Allergy.
	AllergySlice []*Allergy
	// AllergyHook is the signature for custom Allergy hook methods
	AllergyHook func(context.Context, boil.ContextExecutor, *Allergy) error

	allergyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	allergyType                 = reflect.TypeOf(&Allergy{})
	allergyMapping              = queries.MakeStructMapping(allergyType)
	allergyPrimaryKeyMapping, _ = queries.BindMapping(allergyType, allergyMapping, allergyPrimaryKeyColumns)
	allergyInsertCacheMut       sync.RWMutex
	allergyInsertCache          = make(map[string]insertCache)
	allergyUpdateCacheMut       sync.RWMutex
	allergyUpdateCache          = make(map[string]updateCache)
	allergyUpsertCacheMut       sync.RWMutex
	allergyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var allergyBeforeInsertHooks []AllergyHook
var allergyBeforeUpdateHooks []AllergyHook
var allergyBeforeDeleteHooks []AllergyHook
var allergyBeforeUpsertHooks []AllergyHook

var allergyAfterInsertHooks []AllergyHook
var allergyAfterSelectHooks []AllergyHook
var allergyAfterUpdateHooks []AllergyHook
var allergyAfterDeleteHooks []AllergyHook
var allergyAfterUpsertHooks []AllergyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Allergy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Allergy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Allergy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Allergy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Allergy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Allergy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Allergy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Allergy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Allergy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allergyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAllergyHook registers your hook function for all future operations.
func AddAllergyHook(hookPoint boil.HookPoint, allergyHook AllergyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		allergyBeforeInsertHooks = append(allergyBeforeInsertHooks, allergyHook)
	case boil.BeforeUpdateHook:
		allergyBeforeUpdateHooks = append(allergyBeforeUpdateHooks, allergyHook)
	case boil.BeforeDeleteHook:
		allergyBeforeDeleteHooks = append(allergyBeforeDeleteHooks, allergyHook)
	case boil.BeforeUpsertHook:
		allergyBeforeUpsertHooks = append(allergyBeforeUpsertHooks, allergyHook)
	case boil.AfterInsertHook:
		allergyAfterInsertHooks = append(allergyAfterInsertHooks, allergyHook)
	case boil.AfterSelectHook:
		allergyAfterSelectHooks = append(allergyAfterSelectHooks, allergyHook)
	case boil.AfterUpdateHook:
		allergyAfterUpdateHooks = append(allergyAfterUpdateHooks, allergyHook)
	case boil.AfterDeleteHook:
		allergyAfterDeleteHooks = append(allergyAfterDeleteHooks, allergyHook)
	case boil.AfterUpsertHook:
		allergyAfterUpsertHooks = append(allergyAfterUpsertHooks, allergyHook)
	}
}

// One returns a single allergy record from the query.
func (q allergyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Allergy, error) {
	o := &Allergy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for allergies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Allergy records from the query.
func (q allergyQuery) All(ctx context.Context, exec boil.ContextExecutor) (AllergySlice, error) {
	var o []*Allergy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Allergy slice")
	}

	if len(allergyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Allergy records in the query.
func (q allergyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count allergies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q allergyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if allergies exists")
	}

	return count > 0, nil
}

// Allergies retrieves all the records using an executor.
func Allergies(mods ...qm.QueryMod) allergyQuery {
	mods = append(mods, qm.From("\"allergies\""))
	return allergyQuery{NewQuery(mods...)}
}

// FindAllergy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAllergy(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Allergy, error) {
	allergyObj := &Allergy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"allergies\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, allergyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from allergies")
	}

	if err = allergyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return allergyObj, err
	}

	return allergyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Allergy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no allergies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(allergyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	allergyInsertCacheMut.RLock()
	cache, cached := allergyInsertCache[key]
	allergyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			allergyAllColumns,
			allergyColumnsWithDefault,
			allergyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(allergyType, allergyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(allergyType, allergyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"allergies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"allergies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into allergies")
	}

	if !cached {
		allergyInsertCacheMut.Lock()
		allergyInsertCache[key] = cache
		allergyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Allergy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Allergy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	allergyUpdateCacheMut.RLock()
	cache, cached := allergyUpdateCache[key]
	allergyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			allergyAllColumns,
			allergyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update allergies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"allergies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, allergyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(allergyType, allergyMapping, append(wl, allergyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update allergies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for allergies")
	}

	if !cached {
		allergyUpdateCacheMut.Lock()
		allergyUpdateCache[key] = cache
		allergyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q allergyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for allergies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for allergies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AllergySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allergyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"allergies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, allergyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in allergy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all allergy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Allergy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no allergies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(allergyColumnsWithDefault, o)

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

	allergyUpsertCacheMut.RLock()
	cache, cached := allergyUpsertCache[key]
	allergyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			allergyAllColumns,
			allergyColumnsWithDefault,
			allergyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			allergyAllColumns,
			allergyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert allergies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(allergyPrimaryKeyColumns))
			copy(conflict, allergyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"allergies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(allergyType, allergyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(allergyType, allergyMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert allergies")
	}

	if !cached {
		allergyUpsertCacheMut.Lock()
		allergyUpsertCache[key] = cache
		allergyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Allergy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Allergy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Allergy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), allergyPrimaryKeyMapping)
	sql := "DELETE FROM \"allergies\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from allergies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for allergies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q allergyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no allergyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from allergies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for allergies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AllergySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(allergyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allergyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"allergies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, allergyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from allergy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for allergies")
	}

	if len(allergyAfterDeleteHooks) != 0 {
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
func (o *Allergy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAllergy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AllergySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AllergySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allergyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"allergies\".* FROM \"allergies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, allergyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AllergySlice")
	}

	*o = slice

	return nil
}

// AllergyExists checks if the Allergy row exists.
func AllergyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"allergies\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if allergies exists")
	}

	return exists, nil
}
