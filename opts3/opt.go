package opts3

type Opt interface {
	Apply(Env) Env
}

type OptFunc func(Env) Env

func (f OptFunc) Apply(env Env) Env {
	return f(env)
}

type Opts []Opt

func (o Opts) Apply(env Env) Env {
	for _, opt := range o {
		env = opt.Apply(env)
	}
	return env
}