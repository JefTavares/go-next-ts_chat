package user

import (
	"context"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "secret"
)

// timeout é utilizado no contexto
type service struct {
	Repository
	timeout time.Duration
}

// retorna a interface Service, com seus metodos implementados aqui em baixo
func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second, //Timeout é opcional no contexto aqui estou atribuindo 2 segundos
	}
}

// obs: lembrando que o service é chamado na camada de handler
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	//cria o ctx e o cancel
	ctx, cancel := context.WithTimeout(c, s.timeout) //cria um contexto e define os 2 segundos setados no func NewService
	defer cancel()

	//encripta a senha
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	//Cria a struct de usuario com os dados da request
	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	//monta a resposta
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil

}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	//Password que veio na request vs pass word no banco de dados
	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
