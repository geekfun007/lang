## MongoDB 基本命令行详解

### 1. 连接和启动

**启动 MongoDB 服务**

```bash
# 启动MongoDB服务
mongod

# 指定配置文件启动
mongod --config /path/to/mongod.conf

# 指定数据目录启动
mongod --dbpath /path/to/data/directory

# 指定端口启动
mongod --port 27017
```

**连接到 MongoDB**

```bash
# 连接到本地MongoDB
mongo

# 连接到指定主机和端口
mongo --host localhost --port 27017

# 连接到远程MongoDB
mongo mongodb://username:password@host:port/database

# 连接到MongoDB Atlas
mongo "mongodb+srv://cluster.mongodb.net/database" --username username
```

### 2. 数据库操作

**选择和创建数据库**

```javascript
// 选择数据库（如果不存在会自动创建）
use myDatabase

// 查看当前数据库
db

// 查看所有数据库
show dbs

// 删除当前数据库
db.dropDatabase()
```

### 3. 集合（Collection）操作

**集合管理**

```javascript
// 创建集合
db.createCollection("users", {
  validator: {
    $jsonSchema: {
      bsonType: "bsonType",
    }
  },
  validationLevel: "strict",    // 可选
  validationAction: "error"     // 可选
})

db.runCommand({
  collMod: "users",
  validator: {},
  validationLevel: "strict",
  validationAction: "error"
})

// 查看所有集合
show collections

// 删除集合
db.users.drop()

// 查看集合统计信息
db.users.stats()
```

**bson type**

- BSON vs JSON：
  - BSON 是二进制格式，JSON 是文本格式
  - BSON 支持更多数据类型（如 ObjectId、Date、Binary 等）
  - BSON 保持字段顺序，JSON 不保证字段顺序
  - BSON 支持更高效的序列化和反序列化

```javascript
db.createCollection("users", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["_id", "name", "age", "email", "createdAt"],
      properties: {
        _id: {
          bsonType: "objectId",
        },
        age: {
          bsonType: "int",
          minimum: 0,
          maximum: 150,
        },
        bigId: {
          bsonType: "long",
        },
        score: {
          bsonType: "double",
          minimum: 0,
          maximum: 100,
        },
        preciseBalance: {
          bsonType: "decimal",
        },
        name: {
          bsonType: "string",
          minLength: 1,
          maxLength: 100,
        },
        email: {
          bsonType: "string",
          pattern: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
        },
        status: {
          bsonType: "string",
          enum: ["start", "end"],
        },
        isActive: {
          bsonType: "bool",
        },
        createdAt: {
          bsonType: "date",
        },
        updatedAt: {
          bsonType: "timestamp",
        },
        hobbies: {
          bsonType: "array",
          items: {
            bsonType: "string",
          },
          minItems: 1,
          maxItems: 10,
        },
        address: {
          bsonType: "object",
          required: ["city"],
          properties: {
            city: { bsonType: "string" },
            district: { bsonType: "string" },
            zipCode: {
              bsonType: "string",
              pattern: "^\\d{6}$",
            },
          },
        },
      },
    },
  },
});
```

### 4. 文档（Document）操作

**插入文档**

```javascript
// 插入单个文档
db.users.insertOne({
  name: "张三",
  age: 25,
  email: "zhangsan@example.com",
});

// 插入多个文档
db.users.insertMany([
  { name: "李四", age: 30, email: "lisi@example.com" },
  { name: "王五", age: 28, email: "wangwu@example.com" },
]);

// 使用save方法（如果_id存在则更新，否则插入）
db.users.save({ name: "赵六", age: 32 });
```

**查询文档**

> 文档级别作用域：使用 $ 引用当前文档的字段
> 表达式级别作用域：使用 $$ 引用在表达式中定义的变量 $$NOW

```javascript
// 查询所有文档
db.users.find();

// 格式化输出
db.users.find().pretty();
db.users.find().toArray();

// 查询单个文档
db.users.findOne();

// 运算符
db.users.find({ age: { $gt: 25 } }); // 大于
db.users.find({ age: { $gt: new Date("2023-01-01 21:10:21") } }); // 大于
db.users.find({ age: { $gte: 25 } }); // 大于等于
db.users.find({ age: { $lt: 30 } }); // 小于
db.users.find({ age: { $lte: 30 } }); // 小于等于
db.users.find({ age: { $ne: 25 } }); // 不等于
db.users.find({ age: { $in: [25, 30, 35] } }); // 在指定值中
db.users.find({ age: { $nin: [25, 30, 35] } }); // 不在指定值中
db.users.find({ email: { $exists: 1 } }); // 存在字段
db.users.find({ email: { $exists: 0 } }); // 不存在字段

// 文本匹配
db.users.find({ name: /张/i }); // 名字包含"张"
db.users.find({ name: /^张/ }); // 名字以"张"开头
db.users.find({ name: /张$/ }); // 名字以"张"结尾
db.users.find({ name: { $regex: "^张", $options: "i" } }); // 名字以"张"结尾

// 数组
db.users.find({ hobbies: "游泳" }); // 包含特定值
db.users.find({ "hobbies.0": "游泳" }); // 第一个元素是"游泳"
db.users.find({ hobbies: { $all: ["游泳", "篮球"] } }); // 数组包含所有指定元素
db.users.find({ hobbies: { $in: ["游泳", "篮球"] } }); // 数组包含所有任意元素
db.users.find({ hobbies: { $size: 3 } }); // 数组长度

// 嵌套
db.users.find({ "address.city": "北京", "address.district": "朝阳区" }); // 字段查询
db.users.find({ address: { city: "北京", district: "朝阳区" } }); // 完全匹配

// 逻辑操作符
db.users.find({ $and: [{ age: { $gt: 20 } }, { age: { $lt: 40 } }] });
db.users.find({ $or: [{ age: 25 }, { name: "张三" }] });

// 字段比较
db.users.find({
  $expr: {
    $gte: ["$age", "$index"],
  },
});

// 字段投影
db.users.find({}, { name: 1, age: 1, _id: 0 }); // 只返回name和age字段，不返回_id字段
db.users.find(
  {},
  {
    _id: 0,
    name: 1,
    // 默认值
    ifNull: {
      $ifNull: ["$age", 0],
    },
    // 比较
    gt: { $gt: ["$age", 6] },
    gte: { $gte: ["$age", 6] },
    lt: { $lt: ["$age", 6] },
    lte: { $lte: ["$age", 6] },
    eq: { $eq: ["$age", 6] },
    ne: { $ne: ["$age", 6] },
    // 计算
    add: { $add: ["$age", "$index"] },
    subtract: { $subtract: ["$age", "$index"] },
    multiply: { $multiply: ["$age", "$index"] },
    divide: { $divide: ["$age", "$index"] },
    round: { $round: ["$age", 2] },
    // 字符串
    concat: { $concat: ["$name"] },
    toUpper: { $toUpper: "$name" },
    toLower: { $toLower: "$name" },
    substr: { $substr: ["$name", 0, 2] },
    // 日期
    year: { $year: "$create_at" },
    month: { $month: "$create_at" },
    dayOfMonth: { $dayOfMonth: "$create_at" },
    hour: { $hour: "$create_at" },
    minute: { $minute: "$create_at" },
    second: { $second: "$create_at" },
    dateToString: {
      $dateToString: { date: "$create_at", format: "%Y-%m-%d %H:%M:%S" },
    },
    // 数组
    slice: {
      $slice: ["$hobbies", 1, 2],
    },
    arrayElemAt: {
      $arrayElemAt: ["$hobbies", 0],
    },
    size: {
      $size: "$hobbies",
    },
    hobbies: {
      $map: {
        input: "$hobbies",
        as: "hobby",
        in: {
          $toUpper: "$$hobby",
        },
      },
    },
    // 条件
    cond: {
      $cond: {
        if: { $gte: ["$age", 18] },
        then: "adult",
        else: "minor",
      },
    },
    switch: {
      $switch: {
        branches: [
          { case: { $gte: ["$score", 90] }, then: "A" },
          { case: { $gte: ["$score", 80] }, then: "B" },
          { case: { $gte: ["$score", 70] }, then: "C" },
          { case: { $gte: ["$score", 60] }, then: "D" },
          { case: { $gte: ["$score", 0] }, then: "F" },
        ],
        default: "Invalid",
      },
    },
  }
);

// 排序
db.users.find().sort({ age: 1 }); // 按年龄升序
db.users.find().sort({ age: -1 }); // 按年龄降序

// 限制数量
db.users.find().limit(5); // 限制返回5条
db.users.find().skip(10); // 跳过前10条
db.users.find().skip(10).limit(5); // 分页查询

// 计数
db.users.count();
db.users.count({ age: { $gt: 25 } });
```

**更新文档**

```javascript
// 更新单个字段
db.users.updateOne({ name: "张三" }, { $set: { age: 26 } });

// 更新多个字段
db.users.updateOne({ name: "张三" }, { $set: { age: 26, status: "active" } });

// 使用upsert（如果不存在则插入）
db.users.updateOne(
  { name: "新用户" },
  { $set: { age: 25, email: "new@example.com" } },
  { upsert: true }
);

// $set - 设置字段值
db.users.updateOne(
  { name: "张三" },
  { $set: { age: 26, lastLogin: new Date() } }
);

// $unset - 删除字段
db.users.updateOne({ name: "张三" }, { $unset: { oldField: "" } });

// $inc - 数值增减
db.users.updateOne({ name: "张三" }, { $inc: { age: 1, score: 10 } });

// $rename - 重命名字段
db.users.updateOne({ name: "张三" }, { $rename: { oldName: "newName" } });

// 数值
db.users.updateOne({ name: "张三" }, { $push: { hobbies: "跑步" } }); // 添加元素到数组末尾 addToSet 添加唯一元素
db.users.updateOne(
  { name: "张三" },
  { $push: { hobbies: { $each: ["跑步", "健身", "瑜伽"] } } }
); // 添加多个元素到数组末尾
db.users.updateOne({ name: "张三" }, { $pull: { hobbies: "游泳" } }); // 删除匹配的元素
db.users.updateOne(
  { name: "张三" },
  {
    $pull: {
      hobbies: {
        $in: ["跑步", "健身", "瑜伽"],
      },
    },
  }
); // 删除多个匹配的元素
```

**删除文档**

```javascript
// 删除单个文档
db.users.deleteOne({ name: "张三" });

// 删除多个文档
db.users.deleteMany({ age: { $lt: 25 } });

// 删除所有文档
db.users.deleteMany({});
```

### 5. 索引操作

**创建索引**

```javascript
db.users.find().explain("executionStats");

// 创建单字段索引
db.users.createIndex({ name: 1 });

// 创建复合索引
db.users.createIndex({ name: 1, age: -1 });

// 创建唯一索引
db.users.createIndex({ email: 1 }, { unique: true });

// 创建文本索引
db.users.createIndex({ name: "text", description: "text" });

// 查看索引
db.users.getIndexes();

// 删除索引
db.users.dropIndex({ name: 1 });
db.users.dropIndex("index_name");
```

### 6. 聚合操作

**聚合管道**

```javascript
// 过滤
db.users.aggregate([{ $match: { age: { $gt: 25 } } }]);

// 字段投影
db.users.aggregate([
  { $project: { name: 1, age: 1, _id: 0, gte: { $gte: ["$age", 15] } } },
]);

// 排序
db.users.aggregate([{ $sort: { age: -1 } }]);

// 分页
db.users.aggregate([{ $limit: 10, $skip: 20 }]);

// 分组
db.users.aggregate([
  // 单字段分组
  {
    $group: {
      _id: "$status",
      sum: { $sum: "$age" },
      min: { $min: "$age" },
      max: { $max: "$age" },
      avg: { $avg: "$age" },
      hobbies: { $push: "$hobby" }, //  $addToSet
      firstStatus: { $first: "status" },
      lastStatus: { $last: "status" },
    },
  },
  // 多字段分组
  {
    $group: {
      _id: {
        department: "$department",
        status: "$status",
      },
      sum: { $sum: 1 },
    },
  },
  // 所有文档分组
  {
    $group: {
      _id: null,
      sum: { $sum: 1 },
    },
  },
]);

// 展开数组
db.users.aggregate([
  { $unwind: "$hobbies" },
  {
    $unwind: {
      path: "$hobbies",
      preserveNullAndEmptyArrays: true,
    },
  },
]);

// 关联
db.orders.aggregate([
  {
    $lookup: {
      from: "users",
      foreignField: "_id",
      localField: "userId",
      as: "user",
    },
  },
  {
    $lookup: {
      from: "users",
      let: {
        userId: "$_id",
        startDate: { $subtract: ["$$NOW", 30 * 24 * 60 * 60 * 1000] },
      },
      pipeline: [
        {
          $match: {
            // 聚合表达式
            $expr: {
              $and: [
                {
                  $eq: ["$userId", "$$userId"],
                },
                {
                  $gte: ["$orderDate", "$$startDate"],
                },
              ],
            },
          },
        },
      ],
      as: "user",
    },
  },
]);

// 分桶
db.users.aggregate([
  {
    $bucket: {
      groupBy: "$amount",
      boundaries: [0, 100, 200, 300, 400],
      default: "400+",
      output: {
        sum: { $sum: "$age" },
      },
    },
  },
]);
```

### 7. 用户和权限管理

**用户管理**

```javascript
// 创建用户
db.createUser({
  user: "admin",
  pwd: "password",
  roles: ["readWrite", "dbAdmin"],
});

// 查看用户
db.getUsers();

// 删除用户
db.dropUser("admin");

// 修改用户密码
db.changeUserPassword("admin", "newpassword");
```

### 8. 备份和恢复

**备份操作**

```bash
# 备份整个数据库
mongodump --db myDatabase --out /backup/path

# 备份特定集合
mongodump --db myDatabase --collection users --out /backup/path

# 备份到压缩文件
mongodump --db myDatabase --gzip --archive=backup.gz
```

**恢复操作**

```bash
# 恢复数据库
mongorestore --db myDatabase /backup/path/myDatabase

# 从压缩文件恢复
mongorestore --gzip --archive=backup.gz

# 恢复特定集合
mongorestore --db myDatabase --collection users /backup/path/myDatabase/users.bson
```

### 9. 监控和性能

**监控命令**

```javascript
// 查看服务器状态
db.serverStatus();

// 查看数据库统计
db.stats();

// 查看集合统计
db.users.stats();

// 查看当前操作
db.currentOp();

// 查看慢查询
db.setProfilingLevel(2, { slowms: 100 });
db.system.profile.find().sort({ ts: -1 }).limit(5);
```

### 10. 常用工具命令

```bash
# 导出为JSON
mongoexport --db myDatabase --collection users --out users.json

# 从JSON导入
mongoimport --db myDatabase --collection users --file users.json

# 导出为CSV
mongoexport --db myDatabase --collection users --type csv --fields name,age,email --out users.csv

# 从CSV导入
mongoimport --db myDatabase --collection users --type csv --fields name,age,email --file users.csv
```
