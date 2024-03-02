//go:build !nowazero
// +build !nowazero

package wypes

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func (ms Modules) DefineWazero(runtime wazero.Runtime) error {
	for modName, funcs := range ms {
		err := funcs.DefineWazero(runtime, modName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Module) DefineWazero(runtime wazero.Runtime, modName string) error {
	var err error
	mb := runtime.NewHostModuleBuilder(modName)
	for funcName, funcDef := range m {
		fb := mb.NewFunctionBuilder()
		fb = fb.WithGoModuleFunction(
			wazeroAdaptHostFunc(funcDef),
			funcDef.ParamValueTypes(),
			funcDef.ResultValueTypes(),
		)
		mb = fb.Export(funcName)
	}
	_, err = mb.Instantiate(context.Background())
	return err
}

func wazeroAdaptHostFunc(hf HostFunc) api.GoModuleFunction {
	return api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
		adaptedStack := SliceStack(stack)
		store := Store{
			Memory:  mod.Memory(),
			Stack:   &adaptedStack,
			Refs:    map[uint32]any{},
			Context: ctx,
		}
		hf.Call(store)
	})
}
