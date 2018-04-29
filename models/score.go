

package models

type Score struct {
 ScoreID int  // 分数ID
 UsersID string  // 用户ID
 TestpaperID int  // 
 Fraction float64  // 分数
 Img string  // 
}

type scoreOp struct{}

var ScoreOp = &scoreOp{}
var DefaultScore = &Score{}

func (op *scoreOp) Insert(m *Score) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *scoreOp) InsertTx(session *gocql.Session, m *Score) (int64, error) {
	sql := "insert into score(score_id,users_id,testpaper_id,fraction,img) values(?,?,?,?,?)"
	if err := session.Query(
		sql,
		 m.ScoreID ,
		 m.UsersID ,
		 m.TestpaperID ,
		 m.Fraction ,
		 m.Img ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *scoreOp) QueryByMap(m map[string]interface{}, options []string) ([]*Score, error) {
	result := []*Score{}
	var params []interface{}

	sql := "select score_id,users_id,testpaper_id,fraction,img from score"

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

	data := &Score{}
	for iter.Scan(
	 &data.ScoreID,
	 &data.UsersID,
	 &data.TestpaperID,
	 &data.Fraction,
	 &data.Img,
	
	) {
		result = append(result, data)

		data = &Score{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *scoreOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*Score, error) {
	result := []*Score{}
	var params []interface{}

	sql := "select score_id,users_id,testpaper_id,fraction,img from score"

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

	data := &Score{}
	for iter.Scan(
	 &data.ScoreID,
	 &data.UsersID,
	 &data.TestpaperID,
	 &data.Fraction,
	 &data.Img,
	
	) {
		result = append(result, data)

		data = &Score{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


