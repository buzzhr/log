package log

import "log"

func InitLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
