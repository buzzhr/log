package defaultlog

import "log"

func InitLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
