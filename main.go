package main

import (
	kLog "github.com/go-kit/log"
	"os"
)

func main() {
	w := kLog.NewSyncWriter(os.Stderr)
	logger := kLog.NewLogfmtLogger(w)
	logger.Log("question", "what is the meaing of life?")
}
