
nohup go run 7OxCloud-srv/user/main.go > 7OxCloud-srv/user/output.log &

nohup go run 7OxCloud-srv/video/main.go > 7OxCloud-srv/video/output.log &

nohup go run 7OxCloud-srv/interaction/main.go > 7OxCloud-srv/interaction/output.log &

nohup go run 7OxCloud-srv/social/main.go > 7OxCloud-srv/social/output.log &

nohup go run 7OxCloud-srv/cron/main.go > 7OxCloud-srv/cron/output.log &

nohup go run 7OxCloud-api/main.go -p 8080 > 7OxCloud-api/output.log &

#nohup go run 7OxCloud-api/main.go -p 8081 > 7OxCloud-api/output.log &
