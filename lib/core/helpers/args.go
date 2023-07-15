package helpers

import (
	"flag"
	"github.com/jedib0t/go-pretty/v6/table"
	"hogo/lib/core/log"
	"os"
)

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
	var verbosity = flag.String("verbosity", log.InfoLevel, "Level of detailed logging (default: info)")
	var logFormat = flag.String("log-format", log.TextFormatter, "Level of detailed logging (default: text)")

	flag.Parse()

	if *silentMode {
		*verbosity = log.WarningLevel
	}

	if *devMode {
		*verbosity = log.DebugLevel
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Arg", "Value"})
	t.AppendRows([]table.Row{
		[]interface{}{"listen", *listen},
		[]interface{}{"metrics", *metricsEnabled},
		[]interface{}{"silent", *silentMode},
		[]interface{}{"dev", *devMode},
		[]interface{}{"verbosity", *verbosity},
		[]interface{}{"log-format", *logFormat},
	})
	t.Render()

	return Args{
		ListenAt:       *listen,
		MetricsEnabled: *metricsEnabled,
		SilentMode:     *silentMode,
		DevMode:        *devMode,
		Verbosity:      *verbosity,
		LogFormat:      *logFormat,
	}
}
