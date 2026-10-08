# EVENTHUB API
![Go](https://img.shields.io/badge/go-17.11-black?style=for-the-badge&logo=go&logoColor=white&labelColor=blue)
![Gin](https://img.shields.io/badge/Gin--gonic-v1.12.0-black?style=for-the-badge&logo=gin&logoColor=black&labelColor=lightblue)
![Postgresql](https://img.shields.io/badge/Postgresql-17.11-black?style=for-the-badge&logo=postgresql&logoColor=black&labelColor=lightyellow)
![redis](https://img.shields.io/badge/Redis-8.10.02-black?style=for-the-badge&logo=redis&logoColor=white&labelColor=FF4438)


Application for discovering Events and Communities with multi-role based user


## Technologies

- **Go** 1.27.1 — A statically typed programming language used to build the API and manage its backend services.
- **Gin** 1.12.0 — A lightweight HTTP web framework used for routing and handling API requests.
- **PostgreSQL** 17.11 — A relational database used to store and manage event-related data.
- **Redis** 8.10.2 — An in-memory data store used for caching and fast data operations.
- **Swagger** — Generating interactive API documentation and testing endpoints from the API's OpenAPI specification.


## Features

### **Authentication** : 
- Registration
- Login and Logout using Json Web Token (JWT)
### **Users**: 
- Edit Profile with uploading image using multipart.FileHeader
- Join Events and Bookmark Events
- Join and Leave Communities
- Create and update events for organizer

### **Events**:
- Event list
- Discover events with search & filter 

### **Communities**:
- Discover communities with search & filter


# Usage Instruction


### Database Setup

1. Pull Docker Image Postgresql From [dockerhub](https://hub.docker.com/_/postgres)

```sh
docker pull postgres:17.11
```

2. Create container

```sh
docker run --name {container_name} -d  -p {hostPort:5432} -v {volume:/var/lib/postgresql/data} --env-file ./. postgres:17.11
```


### Cache Setup

1. Pull Redis Image from [dockerhub](https://hub.docker.com/_/redis)

```zsh
docker pull redis:8.10.2
```
2. setup your redis configuration
```
bind 0.0.0.0
port 6379

daemonize no
logfile ""
loglevel notice

save 3600 1
save 300 100
save 60 10000

appendonly yes
appendfsync everysec

maxmemory 256mb
maxmemory-policy allkeys-lru

protected-mode yes

user default off
user {superuser} on >{password} ~* &* +@all
user {youruser} on >{yourpass} resetchannels ~{prefixkeys}:* +@string +ping +del +@read
```
3. create docker cache container
```sh
docker run --name {container_name} -p {hostPort:6379} -v {redis.conf}:/usr/local/etc/redis/redis.conf -v {volume}:/data redis:8.10.2 redis-server /usr/local/etc/redis/redis.conf
```

4. SET UP server environment

```sh
#  CONFIG DB
DB_USER={your_db_user}
DB_PASS={your_db_pass}
DB_HOST={your_db_host}
DB_PORT{your_db_port}
DB_NAME={your_db_name}

# PORT BE
PORT={your_server_port}
SERVER_HOST={your_server_host}


# JWT

JWT_SECRET={rand-jwt-secret-key}
JWT_ISSUER={jwt-issuer}

#redis
REDIS_USERNAME={your_redis_username}
REDIS_PASSWORD={your_redis_password}
REDIS_PORT={your_redis_port}
REDIS_HOST={your_redis_host}
```
### Clone Repository

```sh
git clone https://github.com/fajarworks/backend-eventhub.git
```
**Download Dependencies**
```sh
go mod download
```

### API Documentation
Swagger documentation is available at:


```bash
http://localhost:8081/swagger/index.html
```

## License
Distributed under the MIT license. See [License](https://opensource.org/license/mit) for more information