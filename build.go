package sorm

import (
	"fmt"
	"os"
	"reflect"
)

type builder struct {
}

func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) Build(ob ...interface{}) {

	if _, err := os.Stat("./model"); os.IsNotExist(err) {
		fmt.Println("mkdir model")
		os.Mkdir("./model", 0777)
	}

	for _, v := range ob {
		typeOf := reflect.TypeOf(v)
		if typeOf.Kind() != reflect.Struct {
			continue
		}
		fmt.Println("receive : ", typeOf.Name())
		file, err := os.Create("./model/" + typeOf.Name() + ".go")
		if err != nil {
			fmt.Println(err)
			file.Close()
			continue
		}
		_, _ = file.WriteString("package model\n\n")
		file.WriteString("type " + typeOf.Name() + " struct{\n")
		//处理成员
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			fmt.Println(field.Name, field.Type)
			file.WriteString("	" + field.Name + " " + field.Type.String() + "\n")
		}
		file.WriteString("}\n")
		fmt.Println()
		file.Close()
	}
}
