# blog

ブログ

- frontend Nuxt.js 2
- backend Golang + gin-gonic + mgo
- database mongodb

# install (production)

## frontend

```
cd frontend
```

api の URL と 画像の base url を定義

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

api の URL と 画像の base url を定義

env.development.js

```
module.exports = {
  apiBaseUrl: 'https://blog.temo.local/api/',
  imageBaseUrl: 'https://blog.temo.local/image/',
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
ALLOW_ORIGINS=https://blog.temo.local
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

## nginx (development sample)

ssl 作成 (mac)

```
cd docker/ssl
mkcert "*.temo.local"
```

nginx install

```
brew install nginx
```

/usr/local/etc/nginx/servers/blog.conf

```
map $http_accept $webp_suffix {
  default   "";
  "~*image/webp"  ".webp";
}

server {
    listen       443 ssl http2;
    server_name  blog.temo.local;

    index index.html;
    charset utf-8;

    ssl_certificate /repository_path/src/docker/ssl/_wildcard.temo.local.pem;
    ssl_certificate_key /repository_path/docker/ssl/_wildcard.temo.local-key.pem;
    ssl_protocols TLSv1.2;

    # local go
    location /api {
        proxy_pass http://127.0.0.1:8081;
    }

    location /image {
      alias /repository_path/backend/public/image;
      location ~* \.(png|jpe?g)$ {
        add_header Vary Accept;
        try_files $uri$webp_suffix $uri =404;
      }
    }

    # local node dev server
    location / {
        proxy_pass http://127.0.0.1:3333;
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

## mongodb

```
cd docker
docker-compose up -d
```
