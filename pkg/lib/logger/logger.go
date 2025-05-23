package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	stdLog "log"
	"log/slog"
	"path/filepath"
	"runtime"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = "\033[35m" + level + "\033[0m" // Magenta
	case slog.LevelInfo:
		level = "\033[34m" + level + "\033[0m" // Blue
	case slog.LevelWarn:
		level = "\033[33m" + level + "\033[0m" // Yellow
	case slog.LevelError:
		level = "\033[31m" + level + "\033[0m" // Red
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		// Special handling for error values
		if a.Key == "err" || a.Key == "error" {
			if err, ok := a.Value.Any().(error); ok {
				fields[a.Key] = err.Error()
				return true
			}
		}
		fields[a.Key] = a.Value.Any()
		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	if r.Level != slog.LevelInfo {
		pc, file, line, ok := runtime.Caller(4)
		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			funcName = filepath.Base(funcName)
			fileName := filepath.Base(file)
			fields["caller"] = fileName + ":" + funcName + ":" + fmt.Sprintf("%d", line)
		}
	}

	var b []byte
	var err error
	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := "\033[36m" + r.Message + "\033[0m" // Cyan for msg
	h.l.Println(
		timeStr,
		level,
		msg,
		"\033[37m"+string(b)+"\033[0m", // White for fields
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

func SetupPrettySlog(writer io.Writer) *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(writer)

	return slog.New(handler)
}
