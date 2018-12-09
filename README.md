
# Go testアプリ

### 概要

|項目     |内容                          |対応状況 |備考    |
|--------|------------------------------|--------|-------|
|REST    |JWT対応                       |済み   |        |
|        |AWS 署名付きS3 Put返却         |済み   |        |
|        |AWS S3 Put擬似                |済み   |        |
|        |JSON valueをNullまたは空にする方法 |対応未   |    |
|        |Websocket                     |済み   |    |
|microservice |                         |済み   |    |
|             |HTTP Transport<br>(Gorilla/mux http tranport --> Echo/Gin http tranport)|対応未 |    |
|yaml/構成定義読み込み |viper使用              |済み   |    |
|O/R mapper |Gorm connection pool |済み   |    |
|Docker   　|                      |済み   |    |


### Go環境
* package管理 dep前提
```
docker build .
  
  ...
  
  Successfully built <コンテナID>

docker run -p 8080:8080 -td <コンテナID>


```




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