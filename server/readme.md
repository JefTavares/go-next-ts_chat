## Configurando o Makefile

Basicamente o container do postgres

executar o container do postgres
//cria o banco go-chat
docker exec -it postgres15 createdb --username=root --owner=root go-chat

//executa o Makefile,no caso o comando dentro do postgres
make postgres

### Dicas container postgres

//entra no banco
docker exec -it postgres15 psql
//dentro do psql
\l --para listar todos os databases
\c go-chat --conecta no database
\q --sai do cli do postgres
\d --describe schema (Database)
\d users (descreve a tabela users)

### instalação do go migration

//go migration
instalar o brew (https://linux.how2shout.com/how-to-install-brew-ubuntu-20-04-lts-linux/)
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
(echo; echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"') >> /home/$USER/.zshrc
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

brew install golang-migration

depois roda:
//ele cria os arquivos up e down
migrate create -ext sql -dir db/migrations add_users_table

migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

### Pacotes utilizados

sql library: https://pkg.go.dev/database/sql
lib/pq: https://github.com/lib/pq
golang-migrate: https://github.com/golang-migrate/mig...
golang-jwt: https://github.com/golang-jwt/jwt
bcrypt: https://pkg.go.dev/golang.org/x/crypt...
gin: https://github.com/gin-gonic/gin
