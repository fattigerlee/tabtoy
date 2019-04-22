package gosrc

import (
	"fmt"
	"github.com/fattigerlee/tabtoy/util"
	"github.com/fattigerlee/tabtoy/v3/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

// 将定义用的类型，转换为不同语言对应的复合类型

func init() {
	UsefulFunc["GoType"] = func(tf *model.TypeDefine) string {

		convertedType := model.LanguagePrimitive(tf.FieldType, "go")

		if tf.IsArray() {
			// 原始类型
			if model.PrimitiveExists(convertedType) {
				return "[]" + convertedType
			}

			// 自建类型
			return "[]*" + convertedType
		}

		return convertedType
	}

	UsefulFunc["GoTabTag"] = func(fieldType *model.TypeDefine) string {

		var sb strings.Builder

		var kv []string

		if fieldType.Name != "" {
			kv = append(kv, fmt.Sprintf("json:\"%s\",tb_name:\"%s\"", util.CamelToUnderline(fieldType.FieldName), fieldType.Name))
		}

		if len(kv) > 0 {
			sb.WriteString("`")

			for _, s := range kv {
				sb.WriteString(s)
			}

			sb.WriteString("`")
		}

		return sb.String()
	}

	UsefulFunc["JsonTab"] = func(tabType *model.DataTable) string {
		return fmt.Sprintf("`json:\"%s\"`", util.CamelToUnderline(tabType.HeaderType))
	}

	UsefulFunc["JsonTabOmit"] = func() string {
		return "`json:\"-\"`"
	}

}
