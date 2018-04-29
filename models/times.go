

package models

type Times struct {
 TimesID int  // 唯一标识
 TestpaperID int  // 试卷_id
 UserID string  // 用户ID
 DataMin float64  // 
 TimesState int  // 数据状态
}

type timesOp struct{}

var TimesOp = &timesOp{}
var DefaultTimes = &Times{}

func (op *timesOp) Insert(m *Times) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *timesOp) InsertTx(session *gocql.Session, m *Times) (int64, error) {
	sql := "insert into times(times_id,testpaper_id,user_id,data_min,times_state) values(?,?,?,?,?)"
	if err := session.Query(
		sql,
		 m.TimesID ,
		 m.TestpaperID ,
		 m.UserID ,
		 m.DataMin ,
		 m.TimesState ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *timesOp) QueryByMap(m map[string]interface{}, options []string) ([]*Times, error) {
	result := []*Times{}
	var params []interface{}

	sql := "select times_id,testpaper_id,user_id,data_min,times_state from times"

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

	data := &Times{}
	for iter.Scan(
	 &data.TimesID,
	 &data.TestpaperID,
	 &data.UserID,
	 &data.DataMin,
	 &data.TimesState,
	
	) {
		result = append(result, data)

		data = &Times{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *timesOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*Times, error) {
	result := []*Times{}
	var params []interface{}

	sql := "select times_id,testpaper_id,user_id,data_min,times_state from times"

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

	data := &Times{}
	for iter.Scan(
	 &data.TimesID,
	 &data.TestpaperID,
	 &data.UserID,
	 &data.DataMin,
	 &data.TimesState,
	
	) {
		result = append(result, data)

		data = &Times{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


