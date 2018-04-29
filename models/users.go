

package models

type Users struct {
 UserID string  // 用户id
 UserName string  // 用户真实姓名
 UserPass string  // 登陆密码
 Permission int  // 用户权限
}

type usersOp struct{}

var UsersOp = &usersOp{}
var DefaultUsers = &Users{}

func (op *usersOp) Insert(m *Users) (int64, error) {
	return op.InsertTx(mysql, m)
}

func (op *usersOp) InsertTx(session *gocql.Session, m *Users) (int64, error) {
	sql := "insert into users(user_id,user_name,user_pass,permission) values(?,?,?,?)"
	if err := session.Query(
		sql,
		 m.UserID ,
		 m.UserName ,
		 m.UserPass ,
		 m.Permission ,
		
	).Exec(); err != nil {
		return -1, err

	}

	return 0, nil
}

func (op *usersOp) QueryByMap(m map[string]interface{}, options []string) ([]*Users, error) {
	result := []*Users{}
	var params []interface{}

	sql := "select user_id,user_name,user_pass,permission from users"

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

	data := &Users{}
	for iter.Scan(
	 &data.UserID,
	 &data.UserName,
	 &data.UserPass,
	 &data.Permission,
	
	) {
		result = append(result, data)

		data = &Users{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}

func (op *usersOp) QueryByMapComparison(m map[string]interface{}, options []string) ([]*Users, error) {
	result := []*Users{}
	var params []interface{}

	sql := "select user_id,user_name,user_pass,permission from users"

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

	data := &Users{}
	for iter.Scan(
	 &data.UserID,
	 &data.UserName,
	 &data.UserPass,
	 &data.Permission,
	
	) {
		result = append(result, data)

		data = &Users{}
	}

	if err := iter.Close(); err != nil {
		fmt.Println("err:", err)
	}

	return result, nil
}


