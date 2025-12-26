# DataType

```bash
#! /bin/bash

set -e
```

## 数据类型

### 1. 字符串 (String)

Bash 中所有变量默认都是字符串类型。

```bash
# 字符串定义
str1="Hello World"
str2='Single quotes'
str3=NoQuotesSimple

# 字符串拼接
name="John"
greeting="Hello, $name!"  # 结果: Hello, John!

# 字符串长度
echo ${#str1}  # 输出: 11

# 字符串切片
echo ${str1:0:5}   # 输出: Hello (从位置0开始，取5个字符)
echo ${str1:6}     # 输出: World (从位置6到结尾)

# 大小写转换
text="Hello World"
echo ${text^^}           # HELLO WORLD (全部大写)
echo ${text,,}           # hello world (全部小写)
echo ${text^}            # Hello World (首字母大写)

# 模式匹配规则
# 通配符规则
# *       # 匹配任意数量的字符（包括0个）
# ?       # 匹配单个字符
# [...]   # 匹配字符类中的任意一个字符
# [!...]  # 匹配不在字符类中的任意一个字符
# 字符类规则
# [[:alnum:]]  # 字母和数字
# [[:alpha:]]  # 字母
# [[:digit:]]  # 数字
# [[:lower:]]  # 小写字母
# [[:upper:]]  # 大写字母
# [[:space:]]  # 空白字符
# [[:punct:]]  # 标点符号

# 字符串替换
# 替换第一个匹配 ${variable/pattern/replacement}
# 替换所有匹配 ${variable//pattern/replacement}
# 替换开头匹配 ${variable/#pattern/replacement}
# 替换结尾匹配 ${variable/%pattern/replacement}
echo ${str1/World/Bash}     # 输出: Hello Bash (替换第一个匹配)
echo ${str1//l/L}          # 输出: HeLLo WorLd (替换所有匹配)
echo ${str1/#Hello/Hi}     # 输出: Hi Bash (替换开头匹配)
echo ${str1/%World/Bash}          # 输出: HeLLo WorLd (替换结尾匹配)
echo ${str1// /} # 输出: Hello WorLd (删除字符（替换为空）)

# 字符串删除
# # 表示从开头删除 ## 表示从开头删除最长匹配
# % 表示从结尾删除 %% 表示从结尾删除最长匹配
text="filename.tar.gz"
echo ${text#*.}      # 输出: tar.gz (删除最短前缀)
echo ${text##*.}     # 输出: gz (删除最长前缀)
echo ${text%.*}      # 输出: filename.tar (删除最短后缀)
echo ${text%%.*}     # 输出: filename (删除最长后缀)

# 精确匹配 == "$pattern"
filename="document.pdf"
if [[ $filename == "document.pdf" ]]; then
    echo "PDF file"
fi

# 包含匹配 == *"$pattern"*
filename="document.pdf"
if [[ $filename == *"document"* ]]; then
    echo "Document file"
fi

# 模式匹配 == $pattern
filename="document.pdf"
if [[ $filename == *.pdf ]]; then
    echo "PDF file"
fi

# 正则表达式匹配 =~ $pattern
if [[ $filename =~ ^[a-z]+\.pdf$ ]]; then
    echo "Matches pattern"
fi

# 正则匹配规则
test_regex() {
    local text="$1"
    local pattern="$2"

    echo "文本: '$text'"
    echo "模式: '$pattern'"

    if [[ "$text" =~ $pattern ]]; then
        echo "✅ 匹配成功"
        echo "匹配组: ${BASH_REMATCH[@]}"
    else
        echo "❌ 匹配失败"
    fi
    echo "---"
}

# 基本字符类
# 数字字符类
test_regex "123" "[0-9]+"
# 字母字符类
test_regex "abc" "[a-zA-Z]+"
# 字母数字字符类
test_regex "abc123" "[a-zA-Z0-9]+"
# 小写字母
test_regex "hello" "[a-z]+"
# 大写字母
test_regex "WORLD" "[A-Z]+"
# 预定义字符类
# 数字
test_regex "123" "[[:digit:]]+"
# 字母
test_regex "abc" "[[:alpha:]]+"
# 字母数字
test_regex "abc123" "[[:alnum:]]+"
# 小写字母
test_regex "hello" "[[:lower:]]+"
# 大写字母
test_regex "WORLD" "[[:upper:]]+"
# 空白字符
test_regex "hello world" "[[:space:]]+"
# 标点符号
test_regex "hello!" "[[:punct:]]+"

# 量词
# 零个或多个 (*)
test_regex "a" "a*"
# 一个或多个 (+)
test_regex "a" "a+"
# 零个或一个 (?)
test_regex "a" "a?"
# 精确数量 {n}
test_regex "aaa" "a{3}"
# 范围 {n,m}
test_regex "aa" "a{2,4}"
# 至少n个 {n,}
test_regex "aaa" "a{2,}"

# 锚点
# 行首 (^)
test_regex "hello world" "^hello"
# 行尾 ($)
test_regex "hello world" "world$"
# 完整行匹配
test_regex "hello" "^hello$"
# 单词边界 (\b) 不支持 (^|[[:space:]])hello([[:space:]]|$)
test_regex "hello world" "\bhello\b"

# 前瞻和后顾
# 正向前瞻 (?=...)
test_regex "hello123" "hello(?=[0-9])"
# 负向前瞻 (?!...)
test_regex "helloabc" "hello(?!123)"
# 正向后顾 (?<=...)
test_regex "123hello" "(?<=[0-9])hello"
# 负向后顾 (?<!...)
test_regex "abchello" "(?<![0-9])hello"

# 条件匹配
# 或操作 (|)
test_regex "hello123" "(hello|world)[0-9]+"

# 分组
# 基本分组
test_regex "hello world" "(hello) (world)"
# 非捕获分组 不支持 ${BASH_REMATCH[2]}
test_regex "hello world" "(?:hello) (world)"
# 命名分组 不支持 declare -A parts=([greeting]="${BASH_REMATCH[1]}")
if [[ "hello world" =~ (?<greeting>hello) ]]; then
    echo "命名分组匹配成功"
    echo "greeting: ${BASH_REMATCH[1]}"
fi


# 日期
date
date +'%Y-%m-%d %H:%M:%S'
date -v +1H -v +10M +'%Y-%m-%d %H:%M:%S'
date +'%s'
date -r 123123131 +'%Y-%m-%d %H:%M:%S'

cal
cal -3
cal 2025
cal 10 2025
```

### 2. 数字 (Numbers)

Bash 原生不支持浮点数，只支持整数运算。

```bash
# 整数运算
a=10
b=5

# 算术运算 - 使用 $(()) 或 $[]
result=$((a + b))      # 15
result=$((a - b))      # 5
result=$((a * b))      # 50
result=$((a / b))      # 2
result=$((a % b))      # 0

# 自增自减
((a++))                # a = 11
((a--))                # a = 10
((a += 5))             # a = 15

# 比较运算
if (( a > b )); then
    echo "a is greater"
fi

# 进制转换
echo $((16#FF))          # 十六进制FF转十进制: 255
echo $((8#77))           # 八进制77转十进制: 63
echo $((2#1010))         # 二进制1010转十进制: 10

# 随机数
echo $RANDOM            # 0-32767之间的随机数
echo $((RANDOM % 100))  # 0-99之间的随机数
echo $((RANDOM % 10))  # 0-9之间的随机数

# 使用 let 命令
let result=a+b
let a++

# 使用 expr (较老的方法)
result=$(expr $a + $b)

# 浮点数运算需要使用 bc
result=$(echo "scale=2; 10/3" | bc)  # 输出: 3.33
```

### 3. 数组 (Arrays)

#### 索引数组

```bash
# 数组定义
arr1=(apple banana cherry)
arr2[0]="first"
arr2[1]="second"
arr2[10]="eleventh"  # 稀疏数组

# 访问数组元素
echo ${arr1[0]}      # apple
echo ${arr1[1]}      # banana
echo ${arr1[@]}      # 所有元素: apple banana cherry
echo ${arr1[*]}      # 所有元素: apple banana cherry

# 数组长度
echo ${#arr1[@]}     # 3

# 数组索引
echo ${!arr1[@]}     # 0 1 2

# 数组切片
echo ${arr1[@]:1:2}  # banana cherry (从索引1开始，取2个元素)

# 添加元素
arr1+=(date)         # 追加元素
arr1[10]="ten"       # 指定位置

# 删除元素
unset arr1[1]        # 删除索引1的元素

# 数组排序
arr=(3 1 4 1 5 9 2 6)
sorted_arr=($(printf '%s\n' "${arr[@]}" | sort -n))

# 数组去重
# 推荐使用 "${}" 避免空格处理
declare -A seen
unique=()
for item in "${arr[@]}"; do
    if [[ ! ${seen[$item]} ]]; then
        seen[$item]=1
        unique+=("$item")
    fi
done

# 数组连接
arr1=(a b c)
arr2=(d e f)
combined=("${arr1[@]}" "${arr2[@]}")
```

#### 关联数组 (Associative Arrays)

```bash
# 声明关联数组
declare -A colors

# 赋值
colors[red]="#FF0000"
colors[green]="#00FF00"
colors[blue]="#0000FF"

# 或者一次性定义
declare -A colors=([red]="#FF0000" [green]="#00FF00" [blue]="#0000FF")

# 访问
echo ${colors[red]}        # #FF0000
echo ${colors[@]}          # 所有值
echo ${!colors[@]}         # 所有键: red green blue

# 遍历
for value in "${colors[@]}"; do
    echo "$value: ${colors[$value]}"
done

for key in "${!colors[@]}"; do
    echo "$key: ${colors[$key]}"
done
```

## 常用方法和操作

### 输入输出处理

```bash
# 读取用户输入
read -p "Enter your name: " name
read -p "Enter your names: " -ra names  # 数组 + 不转义
read -s -p "Enter password: " password  # 隐藏输入
read -t 5 -p "Quick response (5s): " response  # 超时

# 输入重定向
while IFS= read -r line; do # IFS 字符串分割
    echo "Line: $line"
done < file.txt


cat > config.ini << EOF # 多行字符串读取输入
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password

# 应用配置
APP_NAME=MyApp
APP_PORT=8080
DEBUG=true
EOF

text="apple,banana,orange,grape"
IFS=',' read -ra fruits <<< "$text" # 字符串读取输入

for fruit in "${fruits[@]}"; do
    echo $fruit
done

# 输出重定向
echo "message" > file.txt     # 覆盖
echo "message" >> file.txt    # 追加
echo "error" >&2              # 输出到标准错误 /dev/null
echo "error" 2>&1              # 输出到标准输出 /dev/null
```

### 条件判断

```bash
# 文件测试
if [[ -f "file.txt" ]]; then echo "File exists"; fi
if [[ -d "directory" ]]; then echo "Directory exists"; fi
if [[ -r "file.txt" ]]; then echo "File is readable"; fi
if [[ -w "file.txt" ]]; then echo "File is writable"; fi
if [[ -x "script.sh" ]]; then echo "File is executable"; fi

# 字符串测试
if [[ -z "$str" ]]; then echo "String is empty"; fi
if [[ -n "$str" ]]; then echo "String is not empty"; fi
if [[ "$str1" == "$str2" ]]; then echo "Strings are equal"; fi

# 数值测试
if (( num1 > num2 )); then echo "Numbers are equal"; fi
if (( num1 < num2 )); then echo "num1 is greater"; fi
if (( num1 == num2 )); then echo "num1 is less"; fi

# 包含与(&&)、或(||)、非(!) 的用法
if [[ -f "file.txt" && -r "file.txt" ]]; then
    echo "File exists and is readable"
fi

if [[ "$str" || ! "$str2" ]]; then
    echo "At least one string is empty"
fi

if (( num1 > num2 && num1 < 100 )); then
    echo "num1 is greater than num2 and less than 100"
fi
```

### 循环结构

```bash
# for 循环
for i in {1..10}; do
    echo $i
done

for file in *.txt; do
    echo "Processing: $file"
done

for (( i=0; i<10; i++ )); do
    echo $i
done

# while 循环
counter=0
while (( counter < 5 )); do
    echo $counter
    ((counter++))
done

# until 循环
until (( counter >= 10 )); do
    echo $counter
    ((counter++))
done
```

### 函数定义和使用

```bash
# 函数定义
function greet() {
    local name=$1
    local age=$2
    echo "Hello, $name! You are $age years old."
    return 0
}

# 或者简化语法
greet() {
    echo "Hello, $1!"
}

# 函数调用
greet "John" 25

# 返回值
get_sum() {
    local result=$(($1 + $2))
    echo $result
}

sum=$(get_sum 10 20)
echo "Sum is: $sum"
```

# I/O

## 1. 基本文件操作

### 创建文件

```bash
# 创建空文件
touch filename.txt

# 创建多个文件
touch file1.txt file2.txt file3.txt

# 使用重定向创建文件
echo "内容" > newfile.txt
echo "" > empty_file.txt

# 使用cat创建文件
cat > myfile.txt << EOF
这是第一行
这是第二行
EOF
```

### 删除文件

```bash
# 删除单个文件
rm filename.txt

# 删除多个文件
rm file1.txt file2.txt

# 强制删除（不提示）
rm -f filename.txt

# 交互式删除（确认提示）
rm -i filename.txt

# 删除目录及其内容
rm -r directory/
rm -rf directory/  # 强制递归删除
```

### 复制文件

```bash
# 复制文件
cp source.txt destination.txt

# 复制到目录
cp file.txt /path/to/directory/

# 复制多个文件到目录
cp file1.txt file2.txt /path/to/directory/

# 递归复制目录
cp -r source_dir/ destination_dir/

# 保持属性复制
cp -p file.txt backup.txt

# 交互式复制（确认覆盖）
cp -i source.txt destination.txt
```

### 移动/重命名文件

```bash
# 重命名文件
mv oldname.txt newname.txt

# 移动文件到目录
mv file.txt /path/to/directory/

# 移动多个文件
mv file1.txt file2.txt /path/to/directory/

# 移动目录
mv old_directory/ new_directory/
```

## 2. 文件内容操作

### 查看文件内容

```bash
# 显示完整文件内容
cat filename.txt

# 显示行号
cat -n filename.txt

# 显示文件头部（默认10行）
head filename.txt
head -n 20 filename.txt  # 显示前20行

# 显示文件尾部（默认10行）
tail filename.txt
tail -n 15 filename.txt  # 显示后15行

# 实时查看文件变化
tail -f logfile.txt

# 分页查看
less filename.txt
more filename.txt
```

### 编辑文件内容

```bash
# 使用重定向写入
echo "新内容" > file.txt          # 覆盖
echo "追加内容" >> file.txt       # 追加

# 使用printf（更好的格式控制）
printf "格式化内容: %s\n" "变量" > file.txt

# 字符串输出
printf "Hello %s\n" "World"
# 输出: Hello World

# 数字格式化
printf "整数: %d, 八进制: %o, 十六进制: %x\n" 255 255 255
# 输出: 整数: 255, 八进制: 377, 十六进制: ff

# 浮点数
printf "浮点数: %.2f\n" 3.14159
# 输出: 浮点数: 3.14

# 指定宽度
printf "%10s\n" "hello"        # 右对齐，总宽度10
printf "%-10s\n" "hello"       # 左对齐，总宽度10

# 数字宽度
printf "%5d\n" 42              # 右对齐，宽度5
printf "%05d\n" 42             # 用0填充，宽度5

# 浮点数精度
printf "%.3f\n" 3.14159        # 3位小数
printf "%8.2f\n" 3.14159       # 总宽度8，2位小数

printf "行1\n行2\n"            # \n 换行
printf "制表符:\t内容\n"        # \t 制表符
printf "退格:\b删除\n"          # \b 退格
printf "回车:\r覆盖\n"          # \r 回车
printf "反斜杠:\\\n"           # \\ 反斜杠
printf "百分号:%%\n"           # %% 百分号

# ANSI 颜色代码
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

printf "${RED}错误信息${NC}\n"
printf "${GREEN}成功信息${NC}\n"
printf "${BLUE}信息提示${NC}\n"

# 使用cat写入多行
cat > multiline.txt << EOF
第一行内容
第二行内容
包含$变量的行（不会被展开）
EOF

# 使用sed进行行内替换
sed 's/old_text/new_text/g' file.txt > new_file.txt
sed -i '' 's/old_text/new_text/g' file.txt  # 直接修改原文件

# 删除特定行
sed '3d' file.txt > new_file.txt    # 删除第3行
sed '/pattern/d' file.txt           # 删除匹配模式的行
```

## 3. 文件查找和搜索

### 使用 find 查找文件

```bash
# 按名称查找
find /path -name "filename.txt"
m               # 当前目录下所有txt文件

# 按类型查找
find /path -type f                 # 查找文件
find /path -type d                 # 查找目录

# 按大小查找
# c : 字节 (bytes)
# w : 字 (words, 2字节)
# b : 块 (blocks, 512字节，默认)
# k : KB (1024字节)
# M : MB (1024KB)
# G : GB (1024MB)
find /path -size +100M             # 大于100MB的文件
find /path -size -1M               # 小于1MB的文件

# 按修改时间查找
# N天前被访问 -atime N
# N分钟前被访问 -amin N
# N天前被修改内容 -mtime N
# N分钟前被修改内容 -mmin N
# N分钟前被修改元数据 -ctime N
# N分钟前被修改元数据 -cmin N
#
find /path -mtime -7               # 7天内修改的文件
find /path -mtime +30              # 30天前修改的文件

# 组合条件
find /path -name "*.log" -size +10M -mtime -7

# 执行命令
find /path -name "*.sh" -exec chmod +x {} +
```

### 使用 grep 搜索内容

```bash
# 基本搜索
grep "pattern" filename.txt

# 不区分大小写
grep -i "pattern" filename.txt

# 显示行号
grep -n "pattern" filename.txt

# 递归搜索目录
grep -r "pattern" /path/to/directory/

# 搜索多个文件
grep "pattern" *.txt

# 反向匹配（不包含pattern的行）
grep -v "pattern" filename.txt

# 使用正则表达式
grep -E "^[0-9]+$" filename.txt    # 匹配纯数字行

# 静默模式，不输出匹配的行内容
grep -q "pattern" filename.txt
```

## 4. 文件权限和属性

### 查看文件属性

```bash
# 详细列表
ls -l filename.txt

# 显示隐藏文件
ls -la

# 显示文件大小（人类可读）
ls -lh filename.txt

# 显示文件inode信息
ls -li filename.txt

# 显示文件统计信息
stat filename.txt
```

### 修改文件权限

```bash
# 数字方式
chmod 644 filename.txt    # rw-r--r--
chmod 755 filename.txt    # rwxr-xr-x
chmod 777 filename.txt    # rwxrwxrwx

# 符号方式
chmod u+x filename.txt    # 给所有者添加执行权限
chmod g-w filename.txt    # 移除组写权限
chmod o=r filename.txt    # 设置其他用户只读

# 递归修改目录权限
chmod -R 755 directory/
```

### 修改文件所有者

```bash
# 修改所有者
chown user:group filename.txt
chown user filename.txt

# 递归修改
chown -R user:group directory/
```

## 5. 高级文件操作

### 文件比较

```bash
# 比较两个文件
diff file1.txt file2.txt

# 并排比较
diff -y file1.txt file2.txt

# 忽略空白差异
diff -w file1.txt file2.txt
```

### 文件压缩和解压

```bash
# tar归档
tar -czf archive.tar.gz directory/     # 创建压缩归档
tar -xzf archive.tar.gz               # 解压

# zip压缩
zip -r archive.zip directory/         # 创建zip文件
unzip archive.zip                     # 解压zip文件

# gzip压缩单个文件
gzip filename.txt                     # 压缩（原文件被替换）
gunzip filename.txt.gz               # 解压
```

### 文件链接

> 原理上，硬链接和源文件的 inode 节点号相同，两者互为硬链接。软连接和源文件的 inode 节点号不同，进而指向的 block 也不同，软连接 block 中存放了源文件的路径名。 实际上，硬链接和源文件是同一份文件，而软连接是独立的文件，类似于快捷方式，存储着源文件的位置信息便于指向。

```bash
# 创建硬链接
ln original.txt hardlink.txt

# 创建软链接（符号链接）
ln -s /path/to/original.txt symlink.txt

# 查看链接信息
ls -l symlink.txt
readlink symlink.txt
```

## 6. 批量文件操作

### 使用通配符

```bash
# 操作所有txt文件
cp *.txt backup/
rm *.tmp
chmod 644 *.sh

# 使用花括号展开
cp file.{txt,bak} directory/
touch file{1..10}.txt              # 创建file1.txt到file10.txt
```

### 使用 xargs 进行批量处理

```bash
# 查找并删除
find . -name "*.tmp" | xargs rm

# 批量修改权限
find . -type f -name "*.sh" | xargs chmod +x

# 批量处理文件内容
find . -name "*.txt" | xargs grep "pattern"
```

### 使用 while 循环处理文件列表

```bash
# 读取文件列表并处理
while IFS= read -r file; do
    echo "处理文件: $file"
    cp "$file" "backup_$file"
done < filelist.txt

# 处理find结果
find . -name "*.log" | while read -r logfile; do
    echo "处理日志: $logfile"
    gzip "$logfile"
done
```

# Network

## 1. curl 简介

curl（Client URL）是一个在命令行或脚本中用来传输数据的工具和库。它支持众多协议，包括 HTTP、HTTPS、FTP、FTPS、SCP、SFTP、TFTP、LDAP、DAP、DICT、TELNET、FILE、POP3、IMAP、SMTP 等。

### 核心特性

- 跨平台支持（Linux、macOS、Windows）
- 支持多种协议
- 灵活的认证方式
- 强大的调试功能
- 支持代理和重定向

## 2. 基本语法

```bash
curl [options] [URL...]
```

## 3. 常用选项详解

### 3.1 基础请求选项

| 选项            | 说明                   | 示例                                           |
| --------------- | ---------------------- | ---------------------------------------------- |
| `-X, --request` | 指定 HTTP 请求方法     | `curl -X POST url`                             |
| `-H, --header`  | 添加 HTTP 头           | `curl -H "Content-Type: application/json" url` |
| `-d, --data`    | 发送 POST 数据         | `curl -d "param=value" url`                    |
| `-G, --get`     | 将-d 数据作为 GET 参数 | `curl -G -d "q=search" url`                    |

### 3.2 输出控制选项

| 选项                | 说明                           | 示例                         |
| ------------------- | ------------------------------ | ---------------------------- |
| `-o, --output`      | 保存到指定文件                 | `curl -o file.html url`      |
| `-O, --remote-name` | 使用远程文件名保存             | `curl -O url/file.zip`       |
| `-s, --silent`      | 静默模式                       | `curl -s url`                |
| `-v, --verbose`     | 详细输出                       | `curl -v url`                |
| `-w, --write-out`   | 自定义输出格式                 | `curl -w "%{http_code}" url` |
| `-N, --no-buffer`   | 禁用输出缓冲, 确保数据实时显示 | `curl -N url`                |

### 3.3 认证选项

| 选项               | 说明               | 示例                             |
| ------------------ | ------------------ | -------------------------------- |
| `-u, --user`       | 用户认证           | `curl -u username:password url`  |
| `-b, --cookie`     | 发送 cookie        | `curl -b "session=abc123" url`   |
| `-c, --cookie-jar` | 保存 cookie        | `curl -c cookies.txt url`        |
| `--oauth2-bearer`  | OAuth2 Bearer 令牌 | `curl --oauth2-bearer token url` |

### 3.4 连接控制选项

| 选项                | 说明              | 示例                            |
| ------------------- | ----------------- | ------------------------------- |
| `-L, --location`    | 跟随重定向        | `curl -L url`                   |
| `-m, --max-time`    | 最大执行时间      | `curl -m 30 url`                |
| `--connect-timeout` | 连接超时时间      | `curl --connect-timeout 10 url` |
| `-k, --insecure`    | 忽略 SSL 证书错误 | `curl -k https://url`           |
| `--proxy`           | 使用代理          | `curl --proxy proxy:port url`   |

## 4. 实际应用案例

### 4.1 GET 请求

```bash
# 基本GET请求
curl https://api.example.com/users

# 带参数的GET请求
curl "https://api.example.com/search?q=keyword&limit=10"

# 带自定义头的GET请求
curl -H "Authorization: Bearer token123" \
     -H "Accept: application/json" \
     https://api.example.com/profile
```

### 4.2 POST 请求

```bash
# 发送表单数据
curl -X POST \
     -d "username=john&password=secret" \
     https://api.example.com/login

# 发送JSON数据
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"name":"John","email":"john@example.com"}' \
     https://api.example.com/users

# 从文件发送数据
curl -X POST \
     -H "Content-Type: application/json" \
     -d @data.json \
     https://api.example.com/users
```

### 4.3 文件上传

```bash
# 上传单个文件
curl -F "file=@document.pdf" \
     https://api.example.com/upload

# 上传多个文件
curl -F "file1=@image1.jpg" \
     -F "file2=@image2.jpg" \
     https://api.example.com/upload

# 指定文件类型
curl -F "file=@data.csv;type=text/csv" \
     https://api.example.com/upload
```

### 4.4 下载文件

```bash
# 下载并保存为指定文件名
curl -o myfile.zip https://example.com/file.zip

# 使用远程文件名
curl -O https://example.com/file.zip

# 断点续传
curl -C - -O https://example.com/largefile.zip

# 下载多个文件
curl -O https://example.com/file1.zip \
     -O https://example.com/file2.zip
```

### 4.5 Cookie 处理

```bash
# 保存cookie到文件
curl -c cookies.txt https://example.com/login

# 从文件读取cookie
curl -b cookies.txt https://example.com/dashboard

# 同时保存和发送cookie
curl -b cookies.txt -c cookies.txt https://example.com/api
```

## 5. 高级用法

### 5.1 自定义输出格式

```bash
# 获取HTTP状态码
curl -w "%{http_code}\n" -s -o /dev/null https://example.com

# 获取响应时间
curl -w "Total time: %{time_total}s\n" -s -o /dev/null https://example.com

# 详细的输出格式
curl -w "Status: %{http_code}\nTime: %{time_total}s\nSize: %{size_download} bytes\n" \
     -s -o /dev/null https://example.com
```

### 5.2 条件请求

```bash
# If-Modified-Since
curl -H "If-Modified-Since: Wed, 21 Oct 2015 07:28:00 GMT" \
     https://example.com/resource

# If-None-Match (ETag)
curl -H "If-None-Match: \"abc123\"" \
     https://example.com/resource
```

### 5.3 范围请求

```bash
# 下载文件的前1024字节
curl -r 0-1023 https://example.com/largefile

# 下载文件的最后500字节
curl -r -500 https://example.com/largefile

# 下载多个范围
curl -r 0-499,1000-1499 https://example.com/largefile
```

### 5.4 并发请求

```bash
# 使用xargs并发执行多个curl请求
echo "url1 url2 url3" | xargs -n 1 -P 3 curl -s

# 在脚本中实现并发
#!/bin/bash
urls=("url1" "url2" "url3")
for url in "${urls[@]}"; do
    curl -s "$url" &
done
wait
```

## 6. 调试技巧

### 6.1 详细输出

```bash
# 查看完整的HTTP交互
curl -v https://example.com

# 只显示响应头
curl -I https://example.com

# 追踪重定向
curl -v -L https://example.com
```

### 6.2 保存调试信息

```bash
# 保存头信息到文件
curl -D headers.txt https://example.com

# 保存详细日志
curl -v https://example.com 2>debug.log

# 分别保存头和内容
curl -D headers.txt -o content.html https://example.com
```

### 6.3 测试连接

```bash
# 测试连接时间
curl -w "@curl-format.txt" -s -o /dev/null https://example.com

# curl-format.txt 内容示例：
#     time_namelookup:  %{time_namelookup}s\n
#        time_connect:  %{time_connect}s\n
#     time_appconnect:  %{time_appconnect}s\n
#    time_pretransfer:  %{time_pretransfer}s\n
#       time_redirect:  %{time_redirect}s\n
#  time_starttransfer:  %{time_starttransfer}s\n
#                     ----------\n
#          time_total:  %{time_total}s\n
```

# Process

## 1. 进程基础概念

### 什么是进程

```bash
# 进程是正在运行的程序实例
# 每个进程都有唯一的进程ID (PID)
# 进程有父进程(PPID)和子进程关系
# 进程状态：运行(Running)、睡眠(Sleeping)、停止(Stopped)、僵尸(Zombie)
```

### 查看进程信息

```bash
# 显示当前用户的进程
ps

# 显示所有进程
ps -ef

# 实时显示进程
top
htop  # 需要安装

# 显示进程树
pstree
ps --forest

# 显示特定进程信息
ps -p 1234      # 显示PID为1234的进程
ps -u username  # 显示特定用户的进程
```

## 2. 启动和管理进程

### 前台和后台进程

```bash
# 前台运行（默认）
command

# 后台运行
command &

# 将前台进程放到后台
# 先按 Ctrl+Z 暂停进程，然后：
bg

# 将后台进程调到前台
fg

# 将特定作业调到前台
fg %1    # 调用作业号为1的进程
```

### 作业控制

```bash
# 查看当前作业
jobs

# 查看作业详细信息
jobs -l

# 结束作业
kill %1         # 结束作业号为1的进程
kill %+         # 结束当前作业
kill %-         # 结束上一个作业

# 脱离终端运行
nohup command &
```

## 3. 进程终止和信号

### 终止进程

```bash
# 使用PID终止进程
kill 1234

# 强制终止进程
kill -9 1234
kill -KILL 1234

# 按名称终止进程
killall process_name
pkill process_name

# 交互式终止进程
killall -i process_name
```

### 信号类型

```bash
# 常用信号
kill -l         # 列出所有信号

# SIGTERM (15) - 正常终止请求
kill -15 1234
kill -TERM 1234

# SIGKILL (9) - 强制终止
kill -9 1234
kill -KILL 1234

# SIGHUP (1) - 重新加载配置
kill -1 1234
kill -HUP 1234

# SIGSTOP (19) - 暂停进程
kill -19 1234
kill -STOP 1234

# SIGCONT (18) - 继续进程
kill -18 1234
kill -CONT 1234
```

## 4. 进程监控和分析

### 系统资源监控

```bash
# CPU和内存使用情况
top
htop

# 按CPU使用率排序
top -o %CPU

# 按内存使用率排序
top -o %MEM

# 显示特定用户的进程
top -u username

# 系统负载
uptime
w

# 内存使用情况
free -h
cat /proc/meminfo
```

### 进程详细信息

```bash
# 查看进程详细状态
cat /proc/1234/status

# 查看进程命令行参数
cat /proc/1234/cmdline

# 查看进程环境变量
cat /proc/1234/environ

# 查看进程打开的文件
lsof -p 1234

# 查看进程网络连接
netstat -p | grep 1234
ss -p | grep 1234
```

### 进程资源使用

```bash
# 查看进程资源限制
ulimit -a

# 设置资源限制
ulimit -n 4096      # 设置文件描述符限制
ulimit -u 1024      # 设置用户进程数限制

# 查看进程I/O统计
iotop               # 需要安装
pidstat -d 1        # 需要安装sysstat包
```

## 5. 进程间通信(IPC)

### 管道通信

```bash
# 匿名管道
command1 | command2

# 命名管道(FIFO)
mkfifo mypipe
echo "data" > mypipe &
cat < mypipe
```

### 信号通信

```bash
# 自定义信号处理脚本
#!/bin/bash
trap 'echo "收到SIGUSR1信号"' USR1
trap 'echo "收到SIGUSR2信号"' USR2
trap 'echo "正在清理..."; exit' INT TERM

while true; do
    sleep 1
done
```

## 6. 守护进程和服务

### 创建守护进程

```bash
#!/bin/bash
# daemon.sh - 简单守护进程示例

# 创建守护进程
(
    # 切换到根目录
    cd /

    # 重定向标准输入输出
    exec < /dev/null
    exec > /var/log/mydaemon.log 2>&1

    # 主循环
    while true; do
        echo "$(date): 守护进程运行中..."
        sleep 60
    done
) &

echo $! > /var/run/mydaemon.pid
```

### 使用 systemd 管理服务

```bash
# 查看服务状态
systemctl status service_name

# 启动/停止/重启服务
systemctl start service_name
systemctl stop service_name
systemctl restart service_name

# 开机自启动
systemctl enable service_name
systemctl disable service_name

# 查看服务日志
journalctl -u service_name
journalctl -f -u service_name  # 实时查看
```

## 7. 进程调度和优先级

### 进程优先级

```bash
# 查看进程优先级（NI列）
ps -eo pid,ni,cmd

# 设置进程优先级（nice值：-20到19）
nice -n 10 command          # 低优先级启动
nice --10 command           # 高优先级启动

# 修改运行中进程的优先级
renice -n 5 -p 1234        # 设置PID 1234的nice值为5
renice -n -5 -u username   # 设置用户的所有进程nice值
```

### CPU 亲和性

```bash
# 查看进程CPU亲和性
taskset -p 1234

# 设置进程CPU亲和性
taskset -p 0x3 1234        # 绑定到CPU 0和1
taskset -c 0,1 command     # 在CPU 0和1上启动命令
```

## 8. 进程性能分析

### 性能分析工具

```bash
# 系统调用跟踪
strace -p 1234             # 跟踪运行中的进程
strace command             # 跟踪新进程

# 进程性能统计
time command               # 执行时间统计
/usr/bin/time -v command   # 详细统计信息

# 进程内存映射
pmap 1234
cat /proc/1234/maps

# 进程线程信息
ps -eLf | grep 1234        # 显示线程
top -H -p 1234             # 显示进程的线程
```

### 性能监控脚本

```bash
#!/bin/bash
# monitor_process.sh - 进程监控脚本

PROCESS_NAME="$1"
INTERVAL=5

if [ -z "$PROCESS_NAME" ]; then
    echo "用法: $0 <进程名>"
    exit 1
fi

while true; do
    PID=$(pgrep "$PROCESS_NAME" | head -1)
    if [ -n "$PID" ]; then
        echo "=== $(date) ==="
        echo "进程 $PROCESS_NAME (PID: $PID) 状态:"
        ps -p "$PID" -o pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,cmd
        echo "内存详情:"
        cat /proc/$PID/status | grep -E "(VmPeak|VmSize|VmRSS|VmData)"
        echo
    else
        echo "$(date): 进程 $PROCESS_NAME 未找到"
    fi
    sleep $INTERVAL
done
```

## 9. 批量进程操作

### 批量终止进程

```bash
# 终止所有匹配的进程
killall firefox
pkill -f "python script.py"

# 终止用户的所有进程
pkill -u username
killall -u username

# 条件终止进程
ps aux | grep 'defunct' | awk '{print $2}' | xargs kill -9

# 终止占用端口的进程
lsof -ti:8080 | xargs kill -9
```

### 进程状态统计

```bash
#!/bin/bash
# process_stats.sh - 进程状态统计

echo "=== 进程状态统计 ==="
ps aux --no-headers | awk '
{
    states[$8]++
    total++
}
END {
    for (state in states) {
        printf "%-10s: %d (%.1f%%)\n", state, states[state], states[state]*100/total
    }
    print "总计: " total
}'

echo -e "\n=== 内存使用前10的进程 ==="
ps aux --no-headers --sort=-%mem | head -10 | awk '{printf "%-8s %6s %6s %s\n", $2, $4"%", $6"K", $11}'

echo -e "\n=== CPU使用前10的进程 ==="
ps aux --no-headers --sort=-%cpu | head -10 | awk '{printf "%-8s %6s %6s %s\n", $2, $3"%", $6"K", $11}'
```

## 10. 进程调试和故障排除

### 进程调试

```bash
# 使用gdb调试运行中的进程
gdb -p 1234

# 生成核心转储
gcore 1234                 # 生成进程的核心转储
kill -QUIT 1234            # 发送SIGQUIT信号生成转储

# 查看进程堆栈
pstack 1234                # 显示进程调用栈
```

### 故障排除脚本

```bash
#!/bin/bash
# troubleshoot_process.sh - 进程故障排除脚本

PROCESS_NAME="$1"

if [ -z "$PROCESS_NAME" ]; then
    echo "用法: $0 <进程名或PID>"
    exit 1
fi

# 检查是数字(PID)还是进程名
if [[ "$PROCESS_NAME" =~ ^[0-9]+$ ]]; then
    PID="$PROCESS_NAME"
    PROCESS_NAME=$(ps -p "$PID" -o comm=)
else
    PID=$(pgrep "$PROCESS_NAME" | head -1)
fi

if [ -z "$PID" ]; then
    echo "错误: 进程 $PROCESS_NAME 未找到"
    exit 1
fi

echo "=== 进程 $PROCESS_NAME (PID: $PID) 故障排除 ==="

echo -e "\n1. 基本信息:"
ps -p "$PID" -o pid,ppid,user,%cpu,%mem,vsz,rss,tty,stat,start,time,cmd

echo -e "\n2. 进程状态:"
cat /proc/$PID/status | grep -E "(State|Threads|VmPeak|VmSize|VmRSS)"

echo -e "\n3. 打开的文件描述符数量:"
ls /proc/$PID/fd | wc -l

echo -e "\n4. 网络连接:"
netstat -anp 2>/dev/null | grep "$PID" | head -10

echo -e "\n5. 最近的系统调用:"
timeout 3 strace -p "$PID" 2>&1 | head -20

echo -e "\n6. 内存映射:"
head -10 /proc/$PID/maps

echo -e "\n7. 环境变量数量:"
cat /proc/$PID/environ | tr '\0' '\n' | wc -l
```

## 11. 实用脚本和工具

### 进程监控告警脚本

```bash
#!/bin/bash
# process_monitor.sh - 进程监控告警脚本

CONFIG_FILE="monitor.conf"
LOG_FILE="process_monitor.log"

# 配置文件格式: process_name:max_cpu:max_memory:email
# 例如: nginx:80:500:admin@example.com

log_message() {
    echo "$(date '+%Y-%m-%d %H:%M:%S'): $1" | tee -a "$LOG_FILE"
}

check_process() {
    local process_name="$1"
    local max_cpu="$2"
    local max_mem="$3"
    local email="$4"

    local pid=$(pgrep "$process_name" | head -1)
    if [ -z "$pid" ]; then
        log_message "警告: 进程 $process_name 未运行"
        return
    fi

    local cpu_usage=$(ps -p "$pid" -o %cpu= | tr -d ' ')
    local mem_usage=$(ps -p "$pid" -o %mem= | tr -d ' ')

    cpu_usage=${cpu_usage%.*}  # 取整数部分
    mem_usage=${mem_usage%.*}

    if [ "$cpu_usage" -gt "$max_cpu" ]; then
        log_message "警告: $process_name CPU使用率过高: $cpu_usage%"
    fi

    if [ "$mem_usage" -gt "$max_mem" ]; then
        log_message "警告: $process_name 内存使用率过高: $mem_usage%"
    fi
}

# 读取配置文件并检查进程
while IFS=':' read -r process_name max_cpu max_mem email; do
    [ -z "$process_name" ] && continue
    check_process "$process_name" "$max_cpu" "$max_mem" "$email"
done < "$CONFIG_FILE"
```

### 进程清理脚本

```bash
#!/bin/bash
# cleanup_processes.sh - 进程清理脚本

# 清理僵尸进程
echo "清理僵尸进程..."
ps aux | awk '$8 ~ /^Z/ {print $2}' | while read pid; do
    if [ -n "$pid" ]; then
        echo "发现僵尸进程: $pid"
        # 僵尸进程需要父进程清理，通常需要重启父进程
        ppid=$(ps -o ppid= -p "$pid" 2>/dev/null)
        if [ -n "$ppid" ]; then
            echo "父进程: $ppid"
        fi
    fi
done

# 清理长时间运行的临时进程
echo "清理长时间运行的临时进程..."
ps aux | awk '$11 ~ /tmp/ && $10 > "01:00" {print $2, $11}' | while read pid cmd; do
    echo "发现长时间运行的临时进程: $pid ($cmd)"
    # 可以选择是否自动终止
    # kill "$pid"
done

# 清理高CPU使用率的异常进程
echo "检查高CPU使用率进程..."
ps aux --sort=-%cpu | head -10 | awk 'NR>1 && $3>80 {print $2, $3, $11}' | while read pid cpu cmd; do
    echo "高CPU使用率进程: $pid (CPU: $cpu%) - $cmd"
done
```
