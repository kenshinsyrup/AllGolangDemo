package main

import (
	"allgolangdemo/traversestruct/genfile"
	"fmt"
	"reflect"
)

type SB struct {
	sb string
}

type SU struct {
	Path  *string
	Title *string
}
type SD struct {
	SBField  SB
	Upload   *SU
	DataName string
}
type SA struct {
	Data []SD
}

func main() {
	fmt.Println("Hello, playground")

	// sb1 := SB{}
	// p1 := "sp1 path"
	// t1 := "sp1 title"
	// sp1 := SU{
	// 	Path:  &p1,
	// 	Title: &t1,
	// }
	// sd1 := SD{
	// 	SBField: sb1,
	// 	Upload:  &sp1,
	// }
	// sa1 := SA{
	// 	Data: []SD{sd1},
	// }
	// fmt.Println(sa1)

	// var sa2 SA
	// fmt.Println(sa2)

	// sa3 := SA{
	// 	Data: []SD{{DataName: "sa3 data"}},
	// }
	// fmt.Println(sa3)

	// v := reflect.ValueOf(sa1)
	// if v.Kind() == reflect.Struct {
	// 	parent := v.Type().Name()
	// 	fieldName := v.Type().Field(0).Name
	// 	extract(v, parent, fieldName)
	// }
	// if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
	// 	parent := v.Elem().Type().Name()
	// 	fieldName := v.Elem().Type().Field(0).Name
	// 	extract(v.Elem(), parent, fieldName)
	// }

	// fmt.Println("&&&&&&&&&&&&")
	// ss := []SA{sa1}
	// v = reflect.ValueOf(ss)
	// fmt.Println(v.Kind(), reflect.TypeOf(v.Interface()))

	genfile.Genfile()
}

func extract(v reflect.Value, parent, fieldName string) {
	if v.Kind() == reflect.Struct {
		fmt.Println("******")
		fmt.Println(reflect.TypeOf(v.Interface()))
		if _, ok := v.Interface().(SB); ok {
			return
		}
		if su, ok := v.Interface().(SU); ok {
			if su.Path != nil {
				fmt.Println(parent + fieldName + *su.Path)
			}
			if su.Title != nil {
				fmt.Println(parent + fieldName + *su.Title)
			}
		}

		for i := 0; i < v.NumField(); i++ {
			fmt.Println("i: ", i)
			fmt.Println("type: ", v.Type().Field(i))
			vf := v.Field(i)
			if vf.Kind() == reflect.Ptr {
				vf = vf.Elem()
			}
			if vf.Kind() == reflect.Struct {
				parent := v.Type().Name()
				fieldName := v.Type().Field(i).Name
				extract(vf, parent, fieldName)
			}
			fmt.Println("ready to slice")
			fmt.Println("kind: ", vf.Kind())
			if vf.Kind() == reflect.Slice {
				fmt.Println("(&&&&")
				for i := 0; i < vf.Len(); i++ {
					fmt.Println("index: ", i)
					if vf.Index(i).Kind() == reflect.Struct {
						parent := vf.Index(i).Type().Name()
						fieldName := vf.Index(i).Type().Field(i).Name
						extract(vf.Index(i), parent, fieldName)
					}
					if vf.Index(i).Kind() == reflect.Ptr && vf.Index(i).Elem().Kind() == reflect.Struct {
						parent := vf.Index(i).Elem().Type().Name()
						fieldName := vf.Index(i).Elem().Type().Field(i).Name
						extract(vf.Index(i).Elem(), parent, fieldName)
					}
				}
			}
		}
	}
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Kind() == reflect.Struct {
				fmt.Println("parent struct: ", v.Index(i).Type().Name())
				parent := v.Index(i).Type().Name()
				fieldName := v.Index(i).Type().Field(i).Name
				extract(v.Index(i), parent, fieldName)
			}
			if v.Index(i).Kind() == reflect.Ptr && v.Index(i).Elem().Kind() == reflect.Struct {
				fmt.Println("parent struct: ", v.Index(i).Elem().Type().Name())
				parent := v.Index(i).Elem().Type().Name()
				fieldName := v.Index(i).Elem().Type().Field(i).Name
				extract(v.Index(i).Elem(), parent, fieldName)
			}
		}
	}
}
