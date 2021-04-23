package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

type SentryLogger struct {
	*AbstractLogger
	hub   *sentry.Hub
	level Level
	err   error
}

func NewSentryLogger(hub *sentry.Hub, level Level) *SentryLogger {
	a := &AbstractLogger{}
	s := &SentryLogger{
		a,
		hub,
		level,
		nil,
	}
	a.Logger = s
	return s
}

func (s *SentryLogger) Clone() *SentryLogger {
	h := s.hub.Clone()
	clone := NewSentryLogger(h, s.level)
	clone.err = s.err
	return clone
}

//func (s *SentryLogger) WithField(key string, value interface{}) Logger {
//	return s.WithFields(Fields{key: value})
//}

func (s *SentryLogger) WithFields(fields Fields) Logger {
	clone := s.Clone()
	clone.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtras(fields)
	})
	return clone
}

func (s *SentryLogger) WithError(err error) Logger {
	clone := s.Clone()
	clone.err = err
	return clone
}

func (s *SentryLogger) WithRequest(req *http.Request) Logger {
	clone := s.Clone()
	clone.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetRequest(req)
	})
	return clone
}

func (s *SentryLogger) Log(level Level, args ...interface{}) {
	if level.isLower(s.level) {
		return
	}
	message := fmt.Sprint(args...)
	clone := s.Clone()
	clone.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.Level(level))
		if clone.err != nil {
			scope.AddEventProcessor(func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
				err := clone.err
				err = errors.Wrap(err, message)
				event.Message = err.Error()
				event.Level = sentry.Level(level)
				//event.Fingerprint = append(event.Fingerprint, "{{ default }}", "{{ message }}", "{{ error.type }}", "{{ error.value }}")
				//event.Fingerprint = []string{"{{ default }}", "{{ message }}", "{{ error.type }}", "{{ error.value }}"}
				return event
			})
		}
	})
	if clone.err != nil {
		clone.hub.CaptureException(clone.err)
	} else {
		clone.hub.CaptureMessage(message)
	}
}
