

package models

type TestpaperTests struct {
 TestpaperTestsID int  // 
 TestpaperID int  // 
 QuestionBankID int  // 
}

type testpaperTestsOp struct{}

var TestpaperTestsOp = &testpaperTestsOp{}
var DefaultTestpaperTests = &TestpaperTests{}

func (op *testpaperTestsOp) Insert(m *TestpaperTests) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *testpaperTestsOp) InsertTx(session *gocql.Session, m *TestpaperTests) (int64, error) {
	sql := "insert into testpaper_tests(testpaper_tests_id,testpaper_id,question_bank_id) values(?,?,?)"
	if err := session.Query(
		sql,
		 m.TestpaperTestsID ,
		 m.TestpaperID ,
		 m.QuestionBankID ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *testpaperTestsOp) QueryByMap(m map[string]interface{}, options []string) ([]*TestpaperTests, error) {
	result := []*TestpaperTests{}
	var params []interface{}

	sql := "select testpaper_tests_id,testpaper_id,question_bank_id from testpaper_tests"

	kNo := 0
	for k,v := range m{
		if (kNo==0){
			sql += " where "+ k +" = ?"
		}else{
			sql += " and "+ k +" = ?"
		}

		kNo += 1

		params = append(params, v)
	}

	if len(m) >0 {
		for _, option := range options{
			sql += " " + option
		}
	} 

	iter := mysql.Query(sql, params...).Iter()

	if nil == iter{
		return result, nil
	}

	data := &TestpaperTests{}
	for iter.Scan(
	 &data.TestpaperTestsID,
	 &data.TestpaperID,
	 &data.QuestionBankID,
	
	) {
		result = append(result, data)

		data = &TestpaperTests{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *testpaperTestsOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*TestpaperTests, error) {
	result := []*TestpaperTests{}
	var params []interface{}

	sql := "select testpaper_tests_id,testpaper_id,question_bank_id from testpaper_tests"

	kNo := 0
	for k,v := range m{
		if (kNo==0){
			sql += " where "+ k +" ?"
		}else{
			sql += " and "+ k +" ?"
		}

		kNo += 1

		params = append(params, v)
	}

	if len(m) >0 {
		for _, option := range options{
			sql += " " + option
		}
	} 

	iter := mysql.Query(sql, params...).Iter()

	if nil == iter{
		return result, nil
	}

	data := &TestpaperTests{}
	for iter.Scan(
	 &data.TestpaperTestsID,
	 &data.TestpaperID,
	 &data.QuestionBankID,
	
	) {
		result = append(result, data)

		data = &TestpaperTests{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


