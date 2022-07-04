package adapters

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

// OrderModel represents a row from 'dbo.nxorderline'.
type OrderModel struct {
	ID               int64           `db:"nxorderlineid"`
	ProductID        int64           `db:"nxproductid"`
	PartID           sql.NullInt64   `db:"nxpartid"`
	OrderNo          sql.NullString  `db:"nxororderno"`
	LineNo           int64           `db:"nxollineno"`
	ParentName       sql.NullString  `db:"nxolparentname"`
	Section          sql.NullString  `db:"nxolsection"`
	MrpLineNo        sql.NullString  `db:"nxolmrplineno"`
	Count            int64           `db:"nxolordercount"`
	Rejected         sql.NullInt64   `db:"nxolrejected"`
	DueDate          sql.NullTime    `db:"nxolduedate"`
	NestedCount      sql.NullInt64   `db:"nxolnestedcnt"`
	Rotate           sql.NullInt64   `db:"nxolrotate"`
	Status           sql.NullInt64   `db:"nxolstatus"`
	MatStd           sql.NullString  `db:"nxolmatstd"`
	RowType          int64           `db:"nxolrowtype"`
	InfoTxt          sql.NullString  `db:"nxolinfotxt"`
	TotWeight        sql.NullFloat64 `db:"nxoltotweight"`
	Weight           sql.NullFloat64 `db:"nxolweight"`
	Thick            sql.NullFloat64 `db:"nxolthick"`
	Width            sql.NullFloat64 `db:"nxolwidth"`
	Length           sql.NullFloat64 `db:"nxollength"`
	PurCode          sql.NullInt64   `db:"nxolpurcode"`
	PathName         sql.NullString  `db:"nxolnestixno"`
	Ready            sql.NullInt64   `db:"nxolready"`
	Source           sql.NullInt64   `db:"nxolsource"`
	Type             sql.NullInt64   `db:"nxoltype"`
	Mirror           sql.NullInt64   `db:"nxolmirror"`
	ProdInfo         sql.NullString  `db:"nxolprodinfo"`
	PartKey          sql.NullInt64   `db:"nxolpartkey"`
	AttributeKey     sql.NullString  `db:"nxolattributekey"`
	LogisticalKey    sql.NullInt64   `db:"nxollogisticalkey"`
	CenterOfGravityX sql.NullFloat64 `db:"nxolcenterofgravityx"`
	CenterOfGravityY sql.NullFloat64 `db:"nxolcenterofgravityy"`
	CenterOfGravityZ sql.NullFloat64 `db:"nxolcenterofgravityz"`
	ProjectName      sql.NullString  `db:"nxolprojectname"`
	AssemblySequence sql.NullInt64   `db:"nxolassemblysequence"`
	Name             sql.NullString  `db:"nxolname"`
	CustomerID       sql.NullInt64   `db:"nxcustomerid"`
	CustOrderNo      sql.NullString  `db:"nxolcustorderno"`
	CustMark         sql.NullString  `db:"nxolcustmark"`
	Workphases       sql.NullString  `db:"nxolworkphases"`
	Info1            sql.NullString  `db:"nxolorderinfo1"`
	Info2            sql.NullString  `db:"nxolorderinfo2"`
	Info3            sql.NullString  `db:"nxolorderinfo3"`
	ChargeNo         sql.NullString  `db:"nxolchargeno"`
	InsertDate       time.Time       `db:"nxolinsertdate"`
	Created          sql.NullTime    `db:"nxolcreated"`
	Creator          sql.NullString  `db:"nxolcreator"`
	Changed          sql.NullTime    `db:"nxolchanged"`
	Changer          sql.NullString  `db:"nxolchanger"`
	SourceOrderNo    sql.NullString  `db:"nxolsourceorderno"`
	PartSide         sql.NullInt64   `db:"nxolpartside"`
}

type OrderRepository struct {
	dbs map[string]*sqlx.DB
}

func NewOrderRepository(db map[string]*sqlx.DB) *OrderRepository {
	return &OrderRepository{
		dbs: db,
	}
}

//
// // Insert inserts the OrderModel to the database.
// func (r *OrderRepository) Insert(ctx context.Context, n *OrderModel) error {
// 	// insert (manual)
// 	const sqlstr = `INSERT INTO dbo.nxorderline (
// 		nxorderlineid, nxproductid, nxpartid, nxororderno, nxollineno, nxolparentname, nxolsection, nxolmrplineno, nxolordercount, nxolrejected, nxolduedate, nxolnestedcnt, nxolrotate, nxolstatus, nxolmatstd, nxolrowtype, nxolinfotxt, nxoltotweight, nxolweight, nxolthick, nxolwidth, nxollength, nxolpurcode, nxolnestixno, nxolready, nxolsource, nxoltype, nxolmirror, nxolprodinfo, nxolpartkey, nxolattributekey, nxollogisticalkey, nxolcenterofgravityx, nxolcenterofgravityy, nxolcenterofgravityz, nxolprojectname, nxolassemblysequence, nxolname, nxcustomerid, nxolcustorderno, nxolcustmark, nxolworkphases, nxolorderinfo1, nxolorderinfo2, nxolorderinfo3, nxolchargeno, nxolinsertdate, nxolcreated, nxolcreator, nxolchanged, nxolchanger, nxolsourceorderno, nxolpartside
// 		) VALUES (
// 		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30, @p31, @p32, @p33, @p34, @p35, @p36, @p37, @p38, @p39, @p40, @p41, @p42, @p43, @p44, @p45, @p46, @p47, @p48, @p49, @p50, @p51, @p52, @p53
// 		)`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.ProductID,
// 		n.PartID,
// 		n.OrderNo,
// 		n.LineNo,
// 		n.ParentName,
// 		n.Section,
// 		n.MrpLineNo,
// 		n.Count,
// 		n.Rejected,
// 		n.DueDate,
// 		n.NestedCount,
// 		n.Rotate,
// 		n.Status,
// 		n.MatStd,
// 		n.RowType,
// 		n.InfoTxt,
// 		n.TotWeight,
// 		n.Weight,
// 		n.Thick,
// 		n.Width,
// 		n.Length,
// 		n.PurCode,
// 		n.PathName,
// 		n.Ready,
// 		n.Source,
// 		n.Type,
// 		n.Mirror,
// 		n.ProdInfo,
// 		n.PartKey,
// 		n.AttributeKey,
// 		n.LogisticalKey,
// 		n.CenterOfGravityX,
// 		n.CenterOfGravityY,
// 		n.CenterOfGravityZ,
// 		n.ProjectName,
// 		n.AssemblySequence,
// 		n.Name,
// 		n.CustomerID,
// 		n.CustOrderNo,
// 		n.CustMark,
// 		n.Workphases,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.ChargeNo,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.SourceOrderNo,
// 		n.PartSide,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.ProductID, n.PartID, n.OrderNo, n.LineNo, n.ParentName, n.Section, n.MrpLineNo, n.Count, n.Rejected, n.DueDate, n.NestedCount, n.Rotate, n.Status, n.MatStd, n.RowType, n.InfoTxt, n.TotWeight, n.Weight, n.Thick, n.Width, n.Length, n.PurCode, n.PathName, n.Ready, n.Source, n.Type, n.Mirror, n.ProdInfo, n.PartKey, n.AttributeKey, n.LogisticalKey, n.CenterOfGravityX, n.CenterOfGravityY, n.CenterOfGravityZ, n.ProjectName, n.AssemblySequence, n.Name, n.CustomerID, n.CustOrderNo, n.CustMark, n.Workphases, n.Info1, n.Info2, n.Info3, n.ChargeNo, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.SourceOrderNo, n.PartSide); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }
//
// // Update updates an OrderModel in the database.
// func (r *OrderRepository) Update(ctx context.Context, n *OrderModel) error {
// 	// update with primary key
// 	const sqlstr = `UPDATE dbo.nxorderline SET
// 		nxproductid = @p1, nxpartid = @p2, nxororderno = @p3, nxollineno = @p4, nxolparentname = @p5, nxolsection = @p6, nxolmrplineno = @p7, nxolordercount = @p8, nxolrejected = @p9, nxolduedate = @p10, nxolnestedcnt = @p11, nxolrotate = @p12, nxolstatus = @p13, nxolmatstd = @p14, nxolrowtype = @p15, nxolinfotxt = @p16, nxoltotweight = @p17, nxolweight = @p18, nxolthick = @p19, nxolwidth = @p20, nxollength = @p21, nxolpurcode = @p22, nxolnestixno = @p23, nxolready = @p24, nxolsource = @p25, nxoltype = @p26, nxolmirror = @p27, nxolprodinfo = @p28, nxolpartkey = @p29, nxolattributekey = @p30, nxollogisticalkey = @p31, nxolcenterofgravityx = @p32, nxolcenterofgravityy = @p33, nxolcenterofgravityz = @p34, nxolprojectname = @p35, nxolassemblysequence = @p36, nxolname = @p37, nxcustomerid = @p38, nxolcustorderno = @p39, nxolcustmark = @p40, nxolworkphases = @p41, nxolorderinfo1 = @p42, nxolorderinfo2 = @p43, nxolorderinfo3 = @p44, nxolchargeno = @p45, nxolinsertdate = @p46, nxolcreated = @p47, nxolcreator = @p48, nxolchanged = @p49, nxolchanger = @p50, nxolsourceorderno = @p51, nxolpartside = @p52
// 		WHERE nxorderlineid = @p53`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ProductID,
// 		n.PartID,
// 		n.OrderNo,
// 		n.LineNo,
// 		n.ParentName,
// 		n.Section,
// 		n.MrpLineNo,
// 		n.Count,
// 		n.Rejected,
// 		n.DueDate,
// 		n.NestedCount,
// 		n.Rotate,
// 		n.Status,
// 		n.MatStd,
// 		n.RowType,
// 		n.InfoTxt,
// 		n.TotWeight,
// 		n.Weight,
// 		n.Thick,
// 		n.Width,
// 		n.Length,
// 		n.PurCode,
// 		n.PathName,
// 		n.Ready,
// 		n.Source,
// 		n.Type,
// 		n.Mirror,
// 		n.ProdInfo,
// 		n.PartKey,
// 		n.AttributeKey,
// 		n.LogisticalKey,
// 		n.CenterOfGravityX,
// 		n.CenterOfGravityY,
// 		n.CenterOfGravityZ,
// 		n.ProjectName,
// 		n.AssemblySequence,
// 		n.Name,
// 		n.CustomerID,
// 		n.CustOrderNo,
// 		n.CustMark,
// 		n.Workphases,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.ChargeNo,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.SourceOrderNo,
// 		n.PartSide,
// 		n.ID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ProductID, n.PartID, n.OrderNo, n.LineNo, n.ParentName, n.Section, n.MrpLineNo, n.Count, n.Rejected, n.DueDate, n.NestedCount, n.Rotate, n.Status, n.MatStd, n.RowType, n.InfoTxt, n.TotWeight, n.Weight, n.Thick, n.Width, n.Length, n.PurCode, n.PathName, n.Ready, n.Source, n.Type, n.Mirror, n.ProdInfo, n.PartKey, n.AttributeKey, n.LogisticalKey, n.CenterOfGravityX, n.CenterOfGravityY, n.CenterOfGravityZ, n.ProjectName, n.AssemblySequence, n.Name, n.CustomerID, n.CustOrderNo, n.CustMark, n.Workphases, n.Info1, n.Info2, n.Info3, n.ChargeNo, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.SourceOrderNo, n.PartSide, n.ID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// // Upsert performs an upsert for OrderModel.
// func (r *OrderRepository) Upsert(ctx context.Context, n *OrderModel) error {
// 	// upsert
// 	const sqlstr = `MERGE dbo.nxorderline AS t
// 		USING (
// 		SELECT @p1 nxorderlineid, @p2 nxproductid, @p3 nxpartid, @p4 nxororderno, @p5 nxollineno, @p6 nxolparentname, @p7 nxolsection, @p8 nxolmrplineno, @p9 nxolordercount, @p10 nxolrejected, @p11 nxolduedate, @p12 nxolnestedcnt, @p13 nxolrotate, @p14 nxolstatus, @p15 nxolmatstd, @p16 nxolrowtype, @p17 nxolinfotxt, @p18 nxoltotweight, @p19 nxolweight, @p20 nxolthick, @p21 nxolwidth, @p22 nxollength, @p23 nxolpurcode, @p24 nxolnestixno, @p25 nxolready, @p26 nxolsource, @p27 nxoltype, @p28 nxolmirror, @p29 nxolprodinfo, @p30 nxolpartkey, @p31 nxolattributekey, @p32 nxollogisticalkey, @p33 nxolcenterofgravityx, @p34 nxolcenterofgravityy, @p35 nxolcenterofgravityz, @p36 nxolprojectname, @p37 nxolassemblysequence, @p38 nxolname, @p39 nxcustomerid, @p40 nxolcustorderno, @p41 nxolcustmark, @p42 nxolworkphases, @p43 nxolorderinfo1, @p44 nxolorderinfo2, @p45 nxolorderinfo3, @p46 nxolchargeno, @p47 nxolinsertdate, @p48 nxolcreated, @p49 nxolcreator, @p50 nxolchanged, @p51 nxolchanger, @p52 nxolsourceorderno, @p53 nxolpartside
// 		) AS s
// 		ON s.nxorderlineid = t.nxorderlineid
// 		WHEN MATCHED THEN
// 		UPDATE SET
// 		t.nxproductid = s.nxproductid, t.nxpartid = s.nxpartid, t.nxororderno = s.nxororderno, t.nxollineno = s.nxollineno, t.nxolparentname = s.nxolparentname, t.nxolsection = s.nxolsection, t.nxolmrplineno = s.nxolmrplineno, t.nxolordercount = s.nxolordercount, t.nxolrejected = s.nxolrejected, t.nxolduedate = s.nxolduedate, t.nxolnestedcnt = s.nxolnestedcnt, t.nxolrotate = s.nxolrotate, t.nxolstatus = s.nxolstatus, t.nxolmatstd = s.nxolmatstd, t.nxolrowtype = s.nxolrowtype, t.nxolinfotxt = s.nxolinfotxt, t.nxoltotweight = s.nxoltotweight, t.nxolweight = s.nxolweight, t.nxolthick = s.nxolthick, t.nxolwidth = s.nxolwidth, t.nxollength = s.nxollength, t.nxolpurcode = s.nxolpurcode, t.nxolnestixno = s.nxolnestixno, t.nxolready = s.nxolready, t.nxolsource = s.nxolsource, t.nxoltype = s.nxoltype, t.nxolmirror = s.nxolmirror, t.nxolprodinfo = s.nxolprodinfo, t.nxolpartkey = s.nxolpartkey, t.nxolattributekey = s.nxolattributekey, t.nxollogisticalkey = s.nxollogisticalkey, t.nxolcenterofgravityx = s.nxolcenterofgravityx, t.nxolcenterofgravityy = s.nxolcenterofgravityy, t.nxolcenterofgravityz = s.nxolcenterofgravityz, t.nxolprojectname = s.nxolprojectname, t.nxolassemblysequence = s.nxolassemblysequence, t.nxolname = s.nxolname, t.nxcustomerid = s.nxcustomerid, t.nxolcustorderno = s.nxolcustorderno, t.nxolcustmark = s.nxolcustmark, t.nxolworkphases = s.nxolworkphases, t.nxolorderinfo1 = s.nxolorderinfo1, t.nxolorderinfo2 = s.nxolorderinfo2, t.nxolorderinfo3 = s.nxolorderinfo3, t.nxolchargeno = s.nxolchargeno, t.nxolinsertdate = s.nxolinsertdate, t.nxolcreated = s.nxolcreated, t.nxolcreator = s.nxolcreator, t.nxolchanged = s.nxolchanged, t.nxolchanger = s.nxolchanger, t.nxolsourceorderno = s.nxolsourceorderno, t.nxolpartside = s.nxolpartside
// 		WHEN NOT MATCHED THEN
// 		INSERT (
// 		nxorderlineid, nxproductid, nxpartid, nxororderno, nxollineno, nxolparentname, nxolsection, nxolmrplineno, nxolordercount, nxolrejected, nxolduedate, nxolnestedcnt, nxolrotate, nxolstatus, nxolmatstd, nxolrowtype, nxolinfotxt, nxoltotweight, nxolweight, nxolthick, nxolwidth, nxollength, nxolpurcode, nxolnestixno, nxolready, nxolsource, nxoltype, nxolmirror, nxolprodinfo, nxolpartkey, nxolattributekey, nxollogisticalkey, nxolcenterofgravityx, nxolcenterofgravityy, nxolcenterofgravityz, nxolprojectname, nxolassemblysequence, nxolname, nxcustomerid, nxolcustorderno, nxolcustmark, nxolworkphases, nxolorderinfo1, nxolorderinfo2, nxolorderinfo3, nxolchargeno, nxolinsertdate, nxolcreated, nxolcreator, nxolchanged, nxolchanger, nxolsourceorderno, nxolpartside
// 		) VALUES (
// 		s.nxorderlineid, s.nxproductid, s.nxpartid, s.nxororderno, s.nxollineno, s.nxolparentname, s.nxolsection, s.nxolmrplineno, s.nxolordercount, s.nxolrejected, s.nxolduedate, s.nxolnestedcnt, s.nxolrotate, s.nxolstatus, s.nxolmatstd, s.nxolrowtype, s.nxolinfotxt, s.nxoltotweight, s.nxolweight, s.nxolthick, s.nxolwidth, s.nxollength, s.nxolpurcode, s.nxolnestixno, s.nxolready, s.nxolsource, s.nxoltype, s.nxolmirror, s.nxolprodinfo, s.nxolpartkey, s.nxolattributekey, s.nxollogisticalkey, s.nxolcenterofgravityx, s.nxolcenterofgravityy, s.nxolcenterofgravityz, s.nxolprojectname, s.nxolassemblysequence, s.nxolname, s.nxcustomerid, s.nxolcustorderno, s.nxolcustmark, s.nxolworkphases, s.nxolorderinfo1, s.nxolorderinfo2, s.nxolorderinfo3, s.nxolchargeno, s.nxolinsertdate, s.nxolcreated, s.nxolcreator, s.nxolchanged, s.nxolchanger, s.nxolsourceorderno, s.nxolpartside
// 		);`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.ProductID,
// 		n.PartID,
// 		n.OrderNo,
// 		n.LineNo,
// 		n.ParentName,
// 		n.Section,
// 		n.MrpLineNo,
// 		n.Count,
// 		n.Rejected,
// 		n.DueDate,
// 		n.NestedCount,
// 		n.Rotate,
// 		n.Status,
// 		n.MatStd,
// 		n.RowType,
// 		n.InfoTxt,
// 		n.TotWeight,
// 		n.Weight,
// 		n.Thick,
// 		n.Width,
// 		n.Length,
// 		n.PurCode,
// 		n.PathName,
// 		n.Ready,
// 		n.Source,
// 		n.Type,
// 		n.Mirror,
// 		n.ProdInfo,
// 		n.PartKey,
// 		n.AttributeKey,
// 		n.LogisticalKey,
// 		n.CenterOfGravityX,
// 		n.CenterOfGravityY,
// 		n.CenterOfGravityZ,
// 		n.ProjectName,
// 		n.AssemblySequence,
// 		n.Name,
// 		n.CustomerID,
// 		n.CustOrderNo,
// 		n.CustMark,
// 		n.Workphases,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.ChargeNo,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.SourceOrderNo,
// 		n.PartSide,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.ProductID, n.PartID, n.OrderNo, n.LineNo, n.ParentName, n.Section, n.MrpLineNo, n.Count, n.Rejected, n.DueDate, n.NestedCount, n.Rotate, n.Status, n.MatStd, n.RowType, n.InfoTxt, n.TotWeight, n.Weight, n.Thick, n.Width, n.Length, n.PurCode, n.PathName, n.Ready, n.Source, n.Type, n.Mirror, n.ProdInfo, n.PartKey, n.AttributeKey, n.LogisticalKey, n.CenterOfGravityX, n.CenterOfGravityY, n.CenterOfGravityZ, n.ProjectName, n.AssemblySequence, n.Name, n.CustomerID, n.CustOrderNo, n.CustMark, n.Workphases, n.Info1, n.Info2, n.Info3, n.ChargeNo, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.SourceOrderNo, n.PartSide); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// Delete deletes the OrderModel from the database.
func (r *OrderRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxorderline
		WHERE nxorderlineid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}
	return nil
}

// GetByID retrieves a row from 'dbo.nxorderline' as a OrderModel.
func (r *OrderRepository) GetByID(ctx context.Context, id int64) (*OrderModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxorderlineid, nxproductid, nxpartid, nxororderno, nxollineno, nxolparentname, nxolsection, nxolmrplineno, nxolordercount, nxolrejected, nxolduedate, nxolnestedcnt, nxolrotate, nxolstatus, nxolmatstd, nxolrowtype, nxolinfotxt, nxoltotweight, nxolweight, nxolthick, nxolwidth, nxollength, nxolpurcode, nxolnestixno, nxolready, nxolsource, nxoltype, nxolmirror, nxolprodinfo, nxolpartkey, nxolattributekey, nxollogisticalkey, nxolcenterofgravityx, nxolcenterofgravityy, nxolcenterofgravityz, nxolprojectname, nxolassemblysequence, nxolname, nxcustomerid, nxolcustorderno, nxolcustmark, nxolworkphases, nxolorderinfo1, nxolorderinfo2, nxolorderinfo3, nxolchargeno, nxolinsertdate, nxolcreated, nxolcreator, nxolchanged, nxolchanger, nxolsourceorderno, nxolpartside
		FROM dbo.nxorderline
		WHERE nxorderlineid = @p1`
	logf(tenantID, sqlstr, id)
	var n OrderModel

	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}

func (r *OrderRepository) GetByIDs(ctx context.Context, ids []int64) ([]*OrderModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxorderlineid, nxproductid, nxpartid, nxororderno, nxollineno, nxolparentname, nxolsection, nxolmrplineno, nxolordercount, nxolrejected, nxolduedate, nxolnestedcnt, nxolrotate, nxolstatus, nxolmatstd, nxolrowtype, nxolinfotxt, nxoltotweight, nxolweight, nxolthick, nxolwidth, nxollength, nxolpurcode, nxolnestixno, nxolready, nxolsource, nxoltype, nxolmirror, nxolprodinfo, nxolpartkey, nxolattributekey, nxollogisticalkey, nxolcenterofgravityx, nxolcenterofgravityy, nxolcenterofgravityz, nxolprojectname, nxolassemblysequence, nxolname, nxcustomerid, nxolcustorderno, nxolcustmark, nxolworkphases, nxolorderinfo1, nxolorderinfo2, nxolorderinfo3, nxolchargeno, nxolinsertdate, nxolcreated, nxolcreator, nxolchanged, nxolchanger, nxolsourceorderno, nxolpartside
		FROM dbo.nxorderline
		WHERE nxolrowtype = 0 AND nxorderlineid IN (?)`

	query, args, err := sqlx.In(sqlstr, ids)
	query = r.dbs[tenantID].Rebind(query)

	logf(tenantID, query, args)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*OrderModel
	for rows.Next() {
		var n OrderModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

// GetByID retrieves a row from 'dbo.nxorderline' as a OrderModel.
func (r *OrderRepository) List(ctx context.Context) ([]*OrderModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxorderlineid, nxproductid, nxpartid, nxororderno, nxollineno, nxolparentname, nxolsection, nxolmrplineno, nxolordercount, nxolrejected, nxolduedate, nxolnestedcnt, nxolrotate, nxolstatus, nxolmatstd, nxolrowtype, nxolinfotxt, nxoltotweight, nxolweight, nxolthick, nxolwidth, nxollength, nxolpurcode, nxolnestixno, nxolready, nxolsource, nxoltype, nxolmirror, nxolprodinfo, nxolpartkey, nxolattributekey, nxollogisticalkey, nxolcenterofgravityx, nxolcenterofgravityy, nxolcenterofgravityz, nxolprojectname, nxolassemblysequence, nxolname, nxcustomerid, nxolcustorderno, nxolcustmark, nxolworkphases, nxolorderinfo1, nxolorderinfo2, nxolorderinfo3, nxolchargeno, nxolinsertdate, nxolcreated, nxolcreator, nxolchanged, nxolchanger, nxolsourceorderno, nxolpartside
		FROM dbo.nxorderline
		WHERE nxolrowtype = 0`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*OrderModel
	for rows.Next() {
		var n OrderModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}
