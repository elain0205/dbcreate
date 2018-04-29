package modeltool

import (
	"html/template"
	"strings"
)

type ModelInfo struct {
	BDName       string
	DBConnection string
	TableName    string
	PackageName  string
	ModelName    string
	TableSchema  *[]TABLE_SCHEMA
}

type TABLE_SCHEMA struct {
	Column_name    string
	Data_type      string
	Column_key     string
	Column_comment string
}

func (m *ModelInfo) ColumnNames() []string {
	result := make([]string, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {

		result = append(result, t.Column_name)

	}
	return result
}

func (m *ModelInfo) ColumnCount() int {
	return len(*m.TableSchema)
}

func (m *ModelInfo) PkColumnsSchema() []TABLE_SCHEMA {
	result := make([]TABLE_SCHEMA, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		if t.Column_key == "PRI" {
			result = append(result, t)
		}
	}
	return result
}

func (m *ModelInfo) HavePk() bool {
	return len(m.PkColumnsSchema()) > 0
}

func (m *ModelInfo) NoPkColumnsSchema() []TABLE_SCHEMA {
	result := make([]TABLE_SCHEMA, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		if t.Column_key != "PRI" {
			result = append(result, t)
		}
	}
	return result
}

func (m *ModelInfo) NoPkColumns() []string {
	noPkColumnsSchema := m.NoPkColumnsSchema()
	result := make([]string, 0, len(noPkColumnsSchema))
	for _, t := range noPkColumnsSchema {
		result = append(result, t.Column_name)
	}
	return result
}

func (m *ModelInfo) PkColumns() []string {
	pkColumnsSchema := m.PkColumnsSchema()
	result := make([]string, 0, len(pkColumnsSchema))
	for _, t := range pkColumnsSchema {
		result = append(result, t.Column_name)
	}
	return result
}

func IsUUID(str string) bool {
	return "uuid" == str
}

func FirstCharLower(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func Tags(columnName string) template.HTML {

	return template.HTML("`db:" + `"` + columnName + `"` +
		" json:" + `"` + columnName + "\"`")
}

func ExportColumn(columnName string) string {
	columnItems := strings.Split(columnName, "_")
	columnItems[0] = FirstCharUpper(columnItems[0])
	for i := 0; i < len(columnItems); i++ {
		item := strings.Title(columnItems[i])

		if strings.ToUpper(item) == "ID" {
			item = "ID"
		}

		columnItems[i] = item
	}

	return strings.Join(columnItems, "")

}

func TypeConvert(str string) string {

	switch str {
	case "smallint", "tinyint":
		return "int8"

	case "varchar", "text", "longtext", "char":
		return "string"

	case "date":
		return "string"

	case "int":
		return "int"

	case "timestamp", "datetime":
		return "time.Time"

	case "bigint":
		return "int64"

	case "float", "double", "decimal":
		return "float64"
	case "uuid":
		return "gocql.UUID"

	default:
		return str
	}
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func ColumnAndType(table_schema []TABLE_SCHEMA) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.Column_name+" "+TypeConvert(t.Data_type))
	}
	return strings.Join(result, ",")
}

func ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+Postfix)
	}
	return strings.Join(result, sep)
}

func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}
