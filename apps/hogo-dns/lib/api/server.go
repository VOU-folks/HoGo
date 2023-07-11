package api

import (
	"flag"
	"log"
	"sync"

	"github.com/gin-gonic/gin"

	"hogo/apps/hogo-dns/lib/http"
	"hogo/apps/hogo-dns/lib/http/methods"
	"hogo/apps/hogo-dns/lib/metrics"
)

type Server struct {
	apiServer      *gin.Engine
	listenAt       string
	metricsEnabled bool
	silentMode     bool
	mx             sync.Mutex
}

func CreateServer() *Server {
	var listen = flag.String("listen", "0.0.0.0:8053", "Address to listen at (default 0.0.0.0:80)")
	var metricsEnabled = flag.Bool("metrics", false, "Enable metrics (default false)")
	var silentMode = flag.Bool("silent", true, "Disable verbosity, log only errors (default true)")
	flag.Parse()

	log.Println("listen =", *listen)
	log.Println("metrics =", *metricsEnabled)
	log.Println("silent =", *silentMode)

	apiInstance := createApiInstance(*metricsEnabled, *silentMode)

	return &Server{
		apiServer:      apiInstance,
		listenAt:       *listen,
		metricsEnabled: *metricsEnabled,
		silentMode:     *silentMode,
		mx:             sync.Mutex{},
	}
}

func createApiInstance(withMetrics bool, silentMode bool) *gin.Engine {
	if silentMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.Use(gin.Recovery())

	if withMetrics == true {
		log.Println("Attaching metrics monitor")
		monitor := metrics.GetMonitor()
		monitor.Use(app)
	}

	app.NoMethod(methods.NotFound)

	return app
}

func (s *Server) ListenAndServe() {
	httpServer := http.CreateHttpServer(s.listenAt)
	httpServer.Handler = s.apiServer.Handler()

	log.Println("HoGo DNS HTTP/API Component listening at", s.listenAt)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("httpServer.ListenAndServe:", err.Error())
	}
}

func (s *Server) Close() {
	err := s.apiServer.Close()
	if err != nil {
		log.Fatal("dnsService.Listener:", err.Error())
	}
}
