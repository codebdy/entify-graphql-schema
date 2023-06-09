package script

import (
	"github.com/dop251/goja"
)

func Enable(vm *goja.Runtime) {
	vm.Set("log", Log)
	vm.Set("iFetch", FetchFn)
	vm.Set("writeToCache", WriteToCache)
	vm.Set("readFromCache", ReadFromCache)
}

func GetCommonCodes() string {
	return `
	const debug = {}
	debug.log = log
	`
}
