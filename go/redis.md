## Redis 基本命令行详解

### 1. 连接 Redis 服务器

```bash
# 连接到本地Redis服务器（默认端口6379）
redis-cli

# 连接到指定主机和端口
redis-cli -h hostname -p port

# 连接到指定主机、端口和数据库
redis-cli -h hostname -p port -n database_number

# 使用密码连接
redis-cli -a password

# 连接到远程Redis服务器
redis-cli -h 192.168.1.100 -p 6379 -a mypassword
```

### 2. 基本服务器操作

```bash
# 查看Redis信息
INFO

# 查看特定信息
INFO memory
INFO cpu
INFO replication


# 清空当前数据库
FLUSHDB

# 清空所有数据库
FLUSHALL

# 查看所有键
KEYS *

# 查看匹配模式的键
KEYS user:*
KEYS *session*

# 查看键的总数
DBSIZE

# 查看键的类型
TYPE keyname

# 查看键的生存时间
TTL keyname

# 设置键的生存时间（秒）
EXPIRE keyname 3600

# 设置键的生存时间（毫秒）
PEXPIRE keyname 3600000

# 删除键
DEL keyname

# 检查键是否存在
EXISTS keyname

# 重命名键
RENAME oldkey newkey
```

### 3. 字符串（String）操作

```bash
# 设置键值
SET key value [EX seconds] [PX milliseconds] [NX|XX]

# 获取值
GET key

# 设置多个键值
MSET key1 value1 key2 value2 key3 value3

# 获取多个值
MGET key1 key2 key3

# 追加字符串
APPEND key1 " +1"

# 原子递增指定数值
INCRBY key 5

DECR key # 原子递减
INCR key # 原子递增
DECRBY key 3 # 原子递减指定数值


```

### 4. 哈希（Hash）操作

```bash
# 设置哈希字段
HSET user:1 name "John" age 30 email "john@example.com"

# 获取哈希字段
HGET user:1 name

# 获取多个字段
HMGET user:2 name age city

# 获取所有字段和值
HGETALL user:1

# 获取所有字段名
HKEYS user:1

# 获取所有值
HVALS user:1

# 检查字段是否存在
HEXISTS user:1 name

# 删除字段
HDEL user:1 age

# 原子递增字段值
HINCRBY user:1 age 1

# 获取字段数量
HLEN user:1

```

### 5. 列表（List）操作

```bash
# 从左侧推入元素
LPUSH mylist "item1" "item2" "item3"

# 从右侧推入元素
RPUSH mylist "item4" "item5"

# 从左侧弹出元素
LPOP mylist
LPOP mylist 2 # 删除2个

# 从右侧弹出元素
RPOP mylist
RPOP mylist 2 # 删除2个

# 获取指定索引的元素
LINDEX mylist 0

# 删除指定值的元素
LREM mylist 1 "item" # 删除1个匹配的元素
LREM mylist 0 "item" # 删除所有匹配的元素

# 获取指定范围的元素
LRANGE mylist 0 -1

# 获取列表长度
LLEN mylist

# 设置指定索引的元素
LSET mylist 0 "new item"
```

### 6. 集合（Set）操作

```bash
# 添加元素到集合
SADD myset "member1" "member2" "member3"

# 检查成员是否存在
SISMEMBER myset "member1"

# 删除成员
SREM myset "member1"

# 获取集合大小
SCARD myset

# 获取集合所有成员
SMEMBERS myset
```

### 7. 有序集合（Sorted Set）操作

```bash
# 添加成员和分数
ZADD leaderboard 100 "player1" 200 "player2" 150 "player3"

# 获取成员分数
ZSCORE leaderboard "player1"

# 获取成员排名（从低到高）
ZRANK leaderboard "player1"

# 删除成员
ZREM leaderboard "player1"

# 获取集合大小
ZCARD leaderboard

# 原子递增成员分数
ZINCRBY leaderboard 50 "player2"
```

### 8. 发布订阅（Pub/Sub）

```bash
# 订阅频道
SUBSCRIBE channel1 channel2

# 发布消息
PUBLISH channel1 "Hello World"

# 模式订阅
PSUBSCRIBE news.*

# 取消订阅
UNSUBSCRIBE channel1

# 取消模式订阅
PUNSUBSCRIBE news.*

# 查看活跃频道
PUBSUB CHANNELS

# 查看频道订阅者数量
PUBSUB NUMSUB channel1
```

### 9. 事务操作

```bash
# 开始事务
MULTI

# 添加命令到事务队列
SET key1 value1
SET key2 value2
INCR counter

# 执行事务
EXEC

# 取消事务
DISCARD

# 监视键（乐观锁）
WATCH key1 key2
MULTI
SET key1 newvalue
EXEC
```

### 10. 持久化和备份

```bash
# 手动保存到磁盘
SAVE

# 异步保存到磁盘
BGSAVE

# 获取最后保存时间
LASTSAVE

# 查看持久化配置
CONFIG GET save

# 设置持久化配置
CONFIG SET save "900 1 300 10 60 10000"
```

### 11. 性能监控

```bash
# 查看慢查询日志
SLOWLOG GET 10

# 重置慢查询日志
SLOWLOG RESET

# 查看内存使用情况
MEMORY USAGE keyname

# 查看内存统计
MEMORY STATS

# 查看客户端连接
CLIENT LIST

# 杀死指定客户端
CLIENT KILL ip:port

# 设置客户端名称
CLIENT SETNAME myclient
```

### 12. 配置管理

```bash
# 查看所有配置
CONFIG GET *

# 查看特定配置
CONFIG GET maxmemory

# 设置配置
CONFIG SET maxmemory 100mb

# 重写配置文件
CONFIG REWRITE
```

### 13. 实用技巧

```bash
# 批量删除键（使用管道）
redis-cli --scan --pattern "user:*" | xargs redis-cli DEL

# 查看键的内存使用
redis-cli --bigkeys

# 监控实时命令
MONITOR

# 查看Redis版本
INFO server | grep redis_version

# 退出Redis客户端
EXIT
QUIT
```

这些是 Redis 最常用的命令行操作。Redis 是一个功能强大的内存数据库，支持多种数据结构，这些命令可以帮助您有效地管理和操作数据。
