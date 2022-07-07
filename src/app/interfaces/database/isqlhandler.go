// database層から実際に呼び出すのはこれ
package database

// Execute, Queryの戻り値をResultとRowのインターフェースにしていて、そうすることで依存ルールを守っている。
type ISqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row,error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
