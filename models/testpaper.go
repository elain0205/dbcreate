

package models

type Testpaper struct {
 TestpaperID int  // 
 TestpaperName string  // 试卷名
 TestpaperState int  // 试卷状态 0为无效  1为有效
 StartDate time.Time  // 开始时间
 EndDate time.Time  // 结束时间
 IsStart int  // 是否开始考试 0 关闭 1开始
}

type testpaperOp struct{}

var TestpaperOp = &testpaperOp{}
var DefaultTestpaper = &Testpaper{}

func (op *testpaperOp) Insert(m *Testpaper) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *testpaperOp) InsertTx(session *gocql.Session, m *Testpaper) (int64, error) {
	sql := "insert into testpaper(testpaper_id,testpaper_name,testpaper_state,start_date,end_date,is_start) values(?,?,?,?,?,?)"
	if err := session.Query(
		sql,
		 m.TestpaperID ,
		 m.TestpaperName ,
		 m.TestpaperState ,
		 m.StartDate ,
		 m.EndDate ,
		 m.IsStart ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *testpaperOp) QueryByMap(m map[string]interface{}, options []string) ([]*Testpaper, error) {
	result := []*Testpaper{}
	var params []interface{}

	sql := "select testpaper_id,testpaper_name,testpaper_state,start_date,end_date,is_start from testpaper"

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

	data := &Testpaper{}
	for iter.Scan(
	 &data.TestpaperID,
	 &data.TestpaperName,
	 &data.TestpaperState,
	 &data.StartDate,
	 &data.EndDate,
	 &data.IsStart,
	
	) {
		result = append(result, data)

		data = &Testpaper{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *testpaperOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*Testpaper, error) {
	result := []*Testpaper{}
	var params []interface{}

	sql := "select testpaper_id,testpaper_name,testpaper_state,start_date,end_date,is_start from testpaper"

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

	data := &Testpaper{}
	for iter.Scan(
	 &data.TestpaperID,
	 &data.TestpaperName,
	 &data.TestpaperState,
	 &data.StartDate,
	 &data.EndDate,
	 &data.IsStart,
	
	) {
		result = append(result, data)

		data = &Testpaper{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


