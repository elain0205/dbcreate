

package models

type Options struct {
 OptionsID int  // ID
 Op string  // 选项
 Content string  // 选项内容
 QuestionBankID int  // 
}

type optionsOp struct{}

var OptionsOp = &optionsOp{}
var DefaultOptions = &Options{}

func (op *optionsOp) Insert(m *Options) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *optionsOp) InsertTx(session *gocql.Session, m *Options) (int64, error) {
	sql := "insert into options(options_id,op,content,question_bank_id) values(?,?,?,?)"
	if err := session.Query(
		sql,
		 m.OptionsID ,
		 m.Op ,
		 m.Content ,
		 m.QuestionBankID ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *optionsOp) QueryByMap(m map[string]interface{}, options []string) ([]*Options, error) {
	result := []*Options{}
	var params []interface{}

	sql := "select options_id,op,content,question_bank_id from options"

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

	data := &Options{}
	for iter.Scan(
	 &data.OptionsID,
	 &data.Op,
	 &data.Content,
	 &data.QuestionBankID,
	
	) {
		result = append(result, data)

		data = &Options{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *optionsOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*Options, error) {
	result := []*Options{}
	var params []interface{}

	sql := "select options_id,op,content,question_bank_id from options"

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

	data := &Options{}
	for iter.Scan(
	 &data.OptionsID,
	 &data.Op,
	 &data.Content,
	 &data.QuestionBankID,
	
	) {
		result = append(result, data)

		data = &Options{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


