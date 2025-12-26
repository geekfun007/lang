## Docker 常用命令详解

### 1. 系统信息命令

#### **版本和系统信息**

```bash
# 查看Docker版本
docker version
docker --version

# 查看Docker系统信息
docker info

# 查看Docker系统资源使用
docker system df
docker system events
```

#### **系统管理**

```bash
# 清理未使用的资源
docker system prune
docker system prune -a  # 清理所有未使用的资源
docker system prune --volumes  # 包括卷

# 查看Docker事件
docker events
docker events --filter container=mycontainer
```

### 2. 镜像管理命令

#### **镜像操作**

```bash
# 搜索镜像
docker search nginx
docker search ubuntu --limit 5

# 拉取镜像
docker pull nginx
docker pull nginx:1.20
docker pull ubuntu:20.04

# 查看本地镜像
docker images
docker image ls
docker images -a  # 包括中间层

# 删除镜像
docker rmi nginx
docker rmi nginx:1.20
docker image rm nginx
docker rmi $(docker images -q)  # 删除所有镜像
```

#### **镜像构建和标签**

```bash
# 构建镜像
docker build -t myapp:latest .
docker build -t myapp:v1.0 -f Dockerfile.prod .

# 给镜像打标签
docker tag nginx:latest myregistry/nginx:v1.0
docker tag nginx:latest myregistry/nginx:latest

# 推送镜像
docker push myregistry/nginx:v1.0
docker push myregistry/nginx:latest
```

#### **镜像导入导出**

```bash
# 保存镜像到文件
docker save -o nginx.tar nginx:latest
docker save nginx:latest > nginx.tar

# 从文件加载镜像
docker load -i nginx.tar
docker load < nginx.tar

# 导出容器
docker export container_name > container.tar

# 导入容器为镜像
docker import container.tar myimage:latest
```

### 3. 容器管理命令

#### **容器生命周期**

```bash
# 创建并运行容器
docker run nginx
docker run -d nginx  # 后台运行
docker run -it ubuntu bash  # 交互式运行

# 启动/停止/重启容器
docker start container_name
docker stop container_name
docker restart container_name

# 暂停/恢复容器
docker pause container_name
docker unpause container_name

# 强制终止容器
docker kill container_name
docker kill -s SIGTERM container_name
```

#### **容器查看和删除**

```bash
# 查看容器
docker ps  # 运行中的容器
docker ps -a  # 所有容器
docker ps -l  # 最新创建的容器

# 删除容器
docker rm container_name
docker rm -f container_name  # 强制删除运行中的容器
docker rm $(docker ps -a -q --filter 'status=exited')  # 删除所有容器
```

#### **容器交互**

```bash
# 进入容器
docker exec -it container_name bash
docker exec -it container_name sh

# 在容器中执行命令
docker exec container_name ls -la
docker exec container_name ps aux

# 查看容器日志
docker logs container_name
docker logs -f container_name  # 实时跟踪
docker logs --tail 100 container_name
```

### 4. 网络管理命令

#### **网络操作**

```bash
# 查看网络
docker network ls
docker network inspect bridge

# 创建网络
docker network create mynetwork
docker network create --driver bridge mynetwork

# 连接容器到网络
docker network connect mynetwork container_name
docker network disconnect mynetwork container_name

# 删除网络
docker network rm mynetwork
docker network prune  # 删除未使用的网络
```

#### **端口映射**

```bash
# 端口映射
docker run -p 8080:80 nginx
docker run -p 127.0.0.1:8080:80 nginx
docker run -P nginx  # 随机端口映射

# 查看端口映射
docker port container_name
```

### 5. 数据卷管理命令

#### **卷操作**

```bash
# 查看卷
docker volume ls
docker volume inspect myvolume

# 创建卷
docker volume create myvolume
docker volume create --driver local myvolume

# 删除卷
docker volume rm myvolume
docker volume prune  # 删除未使用的卷

# 使用卷
docker run -v myvolume:/data nginx
docker run -v /host/path:/container/path nginx
```

#### **挂载选项**

```bash
# 绑定挂载
docker run -v /host/path:/container/path nginx
docker run -v /host/path:/container/path:ro nginx  # 只读

# 临时文件系统
docker run --tmpfs /tmp nginx

# 多个挂载
docker run -v /host1:/container1 -v /host2:/container2 nginx
```

### 6. 容器配置命令

#### **资源限制**

```bash
# 内存限制
docker run -m 512m nginx
docker run --memory=1g nginx
docker run --memory-swap=2g nginx

# CPU限制
docker run --cpus=2 nginx
docker run --cpu-shares=512 nginx
docker run --cpuset-cpus="0,1" nginx

# 其他限制
docker run --pids-limit=100 nginx
docker run --ulimit nofile=1024:1024 nginx
```

#### **环境变量和配置**

```bash
# 环境变量
docker run -e VAR1=value1 nginx
docker run -e VAR1=value1 -e VAR2=value2 nginx
docker run --env-file .env nginx

# 主机名和工作目录
docker run --hostname myhost nginx
docker run -w /app nginx

# 用户和权限
docker run --user 1000:1000 nginx
docker run --user root nginx
```

### 7. 容器编排命令

#### **Docker Compose**

```bash
# 启动服务
docker-compose up
docker-compose up -d  # 后台运行
docker-compose up --build  # 重新构建

# 停止服务
docker-compose down
docker-compose stop

# 查看服务状态
docker-compose ps
docker-compose logs

# 扩展服务
docker-compose up --scale web=3
```

```dockerfile
version: '1.0'

services:
  mysql:
    image: 'mysql:latest'
    volumes:
      - ./initdb.d:/docker-entrypoint-initdb.d
    ports:
      - 9910:3306
    environment:
      - MYSQL_DATABASE=gorm
      - MYSQL_USER=gorm
      - MYSQL_PASSWORD=gorm
      - MYSQL_RANDOM_ROOT_PASSWORD="yes"
```

### 8. 监控和调试命令

#### **容器监控**

```bash
# 查看容器资源使用
docker stats
docker stats container_name
docker stats container1 container2

# 查看容器进程
docker top container_name

# 查看容器详细信息
docker inspect container_name
```

#### **日志管理**

```bash
# 查看日志
docker logs container_name
docker logs -f container_name  # 实时跟踪
docker logs --since="2023-01-01" container_name
docker logs --until="2023-12-31" container_name
docker logs --tail 100 container_name
```

### 9. 清理命令

#### **系统清理**

```bash
# 清理未使用的资源
docker system prune
docker system prune -a  # 清理所有未使用的资源
docker system prune --volumes  # 包括卷

# 清理特定资源
docker container prune  # 清理停止的容器
docker image prune  # 清理悬空镜像
docker volume prune  # 清理未使用的卷
docker network prune  # 清理未使用的网络
```

### 10. 实用组合命令

#### **开发环境常用组合**

```bash
# 交互式开发环境
docker run -it --rm -v $(pwd):/app -w /app node:18 bash

# 后台运行服务
docker run -d --name service -p 8080:8080 myapp:latest

# 调试模式
docker run -it --rm --privileged ubuntu:20.04 bash
```

#### **生产环境常用组合**

```bash
# 生产环境部署
docker run -d \
  --name app \
  --restart=always \
  -p 8080:80 \
  -e NODE_ENV=production \
  -v app-data:/data \
  myapp:latest
```

### 11. 常用参数速查表

| 参数        | 说明       | 示例                                   |
| ----------- | ---------- | -------------------------------------- |
| `-d`        | 后台运行   | `docker run -d nginx`                  |
| `-i`        | 交互式     | `docker run -i ubuntu`                 |
| `-t`        | 分配伪终端 | `docker run -t ubuntu`                 |
| `-it`       | 交互式终端 | `docker run -it ubuntu bash`           |
| `--name`    | 容器名称   | `docker run --name myapp nginx`        |
| `-p`        | 端口映射   | `docker run -p 8080:80 nginx`          |
| `-v`        | 数据卷挂载 | `docker run -v /host:/container nginx` |
| `-e`        | 环境变量   | `docker run -e VAR=value nginx`        |
| `--rm`      | 自动删除   | `docker run --rm ubuntu`               |
| `--restart` | 重启策略   | `docker run --restart=always nginx`    |
| `-w`        | 工作目录   | `docker run -w /app nginx`             |
| `--network` | 网络模式   | `docker run --network host nginx`      |
| `-m`        | 内存限制   | `docker run -m 512m nginx`             |
| `--cpus`    | CPU 限制   | `docker run --cpus=2 nginx`            |

### 12. 最佳实践建议

#### **命令使用最佳实践**

1. **使用描述性名称**：`--name web-server`
2. **设置重启策略**：`--restart=always`
3. **合理设置资源限制**：`-m 512m --cpus=1`
4. **使用数据卷持久化**：`-v mydata:/data`
5. **避免使用特权模式**：除非必要
6. **定期清理资源**：`docker system prune`

#### **安全考虑**

1. **使用非 root 用户**：`--user 1000:1000`
2. **限制资源使用**：设置内存和 CPU 限制
3. **使用只读文件系统**：`--read-only`
4. **定期更新镜像**：使用最新版本
5. **扫描镜像漏洞**：`docker scan`

我来为您详细解释 Dockerfile 的语法、指令和最佳实践：

## Dockerfile 详解

### 1. Dockerfile 基本概念

#### **什么是 Dockerfile**

- Dockerfile 是一个文本文件，包含构建 Docker 镜像的指令
- 通过 `docker build` 命令执行 Dockerfile 中的指令
- 每个指令都会创建一个新的镜像层
- 遵循"层缓存"机制，提高构建效率

#### **基本结构**

```dockerfile
# 注释
FROM base_image
# 指令1
# 指令2
# ...
```

### 2. 常用指令详解

#### **FROM - 基础镜像**

```dockerfile
# 指定基础镜像
FROM ubuntu:20.04
FROM node:18-alpine
FROM python:3.9-slim
FROM nginx:latest

# 多阶段构建
FROM node:18 AS builder
FROM nginx:alpine AS production

# 使用官方镜像
FROM node:18
FROM python:3.9
FROM golang:1.19
```

#### **RUN - 执行命令**

```dockerfile
# 基本用法
RUN apt-get update
RUN apt-get install -y nginx

# 合并命令（减少层数，推荐）
RUN apt-get update && \
    apt-get install -y nginx && \
    rm -rf /var/lib/apt/lists/*

# 使用shell形式
RUN echo "Hello World"
RUN npm install

# 使用exec形式（推荐）
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "nginx"]

# 多行命令
RUN apt-get update && \
    apt-get install -y \
    nginx \
    curl \
    wget && \
    rm -rf /var/lib/apt/lists/*
```

#### **COPY - 复制文件**

```dockerfile
# 复制文件
COPY package.json /app/
COPY src/ /app/src/

# 复制目录
COPY . /app/

# 复制多个文件
COPY package.json package-lock.json /app/

# 使用通配符
COPY *.js /app/
COPY src/*.py /app/

# 复制并保持权限
COPY --chown=node:node package.json /app/
```

#### **ADD - 复制文件（增强版）**

```dockerfile
# 基本复制（与COPY相同）
ADD package.json /app/

# 自动解压压缩文件
ADD archive.tar.gz /app/

# 从URL下载文件
ADD https://example.com/file.tar.gz /app/

# 复制并解压
ADD archive.tar.gz /app/
```

#### **WORKDIR - 设置工作目录**

```dockerfile
# 设置工作目录
WORKDIR /app
WORKDIR /usr/src/app

# 后续命令在此目录执行
RUN pwd  # 输出: /app
COPY . .  # 复制到 /app

# 多层工作目录
WORKDIR /app
WORKDIR /app/src
```

#### **ENV - 设置环境变量**

```dockerfile
# 设置环境变量
ENV NODE_ENV=production
ENV PORT=3000
ENV DATABASE_URL=postgresql://localhost:5432/mydb

# 设置多个环境变量
ENV NODE_ENV=production \
    PORT=3000 \
    DATABASE_URL=postgresql://localhost:5432/mydb

# 使用变量
ENV APP_HOME=/app
WORKDIR $APP_HOME
```

#### **ARG - 构建参数**

```dockerfile
# 定义构建参数
ARG NODE_VERSION=18
ARG BUILD_ENV=production

# 使用构建参数
FROM node:${NODE_VERSION}
ENV NODE_ENV=${BUILD_ENV}

# 构建时传递参数
# docker build --build-arg NODE_VERSION=16 .
```

#### **EXPOSE - 暴露端口**

```dockerfile
# 暴露端口
EXPOSE 3000
EXPOSE 80 443

# 暴露多个端口
EXPOSE 3000 8080 9000

# 指定协议
EXPOSE 3000/tcp
EXPOSE 53/udp
```

#### **CMD - 默认命令**

```dockerfile
# 使用exec形式（推荐）
CMD ["npm", "start"]
CMD ["python", "app.py"]
CMD ["node", "server.js"]

# 使用shell形式
CMD npm start
CMD python app.py

# 带参数
CMD ["npm", "start", "--", "--port", "3000"]
```

#### **ENTRYPOINT - 入口点**

```dockerfile
# 基本用法
ENTRYPOINT ["python", "app.py"]

# 与CMD结合使用
ENTRYPOINT ["python", "app.py"]
CMD ["--help"]

# 使用shell形式
ENTRYPOINT python app.py
```

#### **VOLUME - 数据卷**

```dockerfile
# 创建数据卷
VOLUME ["/data"]
VOLUME ["/var/log", "/var/lib/mysql"]

# 指定挂载点
VOLUME /app/data
```

#### **USER - 设置用户**

```dockerfile
# 切换到指定用户
USER node
USER 1000:1000

# 创建用户并切换
RUN adduser --disabled-password --gecos '' appuser
USER appuser
```

### 3. 实际应用示例

#### **Node.js 应用**

```dockerfile
# 使用官方Node.js镜像
FROM node:18-alpine

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm ci --only=production

# 复制应用代码
COPY . .

# 创建非root用户
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001

# 切换到非root用户
USER nextjs

# 暴露端口
EXPOSE 3000

# 启动应用
CMD ["npm", "start"]
```

#### **Python 应用**

```dockerfile
# 使用Python官方镜像
FROM python:3.9-slim

# 设置环境变量
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# 设置工作目录
WORKDIR /app

# 安装系统依赖
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# 复制requirements.txt
COPY requirements.txt .

# 安装Python依赖
RUN pip install --no-cache-dir -r requirements.txt

# 复制应用代码
COPY . .

# 创建非root用户
RUN adduser --disabled-password --gecos '' appuser
USER appuser

# 暴露端口
EXPOSE 8000

# 启动应用
CMD ["python", "app.py"]
```

#### **Nginx 配置**

```dockerfile
# 使用Nginx官方镜像
FROM nginx:alpine

# 复制配置文件
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf

# 复制静态文件
COPY dist/ /usr/share/nginx/html/

# 暴露端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]
```

### 4. 多阶段构建

#### **前端应用多阶段构建**

```dockerfile
# 构建阶段
FROM node:18 AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

# 生产阶段
FROM nginx:alpine

# 复制构建结果
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制Nginx配置
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

#### **Go 应用多阶段构建**

```dockerfile
# 构建阶段
FROM golang:1.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 生产阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# 复制二进制文件
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
```

### 5. 高级指令

#### **HEALTHCHECK - 健康检查**

```dockerfile
# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:3000/health || exit 1

# 使用脚本
HEALTHCHECK --interval=30s --timeout=3s \
    CMD /app/healthcheck.sh || exit 1
```

#### **LABEL - 标签**

```dockerfile
# 添加标签
LABEL maintainer="your-email@example.com"
LABEL version="1.0"
LABEL description="My application"

# 多个标签
LABEL maintainer="your-email@example.com" \
      version="1.0" \
      description="My application"
```

#### **ONBUILD - 构建时触发**

```dockerfile
# 当此镜像被用作基础镜像时触发
ONBUILD COPY package.json /app/
ONBUILD RUN npm install
ONBUILD COPY . /app/
```

### 6. 最佳实践

#### **优化镜像大小**

```dockerfile
# 使用Alpine Linux
FROM node:18-alpine

# 合并RUN命令
RUN apk add --no-cache \
    python3 \
    make \
    g++ \
    && npm install \
    && apk del python3 make g++

# 清理缓存
RUN apt-get update && \
    apt-get install -y nginx && \
    rm -rf /var/lib/apt/lists/*

# 使用.dockerignore
# 排除不必要的文件
```

#### **安全最佳实践**

```dockerfile
# 使用非root用户
RUN adduser --disabled-password --gecos '' appuser
USER appuser

# 使用特定版本
FROM node:18.4.0-alpine

# 最小权限原则
RUN chmod 755 /app

# 只读文件系统
RUN mkdir -p /app/data
VOLUME /app/data
```

#### **性能优化**

```dockerfile
# 利用Docker缓存
# 先复制package.json，再安装依赖
COPY package*.json ./
RUN npm ci

# 然后复制源代码
COPY . .

# 使用多阶段构建
FROM node:18 AS builder
# 构建步骤...

FROM node:18-alpine AS production
# 生产步骤...
```

### 7. .dockerignore 文件

```dockerignore
# 忽略文件
node_modules
npm-debug.log
.git
.gitignore
README.md
.env
.nyc_output
coverage
.nyc_output
.coverage
.coverage/
.pytest_cache
__pycache__
*.pyc
*.pyo
*.pyd
.Python
env
pip-log.txt
pip-delete-this-directory.txt
.tox
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.log
.git
.mypy_cache
.pytest_cache
.hypothesis
```

### 8. 构建和运行

#### **构建镜像**

```bash
# 基本构建
docker build -t myapp:latest .

# 指定Dockerfile
docker build -f Dockerfile.prod -t myapp:prod .

# 传递构建参数
docker build --build-arg NODE_ENV=production -t myapp:prod .

# 多阶段构建
docker build --target production -t myapp:prod .

# 构建时传递参数
docker build --build-arg NODE_VERSION=16 -t myapp:latest .
```

#### **运行容器**

```bash
# 运行容器
docker run -p 3000:3000 myapp:latest

# 后台运行
docker run -d -p 3000:3000 myapp:latest

# 传递环境变量
docker run -e NODE_ENV=production -p 3000:3000 myapp:latest
```

### 9. 调试 Dockerfile

#### **调试技巧**

```dockerfile
# 使用交互式shell
FROM ubuntu:20.04
RUN apt-get update && apt-get install -y vim
CMD ["/bin/bash"]

# 添加调试信息
RUN echo "Building for environment: $NODE_ENV"
RUN ls -la /app
```

#### **构建调试**

```bash
# 查看构建过程
docker build --no-cache -t myapp:debug .

# 查看中间层
docker history myapp:debug

# 进入构建失败的层
docker run -it <intermediate_image_id> /bin/bash
```

### 10. 常用指令总结

| 指令          | 说明             | 示例                                       |
| ------------- | ---------------- | ------------------------------------------ |
| `FROM`        | 基础镜像         | `FROM node:18`                             |
| `RUN`         | 执行命令         | `RUN npm install`                          |
| `COPY`        | 复制文件         | `COPY . /app`                              |
| `ADD`         | 复制文件（增强） | `ADD archive.tar.gz /app`                  |
| `WORKDIR`     | 工作目录         | `WORKDIR /app`                             |
| `ENV`         | 环境变量         | `ENV NODE_ENV=production`                  |
| `ARG`         | 构建参数         | `ARG NODE_VERSION=18`                      |
| `EXPOSE`      | 暴露端口         | `EXPOSE 3000`                              |
| `CMD`         | 默认命令         | `CMD ["npm", "start"]`                     |
| `ENTRYPOINT`  | 入口点           | `ENTRYPOINT ["python", "app.py"]`          |
| `VOLUME`      | 数据卷           | `VOLUME ["/data"]`                         |
| `USER`        | 用户             | `USER node`                                |
| `HEALTHCHECK` | 健康检查         | `HEALTHCHECK CMD curl -f http://localhost` |
| `LABEL`       | 标签             | `LABEL maintainer="email@example.com"`     |
| `ONBUILD`     | 构建时触发       | `ONBUILD COPY . /app`                      |
