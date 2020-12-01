# 作业

我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

答：不应该直接Wrap sql.ErrNoRows，应该自定义一个error，将这个error Wrap后返回。 可以屏蔽底层使用的技术细节， Wrap自定义的error可以增加堆栈信息，便于调试

# 代码

~~~go
// DAO 层

type Dao struct {
}
var ErrRecordNotFound = errors.New("record not found")

func (d *Dao) FindData(key int) (data string, err error) {
    err = sqlxxx()
	if errors.Is(err, sql.ErrNoRows) {
        err = ErrRecordNotFound
	}
    
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("FindData error key: %v error", key))
	}
	return
}


~~~

~~~go
// Service 层
type Service struct {
}

func (s *Service) FindData(key int) (data string, err error) {
    return dao.FindData(key)
}

~~~

~~~
func main() {
	s := Service{}
	
	data, err := s.FindData(1)
	if erros.Is(err, ErrRecordNotFond) {
		fmt.Printf("%+v\n", err)
	} 
}
~~~

