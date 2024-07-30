package logger

import "log/slog"

func LogInternalErr(location string, err error) {
	slog.Error("`" + location + "`" + " internal error: " + err.Error())
}
