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
  imageBaseUrl: 'http://xxxxxxxx/image/',
  analyticsTrackingId: 'Google Analytics Tracking Id'
}
```


```
npm install
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
  imageBaseUrl: 'http://localhost:8080/image/',
  analyticsTrackingId: 'Google Analytics Tracking Id'
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

# nginx (development sample)

```
map $http_accept $webp_suffix {
  default   "";
  "~*image/webp"  ".webp";
}

server {
    listen       8080;
    server_name  dev.blog.temo.xyz;

    root /var/www/blog.temo.xyz/frontend/dist;
    index index.html;
    charset utf-8;

    location /api {
        proxy_pass http://127.0.0.1:3000;
    }

    location /image {
      alias /var/www/blog.temo.xyz/backend/public/image;
      location ~* \.(png|jpe?g)$ {
        add_header Vary Accept;
        try_files $uri$webp_suffix $uri =404;
      }
    }

    gzip on;
    gzip_types text/css application/javascript application/json application/font-woff application/font-tff application/octet-stream;

    etag off;
    add_header X-Frame-Options SAMEORIGIN;
    add_header X-XSS-Protection "1; mode=block";
    add_header X-Content-Type-Options nosniff;
    client_max_body_size 10M;
    client_header_buffer_size 1k;
    large_client_header_buffers 4 8k;
}
```
