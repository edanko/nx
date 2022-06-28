package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

// SheetPathModel represents a row from 'dbo.nxsheetpath'.
type SheetPathModel struct {
	ID                int64           `db:"nxsheetpathid"`
	PathID            sql.NullInt64   `db:"nxpathid"`
	MatOrderlineID    int64           `db:"nxmatorderlineid"`
	InventoryID       sql.NullInt64   `db:"nxinventoryid"`
	NestReservationID sql.NullInt64   `db:"nxnestreservationid"`
	ProductID         sql.NullInt64   `db:"nxproductid"`
	PartID            sql.NullInt64   `db:"nxpartid"`
	CutLength         sql.NullFloat64 `db:"nxcutlength"`
	PowderLength      sql.NullFloat64 `db:"nxpowderlength"`
	RapidLength       sql.NullFloat64 `db:"nxrapidlength"`
	NetArea           sql.NullFloat64 `db:"nxspnetarea"`
	UsedArea          sql.NullFloat64 `db:"nxusedarea"`
	Used              sql.NullFloat64 `db:"nxused"`
	MarkCount         sql.NullInt64   `db:"nxmarkcount"`
	ProfCount         sql.NullInt64   `db:"nxprofcount"`
	IsChanged         sql.NullInt64   `db:"nxspischanged"`
	InsertDate        sql.NullTime    `db:"nxspinsertdate"`
}

type SheetPathRepository struct {
	dbs map[string]*sqlx.DB
}

func NewSheetPathRepository(db map[string]*sqlx.DB) *SheetPathRepository {
	return &SheetPathRepository{
		dbs: db,
	}
}

// // Insert inserts the SheetPathModel to the database.
// func (r *SheetPathRepository) Insert(ctx context.Context, n *SheetPathModel) error {
// 	// insert (manual)
// 	const sqlstr = `INSERT INTO dbo.nxsheetpath (
// 		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
// 		) VALUES (
// 		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17
// 		)`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.PathID,
// 		n.MatOrderlineID,
// 		n.InventoryID,
// 		n.NestReservationID,
// 		n.ProductID,
// 		n.PartID,
// 		n.CutLength,
// 		n.PowderLength,
// 		n.RapidLength,
// 		n.NetArea,
// 		n.UsedArea,
// 		n.Used,
// 		n.MarkCount,
// 		n.ProfCount,
// 		n.IsChanged,
// 		n.InsertDate,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.PathID, n.MatOrderlineID, n.InventoryID, n.NestReservationID, n.ProductID, n.PartID, n.CutLength, n.PowderLength, n.RapidLength, n.NetArea, n.UsedArea, n.Used, n.MarkCount, n.ProfCount, n.IsChanged, n.InsertDate); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// // Update updates a SheetPathModel in the database.
// func (r *SheetPathRepository) Update(ctx context.Context, n *SheetPathModel) error {
// 	// update with primary key
// 	const sqlstr = `UPDATE dbo.nxsheetpath SET ` +
// 		`nxpathid = @p1, nxmatorderlineid = @p2, nxinventoryid = @p3, nxnestreservationid = @p4, nxproductid = @p5, nxpartid = @p6, nxcutlength = @p7, nxpowderlength = @p8, nxrapidlength = @p9, nxspnetarea = @p10, nxusedarea = @p11, nxused = @p12, nxmarkcount = @p13, nxprofcount = @p14, nxspischanged = @p15, nxspinsertdate = @p16 ` +
// 		`WHERE nxsheetpathid = @p17`
// 	logf(tenantID,
// 		sqlstr,
// 		n.PathID,
// 		n.MatOrderlineID,
// 		n.InventoryID,
// 		n.NestReservationID,
// 		n.ProductID,
// 		n.PartID,
// 		n.CutLength,
// 		n.PowderLength,
// 		n.RapidLength,
// 		n.NetArea,
// 		n.UsedArea,
// 		n.Used,
// 		n.MarkCount,
// 		n.ProfCount,
// 		n.IsChanged,
// 		n.InsertDate,
// 		n.ID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.PathID, n.MatOrderlineID, n.InventoryID, n.NestReservationID, n.ProductID, n.PartID, n.CutLength, n.PowderLength, n.RapidLength, n.NetArea, n.UsedArea, n.Used, n.MarkCount, n.ProfCount, n.IsChanged, n.InsertDate, n.ID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// // Upsert performs an upsert for SheetPathModel.
// func (r *SheetPathRepository) Upsert(ctx context.Context, n *SheetPathModel) error {
// 	// upsert
// 	const sqlstr = `MERGE dbo.nxsheetpath AS t
// 		USING (
// 		SELECT @p1 nxsheetpathid, @p2 nxpathid, @p3 nxmatorderlineid, @p4 nxinventoryid, @p5 nxnestreservationid, @p6 nxproductid, @p7 nxpartid, @p8 nxcutlength, @p9 nxpowderlength, @p10 nxrapidlength, @p11 nxspnetarea, @p12 nxusedarea, @p13 nxused, @p14 nxmarkcount, @p15 nxprofcount, @p16 nxspischanged, @p17 nxspinsertdate
// 		) AS s
// 		ON s.nxsheetpathid = t.nxsheetpathid
// 		WHEN MATCHED THEN
// 		UPDATE SET
// 		t.nxpathid = s.nxpathid, t.nxmatorderlineid = s.nxmatorderlineid, t.nxinventoryid = s.nxinventoryid, t.nxnestreservationid = s.nxnestreservationid, t.nxproductid = s.nxproductid, t.nxpartid = s.nxpartid, t.nxcutlength = s.nxcutlength, t.nxpowderlength = s.nxpowderlength, t.nxrapidlength = s.nxrapidlength, t.nxspnetarea = s.nxspnetarea, t.nxusedarea = s.nxusedarea, t.nxused = s.nxused, t.nxmarkcount = s.nxmarkcount, t.nxprofcount = s.nxprofcount, t.nxspischanged = s.nxspischanged, t.nxspinsertdate = s.nxspinsertdate
// 		WHEN NOT MATCHED THEN
// 		INSERT (
// 		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
// 		) VALUES (
// 		s.nxsheetpathid, s.nxpathid, s.nxmatorderlineid, s.nxinventoryid, s.nxnestreservationid, s.nxproductid, s.nxpartid, s.nxcutlength, s.nxpowderlength, s.nxrapidlength, s.nxspnetarea, s.nxusedarea, s.nxused, s.nxmarkcount, s.nxprofcount, s.nxspischanged, s.nxspinsertdate
// 		);`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.PathID,
// 		n.MatOrderlineID,
// 		n.InventoryID,
// 		n.NestReservationID,
// 		n.ProductID,
// 		n.PartID,
// 		n.CutLength,
// 		n.PowderLength,
// 		n.RapidLength,
// 		n.NetArea,
// 		n.UsedArea,
// 		n.Used,
// 		n.MarkCount,
// 		n.ProfCount,
// 		n.IsChanged,
// 		n.InsertDate,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.PathID, n.MatOrderlineID, n.InventoryID, n.NestReservationID, n.ProductID, n.PartID, n.CutLength, n.PowderLength, n.RapidLength, n.NetArea, n.UsedArea, n.Used, n.MarkCount, n.ProfCount, n.IsChanged, n.InsertDate); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// Delete deletes the SheetPathModel from the database.
func (r *SheetPathRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxsheetpath
		WHERE nxsheetpathid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}
	return nil
}

// Get retrieves a row from 'dbo.nxsheetpath' as a SheetPathModel.
func (r *SheetPathRepository) Get(ctx context.Context, id int64) (*SheetPathModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT ` +
		`nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate ` +
		`FROM dbo.nxsheetpath ` +
		`WHERE nxsheetpathid = @p1`
	logf(tenantID, sqlstr, id)
	n := SheetPathModel{}
	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}

func (r *SheetPathRepository) GetByPathID(ctx context.Context, id int64) (*SheetPathModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT ` +
		`nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate ` +
		`FROM dbo.nxsheetpath ` +
		`WHERE nxpathid = @p1`
	logf(tenantID, sqlstr, id)
	n := SheetPathModel{}
	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}

func (r *SheetPathRepository) GetByPathIDs(ctx context.Context, ids []int64) ([]*SheetPathModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT ` +
		`nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate ` +
		`FROM dbo.nxsheetpath ` +
		`WHERE nxpathid IN (?)`
	query, args, err := sqlx.In(sqlstr, ids)
	query = r.dbs[tenantID].Rebind(query)

	logf(tenantID, query, ids)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*SheetPathModel
	for rows.Next() {
		var n SheetPathModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

// // GetByOrderlineID retrieves a row from 'dbo.nxsheetpath' as a SheetPathModel.
// func (r *SheetPathRepository) GetByOrderlineID(ctx context.Context, id int64) (*SheetPathModel, error) {
// 	tenantID, ok := tenant.FromContext(ctx)
// 	if !ok {
// 		return nil, errors.New("tenant id not found in context")
// 	}
//
// 	const sqlstr = `SELECT ` +
// 		`nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate ` +
// 		`FROM dbo.nxsheetpath ` +
// 		`WHERE nxmatorderlineid = @p1`
// 	logf(tenantID, sqlstr, id)
// 	n := SheetPathModel{}
// 	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
// 	err := row.StructScan(&n)
// 	if err != nil {
// 		return nil, logerror(err)
// 	}
// 	return &n, nil
// }
