package globals

import "embed"

type InitArgs struct {
	EmbeddedFiles embed.FS
}

func Init(args InitArgs) {
	ReadAppInfo(args.EmbeddedFiles)
	ReadArgs()
}
