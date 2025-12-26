# SQL 速查表

## 数据库操作

```sql
-- 创建数据库
CREATE DATABASE db_name;
CREATE DATABASE IF NOT EXISTS db_name CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 查看数据库
SHOW DATABASES;
SHOW CREATE DATABASE db_name;

-- 选择数据库
USE db_name;

-- 删除数据库
DROP DATABASE db_name;
DROP DATABASE IF EXISTS db_name;

-- 修改数据库
ALTER DATABASE db_name CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

---

## 数据表操作

```sql
-- 创建表
CREATE TABLE table_name (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建用户表
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    age INT,
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_username (username),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建订单表
CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    order_no VARCHAR(50) NOT NULL UNIQUE,
    total_amount DECIMAL(10, 2) NOT NULL,
    order_status ENUM('pending', 'paid', 'shipped', 'completed', 'cancelled') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_order_no (order_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 查看表
SHOW TABLES;
DESC table_name;
SHOW COLUMNS FROM table_name;
SHOW CREATE TABLE table_name;

-- 修改表
ALTER TABLE table_name ADD COLUMN column_name datatype;
ALTER TABLE table_name DROP COLUMN column_name;
ALTER TABLE table_name MODIFY COLUMN column_name new_datatype;
ALTER TABLE table_name CHANGE old_name new_name datatype;
ALTER TABLE table_name RENAME TO new_table_name;

-- 删除表
DROP TABLE table_name;
DROP TABLE IF EXISTS table_name;
TRUNCATE TABLE table_name;  -- 清空数据但保留结构

-- 复制表结构
CREATE TABLE new_table LIKE old_table;

-- 复制表结构和数据
CREATE TABLE new_table AS SELECT * FROM old_table;
```

---

## 数据类型

### 数值类型

| 类型              | 大小   | 范围                      | 用途     |
| ----------------- | ------ | ------------------------- | -------- |
| TINYINT UNSIGNED  | 1 字节 | -128 到 127               | 小整数   |
| SMALLINT UNSIGNED | 2 字节 | -32768 到 32767           | 普通整数 |
| INT UNSIGNED      | 4 字节 | -2147483648 到 2147483647 | 整数     |
| BIGINT UNSIGNED   | 8 字节 | 非常大的整数              | 大整数   |
| FLOAT             | 4 字节 | 单精度浮点数              | 浮点数   |
| DOUBLE            | 8 字节 | 双精度浮点数              | 浮点数   |
| DECIMAL(M,D)      | 变长   | 精确小数                  | 金额等   |

### 字符串类型

| 类型       | 大小         | 用途         |
| ---------- | ------------ | ------------ |
| CHAR(n)    | 0-255 字节   | 定长字符串   |
| VARCHAR(n) | 0-65535 字节 | 变长字符串   |
| TEXT       | 0-65535 字节 | 长文本       |
| MEDIUMTEXT | 0-16MB       | 中等长度文本 |
| LONGTEXT   | 0-4GB        | 超长文本     |

### 日期时间类型

| 类型      | 格式                | 用途     |
| --------- | ------------------- | -------- |
| DATE      | YYYY-MM-DD          | 日期     |
| TIME      | HH:MM:SS            | 时间     |
| DATETIME  | YYYY-MM-DD HH:MM:SS | 日期时间 |
| TIMESTAMP | YYYY-MM-DD HH:MM:SS | 时间戳   |
| YEAR      | YYYY                | 年份     |

### 其他类型

```sql
-- 枚举类型
ENUM('value1', 'value2', 'value3')

-- 集合类型
SET('value1', 'value2', 'value3')

-- 布尔类型
BOOLEAN  -- 实际是 TINYINT(1)

-- 二进制类型
BLOB, BINARY, VARBINARY
```

---

## 约束

```sql
-- 主键约束
PRIMARY KEY
id INT AUTO_INCREMENT PRIMARY KEY

-- 唯一约束
UNIQUE
email VARCHAR(100) UNIQUE

-- 非空约束
NOT NULL
name VARCHAR(50) NOT NULL

-- 默认值
DEFAULT
status INT DEFAULT 1

-- 检查约束 (MySQL 8.0.16+)
CHECK
age INT CHECK (age >= 0 AND age <= 150)

-- 外键约束
FOREIGN KEY
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE
ON UPDATE CASCADE

-- 自增
AUTO_INCREMENT
id INT AUTO_INCREMENT
```

---

## 增删改查

### 插入 (INSERT)

```sql
-- 插入单条
INSERT INTO table_name (col1, col2) VALUES (val1, val2);

-- 插入多条
INSERT INTO table_name (col1, col2) VALUES
    (val1, val2),
    (val3, val4),
    (val5, val6);

-- 插入查询结果
INSERT INTO table1 (col1, col2)
SELECT col1, col2 FROM table2;

-- 存在则更新
INSERT INTO table_name (id, name) VALUES (1, 'John')
ON DUPLICATE KEY UPDATE name = VALUES(name);

-- 替换
REPLACE INTO table_name (id, name) VALUES (1, 'John');
```

### 查询 (SELECT)

```sql
-- 基本查询
SELECT * FROM table_name;
SELECT col1, col2 FROM table_name;

-- 去重
SELECT DISTINCT col1 FROM table_name;

-- 限制数量
SELECT * FROM table_name LIMIT 10;
SELECT * FROM table_name LIMIT 10 OFFSET 20;

-- 别名
SELECT col1 AS alias1, col2 AS alias2 FROM table_name AS t;
```

### 更新 (UPDATE)

```sql
-- 基本更新
UPDATE table_name SET col1 = val1 WHERE condition;

-- 更新多个字段
UPDATE table_name SET col1 = val1, col2 = val2 WHERE condition;

-- 基于计算更新
UPDATE table_name SET price = price * 1.1 WHERE id > 100;

-- 多表更新
UPDATE table1 t1
INNER JOIN table2 t2 ON t1.id = t2.id
SET t1.col1 = t2.col2
WHERE condition;
```

### 删除 (DELETE)

```sql
-- 基本删除
DELETE FROM table_name WHERE condition;

-- 删除所有记录
DELETE FROM table_name;

-- 使用 TRUNCATE（更快）
TRUNCATE TABLE table_name;

-- 多表删除
DELETE t1, t2
FROM table1 t1
INNER JOIN table2 t2 ON t1.id = t2.id
WHERE condition;
```

---

## 条件查询

```sql
-- 比较运算符
WHERE age = 25
WHERE age != 25
WHERE age <> 25
WHERE age > 25
WHERE age >= 25
WHERE age < 25
WHERE age <= 25

-- 逻辑运算符
WHERE age > 18 AND age < 60
WHERE city = 'Beijing' OR city = 'Shanghai'
WHERE NOT status = 0

-- 范围
WHERE age BETWEEN 18 AND 60
WHERE age NOT BETWEEN 18 AND 60

-- 列表
WHERE id IN (1, 2, 3, 4, 5)
WHERE id NOT IN (1, 2, 3)

-- 空值
WHERE email IS NULL
WHERE email IS NOT NULL

-- 模糊查询
WHERE name LIKE 'Zhang%'      -- 以Zhang开头
WHERE name LIKE '%san'        -- 以san结尾
WHERE name LIKE '%li%'        -- 包含li
WHERE name LIKE '_ohn'        -- 第一个字符任意，后面是ohn

-- 正则表达式
WHERE name REGEXP '^[A-Z]'    -- 以大写字母开头
WHERE email REGEXP '.*@gmail\\.com$'
```

---

## 排序与分组

### 排序 (ORDER BY)

```sql
-- 升序（默认）
SELECT * FROM users ORDER BY age ASC;

-- 降序
SELECT * FROM users ORDER BY age DESC;

-- 多字段排序
SELECT * FROM users ORDER BY age DESC, name ASC;

-- 按表达式排序
SELECT * FROM products ORDER BY price * quantity DESC;

-- 自定义排序
SELECT * FROM users ORDER BY
    CASE status
        WHEN 'active' THEN 1
        WHEN 'pending' THEN 2
        ELSE 3
    END;
```

### 分组 (GROUP BY)

```sql
-- 基本分组
SELECT city, COUNT(*) FROM users GROUP BY city;

-- 多字段分组
SELECT city, age, COUNT(*) FROM users GROUP BY city, age;

-- 过滤分组结果
SELECT city, COUNT(*) AS count
FROM users
GROUP BY city
HAVING count > 100;

-- WITH ROLLUP（汇总）
SELECT city, SUM(amount) FROM orders GROUP BY city WITH ROLLUP;
```

---

## 连接查询

```sql
-- INNER JOIN（内连接）
SELECT * FROM table1 t1
INNER JOIN table2 t2 ON t1.id = t2.id;

-- LEFT JOIN（左连接）
SELECT * FROM table1 t1
LEFT JOIN table2 t2 ON t1.id = t2.id;

-- RIGHT JOIN（右连接）
SELECT * FROM table1 t1
RIGHT JOIN table2 t2 ON t1.id = t2.id;

-- FULL OUTER JOIN（全外连接，MySQL不直接支持）
SELECT * FROM table1 t1
LEFT JOIN table2 t2 ON t1.id = t2.id
UNION
SELECT * FROM table1 t1
RIGHT JOIN table2 t2 ON t1.id = t2.id;

-- CROSS JOIN（交叉连接）
SELECT * FROM table1 CROSS JOIN table2;

-- SELF JOIN（自连接）
SELECT e1.name AS employee, e2.name AS manager
FROM employees e1
LEFT JOIN employees e2 ON e1.manager_id = e2.id;

-- 多表连接
SELECT *
FROM table1 t1
INNER JOIN table2 t2 ON t1.id = t2.t1_id
INNER JOIN table3 t3 ON t2.id = t3.t2_id;
```

---

## 聚合函数

```sql
-- COUNT：计数
SELECT COUNT(*) FROM users;
SELECT COUNT(DISTINCT city) FROM users;

-- SUM：求和
SELECT SUM(price) FROM products;

-- AVG：平均值
SELECT AVG(age) FROM users;

-- MAX：最大值
SELECT MAX(price) FROM products;

-- MIN：最小值
SELECT MIN(age) FROM users;

-- GROUP_CONCAT：连接
SELECT city, GROUP_CONCAT(name) FROM users GROUP BY city;

-- 组合使用
SELECT
    COUNT(*) AS total,
    AVG(price) AS avg_price,
    SUM(quantity) AS total_qty,
    MAX(price) AS max_price,
    MIN(price) AS min_price
FROM products;
```

---

## 字符串函数

```sql
-- 拼接
CONCAT('Hello', ' ', 'World')                    -- Hello World
CONCAT_WS('-', '2023', '11', '18')              -- 2023-11-18

-- 长度
LENGTH('Hello')                                  -- 5
CHAR_LENGTH('你好')                              -- 2

-- 大小写转换
UPPER('hello')                                   -- HELLO
LOWER('HELLO')                                   -- hello

-- 截取
SUBSTRING('Hello World', 1, 5)                   -- Hello
LEFT('Hello', 2)                                 -- He
RIGHT('Hello', 2)                                -- lo

-- 去除空格
TRIM('  Hello  ')                                -- Hello
LTRIM('  Hello')                                 -- Hello
RTRIM('Hello  ')                                 -- Hello

-- 替换
REPLACE('Hello World', 'World', 'MySQL')         -- Hello MySQL

-- 查找位置
LOCATE('o', 'Hello World')                       -- 5
POSITION('World' IN 'Hello World')               -- 7

-- 重复
REPEAT('Ha', 3)                                  -- HaHaHa

-- 反转
REVERSE('Hello')                                 -- olleH

-- 填充
LPAD('5', 3, '0')                               -- 005
RPAD('5', 3, '0')                               -- 500
```

---

## 日期函数

```sql
-- 当前日期时间
NOW()                                            -- 2023-11-18 10:30:45
CURRENT_DATE()                                        -- 2023-11-18
CURRENT_TIME()                                        -- 10:30:45
CURRENT_TIMESTAMP()                              -- 2023-11-18 10:30:45

-- 提取部分
YEAR('2023-11-18')                              -- 2023
MONTH('2023-11-18')                             -- 11
DAY('2023-11-18')                               -- 18
HOUR('10:30:45')                                -- 10
MINUTE('10:30:45')                              -- 30
SECOND('10:30:45')                              -- 45

-- 日期格式化
DATE_FORMAT(NOW(), '%Y-%m-%d')                  -- 2023-11-18
DATE_FORMAT(NOW(), '%Y年%m月%d日')               -- 2023年11月18日
TIME_FORMAT('10:30:45', '%H:%i')                -- 10:30

-- 日期计算
DATE_ADD(NOW(), INTERVAL 1 DAY)                 -- 明天
DATE_SUB(NOW(), INTERVAL 1 YEAR)                -- 去年

-- 日期差
DATEDIFF('2023-12-31', '2023-01-01')           -- 364
TIMESTAMPDIFF(YEAR, '2000-01-01', NOW())       -- 年龄计算

-- 星期
DAYOFWEEK(NOW())                                -- 1-7 (周日-周六)
WEEKDAY(NOW())                                  -- 0-6 (周一-周日)
DAYNAME(NOW())                                  -- Saturday

-- 获取月末
LAST_DAY('2023-02-15')                         -- 2023-02-28

-- Unix时间戳
UNIX_TIMESTAMP()                                -- 1700282445
FROM_UNIXTIME(1700282445)                       -- 2023-11-18 10:30:45
```

---

## 数学函数

```sql
-- 四舍五入
ROUND(3.1415, 2)                                -- 3.14
CEIL(3.14)                                      -- 4
FLOOR(3.99)                                     -- 3

-- 绝对值
ABS(-10)                                        -- 10

-- 幂运算
POW(2, 3)                                       -- 8
SQRT(16)                                        -- 4

-- 随机数
RAND()                                          -- 0-1之间的随机数
FLOOR(RAND() * 100)                             -- 0-99的随机整数

-- 三角函数
SIN(1)
COS(1)
TAN(1)

-- 取模
MOD(10, 3)                                      -- 1

-- 符号
SIGN(-5)                                        -- -1
SIGN(0)                                         -- 0
SIGN(5)                                         -- 1

-- 截断
TRUNCATE(3.1415, 2)                             -- 3.14
```

---

## 事务

```sql
-- 开启事务
START TRANSACTION;
BEGIN;

-- 提交事务
COMMIT;

-- 回滚事务
ROLLBACK;

-- 保存点
SAVEPOINT sp1;
ROLLBACK TO sp1;
RELEASE SAVEPOINT sp1;

-- 查看自动提交状态
SHOW VARIABLES LIKE 'autocommit';

-- 设置自动提交
SET autocommit = 0;  -- 关闭
SET autocommit = 1;  -- 开启

-- 隔离级别
SELECT @@transaction_isolation;
SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;

-- 完整事务示例
START TRANSACTION;

UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;

-- 检查是否成功
SELECT balance FROM accounts WHERE id IN (1, 2);

COMMIT;  -- 或 ROLLBACK;
```

---

## 索引

```sql
-- 创建索引
CREATE INDEX idx_name ON table_name(column_name);
CREATE UNIQUE INDEX idx_email ON users(email);
CREATE INDEX idx_multi ON table_name(col1, col2);

-- 使用 ALTER TABLE 创建
ALTER TABLE table_name ADD INDEX idx_name(column_name);
ALTER TABLE table_name ADD UNIQUE INDEX idx_email(email);

-- 创建全文索引
CREATE FULLTEXT INDEX idx_content ON articles(content);

-- 查看索引
SHOW INDEX FROM table_name;

-- 删除索引
DROP INDEX idx_name ON table_name;
ALTER TABLE table_name DROP INDEX idx_name;

-- 查看索引使用情况
EXPLAIN SELECT * FROM users WHERE email = 'test@example.com';

-- 强制使用索引
SELECT * FROM users FORCE INDEX (idx_email) WHERE email = 'test@example.com';

-- 忽略索引
SELECT * FROM users IGNORE INDEX (idx_email) WHERE email = 'test@example.com';
```

---

## 用户权限

```sql
-- 创建用户
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'username'@'%' IDENTIFIED BY 'password';  -- 允许远程

-- 修改密码
ALTER USER 'username'@'localhost' IDENTIFIED BY 'new_password';
SET PASSWORD FOR 'username'@'localhost' = 'new_password';

-- 授权
GRANT ALL PRIVILEGES ON database.* TO 'username'@'localhost';
GRANT SELECT, INSERT ON database.table TO 'username'@'localhost';
GRANT SELECT ON database.* TO 'username'@'%';

-- 常用权限
ALL PRIVILEGES    -- 所有权限
SELECT           -- 查询
INSERT           -- 插入
UPDATE           -- 更新
DELETE           -- 删除
CREATE           -- 创建
DROP             -- 删除
ALTER            -- 修改结构
INDEX            -- 索引

-- 查看权限
SHOW GRANTS FOR 'username'@'localhost';
SHOW GRANTS;  -- 当前用户

-- 撤销权限
REVOKE SELECT ON database.* FROM 'username'@'localhost';
REVOKE ALL PRIVILEGES ON database.* FROM 'username'@'localhost';

-- 删除用户
DROP USER 'username'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;
```

---

## 常用系统命令

```sql
-- 查看版本
SELECT VERSION();

-- 查看当前用户
SELECT USER();

-- 查看当前数据库
SELECT DATABASE();

-- 查看所有变量
SHOW VARIABLES;
SHOW VARIABLES LIKE '%char%';

-- 查看状态
SHOW STATUS;
SHOW STATUS LIKE '%thread%';

-- 查看进程
SHOW PROCESSLIST;
SHOW FULL PROCESSLIST;

-- 杀死进程
KILL process_id;

-- 查看引擎
SHOW ENGINES;

-- 查看表状态
SHOW TABLE STATUS;

-- 查看警告
SHOW WARNINGS;

-- 查看错误
SHOW ERRORS;
```

---

## 高级查询技巧

```sql
-- CASE 表达式
SELECT name,
    CASE
        WHEN age < 18 THEN '未成年'
        WHEN age BETWEEN 18 AND 60 THEN '成年'
        ELSE '老年'
    END AS age_group
FROM users;

-- IF 函数
SELECT name, IF(age >= 18, '成年', '未成年') AS status FROM users;

-- IFNULL / COALESCE
SELECT name, IFNULL(phone, '未填写') FROM users;
SELECT name, COALESCE(phone, email, '无联系方式') FROM users;

-- NULLIF
SELECT NULLIF(column1, column2) FROM table_name;

-- 子查询
SELECT * FROM users WHERE id IN (SELECT user_id FROM orders);
SELECT * FROM users WHERE EXISTS (SELECT 1 FROM orders WHERE orders.user_id = users.id);

-- 窗口函数 (MySQL 8.0+)
SELECT name, age,
    ROW_NUMBER() OVER (ORDER BY age) AS row_num,
    RANK() OVER (ORDER BY age) AS rank,
    DENSE_RANK() OVER (ORDER BY age) AS dense_rank
FROM users;

-- 分区窗口
SELECT name, class, score,
    ROW_NUMBER() OVER (PARTITION BY class ORDER BY score DESC) AS class_rank
FROM students;

-- UNION
SELECT name FROM students
UNION
SELECT name FROM teachers;

-- UNION ALL（保留重复）
SELECT name FROM students
UNION ALL
SELECT name FROM teachers;
```

---

## 备份与恢复

```bash
# 备份单个数据库
mysqldump -u root -p database_name > backup.sql

# 备份多个数据库
mysqldump -u root -p --databases db1 db2 db3 > backup.sql

# 备份所有数据库
mysqldump -u root -p --all-databases > all_backup.sql

# 备份指定表
mysqldump -u root -p database_name table1 table2 > tables.sql

# 只备份结构
mysqldump -u root -p --no-data database_name > structure.sql

# 只备份数据
mysqldump -u root -p --no-create-info database_name > data.sql

# 恢复数据库
mysql -u root -p database_name < backup.sql

# 在MySQL中恢复
source /path/to/backup.sql
```

---

## 性能分析

```sql
-- 分析查询执行计划
EXPLAIN SELECT * FROM users WHERE age > 20;
EXPLAIN FORMAT=JSON SELECT * FROM users WHERE age > 20;

-- 查看执行时间
SET profiling = 1;
SELECT * FROM large_table;
SHOW PROFILES;
SHOW PROFILE FOR QUERY 1;

-- 优化表
OPTIMIZE TABLE table_name;

-- 分析表
ANALYZE TABLE table_name;

-- 检查表
CHECK TABLE table_name;

-- 修复表
REPAIR TABLE table_name;
```

# MySQL 性能优化详解

## 索引优化

### 1.1 索引的类型

#### 普通索引 (INDEX)

```sql
CREATE INDEX idx_name ON users(name);
```

#### 唯一索引 (UNIQUE)

```sql
CREATE UNIQUE INDEX idx_email ON users(email);
```

#### 主键索引 (PRIMARY KEY)

```sql
ALTER TABLE users ADD PRIMARY KEY (id);
```

#### 全文索引 (FULLTEXT)

```sql
CREATE FULLTEXT INDEX idx_content ON articles(content);
```

#### 组合索引 (复合索引)

```sql
CREATE INDEX idx_name_age ON users(name, age);
```

### 1.2 索引设计原则

#### ✅ 应该创建索引的情况

1. **WHERE 子句中经常使用的字段**

```sql
-- 查询经常用到 email
SELECT * FROM users WHERE email = 'test@example.com';
-- 应该创建索引
CREATE INDEX idx_email ON users(email);
```

2. **ORDER BY 和 GROUP BY 的字段**

```sql
-- 经常按创建时间排序
SELECT * FROM orders ORDER BY created_at DESC;
-- 应该创建索引
CREATE INDEX idx_created_at ON orders(created_at);
```

3. **JOIN 连接的字段**

```sql
-- 连接查询
SELECT * FROM orders o
INNER JOIN users u ON o.user_id = u.id;
-- 应该在 user_id 上创建索引
CREATE INDEX idx_user_id ON orders(user_id);
```

4. **作为查询条件的外键字段**

```sql
CREATE INDEX idx_order_id ON order_items(order_id);
```

#### ❌ 不应该创建索引的情况

1. **频繁更新的字段**

   - 索引会增加 UPDATE、INSERT、DELETE 的开销

2. **区分度低的字段**

```sql
-- 性别字段只有 M/F 两个值，区分度太低
-- 不推荐创建索引
CREATE INDEX idx_gender ON users(gender); -- ❌
```

3. **表数据量很小**

   - 小于 1000 行的表通常不需要索引

4. **很少查询的字段**
   - 索引会占用存储空间

### 1.3 组合索引的最左前缀原则

```sql
-- 创建组合索引
CREATE INDEX idx_abc ON table_name(a, b, c);

-- ✅ 可以使用索引
WHERE a = 1
WHERE a = 1 AND b = 2
WHERE a = 1 AND b = 2 AND c = 3
WHERE a = 1 AND c = 3  -- 只使用 a

-- ❌ 无法使用索引
WHERE b = 2
WHERE c = 3
WHERE b = 2 AND c = 3
```

**示例：**

```sql
-- 创建组合索引
CREATE INDEX idx_name_age_city ON users(name, age, city);

-- 查询1：使用完整索引
SELECT * FROM users WHERE name = '张三' AND age = 25 AND city = '北京';

-- 查询2：使用部分索引（name, age）
SELECT * FROM users WHERE name = '张三' AND age = 25;

-- 查询3：只使用 name
SELECT * FROM users WHERE name = '张三';

-- 查询4：无法使用索引
SELECT * FROM users WHERE age = 25 AND city = '北京';
```

### 1.4 索引优化技巧

#### 1. 覆盖索引

```sql
-- 创建包含所有查询字段的索引
CREATE INDEX idx_user_info ON users(id, name, email);

-- 查询只需要这三个字段，无需回表
SELECT id, name, email FROM users WHERE name = '张三';
```

#### 2. 前缀索引

```sql
-- 对于长字符串，只索引前面部分
CREATE INDEX idx_url ON pages(url(50));

-- 确定最佳前缀长度
SELECT
    COUNT(DISTINCT LEFT(url, 10)) / COUNT(*) AS sel10,
    COUNT(DISTINCT LEFT(url, 20)) / COUNT(*) AS sel20,
    COUNT(DISTINCT LEFT(url, 30)) / COUNT(*) AS sel30,
    COUNT(DISTINCT LEFT(url, 50)) / COUNT(*) AS sel50
FROM pages;
```

#### 3. 使用索引排序

```sql
-- ✅ 可以使用索引排序
SELECT * FROM users ORDER BY age;

-- ❌ 无法使用索引排序（多个字段不同方向）
SELECT * FROM users ORDER BY age ASC, name DESC;

-- ✅ 创建对应的索引
CREATE INDEX idx_age_name ON users(age ASC, name DESC);
```

#### 4. 避免索引失效

```sql
-- ❌ 在索引列上使用函数
SELECT * FROM users WHERE YEAR(created_at) = 2023;

-- ✅ 改为范围查询
SELECT * FROM users
WHERE created_at >= '2023-01-01' AND created_at < '2024-01-01';

-- ❌ 隐式类型转换
SELECT * FROM users WHERE phone = 13800138000;  -- phone 是 VARCHAR

-- ✅ 使用正确的类型
SELECT * FROM users WHERE phone = '13800138000';

-- ❌ LIKE 以通配符开头
SELECT * FROM users WHERE name LIKE '%张%';

-- ✅ LIKE 不以通配符开头
SELECT * FROM users WHERE name LIKE '张%';

-- ❌ OR 条件中有未建索引的字段
SELECT * FROM users WHERE id = 1 OR nickname = '张三';

-- ✅ 使用 UNION
SELECT * FROM users WHERE id = 1
UNION
SELECT * FROM users WHERE nickname = '张三';

-- ❌ NOT、!=、<> 会导致全表扫描
SELECT * FROM users WHERE status != 0;

-- ✅ 使用 IN
SELECT * FROM users WHERE status IN (1, 2, 3);
```

### 1.5 查看和分析索引

```sql
-- 查看表的索引
SHOW INDEX FROM table_name;

-- 查看索引使用情况
EXPLAIN SELECT * FROM users WHERE email = 'test@example.com';

-- 查看索引统计信息
SELECT * FROM information_schema.statistics
WHERE table_schema = 'database_name'
AND table_name = 'table_name';

-- 分析索引效率
SELECT
    index_name,
    cardinality,
    ROUND(cardinality / table_rows * 100, 2) AS selectivity
FROM information_schema.statistics s
JOIN information_schema.tables t
    ON s.table_schema = t.table_schema
    AND s.table_name = t.table_name
WHERE s.table_schema = 'database_name';
```

---

## 查询优化

### 2.1 使用 EXPLAIN 分析查询

```sql
EXPLAIN SELECT * FROM users WHERE age > 20;
```

**EXPLAIN 输出字段说明：**

| 字段          | 说明                                                            |
| ------------- | --------------------------------------------------------------- |
| id            | 查询序号                                                        |
| select_type   | 查询类型（SIMPLE, PRIMARY, SUBQUERY 等）                        |
| table         | 表名                                                            |
| type          | 连接类型（system > const > eq_ref > ref > range > index > ALL） |
| possible_keys | 可能使用的索引                                                  |
| key           | 实际使用的索引                                                  |
| key_len       | 使用的索引长度                                                  |
| ref           | 与索引比较的列                                                  |
| rows          | 扫描的行数（预估）                                              |
| Extra         | 额外信息                                                        |

**type 字段性能排序：**

```
system > const > eq_ref > ref > range > index > ALL

优 <-----------------------------------------> 差
```

**示例分析：**

```sql
-- 查看执行计划
EXPLAIN SELECT * FROM users WHERE email = 'test@example.com';

-- 理想结果
+----+-------------+-------+-------+---------------+-----------+---------+-------+------+-------+
| id | select_type | table | type  | possible_keys | key       | key_len | ref   | rows | Extra |
+----+-------------+-------+-------+---------------+-----------+---------+-------+------+-------+
|  1 | SIMPLE      | users | const | idx_email     | idx_email | 303     | const |    1 | NULL  |
+----+-------------+-------+-------+---------------+-----------+---------+-------+------+-------+
```

### 2.2 查询优化技巧

#### 1. 避免 SELECT \*

```sql
-- ❌ 不推荐
SELECT * FROM users WHERE id = 1;

-- ✅ 推荐
SELECT id, name, email FROM users WHERE id = 1;
```

#### 2. 使用 LIMIT 限制结果

```sql
-- ✅ 限制返回数量
SELECT * FROM users ORDER BY created_at DESC LIMIT 100;

-- ✅ 分页查询优化
-- 避免深度分页
SELECT * FROM users ORDER BY id LIMIT 10000, 10;  -- 慢

-- 使用 WHERE 条件
SELECT * FROM users WHERE id > 10000 ORDER BY id LIMIT 10;  -- 快
```

#### 3. 优化 JOIN 查询

```sql
-- ❌ 不推荐：大表驱动小表
SELECT * FROM big_table b
INNER JOIN small_table s ON b.id = s.big_id;

-- ✅ 推荐：小表驱动大表
SELECT * FROM small_table s
INNER JOIN big_table b ON s.big_id = b.id;

-- ✅ 确保 JOIN 字段有索引
CREATE INDEX idx_big_id ON small_table(big_id);
CREATE INDEX idx_id ON big_table(id);
```

#### 4. 使用 EXISTS 代替 IN

```sql
-- ❌ 大表使用 IN 性能差
SELECT * FROM users
WHERE id IN (SELECT user_id FROM orders);

-- ✅ 使用 EXISTS
SELECT * FROM users u
WHERE EXISTS (SELECT 1 FROM orders o WHERE o.user_id = u.id);
```

#### 5. 优化子查询

```sql
-- ❌ 相关子查询（慢）
SELECT u.name,
    (SELECT COUNT(*) FROM orders WHERE user_id = u.id) AS order_count
FROM users u;

-- ✅ 使用 JOIN（快）
SELECT u.name, COUNT(o.id) AS order_count
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
GROUP BY u.id, u.name;
```

#### 6. 避免在 WHERE 中使用函数

```sql
-- ❌ 索引失效
SELECT * FROM orders WHERE DATE(created_at) = '2023-11-18';

-- ✅ 使用范围查询
SELECT * FROM orders
WHERE created_at >= '2023-11-18 00:00:00'
AND created_at < '2023-11-19 00:00:00';
```

#### 7. 批量操作优化

```sql
-- ❌ 逐条插入（慢）
INSERT INTO users (name, email) VALUES ('user1', 'email1');
INSERT INTO users (name, email) VALUES ('user2', 'email2');
INSERT INTO users (name, email) VALUES ('user3', 'email3');

-- ✅ 批量插入（快）
INSERT INTO users (name, email) VALUES
    ('user1', 'email1'),
    ('user2', 'email2'),
    ('user3', 'email3');

-- ❌ 逐条更新（慢）
UPDATE users SET status = 1 WHERE id = 1;
UPDATE users SET status = 1 WHERE id = 2;
UPDATE users SET status = 1 WHERE id = 3;

-- ✅ 批量更新（快）
UPDATE users SET status = 1 WHERE id IN (1, 2, 3);
```

### 2.3 查询缓存（MySQL 5.7，8.0 已移除）

```sql
-- 查看查询缓存配置
SHOW VARIABLES LIKE 'query_cache%';

-- 启用查询缓存（MySQL 5.7）
SET GLOBAL query_cache_type = ON;
SET GLOBAL query_cache_size = 67108864;  -- 64MB

-- 查看缓存状态
SHOW STATUS LIKE 'Qcache%';

-- 清除查询缓存
FLUSH QUERY CACHE;
RESET QUERY CACHE;
```

---

## 表结构优化

### 3.1 选择合适的数据类型

#### 1. 数值类型选择

```sql
-- ❌ 不合理
age BIGINT  -- 年龄用 BIGINT 太大

-- ✅ 合理
age TINYINT UNSIGNED  -- 0-255 足够

-- 状态字段
status TINYINT  -- 不要用 INT

-- 金额字段
price DECIMAL(10, 2)  -- 不要用 FLOAT/DOUBLE
```

#### 2. 字符串类型选择

```sql
-- 固定长度用 CHAR
country_code CHAR(2)  -- 'CN', 'US'

-- 可变长度用 VARCHAR
name VARCHAR(50)
email VARCHAR(100)

-- 长文本用 TEXT
article_content TEXT

-- ❌ 不要滥用 VARCHAR(255)
-- ✅ 根据实际需要设置长度
```

#### 3. 日期时间类型

```sql
-- 存储日期时间
created_at DATETIME
updated_at TIMESTAMP  -- 自动更新

-- 只存储日期
birth_date DATE

-- 只存储时间
open_time TIME
```

#### 4. 枚举类型

```sql
-- ✅ 使用 ENUM 代替字符串
status ENUM('pending', 'paid', 'shipped', 'completed')

-- 比这样更高效
status VARCHAR(20)  -- ❌
```

### 3.2 表设计优化

#### 1. 垂直分割（列分割）

```sql
-- ❌ 所有字段放在一起
CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(255),
    profile_text TEXT,           -- 大字段
    avatar_data BLOB,            -- 大字段
    last_login DATETIME
);

-- ✅ 分割成多个表
CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(255),
    last_login DATETIME
);

CREATE TABLE user_profiles (
    user_id INT PRIMARY KEY,
    profile_text TEXT,
    avatar_data BLOB,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

#### 2. 水平分割（行分割）

```sql
-- 按时间分表
CREATE TABLE orders_2023_01 (...);
CREATE TABLE orders_2023_02 (...);
CREATE TABLE orders_2023_03 (...);

-- 按地区分表
CREATE TABLE users_beijing (...);
CREATE TABLE users_shanghai (...);
CREATE TABLE users_guangzhou (...);

-- 使用分区表（MySQL 5.1+）
CREATE TABLE orders (
    id INT,
    order_date DATE,
    ...
) PARTITION BY RANGE (YEAR(order_date)) (
    PARTITION p2021 VALUES LESS THAN (2022),
    PARTITION p2022 VALUES LESS THAN (2023),
    PARTITION p2023 VALUES LESS THAN (2024),
    PARTITION p_future VALUES LESS THAN MAXVALUE
);
```

#### 3. 反范式化设计

```sql
-- 规范化设计
CREATE TABLE orders (
    id INT,
    user_id INT,
    ...
);

-- 每次查询都要 JOIN
SELECT o.*, u.username, u.email
FROM orders o
JOIN users u ON o.user_id = u.id;

-- 反范式化：冗余常用字段
CREATE TABLE orders (
    id INT,
    user_id INT,
    username VARCHAR(50),  -- 冗余
    user_email VARCHAR(100),  -- 冗余
    ...
);

-- 查询更快，无需 JOIN
SELECT * FROM orders WHERE id = 1;
```

### 3.3 表维护

```sql
-- 分析表
ANALYZE TABLE table_name;

-- 优化表
OPTIMIZE TABLE table_name;

-- 检查表
CHECK TABLE table_name;

-- 修复表
REPAIR TABLE table_name;

-- 查看表碎片率
SELECT
    table_name,
    data_free / (data_length + index_length) AS fragmentation_ratio
FROM information_schema.tables
WHERE table_schema = 'database_name'
AND data_free > 0;
```

---

## 配置优化

### 4.1 InnoDB 配置

```ini
[mysqld]
# InnoDB 缓冲池大小（最重要的参数）
# 建议设置为物理内存的 50-80%
innodb_buffer_pool_size = 4G

# InnoDB 缓冲池实例数
# 建议：buffer_pool_size >= 1G 时，设置多个实例
innodb_buffer_pool_instances = 4

# InnoDB 日志文件大小
innodb_log_file_size = 512M

# InnoDB 日志缓冲区大小
innodb_log_buffer_size = 16M

# InnoDB 刷新日志策略
# 0: 每秒刷新（性能最好，但可能丢失1秒数据）
# 1: 每次事务提交刷新（最安全，性能较差）
# 2: 每次事务提交写入，每秒刷新（折中）
innodb_flush_log_at_trx_commit = 2

# InnoDB 刷新方法
innodb_flush_method = O_DIRECT

# InnoDB IO 线程数
innodb_read_io_threads = 8
innodb_write_io_threads = 8

# InnoDB 并发线程数
innodb_thread_concurrency = 0  # 0 表示不限制

# 开启 InnoDB 文件表
innodb_file_per_table = 1
```

### 4.2 连接配置

```ini
[mysqld]
# 最大连接数
max_connections = 500

# 最大用户连接数
max_user_connections = 450

# 连接超时时间
connect_timeout = 10

# 交互式连接超时时间
interactive_timeout = 28800

# 非交互式连接超时时间
wait_timeout = 28800

# 最大允许的包大小
max_allowed_packet = 64M
```

### 4.3 缓存配置

```ini
[mysqld]
# 表缓存
table_open_cache = 4000

# 表定义缓存
table_definition_cache = 2000

# 线程缓存
thread_cache_size = 100

# 排序缓冲区
sort_buffer_size = 2M

# 连接缓冲区
join_buffer_size = 2M

# 读缓冲区
read_buffer_size = 2M

# 随机读缓冲区
read_rnd_buffer_size = 4M
```

### 4.4 慢查询日志

```ini
[mysqld]
# 开启慢查询日志
slow_query_log = 1

# 慢查询日志文件
slow_query_log_file = /var/log/mysql/slow.log

# 慢查询阈值（秒）
long_query_time = 2

# 记录未使用索引的查询
log_queries_not_using_indexes = 1

# 限制每分钟记录的未使用索引查询数
log_throttle_queries_not_using_indexes = 10
```

**查看慢查询：**

```bash
# 使用 mysqldumpslow 分析
mysqldumpslow -s t -t 10 /var/log/mysql/slow.log

# 使用 pt-query-digest 分析
pt-query-digest /var/log/mysql/slow.log
```

### 4.5 查看和修改配置

```sql
-- 查看所有配置
SHOW VARIABLES;

-- 查看特定配置
SHOW VARIABLES LIKE 'innodb%';
SHOW VARIABLES LIKE '%timeout%';

-- 查看全局变量
SHOW GLOBAL VARIABLES LIKE 'max_connections';

-- 查看会话变量
SHOW SESSION VARIABLES LIKE 'autocommit';

-- 临时修改配置（重启失效）
SET GLOBAL max_connections = 500;
SET SESSION sort_buffer_size = 4194304;

-- 永久修改：编辑 my.cnf 或 my.ini
```

---

## 架构优化

### 5.1 读写分离

```
          应用层
             |
        +---------+
        |  Master |  (写)
        +---------+
         /       \
        /         \
   +-------+   +-------+
   | Slave |   | Slave |  (读)
   +-------+   +-------+
```

**实现方式：**

1. 应用层实现（代码中区分读写）
2. 中间件实现（如 MySQL Proxy、ProxySQL）
3. 框架支持（如 Spring 的 AbstractRoutingDataSource）

### 5.2 分库分表

#### 垂直分库

```
-- 按业务模块分库
user_db      -- 用户相关
order_db     -- 订单相关
product_db   -- 商品相关
```

#### 水平分库分表

```
-- 按用户 ID 分库（取模）
user_db_0    -- user_id % 4 = 0
user_db_1    -- user_id % 4 = 1
user_db_2    -- user_id % 4 = 2
user_db_3    -- user_id % 4 = 3

-- 每个库内再分表
users_0, users_1, users_2, users_3
```

**分片策略：**

- 哈希分片：`user_id % N`
- 范围分片：`user_id < 10000 -> db0`
- 地理分片：按地区分
- 时间分片：按日期分

**中间件：**

- Sharding-JDBC
- MyCAT
- Vitess

### 5.3 缓存优化

#### Redis 缓存

```python
# 查询流程
def get_user(user_id):
    # 1. 先查缓存
    user = redis.get(f"user:{user_id}")
    if user:
        return user

    # 2. 缓存未命中，查数据库
    user = db.query("SELECT * FROM users WHERE id = ?", user_id)

    # 3. 写入缓存
    if user:
        redis.setex(f"user:{user_id}", 3600, user)

    return user
```

#### 缓存策略

1. **Cache Aside（旁路缓存）**

   - 读：先读缓存，miss 则读 DB 并更新缓存
   - 写：先写 DB，再删除缓存

2. **Read Through**

   - 应用只和缓存交互
   - 缓存负责从 DB 加载数据

3. **Write Through**

   - 写操作通过缓存，缓存负责写入 DB

4. **Write Behind（异步写）**
   - 写入缓存后立即返回
   - 缓存异步批量写入 DB

### 5.4 连接池优化

```java
// HikariCP 配置示例
HikariConfig config = new HikariConfig();
config.setJdbcUrl("jdbc:mysql://localhost:3306/db");
config.setUsername("user");
config.setPassword("password");

// 连接池大小
config.setMaximumPoolSize(20);      // 最大连接数
config.setMinimumIdle(5);           // 最小空闲连接数

// 超时设置
config.setConnectionTimeout(30000);  // 连接超时 30秒
config.setIdleTimeout(600000);      // 空闲超时 10分钟
config.setMaxLifetime(1800000);     // 最大生命周期 30分钟

// 连接测试
config.setConnectionTestQuery("SELECT 1");
```

---

## 监控与诊断

### 6.1 性能监控指标

```sql
-- 查看连接数
SHOW STATUS LIKE 'Threads_connected';
SHOW STATUS LIKE 'Max_used_connections';

-- 查看 QPS/TPS
SHOW GLOBAL STATUS LIKE 'Questions';
SHOW GLOBAL STATUS LIKE 'Com_commit';
SHOW GLOBAL STATUS LIKE 'Com_rollback';

-- 查看缓冲池命中率
SHOW STATUS LIKE 'Innodb_buffer_pool_read%';

-- 计算命中率
SELECT
    (1 - Innodb_buffer_pool_reads / Innodb_buffer_pool_read_requests) * 100
    AS hit_rate;

-- 查看锁等待
SHOW STATUS LIKE 'Innodb_row_lock%';

-- 查看慢查询数量
SHOW STATUS LIKE 'Slow_queries';

-- 查看正在运行的线程
SHOW PROCESSLIST;
SHOW FULL PROCESSLIST;
```

### 6.2 诊断工具

#### 1. EXPLAIN 分析

```sql
EXPLAIN SELECT * FROM users WHERE age > 20;
EXPLAIN FORMAT=JSON SELECT * FROM users WHERE age > 20;
```

#### 2. SHOW PROFILE

```sql
SET profiling = 1;
SELECT * FROM large_table WHERE condition;
SHOW PROFILES;
SHOW PROFILE FOR QUERY 1;
SHOW PROFILE CPU, BLOCK IO FOR QUERY 1;
```

#### 3. Performance Schema

```sql
-- 启用 Performance Schema
UPDATE performance_schema.setup_instruments
SET ENABLED = 'YES', TIMED = 'YES';

-- 查看最慢的语句
SELECT * FROM performance_schema.events_statements_summary_by_digest
ORDER BY SUM_TIMER_WAIT DESC LIMIT 10;

-- 查看表的 IO 统计
SELECT * FROM performance_schema.table_io_waits_summary_by_table
WHERE OBJECT_SCHEMA = 'your_database'
ORDER BY SUM_TIMER_WAIT DESC;
```

#### 4. 第三方工具

- **pt-query-digest**：分析慢查询日志
- **mytop / innotop**：实时监控
- **Percona Monitoring**：综合监控
- **Prometheus + Grafana**：现代化监控
- **Zabbix**：企业级监控

### 6.3 常用监控查询

```sql
-- 查看数据库大小
SELECT
    table_schema AS 'Database',
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)'
FROM information_schema.tables
GROUP BY table_schema;

-- 查看表大小
SELECT
    table_name AS 'Table',
    ROUND(((data_length + index_length) / 1024 / 1024), 2) AS 'Size (MB)'
FROM information_schema.tables
WHERE table_schema = 'your_database'
ORDER BY (data_length + index_length) DESC;

-- 查看索引使用情况
SELECT
    t.table_schema,
    t.table_name,
    s.index_name,
    s.cardinality
FROM information_schema.tables t
LEFT JOIN information_schema.statistics s
    ON t.table_schema = s.table_schema
    AND t.table_name = s.table_name
WHERE t.table_schema = 'your_database';

-- 查看未使用的索引
SELECT * FROM sys.schema_unused_indexes;

-- 查看冗余索引
SELECT * FROM sys.schema_redundant_indexes;
```

---

## 常见性能问题

### 7.1 慢查询优化案例

#### 案例 1：缺少索引

```sql
-- 问题查询（慢）
SELECT * FROM orders WHERE user_id = 1000;

-- 分析
EXPLAIN SELECT * FROM orders WHERE user_id = 1000;
-- type: ALL（全表扫描）

-- 解决方案
CREATE INDEX idx_user_id ON orders(user_id);

-- 验证
EXPLAIN SELECT * FROM orders WHERE user_id = 1000;
-- type: ref（使用索引）
```

#### 案例 2：索引失效

```sql
-- 问题查询（慢）
SELECT * FROM users WHERE YEAR(created_at) = 2023;

-- 分析
EXPLAIN SELECT * FROM users WHERE YEAR(created_at) = 2023;
-- type: ALL（函数导致索引失效）

-- 解决方案
SELECT * FROM users
WHERE created_at >= '2023-01-01' AND created_at < '2024-01-01';
```

#### 案例 3：深度分页

```sql
-- 问题查询（慢）
SELECT * FROM users ORDER BY id LIMIT 1000000, 10;

-- 解决方案1：使用 WHERE 条件
SELECT * FROM users WHERE id > 1000000 ORDER BY id LIMIT 10;

-- 解决方案2：子查询优化
SELECT * FROM users
WHERE id >= (SELECT id FROM users ORDER BY id LIMIT 1000000, 1)
LIMIT 10;

-- 解决方案3：延迟关联
SELECT u.* FROM users u
INNER JOIN (
    SELECT id FROM users ORDER BY id LIMIT 1000000, 10
) AS t ON u.id = t.id;
```

### 7.2 锁问题

#### 死锁检测

```sql
-- 查看 InnoDB 状态（包括死锁信息）
SHOW ENGINE INNODB STATUS;

-- 查看当前锁等待
SELECT * FROM information_schema.innodb_locks;
SELECT * FROM information_schema.innodb_lock_waits;
SELECT * FROM information_schema.innodb_trx;

-- 查看锁等待详情
SELECT
    r.trx_id waiting_trx_id,
    r.trx_mysql_thread_id waiting_thread,
    r.trx_query waiting_query,
    b.trx_id blocking_trx_id,
    b.trx_mysql_thread_id blocking_thread,
    b.trx_query blocking_query
FROM information_schema.innodb_lock_waits w
INNER JOIN information_schema.innodb_trx b ON b.trx_id = w.blocking_trx_id
INNER JOIN information_schema.innodb_trx r ON r.trx_id = w.requesting_trx_id;
```

#### 避免死锁

1. 按相同顺序访问表和行
2. 尽量使用索引访问数据
3. 缩小事务范围
4. 降低隔离级别（谨慎）
5. 添加合理的超时时间

### 7.3 表锁问题

```sql
-- 查看表锁
SHOW OPEN TABLES WHERE In_use > 0;

-- 查看表元数据锁
SELECT * FROM performance_schema.metadata_locks;

-- 解决方案
-- 1. 尽量使用 InnoDB（行锁）
-- 2. 避免长事务
-- 3. 避免 DDL 操作阻塞业务
```

# Gorm

## 类型

### 一、基础数据类型

#### 1.1 整数类型

```go
type User struct {
    ID        uint           `gorm:"primaryKey"`
    Age       int            // int
    Age8      int8           // TINYINT
    Age16     int16          // SMALLINT
    Age32     int32          // INT
    Age64     int64          // BIGINT

    UAge      uint           // UNSIGNED INT
    UAge8     uint8          // UNSIGNED TINYINT
    UAge16    uint16         // UNSIGNED SMALLINT
    UAge32    uint32         // UNSIGNED INT
    UAge64    uint64         // UNSIGNED BIGINT

    Score     byte           // UNSIGNED TINYINT (uint8 别名)
    Rune      rune           // INT (int32 别名)
}

// 自定义大小
type Product struct {
    ID    uint   `gorm:"primaryKey"`
    Stock int    `gorm:"type:int(10)"`          // 指定长度
    Count uint   `gorm:"type:int unsigned"`     // 无符号
}
```

#### 1.2 浮点数和小数

```go
type Product struct {
    ID          uint    `gorm:"primaryKey"`

    // 浮点数（不精确）
    Price32     float32 `gorm:"type:float"`        // FLOAT
    Price64     float64 `gorm:"type:double"`       // DOUBLE

    // 定点数（精确，用于金额）
    Amount      float64 `gorm:"type:decimal(10,2)"` // DECIMAL(10,2)
    Total       float64 `gorm:"type:numeric(12,4)"` // NUMERIC(12,4)

    // 推荐：用于货币
    Price       float64 `gorm:"type:decimal(10,2);not null;default:0.00"`
}
```

#### 1.3 字符串类型

```go
type Article struct {
    ID       uint   `gorm:"primaryKey"`

    // VARCHAR - 可变长度字符串
    Title    string `gorm:"size:200"`              // VARCHAR(200)
    Author   string `gorm:"type:varchar(100)"`     // VARCHAR(100)

    // TEXT - 长文本
    Content  string `gorm:"type:text"`             // TEXT (65,535 字节)
    Summary  string `gorm:"type:tinytext"`         // TINYTEXT (255 字节)
    Detail   string `gorm:"type:mediumtext"`       // MEDIUMTEXT (16MB)
    Article  string `gorm:"type:longtext"`         // LONGTEXT (4GB)

    // CHAR - 固定长度
    Code     string `gorm:"type:char(10)"`         // CHAR(10)

    // 默认字符串（不指定 size 时）
    Name     string // VARCHAR(255)
}
```

#### 1.4 布尔类型

```go
type User struct {
    ID        uint `gorm:"primaryKey"`

    // bool -> TINYINT(1)
    Active    bool `gorm:"default:true"`
    IsAdmin   bool `gorm:"default:false"`
    Verified  bool

    // 自定义存储方式
    Status    bool `gorm:"type:boolean"`          // BOOLEAN (MySQL 会转为 TINYINT(1))
    Enabled   bool `gorm:"type:tinyint(1)"`       // TINYINT(1)
}

// 布尔值在数据库中的存储
// MySQL: 0 (false), 1 (true)
// PostgreSQL: false, true
// SQLite: 0, 1
```

#### 1.5 时间类型

```go
import "time"

type Event struct {
    ID          uint      `gorm:"primaryKey"`

    // time.Time -> DATETIME
    CreatedAt   time.Time `gorm:"autoCreateTime"`           // 自动设置创建时间
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`           // 自动更新时间
    DeletedAt   gorm.DeletedAt `gorm:"index"`               // 软删除

    // 不同的时间类型
    Birthday    time.Time `gorm:"type:date"`                // DATE (只存储日期)
    StartTime   time.Time `gorm:"type:time"`                // TIME (只存储时间)
    EventTime   time.Time `gorm:"type:datetime"`            // DATETIME

    // 时间戳
    UnixTime    int64     `gorm:"autoCreateTime"`           // Unix 时间戳（秒）
    UnixMilli   int64     `gorm:"autoCreateTime:milli"`     // Unix 时间戳（毫秒）
    UnixNano    int64     `gorm:"autoCreateTime:nano"`      // Unix 时间戳（纳秒）

    // 可空时间
    ExpiredAt   *time.Time                                  // 可为 NULL
    PublishedAt *time.Time `gorm:"default:null"`
}

// 时间精度
type Log struct {
    ID        uint      `gorm:"primaryKey"`
    Timestamp time.Time `gorm:"type:datetime(6)"` // 微秒精度
    CreatedAt time.Time `gorm:"type:datetime(3)"` // 毫秒精度
}
```

#### 1.6 字节和二进制

```go
type File struct {
    ID       uint   `gorm:"primaryKey"`

    // []byte -> BLOB
    Data     []byte `gorm:"type:blob"`            // BLOB (65KB)
    TinyData []byte `gorm:"type:tinyblob"`        // TINYBLOB (255 字节)
    MedData  []byte `gorm:"type:mediumblob"`      // MEDIUMBLOB (16MB)
    LongData []byte `gorm:"type:longblob"`        // LONGBLOB (4GB)

    // VARBINARY
    Hash     []byte `gorm:"type:varbinary(32)"`   // VARBINARY(32)

    // BINARY - 固定长度
    UUID     []byte `gorm:"type:binary(16)"`      // BINARY(16)
}
```

### 二、特殊类型

#### 2.1 枚举类型

```go
// 使用字符串
type Order struct {
    ID     uint   `gorm:"primaryKey"`
    Status string `gorm:"type:enum('pending','paid','shipped','completed');default:'pending'"`
}

// 使用整数（推荐）
type OrderStatus int

const (
    OrderPending OrderStatus = iota
    OrderPaid
    OrderShipped
    OrderCompleted
)

type Order struct {
    ID     uint        `gorm:"primaryKey"`
    Status OrderStatus `gorm:"type:int;default:0"`
}

// 使用自定义类型
type Status string

const (
    StatusActive   Status = "active"
    StatusInactive Status = "inactive"
    StatusDeleted  Status = "deleted"
)

type User struct {
    ID     uint   `gorm:"primaryKey"`
    Status Status `gorm:"type:varchar(20);default:'active'"`
}
```

#### 2.2 JSON 类型

```go
import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

// 使用 datatypes.JSON
import "gorm.io/datatypes"

type User struct {
    ID       uint           `gorm:"primaryKey"`
    Name     string

    // MySQL 5.7+ / PostgreSQL
    Settings datatypes.JSON `gorm:"type:json"`
    Metadata datatypes.JSON `gorm:"type:jsonb"` // PostgreSQL JSONB
}

// 使用示例
user := User{
    Name: "张三",
    Settings: datatypes.JSON([]byte(`{"theme":"dark","language":"zh"}`)),
}
db.Create(&user)

// 查询
var user User
db.First(&user)
var settings map[string]interface{}
json.Unmarshal(user.Settings, &settings)

// 自定义 JSON 类型
type UserInfo struct {
    Age      int      `json:"age"`
    Hobbies  []string `json:"hobbies"`
    Address  string   `json:"address"`
}

// 实现 Scanner 和 Valuer 接口
func (u *UserInfo) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to unmarshal JSON value")
    }
    return json.Unmarshal(bytes, u)
}

func (u UserInfo) Value() (driver.Value, error) {
    return json.Marshal(u)
}

type User struct {
    ID   uint     `gorm:"primaryKey"`
    Name string
    Info UserInfo `gorm:"type:json"`
}

// 使用
user := User{
    Name: "张三",
    Info: UserInfo{
        Age:     25,
        Hobbies: []string{"阅读", "运动"},
        Address: "北京",
    },
}
db.Create(&user)
```

#### 2.3 数组和切片

```go
// 使用 PostgreSQL 数组
import "github.com/lib/pq"

type Article struct {
    ID   uint         `gorm:"primaryKey"`
    Tags pq.StringArray `gorm:"type:text[]"` // PostgreSQL 数组
}

// 使用 JSON 存储切片
type Product struct {
    ID     uint           `gorm:"primaryKey"`
    Images datatypes.JSON `gorm:"type:json"` // 存储 []string
    Tags   datatypes.JSON `gorm:"type:json"` // 存储 []string
}

images := []string{"img1.jpg", "img2.jpg"}
imagesJSON, _ := json.Marshal(images)
product := Product{
    Images: datatypes.JSON(imagesJSON),
}

// 自定义切片类型
type StringSlice []string

func (s *StringSlice) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan StringSlice")
    }
    return json.Unmarshal(bytes, s)
}

func (s StringSlice) Value() (driver.Value, error) {
    if len(s) == 0 {
        return nil, nil
    }
    return json.Marshal(s)
}

type Article struct {
    ID   uint        `gorm:"primaryKey"`
    Tags StringSlice `gorm:"type:json"`
}
```

#### 2.4 Map 类型

```go
// 使用 JSON 存储 Map
type User struct {
    ID       uint           `gorm:"primaryKey"`
    Settings datatypes.JSON `gorm:"type:json"`
}

settings := map[string]interface{}{
    "theme":    "dark",
    "language": "zh",
    "notifications": map[string]bool{
        "email": true,
        "sms":   false,
    },
}
settingsJSON, _ := json.Marshal(settings)
user := User{
    Settings: datatypes.JSON(settingsJSON),
}

// 自定义 Map 类型
type StringMap map[string]string

func (m *StringMap) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan StringMap")
    }
    return json.Unmarshal(bytes, m)
}

func (m StringMap) Value() (driver.Value, error) {
    if len(m) == 0 {
        return nil, nil
    }
    return json.Marshal(m)
}

type Config struct {
    ID      uint      `gorm:"primaryKey"`
    Options StringMap `gorm:"type:json"`
}
```

#### 2.5 UUID 类型

```go
import (
    "github.com/google/uuid"
    "gorm.io/datatypes"
)

type User struct {
    // 使用 UUID 作为主键
    ID   uuid.UUID `gorm:"type:char(36);primaryKey"`
    Name string
}

// 创建时自动生成 UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }
    return nil
}

// 使用 binary 存储（更省空间）
type User struct {
    ID   uuid.UUID `gorm:"type:binary(16);primaryKey"`
    Name string
}

// PostgreSQL UUID 类型
type User struct {
    ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Name string
}
```

#### 2.6 IP 地址类型

```go
import "net"

// 自定义 IP 类型
type IPAddress string

func (ip *IPAddress) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan IPAddress")
    }
    *ip = IPAddress(bytes)
    return nil
}

func (ip IPAddress) Value() (driver.Value, error) {
    return string(ip), nil
}

type Log struct {
    ID        uint      `gorm:"primaryKey"`
    IPAddress IPAddress `gorm:"type:varchar(45)"` // 支持 IPv4 和 IPv6
}

// 使用 net.IP
type Request struct {
    ID        uint   `gorm:"primaryKey"`
    IPAddress string `gorm:"type:varchar(45)"`
}

// 存储和读取
ipStr := "192.168.1.1"
request := Request{IPAddress: ipStr}
db.Create(&request)

var req Request
db.First(&req)
ip := net.ParseIP(req.IPAddress)
```

### 三、GORM 特殊类型

#### 3.1 软删除（DeletedAt）

```go
import "gorm.io/gorm"

type User struct {
    ID        uint           `gorm:"primaryKey"`
    Name      string
    DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除字段
}

// 软删除
db.Delete(&user) // UPDATE users SET deleted_at = NOW() WHERE id = 1;

// 查询自动过滤软删除记录
db.Find(&users) // SELECT * FROM users WHERE deleted_at IS NULL;

// 包含软删除记录
db.Unscoped().Find(&users)

// 永久删除
db.Unscoped().Delete(&user)

// 使用 Unix 时间戳
type User struct {
    ID        uint  `gorm:"primaryKey"`
    Name      string
    DeletedAt int64 `gorm:"index"`
}
```

#### 3.2 数据类型包（datatypes）

```go
import "gorm.io/datatypes"

type User struct {
    ID       uint           `gorm:"primaryKey"`

    // JSON
    Settings datatypes.JSON `gorm:"type:json"`

    // Date
    Birthday datatypes.Date `gorm:"type:date"`

    // Time
    WorkTime datatypes.Time `gorm:"type:time"`

    // URL（PostgreSQL）
    Website  datatypes.URL  `gorm:"type:text"`
}
```

### 四、自定义类型

#### 4.1 实现 Scanner 和 Valuer

```go
import (
    "database/sql/driver"
    "errors"
    "fmt"
    "strings"
)

// 自定义字符串类型
type UpperString string

func (s *UpperString) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan UpperString")
    }
    *s = UpperString(strings.ToUpper(string(bytes)))
    return nil
}

func (s UpperString) Value() (driver.Value, error) {
    return strings.ToUpper(string(s)), nil
}

type User struct {
    ID   uint        `gorm:"primaryKey"`
    Code UpperString `gorm:"type:varchar(50)"`
}

// 使用
user := User{Code: "abc123"}
db.Create(&user) // 存储为 "ABC123"

// 自定义加密类型
type EncryptedString string

func (s *EncryptedString) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan EncryptedString")
    }
    // 解密
    decrypted := decrypt(bytes)
    *s = EncryptedString(decrypted)
    return nil
}

func (s EncryptedString) Value() (driver.Value, error) {
    // 加密
    encrypted := encrypt(string(s))
    return encrypted, nil
}

type User struct {
    ID       uint            `gorm:"primaryKey"`
    Password EncryptedString `gorm:"type:varchar(255)"`
}
```

#### 4.2 自定义复杂类型

```go
// 金额类型（分为单位）
type Money int64

func (m *Money) Scan(value interface{}) error {
    switch v := value.(type) {
    case int64:
        *m = Money(v)
    case []byte:
        val, _ := strconv.ParseInt(string(v), 10, 64)
        *m = Money(val)
    default:
        return errors.New("failed to scan Money")
    }
    return nil
}

func (m Money) Value() (driver.Value, error) {
    return int64(m), nil
}

// 辅助方法
func (m Money) Yuan() float64 {
    return float64(m) / 100.0
}

func NewMoney(yuan float64) Money {
    return Money(yuan * 100)
}

type Order struct {
    ID     uint  `gorm:"primaryKey"`
    Amount Money `gorm:"type:bigint;not null;default:0"`
}

// 使用
order := Order{
    Amount: NewMoney(99.99), // 存储为 9999
}
db.Create(&order)

var order Order
db.First(&order)
fmt.Printf("金额: %.2f 元\n", order.Amount.Yuan()) // 输出: 99.99 元
```

#### 4.3 Point 类型（地理位置）

```go
// 地理坐标
type Point struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

func (p *Point) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan Point")
    }
    return json.Unmarshal(bytes, p)
}

func (p Point) Value() (driver.Value, error) {
    return json.Marshal(p)
}

type Location struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    Position Point  `gorm:"type:json"`
}

// 使用
location := Location{
    Name: "故宫",
    Position: Point{
        Lat: 39.9163,
        Lng: 116.3972,
    },
}
db.Create(&location)

// MySQL 原生 Point 类型
type Location struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    Position string `gorm:"type:point"`
}

// 存储
location := Location{
    Name:     "故宫",
    Position: fmt.Sprintf("POINT(%f %f)", 116.3972, 39.9163),
}
db.Create(&location)
```

### 五、类型转换和序列化

#### 5.1 GormDataType 接口

```go
type MyType struct {
    Value string
}

// 实现 GormDataTypeInterface
func (MyType) GormDataType() string {
    return "varchar(100)"
}

// 实现 GormDBDataTypeInterface（不同数据库使用不同类型）
func (MyType) GormDBDataType(db *gorm.DB, field *schema.Field) string {
    switch db.Dialector.Name() {
    case "mysql", "sqlite":
        return "VARCHAR(100)"
    case "postgres":
        return "TEXT"
    }
    return ""
}

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Data MyType
}
```

#### 5.2 类型序列化示例

```go
// 完整的自定义类型示例
type Tags []string

// 实现数据库类型
func (Tags) GormDataType() string {
    return "json"
}

// 实现扫描
func (t *Tags) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to unmarshal tags")
    }
    return json.Unmarshal(bytes, t)
}

// 实现写入
func (t Tags) Value() (driver.Value, error) {
    if len(t) == 0 {
        return "[]", nil
    }
    return json.Marshal(t)
}

// 辅助方法
func (t Tags) Contains(tag string) bool {
    for _, v := range t {
        if v == tag {
            return true
        }
    }
    return false
}

func (t *Tags) Add(tag string) {
    if !t.Contains(tag) {
        *t = append(*t, tag)
    }
}

type Article struct {
    ID    uint   `gorm:"primaryKey"`
    Title string
    Tags  Tags   `gorm:"type:json"`
}

// 使用
article := Article{
    Title: "Go 语言教程",
    Tags:  Tags{"go", "编程", "教程"},
}
db.Create(&article)

article.Tags.Add("后端")
db.Save(&article)
```

### 六、类型映射表

#### 6.1 Go 类型到 MySQL 类型

| Go 类型                             | MySQL 类型                          | 说明           |
| ----------------------------------- | ----------------------------------- | -------------- |
| bool                                | TINYINT(1)                          | 布尔值         |
| int, int8, int16, int32, int64      | INT, TINYINT, SMALLINT, INT, BIGINT | 有符号整数     |
| uint, uint8, uint16, uint32, uint64 | UNSIGNED INT                        | 无符号整数     |
| float32, float64                    | FLOAT, DOUBLE                       | 浮点数         |
| string                              | VARCHAR(255)                        | 字符串（默认） |
| string with size                    | VARCHAR(size)                       | 指定长度       |
| string with type:text               | TEXT                                | 长文本         |
| []byte                              | BLOB                                | 二进制数据     |
| time.Time                           | DATETIME                            | 日期时间       |
| datatypes.JSON                      | JSON                                | JSON 数据      |
| datatypes.Date                      | DATE                                | 日期           |
| datatypes.Time                      | TIME                                | 时间           |

#### 6.2 标签指定类型

```go
type Example struct {
    // 基础类型指定
    Field1 string `gorm:"type:varchar(100)"`
    Field2 int    `gorm:"type:int unsigned"`
    Field3 bool   `gorm:"type:boolean"`

    // 精度指定
    Price  float64 `gorm:"type:decimal(10,2)"`
    Amount float64 `gorm:"precision:12;scale:4"`

    // 长度指定
    Name   string `gorm:"size:100"`
    Code   string `gorm:"size:50"`

    // 组合使用
    Status string `gorm:"type:varchar(20);not null;default:'active'"`
}
```

### 七、最佳实践

#### 7.1 类型选择建议

```go
// ✅ 推荐
type Product struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"size:200;not null"`              // 明确指定大小
    Price       float64        `gorm:"type:decimal(10,2);not null"`    // 金额用 decimal
    Stock       int            `gorm:"not null;default:0"`             // 库存用 int
    Status      int            `gorm:"type:tinyint;default:1"`         // 状态用小整数
    Description string         `gorm:"type:text"`                      // 长文本用 text
    CreatedAt   time.Time      `gorm:"autoCreateTime"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// ❌ 不推荐
type Product struct {
    ID          int       // 应该用 uint
    Name        string    // 没指定大小
    Price       float64   // 金额应该用 decimal
    Description string    // 长文本应该用 text 类型
}
```

#### 7.2 性能优化

```go
// 选择合适的类型大小
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Age      uint8  `gorm:"type:tinyint unsigned"`  // 年龄 0-255 够用
    Status   uint8  `gorm:"type:tinyint unsigned"`  // 状态值不多用 tinyint
    ViewCount uint32 `gorm:"type:int unsigned"`      // 浏览量可能很大
}

// 索引优化
type Article struct {
    ID        uint   `gorm:"primaryKey"`
    Title     string `gorm:"size:200;index"`          // 经常查询的字段加索引
    Category  string `gorm:"size:50;index"`
    Tags      string `gorm:"type:text"`               // 长文本不加索引
}
```

---

## 数据库操作

```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

### 1.1 什么是 \*gorm.DB

```go
// *gorm.DB 是 GORM 的核心结构
type DB struct {
    *Config                    // 配置
    Error        error         // 错误
    RowsAffected int64         // 影响行数
    Statement    *Statement    // SQL 语句构建器
    clone        int           // 克隆级别
}

// 每次操作都会返回新的 *gorm.DB 实例（链式调用）
db1 := db.Where("age > ?", 18)  // 新实例
db2 := db1.Order("name")        // 又一个新实例
// db, db1, db2 互不影响
```

### 1.2 链式方法 vs 终止方法

```go
// 链式方法（Chain Methods）
// - 返回 *gorm.DB
// - 可以继续链式调用
// - 不执行 SQL
// - 用于构建查询条件

db.Where("age > ?", 18).     // 链式方法
   Order("name").            // 链式方法
   Limit(10)                 // 链式方法
   // 此时还没有执行 SQL

// 终止方法（Finisher Methods）
// - 返回 *gorm.DB（但通常检查 Error）
// - 执行 SQL
// - 返回结果
// - 链式调用的终点

db.Where("age > ?", 18).
   Order("name").
   Limit(10).
   Find(&users)  // 终止方法，执行 SQL
```

## 链式方法详解

### 2.1 条件方法

#### Where - 添加查询条件

```go
// 字符串条件
db.Where("name = ?", "张三").Find(&users)
db.Where("name = ? AND age >= ?", "张三", 20).Find(&users)

// 结构体条件（只查询非零值）
db.Where(&User{Name: "张三", Age: 20}).Find(&users)
// 注意：Age = 0 不会被查询

// Map 条件（可以查询零值）
db.Where(map[string]interface{}{
    "name": "张三",
    "age":  0,  // 会查询 age = 0
}).Find(&users)

// IN 查询
db.Where("name IN ?", []string{"张三", "李四"}).Find(&users)

// LIKE 查询
db.Where("name LIKE ?", "%张%").Find(&users)

// BETWEEN
db.Where("age BETWEEN ? AND ?", 20, 30).Find(&users)

// 复杂条件
db.Where("(name = ? OR email = ?) AND age > ?", "张三", "zhangsan@example.com", 18).Find(&users)

// 子查询
subQuery := db.Model(&Order{}).Select("user_id").Where("amount > ?", 1000)
db.Where("id IN (?)", subQuery).Find(&users)

// 返回值：*gorm.DB（链式）
result := db.Where("age > ?", 18)
// 此时没有执行 SQL
result.Find(&users)  // 这里才执行
```

#### Or - 或条件

```go
// 基础 OR
db.Where("name = ?", "张三").
   Or("name = ?", "李四").
   Find(&users)
// SELECT * FROM users WHERE name = '张三' OR name = '李四';

// 结构体 OR
db.Where("age > ?", 30).
   Or(User{Name: "张三"}).
   Find(&users)

// Map OR
db.Where("status = ?", 1).
   Or(map[string]interface{}{"age": 18}).
   Find(&users)

// 复杂 OR
db.Where(db.Where("name = ?", "张三").Or("email = ?", "zhangsan@example.com")).
   Where("status = ?", 1).
   Find(&users)
// SELECT * FROM users WHERE (name = '张三' OR email = '...') AND status = 1;

// 返回值：*gorm.DB（链式）
```

#### Not - 非条件

```go
// 基础 NOT
db.Not("name = ?", "张三").Find(&users)
// SELECT * FROM users WHERE NOT name = '张三';

// NOT IN
db.Not(map[string]interface{}{
    "name": []string{"张三", "李四"},
}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ('张三', '李四');

// 结构体 NOT
db.Not(User{Name: "张三"}).Find(&users)

// 主键 NOT IN
db.Not([]int64{1, 2, 3}).Find(&users)
// SELECT * FROM users WHERE id NOT IN (1,2,3);

// 返回值：*gorm.DB（链式）
```

### 2.2 查询构建方法

#### Select - 选择字段

```go
// 选择特定字段
db.Select("name", "age").Find(&users)
db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Select("*").Find(&users) // 选择所有字段
db.Select("*").Omit("Role").Find(&users) // 选择除 Role 外的所有字段

// 排除字段（使用 Omit）
db.Omit("password", "secret").Find(&users)
// SELECT ... (除了 password 和 secret 的所有字段)

// 使用表达式
db.Select("name, age * 2 AS double_age").Find(&users)
db.Select("COALESCE(name, 'Unknown') AS name").Find(&users)

// 子查询
db.Select("name, (SELECT COUNT(*) FROM orders WHERE user_id = users.id) AS order_count").
   Find(&users)

// 聚合函数
db.Select("COUNT(*) AS count, AVG(age) AS avg_age").
   Table("users").
   Scan(&result)

// 返回值：*gorm.DB（链式）
```

#### Omit - 忽略字段

```go
// 忽略单个字段
db.Omit("password").Find(&users)

// 忽略多个字段
db.Omit("password", "secret", "internal_notes").Find(&users)

// 创建时忽略字段
db.Omit("CreatedAt").Create(&user)

// 更新时忽略字段
db.Omit("UpdatedAt").Updates(&user)

// 忽略关联
db.Omit("Profile").Create(&user)  // 不创建 Profile

import "gorm.io/gorm/clause"
// 忽略所有关联
db.Omit(clause.Associations).Create(&user)

// 返回值：*gorm.DB（链式）
```

#### Order - 排序

```go
// 单字段排序
db.Order("age").Find(&users)        // 升序
db.Order("age DESC").Find(&users)   // 降序

// 多字段排序
db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age DESC, name ASC;

// 使用字段表达式
db.Order("FIELD(status, 'pending', 'processing', 'completed')").Find(&orders)

// 动态排序
orderBy := "age DESC"
db.Order(orderBy).Find(&users)

// 多次调用 Order（会累加）
db.Order("age DESC").Order("name").Find(&users)
// ORDER BY age DESC, name

// 使用 Reorder 覆盖之前的排序
db.Order("age DESC").Reorder("name").Find(&users)
// ORDER BY name（覆盖了 age DESC）

// 清除排序
db.Order("age DESC").Reorder("").Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Limit - 限制数量

```go
// 基础限制
db.Limit(10).Find(&users)
// SELECT * FROM users LIMIT 10;

// 取消限制（使用 -1）
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// users2 查询时没有限制

// 分页配合 Offset
db.Limit(10).Offset(20).Find(&users)
// SELECT * FROM users LIMIT 10 OFFSET 20;

// 动态限制
limit := 20
if limit > 0 {
    db = db.Limit(limit)
}

// 返回值：*gorm.DB（链式）
```

#### Offset - 偏移量

```go
// 基础偏移
db.Offset(10).Find(&users)
// SELECT * FROM users OFFSET 10;

// 分页使用
page := 2
pageSize := 10
db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)

// 取消偏移（使用 -1）
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)

// 分页函数
func Paginate(page, pageSize int) func(*gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if page <= 0 {
            page = 1
        }
        offset := (page - 1) * pageSize

        return db.Offset(offset).Limit(pageSize)
    }
}

// 使用
db.Scopes(Paginate(2, 10)).Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Group - 分组

```go
// 基础分组
db.Select("name, COUNT(*) AS count").
   Group("name").
   Find(&results)
// SELECT name, COUNT(*) AS count FROM users GROUP BY name;

// 多字段分组
db.Select("status, age, COUNT(*) AS count").
   Group("status, age").
   Scan(&results)

// 配合 Having
db.Select("age, COUNT(*) AS count").
   Group("age").
   Having("COUNT(*) > ?", 5).
   Scan(&results)
// SELECT age, COUNT(*) AS count FROM users
// GROUP BY age HAVING COUNT(*) > 5;

// 复杂分组
db.Select("DATE(created_at) AS date, COUNT(*) AS count").
   Group("DATE(created_at)").
   Order("date DESC").
   Scan(&results)

// 返回值：*gorm.DB（链式）
```

#### Having - Having 条件

```go
// 基础 Having
db.Select("name, COUNT(*) AS count").
   Group("name").
   Having("COUNT(*) > ?", 1).
   Scan(&results)

// 多个 Having 条件
db.Select("age, AVG(score) AS avg_score, COUNT(*) AS count").
   Group("age").
   Having("AVG(score) > ?", 60).
   Having("COUNT(*) >= ?", 5).
   Scan(&results)

// 复杂 Having
db.Select("category, SUM(amount) AS total").
   Group("category").
   Having("SUM(amount) > ? AND COUNT(*) > ?", 1000, 5).
   Scan(&results)

// 返回值：*gorm.DB（链式）
```

#### Distinct - 去重

```go
// 去重单字段
db.Distinct("name").Find(&users)
// SELECT DISTINCT name FROM users;

// 去重多字段
db.Distinct("name", "email").Find(&users)
// SELECT DISTINCT name, email FROM users;

// 配合其他条件
db.Distinct("name").
   Where("age > ?", 18).
   Order("name").
   Find(&users)

// 返回值：*gorm.DB（链式）
```

### 2.3 关联方法

#### Joins - 连接查询 / Profile Profile 一对一

```go
// 基础 Joins（LEFT JOIN）
db.Joins("Profile").Find(&users)
// SELECT users.*, profiles.* FROM users
// LEFT JOIN profiles ON profiles.user_id = users.id;

// INNER JOIN
db.InnerJoins("Profile").Find(&users)

// 带条件的 Joins
db.Joins("Profile", db.Where(&Profile{Age: 18})).Find(&users)
// LEFT JOIN profiles ON ... AND profiles.age = 18

// 自定义 JOIN
db.Joins("LEFT JOIN profiles ON profiles.user_id = users.id").Find(&users)

// 多个 Joins
db.Joins("Profile").Joins("Company").Find(&users)

// JOIN 子查询
subQuery := db.Table("orders").
    Select("user_id, COUNT(*) AS count").
    Group("user_id")
db.Joins("LEFT JOIN (?) AS order_stats ON order_stats.user_id = users.id", subQuery).
   Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Preload - 预加载 Articles []Article 一对多

```go
// 基础预加载
db.Preload("Articles").Find(&users)

// 多个预加载
db.Preload("Articles").Preload("Profile").Find(&users)

// 嵌套预加载
db.Preload("Articles.Comments").Find(&users)

// 条件预加载
db.Preload("Articles", "status = ?", 1).Find(&users)

db.Preload("Articles", func(db *gorm.DB) *gorm.DB {
    return db.Where("status = ?", 1).Order("created_at DESC").Limit(5)
}).Find(&users)

// 预加载所有
import "gorm.io/gorm/clause"
db.Preload(clause.Associations).Find(&users)

// 返回值：*gorm.DB（链式）
```

### 2.4 其他链式方法

#### Model - 指定模型

```go
// 指定模型
db.Model(&User{}).Where("age > ?", 18).Find(&users)

// 用于更新
db.Model(&user).Update("name", "新名字")

// 批量操作
db.Model(&User{}).Where("status = ?", 0).Update("status", 1)

// 返回值：*gorm.DB（链式）
```

#### Table - 指定表名

```go
// 指定表名
db.Table("users").Where("age > ?", 18).Find(&users)

// 动态表名
tableName := "users_2024"
db.Table(tableName).Find(&users)

// 表别名
db.Table("users u").
   Joins("LEFT JOIN orders o ON o.user_id = u.id").
   Select("u.*, COUNT(o.id) AS order_count").
   Group("u.id").
   Scan(&results)

// 返回值：*gorm.DB（链式）
```

#### Scopes - 作用域

```go
// 定义作用域
func ActiveUsers(db *gorm.DB) *gorm.DB {
    return db.Where("status = ?", 1)
}

func AgeGreaterThan(age int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("age > ?", age)
    }
}

// 使用作用域
db.Scopes(ActiveUsers, AgeGreaterThan(18)).Find(&users)

// 多个作用域组合
db.Scopes(
    ActiveUsers,
    AgeGreaterThan(18),
    func(db *gorm.DB) *gorm.DB {
        return db.Order("created_at DESC").Limit(10)
    },
).Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Clauses - 添加子句

```go
import "gorm.io/gorm/clause"

// OnConflict（Upsert）
db.Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "email"}},
    DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
}).Create(&users)

// Locking（锁）
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Session - 会话配置

```go
// 创建新会话
db.Session(&gorm.Session{
    DryRun:      true,  // 只生成 SQL 不执行
    PrepareStmt: true,  // 预编译
    SkipHooks:   true,  // 跳过钩子
}).Find(&users)

// 返回值：*gorm.DB（链式）
```

#### WithContext - 上下文

```go
import "context"

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 带超时的查询
db.WithContext(ctx).Find(&users)

// 返回值：*gorm.DB（链式）
```

#### Debug - 调试模式

```go
// 临时开启 SQL 日志
db.Debug().Where("age > ?", 18).Find(&users)
// 会打印执行的 SQL

// 返回值：*gorm.DB（链式）
```

## 终止方法详解

### 3.1 查询终止方法

#### Find - 查询多条记录

```go
// 查询所有
var users []User
result := db.Find(&users)
// 检查错误
if result.Error != nil {
    // 处理错误
}
// 检查行数
fmt.Println(result.RowsAffected)

// 带条件查询
db.Where("age > ?", 18).Find(&users)

// 根据主键查询
db.Find(&users, []int{1, 2, 3})
// SELECT * FROM users WHERE id IN (1,2,3);

// 空结果不报错
db.Find(&users)  // users 为空切片，Error 为 nil

// 返回值：*gorm.DB
// - Error: 错误信息
// - RowsAffected: 影响行数
```

#### First - 查询第一条

```go
// 查询第一条（按主键排序）
var user User
result := db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    // 记录不存在
}

// 根据主键查询
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = '10';

// 带条件查询
db.Where("name = ?", "张三").First(&user)

// 返回值：*gorm.DB
// - Error: 错误（找不到时返回 ErrRecordNotFound）
// - RowsAffected: 0 或 1
```

#### Last - 查询最后一条

```go
// 查询最后一条（按主键倒序）
var user User
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 带条件
db.Where("age > ?", 18).Last(&user)

// 返回值：*gorm.DB
// - Error: 错误（找不到时返回 ErrRecordNotFound）
```

#### Take - 查询一条（无排序）

```go
// 查询一条（不指定顺序）
var user User
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 根据主键
db.Take(&user, 10)

// 带条件
db.Where("name = ?", "张三").Take(&user)

// 返回值：*gorm.DB
// - Error: 错误（找不到时返回 ErrRecordNotFound）
```

#### Count - 统计数量

```go
// 统计数量
var count int64
db.Model(&User{}).Count(&count)
// SELECT COUNT(*) FROM users;

// 带条件统计
db.Model(&User{}).Where("age > ?", 18).Count(&count)

// 统计 Distinct
db.Model(&User{}).Distinct("name").Count(&count)
// SELECT COUNT(DISTINCT name) FROM users;

// 返回值：*gorm.DB
// - Error: 错误
// count 变量被填充
```

#### Row - 获取单行

```go
// 获取 *sql.Row
row := db.Table("users").
    Where("name = ?", "张三").
    Select("name, age, email").
    Row()

// 扫描结果
row.Scan(&results)

// 返回值：*sql.Row
```

#### Rows - 获取多行

```go
// 获取 *sql.Rows
rows, err := db.Model(&User{}).
    Where("age > ?", 18).
    Select("name, age").
    Rows()
defer rows.Close()

if err != nil {
    // 处理错误
}

// 遍历结果
for rows.Next() {
    rows.Scan(&results)
    fmt.Println(name, age)
}

// 返回值：*sql.Rows, error
```

#### Scan - 扫描到结构体

```go
type Result struct {
    Name  string
    Count int64
}

// 扫描到自定义结构体
var results []Result
db.Table("users").
    Select("name, COUNT(*) AS count").
    Group("name").
    Scan(&results)

// 扫描单条
var result Result
db.Table("users").
    Select("name, age").
    Where("id = ?", 1).
    Scan(&result)

// 扫描到 Map
var result map[string]interface{}
db.Table("users").Take(&result)

// 返回值：*gorm.DB
// - Error: 错误
```

#### Pluck - 查询单列

```go
// 查询单列到切片
var names []string
db.Model(&User{}).Pluck("name", &names)
// SELECT name FROM users;

var ages []int
db.Model(&User{}).Pluck("age", &ages)

// 带条件
db.Model(&User{}).Where("age > ?", 18).Pluck("email", &emails)

// 返回值：*gorm.DB
// - Error: 错误
// names/ages 变量被填充
```

### 3.2 创建终止方法

#### Create - 创建记录 BeforeSave, BeforeCreate, AfterSave, AfterCreate

```go
// 创建单条
user := User{Name: "张三", Age: 25}
result := db.Create(&user)

// 检查结果
if result.Error != nil {
    // 处理错误
}
fmt.Println(user.ID)           // 自动填充的 ID
fmt.Println(result.RowsAffected) // 影响行数

// 批量创建
users := []User{
    {Name: "张三"},
    {Name: "李四"},
}
db.Create(&users)

// 指定字段创建
db.Select("Name", "Age").Create(&user)

// 忽略字段创建
db.Omit("CreatedAt").Create(&user)

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
// user.ID 会被填充
```

#### CreateInBatches - 分批创建

```go
// 分批创建（避免单次插入过多）
users := []User{...}  // 1000 条记录
db.CreateInBatches(users, 100)  // 每批 100 条

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 总影响行数
```

#### Save - 保存（创建或更新）

```go
// 保存记录
// 如果主键为空，执行 INSERT
// 如果主键有值，执行 UPDATE（更新所有字段）

user := User{Name: "张三", Age: 25}
db.Save(&user)  // INSERT

user.Age = 30
db.Save(&user)  // UPDATE

// 注意：Save 会更新所有字段（包括零值）
user.Age = 0
db.Save(&user)  // age 会被更新为 0

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
```

### 3.3 更新终止方法

#### Update - 更新单个字段

```go
// 更新单个字段
db.Model(&user).Update("name", "新名字")
// UPDATE users SET name='新名字', updated_at='...' WHERE id=1;

// 批量更新
db.Model(&User{}).Where("age < ?", 18).Update("active", false)

// 使用表达式
db.Model(&User{}).Where("id = ?", 1).Update("age", gorm.Expr("age + ?", 1))

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
```

#### Updates - 更新多个字段 BeforeSave, BeforeUpdate, AfterSave, AfterUpdate

```go
// 使用 Struct 更新（零值字段不更新） -> Select（可以更新零值）
db.Model(&user).Updates(User{Name: "新名字", Age: 30})

// 使用 Map 更新（可以更新零值）
db.Model(&user).Updates(map[string]interface{}{
    "name": "新名字",
    "age":  0,  // 会更新
})

// 批量更新
db.Model(&User{}).
    Where("status = ?", 0).
    Updates(map[string]interface{}{"status": 1})

// 选择字段更新
db.Model(&user).Select("Name", "Age").Updates(User{Name: "新名字", Age: 0})

// 忽略字段更新
db.Model(&user).Omit("Email").Updates(User{Name: "新名字", Email: "new@example.com"})

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
```

#### UpdateColumn/UpdateColumns - 更新（不触发钩子）

```go
// 更新单列（不触发钩子，不更新 updated_at）
db.Model(&user).UpdateColumn("name", "新名字")

// 更新多列
db.Model(&user).UpdateColumns(User{Name: "新名字", Age: 30})

// 返回值：*gorm.DB
```

### 3.4 删除终止方法

#### Delete - 删除记录 BeforeDelete、AfterDelete

```go
// 删除记录
db.Delete(&user)
// DELETE FROM users WHERE id = 1;

// 根据主键删除
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, []int{1, 2, 3})
// DELETE FROM users WHERE id IN (1,2,3);

// 条件删除
db.Where("age < ?", 18).Delete(&User{})

// 软删除（模型有 DeletedAt 字段）
db.Delete(&user)
// UPDATE users SET deleted_at='...' WHERE id = 1;

// 永久删除
db.Unscoped().Delete(&user)
// DELETE FROM users WHERE id = 1;

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
```

### 3.5 其他终止方法

#### Exec - 执行原生 SQL

```go
// 执行原生 SQL
db.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "张三")

// 执行多条语句
db.Exec("DELETE FROM users WHERE age < 18; OPTIMIZE TABLE users;")

// 返回值：*gorm.DB
// - Error: 错误
// - RowsAffected: 影响行数
```

#### Raw - 原生查询

```go
// 原生查询
var users []User
db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&users)

// 查询单行
var user User
db.Raw("SELECT * FROM users WHERE name = ? LIMIT 1", "张三").Scan(&user)

// 返回值：*gorm.DB
```

#### FirstOrInit - 查询或初始化

```go
// 查询，不存在则初始化（不保存）
var user User
db.Where(User{Name: "张三"}).FirstOrInit(&user)
// 如果找到，user 包含查询结果
// 如果没找到，user 被初始化（Name = "张三"）但不保存

// 指定初始化属性
db.Where(User{Name: "张三"}).
    Attrs(User{Age: 20}).
    FirstOrInit(&user)
// 没找到时，user = {Name: "张三", Age: 20}

// 返回值：*gorm.DB
```

#### FirstOrCreate - 查询或创建

```go
// 查询，不存在则创建
var user User
db.Where(User{Name: "张三"}).FirstOrCreate(&user)
// 如果找到，user 包含查询结果
// 如果没找到，创建新记录

// 指定创建属性
db.Where(User{Name: "张三"}).
    Attrs(User{Age: 20}).
    FirstOrCreate(&user)
// 创建时设置 Age = 20

// 无论是否找到都更新
db.Where(User{Name: "张三"}).
    Assign(User{Age: 20}).
    FirstOrCreate(&user)
// 找到或创建后，Age 都会被设置为 20

// 返回值：*gorm.DB
```

#### Association - 关联操作

```go
// 获取关联操作对象
var user User
db.First(&user, 1)

// 查找关联
var articles []Article
db.Model(&user).Association("Articles").Find(&articles)

// 追加关联
article := Article{Title: "新文章"}
db.Model(&user).Association("Articles").Append(&article)
db.Model(&user).Association("Articles").Append(&article1, &article2)

// 替换关联
db.Model(&user).Association("Articles").Replace(&article)

// 删除关联
db.Model(&user).Association("Articles").Delete(&article)

// 清空关联
db.Model(&user).Association("Articles").Clear()

// 统计关联数量
count := db.Model(&user).Association("Articles").Count()

// 返回值：关联操作对象
```

#### Transaction - 手动事务

```go
// 开始事务
tx := db.Begin()

// 执行操作
if err := tx.Create(&user1).Error; err != nil {
    tx.Rollback()
    return err
}

if err := tx.Create(&user2).Error; err != nil {
    tx.Rollback()
    return err
}

// 提交事务
if err := tx.Commit().Error; err != nil {
    return err
}

// 返回值：*gorm.DB（事务对象）

// db.Transaction 自动处理提交和回滚
err := db.Transaction(func(tx *gorm.DB) error {
    // 在事务中执行操作
    if err := tx.Create(&user).Error; err != nil {
        return err  // 返回错误会自动回滚
    }

    if err := tx.Create(&profile).Error; err != nil {
        return err  // 自动回滚
    }

    return nil  // 返回 nil 会自动提交
})

if err != nil {
    // 事务失败，已经回滚
    log.Println("Transaction failed:", err)
}

// 事务成功，已经提交
```

#### FindInBatches - 批量查询

```go
// 分批处理大量数据
result := db.Where("age > ?", 18).FindInBatches(&users, 100, func(tx *gorm.DB, batch int) error {
    // 处理每批数据
    for _, user := range users {
        // 处理 user
    }

    // 返回 error 会停止后续批次
    return nil
})

// 检查结果
if result.Error != nil {
    // 处理错误
}
fmt.Println(result.RowsAffected)  // 总行数

// 返回值：*gorm.DB
```

## 五、实战模式

### 5.1 查询模式

```go
// 模式1: 简单查询
db.Where("age > ?", 18).Find(&users)

// 模式2: 复杂条件查询
db.Where("age > ?", 18).
   Where("status = ?", 1).
   Or("vip = ?", true).
   Order("created_at DESC").
   Limit(10).
   Find(&users)

// 模式3: 动态查询构建
query := db.Model(&User{})
if name != "" {
    query = query.Where("name LIKE ?", "%"+name+"%")
}
if ageMin > 0 {
    query = query.Where("age >= ?", ageMin)
}
query.Find(&users)

// 模式4: 预加载关联
db.Preload("Profile").
   Preload("Articles", func(db *gorm.DB) *gorm.DB {
       return db.Order("created_at DESC").Limit(5)
   }).
   Find(&users)

// 模式5: JOIN查询
db.Joins("Profile").
   Where("profiles.age > ?", 18).
   Find(&users)

// 模式6: 统计查询
var result struct {
    Status string
    Count  int64
    AvgAge float64
}
db.Model(&User{}).
   Select("status, COUNT(*) AS count, AVG(age) AS avg_age").
   Group("status").
   Having("COUNT(*) > ?", 10).
   Scan(&results)
```

### 5.2 创建模式

```go
// 模式1: 简单创建
user := User{Name: "张三", Age: 25}
db.Create(&user)

// 模式2: 批量创建
users := []User{{Name: "张三"}, {Name: "李四"}}
db.Create(&users)

// 模式3: 分批创建
db.CreateInBatches(users, 100)

// 模式4: 选择性创建
db.Select("Name", "Age").Create(&user)
db.Omit("CreatedAt").Create(&user)

// 模式5: 创建关联
user := User{
    Name: "张三",
    Profile: Profile{Bio: "简介"},
    Articles: []Article{{Title: "文章1"}},
}
db.Create(&user)

// 模式6: Upsert
db.Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "email"}},
    DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
}).Create(&users)
```

### 5.3 更新模式

```go
// 模式1: 更新单字段
db.Model(&user).Update("name", "新名字")

// 模式2: 更新多字段（Map）
db.Model(&user).Updates(map[string]interface{}{
    "name": "新名字",
    "age":  0,  // Map可以更新零值
})

// 模式3: 批量更新
db.Model(&User{}).
   Where("status = ?", 0).
   Updates(map[string]interface{}{"status": 1})

// 模式4: 使用表达式
db.Model(&User{}).
   Where("id = ?", 1).
   Update("age", gorm.Expr("age + ?", 1))

// 模式5: 选择性更新
db.Model(&user).
   Select("Name", "Age").
   Updates(User{Name: "新名字", Age: 0})

// 模式6: 事务更新
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Model(&user).Update("balance", newBalance).Error; err != nil {
        return err
    }

    if err := tx.Create(&transaction).Error; err != nil {
        return err
    }

    return nil
})
```

### 5.4 删除模式

```go
// 模式1: 删除单条
db.Delete(&user)

// 模式2: 根据主键删除
db.Delete(&User{}, 1)
db.Delete(&User{}, []int{1, 2, 3})

// 模式3: 条件删除
db.Where("age < ?", 18).Delete(&User{})

// 模式4: 软删除
db.Delete(&user)  // 设置 deleted_at

// 模式5: 永久删除
db.Unscoped().Delete(&user)

// 模式6: 关联删除
db.Select("Articles").Delete(&user)  // 同时删除关联
```

### 5.5 完整示例

```go
// 用户服务完整示例
type UserService struct {
    db *gorm.DB
}

// 查询用户列表（最佳实践）
func (s *UserService) ListUsers(req ListUsersRequest) (*ListUsersResponse, error) {
    var users []User
    var total int64

    // 构建查询
    query := s.db.Model(&User{})

    // 动态条件
    if req.Name != "" {
        query = query.Where("name LIKE ?", "%"+req.Name+"%")
    }
    if req.MinAge > 0 {
        query = query.Where("age >= ?", req.MinAge)
    }
    if len(req.Statuses) > 0 {
        query = query.Where("status IN ?", req.Statuses)
    }

    // 统计总数
    if err := query.Count(&total).Error; err != nil {
        return nil, err
    }

    // 查询数据（限制字段、分页、排序）
    err := query.
        Select("id", "name", "email", "age", "status", "created_at").
        Scopes(Paginate(req.Page, req.PageSize)).
        Order("created_at DESC").
        Find(&users).Error

    if err != nil {
        return nil, err
    }

    return &ListUsersResponse{
        Users: users,
        Total: total,
        Page:  req.Page,
        PageSize: req.PageSize,
    }, nil
}

// 获取用户详情（含关联）
func (s *UserService) GetUserDetail(id uint) (*User, error) {
    var user User

    err := s.db.
        Preload("Profile").
        Preload("Articles", func(db *gorm.DB) *gorm.DB {
            return db.Where("status = ?", 1).
                Order("created_at DESC").
                Limit(10)
        }).
        Preload("Roles").
        First(&user, id).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("用户不存在")
    }

    return &user, err
}

// 更新用户（事务）
func (s *UserService) UpdateUser(id uint, req UpdateUserRequest) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        // 检查用户是否存在
        var user User
        if err := tx.First(&user, id).Error; err != nil {
            return err
        }

        // 更新用户
        updates := make(map[string]interface{})
        if req.Name != nil {
            updates["name"] = *req.Name
        }
        if req.Age != nil {
            updates["age"] = *req.Age
        }

        if len(updates) > 0 {
            if err := tx.Model(&user).Updates(updates).Error; err != nil {
                return err
            }
        }

        // 更新资料
        if req.Bio != nil {
            if err := tx.Model(&Profile{}).
                Where("user_id = ?", id).
                Update("bio", *req.Bio).Error; err != nil {
                return err
            }
        }

        return nil
    })
}
```

# Thrift IDL

## 基本类型（Base Types）

```thrift
// 布尔型
bool isActive = true

// 8位有符号整数
byte age = 25

// 16位有符号整数
i16 port = 8080

// 32位有符号整数
i32 count = 100

// 64位有符号整数
i64 timestamp = 1234567890

// 双精度浮点数
double price = 99.99

// 字符串（UTF-8）
string name = "Alice"

// 二进制数据
binary data
```

## 容器类型（Container Types）

### 1. List（列表）

```thrift
// 有序可重复元素列表
list<i32> numbers
list<string> names
list<User> users

// 示例定义
struct Response {
    1: list<string> messages
    2: list<i32> errorCodes
}
```

### 2. Set（集合）

```thrift
// 无序不重复元素集合
set<string> tags
set<i32> userIds

// 示例
struct Article {
    1: string title
    2: set<string> tags
    3: set<i64> likedBy
}
```

### 3. Map（映射）

```thrift
// 键值对映射
map<string, string> metadata
map<i32, User> userMap
map<string, list<string>> categories

// 示例
struct Config {
    1: map<string, string> settings
    2: map<i32, double> prices
}
```

## 结构体（Struct）

```thrift
// 基本结构体
struct User {
    1: required i64 id           // 必填字段
    2: required string name      // 必填
    3: optional string email     // 可选字段
    4: string phone              // 默认 optional
    5: i32 age = 0               // 带默认值
}

// 嵌套结构体
struct Address {
    1: string street
    2: string city
    3: string country
}

struct UserProfile {
    1: User user
    2: Address address
    3: list<string> hobbies
}

// 结构体继承（不推荐使用，已废弃）
// struct Employee extends User {
//     1: string employeeId
// }
```

### 字段修饰符

```thrift
struct Example {
    1: required string id        // 必须存在
    2: optional string name      // 可以不存在
    3: string description        // 默认 optional
}
```

## 枚举（Enum）

```thrift
// 基本枚举
enum Status {
    ACTIVE = 1
    INACTIVE = 2
    PENDING = 3
    DELETED = 4
}

// 不指定值（自动从0开始）
enum Color {
    RED
    GREEN
    BLUE
}

// 使用枚举
struct User {
    1: string name
    2: Status status
}
```

## 异常（Exception）

```thrift
// 定义异常（类似 struct）
exception UserNotFoundException {
    1: string message
    2: i32 code
}

exception ValidationException {
    1: string message
    2: map<string, string> errors
}

// 在服务中使用
service UserService {
    User getUser(1: i64 id) throws (1: UserNotFoundException notFound)
}
```

## 联合体（Union）

```thrift
// 只能有一个字段被设置
union Result {
    1: string success
    2: string error
}

union SearchResult {
    1: User user
    2: Article article
    3: Comment comment
}

// 使用示例
struct Response {
    1: Result result
}
```

## 常量（Const）

```thrift
// 基本类型常量
const i32 MAX_USERS = 1000
const string DEFAULT_NAME = "Guest"
const double PI = 3.14159

// 容器常量
const list<string> ALLOWED_ROLES = ["admin", "user", "guest"]
const map<string, i32> ERROR_CODES = {
    "NOT_FOUND": 404,
    "FORBIDDEN": 403,
    "SERVER_ERROR": 500
}

// 枚举常量
enum Status {
    ACTIVE = 1
    INACTIVE = 2
}
const Status DEFAULT_STATUS = Status.ACTIVE
```

## 类型定义（Typedef）

```thrift
// 类型别名
typedef i64 UserId
typedef string Email
typedef map<string, string> Metadata

// 使用别名
struct User {
    1: UserId id
    2: Email email
    3: Metadata metadata
}
```

## 服务定义（Service）

```thrift
// 基本服务
service UserService {
    // 返回 User，无参数
    User getCurrentUser()

    // 带参数
    User getUser(1: i64 id)

    // 多个参数
    list<User> searchUsers(1: string keyword, 2: i32 limit, 3: i32 offset)

    // void 返回类型
    void deleteUser(1: i64 id)

    // 抛出异常
    User updateUser(1: User user) throws (
        1: UserNotFoundException notFound,
        2: ValidationException validation
    )

    // oneway：不等待响应
    oneway void logEvent(1: string event)
}

// 服务继承
service AdminService extends UserService {
    list<User> getAllUsers()
    void banUser(1: i64 id)
}
```

## 命名空间（Namespace）

```thrift
// 为不同语言指定命名空间/包名
namespace java com.example.thrift
namespace go example.thrift
namespace py example.thrift
namespace cpp example.thrift
namespace php example.thrift
namespace js example.thrift
```

## Include（包含其他文件）

```thrift
// user.thrift
namespace java com.example.user

struct User {
    1: i64 id
    2: string name
}

// post.thrift
include "user.thrift"

namespace java com.example.post

struct Post {
    1: i64 id
    2: string title
    3: user.User author  // 使用 include 的类型
}
```

## 注释

```thrift
/**
 * 多行注释
 * 用户结构体定义
 */
struct User {
    1: i64 id        // 单行注释
    2: string name   /* 行内注释 */
    // 3: string deprecated  // 注释掉的字段
}

// 单行注释

# Shell 风格注释也支持
```

## 完整示例：电商系统

```thrift
namespace java com.example.shop
namespace go shop
namespace py shop

// ============= 常量 =============
const i32 MAX_PAGE_SIZE = 100
const string DEFAULT_CURRENCY = "USD"

// ============= 枚举 =============
enum OrderStatus {
    PENDING = 1
    PAID = 2
    SHIPPED = 3
    DELIVERED = 4
    CANCELLED = 5
}

enum PaymentMethod {
    CREDIT_CARD = 1
    PAYPAL = 2
    ALIPAY = 3
    WECHAT = 4
}

// ============= 类型定义 =============
typedef i64 Timestamp
typedef string Currency

// ============= 结构体 =============
struct Money {
    1: double amount
    2: Currency currency = DEFAULT_CURRENCY
}

struct Address {
    1: string street
    2: string city
    3: string state
    4: string zipCode
    5: string country
}

struct User {
    1: required i64 id
    2: required string email
    3: required string name
    4: optional string phone
    5: optional Address address
    6: Timestamp createdAt
}

struct Product {
    1: required i64 id
    2: required string name
    3: required string description
    4: required Money price
    5: i32 stock
    6: set<string> tags
    7: map<string, string> attributes
}

struct OrderItem {
    1: Product product
    2: i32 quantity
    3: Money subtotal
}

struct Order {
    1: required i64 id
    2: required User user
    3: required list<OrderItem> items
    4: required Money total
    5: required OrderStatus status
    6: required PaymentMethod paymentMethod
    7: optional Address shippingAddress
    8: Timestamp createdAt
    9: optional Timestamp paidAt
}

// ============= 异常 =============
exception ProductNotFoundException {
    1: string message
    2: i64 productId
}

exception InsufficientStockException {
    1: string message
    2: i64 productId
    3: i32 requestedQuantity
    4: i32 availableStock
}

exception OrderNotFoundException {
    1: string message
    2: i64 orderId
}

exception PaymentFailedException {
    1: string message
    2: string reason
}

// ============= Union =============
union SearchResult {
    1: list<Product> products
    2: list<User> users
    3: list<Order> orders
}

// ============= 服务 =============
service ProductService {
    /**
     * 获取产品详情
     */
    Product getProduct(1: i64 productId) throws (
        1: ProductNotFoundException notFound
    )

    /**
     * 搜索产品
     */
    list<Product> searchProducts(
        1: string keyword,
        2: i32 page = 1,
        3: i32 pageSize = 20
    )

    /**
     * 创建产品
     */
    Product createProduct(1: Product product)

    /**
     * 更新库存
     */
    void updateStock(1: i64 productId, 2: i32 quantity) throws (
        1: ProductNotFoundException notFound
    )
}

service OrderService {
    /**
     * 创建订单
     */
    Order createOrder(
        1: i64 userId,
        2: list<OrderItem> items,
        3: PaymentMethod paymentMethod
    ) throws (
        1: ProductNotFoundException notFound,
        2: InsufficientStockException insufficientStock
    )

    /**
     * 获取订单
     */
    Order getOrder(1: i64 orderId) throws (
        1: OrderNotFoundException notFound
    )

    /**
     * 获取用户订单列表
     */
    list<Order> getUserOrders(
        1: i64 userId,
        2: optional OrderStatus status,
        3: i32 page = 1,
        4: i32 pageSize = 20
    )

    /**
     * 取消订单
     */
    void cancelOrder(1: i64 orderId) throws (
        1: OrderNotFoundException notFound
    )

    /**
     * 处理支付
     */
    bool processPayment(1: i64 orderId) throws (
        1: OrderNotFoundException notFound,
        2: PaymentFailedException paymentFailed
    )

    /**
     * 异步记录订单事件（不等待响应）
     */
    oneway void logOrderEvent(1: i64 orderId, 2: string event)
}

// 管理员服务（继承）
service AdminService extends ProductService {
    list<Order> getAllOrders(1: i32 page, 2: i32 pageSize)
    void deleteProduct(1: i64 productId)
    map<string, i32> getStatistics()
}
```

## Thrift 数据类型对照表

| Thrift 类型 | Java    | Go             | Python | C++         |
| ----------- | ------- | -------------- | ------ | ----------- |
| bool        | boolean | bool           | bool   | bool        |
| byte        | byte    | int8           | int    | int8_t      |
| i16         | short   | int16          | int    | int16_t     |
| i32         | int     | int32          | int    | int32_t     |
| i64         | long    | int64          | long   | int64_t     |
| double      | double  | float64        | float  | double      |
| string      | String  | string         | str    | std::string |
| binary      | byte[]  | []byte         | bytes  | std::string |
| list        | List    | []T            | list   | std::vector |
| set         | Set     | map[T]struct{} | set    | std::set    |
| map         | Map     | map[K]V        | dict   | std::map    |
