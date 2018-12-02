# blog.temo.xyz

ブログ

- frontend Nuxt.js 2
- backend Golang + gin-gonic + mgo
- database mongodb

# install (production)

## frontend

```
cd frontend
```

api のURLと 画像の base url を定義

env.production.js
```
module.exports = {
  apiBaseUrl: 'https://xxxxxxxx/api/',
  imageBaseUrl: 'http://xxxxxxxx/image/'
}
```


```
npm run build
```

## backend

cwebp が使えること
```
Centos7
yum install libwebp-tools

mac
brew install webp
```

```
cd backend
```

.env 作成
```
DB_NAME=databaseName
DB_HOST=localhost
ALLOW_ORIGINS=https://xxxxxxxx
```

collection + index 作成
```
go run task/createCollection.go
```

ユーザ作成
```
go run task/createUser.go アカウント パスワード
```

build + start
```
go build main.go
./main &
```


# install (development)

## frontend

```
cd frontend
```

api のURLと 画像の base url を定義

env.development.js
```
module.exports = {
  apiBaseUrl: 'http://localhost:8080/api/',
  imageBaseUrl: 'http://localhost:8080/image/'
}
```


```
npm install
npm run dev
```

## backend

cwebp が使えること
```
Centos7
yum install libwebp-tools

mac
brew install webp
```

```
cd backend
```

.env 作成
```
DB_NAME=databaseName
DB_HOST=localhost
ALLOW_ORIGINS=http://localhost:3333
```

collection + index 作成
```
go run task/createCollection.go
```

ユーザ作成
```
go run task/createUser.go アカウント パスワード
```


```
gin run main.go
```

