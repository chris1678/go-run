package logger

type Option func(*options)

type options struct {
	path   string
	level  string
	stdout string
	cap    uint
}

func setDefault() options {
	return options{
		path:   "temp/logs",
		level:  "warn",
		stdout: "default",
	}
}

func WithPath(s string) Option {
	return func(o *options) {
		o.path = s
	}
}

func WithLevel(s string) Option {
	return func(o *options) {
		o.level = s
	}
}

func WithStdout(s string) Option {
	return func(o *options) {
		o.stdout = s
	}
}

func WithCap(n uint) Option {
	return func(o *options) {
		o.cap = n
	}
}
