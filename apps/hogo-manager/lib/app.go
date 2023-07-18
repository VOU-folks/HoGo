package lib

import (
	"github.com/gin-gonic/gin"
	"hogo/apps/hogo-manager/lib/routes"
	. "hogo/lib/core/db/common/sqlite"
	. "hogo/lib/core/helpers"
	"hogo/lib/core/http/handlers"
	"hogo/lib/core/interfaces"
	"hogo/lib/core/log"
	"hogo/lib/core/metrics"
	"net/http"
)

type HogoManagerApp struct {
	stopped       bool
	stopRequested bool

	args         Args
	dbConnection interfaces.DBConnection
	instance     *gin.Engine
}

func (a *HogoManagerApp) Init(args Args) {
	a.args = args
	a.stopped = false
	a.stopRequested = false

	a.dbConnection = SQLiteConnector{}
	err := a.dbConnection.Open()
	if err != nil {
		log.Fatal("db.Open:", err.Error())
	}
	log.Info("DB connection established")

	if a.args.SilentMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if a.args.DevMode {
		gin.SetMode(gin.DebugMode)
	}

	a.instance = gin.New()
	a.instance.Use(gin.Recovery())

	if args.MetricsEnabled == true {
		monitor := metrics.GetMonitor()
		monitor.Use(a.instance)
		log.Info("Metrics monitor attached")
	}

	a.instance.NoMethod(handlers.NotFound)
	routes.Attach(a.instance)
}

func (a *HogoManagerApp) BindToHttpServer(server *http.Server) {
	server.Handler = a.instance
}

func (a *HogoManagerApp) Stop() {
	if a.stopRequested {
		return
	}
	a.stopRequested = true

	if a.dbConnection != nil {
		err := a.dbConnection.Close()
		if err != nil {
			log.Error("db.Close:", err.Error())
		}
	}

	a.stopped = true
}

func (a *HogoManagerApp) IsStopped() bool {
	return a.IsStopped()
}
