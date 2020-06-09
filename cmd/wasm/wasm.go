package main

import (
	"errors"
	"github.com/yanishoss/tml"
	"syscall/js"
)

func main() {
	defaultConfigFn := js.FuncOf(defaultConfig)
	parseFn := js.FuncOf(parseTML)

	js.Global().Set("TML_parse", parseFn)
	js.Global().Set("TML_withDefaultConfig", defaultConfigFn)
}

func parseTML(_ js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return errors.New("an input string must be provided")
	}

	input := args[0].String()
	conf := tml.WithDefaultConfig()

	if len(args) >= 2 {
		defaultUnit := args[1].Get("defaultUnit")
		rpeRange := args[1].Get("rpeRange")
		validUnits := args[1].Get("validUnits")


		if !defaultUnit.IsUndefined() && !defaultUnit.IsNull() {
			conf.DefaultUnit = defaultUnit.String()
		}

		if !rpeRange.IsUndefined() && !rpeRange.IsNull() && rpeRange.Get("length").Int() == 2 {
			conf.RPERange[0] = rpeRange.Index(0).Float()
			conf.RPERange[1] = rpeRange.Index(1).Float()
		}

		if !validUnits.IsUndefined() && !validUnits.IsNull() {
			length := validUnits.Get("length").Int()
			units := make([]string, length)

			for i := 0; i < length; i++ {
				units[i] = validUnits.Index(i).String()
			}

			conf.ValidUnits = units
		}
	}

	w, err := tml.Parse(input, conf)

	if err != nil {
		return err
	}

	return w
}

func defaultConfig(_ js.Value, _ []js.Value) interface{} {
	return tml.WithDefaultConfig()
}