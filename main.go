package main

import (
	"dbcreate/dbhelper"
	"dbcreate/modeltool"
	_ "dbcreate/routers"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	//dbcreate.exe "mysql" "" "3306" "exam" "root" "123456" "./modeltool/model.tpl"
	dbtype := os.Args[1]
	ip := os.Args[2]
	port := os.Args[3]
	dbName := os.Args[4]
	user := os.Args[5]
	pass := os.Args[6]
	tplFile := os.Args[7]

	dbhelper.Initdb(dbtype, ip, port, dbName, user, pass)

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	//获取表名
	var list orm.ParamsList
	num, err := o.Raw("select table_name from information_schema.TABLES where TABLE_SCHEMA=?", dbName).ValuesFlat(&list)
	if err != nil {
		fmt.Println(err)
	}

	if err == nil && num > 0 {
		for i := 0; i < len(list); i++ {
			tableName := list[i]

			fmt.Println(tableName)

			res := make([]modeltool.TABLE_SCHEMA, 0)
			nums, err := o.Raw("SELECT column_name,data_type,column_key,column_comment from information_schema.COLUMNS where TABLE_NAME=? and table_schema = ?", tableName, dbName).QueryRows(&res)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(nums)
			fmt.Println(res)
			value, ok := tableName.(string)
			if !ok {
				fmt.Println("It's not ok for type string")
				return
			}
			genModelFile(res, value, tplFile, dbtype)
		}
	}

}

func genModelFile(tableSchema []modeltool.TABLE_SCHEMA, tableName, tplFile, dbConnection string) {

	modelFolder := "./models/"
	packageName := "models"
	fileTail := ""

	logDir, _ := filepath.Abs(modelFolder)
	if _, err := os.Stat(logDir); err != nil {
		os.Mkdir(logDir, os.ModePerm)
	}

	data, err := ioutil.ReadFile(tplFile)
	if nil != err {
		fmt.Println("read tplFile err:", err)
		return
	}

	render := template.Must(template.New("statscassandra").
		Funcs(template.FuncMap{
			"FirstCharUpper":       modeltool.FirstCharUpper,
			"FirstCharLower":       modeltool.FirstCharLower,
			"TypeConvert":          modeltool.TypeConvert,
			"Tags":                 modeltool.Tags,
			"ExportColumn":         modeltool.ExportColumn,
			"Join":                 modeltool.Join,
			"MakeQuestionMarkList": modeltool.MakeQuestionMarkList,
			"ColumnAndType":        modeltool.ColumnAndType,
			"ColumnWithPostfix":    modeltool.ColumnWithPostfix,
			"IsUUID":               modeltool.IsUUID,
		}).
		Parse(string(data)))

	model := &modeltool.ModelInfo{
		PackageName:  packageName,
		DBConnection: dbConnection,
		TableName:    tableName,
		ModelName:    tableName,
		TableSchema:  &tableSchema}

	fileName := modelFolder + strings.ToLower(tableName) + fileTail + ".go"

	os.Remove(fileName)
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if err := render.Execute(f, model); err != nil {

		log.Fatal("", err)
	}
	fmt.Println("fileName", fileName)
	cmd := exec.Command("goimports", "-w", fileName)
	cmd.Run()

}
