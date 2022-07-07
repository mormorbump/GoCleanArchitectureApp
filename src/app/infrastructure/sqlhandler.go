// DB接続には外部パッケージを使用しているのでinfrastructure層に定義し、外側のルールを内側に持ち込まないようにする。
package infrastructure

import (
	// 自分より内側のinterfaces/databaseに依存できる。そこでISqlHandlerインターフェースに依存する。
	"app/interfaces/database"
	// こいつはdatabase driver(下のやつとか)と一緒に使う必要がある。 https://pkg.go.dev/database/sql
	"database/sql"
	// driverをインストールするだけでdatabase/sqlの完全なAPIを使用できる https://github.com/go-sql-driver/mysql#usage
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB // sql.DB型のポインタ
}

// SqlHandlerのstruct型
func NewSqlHandler() *SqlHandler {
	// userName@tcp(コンテナ名:3306)/databaseNameって感じ
	conn, err := sql.Open("mysql", "root@tcp(db:3306)/cleanachitecturedb")
	if err != nil {
		panic(err.Error)
	}
	// new(Type)は、代入した型のポインタを消す変数。うけとった型をゼロ値で初期化すると、ポインタを返す性質がある。
	// ゼロ値が入った状態で初期化されるので、その後の取り回しがちょい早い。
	// https://code-graffiti.com/new-function-and-pointers-in-golang/#:~:text=%E6%9C%80%E5%BE%8C%E3%81%AB-,new(),%E8%BF%94%E3%81%99%E3%81%A8%E3%81%84%E3%81%86%E6%80%A7%E8%B3%AA%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler // この返り値をインターフェースにすることによってインターフェースを見たしている(何がきても大丈夫)
}

// interfaces.ISqlHandlerで実装した通りに引数と返り値を設定。
// 保存処理を行うとき(返り値が必要ない)の関数のラッパー
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	// Exec は、行を返すことなくクエリを実行する。
	// args はクエリ内の任意のプレースホルダパラメータ
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil // この戻り値もdatabase.Result、つまりインターフェース(何がきても大丈夫)
}

// 検索して行を返すときの関数
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	row := new(SqlRow)
	if err != nil {
		return row, err // エラーが出た時は検索結果なし、の意味も込めてSqlRowのゼロ値が入ったポインタを返す
	}
	row.Rows = rows
	return row, nil
}

// SQLから返ってきた内容を保管し、メソッドによって整形して返すためのstruct
// database.sql_handlerよりLastInsertIdとRowsAffectedを実装することでResultインターフェースを実装しているが、
//同時に sql.ResultインターフェースもLastInsertIdと RowsInsertIdが定義されているので、sql.Resultインターフェースを実装したこととも同じとなる。
// いわゆるインターフェースを満たすために外部から提供されている機能をラップしている感じ
type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// Rows はクエリの結果です。カーソルは結果セットの最初の行の手前から始まります。
// 行から行へ進むには、Next を使用します。
type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

// Rowsを閉じて、それ以上の列挙を防ぐ。そのあとNExtが呼び誰てもfalseを返す。
// それ以上の結果が存在しない場合は行は自動的に閉じられ、Errの結果をチェックするだけで良い。
func (r SqlRow) Close() error {
	return r.Rows.Close()
}