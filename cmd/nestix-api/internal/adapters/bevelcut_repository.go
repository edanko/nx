package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

// BevelCutModel represents a row from 'dbo.nxbevelcut'.
type BevelCutModel struct {
	ID           int64           `db:"nxbevelcutid"`
	Name         string          `db:"nxbvlcutname"`
	AngleMin     sql.NullFloat64 `db:"nxbvlanglemin"`
	AngleMax     sql.NullFloat64 `db:"nxbvlanglemax"`
	RootWidthMin sql.NullFloat64 `db:"nxbvlrootwidthmin"`
	RootWidthMax sql.NullFloat64 `db:"nxbvlrootwidthmax"`
	EdgeDistMin  sql.NullFloat64 `db:"nxbvledgedistmin"`
	EdgeDistMax  sql.NullFloat64 `db:"nxbvledgedistmax"`
	ThickMin     float32         `db:"nxbvlthickmin"`
	ThickMax     float32         `db:"nxbvlthickmax"`
	Speed        float32         `db:"nxbvlspeed"`
	CorrAngle    float32         `db:"nxbvlcorrangle"`
	CorrWidth    float32         `db:"nxbvlcorrwidth"`
	Tech         sql.NullString  `db:"nxbvltech"`
	TechGroup    sql.NullString  `db:"nxbvltechgroup"`
	MatGroup     sql.NullString  `db:"nxbvlmatgroup"`
	Creator      sql.NullString  `db:"nxbvlcreator"`
	Created      sql.NullTime    `db:"nxbvlcreated"`
	Changer      sql.NullString  `db:"nxbvlchanger"`
	Changed      sql.NullTime    `db:"nxbvlchanged"`
}

type BevelCutRepository struct {
	dbs map[string]*sqlx.DB
}

func NewBevelCutRepository(db map[string]*sqlx.DB) *BevelCutRepository {
	return &BevelCutRepository{
		dbs: db,
	}
}

// Insert inserts the BevelCutModel to the database.
func (r *BevelCutRepository) Insert(ctx context.Context, n *BevelCutModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxbevelcut (
		nxbevelcutid, nxbvlcutname, nxbvlanglemin, nxbvlanglemax, nxbvlrootwidthmin, nxbvlrootwidthmax, nxbvledgedistmin, nxbvledgedistmax, nxbvlthickmin, nxbvlthickmax, nxbvlspeed, nxbvlcorrangle, nxbvlcorrwidth, nxbvltech, nxbvltechgroup, nxbvlmatgroup, nxbvlcreator, nxbvlcreated, nxbvlchanger, nxbvlchanged
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20
		)`
	logf(tenantID,
		sqlstr,
		n.ID,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
	)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.Name, n.AngleMin, n.AngleMax, n.RootWidthMin, n.RootWidthMax, n.EdgeDistMin, n.EdgeDistMax, n.ThickMin, n.ThickMax, n.Speed, n.CorrAngle, n.CorrWidth, n.Tech, n.TechGroup, n.MatGroup, n.Creator, n.Created, n.Changer, n.Changed); err != nil {
		return logerror(err)
	}

	return nil
}

// Update updates a BevelCutModel in the database.
func (r *BevelCutRepository) Update(ctx context.Context, n *BevelCutModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxbevelcut SET ` +
		`nxbvlcutname = @p1, nxbvlanglemin = @p2, nxbvlanglemax = @p3, nxbvlrootwidthmin = @p4, nxbvlrootwidthmax = @p5, nxbvledgedistmin = @p6, nxbvledgedistmax = @p7, nxbvlthickmin = @p8, nxbvlthickmax = @p9, nxbvlspeed = @p10, nxbvlcorrangle = @p11, nxbvlcorrwidth = @p12, nxbvltech = @p13, nxbvltechgroup = @p14, nxbvlmatgroup = @p15, nxbvlcreator = @p16, nxbvlcreated = @p17, nxbvlchanger = @p18, nxbvlchanged = @p19 ` +
		`WHERE nxbevelcutid = @p20`
	logf(tenantID,
		sqlstr,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.Name, n.AngleMin, n.AngleMax, n.RootWidthMin, n.RootWidthMax, n.EdgeDistMin, n.EdgeDistMax, n.ThickMin, n.ThickMax, n.Speed, n.CorrAngle, n.CorrWidth, n.Tech, n.TechGroup, n.MatGroup, n.Creator, n.Created, n.Changer, n.Changed, n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// // Save saves the BevelCutModel to the database.
// func (n *BevelCutModel) Save(ctx context.Context, db *sqlx.DB) error {
// 	if n.Exists() {
// 		return n.Update(ctx, db)
// 	}
// 	return n.Insert(ctx, db)
// }

// Upsert performs an upsert for BevelCutModel.
func (r *BevelCutRepository) Upsert(ctx context.Context, n *BevelCutModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `MERGE dbo.nxbevelcut AS t ` +
		`USING (` +
		`SELECT @p1 nxbevelcutid, @p2 nxbvlcutname, @p3 nxbvlanglemin, @p4 nxbvlanglemax, @p5 nxbvlrootwidthmin, @p6 nxbvlrootwidthmax, @p7 nxbvledgedistmin, @p8 nxbvledgedistmax, @p9 nxbvlthickmin, @p10 nxbvlthickmax, @p11 nxbvlspeed, @p12 nxbvlcorrangle, @p13 nxbvlcorrwidth, @p14 nxbvltech, @p15 nxbvltechgroup, @p16 nxbvlmatgroup, @p17 nxbvlcreator, @p18 nxbvlcreated, @p19 nxbvlchanger, @p20 nxbvlchanged ` +
		`) AS s ` +
		`ON s.nxbevelcutid = t.nxbevelcutid ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.nxbvlcutname = s.nxbvlcutname, t.nxbvlanglemin = s.nxbvlanglemin, t.nxbvlanglemax = s.nxbvlanglemax, t.nxbvlrootwidthmin = s.nxbvlrootwidthmin, t.nxbvlrootwidthmax = s.nxbvlrootwidthmax, t.nxbvledgedistmin = s.nxbvledgedistmin, t.nxbvledgedistmax = s.nxbvledgedistmax, t.nxbvlthickmin = s.nxbvlthickmin, t.nxbvlthickmax = s.nxbvlthickmax, t.nxbvlspeed = s.nxbvlspeed, t.nxbvlcorrangle = s.nxbvlcorrangle, t.nxbvlcorrwidth = s.nxbvlcorrwidth, t.nxbvltech = s.nxbvltech, t.nxbvltechgroup = s.nxbvltechgroup, t.nxbvlmatgroup = s.nxbvlmatgroup, t.nxbvlcreator = s.nxbvlcreator, t.nxbvlcreated = s.nxbvlcreated, t.nxbvlchanger = s.nxbvlchanger, t.nxbvlchanged = s.nxbvlchanged ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`nxbevelcutid, nxbvlcutname, nxbvlanglemin, nxbvlanglemax, nxbvlrootwidthmin, nxbvlrootwidthmax, nxbvledgedistmin, nxbvledgedistmax, nxbvlthickmin, nxbvlthickmax, nxbvlspeed, nxbvlcorrangle, nxbvlcorrwidth, nxbvltech, nxbvltechgroup, nxbvlmatgroup, nxbvlcreator, nxbvlcreated, nxbvlchanger, nxbvlchanged` +
		`) VALUES (` +
		`s.nxbevelcutid, s.nxbvlcutname, s.nxbvlanglemin, s.nxbvlanglemax, s.nxbvlrootwidthmin, s.nxbvlrootwidthmax, s.nxbvledgedistmin, s.nxbvledgedistmax, s.nxbvlthickmin, s.nxbvlthickmax, s.nxbvlspeed, s.nxbvlcorrangle, s.nxbvlcorrwidth, s.nxbvltech, s.nxbvltechgroup, s.nxbvlmatgroup, s.nxbvlcreator, s.nxbvlcreated, s.nxbvlchanger, s.nxbvlchanged` +
		`);`
	logf(tenantID,
		sqlstr,
		n.ID,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
	)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.Name, n.AngleMin, n.AngleMax, n.RootWidthMin, n.RootWidthMax, n.EdgeDistMin, n.EdgeDistMax, n.ThickMin, n.ThickMax, n.Speed, n.CorrAngle, n.CorrWidth, n.Tech, n.TechGroup, n.MatGroup, n.Creator, n.Created, n.Changer, n.Changed); err != nil {
		return logerror(err)
	}

	return nil
}

// Delete deletes the BevelCutModel from the database.
func (r *BevelCutRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxbevelcut ` +
		`WHERE nxbevelcutid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}

	return nil
}
