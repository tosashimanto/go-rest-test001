
# Go testアプリ

### 概要

|項目     |内容                          |       |
|--------|------------------------------|------|
|REST    |AW  S3　put 擬似               |済み   |
|        |JSON valueをNullまたは空にする   |未対応   |
|        |Websocket                     |済み   |
|microservice |Gorilla/mux http tranport --> Echo/Gin http tranport|未対応   |
|yaml/構成定義   | viper            |済み   |
|O/R mapper |Gorm connection pool |
|Docker   　|                      |未対応   |


### Go環境

docker build .
  
  Successfully built <コンテナID>

docker run -p 8080:8080 -td <コンテナID>
docker ps -a




### パッケージ管理
```

  - govendorの場合(Heroku用)
  go get -u github.com/kardianos/govendor
  govendor init
  govendor fetch +out

  - dep の場合 
  go get -u github.com/golang/dep/cmd/dep

  dep init                        
  dep ensure
  
  参考:                        
  dep ensure -update                # 全ての依存関係のバージョンを更新
  dep ensure github.com/pkg/errors  # プロジェクトに依存関係を追加
  
```



### Build、実行
```
  go build main.go
  go run main.go
```