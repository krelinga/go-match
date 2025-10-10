package opts

import "reflect"

type Option func(Options) Options

type Options interface {
	With(typeKey reflect.Type, opKey any, val any) Options
	Get(typeKey reflect.Type, opKey any) any
}

func NewOptions() Options {
	var o *options
	return o
}

type options struct {
	parent  Options
	typeKey reflect.Type
	opKey   any
	val     any
}

func (o *options) With(typeKey reflect.Type, opKey any, val any) Options {
	return &options{
		parent:  o,
		typeKey: typeKey,
		opKey:   opKey,
		val:     val,
	}
}

func (o *options) Get(typeKey reflect.Type, opKey any) any {
	if o == nil {
		return nil
	}
	typeMatch := o.typeKey == nil || o.typeKey == typeKey
	if typeMatch && o.opKey == opKey {
		return o.val
	}
	return o.parent.Get(typeKey, opKey)
}
