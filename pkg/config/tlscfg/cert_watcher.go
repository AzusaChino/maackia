package tlscfg

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type certWatcher struct {
	opts    Options
	watcher *fsnotify.Watcher
	cert    *tls.Certificate
	logger  *zap.Logger
	mu      *sync.RWMutex
}

var _ io.Closer = (*certWatcher)(nil)

// constructor
func newCertWatcher(opts Options, logger *zap.Logger) (*certWatcher, error) {

	return nil, nil
}

func (w *certWatcher) Close() error {
	return w.watcher.Close()
}

func (w *certWatcher) certificate() *tls.Certificate {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.cert
}

func (w *certWatcher) watchChangesLoop(rootCAs, clientCAs *x509.CertPool) {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			// ignore if the event is chmod
			if event.Op&fsnotify.Chmod == fsnotify.Chmod {
				continue
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				w.logger.Warn("Certificate has been removed, using the last known version",
					zap.String("certificate", event.Name))
				continue
			}
			w.logger.Info("Loading modified certificate",
				zap.String("certificate", event.Name),
				zap.String("event", event.Op.String()))
			var err error
			switch event.Name {
			case w.opts.CAPath:
				err = addCertToPool(w.opts.CAPath, rootCAs)
			case w.opts.ClientCAPath:
				err = addCertToPool(w.opts.ClientCAPath, clientCAs)
			case w.opts.CertPath, w.opts.KeyPath:
				w.mu.Lock()
				c, e := tls.LoadX509KeyPair(filepath.Clean(w.opts.CertPath), filepath.Clean(w.opts.KeyPath))
				if e == nil {
					w.cert = &c
				}
				w.mu.Unlock()
				err = e
			}
			if err == nil {
				w.logger.Info("Loaded modified certificate",
					zap.String("certificate", event.Name),
					zap.String("event", event.Op.String()))

			} else {
				w.logger.Error("Failed to load certificate",
					zap.String("certificate", event.Name),
					zap.String("event", event.Op.String()),
					zap.Error(err))
			}
		case err := <-w.watcher.Errors:
			w.logger.Error("Watcher got error", zap.Error(err))
		}
	}
}

func addCertsToWatch(watcher *fsnotify.Watcher, opts Options) error {
	if len(opts.CAPath) != 0 {
		err := watcher.Add(opts.CAPath)
		if err != nil {
			return err
		}
	}
	if len(opts.ClientCAPath) != 0 {
		err := watcher.Add(opts.ClientCAPath)
		if err != nil {
			return err
		}
	}
	if len(opts.CertPath) != 0 {
		err := watcher.Add(opts.CertPath)
		if err != nil {
			return err
		}
	}
	if len(opts.KeyPath) != 0 {
		err := watcher.Add(opts.KeyPath)
		if err != nil {
			return err
		}
	}
	return nil
}
