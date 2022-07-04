package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

// SheetPathDetModel represents a row from 'dbo.nxsheetpathdet'.
type SheetPathDetModel struct {
	ID          int64           `db:"nxsheetpathdetid"`
	SheetPathID int64           `db:"nxsheetpathid"`
	OrderlineID sql.NullInt64   `db:"nxorderlineid"`
	PartID      sql.NullInt64   `db:"nxpartid"`
	ProductID   sql.NullInt64   `db:"nxproductid"`
	DetailCount sql.NullInt64   `db:"nxdetailcount"`
	DetailCode  sql.NullString  `db:"nxdetailcode"`
	UsedArea    sql.NullFloat64 `db:"nxusedarea"`
	Area        sql.NullFloat64 `db:"nxarea"`
	SlagArea    sql.NullFloat64 `db:"nxslagarea"`
	NestScrap   sql.NullFloat64 `db:"nxspdnestscrap"`
	AddScrap    sql.NullFloat64 `db:"nxspdaddscrap"`
	MachTime    sql.NullFloat64 `db:"nxspdmachtime"`
	ClaimArea   sql.NullFloat64 `db:"nxspdclaimarea"`
	ReturnArea  sql.NullFloat64 `db:"nxspdreturnarea"`
	MachTimePct sql.NullFloat64 `db:"nxspdmachtimepct"`
	Type        sql.NullInt64   `db:"nxspdtype"`
	SequenceNo  sql.NullInt64   `db:"nxspdsequenceno"`
}

type SheetPathDetRepository struct {
	dbs map[string]*sqlx.DB
}

func NewSheetPathDetRepository(db map[string]*sqlx.DB) *SheetPathDetRepository {
	return &SheetPathDetRepository{
		dbs: db,
	}
}

// // Insert inserts the SheetPathDetModel to the database.
// func (r *SheetPathDetRepository) Insert(ctx context.Context, n *SheetPathDetModel) error {
// 	// insert (manual)
// 	const sqlstr = `INSERT INTO dbo.nxsheetpathdet (
// 		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
// 		) VALUES (
// 		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18
// 		)`
// 		sqlstr,
// 		n.ID,
// 		n.SheetPathID,
// 		n.OrderlineID,
// 		n.PartID,
// 		n.ProductID,
// 		n.DetailCount,
// 		n.DetailCode,
// 		n.UsedArea,
// 		n.Area,
// 		n.SlagArea,
// 		n.NestScrap,
// 		n.AddScrap,
// 		n.MachTime,
// 		n.ClaimArea,
// 		n.ReturnArea,
// 		n.MachTimePct,
// 		n.Type,
// 		n.SequenceNo,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.SheetPathID, n.OrderlineID, n.PartID, n.ProductID, n.DetailCount, n.DetailCode, n.UsedArea, n.Area, n.SlagArea, n.NestScrap, n.AddScrap, n.MachTime, n.ClaimArea, n.ReturnArea, n.MachTimePct, n.Type, n.SequenceNo); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// // Update updates a SheetPathDetModel in the database.
// func (r *SheetPathDetRepository) Update(ctx context.Context, n *SheetPathDetModel) error {
// 	// update with primary key
// 	const sqlstr = `UPDATE dbo.nxsheetpathdet SET
// 		nxsheetpathid = @p1, nxorderlineid = @p2, nxpartid = @p3, nxproductid = @p4, nxdetailcount = @p5, nxdetailcode = @p6, nxusedarea = @p7, nxarea = @p8, nxslagarea = @p9, nxspdnestscrap = @p10, nxspdaddscrap = @p11, nxspdmachtime = @p12, nxspdclaimarea = @p13, nxspdreturnarea = @p14, nxspdmachtimepct = @p15, nxspdtype = @p16, nxspdsequenceno = @p17
// 		WHERE nxsheetpathdetid = @p18`
// 	logf(tenantID,
// 		sqlstr,
// 		n.SheetPathID,
// 		n.OrderlineID,
// 		n.PartID,
// 		n.ProductID,
// 		n.DetailCount,
// 		n.DetailCode,
// 		n.UsedArea,
// 		n.Area,
// 		n.SlagArea,
// 		n.NestScrap,
// 		n.AddScrap,
// 		n.MachTime,
// 		n.ClaimArea,
// 		n.ReturnArea,
// 		n.MachTimePct,
// 		n.Type,
// 		n.SequenceNo,
// 		n.ID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.SheetPathID, n.OrderlineID, n.PartID, n.ProductID, n.DetailCount, n.DetailCode, n.UsedArea, n.Area, n.SlagArea, n.NestScrap, n.AddScrap, n.MachTime, n.ClaimArea, n.ReturnArea, n.MachTimePct, n.Type, n.SequenceNo, n.ID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// // Upsert performs an upsert for SheetPathDetModel.
// func (r *SheetPathDetRepository) Upsert(ctx context.Context, n *SheetPathDetModel) error {
// 	// upsert
// 	const sqlstr = `MERGE dbo.nxsheetpathdet AS t
// 		USING (
// 		SELECT @p1 nxsheetpathdetid, @p2 nxsheetpathid, @p3 nxorderlineid, @p4 nxpartid, @p5 nxproductid, @p6 nxdetailcount, @p7 nxdetailcode, @p8 nxusedarea, @p9 nxarea, @p10 nxslagarea, @p11 nxspdnestscrap, @p12 nxspdaddscrap, @p13 nxspdmachtime, @p14 nxspdclaimarea, @p15 nxspdreturnarea, @p16 nxspdmachtimepct, @p17 nxspdtype, @p18 nxspdsequenceno
// 		) AS s
// 		ON s.nxsheetpathdetid = t.nxsheetpathdetid
// 		WHEN MATCHED THEN
// 		UPDATE SET
// 		t.nxsheetpathid = s.nxsheetpathid, t.nxorderlineid = s.nxorderlineid, t.nxpartid = s.nxpartid, t.nxproductid = s.nxproductid, t.nxdetailcount = s.nxdetailcount, t.nxdetailcode = s.nxdetailcode, t.nxusedarea = s.nxusedarea, t.nxarea = s.nxarea, t.nxslagarea = s.nxslagarea, t.nxspdnestscrap = s.nxspdnestscrap, t.nxspdaddscrap = s.nxspdaddscrap, t.nxspdmachtime = s.nxspdmachtime, t.nxspdclaimarea = s.nxspdclaimarea, t.nxspdreturnarea = s.nxspdreturnarea, t.nxspdmachtimepct = s.nxspdmachtimepct, t.nxspdtype = s.nxspdtype, t.nxspdsequenceno = s.nxspdsequenceno
// 		WHEN NOT MATCHED THEN
// 		INSERT (
// 		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
// 		) VALUES (
// 		s.nxsheetpathdetid, s.nxsheetpathid, s.nxorderlineid, s.nxpartid, s.nxproductid, s.nxdetailcount, s.nxdetailcode, s.nxusedarea, s.nxarea, s.nxslagarea, s.nxspdnestscrap, s.nxspdaddscrap, s.nxspdmachtime, s.nxspdclaimarea, s.nxspdreturnarea, s.nxspdmachtimepct, s.nxspdtype, s.nxspdsequenceno
// 		);`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.SheetPathID,
// 		n.OrderlineID,
// 		n.PartID,
// 		n.ProductID,
// 		n.DetailCount,
// 		n.DetailCode,
// 		n.UsedArea,
// 		n.Area,
// 		n.SlagArea,
// 		n.NestScrap,
// 		n.AddScrap,
// 		n.MachTime,
// 		n.ClaimArea,
// 		n.ReturnArea,
// 		n.MachTimePct,
// 		n.Type,
// 		n.SequenceNo,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.SheetPathID, n.OrderlineID, n.PartID, n.ProductID, n.DetailCount, n.DetailCode, n.UsedArea, n.Area, n.SlagArea, n.NestScrap, n.AddScrap, n.MachTime, n.ClaimArea, n.ReturnArea, n.MachTimePct, n.Type, n.SequenceNo); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// Delete deletes the SheetPathDetModel from the database.
func (r *SheetPathDetRepository) Delete(ctx context.Context, n *SheetPathDetModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxsheetpathdet
		WHERE nxsheetpathdetid = @p1`
	logf(tenantID, sqlstr, n.ID)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

func (r *SheetPathDetRepository) ListPartsBySheetPathID(ctx context.Context, id int64) ([]*SheetPathDetModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
		FROM dbo.nxsheetpathdet
		WHERE nxsheetpathid = @p1
		ORDER BY nxdetailcode ASC`
	logf(tenantID, sqlstr, id)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr, id)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*SheetPathDetModel

	for rows.Next() {
		n := SheetPathDetModel{}

		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}
	return nn, nil
}
