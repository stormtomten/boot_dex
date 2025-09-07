package main

import (
	"fmt"
	"reflect"
	"strings"
)

func commandInspect(args []string, cfg *config) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <name>")
	}
	name := args[0]
	p, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("Pokemon not owned")
	}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		// s.Stat.Name like "hp", s.BaseStat like 40
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		// t.Type.Name like "normal"
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func printFields(val interface{}, prefix string, depth int) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	indent := strings.Repeat("  ", depth) // Indentation based on depth
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldName := prefix + field.Name
		if value.Kind() == reflect.Struct && field.Type.PkgPath() == "" {
			// Print struct name as header, then recurse
			fmt.Printf("%s%s (struct):\n", indent, fieldName)
			printFields(value.Interface(), fieldName+".", depth+1)
		} else {
			// Print simple field
			fmt.Printf("%s%s: %v\n", indent, fieldName, value.Interface())
		}
	}
}
