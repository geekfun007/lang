# ============================================
# 多阶段构建 Dockerfile - 前端 React 应用
# ============================================

# -------------------- 阶段 1: 构建阶段 --------------------
FROM node:18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 添加标签
LABEL stage=builder
LABEL description="Frontend build stage"

# 复制 package 文件（利用 Docker 缓存）
COPY package*.json ./

# 安装依赖
RUN npm ci --only=production && \
    npm cache clean --force

# 复制源代码
COPY . .

# 构建应用
RUN npm run build

# -------------------- 阶段 2: 生产阶段 --------------------
FROM nginx:alpine

# 添加标签
LABEL maintainer="your-email@example.com"
LABEL version="1.0"
LABEL description="Frontend React application with Nginx"

# 安装 curl 用于健康检查
RUN apk add --no-cache curl

# 从构建阶段复制构建产物
COPY --from=builder /app/build /usr/share/nginx/html

# 复制自定义 Nginx 配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost/ || exit 1

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
