

package models

type QuestionBank struct {
 QuestionBankID int  // 
 Stem string  // 题目
 Answer string  // 正确答案
 TestsType int  // 题目类型（判断 0或者选择 1）
 State int  // 状态 0 删除 1有效
}

type questionBankOp struct{}

var QuestionBankOp = &questionBankOp{}
var DefaultQuestionBank = &QuestionBank{}

func (op *questionBankOp) Insert(m *QuestionBank) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *questionBankOp) InsertTx(session *gocql.Session, m *QuestionBank) (int64, error) {
	sql := "insert into question_bank(question_bank_id,stem,answer,tests_type,state) values(?,?,?,?,?)"
	if err := session.Query(
		sql,
		 m.QuestionBankID ,
		 m.Stem ,
		 m.Answer ,
		 m.TestsType ,
		 m.State ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *questionBankOp) QueryByMap(m map[string]interface{}, options []string) ([]*QuestionBank, error) {
	result := []*QuestionBank{}
	var params []interface{}

	sql := "select question_bank_id,stem,answer,tests_type,state from question_bank"

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

	data := &QuestionBank{}
	for iter.Scan(
	 &data.QuestionBankID,
	 &data.Stem,
	 &data.Answer,
	 &data.TestsType,
	 &data.State,
	
	) {
		result = append(result, data)

		data = &QuestionBank{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *questionBankOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*QuestionBank, error) {
	result := []*QuestionBank{}
	var params []interface{}

	sql := "select question_bank_id,stem,answer,tests_type,state from question_bank"

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

	data := &QuestionBank{}
	for iter.Scan(
	 &data.QuestionBankID,
	 &data.Stem,
	 &data.Answer,
	 &data.TestsType,
	 &data.State,
	
	) {
		result = append(result, data)

		data = &QuestionBank{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


