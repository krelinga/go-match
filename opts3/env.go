package opts3

import "reflect"

type Env interface {
	Get(tag any, typ reflect.Type) (any, bool)
}

type envImpl struct {
	parent Env
	tag    any
	typ    reflect.Type
	value  any
}

func (o *envImpl) Get(tag any, typ reflect.Type) (any, bool) {
	if o == nil {
		return nil, false
	}
	typeMatched := o.typ == typ || o.typ == nil
	if o.tag == tag && typeMatched {
		return o.value, true
	}
	return o.parent.Get(tag, typ)
}

func NewEnv() Env {
	var o envImpl
	return &o
}

func WrapEnv(parent Env, tag any, typ reflect.Type, value any) Env {
	return &envImpl{
		parent: parent,
		tag:    tag,
		typ:    typ,
		value:  value,
	}
}
