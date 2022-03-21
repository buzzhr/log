package defaultlog

import "log"

func SetDefaultLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
