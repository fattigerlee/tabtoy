package gen

import "github.com/fattigerlee/tabtoy/v3/model"

type GenFunc func(globals *model.Globals) (data []byte, err error)
