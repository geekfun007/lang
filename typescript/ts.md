# pnpm

### 安装与初始化

```bash
# 全局安装 pnpm
npm install -g pnpm

# 检查 pnpm 版本
pnpm --version

# 初始化新项目
pnpm init

# 从现有 package.json 安装依赖
pnpm install
# 简写
pnpm i
```

### 依赖管理

```bash
# 安装生产依赖
pnpm add <package-name>
pnpm add -save-dev <package-name> # 安装开发依赖
pnpm add -save-peer <package-name> # 安装对等依赖
pnpm add -save-optional <package-name> # 安装可选依赖
pnpm add -g <package-name> # 安装全局包

# 安装指定版本
pnpm add express@4.18.0
pnpm add express@^4.18.0
pnpm add express@~4.18.0

# 移除依赖
pnpm remove <package-name>
pnpm rm <package-name>

# 查看已安装的包
pnpm list
pnpm ls

# 查看过时的包
pnpm outdated

# 更新依赖
pnpm update
pnpm up

# 更新指定包
pnpm update <package-name>
pnpm up <package-name>
```

### 脚本执行

```bash
# 运行 package.json 中的脚本
pnpm run <script-name>
pnpm run build
pnpm run dev

# 简写（对于常用脚本）
pnpm build
pnpm dev
pnpm start
pnpm test

# 运行任意命令
pnpm exec <command>
pnpm exec tsc --version
```

## package.json 配置详解

### 基本结构

```json
{
  "name": "project-name", // 项目名称
  "version": "1.0.0", // 版本号
  "description": "项目描述", // 项目描述
  "main": "dist/index.js", // 入口文件
  "type": "module", // 模块类型 (module/commonjs)
  "packageManager": "pnpm@8.0.0", // 包管理器版本

  // 脚本命令
  "scripts": {
    "build": "tsc",
    "build:watch": "tsc --watch",
    "dev": "ts-node src/index.ts",
    "start": "node dist/index.js"
  },

  // 生产依赖
  "dependencies": {
    "express": "^4.18.0"
  },

  // 开发依赖
  "devDependencies": {
    "typescript": "^5.0.0"
  },

  // 对等依赖
  "peerDependencies": {
    "react": ">=16.8.0"
  },

  // 可选依赖
  "optionalDependencies": {
    "fsevents": "^2.3.0"
  }
}
```

### 版本号规则

```json
{
  "dependencies": {
    "exact": "1.0.0", // 精确版本
    "caret": "^1.0.0", // 兼容版本 (1.x.x)
    "tilde": "~1.0.0", // 补丁版本 (1.0.x)
    "range": ">=1.0.0 <2.0.0", // 版本范围
    "latest": "*", // 最新版本
    "beta": "1.0.0-beta.1", // 预发布版本
    "git": "git+https://github.com/user/repo.git",
    "file": "file:../local-package"
  }
}
```

## 高级用法

### 环境变量和配置

```bash
# 设置环境变量
NODE_ENV=production pnpm start

# 使用 .env 文件
pnpm add -D dotenv
```

### 发布和版本管理

```bash
# 发布包
pnpm publish

# 发布到特定注册表
pnpm publish --registry https://registry.npmjs.org/

# 版本管理
pnpm version patch    # 1.0.0 -> 1.0.1
pnpm version minor    # 1.0.0 -> 1.1.0
pnpm version major    # 1.0.0 -> 2.0.0
pnpm version prerelease # 1.0.0 -> 1.0.1-0
```

# DataType

## 基础类型

### 1. 原始类型

```typescript
// 数字类型
let age: number = 25;
let price: number = 99.99;

// 字符串类型
let name: string = "Alice";
let message: string = `Hello, ${name}`;

// 布尔类型
let isActive: boolean = true;
let isCompleted: boolean = false;

// null 和 undefined
let nullValue: null = null;
let undefinedValue: undefined = undefined;

// symbol 类型
let sym: symbol = Symbol("key");

// bigint 类型
let bigNumber: bigint = 100n;
```

### 2. 复合类型

```typescript
// 数组类型
let numbers: number[] = [1, 2, 3, 4];
let names: Array<string> = ["Alice", "Bob", "Charlie"];

// 元组类型
let person: [string, number] = ["Alice", 25];
let rgb: [number, number, number] = [255, 0, 0];

// 枚举类型
enum Color {
  Red = "red",
  Green = "green",
  Blue = "blue",
}
let favoriteColor: Color = Color.Red;

// any 类型（尽量避免使用）
let anything: any = "hello";
anything = 42;
anything = true;

// unknown 类型（更安全的 any）
let userInput: unknown = "hello";
if (typeof userInput === "string") {
  console.log(userInput.toUpperCase());
}

// void 类型（主要用于函数返回值）
function logMessage(): void {
  console.log("Hello World");
}

// never 类型（永远不会返回的函数）
function throwError(): never {
  throw new Error("Something went wrong");
}
```

## 对象类型

### 1. 对象字面量类型

```typescript
// 基本对象类型
let user: {
  name: string;
  age: number;
  email?: string; // 可选属性
  readonly id: number; // 只读属性
} = {
  name: "Alice",
  age: 25,
  id: 1,
};

// 索引签名
let scores: {
  [subject: string]: number;
} = {
  math: 95,
  english: 87,
  science: 92,
};
```

### 2. 接口（Interface）

```typescript
interface User {
  readonly id: number;
  name: string;
  age: number;
  email?: string;
  greet(): string;
}

interface Admin extends User {
  permissions: string[];
  manageUsers(): void;
}

// 实现接口
class AdminUser implements Admin {
  readonly id: number;
  name: string;
  age: number;
  email?: string;
  permissions: string[];

  constructor(id: number, name: string, age: number, permissions: string[]) {
    this.id = id;
    this.name = name;
    this.age = age;
    this.permissions = permissions;
  }

  greet(): string {
    return `Hello, I'm ${this.name}`;
  }

  manageUsers(): void {
    console.log("Managing users...");
  }
}
```

### 3. 类型别名（Type Alias）

```typescript
// 基本类型别名
type UserID = string | number;
type Status = "pending" | "approved" | "rejected";

// 对象类型别名
type Point = {
  x: number;
  y: number;
};

type Circle = {
  center: Point;
  radius: number;
};

// 函数类型别名
type EventHandler = (event: string) => void;
type Calculator = (a: number, b: number) => number;
```

## 联合类型与交叉类型

### 1. 联合类型（Union Types）

```typescript
// 基本联合类型
let id: string | number = "abc123";
id = 12345; // 也可以是数字

// 字面量联合类型
type Theme = "light" | "dark" | "auto";
let currentTheme: Theme = "dark";

// 函数参数联合类型
function formatValue(value: string | number): string {
  if (typeof value === "string") {
    return value.toUpperCase();
  }
  return value.toString();
}
```

### 2. 交叉类型（Intersection Types）

```typescript
type Person = {
  name: string;
  age: number;
};

type Employee = {
  employeeId: string;
  department: string;
};

// 交叉类型
type EmployeePerson = Person & Employee;

let employee: EmployeePerson = {
  name: "Alice",
  age: 30,
  employeeId: "EMP001",
  department: "Engineering",
};
```

## 函数类型

### 1. 函数声明与表达式

```typescript
// 函数声明
function add(a: number, b: number): number {
  return a + b;
}

// 函数表达式
const multiply = (a: number, b: number): number => {
  return a * b;
};

// 简化箭头函数
const subtract = (a: number, b: number): number => a - b;

// 可选参数
function greet(name: string, title?: string): string {
  return title ? `Hello, ${title} ${name}` : `Hello, ${name}`;
}

// 默认参数
function createUser(name: string, age: number = 18): User {
  return { id: 1, name, age };
}

// 剩余参数
function sum(...numbers: number[]): number {
  return numbers.reduce((total, num) => total + num, 0);
}
```

### 2. 函数重载

```typescript
// 重载签名
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: number[], b: number[]): number[];

// 实现签名
function add(a: any, b: any): any {
  if (typeof a === "number" && typeof b === "number") {
    return a + b;
  }
  if (typeof a === "string" && typeof b === "string") {
    return a + b;
  }
  if (Array.isArray(a) && Array.isArray(b)) {
    return [...a, ...b];
  }
}

// 使用
console.log(add(1, 2)); // 3
console.log(add("hello", "world")); // "helloworld"
console.log(add([1, 2], [3, 4])); // [1, 2, 3, 4]

// 根据参数数量的重载
function createElement(tag: string): HTMLElement;
function createElement(tag: string, content: string): HTMLElement;
function createElement(
  tag: string,
  content: string,
  className: string
): HTMLElement;

function createElement(
  tag: string,
  content?: string,
  className?: string
): HTMLElement {
  const element = document.createElement(tag);
  if (content) {
    element.textContent = content;
  }
  if (className) {
    element.className = className;
  }
  return element;
}

// 使用
const div1 = createElement("div");
const div2 = createElement("div", "Hello");
const div3 = createElement("div", "Hello", "my-class");

// 泛型重载签名
function process<T>(data: T[]): T[];
function process<T, U>(data: T[], transformer: (item: T) => U): U[];

// 实现
function process<T, U = T>(data: T[], transformer?: (item: T) => U): U[] {
  if (transformer) {
    return data.map(transformer);
  }
  return data;
}

// 使用
const numbers = [1, 2, 3];
const doubled = process(numbers, (x) => x * 2); // number[]
const strings = process(numbers, (x) => x.toString()); // string[]
const original = process(numbers); // number[]
```

## 泛型（Generics）

### 1. 基本泛型

```typescript
// 泛型函数
function identity<T>(arg: T): T {
  return arg;
}

let stringResult = identity<string>("hello");
let numberResult = identity<number>(42);

// 泛型数组
function getFirstElement<T>(arr: T[]): T | undefined {
  return arr.length > 0 ? arr[0] : undefined;
}

// 泛型接口
interface Container<T> {
  value: T;
  getValue(): T;
  setValue(value: T): void;
}

class Box<T> implements Container<T> {
  constructor(public value: T) {}

  getValue(): T {
    return this.value;
  }

  setValue(value: T): void {
    this.value = value;
  }
}
```

### 2. 泛型约束

```typescript
// 基本约束
interface Lengthwise {
  length: number;
}

function logLength<T extends Lengthwise>(arg: T): T {
  console.log(arg.length);
  return arg;
}

// keyof 约束
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}

let person = { name: "Alice", age: 30, email: "alice@example.com" };
let name = getProperty(person, "name"); // string
let age = getProperty(person, "age"); // number

// infer 类型推导（只能在条件类型的 extends 子句中使用）
type ExtractGeneric<T> = T extends Array<infer U>
  ? U
  : T extends Promise<infer V>
  ? V
  : T;

type Test1 = ExtractGeneric<string[]>; // string
type Test2 = ExtractGeneric<Promise<number>>; // number
```

## 高级类型

### 1. 映射类型

```typescript
// Partial - 所有属性变为可选
type PartialUser = Partial<User>;

// Required - 所有属性变为必需
type RequiredUser = Required<User>;

// Readonly - 所有属性变为只读
type ReadonlyUser = Readonly<User>;

// Pick - 选择特定属性
type UserBasic = Pick<User, "id" | "name">;

// Omit - 排除特定属性
type UserWithoutId = Omit<User, "id">;

// Record - 创建键值对类型
type UserRoles = Record<string, string[]>;
```

### 2. 条件类型

```typescript
// 基本条件类型
type IsString<T> = T extends string ? true : false;

type Test1 = IsString<string>; // true
type Test2 = IsString<number>; // false

// 实用条件类型
type NonNullable<T> = T extends null | undefined ? never : T;

type SafeString = NonNullable<string | null>; // string
```

### 3. 工具类型

```typescript
// ReturnType - 获取函数返回类型
type AddReturn = ReturnType<typeof add>; // number

// Parameters - 获取函数参数类型
type AddParams = Parameters<typeof add>; // [number, number]

// ConstructorParameters - 获取构造函数参数类型
type BoxParams = ConstructorParameters<typeof Box>; // [unknown]

// InstanceType - 获取构造函数实例类型
type BoxInstance = InstanceType<typeof Box>; // Box<unknown>
```

## 类与继承

### 1. 基本类

```typescript
class Animal {
  protected name: string;
  private age: number;
  public species: string;

  constructor(name: string, age: number, species: string) {
    this.name = name;
    this.age = age;
    this.species = species;
  }

  public makeSound(): void {
    console.log(`${this.name} makes a sound`);
  }

  protected getAge(): number {
    return this.age;
  }
}

class Dog extends Animal {
  private breed: string;

  constructor(name: string, age: number, breed: string) {
    super(name, age, "Canine");
    this.breed = breed;
  }

  public makeSound(): void {
    console.log(`${this.name} barks`);
  }

  public getInfo(): string {
    return `${this.name} is a ${this.breed}, age ${this.getAge()}`;
  }
}
```

### 2. 抽象类

```typescript
abstract class Shape {
  abstract calculateArea(): number;

  public displayArea(): void {
    console.log(`Area: ${this.calculateArea()}`);
  }
}

class Rectangle extends Shape {
  constructor(private width: number, private height: number) {
    super();
  }

  calculateArea(): number {
    return this.width * this.height;
  }
}

class Circle extends Shape {
  constructor(private radius: number) {
    super();
  }

  calculateArea(): number {
    return Math.PI * this.radius * this.radius;
  }
}
```

## 装饰器（Decorators）

```typescript
// 类装饰器
function Component(target: any) {
  target.prototype.isComponent = true;
}

// 方法装饰器
function Log(target: any, propertyKey: string, descriptor: PropertyDescriptor) {
  const originalMethod = descriptor.value;
  descriptor.value = function (...args: any[]) {
    console.log(`Calling ${propertyKey} with args:`, args);
    return originalMethod.apply(this, args);
  };
}

// 属性装饰器
function Required(target: any, propertyKey: string) {
  // 属性验证逻辑
}

@Component
class UserService {
  @Required
  private apiUrl: string;

  @Log
  getUser(id: number): User {
    // 获取用户逻辑
    return { id, name: "User", age: 25 };
  }
}
```

## 实用技巧

### 1. 类型守卫

```typescript
// typeof 类型守卫
function processValue(value: string | number) {
  if (typeof value === "string") {
    // 在这里 value 被推断为 string
    return value.toUpperCase();
  }
  // 在这里 value 被推断为 number
  return value.toFixed(2);
}

// instanceof 类型守卫
function handleError(error: Error | string) {
  if (error instanceof Error) {
    console.log(error.message);
  } else {
    console.log(error);
  }
}

// 自定义类型守卫
function isString(value: any): value is string {
  return typeof value === "string";
}
```

### 2. 类型断言

```typescript
// 类型断言
let someValue: unknown = "hello world";
let strLength: number = (someValue as string).length;

// 非空断言
function processUser(user: User | null) {
  // 确定 user 不为 null
  console.log(user!.name);
}
```

### 3. 字面量类型

```typescript
// 字符串字面量类型
type Direction = "up" | "down" | "left" | "right";

// 数字字面量类型
type Dice = 1 | 2 | 3 | 4 | 5 | 6;

// 模板字面量类型
type EventName<T extends string> = `on${Capitalize<T>}`;
type ButtonEvent = EventName<"click">; // "onClick"

type UnitType = "s" | "m" | "h";
type Unit = `${number}${UnitType}`;
```

## String 字符串类型方法

### 基本字符串方法

> ASCII（American Standard Code for Information Interchange，美国信息交换标准代码）是一套字符编码标准，用于将字符转换为数字，以便计算机能够处理和存储文本信息
> 由于标准 ASCII 只能表示 128 个字符，后来出现了扩展 ASCII，使用 8 位编码，可以表示 256 个字符（0-255），增加了更多的特殊符号和非英语字符。
> ASCII 是现代字符编码的基础，UTF-8 等现代编码方案都保持了对 ASCII 的向后兼容性。
> 编码范围：ASCII 使用 7 位二进制数，可以表示 128 个字符（0-127）
> 字符类型：
> 控制字符（0-31）：如换行符、制表符等
> 可打印字符（32-126）：包括空格、数字、字母、标点符号
> DEL 字符（127）：删除字符

```typescript
let str: string = "Hello TypeScript";

// 长度属性
console.log(str.length); // 16

// 字符访问
console.log(str[0]); // "H" console.log(str.charAt(0));
console.log(str.charCodeAt(0)); // 72 (H的ASCII码)
console.log(String.fromCharCode(32)); // -

// 字符串连接
let template: string = `${str} is awesome!`; // 模板字符串 + str.concat(" World!");

// 字符串查找
console.log(str.includes("Script")); // true str.search(/Type/) 6 (正则搜索) str.indexOf("Type") str.lastIndexOf("e")
console.log(str.startsWith("Hello")); // true
console.log(str.endsWith("Script")); // true
```

### 字符串截取与操作

```typescript
let text: string = "JavaScript and TypeScript";

// 截取方法
console.log(text.slice(0, 10)); // "JavaScript" text.substring(15, 25)

// 替换方法
console.log(text.replace("JavaScript", "JS")); // "JS and TypeScript"
console.log(text.replaceAll("Script", "CODE")); // "JavaCODE and TypeCODE"

// 大小写转换
console.log(text.toLowerCase()); // "javascript and typescript"
console.log(text.toUpperCase()); // "JAVASCRIPT AND TYPESCRIPT"

// 去除空格
let spaceStr: string = "  Hello World  ";
console.log(spaceStr.trim()); // "Hello World"
console.log(spaceStr.trimStart()); // "Hello World  "
console.log(spaceStr.trimEnd()); // "  Hello World"

// 填充方法
console.log("5".padStart(3)); // "  5"
console.log("5".padStart(3, "0")); // "005"
console.log("5".padEnd(3, "0")); // "500"

// 重复方法
console.log("x".repeat(3)); // "xxx"
```

### 字符串分割与匹配

```typescript
// 正则表达式创建方式
let regex1: RegExp = /pattern/flags;
let regex2: RegExp = new RegExp("pattern", "flags");

// 常用标志 (flags)
let globalRegex: RegExp = /hello/g; // g: 全局搜索
let caseInsensitive: RegExp = /hello/i; // i: 忽略大小写
let multiline: RegExp = /^hello/m; // m: 多行模式
let singleline: RegExp = /hello./s; // s: . 匹配换行符
let unicode: RegExp = /\u{1F600}/u; // u: Unicode 模式
let sticky: RegExp = /hello/y; // y: 粘性搜索

// 基本正则方法
let text: string = "Hello World! Hello TypeScript!";
let pattern: RegExp = /Hello/g;

// test() - 测试是否匹配
console.log(pattern.test(text)); // true

// exec() - 执行搜索，返回匹配信息
let match = pattern.exec(text);
console.log(match); // ["Hello", index: 0, input: "Hello World! Hello TypeScript!", groups: undefined]

// 字符串的正则方法
console.log(text.match(/Hello/g)); // ["Hello", "Hello"] (所有匹配) text.search(/World/)
console.log(text.matchAll(/Hello/g)); // 迭代器，包含所有匹配的详细信息

string.match(regexp: RegExp | string): RegExpMatchArray | null

// 返回值类型定义
interface RegExpMatchArray extends Array<string> {
  index?: number;        // 匹配开始的位置
  input?: string;        // 原始字符串
  groups?: { [key: string]: string }; // 命名捕获组
}

let text: string = "Hello 123 World 456 TypeScript 789";

// 非全局模式 - 只返回第一个匹配及其详细信息
let nonGlobalPattern: RegExp = /\d+/;
let nonGlobalMatch = text.match(nonGlobalPattern);
console.log(nonGlobalMatch);
// 输出: ["123", index: 6, input: "Hello 123 World 456 TypeScript 789", groups: undefined]

// 全局模式 - 返回所有匹配的数组，但没有详细信息
let globalPattern: RegExp = /\d+/g;
let globalMatch = text.match(globalPattern);
console.log(globalMatch);
// 输出: ["123", "456", "789"]

// 没有匹配时返回 null
let noMatch = text.match(/xyz/);
console.log(noMatch); // null

string.matchAll(regexp: RegExp): IterableIterator<RegExpMatchArray>
// 注意：regexp 必须包含全局标志 'g'，否则会抛出 TypeError

let text: string = "Hello 123 World 456 TypeScript 789";
let pattern: RegExp = /\d+/g; // 必须有全局标志

// 方法1：使用 for...of 迭代
for (const match of text.matchAll(pattern)) {
  console.log(`找到数字: ${match[0]}, 位置: ${match.index}`);
}
// 输出:
// 找到数字: 123, 位置: 6
// 找到数字: 456, 位置: 16
// 找到数字: 789, 位置: 28

// 方法2：转换为数组
let allMatches = Array.from(text.matchAll(pattern));
console.log(allMatches);
// 每个元素都是完整的 RegExpMatchArray

// 方法3：使用展开运算符
let spreadMatches = [...text.matchAll(pattern)];
console.log(spreadMatches);

// replace 和 replaceAll 与正则
console.log(text.replace(/Hello/g, "Hi")); // "Hi World! Hi TypeScript!"
console.log(text.replaceAll(/Hello/g, "Hi")); // 同上，更明确的全局替换

string.replace(searchValue: string | RegExp, replaceValue: string | ((substring: string, ...args: any[]) => string)): string // 注意：非全局正则只替换第一个匹配

let text: string = "John Doe, Jane Smith, Bob Johnson";

// $& - 整个匹配
let bracketNames = text.replace(/\w+ \w+/g, "[$&]");
console.log(bracketNames); // "[John Doe], [Jane Smith], [Bob Johnson]"

// $1, $2, ... - 捕获组
let swapNames = text.replace(/(\w+) (\w+)/g, "$2, $1");
console.log(swapNames); // "Doe, John, Smith, Jane, Johnson, Bob"

// 函数签名
type ReplacerFunction = (
  match: string,           // 完整匹配
  ...groups: string[],     // 捕获组
  offset: number,          // 匹配位置
  string: string,          // 原始字符串
  namedGroups?: { [key: string]: string } // 命名捕获组
) => string;

// split 与正则
let sentence: string = "apple,banana;orange:grape";
console.log(sentence.split(/[,;:]/)); // ["apple", "banana", "orange", "grape"]
console.log(sentence.split(/[,;:]/, 2)); // ["apple", "banana"] limit

// 字符类
let digitPattern: RegExp = /\d/; // 匹配数字 [0-9]
let wordPattern: RegExp = /\w/; // 匹配单词字符 [a-zA-Z0-9_]
let spacePattern: RegExp = /\s/; // 匹配空白字符
let nonDigitPattern: RegExp = /\D/; // 匹配非数字
let nonWordPattern: RegExp = /\W/; // 匹配非单词字符
let nonSpacePattern: RegExp = /\S/; // 匹配非空白字符

// 量词
let oneOrMore: RegExp = /a+/; // 一个或多个 a
let zeroOrMore: RegExp = /a*/; // 零个或多个 a
let zeroOrOne: RegExp = /a?/; // 零个或一个 a
let exactly: RegExp = /a{3}/; // 恰好 3 个 a
let range: RegExp = /a{2,5}/; // 2-5 个 a
let atLeast: RegExp = /a{3,}/; // 至少 3 个 a

// 位置锚点
let startWith: RegExp = /^Hello/; // 以 Hello 开始
let endWith: RegExp = /World$/; // 以 World 结束
let wordBoundary: RegExp = /\bword\b/; // 单词边界
let nonWordBoundary: RegExp = /\Bword\B/; // 非单词边界

// 分组和捕获
let captureGroup: RegExp = /(Hello) (World)/; // 捕获分组
let namedGroup: RegExp = /(?<greeting>Hello) (?<target>World)/; // 命名捕获组
let nonCapture: RegExp = /(?:Hello) World/; // 非捕获分组

// 先行断言和后行断言
let positiveLookahead: RegExp = /Hello(?=World)/; // 正向先行断言
let negativeLookahead: RegExp = /Hello(?!World)/; // 负向先行断言
let positiveLookbehind: RegExp = /(?<=Hello)World/; // 正向后行断言
let negativeLookbehind: RegExp = /(?<!Hello)World/; // 负向后行断言

// 字符集
let customSet: RegExp = /[aeiou]/; // 元音字母
let rangeSet: RegExp = /[a-z]/; // 小写字母
let negativeSet: RegExp = /[^0-9]/; // 非数字
let multipleRanges: RegExp = /[a-zA-Z0-9]/; // 字母数字
```

### 高级字符串方法

```typescript
// 字符串模板与标签
function highlight(strings: TemplateStringsArray, ...values: any[]): string {
  return strings.reduce((result, string, i) => {
    let value = "";

    if (i < values.length) {
      value = `<mark>${values[i]}</mark>`;
    }

    return result + string + value;
  }, "");
}

let name: string = "TypeScript";
let version: string = "4.9";
let result: string = highlight`Welcome to ${name} version ${version}!`;
```

## Number 数字类型方法

### 基本数字方法

```typescript
let num: number = 123.456789;

// 数字转换
console.log(num.toString()); // "123.456789"
console.log(num.toString(2)); // "1111011.01110101..." (二进制)
console.log(num.toString(16)); // "7b.74bc..." (十六进制)

// 精度控制
console.log(num.toFixed());
console.log(num.toFixed(2)); // "123.46" (保留2位小数) num.toPrecision(5) num.toExponential(2)
```

### Number 静态方法

```typescript
// 类型检查
console.log(Number.isFinite(123)); // true
console.log(Number.isFinite(Infinity)); // false
console.log(Number.isInteger(123)); // true
console.log(Number.isInteger(123.45)); // false
console.log(Number.isNaN(NaN)); // true
console.log(Number.isNaN("hello")); // false
console.log(Number.isSafeInteger(123)); // true

// 转换方法
console.log(Number.parseInt("123px")); // 123
console.log(Number.parseInt("ff", 16)); // 255 (十六进制)
console.log(Number.parseFloat("123.45px")); // 123.45

// 常量
console.log(Number.MAX_VALUE); // 最大数值
console.log(Number.MIN_VALUE); // 最小正数值
console.log(Number.MAX_SAFE_INTEGER); // 最大安全整数
console.log(Number.MIN_SAFE_INTEGER); // 最小安全整数
console.log(Number.POSITIVE_INFINITY); // 正无穷
console.log(Number.NEGATIVE_INFINITY); // 负无穷
console.log(Number.NaN); // NaN
```

### Math 对象方法

```typescript
// 基本运算
console.log(Math.abs(-5)); // 5 (绝对值)
console.log(Math.ceil(4.3)); // 5 (向上取整)
console.log(Math.floor(4.7)); // 4 (向下取整)
console.log(Math.round(4.5)); // 5 (四舍五入)
console.log(Math.trunc(4.7)); // 4 (截断小数部分)

// 最值运算
console.log(Math.max(1, 5, 3, 9, 2)); // 9
console.log(Math.min(1, 5, 3, 9, 2)); // 1

// 幂运算
console.log(Math.pow(2, 3)); // 8 (2的3次方)
console.log(Math.sqrt(16)); // 4 (平方根)
console.log(Math.cbrt(27)); // 3 (立方根)
console.log(Math.exp(1)); // e (自然对数的底)
console.log(Math.log(Math.E)); // 1 (自然对数)
console.log(Math.log10(100)); // 2 (以10为底的对数)

// 三角函数
console.log(Math.sin(Math.PI / 2)); // 1
console.log(Math.cos(0)); // 1
console.log(Math.tan(Math.PI / 4)); // 1

// 随机数
console.log(Math.random()); // 0-1之间的随机数
function getRandomInt(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

### bigint
// BigInt 创建方式
let bigInt1: bigint = 123n;                    // 字面量形式
let bigInt2: bigint = BigInt(123);             // 构造函数形式

// 注意：不能对浮点数使用 BigInt()
// let invalid = BigInt(3.14); // TypeError!

// BigInt 的范围
console.log(Number.MAX_SAFE_INTEGER);          // 9007199254740991
console.log(BigInt(Number.MAX_SAFE_INTEGER));  // 9007199254740991n
console.log(BigInt(Number.MAX_SAFE_INTEGER) + 1n); // 9007199254740992n - 超出Number安全范围

// 超大数运算
let largeNumber: bigint = 123456789012345678901234567890n;
console.log(largeNumber * 2n); // 246913578024691357802469135780n

let a: bigint = 10n;
let b: bigint = 3n;

// 基本算术运算
console.log(a + b);     // 13n
console.log(a - b);     // 7n
console.log(a * b);     // 30n
console.log(a / b);     // 3n (整数除法，自动截断)
console.log(a % b);     // 1n (取模)
console.log(a ** b);    // 1000n (幂运算)

// 比较运算
console.log(a > b);     // true
console.log(a === 10n); // true
console.log(a == 10);   // false (类型不同)


function toNumber(value: bigint | number): number {
  if (typeof value === 'bigint') {
    if (value > BigInt(Number.MAX_SAFE_INTEGER) ||
        value < BigInt(Number.MIN_SAFE_INTEGER)) {
      console.warn("BigInt 转换为 Number 可能丢失精度");
    }
    return Number(value);
  }
  return value;
}

function toBigInt(value: number | bigint): bigint {
  if (typeof value === 'number') {
    if (!Number.isInteger(value)) {
      throw new Error("不能将浮点数转换为 BigInt");
    }
    return BigInt(value);
  }
  return value;
}
```

## Boolean 布尔类型方法

### Boolean 基本方法

```typescript
let bool: boolean = true;

// 转换为字符串
console.log(bool.toString()); // "true"
console.log(bool.valueOf()); // true

// Boolean 构造函数
console.log(Boolean(1)); // true
console.log(Boolean(0)); // false
console.log(Boolean("")); // false
console.log(Boolean("hello")); // true
console.log(Boolean(null)); // false
console.log(Boolean(undefined)); // false
console.log(Boolean([])); // true
console.log(Boolean({})); // true

// 逻辑运算
let a: boolean = true;
let b: boolean = false;
console.log(a && b); // false (逻辑与)
console.log(a || b); // true (逻辑或)
console.log(!a); // false (逻辑非)
```

### 布尔值的实用技巧

```typescript
// 双重否定转换
let value: any = "hello";
let isTruthy: boolean = !!value; // true

// 短路运算
let name: string = "" || "默认名称"; // "默认名称"
let user = { name: "Alice" };
user && console.log(user.name); // 只有user存在时才执行

// 空值合并运算符
let config = {
  timeout: 0,
  retries: null,
};
let timeout = config.timeout ?? 5000; // 0 (因为0不是null/undefined)
let retries = config.retries ?? 3; // 3 (因为null)

// 可选链
let response: any = { data: { users: [{ name: "Alice" }] } };
let userName = response?.data?.users?.[0]?.name; // "Alice"
```

## Symbol 的基本概念

> Symbol 是 ES6 引入的一种新的原始数据类型，每个 Symbol 值都是唯一的，主要用于创建对象的私有属性或避免属性名冲突

```typescript
// 创建 Symbol
const sym1 = Symbol();
const sym2 = Symbol();

console.log(sym1 === sym2); // false，每个 Symbol 都是唯一的

// 带描述的 Symbol
const sym3 = Symbol("description");
console.log(sym3.toString()); // Symbol(description)

// 唯一性
const a = Symbol("key");
const b = Symbol("key");
console.log(a === b); // false

// 不可枚举
const obj = {
  name: "John",
  [Symbol("id")]: 123,
};

console.log(Object.keys(obj)); // ['name']
console.log(Object.getOwnPropertyNames(obj)); // ['name']
console.log(Object.getOwnPropertySymbols(obj)); // [Symbol(id)]

// 基本 Symbol 类型
let sym: symbol = Symbol();

// 具体的 Symbol 类型
const uniqueId: unique symbol = Symbol("id");

// Symbol 作为对象键的类型
interface MyObject {
  [key: symbol]: any;
  name: string;
}

// 实现迭代器
class NumberRange {
  constructor(private start: number, private end: number) {}

  *[Symbol.iterator]() {
    for (let i = this.start; i <= this.end; i++) {
      yield i;
    }
  }
}

const range = new NumberRange(1, 5);
for (const num of range) {
  console.log(num); // 1, 2, 3, 4, 5
}

// 自定义 toString 行为
class CustomClass {
  constructor(private value: string) {}

  [Symbol.toPrimitive](hint: string) {
    if (hint === "string") {
      return `CustomClass: ${this.value}`;
    }
    if (hint === "number") {
      return this.value.length;
    }
    return this.value;
  }
}

const instance = new CustomClass("hello");
console.log(String(instance)); // "CustomClass: hello"
console.log(Number(instance)); // 5

// 创建全局 Symbol
const globalSym = Symbol.for("global-key");
const sameSym = Symbol.for("global-key");

console.log(globalSym === sameSym); // true

// 获取全局 Symbol 的 key
console.log(Symbol.keyFor(globalSym)); // 'global-key'
```

## Array 数组类型方法

### 数组基本操作

```typescript
let numbers: number[] = [1, 2, 3, 4, 5];
let fruits: string[] = ["apple", "banana", "orange"];

// 长度属性
console.log(numbers.length); // 5

// 添加元素
numbers.push(6); // 末尾添加，返回新长度 push(7, 8, 9)
numbers.unshift(0); // 开头添加，返回新长度 unshift(1, 2, 3)
console.log(numbers); // [0, 1, 2, 3, 4, 5, 6]

// 删除元素
let lastItem = numbers.pop(); // 删除并返回最后一个元素
let firstItem = numbers.shift(); // 删除并返回第一个元素
console.log(numbers); // [1, 2, 3, 4, 5]

// 截取和插入
let removed = numbers.splice(2, 1, 10, 11); // 从索引2删除1个元素，插入10,11
console.log(numbers); // [1, 2, 10, 11, 4, 5]
console.log(removed); // [3]
```

### 数组查找与判断

```typescript
let arr: number[] = [1, 2, 3, 4, 5, 3];

// 查找方法
console.log(arr.includes(4)); // true arr.indexOf(3) arr.lastIndexOf(3)

// 高级查找
let found = arr.find((x) => x > 3); // 4 (第一个满足条件的元素) arr.findIndex((x) => x > 3);

// 判断方法
let allPositive = arr.every((x) => x > 0); // true (是否所有元素都满足条件) [] is true
let hasEven = arr.some((x) => x % 2 === 0); // true (是否存在满足条件的元素) [] is false
```

### 数组转换方法

```typescript
let numbers: number[] = [1, 2, 3, 4, 5];

// map - 转换每个元素
let doubled: number[] = numbers.map((x) => x * 2); // [2, 4, 6, 8, 10]
let strings: string[] = numbers.map((x) => x.toString()); // ["1", "2", "3", "4", "5"]

// filter - 过滤元素
let evens: number[] = numbers.filter((x) => x % 2 === 0); // [2, 4]
let odds: number[] = numbers.filter((x) => x % 2 === 1); // [1, 3, 5]

// reduce - 归约
let sum: number = numbers.reduce((acc, cur) => acc + cur, 0); // 15
let product: number = numbers.reduce((acc, cur) => acc * cur, 1); // 120

// 带索引的reduce
let indexed = numbers.reduce((acc: { [key: number]: number }, cur, index) => {
  acc[index] = cur * cur;
  return acc;
}, {}); // {0: 1, 1: 4, 2: 9, 3: 16, 4: 25}

// 扁平化
let deepNested: number[][][] = [[[1, 2]], [[3, 4]]];
console.log(deepNested.flat()); // [[1, 2], [3, 4]] default as 1
console.log(deepNested.flat(2)); // [1, 2, 3, 4]
console.log(deepNested.flat(Infinity)); // 完全扁平化

// flatMap - 扁平化映射 map + flat
let nested: number[][] = [[1, 2], [3, 4], [5]];
let flattened: number[] = nested.flatMap((x) => x); // [1, 2, 3, 4, 5]

// 填充数组
let zeros: number[] = new Array(5).fill(0); // [0, 0, 0, 0, 0]
let sequence: number[] = Array.from({ length: 5 }, (_, i) => i + 1); // [1, 2, 3, 4, 5]
```

### 数组排序与反转

```typescript
let arr: number[] = [3, 1, 4, 1, 5, 9, 2, 6];
let fruits: string[] = ["banana", "apple", "orange", "grape"];

// 排序 (会修改原数组)
arr.sort(); // 默认字符串排序: [1, 1, 2, 3, 4, 5, 6, 9]
arr.sort((a, b) => a - b); // 数字升序: [1, 1, 2, 3, 4, 5, 6, 9]
arr.sort((a, b) => b - a); // 数字降序: [9, 6, 5, 4, 3, 2, 1, 1]

fruits.sort(); // 字母排序: ["apple", "banana", "grape", "orange"]
fruits.sort((a, b) => a.localeCompare(b)); // 本地化排序

// 反转 (会修改原数组)
arr.reverse(); // [1, 1, 2, 3, 4, 5, 6, 9]
```

### 数组连接与截取

```typescript
let arr1: number[] = [1, 2, 3];
let arr2: number[] = [4, 5, 6];

// 连接数组 (不修改原数组)
let combined: number[] = arr1.concat(arr2); // [1, 2, 3, 4, 5, 6]
let spread: number[] = [...arr1, ...arr2]; // 使用展开运算符

// 截取数组 (不修改原数组)
let sliced: number[] = combined.slice(1, 4); // [2, 3, 4]
let copy: number[] = combined.slice(); // 复制整个数组

// 转换为字符串
console.log(arr1.join()); // "1,2,3"
console.log(arr1.join(" - ")); // "1 - 2 - 3"
console.log(arr1.toString()); // "1,2,3"
```

### 现代数组方法

```typescript
let arr: number[] = [1, 2, 3, 4, 5];

// 数组迭代器
for (let value of arr) {
  console.log(value); // 1, 2, 3, 4, 5
}

for (let value of arr.keys()) {
  console.log(value); // 1, 2, 3, 4, 5
}

for (let value of arr.values()) {
  console.log(value); // 1, 2, 3, 4, 5
}

for (let [index, value] of arr.entries()) {
  console.log(index, value); // 0 1, 1 2, 2 3, 3 4, 4 5
}

// 数组方法链式调用
let result: number[] = arr
  .filter((x) => x % 2 === 1) // 过滤奇数
  .map((x) => x * x) // 平方
  .sort((a, b) => b - a); // 降序排列
console.log(result); // [25, 9, 1]
```

## Object 对象类型方法

### Object 静态方法

```typescript
let person = {
  name: "Alice",
  age: 30,
  city: "New York",
};

// 获取属性
console.log(Object.keys(person)); // ["name", "age", "city"]
console.log(Object.values(person)); // ["Alice", 30, "New York"]
console.log(Object.entries(person)); // [["name", "Alice"], ["age", 30], ["city", "New York"]]

// Object.keys() 只返回自有属性（不包含原型链）可枚举属性
// for in 遍历自有和继承的可枚举属性

// 键值对数组转换为对象
const entries = [
  ["name", "Bob"],
  ["age", 25],
  ["city", "Beijing"],
];
const objFromEntries = Object.fromEntries(entries);
console.log(objFromEntries); // { name: "Bob", age: 25, city: "Beijing" }

const map = new Map([
  ["x", 100],
  ["y", 200],
]);
const objFromMap = Object.fromEntries(map);
console.log(objFromMap); // { x: 100, y: 200 }

// 与 Object.entries 结合，实现对象的映射转换
const original = { a: 1, b: 2, c: 3 };
const mapped = Object.fromEntries(
  Object.entries(original).map(([key, value]) => [key, value * 10])
);
console.log(mapped); // { a: 10, b: 20, c: 30 }

const user = { name: "Alice", age: 30, password: "secret" };
const filtered = Object.fromEntries(
  Object.entries(user).filter(([key]) => key !== "password")
);
console.log(filtered); // { name: "Alice", age: 30 }

// 对象操作
let copy = Object.assign({}, person); // 浅拷贝
let merged = Object.assign({}, person, { country: "USA" }); // 合并对象

// 属性描述符
Object.defineProperty(person, "id", {
  value: 123,
  writable: false,
  enumerable: false,
  configurable: false,
});

// 冻结和密封
Object.freeze(person); // 冻结对象，不可修改
Object.seal(person); // 密封对象，不可添加/删除属性
console.log(Object.isFrozen(person)); // true
console.log(Object.isSealed(person)); // true
```

### 对象原型方法

```typescript
let obj = { name: "Alice" };

// 基本原型方法
console.log(obj.toString()); // "[object Object]"
console.log(obj.valueOf()); // { name: "Alice" }
console.log(obj.hasOwnProperty("name")); // true
console.log(obj.propertyIsEnumerable("name")); // true

// 原型链检查
console.log(obj.isPrototypeOf({})); // false
console.log(Object.prototype.isPrototypeOf(obj)); // true

// 获取和设置原型
console.log(Object.getPrototypeOf(obj)); // Object.prototype
let newProto = { type: "person" };
Object.setPrototypeOf(obj, newProto);
```

### 现代对象方法

```typescript
// 对象解构
let { name, age, city = "Unknown" } = person;

// 属性简写
let createPerson = (name: string, age: number) => ({ name, age });

// 计算属性名
let propertyName = "dynamicProp";
let dynamicObj = {
  [propertyName]: "value",
  [`${propertyName}2`]: "value2",
};

// 对象展开
let extended = { ...person, country: "USA", age: 31 };

// 可选链
let user: any = { profile: { settings: { theme: "dark" } } };
let theme = user?.profile?.settings?.theme; // "dark"

// 空值合并
let config = { timeout: null };
let timeout = config.timeout ?? 5000; // 5000
```

## Set 集合类型详解

### Set 基础操作

```typescript
// Set 创建方式
let set1: Set<number> = new Set(); // 空 Set
let set2: Set<string> = new Set(["a", "b", "c"]); // 从数组创建
let set3: Set<number> = new Set([1, 2, 2, 3, 3]); // 自动去重: {1, 2, 3}

// 基本属性和方法
console.log(set2.size); // 3 (Set 大小)

// 添加元素
set1.add(1);
set1.add(2);
set1.add(2); // 重复添加无效
console.log(set1); // Set {1, 2}

// 链式添加
set1.add(3).add(4).add(5);
console.log(set1); // Set {1, 2, 3, 4, 5}

// 检查元素存在
console.log(set1.has(3)); // true
console.log(set1.has(10)); // false

// 删除元素
console.log(set1.delete(3)); // true (删除成功)
console.log(set1.delete(10)); // false (元素不存在)
console.log(set1); // Set {1, 2, 4, 5}

// 清空 Set
set1.clear();
console.log(set1.size); // 0
```

### Set 遍历方法

```typescript
let colors = new Set(["red", "green", "blue", "red"]); // {red, green, blue}

// forEach 遍历
console.log("\n=== forEach 遍历 ===");
colors.forEach((value, key, set) => {
  console.log(`值: ${value}, 键: ${key}, 相同: ${value === key}`);
  // 注意：Set 中 key 和 value 相同
});

// for...of 遍历值
console.log("=== for...of 遍历 ===");
for (const color of colors) {
  console.log(color);
}

for (const [key, value] of colors.entries()) {
  console.log(`[${key}, ${value}]`);
}

// 使用迭代器
console.log("\n=== 迭代器方法 ===");
console.log("values():", Array.from(colors.values())); // ["red", "green", "blue"]
console.log("keys():", Array.from(colors.keys())); // ["red", "green", "blue"] (与values相同)
console.log("entries():", Array.from(colors.entries())); // [ [ 'red', 'red' ], [ 'green', 'green' ], [ 'blue', 'blue' ] ]

// 转换为数组
let colorArray = [...colors]; // 使用展开运算符
let colorArray2 = Array.from(colors); // 使用 Array.from
console.log("转为数组:", colorArray);
```

## Map 映射类型详解

### Map 基础操作

```typescript
// Map 创建方式
let map1: Map<string, number> = new Map(); // 空 Map
let map2: Map<string, number> = new Map([
  ["apple", 5],
  ["banana", 3],
  ["orange", 8],
]); // 从数组创建

// 从对象创建 Map
let obj = { a: 1, b: 2, c: 3 };
let map3: Map<string, number> = new Map(Object.entries(obj));

// 基本属性和方法
console.log(map2.size); // 3

// 设置键值对
map1.set("first", 1);
map1.set("second", 2);
map1.set("third", 3);

// 链式设置
map1.set("fourth", 4).set("fifth", 5);

// 获取值
console.log(map1.get("first")); // 1
console.log(map1.get("nonexistent")); // undefined

// 检查键是否存在
console.log(map1.has("second")); // true
console.log(map1.has("sixth")); // false

// 删除键值对
console.log(map1.delete("third")); // true (删除成功)
console.log(map1.delete("nonexistent")); // false (键不存在)

// 清空 Map
map1.clear();
console.log(map1.size); // 0
```

### Map 遍历方法

```typescript
let fruits = new Map([
  ["apple", 5],
  ["banana", 3],
  ["orange", 8],
  ["grape", 12],
]);

// forEach 遍历
console.log("\n=== forEach 遍历 ===");
fruits.forEach((value, key, map) => {
  console.log(`水果: ${key}, 数量: ${value}`);
});

// for...of 遍历键值对
console.log("=== for...of 遍历键值对 ===");
for (const [key, value] of fruits) {
  console.log(`${key}: ${value}`);
}

console.log("\n=== 遍历键 ===");
for (const key of fruits.keys()) {
  console.log(key);
}

console.log("\n=== 遍历值 ===");
for (const value of fruits.values()) {
  console.log(value);
}

// 方使用迭代器
console.log("\n=== 迭代器方法 ===");
console.log("所有键:", Array.from(fruits.keys()));
console.log("所有值:", Array.from(fruits.values()));
console.log("所有条目:", Array.from(fruits.entries()));

// 转换操作
console.log("\n=== 转换操作 ===");
let fruitsArray = [...fruits]; // 转为数组
let fruitsObject = Object.fromEntries(fruits); // 转为对象
console.log("转为数组:", fruitsArray);
console.log("转为对象:", fruitsObject);
```

## JSON 对象详解

### JSON.stringify() 详解

```typescript
// 基本用法
let obj = {
  name: "Alice",
  age: 30,
  hobbies: ["reading", "swimming"],
  address: {
    city: "New York",
    country: "USA",
  },
};

console.log("=== JSON.stringify() 基本用法 ===");
console.log("基本序列化:", JSON.stringify(obj));
console.log("格式化输出:", JSON.stringify(obj, null, 2));
console.log("紧凑格式:", JSON.stringify(obj, null, 0));

// 高级参数用法
class JSONStringifyAdvanced {
  static demonstrateReplacer() {
    console.log("\n=== replacer 参数详解 ===");

    const data = {
      id: 123,
      name: "Product",
      price: 99.99,
      secret: "confidential",
      internal: "private",
      description: "A great product",
    };

    // 1. 数组形式的 replacer - 只序列化指定字段
    const publicFields = ["id", "name", "price", "description"];
    console.log("只序列化公开字段:");
    console.log(JSON.stringify(data, publicFields, 2));

    // 2. 函数形式的 replacer - 自定义过滤逻辑
    const filterReplacer = (key: string, value: any) => {
      // 过滤敏感字段
      if (key.includes("secret") || key.includes("internal")) {
        return undefined;
      }
      // 格式化价格
      if (key === "price" && typeof value === "number") {
        return `${value.toFixed(2)}`;
      }
      return value;
    };

    console.log("使用函数过滤:");
    console.log(JSON.stringify(data, filterReplacer, 2));

    // 3. 类型转换
    const typeConverter = (key: string, value: any) => {
      if (typeof value === "number") {
        return value.toString(); // 数字转字符串
      }
      if (value instanceof Date) {
        return value.toISOString(); // 日期转ISO字符串
      }
      return value;
    };

    const dataWithDate = { ...data, createdAt: new Date() };
    console.log("类型转换:");
    console.log(JSON.stringify(dataWithDate, typeConverter, 2));
  }

  static demonstrateSpace() {
    console.log("\n=== space 参数详解 ===");

    const data = { a: 1, b: { c: 2, d: 3 } };

    // 数字：缩进空格数
    console.log("2个空格缩进:");
    console.log(JSON.stringify(data, null, 2));

    console.log("4个空格缩进:");
    console.log(JSON.stringify(data, null, 4));

    // 字符串：使用指定字符缩进
    console.log("使用制表符缩进:");
    console.log(JSON.stringify(data, null, "\t"));

    console.log("自定义缩进字符:");
    console.log(JSON.stringify(data, null, "-> "));
  }

  static handleSpecialValues() {
    console.log("\n=== 特殊值处理 ===");

    const specialData = {
      number: 42,
      string: "hello",
      boolean: true,
      nullValue: null,
      undefinedValue: undefined,
      symbol: Symbol("sym"),
      functionValue: function () {
        return "test";
      },
      date: new Date(),
      regex: /test/g,
      infinity: Infinity,
      nan: NaN,
    };

    console.log("原始对象keys:", Object.keys(specialData));
    console.log("JSON序列化结果:");
    console.log(JSON.stringify(specialData, null, 2));

    console.log("\n特殊值处理规则:");
    console.log("- undefined: 被忽略");
    console.log("- Symbol: 被忽略");
    console.log("- Function: 被忽略");
    console.log("- Date: 转为 ISO 字符串");
    console.log("- RegExp: 转为空对象 {}");
    console.log("- Infinity/NaN: 转为 null");
  }
}

// 执行示例
JSONStringifyAdvanced.demonstrateReplacer();
JSONStringifyAdvanced.demonstrateSpace();
JSONStringifyAdvanced.handleSpecialValues();
```

### JSON.parse() 详解

```typescript
class JSONParseAdvanced {
  static demonstrateBasicParsing() {
    console.log("=== JSON.parse() 基本用法 ===");

    const jsonString =
      '{"name":"Alice","age":30,"hobbies":["reading","swimming"]}';
    const parsed = JSON.parse(jsonString);

    console.log("原始JSON字符串:", jsonString);
    console.log("解析结果:", parsed);
    console.log("类型:", typeof parsed);
  }

  static demonstrateReviver() {
    console.log("\n=== reviver 参数详解 ===");

    const jsonData = `{
      "id": "123",
      "name": "Product",
      "price": "99.99",
      "createdAt": "2023-12-25T10:30:00.000Z",
      "tags": "tag1,tag2,tag3",
      "isActive": "true"
    }`;

    // 1. 类型转换 reviver
    const typeReviver = (key: string, value: any) => {
      // 转换日期字符串
      if (key.includes("At") && typeof value === "string") {
        return new Date(value);
      }
      // 转换数字字符串
      if (key === "price" && typeof value === "string") {
        return parseFloat(value);
      }
      // 转换布尔字符串
      if (key === "isActive" && typeof value === "string") {
        return value === "true";
      }
      // 转换逗号分隔的字符串为数组
      if (key === "tags" && typeof value === "string") {
        return value.split(",");
      }
      return value;
    };

    console.log("使用 reviver 转换类型:");
    const parsedWithReviver = JSON.parse(jsonData, typeReviver);
    console.log(parsedWithReviver);
    console.log("createdAt 类型:", typeof parsedWithReviver.createdAt);
    console.log("price 类型:", typeof parsedWithReviver.price);

    // 2. 数据验证 reviver
    const validationReviver = (key: string, value: any) => {
      // 验证必需字段
      if (key === "name" && (!value || value.length === 0)) {
        throw new Error("Name is required");
      }
      // 验证数值范围
      if (key === "price" && typeof value === "number" && value < 0) {
        throw new Error("Price cannot be negative");
      }
      return value;
    };

    try {
      const validData = '{"name":"Product","price":99.99}';
      console.log("有效数据解析:", JSON.parse(validData, validationReviver));

      const invalidData = '{"name":"","price":-10}';
      JSON.parse(invalidData, validationReviver);
    } catch (error) {
      console.log("验证错误:", error.message);
    }
  }

  static handleParsingErrors() {
    console.log("\n=== 解析错误处理 ===");

    const invalidJSONs = [
      '{"name": "Alice", "age": 30,}', // 尾随逗号
      '{"name": "Alice", "age": undefined}', // undefined 值
      "{'name': 'Alice'}", // 单引号
      '{"name": "Alice" "age": 30}', // 缺少逗号
      '{"name": "Alice", "age": 30', // 不完整
    ];

    invalidJSONs.forEach((json, index) => {
      try {
        JSON.parse(json);
        console.log(`JSON ${index + 1}: 解析成功`);
      } catch (error) {
        console.log(`JSON ${index + 1}: ${error.message}`);
      }
    });

    // 安全解析函数
    const safeJSONParse = <T>(jsonString: string, defaultValue: T): T => {
      try {
        return JSON.parse(jsonString);
      } catch (error) {
        console.warn("JSON 解析失败，返回默认值:", error.message);
        return defaultValue;
      }
    };

    console.log("安全解析示例:");
    const result = safeJSONParse("invalid json", { error: "default" });
    console.log("结果:", result);
  }
}

// 执行示例
JSONParseAdvanced.demonstrateBasicParsing();
JSONParseAdvanced.demonstrateReviver();
JSONParseAdvanced.handleParsingErrors();
```

## 1. Date 对象完整 API

### 1.1 Date 构造函数

> 时间戳是毫秒级时间戳，秒级时间戳 \* 1000

```typescript
// 构造函数重载
new Date(): Date;                              // 当前时间 new Date()
new Date(value: number): Date;                 // 时间戳 new Date(1640995200000)
new Date(dateString: string): Date;            // 日期字符串 new Date('2024-01-01T00:00:00Z') ne Date(2020-10-28 12:00:00)
new Date(year: number, monthIndex: number, day?: number,
         hours?: number, minutes?: number, seconds?: number,
         milliseconds?: number): Date;         // 组件构造 new Date(2024, 0, 1, 12, 30, 45, 500)
```

### 1.2 静态方法

```typescript
class DateStaticMethods {
  static demonstrateStaticMethods() {
    // Date.now() - 返回当前时间戳
    const timestamp: number = Date.now();

    // Date.parse() - 返回当前时间戳
    const parsed: number = Date.parse("2024-01-01T00:00:00Z");
  }
}
```

### 1.3 Getter 方法（获取日期组件）

```typescript
class DateGetters {
  private date: Date = new Date("2024-03-15T14:30:45.123Z");

  demonstrateGetters() {
    // 年份
    const fullYear: number = this.date.getFullYear(); // 2024

    // 月份（0-11）
    const month: number = this.date.getMonth(); // 2 (March)

    // 日期（1-31）
    const dateNum: number = this.date.getDate(); // 15

    // 星期几（0-6，0=Sunday）
    const day: number = this.date.getDay(); // 5 (Friday)

    // 小时（0-23）
    const hours: number = this.date.getHours(); // 14

    // 分钟（0-59）
    const minutes: number = this.date.getMinutes(); // 30

    // 秒（0-59）
    const seconds: number = this.date.getSeconds(); // 45

    // 毫秒（0-999）
    const milliseconds: number = this.date.getMilliseconds(); // 123

    // 时间戳
    const time: number = this.date.getTime(); // 1710509445123

    // 时区偏移（分钟）
    const timezoneOffset: number = this.date.getTimezoneOffset();
  }
}
```

### 1.4 Setter 方法（设置日期组件）

```typescript
class DateSetters {
  private date: Date = new Date();

  demonstrateSetters() {
    // 设置年份
    this.date.setFullYear(2024);
    this.date.setFullYear(2024, 11); // 年份和月份
    this.date.setFullYear(2024, 11, 25); // 年份、月份、日期

    // 设置月份（0-11）
    this.date.setMonth(5); // June
    this.date.setMonth(5, 15); // June 15th

    // 设置日期（1-31）
    this.date.setDate(20);

    // 设置小时（0-23）
    this.date.setHours(15);
    this.date.setHours(15, 30); // 小时和分钟
    this.date.setHours(15, 30, 45); // 小时、分钟、秒
    this.date.setHours(15, 30, 45, 500); // 小时、分钟、秒、毫秒

    // 设置分钟（0-59）
    this.date.setMinutes(45);
    this.date.setMinutes(45, 30); // 分钟和秒
    this.date.setMinutes(45, 30, 250); // 分钟、秒、毫秒

    // 设置秒（0-59）
    this.date.setSeconds(30);
    this.date.setSeconds(30, 750); // 秒和毫秒

    // 设置毫秒（0-999）
    this.date.setMilliseconds(123);

    // 设置时间戳
    this.date.setTime(1640995200000);

    console.log("Modified date:", this.date);
  }
}
```

### 1.5 格式化方法

```typescript
class DateFormatting {
  private date: Date = new Date("2024-03-15T14:30:45.123Z");

  demonstrateFormatting() {
    // 基础字符串方法
    const toString: string = this.date.toString();
    // "Fri Mar 15 2024 15:30:45 GMT+0100 (CET)"

    const toDateString: string = this.date.toDateString();
    // "Fri Mar 15 2024"

    const toTimeString: string = this.date.toTimeString();
    // "15:30:45 GMT+0100 (CET)"

    // ISO 格式
    const toISOString: string = this.date.toISOString();
    // "2024-03-15T14:30:45.123Z"

    // JSON 格式（同 toISOString）
    const toJSON: string = this.date.toJSON();
    // "2024-03-15T14:30:45.123Z"
  }

  // 本地化格式化
  demonstrateLocalization() {
    // toLocaleString 系列
    const localeString: string = this.date.toLocaleString();
    const localeString_US: string = this.date.toLocaleString("en-US");
    const localeString_CN: string = this.date.toLocaleString("zh-CN");

    const localeDateString: string = this.date.toLocaleDateString();
    const localeDateString_US: string = this.date.toLocaleDateString("en-US");

    const localeTimeString: string = this.date.toLocaleTimeString();
    const localeTimeString_US: string = this.date.toLocaleTimeString("en-US");

    // 详细格式化选项
    const detailedFormat: string = this.date.toLocaleString("zh-CN", {
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      timeZoneName: "short",
    });

    console.log({
      localeString,
      localeString_US,
      localeString_CN,
      localeDateString,
      localeDateString_US,
      localeTimeString,
      localeTimeString_US,
      detailedFormat,
    });
  }
}
```

### 1.6 其他实用方法

```typescript
class DateUtilities {
  static sub(): number {
    return new Date() - 10 * 60 * 100;
  }

  static isBefore() {
    return new Date() > new Date("2024-10-28");
  }

  static valueOf(date: Date): number {
    return date.valueOf(); // 返回时间戳，等同于 getTime()
  }

  static getWeekNumber(date: Date): number {
    const startOfYear = new Date(date.getFullYear(), 0, 1);
    const pastDaysOfYear = (date.getTime() - startOfYear.getTime()) / 86400000;
    return Math.ceil((pastDaysOfYear + startOfYear.getDay() + 1) / 7);
  }

  static isLeapYear(year: number): boolean {
    return (year % 4 === 0 && year % 100 !== 0) || year % 400 === 0;
  }

  static getDaysInMonth(year: number, month: number): number {
    return new Date(year, month + 1, 0).getDate();
  }

  static addTime(
    date: Date,
    options: {
      years?: number;
      months?: number;
      days?: number;
      hours?: number;
      minutes?: number;
      seconds?: number;
      milliseconds?: number;
    }
  ): Date {
    const result = new Date(date);

    if (options.years) result.setFullYear(result.getFullYear() + options.years);
    if (options.months) result.setMonth(result.getMonth() + options.months);
    if (options.days) result.setDate(result.getDate() + options.days);
    if (options.hours) result.setHours(result.getHours() + options.hours);
    if (options.minutes)
      result.setMinutes(result.getMinutes() + options.minutes);
    if (options.seconds)
      result.setSeconds(result.getSeconds() + options.seconds);
    if (options.milliseconds)
      result.setMilliseconds(result.getMilliseconds() + options.milliseconds);

    return result;
  }

  static formatDuration(milliseconds: number): string {
    const seconds = Math.floor(milliseconds / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);

    if (days > 0) return `${days}d ${hours % 24}h ${minutes % 60}m`;
    if (hours > 0) return `${hours}h ${minutes % 60}m ${seconds % 60}s`;
    if (minutes > 0) return `${minutes}m ${seconds % 60}s`;
    return `${seconds}s`;
  }
}
```

## TypeScript Promise 详解

Promise 是 JavaScript/TypeScript 中处理异步操作的对象，它表示一个异步操作的最终完成或失败，以及其结果值。

### Promise 的三种状态

- **Pending（待定）**: 初始状态，既没有被兑现，也没有被拒绝
- **Fulfilled（已兑现）**: 操作成功完成
- **Rejected（已拒绝）**: 操作失败

## 2. TypeScript 中的 Promise 类型

```typescript
// Promise 是一个泛型类型，T 表示成功时返回的数据类型
Promise<T>;

// 示例
const stringPromise: Promise<string> = Promise.resolve("Hello");
const numberPromise: Promise<number> = Promise.resolve(42);
const userPromise: Promise<User> = fetchUser();

interface User {
  id: number;
  name: string;
  email: string;
}
```

## 3. 创建 Promise

### 3.1 使用 Promise 构造函数

```typescript
const myPromise = new Promise<string>((resolve, reject) => {
  const success = Math.random() > 0.5;

  setTimeout(() => {
    if (success) {
      resolve("操作成功！");
    } else {
      reject(new Error("操作失败！"));
    }
  }, 1000);
});
```

### 3.2 使用 Promise.resolve() 和 Promise.reject()

```typescript
// 创建已解决的 Promise
const resolvedPromise: Promise<number> = Promise.resolve(42);

// 创建已拒绝的 Promise
const rejectedPromise: Promise<never> = Promise.reject(new Error("错误信息"));
```

## 4. Promise 方法

### 4.1 then() 方法

```typescript
const promise: Promise<string> = fetchData();

promise
  .then((data: string) => {
    console.log("成功:", data);
    return data.toUpperCase(); // 返回新值
  })
  .then((upperData: string) => {
    console.log("转换后:", upperData);
  });
```

### 4.2 catch() 方法

```typescript
promise
  .then((data: string) => {
    return processData(data);
  })
  .catch((error: Error) => {
    console.error("错误:", error.message);
    return "默认值"; // 返回默认值或重新抛出错误
  });
```

### 4.3 finally() 方法

```typescript
promise
  .then((data: string) => {
    return processData(data);
  })
  .catch((error: Error) => {
    console.error("错误:", error);
  })
  .finally(() => {
    console.log("清理工作"); // 无论成功失败都会执行
  });
```

## 5. async/await 语法

### 5.1 基本用法

```typescript
async function fetchUserData(id: number): Promise<User> {
  try {
    const response = await fetch(`/api/users/${id}`);
    const user: User = await response.json();
    return user;
  } catch (error) {
    console.error("获取用户数据失败:", error);
    throw error;
  }
}

// 使用
async function handleUser() {
  try {
    const user = await fetchUserData(1);
    console.log("用户:", user.name);
  } catch (error) {
    console.error("处理用户数据失败:", error);
  }
}
```

### 5.2 函数返回类型

```typescript
// async 函数总是返回 Promise
async function getString(): Promise<string> {
  return "Hello"; // 等同于 Promise.resolve("Hello")
}

async function getNumber(): Promise<number> {
  return 42;
}

// 如果函数可能抛出错误，类型仍然是 Promise<T>
async function riskyOperation(): Promise<string> {
  if (Math.random() > 0.5) {
    throw new Error("随机错误");
  }
  return "成功";
}
```

## 6. Promise 工具方法

### 6.1 Promise.all()

```typescript
// 并行执行多个 Promise，全部成功才成功，适合依赖关系
async function fetchAllData(): Promise<[User[], Post[], Comment[]]> {
  const [users, posts, comments] = await Promise.all([
    fetchUsers(),
    fetchPosts(),
    fetchComments(),
  ]);

  return [users, posts, comments];
}

// 类型推导示例
const promises = [
  Promise.resolve("string"),
  Promise.resolve(42),
  Promise.resolve(true),
] as const;

// result 类型为 [string, number, boolean]
const result = await Promise.all(promises);
```

### 6.2 Promise.allSettled()

```typescript
// 返回所有结果，适合独立操作
interface SettledResult<T> {
  status: "fulfilled" | "rejected";
  value?: T;
  reason?: any;
}

async function fetchDataSafely() {
  const results = await Promise.allSettled([
    fetchUsers(),
    fetchPosts(),
    fetchComments(),
  ]);

  results.forEach((result, index) => {
    if (result.status === "fulfilled") {
      console.log(`Promise ${index} 成功:`, result.value);
    } else {
      console.log(`Promise ${index} 失败:`, result.reason);
    }
  });
}
```

### 6.3 Promise.race()

```typescript
// 返回最先完成的 Promise 结果
async function fetchWithTimeout<T>(
  promise: Promise<T>,
  timeout: number
): Promise<T> {
  const timeoutPromise = new Promise<never>((_, reject) => {
    setTimeout(() => reject(new Error("超时")), timeout);
  });

  return Promise.race([promise, timeoutPromise]);
}

// 使用
try {
  const user = await fetchWithTimeout(fetchUser(1), 5000);
  console.log("用户:", user);
} catch (error) {
  console.error("获取用户超时或失败:", error);
}
```

### 6.4 Promise.any()

```typescript
// 返回第一个成功的 Promise，所有都失败才失败
async function fetchFromMultipleSources(): Promise<string> {
  try {
    const data = await Promise.any([
      fetch("/api/source1").then((r) => r.text()),
      fetch("/api/source2").then((r) => r.text()),
      fetch("/api/source3").then((r) => r.text()),
    ]);
    return data;
  } catch (aggregateError) {
    // 所有源都失败
    console.error("所有数据源都失败:", aggregateError);
    throw new Error("无法获取数据");
  }
}
```

## 7. 错误处理最佳实践

### 7.1 自定义错误类型

```typescript
class APIError extends Error {
  constructor(message: string, public status: number, public code: string) {
    super(message);
    this.name = "APIError";
  }
}

async function fetchAPI<T>(url: string): Promise<T> {
  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new APIError(
        `请求失败: ${response.statusText}`,
        response.status,
        "FETCH_ERROR"
      );
    }

    return await response.json();
  } catch (error) {
    if (error instanceof APIError) {
      throw error;
    }
    throw new APIError("网络错误", 0, "NETWORK_ERROR");
  }
}
```

### 7.2 错误类型守卫

```typescript
function isAPIError(error: unknown): error is APIError {
  return error instanceof APIError;
}

async function handleDataFetch() {
  try {
    const data = await fetchAPI<User[]>("/api/users");
    return data;
  } catch (error) {
    if (isAPIError(error)) {
      console.error(`API 错误 [${error.code}]:`, error.message);
      if (error.status === 404) {
        return []; // 返回空数组作为默认值
      }
    } else {
      console.error("未知错误:", error);
    }
    throw error;
  }
}
```

## 8. 实际应用示例

### 8.1 数据获取和缓存

```typescript
class DataService {
  private cache = new Map<string, any>();

  async fetchWithCache<T>(
    key: string,
    fetcher: () => Promise<T>,
    ttl: number = 5 * 60 * 1000 // 5分钟
  ): Promise<T> {
    const cached = this.cache.get(key);

    if (cached && Date.now() - cached.timestamp < ttl) {
      return cached.data;
    }

    try {
      const data = await fetcher();
      this.cache.set(key, {
        data,
        timestamp: Date.now(),
      });
      return data;
    } catch (error) {
      // 如果有缓存数据，即使过期也返回
      if (cached) {
        console.warn("使用过期缓存数据");
        return cached.data;
      }
      throw error;
    }
  }
}

// 使用示例
const dataService = new DataService();

async function getUsers(): Promise<User[]> {
  return dataService.fetchWithCache(
    "users",
    () => fetch("/api/users").then((r) => r.json()),
    10 * 60 * 1000 // 10分钟缓存
  );
}
```

### 8.2 重试机制

```typescript
async function retryOperation<T>(
  operation: () => Promise<T>,
  maxRetries: number = 3,
  delay: number = 1000
): Promise<T> {
  let lastError: Error;

  for (let attempt = 1; attempt <= maxRetries; attempt++) {
    try {
      return await operation();
    } catch (error) {
      lastError = error instanceof Error ? error : new Error(String(error));

      if (attempt === maxRetries) {
        break;
      }

      console.warn(
        `第 ${attempt} 次尝试失败，${delay}ms 后重试:`,
        lastError.message
      );
      await new Promise((resolve) => setTimeout(resolve, delay));
      delay *= 2; // 指数退避
    }
  }

  throw new Error(`操作失败，已重试 ${maxRetries} 次: ${lastError!.message}`);
}

// 使用示例
async function fetchImportantData(): Promise<string> {
  return retryOperation(
    () =>
      fetch("/api/important-data").then((r) => {
        if (!r.ok) throw new Error(`HTTP ${r.status}`);
        return r.text();
      }),
    3,
    1000
  );
}
```

## 9. 类型技巧

### 9.1 Promise 类型提取

```typescript
// 提取 Promise 的返回类型
type Awaited<T> = T extends Promise<infer U> ? U : T;

type UserType = Awaited<Promise<User>>; // User
type StringType = Awaited<Promise<string>>; // string

// 实用工具类型
type PromiseReturnType<T extends (...args: any[]) => Promise<any>> = Awaited<
  ReturnType<T>
>;

declare function fetchUser(): Promise<User>;
type FetchUserReturn = PromiseReturnType<typeof fetchUser>; // User
```

### 9.2 条件 Promise

```typescript
type MaybePromise<T> = T | Promise<T>;

function processValue<T>(value: MaybePromise<T>): Promise<T> {
  return Promise.resolve(value);
}

// 使用
const syncValue = processValue("hello"); // Promise<string>
const asyncValue = processValue(Promise.resolve("hello")); // Promise<string>
```

## 10. 性能考虑

### 10.1 避免不必要的 await

```typescript
// ❌ 不好：串行执行
async function badExample() {
  const user = await fetchUser();
  const posts = await fetchPosts();
  return { user, posts };
}

// ✅ 好：并行执行
async function goodExample() {
  const [user, posts] = await Promise.all([fetchUser(), fetchPosts()]);
  return { user, posts };
}
```

### 10.2 合理使用 Promise.all vs Promise.allSettled

```typescript
// 当需要所有操作都成功时使用 Promise.all
async function criticalOperations() {
  return Promise.all([saveUser(), saveOrder(), sendEmail()]);
}

// 当某些操作失败不影响整体时使用 Promise.allSettled
async function nonCriticalOperations() {
  const results = await Promise.allSettled([
    updateAnalytics(),
    sendNotification(),
    updateCache(),
  ]);

  // 处理失败的操作
  results.forEach((result, index) => {
    if (result.status === "rejected") {
      console.warn(`操作 ${index} 失败:`, result.reason);
    }
  });
}
```

## 迭代器全面详解

### 1.1 什么是迭代器？

迭代器是一种设计模式，它提供了一种方法来顺序访问一个聚合对象中的各个元素，而不需要暴露该对象的内部表示。

### 1.2 TypeScript 中的迭代器接口

```typescript
interface Iterator<T> {
  next(): IteratorResult<T>;
}

interface IteratorResult<T> {
  done: boolean;
  value: T;
}

interface Iterable<T> {
  [Symbol.iterator](): Iterator<T>;
}
```

## 2. 基本迭代器实现

### 2.1 简单数组迭代器

```typescript
class ArrayIterator<T> implements Iterator<T> {
  private index = 0;

  constructor(private array: T[]) {}

  next(): IteratorResult<T> {
    if (this.index < this.array.length) {
      return {
        done: false,
        value: this.array[this.index++],
      };
    }
    return {
      done: true,
      value: undefined as any,
    };
  }
}

// 使用示例
const arr = [1, 2, 3, 4, 5];
const iterator = new ArrayIterator(arr);

console.log(iterator.next()); // { done: false, value: 1 }
console.log(iterator.next()); // { done: false, value: 2 }
console.log(iterator.next()); // { done: false, value: 3 }
```

### 2.2 可迭代对象实现

```typescript
class NumberRange implements Iterable<number> {
  constructor(private start: number, private end: number) {}

  [Symbol.iterator](): Iterator<number> {
    let current = this.start;
    const end = this.end;

    return {
      next(): IteratorResult<number> {
        if (current <= end) {
          return {
            done: false,
            value: current++,
          };
        }
        return {
          done: true,
          value: undefined as any,
        };
      },
    };
  }
}

// 使用示例
const range = new NumberRange(1, 5);

// 使用 for...of 循环
for (const num of range) {
  console.log(num); // 1, 2, 3, 4, 5
}

// 手动迭代
const rangeIterator = range[Symbol.iterator]();
console.log(rangeIterator.next()); // { done: false, value: 1 }
```

## 3. 生成器（Generator）

### 3.1 生成器函数

生成器函数是创建迭代器的最简单方式：

```typescript
function* numberGenerator(start: number, end: number): Generator<number> {
  for (let i = start; i <= end; i++) {
    yield i;
  }
}

// 使用生成器
const gen = numberGenerator(1, 5);
console.log(gen.next()); // { done: false, value: 1 }
console.log(gen.next()); // { done: false, value: 2 }

// 手动执行器
function runGenerator(gen) {
  let result = gen.next();

  while (!result.done) {
    try {
      const value = result.value;
      result = gen.next(value);
    } catch (error) {
      result = gen.throw(error);
    }
  }
}

function runGeneratorRecursive(gen: Generator<number>) {
  function recursive(iter: IteratorResult<number>) {
    if (iter.done) {
      return iter.value;
    }

    try {
      return recursive(gen.next(iter.value));
    } catch (error) {
      return gen.throw(error);
    }
  }

  return recursive(gen.next());
}

runGenerator(gen);
runGeneratorRecursive(gen);

// 或使用 for...of
for (const num of numberGenerator(1, 3)) {
  console.log(num); // 1, 2, 3
}
```

### 3.2 无限生成器

```typescript
function* fibonacci(): Generator<number> {
  let a = 0,
    b = 1;
  while (true) {
    yield a;
    [a, b] = [b, a + b];
  }
}

// 获取前10个斐波那契数
const fib = fibonacci();
for (let i = 0; i < 10; i++) {
  console.log(fib.next().value);
}
```

### 3.3 异步生成器

```typescript
async function* asyncNumberGenerator(
  start: number,
  end: number
): AsyncGenerator<number> {
  for (let i = start; i <= end; i++) {
    await new Promise((resolve) => setTimeout(resolve, 1000)); // 延迟1秒
    yield i;
  }
}

// 使用异步生成器
async function useAsyncGenerator() {
  for await (const num of asyncNumberGenerator(1, 3)) {
    console.log(num); // 每秒输出一个数字
  }
}

useAsyncGenerator();
```

## 4. 内置可迭代对象

### 4.1 数组、字符串、Map、Set

```typescript
// 数组
const arr = [1, 2, 3];
for (const item of arr) {
  console.log(item);
}

// 字符串
const str = "hello";
for (const char of str) {
  console.log(char); // h, e, l, l, o
}

// Map
const map = new Map([
  ["a", 1],
  ["b", 2],
]);
for (const [key, value] of map) {
  console.log(key, value);
}

// Set
const set = new Set([1, 2, 3]);
for (const value of set) {
  console.log(value);
}
```

### 4.2 获取不同类型的迭代器

```typescript
const arr = [1, 2, 3];

// 值迭代器
for (const value of arr.values()) {
  console.log(value); // 1, 2, 3
}

// 键迭代器
for (const key of arr.keys()) {
  console.log(key); // 0, 1, 2
}

// 键值对迭代器
for (const [key, value] of arr.entries()) {
  console.log(key, value); // [0,1], [1,2], [2,3]
}
```

## 5. 高级迭代器模式

### 5.1 链式迭代器

```typescript
class ChainableIterator<T> {
  constructor(private iterable: Iterable<T>) {}

  map<U>(fn: (value: T) => U): ChainableIterator<U> {
    const iterable = this.iterable;
    return new ChainableIterator({
      *[Symbol.iterator]() {
        for (const value of iterable) {
          yield fn(value);
        }
      },
    });
  }

  filter(predicate: (value: T) => boolean): ChainableIterator<T> {
    const iterable = this.iterable;
    return new ChainableIterator({
      *[Symbol.iterator]() {
        for (const value of iterable) {
          if (predicate(value)) {
            yield value;
          }
        }
      },
    });
  }

  take(count: number): ChainableIterator<T> {
    const iterable = this.iterable;
    return new ChainableIterator({
      *[Symbol.iterator]() {
        let taken = 0;
        for (const value of iterable) {
          if (taken >= count) break;
          yield value;
          taken++;
        }
      },
    });
  }

  toArray(): T[] {
    return Array.from(this.iterable);
  }
}

// 使用示例
const result = new ChainableIterator([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
  .filter((x) => x % 2 === 0) // 筛选偶数
  .map((x) => x * 2) // 乘以2
  .take(3) // 取前3个
  .toArray();

console.log(result); // [4, 8, 12]
```

### 5.2 延迟求值迭代器

```typescript
class LazyIterator<T> {
  constructor(private generator: () => Generator<T>) {}

  static from<T>(iterable: Iterable<T>): LazyIterator<T> {
    return new LazyIterator(function* () {
      yield* iterable;
    });
  }

  map<U>(fn: (value: T) => U): LazyIterator<U> {
    const gen = this.generator;
    return new LazyIterator(function* () {
      for (const value of gen()) {
        yield fn(value);
      }
    });
  }

  filter(predicate: (value: T) => boolean): LazyIterator<T> {
    const gen = this.generator;
    return new LazyIterator(function* () {
      for (const value of gen()) {
        if (predicate(value)) {
          yield value;
        }
      }
    });
  }

  *[Symbol.iterator]() {
    yield* this.generator();
  }
}

// 使用示例 - 只在需要时计算
const lazy = LazyIterator.from([1, 2, 3, 4, 5])
  .map((x) => {
    console.log(`Processing: ${x}`);
    return x * 2;
  })
  .filter((x) => x > 4);

// 此时还没有执行任何计算
console.log("Starting iteration...");

// 只有在迭代时才会执行计算
for (const value of lazy) {
  console.log(`Result: ${value}`);
  if (value > 6) break; // 提前退出，后续的计算不会执行
}
```

## 6. 实际应用场景

### 6.1 分页数据迭代器

```typescript
class PaginatedDataIterator<T> implements AsyncIterable<T> {
  private currentPage = 1;

  constructor(
    private fetchData: (
      page: number
    ) => Promise<{ data: T[]; hasMore: boolean }>,
    private pageSize: number = 10
  ) {}

  async *[Symbol.asyncIterator](): AsyncGenerator<T> {
    let hasMore = true;

    while (hasMore) {
      const result = await this.fetchData(this.currentPage);

      for (const item of result.data) {
        yield item;
      }

      hasMore = result.hasMore;
      this.currentPage++;
    }
  }
}

// 使用示例
async function fetchUsersPage(page: number) {
  // 模拟API调用
  const users = Array.from({ length: 5 }, (_, i) => ({
    id: (page - 1) * 5 + i + 1,
    name: `User ${(page - 1) * 5 + i + 1}`,
  }));

  return {
    data: users,
    hasMore: page < 3, // 假设只有3页数据
  };
}

const userIterator = new PaginatedDataIterator(fetchUsersPage);

// 遍历所有用户，无需关心分页逻辑
for await (const user of userIterator) {
  console.log(user);
}
```

### 6.2 文件行迭代器

```typescript
class FileLineIterator implements Iterable<string> {
  constructor(private content: string) {}

  *[Symbol.iterator](): Generator<string> {
    const lines = this.content.split("\n");
    for (const line of lines) {
      if (line.trim()) {
        // 跳过空行
        yield line;
      }
    }
  }
}

// 使用示例
const fileContent = `
line 1
line 2

line 4
line 5
`;

const lineIterator = new FileLineIterator(fileContent);

for (const line of lineIterator) {
  console.log(`Processing: ${line}`);
}
```

## 7. 性能考虑和最佳实践

### 7.1 内存效率

```typescript
// 好的做法：使用迭代器处理大数据集
function* processLargeDataset(data: number[]): Generator<number> {
  for (const item of data) {
    // 只处理当前项，不需要存储整个结果数组
    yield item * 2;
  }
}

// 不好的做法：一次性处理所有数据
function processLargeDatasetBad(data: number[]): number[] {
  return data.map((item) => item * 2); // 可能导致内存溢出
}
```

### 7.2 错误处理

```typescript
function* safeIterator<T>(iterable: Iterable<T>): Generator<T> {
  try {
    for (const item of iterable) {
      yield item;
    }
  } catch (error) {
    console.error("Iterator error:", error);
    // 可以选择重新抛出错误或返回默认值
    throw error;
  }
}
```

### 7.3 类型安全

```typescript
// 使用泛型确保类型安全
interface DataProcessor<T, U> {
  process(input: Iterable<T>): Generator<U>;
}

class NumberProcessor implements DataProcessor<number, string> {
  *process(input: Iterable<number>): Generator<string> {
    for (const num of input) {
      yield num.toString();
    }
  }
}
```

## TypeScript 函数详解

### 函数声明

```typescript
function greet(name: string): string {
  return `Hello, ${name}!`;
}
```

### 函数表达式

```typescript
const greet = function (name: string): string {
  return `Hello, ${name}!`;
};
```

### 箭头函数

```typescript
const greet = (name: string): string => {
  return `Hello, ${name}!`;
};

// 简写形式
const greet = (name: string): string => `Hello, ${name}!`;
```

## 2. 参数类型

### 基本参数类型

```typescript
function add(a: number, b: number): number {
  return a + b;
}

function isAdult(age: number): boolean {
  return age >= 18;
}

function processUser(user: { name: string; age: number }): void {
  console.log(`Processing ${user.name}, age ${user.age}`);
}
```

### 可选参数

```typescript
function createUser(name: string, age?: number): string {
  if (age !== undefined) {
    return `${name} is ${age} years old`;
  }
  return `${name} age unknown`;
}

createUser("Alice"); // ✅ 正确
createUser("Bob", 25); // ✅ 正确
```

### 默认参数

```typescript
function greet(name: string, greeting: string = "Hello"): string {
  return `${greeting}, ${name}!`;
}

greet("Alice"); // "Hello, Alice!"
greet("Bob", "Hi"); // "Hi, Bob!"
```

### 剩余参数

```typescript
function sum(...numbers: number[]): number {
  return numbers.reduce((total, num) => total + num, 0);
}

sum(1, 2, 3, 4, 5); // 15

function logMessages(prefix: string, ...messages: string[]): void {
  messages.forEach((msg) => console.log(`${prefix}: ${msg}`));
}
```

## 3. 返回值类型

### 明确指定返回类型

```typescript
function multiply(a: number, b: number): number {
  return a * b;
}

function getUser(): { name: string; age: number } {
  return { name: "Alice", age: 30 };
}

function logMessage(message: string): void {
  console.log(message);
  // 函数不返回值
}
```

### 返回类型推断

```typescript
// TypeScript 会自动推断返回类型为 number
function divide(a: number, b: number) {
  return a / b;
}
```

## 4. 函数类型

### 函数类型注解

```typescript
// 定义函数类型
type MathOperation = (a: number, b: number) => number;

const add: MathOperation = (a, b) => a + b;
const subtract: MathOperation = (a, b) => a - b;

// 作为参数传递
function calculate(operation: MathOperation, x: number, y: number): number {
  return operation(x, y);
}

calculate(add, 5, 3); // 8
calculate(subtract, 5, 3); // 2
```

### 接口定义函数类型

```typescript
interface Calculator {
  (a: number, b: number): number;
}

const multiply: Calculator = (a, b) => a * b;
```

## 5. 函数重载

```typescript
// 重载签名
function process(input: string): string;
function process(input: number): number;
function process(input: boolean): boolean;

// 实现签名
function process(input: string | number | boolean): string | number | boolean {
  if (typeof input === "string") {
    return input.toUpperCase();
  } else if (typeof input === "number") {
    return input * 2;
  } else {
    return !input;
  }
}

process("hello"); // 返回 "HELLO"
process(5); // 返回 10
process(true); // 返回 false
```

## 6. 泛型函数

### 基本泛型函数

```typescript
function identity<T>(arg: T): T {
  return arg;
}

const stringResult = identity<string>("hello"); // string
const numberResult = identity<number>(42); // number
const autoResult = identity("world"); // 自动推断为 string
```

### 多个泛型参数

```typescript
function pair<T, U>(first: T, second: U): [T, U] {
  return [first, second];
}

const result = pair<string, number>("age", 25); // [string, number]
```

### 泛型约束

```typescript
interface Lengthwise {
  length: number;
}

function logLength<T extends Lengthwise>(arg: T): T {
  console.log(arg.length);
  return arg;
}

logLength("hello"); // ✅ 字符串有 length 属性
logLength([1, 2, 3]); // ✅ 数组有 length 属性
// logLength(123);          // ❌ 数字没有 length 属性
```

## 7. 高级函数模式

### 回调函数

```typescript
type Callback<T> = (result: T) => void;

function fetchData<T>(url: string, callback: Callback<T>): void {
  // 模拟异步操作
  setTimeout(() => {
    // 假设获取到数据
    const data = { message: "success" } as T;
    callback(data);
  }, 1000);
}

fetchData<{ message: string }>("/api/data", (result) => {
  console.log(result.message);
});
```

### Promise 函数

```typescript
async function fetchUser(id: number): Promise<{ name: string; age: number }> {
  // 模拟 API 调用
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({ name: "Alice", age: 30 });
    }, 1000);
  });
}

// 使用
async function main() {
  try {
    const user = await fetchUser(1);
    console.log(user.name);
  } catch (error) {
    console.error("Error fetching user:", error);
  }
}
```

### 高阶函数

```typescript
function createMultiplier(factor: number): (value: number) => number {
  return (value: number) => value * factor;
}

const double = createMultiplier(2);
const triple = createMultiplier(3);

console.log(double(5)); // 10
console.log(triple(5)); // 15
```

## 8. 函数参数解构

### 对象解构

```typescript
interface User {
  name: string;
  age: number;
  email?: string;
}

function greetUser({ name, age, email = "未提供" }: User): string {
  return `Hello ${name}, ${age} years old, email: ${email}`;
}

greetUser({ name: "Alice", age: 30 });
```

### 数组解构

```typescript
function processCoordinates([x, y, z = 0]: [number, number, number?]): string {
  return `Position: (${x}, ${y}, ${z})`;
}

processCoordinates([10, 20]); // "Position: (10, 20, 0)"
processCoordinates([1, 2, 3]); // "Position: (1, 2, 3)"
```

## 9. this 参数

```typescript
interface Counter {
  count: number;
  increment(this: Counter): void;
  getValue(this: Counter): number;
}

const counter: Counter = {
  count: 0,
  increment(this: Counter) {
    this.count++;
  },
  getValue(this: Counter) {
    return this.count;
  },
};
```

## 10. 实用技巧

### 函数类型守卫

```typescript
function isString(value: unknown): value is string {
  return typeof value === "string";
}

function processValue(value: unknown) {
  if (isString(value)) {
    // 这里 TypeScript 知道 value 是 string 类型
    console.log(value.toUpperCase());
  }
}
```

### 条件返回类型

```typescript
function processInput<T extends string | number>(
  input: T
): T extends string ? string : number {
  if (typeof input === "string") {
    return input.toUpperCase() as any;
  }
  return (input * 2) as any;
}

const stringResult = processInput("hello"); // string
const numberResult = processInput(42); // number
```

## 总结

TypeScript 函数的主要特点：

1. **类型安全**：参数和返回值都有类型检查
2. **灵活性**：支持可选参数、默认参数、剩余参数
3. **函数重载**：同一函数名支持多种调用方式
4. **泛型支持**：编写可重用的通用函数
5. **高级特性**：类型守卫、条件类型等

## Blob Buffer Stream

### Blob

> Blob 对象表示一个不可变的、原始数据的类文件对象。它可以包含任意类型的数据，如文本、图片、音频、视频等。

```typescript
interface Blob {
  readonly size: number;
  readonly type: string;
  arrayBuffer(): Promise<ArrayBuffer>;
  text(): Promise<string>;
  stream(): ReadableStream<Uint8Array>;
  slice(start?: number, end?: number, contentType?: string): Blob;
}

interface BlobPropertyBag {
  type?: string;
  endings?: "transparent" | "native";
}

declare var Blob: {
  prototype: Blob;
  new (blobParts?: BlobPart[], options?: BlobPropertyBag): Blob;
};

type BlobPart = BufferSource | Blob | string;

// 核心属性
const blob = new Blob(["Hello World"]);

const textBlob = new Blob(["Hello, World!"], { type: "text/plain" }); // 从字符串创建
const mixedBlob = new Blob(["Hello", " ", "World"], { type: "text/plain" }); // 从多个部分创建
const jsonBlob = new Blob([JSON.stringify({ name: "test" })], {
  type: "application/json",
}); // 从JSON创建

const buffer = new ArrayBuffer(8);
const view = new Uint8Array(buffer);
view[0] = 72; // 'H'
view[1] = 101; // 'e'

const blobFromBuffer = new Blob([buffer], { type: "application/octet-stream" });

const originalBlob = new Blob(["Hello"]);
const newBlob = new Blob([originalBlob, " World"], { type: "text/plain" });

console.log(blob.size); // 11

// 核心方法
const blob = new Blob(['{"name": "test"}'], { type: "application/json" });
console.log(blob.type); // "application/json"

await blob.arrayBuffer();
await blob.text();

const stream = blob.stream();

const blob = new Blob(["Hello World"]);
const sliced = blob.slice(0, 5); // 包含 "Hello"

// 文件下载
function downloadFile(data: string, filename: string, mimeType: string): void {
  const blob = new Blob([data], { type: mimeType });
  const url = URL.createObjectURL(blob);

  const link = document.createElement("a");
  link.href = url;
  link.download = filename;
  link.click();
}

downloadFile("Hello World", "hello.txt", "text/plain");

class BlobUtils {
  // 将多个Blob合并
  static merge(blobs: Blob[], type?: string): Blob {
    return new Blob(blobs, { type });
  }

  // 检查Blob是否为空
  static isEmpty(blob: Blob): boolean {
    return blob.size === 0;
  }

  // 比较两个Blob的大小
  static async equals(blob1: Blob, blob2: Blob): Promise<boolean> {
    if (blob1.size !== blob2.size) return false;

    const [buffer1, buffer2] = await Promise.all([
      blob1.arrayBuffer(),
      blob2.arrayBuffer(),
    ]);

    const view1 = new Uint8Array(buffer1);
    const view2 = new Uint8Array(buffer2);

    return view1.every((byte, index) => byte === view2[index]);
  }
}
```

### Buffer 概述

> Buffer 是 Node.js 中用于处理二进制数据的全局类，在 TypeScript 中有完整的类型定义支持。它类似于整数数组，但专门用于存储原始二进制数据

```typescript
// Buffer的类型定义示例
interface BufferConstructor {
  alloc(
    size: number,
    fill?: string | number | Buffer,
    encoding?: BufferEncoding
  ): Buffer;
  allocUnsafe(size: number): Buffer;
  from(str: string, encoding?: BufferEncoding): Buffer;
  from(data: number[]): Buffer;
  from(arrayBuffer: ArrayBuffer): Buffer;
  concat(list: Buffer[], totalLength?: number): Buffer;
  // ... 更多方法
}

interface Buffer {
  length: number;
  write(string: string, offset?: number, encoding?: BufferEncoding): number;
  toString(encoding?: BufferEncoding, start?: number, end?: number): string;
  slice(start?: number, end?: number): Buffer;
  copy(
    target: Buffer,
    targetStart?: number,
    sourceStart?: number,
    sourceEnd?: number
  ): number;
  // ... 更多方法
}

// 2. 创建指定大小的Buffer（内容初始化为0，安全）
const buf2 = Buffer.alloc(10);
const buf3 = Buffer.alloc(10, 1); // 用1填充

// 3. 从字符串创建
const buf4 = Buffer.from("hello", "utf8");
const buf5 = Buffer.from("SGVsbG8=", "base64");

// 4. 从数组创建
const buf6 = Buffer.from([1, 2, 3, 4, 5]);

// 5. 从ArrayBuffer创建
const arrayBuffer = new ArrayBuffer(10);
const buf7 = Buffer.from(arrayBuffer);

// 拼接
const buf8 = Buffer.concat([buf1, buf2]);

// 切片（返回新的Buffer，但共享内存）
const buf9 = buf1.slice(1, 4);

// 读取为字符串
const str1 = buf.toString(); // 默认utf8
const str2 = buf.toString("hex"); // 16进制字符串
const str3 = buf.toString("base64"); // base64字符串

// 写入字符串
buf.write("hello");
buf.write("hello", 0, "utf8");
```

### Stream

#### Readable Stream（可读流）

```typescript
interface StreamOptions {
  highWaterMark?: number;
  encoding?: BufferEncoding;
  objectMode?: boolean;
}

interface ReadableOptions extends StreamOptions {
  read?(size: number): void;
}

import { Readable } from "stream";

class MyReadableStream extends Readable {
  private current = 0;

  _read() {
    if (this.current < 5) {
      this.push(`data-${this.current++}`);
    } else {
      this.push(null); // 结束流
    }
  }
}

const readable = new MyReadableStream();

readable
  .on("data", (chunk) => {
    console.log("接收到数据:", chunk.toString());
  })
  .on("end", () => {
    console.log("流结束");
  })
  .on("error", console.error);
```

#### Writable Stream（可写流）

```typescript
interface StreamOptions {
  highWaterMark?: number;
  encoding?: BufferEncoding;
  objectMode?: boolean;
}

interface WritableOptions extends StreamOptions {
  write?(chunk: any, encoding: BufferEncoding, callback: Function): void;
  write?(
    chunks: Array<{ chunk: any; encoding: BufferEncoding }>,
    callback: Function
  ): void;
}

import { Writable } from "stream";
import { WriteStream } from "fs";

class MyWritableStream extends Writable {
  _write(chunk: any, encoding: BufferEncoding, callback: Function) {
    console.log("写入数据:", chunk.toString());
    callback(); // 完成写入
  }
}

const writable = new MyWritableStream();

writable.write("Hello ").write("World!").end();

writable
  .on("finish", () => {
    console.log("写入完成");
  })
  .on("error", (error) => {
    console.error("发生错误:", error);
  });
```

#### Duplex Stream（双工流）

```typescript
import { Duplex } from "stream";

class MyDuplexStream extends Duplex {
  _read() {
    this.push("从双工流读取的数据");
    this.push(null);
  }

  _write(chunk: any, encoding: BufferEncoding, callback: Function) {
    console.log("双工流写入:", chunk.toString());
    callback();
  }
}
```

#### Transform Stream（转换流）

```typescript
import { Transform } from "stream";

class UpperCaseTransform extends Transform {
  _transform(chunk: any, encoding: BufferEncoding, callback: Function) {
    this.push(chunk.toString().toUpperCase());
    callback();
  }
}

const upperCaseTransform = new UpperCaseTransform();
process.stdin.pipe(upperCaseTransform).pipe(process.stdout);
```

#### 文件处理

```typescript
import { createReadStream, createWriteStream } from "fs";
import { Transform } from "stream";

// 创建一个转换流来处理 JSON 数据
class JSONProcessor extends Transform {
  _transform(chunk: any, encoding: BufferEncoding, callback: Function) {
    try {
      const data = JSON.parse(chunk.toString());
      const processed = { ...data, processed: true };
      this.push(JSON.stringify(processed) + "\n");
      callback();
    } catch (error) {
      callback(error);
    }
  }
}

// 管道操作
const readStream = new createReadStream("input.json");
const writeStream = new createWriteStream("output.json");
const processor = new JSONProcessor();

readStream
  .pipe(processor)
  .pipe(writeStream)
  .on("end", () => {
    console.log("文件处理完成");
  });
```

#### 文件处理

```typescript
import { IncomingMessage } from "http";
import { Transform } from "stream";

class ResponseProcessor extends Transform {
  constructor() {
    super({ objectMode: true });
  }

  _transform(chunk: any, encoding: BufferEncoding, callback: Function) {
    // 处理 HTTP 响应数据
    const processed = chunk.toString().replace(/\r\n/g, "\n");
    this.push(processed);
    callback();
  }
}

function processHttpResponse(response: IncomingMessage) {
  const processor = new ResponseProcessor();

  response
    .pipe(processor)
    .on("data", (chunk) => {
      console.log("处理后的数据:", chunk);
    })
    .on("end", () => {
      console.log("响应处理完成");
    });
}
```

#### 异步迭代器与流

```typescript
import { Readable } from "stream";

async function processStreamWithAsyncIterator(stream: Readable) {
  try {
    for await (const chunk of stream) {
      console.log("异步处理:", chunk.toString());
      // 可以在这里进行异步操作
      await new Promise((resolve) => setTimeout(resolve, 100));
    }
  } catch (error) {
    console.error("流处理错误:", error);
  }
}

// 使用示例
const readable = new Readable({
  read() {
    this.push("chunk 1");
    this.push("chunk 2");
    this.push(null);
  },
});

processStreamWithAsyncIterator(readable);
```

## 模块与命名空间

### 1. 模块导入导出

```typescript
// math.ts - 导出
export function add(a: number, b: number): number {
  return a + b;
}

export const PI = 3.14159;

export default class Calculator {
  add(a: number, b: number): number {
    return a + b;
  }
}

// main.ts - 导入
import Calculator, { add, PI } from "./math";
import * as MathUtils from "./math";

const calc = new Calculator();
const result = add(5, 3);
```

### 2. 命名空间

```typescript
namespace Geometry {
  export interface Point {
    x: number;
    y: number;
  }

  export class Circle {
    constructor(public center: Point, public radius: number) {}

    area(): number {
      return Math.PI * this.radius * this.radius;
    }
  }

  export function distance(p1: Point, p2: Point): number {
    return Math.sqrt((p2.x - p1.x) ** 2 + (p2.y - p1.y) ** 2);
  }
}

// 使用命名空间
const center: Geometry.Point = { x: 0, y: 0 };
const circle = new Geometry.Circle(center, 5);
```

# I/O

## path

绝对路径：从文件系统根目录开始的完整路径

- Windows: C:\Users\username\project\file.js
- Unix/Linux/macOS: /home/username/project/file.js

相对路径：相对于当前工作目录或当前文件的路径

- ./file.js - 当前目录下的文件
- ../parent/file.js - 上级目录中的文件
- folder/file.js - 子目录中的文件

```typescript
// 假设文件位置: /project/src/app.js
console.log(__dirname); // /project/src 当前执行文件所在的目录的绝对路径
console.log(process.cwd()); // /project Node.js 进程启动时的工作目录

const path = require("path");

// 危险 - 依赖于启动目录
fs.readFileSync("./config.json");

// 安全 - 相对于当前文件
fs.readFileSync(path.join(__dirname, "config.json"));

// 路径拼接 - 自动处理分隔符
path.join("/users", "john", "documents", "file.txt");
// 结果: /users/john/documents/file.txt

// 路径解析 - 返回绝对路径
path.resolve("file.txt");
path.resolve("/users", "john", "../jane", "file.txt");

// 路径解析
path.dirname("/path/to/file.txt"); // /path/to
path.basename("/path/to/file.txt"); // file.txt
path.extname("/path/to/file.txt"); // .txt

// 路径信息解析
path.parse("/path/to/file.txt");
// { root: '/', dir: '/path/to', base: 'file.txt', ext: '.txt', name: 'file' }
```

## 基础概念

```typescript
const fs = require("fs");
const path = require("path");

// Promise 版本 (推荐)
const fs.promises = require("fs").promises;
// 或者
const fs = require("fs/promises");
```

## 同步文件操作

### 读取文件

```typescript
// 回调方式
fs.readFile("example.txt", "utf8", (err, data) => {
  if (err) {
    console.error("读取文件失败:", err);
    return;
  }
  console.log("文件内容:", data);
});

// Promise 方式
fs.promises
  .readFile("example.txt", "utf8")
  .then((data) => {
    console.log("文件内容:", data);
  })
  .catch((err) => {
    console.error("读取文件失败:", err);
  });

// async/await 方式 (推荐)
async function readFileExample() {
  try {
    const data = await fs.promises.readFile("example.txt", "utf8");
    console.log("文件内容:", data);
  } catch (err) {
    console.error("读取文件失败:", err);
  }
}

// 同步读取文件
try {
  const data = fs.readFileSync("example.txt", "utf8");
  console.log("文件内容:", data);
} catch (err) {
  console.error("读取文件失败:", err);
}

//  流式读取 (适合大文件)
const readStream = fs.createReadStream("large-file.txt", {
  encoding: "utf8",
  highWaterMark: 16 * 1024, // 16KB 缓冲区
});

readStream.on("data", (chunk) => {
  console.log("读取数据块:", chunk.length);
});

readStream.on("end", () => {
  console.log("文件读取完成");
});

readStream.on("error", (err) => {
  console.error("读取出错:", err);
});
```

### 写入文件

```typescript
// Promise 方式
async function writeFileExample() {
  try {
    await fs.promises.writeFile("output.txt", "这是写入的内容", "utf8");
    console.log("文件写入成功");
  } catch (err) {
    console.error("写入失败:", err);
  }
}

// 追加内容
async function appendFileExample() {
  try {
    await fs.promises.appendFile("output.txt", "\n追加的内容", "utf8");
    console.log("内容追加成功");
  } catch (err) {
    console.error("追加失败:", err);
  }
}

// 流式写入 (适合大文件)
const writeStream = fs.createWriteStream("output.txt", { encoding: "utf8" });

writeStream.write("第一行内容\n");
writeStream.write("第二行内容\n");
writeStream.end(); // 关闭流

writeStream.on("finish", () => {
  console.log("写入完成");
});

writeStream.on("error", (err) => {
  console.error("写入出错:", err);
});
```

### 文件信息

```typescript
enum Permission {
  OwnerRead = 0o400,
  OwnerWrite = 0o200,
  OwnerExecute = 0o100,
  GroupRead = 0o040,
  GroupWrite = 0o020,
  GroupExecute = 0o010,
  OtherRead = 0o004,
  OtherWrite = 0o002,
  OtherExecute = 0o001,
}

// 检查文件状态
async function getFileStats() {
  try {
    const stats = await fs.promises.stat("example.txt");

    console.log("文件大小:", stats.size, "bytes");
    console.log("是否为文件:", stats.isFile());
    console.log("是否为目录:", stats.isDirectory());

    console.log("创建时间:", stats.birthtime);
    console.log("修改时间:", stats.mtime);
    console.log("访问时间:", stats.atime);
    console.log("权限(mode):", Boolean(stats.mode & 0o400)); // 八进制权限

    console.log("所有者用户ID(uid):", stats.uid);
    console.log("所有者群组ID(gid):", stats.gid);
  } catch (err) {
    console.error("获取文件信息失败:", err);
  }
}

// 检查文件是否存在
async function checkFileExists(filePath) {
  try {
    await fs.promises.access(filePath, fs.constants.F_OK);
    console.log("文件存在");
    return true;
  } catch (err) {
    console.log("文件不存在");
    return false;
  }
}

// 检查权限
async function checkPermissions(filePath) {
  try {
    // 检查读写权限
    await fs.promises.access(filePath, fs.constants.R_OK | fs.constants.W_OK);
    console.log("文件可读写");
  } catch (err) {
    console.log("文件权限不足");
  }
}
```

## 目录操作

### 创建和删除目录

```typescript
// 创建目录
async function createDirectory() {
  try {
    // 创建单级目录
    await fs.promises.mkdir("new-directory");

    // 创建多级目录
    await fs.promises.mkdir("path/to/deep/directory", { recursive: true });

    console.log("目录创建成功");
  } catch (err) {
    console.error("创建目录失败:", err);
  }
}

// 删除目录
async function removeDirectory() {
  try {
    // 删除空目录
    await fs.promises.rmdir("empty-directory");

    // 递归删除目录及其内容
    await fs.promises.rm("directory-with-files", {
      recursive: true,
      force: true,
    });

    console.log("目录删除成功");
  } catch (err) {
    console.error("删除目录失败:", err);
  }
}

// 读取目录内容
async function readDirectory() {
  try {
    // 读取文件名列表
    const files = await fs.promises.readdir("./");
    console.log("目录内容:", files);

    // 获取详细信息
    const filesWithStats = await fs.promises.readdir("./", {
      withFileTypes: true,
    });
    filesWithStats.forEach((dirent) => {
      console.log(`${dirent.name} - ${dirent.isDirectory() ? "目录" : "文件"}`);
    });
  } catch (err) {
    console.error("读取目录失败:", err);
  }
}
```

### 文件操作

```typescript
// 文件复制
async function copyFile() {
  try {
    await fs.promises.copyFile("source.txt", "destination.txt");
    console.log("文件复制成功");
  } catch (err) {
    console.error("复制失败:", err);
  }
}

// 文件移动/重命名
async function moveFile() {
  try {
    await fs.promises.rename("old-name.txt", "new-name.txt");
    console.log("文件重命名成功");
  } catch (err) {
    console.error("重命名失败:", err);
  }
}

// 文件删除
async function deleteFile() {
  try {
    await fs.promises.unlink("file-to-delete.txt");
    console.log("文件删除成功");
  } catch (err) {
    console.error("删除失败:", err);
  }
}

// 修改文件权限
async function changeFileMode() {
  try {
    // 0o644: 所有者可读写，组和其他用户可读
    await fs.promises.chmod("file.txt", 0o644);
    console.log("文件权限修改成功");
  } catch (err) {
    console.error("修改权限失败:", err);
  }
}

// 修改文件所有者和群组
async function changeFileOwner() {
  try {
    // 1000 和 1000 通常是当前用户和群组的 id，请根据实际情况修改
    await fs.promises.chown("file.txt", 1000, 1000);
    console.log("文件所有者/群组修改成功");
  } catch (err) {
    console.error("修改所有者/群组失败:", err);
  }
}
```

## 文件监控

```typescript
// 监控文件变化
function watchFile(filePath: string): void {
  fs.watchFile(filePath, (curr: fs.Stats, prev: fs.Stats) => {
    console.log("文件发生变化");
    console.log("当前大小:", curr.size);
    console.log("之前大小:", prev.size);
    console.log("修改时间:", curr.mtime);
  });
}

// 监控目录变化
function watchDirectory(dirPath: string): void {
  const watcher = fs.watch(dirPath, (eventType: string, filename: string) => {
    console.log("事件类型:", eventType);
    console.log("文件名:", filename);
  });

  // 停止监控
  setTimeout(() => {
    watcher.close();
    console.log("停止监控目录");
  }, 10000);
}
```

## 错误处理

```typescript
async function robustFileOperation(filePath) {
  try {
    // 检查文件是否存在
    await fs.promises.access(filePath, fs.constants.F_OK);

    // 读取文件
    const data = await fs.promises.readFile(filePath, "utf8");
    return data;
  } catch (err) {
    switch (err.code) {
      case "ENOENT":
        console.error("文件不存在:", filePath);
        break;
      case "EACCES":
        console.error("权限不足:", filePath);
        break;
      case "EISDIR":
        console.error("这是一个目录，不是文件:", filePath);
        break;
      default:
        console.error("未知错误:", err.message);
    }
    throw err;
  }
}
```

# Network

## server

```typescript
const http = require("http");
const fs = require("fs");
const path = require("path");

// 中间件
function compose(middlewares) {
  return function (req, res, next) {
    let index = 0;

    (function recursive() {
      if (index > middlewares.length) {
        return;
      }

      if (index === middlewares.length) {
        next();
        return;
      }

      const middleware = middlewares[index++];

      middleware(req, res, recursive);
    })();
  };
}

const accessLogStream = fs.createWriteStream(
  path.join(__dirname, "access.log"),
  { flags: "a" }
);

function infoLogger(req, res, next) {
  const start = Date.now();

  // 收集请求信息
  const requestInfo = {
    timestamp: new Date().toISOString(),
    method: req.method,
    url: req.url,
    headers: req.headers,
    ip: req.socket.remoteAddress,
    body: null,
  };

  // 记录请求体
  let body = [];

  req
    .on("data", (chunk) => {
      body.push(chunk);
    })
    .on("end", () => {
      if (body.length > 0) {
        requestInfo.body = Buffer.concat(body).toString();
      }
    });

  // 记录响应信息
  const originalEnd = res.end;

  res.end = function (chunk, encoding) {
    const responseInfo = {
      statusCode: res.statusCode,
      duration: Date.now() - start,
      responseHeaders: res.getHeaders(),
      responseBody: chunk.toString(),
    };

    // 组合日志信息
    const logEntry = {
      request: requestInfo,
      response: responseInfo,
    };

    // 写入日志文件
    accessLogStream.write(JSON.stringify(logEntry) + "\n");

    // 控制台输出
    console.log(
      `[${requestInfo.timestamp}] ${requestInfo.method} ${requestInfo.url} ${responseInfo.statusCode} - ${responseInfo.duration}ms`
    );

    // 调用原始的 end 方法
    originalEnd.call(this, chunk, encoding);
  };

  next();
}

const errorLogStream = fs.createWriteStream(path.join(__dirname, "error.log"), {
  flags: "a",
});

function errorLogger(req, res, next) {
  const start = Date.now();

  // 记录请求信息
  const requestInfo = {
    timestamp: new Date().toISOString(),
    method: req.method,
    url: req.url,
    headers: req.headers,
    ip: req.socket.remoteAddress,
  };

  // 错误处理
  req.on("error", (error) => {
    const errorInfo = {
      timestamp: new Date().toISOString(),
      type: "request_error",
      error: error.message,
      stack: error.stack,
      request: requestInfo,
    };

    errorLogStream.write(JSON.stringify(errorInfo) + "\n");
    console.error("Request Error:", error);
  });

  res.on("error", (error) => {
    const errorInfo = {
      timestamp: new Date().toISOString(),
      type: "response_error",
      error: error.message,
      stack: error.stack,
      request: requestInfo,
    };

    errorLogStream.write(JSON.stringify(errorInfo) + "\n");
    console.error("Response Error:", error);
  });

  // 记录响应信息
  const originalEnd = res.end;

  res.end = function (chunk, encoding) {
    const responseInfo = {
      statusCode: res.statusCode,
      duration: Date.now() - start,
      responseHeaders: res.getHeaders(),
      responseBody: chunk.toString(),
    };

    // 记录错误状态码
    if (res.statusCode >= 400) {
      const errorInfo = {
        timestamp: new Date().toISOString(),
        type: "http_error",
        statusCode: res.statusCode,
        request: requestInfo,
        response: responseInfo,
      };

      errorLogStream.write(JSON.stringify(errorInfo) + "\n");
    }

    originalEnd.call(this, chunk, encoding);
  };

  next();
}

// 路由拆分
const routes = {
  "/": (req, res) => {
    // 设置响应头
    res.writeHead(200, {
      "Content-Type": "application/json",
    });

    // 发送响应
    res.end(
      JSON.stringify({
        message: "Home page",
        url: req.url,
        headers: req.headers,
      })
    );
  },
};

function handler(req, res) {
  const reqPath = req.url.split("?")[0];
  const func = routes[reqPath];

  if (func) {
    func(req, res);
    return;
  }

  res.writeHead(404);
  res.end("Not Found");
}

http
  .createServer((req, res) => {
    // 设置超时
    req.setTimeout(5000);

    // 中间件
    const middleware = compose([infoLogger, errorLogger]);

    middleware(req, res, () => {
      // 路由拆分
      switch (req.method) {
        case "GET":
        case "POST":
          handler(req, res);
          break;
        default:
          res.writeHead(405);
          res.end("Method Not Allowed");
      }
    });
  })
  .on("error", (error) => {
    console.error("Server error:", error);
  })
  .listen(3000);
```

## client

```typescript
// http.request
const http = require("http");

const request = http.request(
  {
    hostname: "localhost",
    port: 3000,
    path: "/",
    method: "POST",
  },
  (res) => {
    console.log(`STATUS: ${res.statusCode}`);
    console.log(`HEADERS: ${JSON.stringify(res.headers)}`);

    res
      .on("data", (chunk) => {
        console.log(`BODY: ${chunk}`);
      })
      .on("end", () => {
        console.log("No more data in response.");
      });
  }
);

request.on("error", (error) => {
  console.error(`problem with request: ${error.message}`);
});

request.end();

// fetch 函数签名
declare function fetch(
  input: RequestInfo | URL,
  init?: RequestInit
): Promise<Response>;

// RequestInfo 类型
type RequestInfo = Request | string;

// RequestInit 配置对象
interface RequestInit {
  method?: string;
  headers?: HeadersInit;
  body?: BodyInit | null;
  mode?: RequestMode;
  credentials?: RequestCredentials;
  cache?: RequestCache;
  redirect?: RequestRedirect;
  referrer?: string;
  referrerPolicy?: ReferrerPolicy;
  integrity?: string;
  signal?: AbortSignal | null;
}

interface Response {
  readonly status: number; // HTTP 状态码
  readonly statusText: string; // 状态文本
  readonly ok: boolean; // 200-299 范围内为 true
  readonly headers: Headers; // 响应头
  readonly url: string; // 请求的 URL
  readonly redirected: boolean; // 是否被重定向
  json(): Promise<any>; // 解析 JSON
  text(): Promise<string>; // 获取文本
  blob(): Promise<Blob>; // 获取 Blob
  arrayBuffer(): Promise<ArrayBuffer>; // 获取 ArrayBuffer
  formData(): Promise<FormData>; // 获取 FormData
  clone(): Response; // 克隆响应对象
}

(async function () {
  try {
    // 取消请求
    const controller = new AbortController();

    const timeoutId = setTimeout(() => controller.abort(), timeout);

    // 请求参数，需手动过滤 null、undefined
    const params = new URLSearchParams({
      name: "John",
      age: "30",
    });
    const jsonHeaders = new Headers({
      "Content-Type": "application/json",
    });
    const jsonData = JSON.stringify({
      name: "John",
      age: "30",
    });

    // form
    const formData = new FormData();

    formData.append("file", file);

    const response = await fetch(`http://localhost:3000?${params}`, {
      method: "POST",
      headers,
      body: jsonData,
      signal: controller.signal,
    });

    if (!response.ok) {
      throw new Error("POST Error:", response.status);
    }

    const contentType = response.headers.get("content-type");

    if (!contentType) {
      throw new Error("No Content-Type");
    }

    if (contentType.includes("application/json")) {
      const json = await response.json();
      console.log(json);
    } else if (contentType.includes("text/html")) {
      const text = await response.text();
      console.log(text);
    } else if (contentType.includes("application/octet-stream")) {
      const reader = response.body.getReader();

      const chunks = [];
      while (true) {
        const { done, value } = await reader.read();
        if (done) {
          break;
        }
        chunks.push(value);
      }

      const text = Buffer.concat(chunks).toString();
      console.log(text);
    } else {
      const blob = await response.blob();
      console.log(blob);
    }
  } catch (error) {
    if (error instanceof Error && error.name === "AbortError") {
      throw new Error("Request timeout");
    }

    throw error;
  }
})();
```

# Process

## 事件循环

### [事件](https://www.digitalocean.com/community/tutorials/using-event-emitters-in-node-js)

```typescript
const EventEmitter = require("events");

class TicketManager extends EventEmitter {
  constructor(supply) {
    super();

    this.supply = supply;
  }

  buy(email, price) {
    if (this.supply > 0) {
      this.supply--;
      this.emit("buy", email, price, Date.now());
      return;
    }

    this.emit("error", new Error("There are no more tickets left to purchase"));
  }
}

const ticketManager = new TicketManager(10);

ticketManager
  .on("buy", (email, price, timestamp) => {
    EmailService.send(email);
    DatabaseService.save(email, price, timestamp);
  })
  .on("error", (error) => {
    console.error(`Gracefully handling our error: ${error}`);
  });

ticketManager.buy("test@email.com", 20);
```

### [Javascript Runtime](https://www.lydiahallie.com/blog/event-loop)

- Call Stack 堆栈管理我们程序的执行，当我们执行程序时，将创建一个新的执行上下文，将其推入堆栈。当功能完成其执行时，执行上下文会从呼叫堆栈中弹出。JavaScript 只能一次处理一个任务；如果一个任务花费太长以无法弹出，则无法处理其他任务，这意味着更长的运行任务可以阻止任何其他任务执行，从本质上讲我们的程序。

- Event Loop 事件循环的责任是连续检查呼叫堆栈是否为空。每当呼叫堆栈为空时（这意味着当前没有运行的任务），它将从任务队列中获取第一个可用任务，并将其移至呼叫堆栈中，该任务被执行。在移至任务队列之前先从微任务队列中处理所有微任务。

- Task Queue 任务队列，将 SuccessCallback 添加到任务队列（由于这个确切的原因，也称为回调队列）。任务队列保留 Web API 回调和事件处理程序，等待将来在某个时候执行

- MicroTask Queue 微任务队列，通过 Promise 处理程序（或使用等待）而不是使用回调来处理返回的数据，微任务队列是运行时的另一个队列，优先级高于任务队列

  ```typescript
  Promise.then(callback).catch(callback).finally(callback);
  queueMicrotask(callback);
  new MutationObserver(callback).observe(target).disconnect();

  async function() { await }

  // Generator + Promise 实现 async
  function* generator() {
    const result1 = yield Promise.resolve(1);
    console.log(result1); // 1

    const result2 = yield Promise.reject(new Error('Not found.'));
    console.log(result2); // 2

    return '完成';
  }

  function run(generator) {
    const iterator = generator();

    function recursive(result) {
      if (result.done) {
        return result.value;
      }

      return Promise.resolve(result.value)
        .then(value => recursive(iterator.next(value)))
        .catch(error => iterator.throw(error));
    }

    return recursive(iterator.next());
  }

  run(generator)
    .then(result => console.log(result))
    .catch(error => console.error(error));

  ```

- process.nextTick Queue 优先级高于微任务队列

  ```typescript
  process.nextTick(callback);
  ```

- Web APIS 提供了一组接口，可与浏览器的功能进行交互。

- APIS

  - Rendering Engines 渲染引擎
  - Network Service 网络请求

    ```typescript
    const xhr = new XMLHttpRequest();

    xhr.open("GET", "https://api.github.com");

    xhr
      .on("onload", () => callback(null, xhr.response))
      .on("error", (error) => callback(error));

    xhr.send();

    fetch("https://api.github.com");
    ```

  - Timer 定时器

    ```typescript
    setTimeout(callback, ms);
    setInterval(callback, ms);

    // nodejs
    setImmediate(callback);
    ```

  - GPU
  - Storage Engines
  - File System
  - Sensors 传感器
  - Camera 相机
  - Microphone 麦克风
  - Geolocation 地理位置

- Callback-based APIs 基于回调的 API

  ```typescript
  navigator.geolocation.getCurrentPosition(
    (position) => console.log(position),
    (error) => console.error(error)
  );
  ```

- Promisifying callback-based APIS

  ```typescript
  function getCurrentPosition() {
    return new Promise((resolve, reject) =>
      navigator.geolocation.getCurrentPosition(resolve, reject)
    );
  }

  function delay(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms);)
  }

  Promise.all(iterable).then(values => console.log(values)).catch(error => console.error(error))
  Promise.any(iterable).then(value => console.log(value)).catch(error => console.error(error.errors))
  Promise.race(iterable).then(value => console.log(value)).catch(error => console.error(error))
  Promise.allSettled(iterable).forEach(result => console.log(result))

  Promise.resolve()
  Promise.reject()
  ```

## 协程 Coroutine

> 定义：协程是用户态的轻量级线程，可以在执行过程中暂停和恢复

| 方法                 | 等待条件     | 失败条件     | 返回值         | 适用场景     |
| -------------------- | ------------ | ------------ | -------------- | ------------ |
| `Promise.all`        | 所有成功     | 任意一个失败 | 所有结果数组   | 全部必须成功 |
| `Promise.any`        | 任意一个成功 | 所有失败     | 第一个成功结果 | 只需一个成功 |
| `Promise.allSettled` | 所有完成     | 永不失败     | 所有状态和结果 | 需要所有结果 |
| `Promise.race`       | 任意一个完成 | 第一个失败   | 第一个完成结果 | 竞速场景     |

### 详细解析和示例

#### 1. Promise.all() - "全部成功才成功"

**特点：**

- 等待所有 Promise 成功完成
- 任何一个失败就立即失败
- 保持结果顺序

```javascript
// 模拟异步操作
function fetchUser(id, delay, shouldFail = false) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (shouldFail) {
        reject(new Error(`用户 ${id} 获取失败`));
      } else {
        resolve({ id, name: `用户${id}` });
      }
    }, delay);
  });
}

// Promise.all 示例
async function testPromiseAll() {
  console.log("=== Promise.all 测试 ===");

  try {
    // 所有成功的情况
    const results = await Promise.all([
      fetchUser(1, 1000),
      fetchUser(2, 2000),
      fetchUser(3, 500),
    ]);
    console.log("所有用户获取成功:", results);
    // 结果按照原始顺序：[用户1, 用户2, 用户3]
  } catch (error) {
    console.error("Promise.all 失败:", error.message);
  }

  try {
    // 有一个失败的情况
    const results = await Promise.all([
      fetchUser(1, 1000),
      fetchUser(2, 2000, true), // 这个会失败
      fetchUser(3, 500),
    ]);
  } catch (error) {
    console.error("Promise.all 失败:", error.message);
    // 即使用户1和用户3成功，整体也失败
  }
}
```

#### 2. Promise.any() - "任意一个成功就成功"

**特点：**

- 只要有一个成功就成功
- 所有都失败才失败
- 返回第一个成功的结果

```javascript
async function testPromiseAny() {
  console.log("=== Promise.any 测试 ===");

  try {
    // 至少一个成功的情况
    const result = await Promise.any([
      fetchUser(1, 3000, true), // 失败
      fetchUser(2, 1000), // 最快成功
      fetchUser(3, 2000), // 也会成功但较慢
    ]);
    console.log("第一个成功的用户:", result);
    // 只返回用户2，因为它最快成功
  } catch (error) {
    console.error("Promise.any 失败:", error);
  }

  try {
    // 所有都失败的情况
    const result = await Promise.any([
      fetchUser(1, 1000, true),
      fetchUser(2, 2000, true),
      fetchUser(3, 3000, true),
    ]);
  } catch (error) {
    console.error("Promise.any 所有都失败:", error.name);
    // AggregateError: 包含所有失败原因
    console.log("失败详情:", error.errors);
  }
}
```

#### 3. Promise.allSettled() - "等待所有完成，不管成败"

**特点：**

- 等待所有 Promise 完成（成功或失败）
- 永远不会失败
- 返回每个 Promise 的状态和结果

```javascript
async function testPromiseAllSettled() {
  console.log("=== Promise.allSettled 测试 ===");

  const results = await Promise.allSettled([
    fetchUser(1, 1000), // 成功
    fetchUser(2, 2000, true), // 失败
    fetchUser(3, 1500), // 成功
    fetchUser(4, 500, true), // 失败
  ]);

  console.log("所有操作完成:");
  results.forEach((result, index) => {
    if (result.status === "fulfilled") {
      console.log(`操作 ${index + 1} 成功:`, result.value);
    } else {
      console.log(`操作 ${index + 1} 失败:`, result.reason.message);
    }
  });

  // 分离成功和失败的结果
  const successful = results
    .filter((result) => result.status === "fulfilled")
    .map((result) => result.value);

  const failed = results
    .filter((result) => result.status === "rejected")
    .map((result) => result.reason);

  console.log("成功的操作:", successful);
  console.log(
    "失败的操作:",
    failed.map((err) => err.message)
  );
}
```

#### 4. Promise.race() - "第一个完成就完成"

**特点：**

- 返回第一个完成的 Promise 结果
- 可能是成功也可能是失败
- 常用于超时控制

```javascript
async function testPromiseRace() {
  console.log("=== Promise.race 测试 ===");

  try {
    // 正常竞速
    const result = await Promise.race([
      fetchUser(1, 2000),
      fetchUser(2, 1000), // 最快
      fetchUser(3, 3000),
    ]);
    console.log("最快完成的用户:", result);
  } catch (error) {
    console.error("最快完成的是失败:", error.message);
  }

  // 超时控制示例
  function withTimeout(promise, timeout) {
    const timeoutPromise = new Promise((_, reject) => {
      setTimeout(() => {
        reject(new Error(`操作超时 ${timeout}ms`));
      }, timeout);
    });

    return Promise.race([promise, timeoutPromise]);
  }

  try {
    const result = await withTimeout(fetchUser(1, 3000), 2000);
    console.log("在超时前完成:", result);
  } catch (error) {
    console.error("操作超时:", error.message);
  }
}
```

### 实际应用场景

#### 1. Promise.all - 批量数据获取

```javascript
// 需要所有数据都获取成功才能渲染页面
async function loadPageData() {
  try {
    const [userInfo, orders, settings] = await Promise.all([
      fetchUserInfo(),
      fetchUserOrders(),
      fetchUserSettings(),
    ]);

    renderPage({ userInfo, orders, settings });
  } catch (error) {
    showErrorPage(error);
  }
}
```

#### 2. Promise.any - 多服务器请求

```javascript
// 从多个镜像服务器获取数据，任意一个成功即可
async function fetchFromMirrors(path) {
  const servers = [
    "https://api1.example.com",
    "https://api2.example.com",
    "https://api3.example.com",
  ];

  try {
    const result = await Promise.any(
      servers.map((server) => fetch(`${server}${path}`))
    );
    return await result.json();
  } catch (error) {
    throw new Error("所有服务器都无法访问");
  }
}
```

#### 3. Promise.allSettled - 批量操作总结

```javascript
// 批量删除文件，需要知道每个文件的删除结果
async function batchDeleteFiles(filePaths) {
  const results = await Promise.allSettled(
    filePaths.map((path) => fs.unlink(path))
  );

  const summary = {
    successful: [],
    failed: [],
  };

  results.forEach((result, index) => {
    if (result.status === "fulfilled") {
      summary.successful.push(filePaths[index]);
    } else {
      summary.failed.push({
        file: filePaths[index],
        error: result.reason.message,
      });
    }
  });

  return summary;
}
```

#### 4. Promise.race - 超时和取消

```javascript
// 带超时的网络请求
function fetchWithTimeout(url, timeout = 5000) {
  const controller = new AbortController();

  const fetchPromise = fetch(url, {
    signal: controller.signal,
  });

  const timeoutPromise = new Promise((_, reject) => {
    setTimeout(() => {
      controller.abort();
      reject(new Error("请求超时"));
    }, timeout);
  });

  return Promise.race([fetchPromise, timeoutPromise]);
}
```

## 线程 Threads

> 线程是 CPU 调度的基本单位，同一进程内的线程共享内存空间

```typescript
const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");

if (isMainThread) {
  // 可以共享的类型
  const sharedBuffer = new SharedArrayBuffer(4 * 1024); // 4KB
  const sharedArray = new Int32Array(sharedBuffer);

  // 不能共享的类型：
  const normalArray = [1, 2, 3];
  const normalObject = { key: "value" };

  // 创建工作线程
  const worker = new Worker(__filename, {
    workerData: { sharedArray, normalArray, normalObject },
  });

  // 写入数据
  sharedArray[0] = 42;

  // 监听工作线程消息
  worker.on("message", (message) => {
    console.log("主线程收到消息:", message);
    console.log("共享内存值:", sharedArray[0]);
  });
} else {
  const { sharedArray } = workerData;

  // 读取共享内存
  console.log("工作线程读取共享内存:", sharedArray[0]);

  // 修改共享内存
  sharedArray[0] = 100;

  // 通知主线程
  parentPort.postMessage("共享内存已更新");
}
```

## 进程 Process

> 进程是操作系统资源分配的基本单位，拥有独立的内存空间

### 子进程

```typescript
const { exec, spawn, fork } = require("child_process");

// exec - 用于执行简单命令，适合小数据量
exec("ls -la", (error, stdout, stderr) => {
  if (error) {
    console.error(`执行错误: ${error}`);
    return;
  }
  console.log(`标准输出: ${stdout}`);
});

// spawn - 用于启动子进程，适合处理大量数据
const ls = spawn("ls", ["-lh", "/usr"], {
  cwd: process.cwd(), // 工作目录
  env: { ...process.env, NODE_ENV: "development" }, // 环境变量
  stdio: ["pipe", "pipe", "pipe"], // array|string 标准输入输出配置，pipe 创建管道 ignore 忽略 inherit 继承父进程 ipc 进程间通信
  shell: true, // 是否使用 shell
  detached: true, // 是否分离进程
});

let output = "";
let error = "";

ls.stdout.on("data", (data) => {
  output += data;
});

ls.stderr.on("data", (data) => {
  error += data;
});

ls.on("close", (code) => {
  if (code === 0) {
    console.log(output);
  } else {
    console.log(error);
  }
}).on("error", (error) => {
  console.error(`子进程启动失败: ${error}`);
});

// fork - 专门用于创建 Node.js 子进程
const child = fork("child.js");

child
  .on("message", (message) => {
    if (message.type === "result") {
      console.log("收到子进程计算结果:", message.data);
    }
  })
  .on("exit", (code, signal) => {
    if (code !== 0) {
      console.error(`子进程异常退出，退出码 ${code}`);
    }
  })
  .on("error", (error) => {
    console.error("子进程错误:", error);
  });

child.send({ type: "compute", data: [1, 2, 3, 4, 5] });
```

```typescript
// child process
process.on("message", (message) => {
  if (message.type === "compute") {
    console.log("收到父进程数据:", message.data);

    const result = message.data.reduce((a, b) => a + b, 0);

    process.send({ type: "result", data: result });
  }
});
```

### 集群

```typescript
const cluster = require("cluster");
const http = require("http");
const numCPUs = require("os").cpus().length;

if (cluster.isMaster) {
  console.log(`主进程 ${process.pid} 正在运行`);

  // 根据 CPU 核心数创建子进程
  for (let i = 0; i < numCPUs; i++) {
    const worker = cluster.fork();

    // 监听子进程消息
    worker.on("message", (message) => {
      console.log(`收到来自工作进程 ${worker.process.pid} 的消息:`, message);
    });

    // 发送消息到子进程
    worker.send({ type: "task", data: "some data" });
  }

  cluster.on("exit", (worker, code, signal) => {
    if (code !== 0) {
      console.error(`工作进程 ${worker.process.pid} 异常退出，退出码 ${code}`);
    }
  });
} else {
  http
    .createServer((req, res) => {
      res.writeHead(200);
      res.end("Hello World\n");
    })
    .listen(3000);

  console.log(`工作进程 ${process.pid} 启动`);

  // 监听主进程消息
  process.on("message", (message) => {
    console.log("收到主进程消息:", message);

    // 处理消息
    if (message.type === "task") {
      process.send({ type: "result", data: `处理结果: ${message.data}` });
    }
  });
}
```

## 三者对比

| 特性     | 进程         | 线程         | 协程         |
| -------- | ------------ | ------------ | ------------ |
| 内存     | 独立内存空间 | 共享进程内存 | 共享线程内存 |
| 创建开销 | 大           | 中等         | 小           |
| 切换开销 | 大           | 中等         | 很小         |
| 通信方式 | IPC          | 共享内存     | 函数调用     |
| 安全性   | 高           | 中等         | 低           |
| 并发模型 | 抢占式       | 抢占式       | 协作式       |

## os

- pm2

  ```bash
  pm2 start <filename> --name <name>
  pm2 stop <name>
  pm2 restart <name>
  pm2 delete <name>

  pm2 list
  pm2 monit
  pm2 logs <name>
  pm2 info <name>

  pm2 pid <name>
  ```

- exit

  ```typescript
  process.exit(1);
  process.kill(process.pid, "SIGTERM");
  ```

- env

  ```typescript
  // PORT=3000 node index.js
  process.env.PORT;
  ```

- argv

  ```typescript
  // node index.js key=value
  process.argv.slice(2);
  ```
