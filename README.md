
# repositorio
docker push vremerson/codeeducation:1.0

# ver tamnho da imagem
docker images | grep vremerson/codeeducation

# Info da imagem
vremerson/codeeducation    1.0   750afd9f9240  5 seconds ago       1.46MB

# Comandos GO
go mod init example

go build ./...

go test ./...

go test -v ./...

go test -cover ./...

go test -cover -coverprofile=c.out ./...

go tool cover -html=c.out -o coverage.html

godoc -http=:6060

godoc -play -http=:6060