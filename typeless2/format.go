package typeless2

import "fmt"

type Fmt func(any) (string, error)

func FmtDef(val any) (string, error) {
	return fmt.Sprintf("%#v", val), nil
}