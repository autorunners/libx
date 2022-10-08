package logx

func Debug(v ...interface{}) {
	defaultLogx.Debug(v...)
	return
}

func Info(v ...interface{}) {
	defaultLogx.Info(v...)
	return
}
func Infof(format string, v ...interface{}) {
	defaultLogx.Infof(format, v...)
	return
}

func Warn(v ...interface{}) {
	defaultLogx.Warn(v...)
	return
}

func Error(v ...interface{}) {
	defaultLogx.Error(v...)
	return
}
