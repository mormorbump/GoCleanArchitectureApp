//UserControllerはusecase.UserInteractorを内包し、
//func NewUserController(sqlHandler database.ISqlHandler) *UserController
//でインスタンスを作る際にdatabase.ISqlHandlerを引数にとる。
// こうしておいて、
//handlerやrouterでNewUserController()する際にNewSqlHandler()を引数にとれば、
//どちらのインスタンスも無事生成され、Controller構造体にInteractorを入れ、InteractorにRepositoryを入れ、RepositoryにSqlHandlerを入れることができる。

package controllers

// 外から二番目なのでinfrastructure意外依存できる
import (
	"app/domain"
	"app/interfaces/database"
	"app/usecases"
	"strconv"
)

type UserController struct {
	Interactor usecases.UserInteractor
}

func NewUserController(sqlHandler database.ISqlHandler) *UserController {
	// UserControllerインスタンスを作るために、Interactorのインスタンスを定義するために、RepositoryとSqlHandlerを代入している。
	// 同じ層のものは依存してオッケーなのでこれでよし！
	return &UserController{
		Interactor: usecases.UserInteractor{
			UserRepository: &database.UserRepository{
				ISqlHandler: sqlHandler,
			},
		},
	}
}

// Contextはgin.Contextを使いたいので、同じ機能をもったインターフェースを同階層に定義。
//　今回はHTTPリクエストの時にくる情報をひとまとめにしてるものがContextって感じ。
func (controller UserController) Create(c Context)  {
	u := domain.User{}
	c.Bind(&u) // Bindの引数はinterface{}なのでどのドメインも入る。
	user, err := controller.Interactor.Add(u) // リポジトリでの永続化(Store)を行うため
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user) // gin.Contextに格納されるのでレスポンスはよしなにしてくれる。
}

func (controller *UserController) Index(c Context)  {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller UserController) Show(c Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
	}
	c.JSON(200, user)
}