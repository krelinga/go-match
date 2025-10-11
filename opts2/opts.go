package opts2

import "reflect"

type Opts interface {
	With(tag any, typ reflect.Type, val any) Opts
	Get(tag any, typ reflect.Type) any
}

func NewOpts() Opts {
	var opts *opts
	return opts
}

type Opt func(Opts) Opts

type opts struct {
	parent Opts
	tag any
	typ reflect.Type
	val any
}

func (o *opts) With(tag any, typ reflect.Type, val any) Opts {
	return &opts{parent: o, tag: tag, typ: typ, val: val}
}

func (o *opts) Get(tag any, typ reflect.Type) any {
	if o == nil {
		return nil
	}
	typeMatch := o.typ == nil || typ == o.typ
	if o.tag == tag && typeMatch {
		return o.val
	}
	return o.parent.Get(tag, typ)
}
