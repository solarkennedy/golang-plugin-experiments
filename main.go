package main

import (
	"github.com/solarkennedy/golang-plugin-experiments/types"
	"log"
	"plugin"
)

func LoadPlugins(plugins []string) {
	d := types.InData{V: 1}
	log.Printf("Invoking pipeline with data %#v\n", d)
	o := types.OutData{}
	for _, p := range plugins {
		p, err := plugin.Open(p)
		if err != nil {
			panic(err)
		}
		pName, _ := p.Lookup("Name")

		log.Printf("Invoking plugin %s\n", pName)

		input, _ := p.Lookup("Input")
		f, _ := p.Lookup("F")

		*input.(*types.InData) = d
		f.(func())()

		output, _ := p.Lookup("Output")

		d = types.InData{
			V: output.(*types.OutData).V,
		}
		*input.(*types.InData) = d

		o = *output.(*types.OutData)
	}

	log.Printf("Final results %+v\n", o)
}

func main() {
	plugins := []string{"plugin1/plugin1.so", "plugin2/plugin2.so"}
	LoadPlugins(plugins)
}
