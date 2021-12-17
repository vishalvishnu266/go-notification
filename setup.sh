# docker run -d -v ~/mongoDockerdata:/data/db --name mongodb -p 27017:27017  mongo:4.4.6 
# docker run -d --name rp -p 6379:6379 redis
#  npm install -g redis-commander
# redis-commander
# use push_notification;
# db.createUser(
#   {
#     user: "apptino",
#     pwd: "password",
#     roles: [ { role: "readWrite", db: "push_notification" } ]
#   }
# )

# go mod init github.com/growmax/noti 

# go get -u github.com/gin-gonic/gin
# go get -u google.golang.org/grpc
# go get github.com/joho/godotenv
# go get go.mongodb.org/mongo-driver/mongo
# go get -u github.com/SherClockHolmes/webpush-go
# go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
# go get -u github.com/golang-jwt/jwt
# wget https://raw.githubusercontent.com/centrifugal/centrifugo/master/internal/apiproto/api.proto -O centrifugo.proto
# protoc -I ./ centrifugo.proto --go_out=. --go-grpc_out=.


go fmt ./...

