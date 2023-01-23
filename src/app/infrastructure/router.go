package infrastructure

import (
	// routerかくのめんどいのでginで作る
	// 1.3.0のginはバグがあるので注意。フツーにgithubからinportで良さげ https://github.com/gin-gonic/gin/issues/1887#issuecomment-490708131
	"github.com/gin-gonic/gin"
	// infrastractureは一番外側なのでもちろんcontrollersも依存できる。
	"app/interfaces/controllers"
)

// Router gin.Engineの型を定義。
// Engineはフレームワークのインスタンスで、マルチプレクサ、ミドルウェア、コンフィギュレーション設定などが含まれる。
// New()やDefault()を使って、Engineのインスタンスを生成
var Router *gin.Engine

func init() {
	// Engineインスタンスを生成している。
	// デフォルトでは、LoggerとRecoveryミドルウェアが既にアタッチされたEngineインスタンスが返され流。
	routerConfig := gin.Default()

	// 同じ階層のNewSqlHandlerはimportなしで呼べるよ
	userController := controllers.NewUserController(NewSqlHandler())

	// ginのお作法。第一引数にpathを、第二引数にhandler関数を代入。
	// なので*gin.Contextを引数にとった無名関数を代入し、中でuserController.Create(c)を実行。
	// *gin.Contextにリクエストなどの情報が入ってるので、それをcとしてControllerにも代入している。
	// こうすることによって情報をもったままPOST /usersが呼ばれた時にUserController#Createが実行される。
	routerConfig.POST("/users", func(c *gin.Context) { userController.Create(c) })
	routerConfig.GET("/users", func(c *gin.Context) { userController.Index(c) })
	routerConfig.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = routerConfig // 外から呼び出すのはvar Routerなので、Configを代入。
}
