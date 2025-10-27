package opts3

import "reflect"

type Opts interface {
	Get(tag any, typ reflect.Type) (any, bool)
}

type optsImpl struct {
	parent Opts
	tag    any
	typ    reflect.Type
	value  any
}

func (o *optsImpl) Get(tag any, typ reflect.Type) (any, bool) {
	if o == nil {
		return nil, false
	}
	typeMatched := o.typ == typ || o.typ == nil
	if o.tag == tag && typeMatched {
		return o.value, true
	}
	return o.parent.Get(tag, typ)
}

func NewOpts() Opts {
	var o optsImpl
	return &o
}

func WrapOpts(parent Opts, tag any, typ reflect.Type, value any) Opts {
	return &optsImpl{
		parent: parent,
		tag:    tag,
		typ:    typ,
		value:  value,
	}
}
