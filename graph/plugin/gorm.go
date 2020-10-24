package plugin

import (
	"fmt"
	"go/types"
	"strings"
	"unicode"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func gorm(b *modelgen.ModelBuild) {

	var cfg config.Config
	typ := types.NewNamed(
		types.NewTypeName(0, cfg.Model.Pkg(), "time.Time", nil),
		nil,
		nil,
	)
	typP := types.NewNamed(
		types.NewTypeName(0, cfg.Model.Pkg(), "*time.Time", nil),
		nil,
		nil,
	)

	for _, model := range b.Models {
		model.Fields = append(model.Fields,
			&modelgen.Field{
				Name:        "CreatedAt",
				Type:        typ,
				Description: "gorm.Model",
			},
			&modelgen.Field{
				Name: "UpdatedAt",
				Type: typ,
			},
			&modelgen.Field{
				Name: "DeletedAt",
				Type: typP,
				Tag:  `sql:"index"`,
			},
		)

		for _, field := range model.Fields {

			fieldType := field.Type.String()[strings.LastIndex(field.Type.String(), ".")+1:]

			// many2many tag
			if strings.HasPrefix(fieldType, "[]") {
				name := m2mName(model.Name, fieldType)
				field.Tag += fmt.Sprintf(` gorm:"many2many:%ss`, name)
			}

			// ID tag
			if field.Name == "id" {
				field.Tag += ` gorm:"primary_key"`
				// Why?
				// ` gorm:"primary_key;type:uuid;default:uuid_generate_v4()`
			}
		}

	}
}

func m2mName(str1, str2 string) string {
	var m2mName string
	if str1 > str2 {
		m2mName = str1 + str2
	} else {
		m2mName = str2 + str1
	}
	return snakeCase(m2mName)
}

func snakeCase(str string) string {
	var newStr string
	for i, c := range str {
		if unicode.IsUpper(c) && i != 0 {
			newStr += "_"
		}
		newStr += string(unicode.ToLower(c))
	}
	return newStr
}
