package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/edanko/nx/cmd/nestix-api/internal/adapters"
	"github.com/edanko/nx/cmd/nestix-api/internal/config"
	"github.com/edanko/nx/cmd/nestix-api/internal/service"
	"github.com/edanko/nx/cmd/nestix-api/pkg/tenant"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	loggerOutput := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
	log.Logger = zerolog.New(loggerOutput).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Str("app", "nestix-api").
		Logger()

	cfg := config.GetConfig()

	if cfg.App.Environment == "development" {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	}

	connMap := make(map[string]*sqlx.DB, 10)

	for name, t := range cfg.Tenants {
		conn, err := sqlx.ConnectContext(
			ctx,
			"sqlserver",
			fmt.Sprintf(
				"sqlserver://%s:%s@%s?database=%s",
				t.DB.User,
				t.DB.Password,
				t.DB.Host,
				t.DB.Database,
			),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("error connecting to database")
		}
		defer conn.Close()
		connMap[name] = conn
	}

	// sqlLogger := func(s string, v ...any) {
	// 	log.Info().Str("query", s).Interface("args", v).Msg("")
	// }
	// adapters.SetLogger(sqlLogger)

	// sqlErrorLogger := func(s string, v ...any) {
	// 	log.Error().Interface("args", v).Msg("")
	// }
	// adapters.SetErrorLogger(sqlErrorLogger)

	pathRepo := adapters.NewPathRepository(connMap)
	sheetPathRepo := adapters.NewSheetPathRepository(connMap)
	sheetPathDetRepo := adapters.NewSheetPathDetRepository(connMap)
	orderRepo := adapters.NewOrderRepository(connMap)
	productRepo := adapters.NewProductRepository(connMap)
	visualRepo := adapters.NewVisualRepository(connMap)
	machineRepo := adapters.NewMachineRepository(connMap)

	svc := service.New(
		pathRepo,
		sheetPathRepo,
		sheetPathDetRepo,
		orderRepo,
		productRepo,
		visualRepo,
		machineRepo,
	)

	svc.Something(ctx)

	ctx = tenant.ContextWithTenantID(ctx, "IBSV")

	paths, err := pathRepo.SearchByName(ctx, "10-562003")
	if err != nil {
		log.Fatal().Err(err).Msg("search by name failed")
	}
	pathsIDs := lo.Map[*adapters.PathModel, int64](paths, func(x *adapters.PathModel, _ int) int64 {
		return x.ID
	})

	sheetpaths, _ := sheetPathRepo.GetByPathIDs(ctx, pathsIDs)

	var details [][]*adapters.SheetPathDetModel

	for _, sheetpath := range sheetpaths {
		sheetpathdet, err := sheetPathDetRepo.ListPartsBySheetPathID(ctx, sheetpath.ID)
		if err != nil {
			log.Fatal().Err(err).Msg("4get failed")
		}
		details = append(details, sheetpathdet)
	}

	var pathid int64 = 16901 // 20631

	path, err := pathRepo.GetByID(ctx, pathid)
	if err != nil {
		log.Fatal().Err(err).Msg("2get failed")
	}

	fmt.Println("nxname", path.Name)

	sheetpath, err := sheetPathRepo.GetByPathID(ctx, pathid)
	if err != nil {
		log.Fatal().Err(err).Msg("3get failed")
	}

	sheetpathdet, err := sheetPathDetRepo.ListPartsBySheetPathID(ctx, sheetpath.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("4get failed")
	}

	orderIDs := lo.Map[*adapters.SheetPathDetModel, int64](sheetpathdet, func(x *adapters.SheetPathDetModel, _ int) int64 {
		if !x.OrderlineID.Valid {
			panic("orderline id is not valid")
		}
		return x.OrderlineID.Int64
	})

	// spew.Dump(orderIDs)

	orders, _ := orderRepo.GetByIDs(ctx, orderIDs)
	_ = orders

	// for i, order := range orders {
	// 	fmt.Println("idx", i+1, "section", order.Section.String)
	// }

	products, _ := productRepo.GetByIDs(ctx, orderIDs)
	_ = products

	for i := range orderIDs {

		// for _, n := range sheetpathdet {
		// 	if !n.OrderlineID.Valid {
		// 		log.Info().Int64("id", n.ID).Msg("CHECK THIS")
		// 		continue
		// 	}
		// 	orderline, err := orderRepo.GetByID(ctx, n.OrderlineID.Int64)
		// 	if err != nil {
		// 		log.Fatal().Err(err).Msg("5get failed")
		// 	}
		n := sheetpathdet[i]

		orderline := orders[i]

		fmt.Println("detail code", n.DetailCode)
		fmt.Println("order", orderline.OrderNo.String)
		fmt.Println("section", orderline.Section.String)
		// product, err := productRepo.GetByID(ctx, orderline.PartID.Int64)
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("6get failed")
		// }
		product := products[i]
		fmt.Println("pos", product.PartNo.String)
		q := n.DetailCount.Int64 * orderline.Count
		fmt.Println("quantity", q)

		w := n.Area.Float64 * orderline.Thick.Float64 * product.Density.Float64
		fmt.Println("weight", w)
		fmt.Println("total weight", float64(q)*w)
		fmt.Println("quality", product.Quality.String)
		fmt.Println("len", product.Length.Float64)
		fmt.Println("wid", product.Width.Float64)

		fmt.Println("thickness", product.Thick.Float64)
		fmt.Println()
	}

	// mr := adapters.NewMasterRepository(cfg.Tenants["MR"].Master, cfg.Tenants["MR"].Site)
	// n, err := mr.ReadNest("15800")
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("read nest failed")
	// }
	//
	// spew.Dump(n)
}
