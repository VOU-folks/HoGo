package helpers

import (
	"flag"
	"hogo/lib/core/log"
)

var GlobalArgs Args

type Args struct {
	ListenAt       string
	MetricsEnabled bool
	SilentMode     bool
	DevMode        bool
	Verbosity      string
	LogFormat      string
}

func GetArgs() Args {
	var listen = flag.String("listen", "127.0.0.1:8000", "IP:PORT combination to listen")
	var metricsEnabled = flag.Bool("metrics", false, "Enable metrics (optional, default: false), currently not implemented")
	var silentMode = flag.Bool("silent", true, "Disable verbosity, log only errors (optional, default: true)")
	var devMode = flag.Bool("dev", false, "Maximum verbosity, use for development (optional, default: false)")
	var verbosity = flag.String("verbosity", log.WarningLevel, "Level of detailed logging (default: info)")
	var logFormat = flag.String("log-format", log.TextFormatter, "Level of detailed logging (default: text)")

	flag.Parse()

	if *silentMode {
		*verbosity = log.ErrorLevel
	}

	if *devMode {
		*verbosity = log.DebugLevel
	}

	log.Println("listen =", *listen)
	log.Println("metrics =", *metricsEnabled)
	log.Println("silent =", *silentMode)
	log.Println("dev =", *devMode)
	log.Println("verbosity =", *verbosity)
	log.Println("log-format =", *logFormat)

	GlobalArgs = Args{
		ListenAt:       *listen,
		MetricsEnabled: *metricsEnabled,
		SilentMode:     *silentMode,
		DevMode:        *devMode,
		Verbosity:      *verbosity,
		LogFormat:      *logFormat,
	}

	return GlobalArgs
}
