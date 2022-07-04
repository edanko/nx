package adapters

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

// ProductModel represents a row from 'dbo.nxproduct'.
type ProductModel struct {
	ID              int64           `db:"nxproductid"`
	VersNo          sql.NullString  `db:"nxprversno"`
	PartNo          sql.NullString  `db:"nxprpartno"`
	ProductNo       string          `db:"nxproductno"`
	Type            sql.NullInt64   `db:"nxprtype"`
	Name            sql.NullString  `db:"nxprname"`
	Length          sql.NullFloat64 `db:"nxprlength"`
	Width           sql.NullFloat64 `db:"nxprwidth"`
	Thick           sql.NullFloat64 `db:"nxprthick"`
	MatGroup        sql.NullString  `db:"nxprmatgroup"`
	Quality         sql.NullString  `db:"nxprquality"`
	Density         sql.NullFloat64 `db:"nxprdensity"`
	UseDate         sql.NullTime    `db:"nxprusedate"`
	Site            sql.NullString  `db:"nxprsite"`
	Meta            sql.NullString  `db:"nxprmeta"`
	Checksum        sql.NullInt64   `db:"nxprchecksum"`
	PowderLen       sql.NullFloat64 `db:"nxprpowderlen"`
	ProfCnt         sql.NullInt64   `db:"nxprprofcnt"`
	CutLen          sql.NullFloat64 `db:"nxprcutlen"`
	Area            sql.NullFloat64 `db:"nxprarea"`
	Drawed          sql.NullInt64   `db:"nxprdrawed"`
	PointMark       sql.NullInt64   `db:"nxprpointmark"`
	MinRAng         sql.NullFloat64 `db:"nxprminrang"`
	MinRLen         sql.NullFloat64 `db:"nxprminrlen"`
	MinRWidth       sql.NullFloat64 `db:"nxprminrwidth"`
	InfoTxt         sql.NullString  `db:"nxprinfotxt"`
	Weight          sql.NullFloat64 `db:"nxprweight"`
	WeightUnit      sql.NullInt64   `db:"nxprweightunit"`
	Filename        sql.NullString  `db:"nxprfilename"`
	Attributes      []byte          `db:"nxprattributes"`
	InnerArea       sql.NullFloat64 `db:"nxprinnerarea"`
	TextCnt         sql.NullInt64   `db:"nxprtextcnt"`
	Mirroring       sql.NullInt64   `db:"nxprmirroring"`
	Macro           sql.NullString  `db:"nxprmacro"`
	ToolInfos       []byte          `db:"nxprtoolinfos"`
	Assembly        sql.NullString  `db:"nxprassembly"`
	Info1           sql.NullString  `db:"nxprinfo1"`
	Info2           sql.NullString  `db:"nxprinfo2"`
	Info3           sql.NullString  `db:"nxprinfo3"`
	MacroData       []byte          `db:"nxprmacrodata"`
	InsertDate      time.Time       `db:"nxprinsertdate"`
	Created         sql.NullTime    `db:"nxprcreated"`
	Creator         sql.NullString  `db:"nxprcreator"`
	Changed         sql.NullTime    `db:"nxprchanged"`
	Changer         sql.NullString  `db:"nxprchanger"`
	ChangeType      sql.NullInt64   `db:"nxprchangetype"`
	SectionCopyData []byte          `db:"nxprsectioncopydata"`
	PartChecksum    sql.NullString  `db:"nxprpartchecksum"`
	ProjectNo       sql.NullString  `db:"nxprprojectno"`
	DrawnSide       sql.NullInt64   `db:"nxprdrawnside"`
	OwnerID         sql.NullInt64   `db:"nxownerid"`
}

type ProductRepository struct {
	dbs map[string]*sqlx.DB
}

func NewProductRepository(db map[string]*sqlx.DB) *ProductRepository {
	return &ProductRepository{
		dbs: db,
	}
}

//
// // Insert inserts the ProductModel to the database.
// func (r *ProductRepository) Insert(ctx context.Context, n *ProductModel) error {
// 	// insert (manual)
// 	const sqlstr = `INSERT INTO dbo.nxproduct (` +
// 		`nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid` +
// 		`) VALUES (` +
// 		`@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30, @p31, @p32, @p33, @p34, @p35, @p36, @p37, @p38, @p39, @p40, @p41, @p42, @p43, @p44, @p45, @p46, @p47, @p48, @p49, @p50, @p51` +
// 		`)`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.VersNo,
// 		n.PartNo,
// 		n.ProductNo,
// 		n.Type,
// 		n.Name,
// 		n.Length,
// 		n.Width,
// 		n.Thick,
// 		n.MatGroup,
// 		n.Quality,
// 		n.Density,
// 		n.UseDate,
// 		n.Site,
// 		n.Meta,
// 		n.Checksum,
// 		n.PowderLen,
// 		n.ProfCnt,
// 		n.CutLen,
// 		n.Area,
// 		n.Drawed,
// 		n.PointMark,
// 		n.MinRAng,
// 		n.MinRLen,
// 		n.MinRWidth,
// 		n.InfoTxt,
// 		n.Weight,
// 		n.WeightUnit,
// 		n.Filename,
// 		n.Attributes,
// 		n.InnerArea,
// 		n.TextCnt,
// 		n.Mirroring,
// 		n.Macro,
// 		n.ToolInfos,
// 		n.Assembly,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.MacroData,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.ChangeType,
// 		n.SectionCopyData,
// 		n.PartChecksum,
// 		n.ProjectNo,
// 		n.DrawnSide,
// 		n.OwnerID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.VersNo, n.PartNo, n.ProductNo, n.Type, n.Name, n.Length, n.Width, n.Thick, n.MatGroup, n.Quality, n.Density, n.UseDate, n.Site, n.Meta, n.Checksum, n.PowderLen, n.ProfCnt, n.CutLen, n.Area, n.Drawed, n.PointMark, n.MinRAng, n.MinRLen, n.MinRWidth, n.InfoTxt, n.Weight, n.WeightUnit, n.Filename, n.Attributes, n.InnerArea, n.TextCnt, n.Mirroring, n.Macro, n.ToolInfos, n.Assembly, n.Info1, n.Info2, n.Info3, n.MacroData, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.ChangeType, n.SectionCopyData, n.PartChecksum, n.ProjectNo, n.DrawnSide, n.OwnerID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }
//
// // Update updates a ProductModel in the database.
// func (r *ProductRepository) Update(ctx context.Context, n *ProductModel) error {
// 	// update with primary key
// 	const sqlstr = `UPDATE dbo.nxproduct SET ` +
// 		`nxprversno = @p1, nxprpartno = @p2, nxproductno = @p3, nxprtype = @p4, nxprname = @p5, nxprlength = @p6, nxprwidth = @p7, nxprthick = @p8, nxprmatgroup = @p9, nxprquality = @p10, nxprdensity = @p11, nxprusedate = @p12, nxprsite = @p13, nxprmeta = @p14, nxprchecksum = @p15, nxprpowderlen = @p16, nxprprofcnt = @p17, nxprcutlen = @p18, nxprarea = @p19, nxprdrawed = @p20, nxprpointmark = @p21, nxprminrang = @p22, nxprminrlen = @p23, nxprminrwidth = @p24, nxprinfotxt = @p25, nxprweight = @p26, nxprweightunit = @p27, nxprfilename = @p28, nxprattributes = @p29, nxprinnerarea = @p30, nxprtextcnt = @p31, nxprmirroring = @p32, nxprmacro = @p33, nxprtoolinfos = @p34, nxprassembly = @p35, nxprinfo1 = @p36, nxprinfo2 = @p37, nxprinfo3 = @p38, nxprmacrodata = @p39, nxprinsertdate = @p40, nxprcreated = @p41, nxprcreator = @p42, nxprchanged = @p43, nxprchanger = @p44, nxprchangetype = @p45, nxprsectioncopydata = @p46, nxprpartchecksum = @p47, nxprprojectno = @p48, nxprdrawnside = @p49, nxownerid = @p50 ` +
// 		`WHERE nxproductid = @p51`
// 	logf(tenantID,
// 		sqlstr,
// 		n.VersNo,
// 		n.PartNo,
// 		n.ProductNo,
// 		n.Type,
// 		n.Name,
// 		n.Length,
// 		n.Width,
// 		n.Thick,
// 		n.MatGroup,
// 		n.Quality,
// 		n.Density,
// 		n.UseDate,
// 		n.Site,
// 		n.Meta,
// 		n.Checksum,
// 		n.PowderLen,
// 		n.ProfCnt,
// 		n.CutLen,
// 		n.Area,
// 		n.Drawed,
// 		n.PointMark,
// 		n.MinRAng,
// 		n.MinRLen,
// 		n.MinRWidth,
// 		n.InfoTxt,
// 		n.Weight,
// 		n.WeightUnit,
// 		n.Filename,
// 		n.Attributes,
// 		n.InnerArea,
// 		n.TextCnt,
// 		n.Mirroring,
// 		n.Macro,
// 		n.ToolInfos,
// 		n.Assembly,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.MacroData,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.ChangeType,
// 		n.SectionCopyData,
// 		n.PartChecksum,
// 		n.ProjectNo,
// 		n.DrawnSide,
// 		n.OwnerID,
// 		n.ID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.VersNo, n.PartNo, n.ProductNo, n.Type, n.Name, n.Length, n.Width, n.Thick, n.MatGroup, n.Quality, n.Density, n.UseDate, n.Site, n.Meta, n.Checksum, n.PowderLen, n.ProfCnt, n.CutLen, n.Area, n.Drawed, n.PointMark, n.MinRAng, n.MinRLen, n.MinRWidth, n.InfoTxt, n.Weight, n.WeightUnit, n.Filename, n.Attributes, n.InnerArea, n.TextCnt, n.Mirroring, n.Macro, n.ToolInfos, n.Assembly, n.Info1, n.Info2, n.Info3, n.MacroData, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.ChangeType, n.SectionCopyData, n.PartChecksum, n.ProjectNo, n.DrawnSide, n.OwnerID, n.ID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// Save saves the ProductModel to the database.
// func (r *ProductRepository) Save(ctx context.Context, n *ProductModel) error {
// 	if n.Exists() {
// 		return n.Update(ctx, db)
// 	}
// 	return n.Insert(ctx, db)
// }
//
// // Upsert performs an upsert for ProductModel.
// func (r *ProductRepository) Upsert(ctx context.Context, n *ProductModel) error {
// 	// upsert
// 	const sqlstr = `MERGE dbo.nxproduct AS t ` +
// 		`USING (` +
// 		`SELECT @p1 nxproductid, @p2 nxprversno, @p3 nxprpartno, @p4 nxproductno, @p5 nxprtype, @p6 nxprname, @p7 nxprlength, @p8 nxprwidth, @p9 nxprthick, @p10 nxprmatgroup, @p11 nxprquality, @p12 nxprdensity, @p13 nxprusedate, @p14 nxprsite, @p15 nxprmeta, @p16 nxprchecksum, @p17 nxprpowderlen, @p18 nxprprofcnt, @p19 nxprcutlen, @p20 nxprarea, @p21 nxprdrawed, @p22 nxprpointmark, @p23 nxprminrang, @p24 nxprminrlen, @p25 nxprminrwidth, @p26 nxprinfotxt, @p27 nxprweight, @p28 nxprweightunit, @p29 nxprfilename, @p30 nxprattributes, @p31 nxprinnerarea, @p32 nxprtextcnt, @p33 nxprmirroring, @p34 nxprmacro, @p35 nxprtoolinfos, @p36 nxprassembly, @p37 nxprinfo1, @p38 nxprinfo2, @p39 nxprinfo3, @p40 nxprmacrodata, @p41 nxprinsertdate, @p42 nxprcreated, @p43 nxprcreator, @p44 nxprchanged, @p45 nxprchanger, @p46 nxprchangetype, @p47 nxprsectioncopydata, @p48 nxprpartchecksum, @p49 nxprprojectno, @p50 nxprdrawnside, @p51 nxownerid ` +
// 		`) AS s ` +
// 		`ON s.nxproductid = t.nxproductid ` +
// 		`WHEN MATCHED THEN ` +
// 		`UPDATE SET ` +
// 		`t.nxprversno = s.nxprversno, t.nxprpartno = s.nxprpartno, t.nxproductno = s.nxproductno, t.nxprtype = s.nxprtype, t.nxprname = s.nxprname, t.nxprlength = s.nxprlength, t.nxprwidth = s.nxprwidth, t.nxprthick = s.nxprthick, t.nxprmatgroup = s.nxprmatgroup, t.nxprquality = s.nxprquality, t.nxprdensity = s.nxprdensity, t.nxprusedate = s.nxprusedate, t.nxprsite = s.nxprsite, t.nxprmeta = s.nxprmeta, t.nxprchecksum = s.nxprchecksum, t.nxprpowderlen = s.nxprpowderlen, t.nxprprofcnt = s.nxprprofcnt, t.nxprcutlen = s.nxprcutlen, t.nxprarea = s.nxprarea, t.nxprdrawed = s.nxprdrawed, t.nxprpointmark = s.nxprpointmark, t.nxprminrang = s.nxprminrang, t.nxprminrlen = s.nxprminrlen, t.nxprminrwidth = s.nxprminrwidth, t.nxprinfotxt = s.nxprinfotxt, t.nxprweight = s.nxprweight, t.nxprweightunit = s.nxprweightunit, t.nxprfilename = s.nxprfilename, t.nxprattributes = s.nxprattributes, t.nxprinnerarea = s.nxprinnerarea, t.nxprtextcnt = s.nxprtextcnt, t.nxprmirroring = s.nxprmirroring, t.nxprmacro = s.nxprmacro, t.nxprtoolinfos = s.nxprtoolinfos, t.nxprassembly = s.nxprassembly, t.nxprinfo1 = s.nxprinfo1, t.nxprinfo2 = s.nxprinfo2, t.nxprinfo3 = s.nxprinfo3, t.nxprmacrodata = s.nxprmacrodata, t.nxprinsertdate = s.nxprinsertdate, t.nxprcreated = s.nxprcreated, t.nxprcreator = s.nxprcreator, t.nxprchanged = s.nxprchanged, t.nxprchanger = s.nxprchanger, t.nxprchangetype = s.nxprchangetype, t.nxprsectioncopydata = s.nxprsectioncopydata, t.nxprpartchecksum = s.nxprpartchecksum, t.nxprprojectno = s.nxprprojectno, t.nxprdrawnside = s.nxprdrawnside, t.nxownerid = s.nxownerid ` +
// 		`WHEN NOT MATCHED THEN ` +
// 		`INSERT (` +
// 		`nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid` +
// 		`) VALUES (` +
// 		`s.nxproductid, s.nxprversno, s.nxprpartno, s.nxproductno, s.nxprtype, s.nxprname, s.nxprlength, s.nxprwidth, s.nxprthick, s.nxprmatgroup, s.nxprquality, s.nxprdensity, s.nxprusedate, s.nxprsite, s.nxprmeta, s.nxprchecksum, s.nxprpowderlen, s.nxprprofcnt, s.nxprcutlen, s.nxprarea, s.nxprdrawed, s.nxprpointmark, s.nxprminrang, s.nxprminrlen, s.nxprminrwidth, s.nxprinfotxt, s.nxprweight, s.nxprweightunit, s.nxprfilename, s.nxprattributes, s.nxprinnerarea, s.nxprtextcnt, s.nxprmirroring, s.nxprmacro, s.nxprtoolinfos, s.nxprassembly, s.nxprinfo1, s.nxprinfo2, s.nxprinfo3, s.nxprmacrodata, s.nxprinsertdate, s.nxprcreated, s.nxprcreator, s.nxprchanged, s.nxprchanger, s.nxprchangetype, s.nxprsectioncopydata, s.nxprpartchecksum, s.nxprprojectno, s.nxprdrawnside, s.nxownerid` +
// 		`);`
// 	logf(tenantID,
// 		sqlstr,
// 		n.ID,
// 		n.VersNo,
// 		n.PartNo,
// 		n.ProductNo,
// 		n.Type,
// 		n.Name,
// 		n.Length,
// 		n.Width,
// 		n.Thick,
// 		n.MatGroup,
// 		n.Quality,
// 		n.Density,
// 		n.UseDate,
// 		n.Site,
// 		n.Meta,
// 		n.Checksum,
// 		n.PowderLen,
// 		n.ProfCnt,
// 		n.CutLen,
// 		n.Area,
// 		n.Drawed,
// 		n.PointMark,
// 		n.MinRAng,
// 		n.MinRLen,
// 		n.MinRWidth,
// 		n.InfoTxt,
// 		n.Weight,
// 		n.WeightUnit,
// 		n.Filename,
// 		n.Attributes,
// 		n.InnerArea,
// 		n.TextCnt,
// 		n.Mirroring,
// 		n.Macro,
// 		n.ToolInfos,
// 		n.Assembly,
// 		n.Info1,
// 		n.Info2,
// 		n.Info3,
// 		n.MacroData,
// 		n.InsertDate,
// 		n.Created,
// 		n.Creator,
// 		n.Changed,
// 		n.Changer,
// 		n.ChangeType,
// 		n.SectionCopyData,
// 		n.PartChecksum,
// 		n.ProjectNo,
// 		n.DrawnSide,
// 		n.OwnerID,
// 	)
// 	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID, n.VersNo, n.PartNo, n.ProductNo, n.Type, n.Name, n.Length, n.Width, n.Thick, n.MatGroup, n.Quality, n.Density, n.UseDate, n.Site, n.Meta, n.Checksum, n.PowderLen, n.ProfCnt, n.CutLen, n.Area, n.Drawed, n.PointMark, n.MinRAng, n.MinRLen, n.MinRWidth, n.InfoTxt, n.Weight, n.WeightUnit, n.Filename, n.Attributes, n.InnerArea, n.TextCnt, n.Mirroring, n.Macro, n.ToolInfos, n.Assembly, n.Info1, n.Info2, n.Info3, n.MacroData, n.InsertDate, n.Created, n.Creator, n.Changed, n.Changer, n.ChangeType, n.SectionCopyData, n.PartChecksum, n.ProjectNo, n.DrawnSide, n.OwnerID); err != nil {
// 		return logerror(err)
// 	}
// 	return nil
// }

// Delete deletes the ProductModel from the database.
func (r *ProductRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxproduct ` +
		`WHERE nxproductid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}
	return nil
}

// GetByID retrieves a row from 'dbo.nxproduct' as a ProductModel.
func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT ` +
		`nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid ` +
		`FROM dbo.nxproduct ` +
		`WHERE nxproductid = @p1`
	logf(tenantID, sqlstr, id)
	n := ProductModel{}
	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}

func (r *ProductRepository) GetByIDs(ctx context.Context, ids []int64) ([]*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid
		FROM dbo.nxproduct
		WHERE nxproductid IN (?)`

	query, args, err := sqlx.In(sqlstr, ids)
	query = r.dbs[tenantID].Rebind(query)

	logf(tenantID, query, ids)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*ProductModel
	for rows.Next() {
		var n ProductModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

func (r *ProductRepository) ListArticles(ctx context.Context) ([]*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid
		FROM dbo.nxproduct
		WHERE nxprtype = 1`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*ProductModel
	for rows.Next() {
		var n ProductModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

func (r *ProductRepository) ListDrawnParts(ctx context.Context) ([]*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid
		FROM dbo.nxproduct
		WHERE nxprtype = 9 AND nxprdrawed = 1`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*ProductModel
	for rows.Next() {
		var n ProductModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

func (r *ProductRepository) ListNeedParts(ctx context.Context) ([]*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid
		FROM dbo.nxproduct
		WHERE nxprtype = 9 AND nxprdrawed = 0`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*ProductModel
	for rows.Next() {
		var n ProductModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

func (r *ProductRepository) ListChangedParts(ctx context.Context) ([]*ProductModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxproductid, nxprversno, nxprpartno, nxproductno, nxprtype, nxprname, nxprlength, nxprwidth, nxprthick, nxprmatgroup, nxprquality, nxprdensity, nxprusedate, nxprsite, nxprmeta, nxprchecksum, nxprpowderlen, nxprprofcnt, nxprcutlen, nxprarea, nxprdrawed, nxprpointmark, nxprminrang, nxprminrlen, nxprminrwidth, nxprinfotxt, nxprweight, nxprweightunit, nxprfilename, nxprattributes, nxprinnerarea, nxprtextcnt, nxprmirroring, nxprmacro, nxprtoolinfos, nxprassembly, nxprinfo1, nxprinfo2, nxprinfo3, nxprmacrodata, nxprinsertdate, nxprcreated, nxprcreator, nxprchanged, nxprchanger, nxprchangetype, nxprsectioncopydata, nxprpartchecksum, nxprprojectno, nxprdrawnside, nxownerid
		FROM dbo.nxproduct
		WHERE nxprchangetype > 0`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*ProductModel
	for rows.Next() {
		var n ProductModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}
