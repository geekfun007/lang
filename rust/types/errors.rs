// 错误处理：Result, Option, 自定义错误类型

use anyhow::{Context, Result as AnyhowResult};
use std::error::Error;
use std::fmt;
use std::fs::File;
use std::io::{self, Read};
use thiserror::Error;

/// # Result 基础
pub fn result_basics_demo() {
    println!("\n=== Result 基础 ===");

    fn divide(a: f64, b: f64) -> Result<f64, String> {
        if b == 0.0 {
            Err(String::from("除数不能为零"))
        } else {
            Ok(a / b)
        }
    }

    // 使用 match 处理
    match divide(10.0, 2.0) {
        Ok(result) => println!("10 / 2 = {}", result),
        Err(e) => println!("错误: {}", e),
    }

    match divide(10.0, 0.0) {
        Ok(result) => println!("10 / 0 = {}", result),
        Err(e) => println!("错误: {}", e),
    }

    // unwrap - 成功返回值，失败则 panic
    let result = divide(10.0, 2.0).unwrap();
    println!("unwrap: {}", result);

    // expect - 带自定义错误消息
    let result = divide(10.0, 2.0).expect("除法失败");
    println!("expect: {}", result);

    // unwrap_or - 提供默认值
    let result = divide(10.0, 0.0).unwrap_or(0.0);
    println!("unwrap_or: {}", result);

    // unwrap_or_else - 提供默认值计算函数
    let result = divide(10.0, 0.0).unwrap_or_else(|_| {
        println!("计算默认值");
        0.0
    });
    println!("unwrap_or_else: {}", result);
}

/// # ? 操作符
pub fn question_mark_demo() {
    println!("\n=== ? 操作符 ===");

    fn divide(a: f64, b: f64) -> Result<f64, String> {
        if b == 0.0 {
            Err(String::from("除数不能为零"))
        } else {
            Ok(a / b)
        }
    }

    fn calculate() -> Result<f64, String> {
        let x = divide(10.0, 2.0)?; // 如果错误，提前返回
        let y = divide(x, 2.0)?;
        let z = divide(y, 2.0)?;
        Ok(z)
    }

    match calculate() {
        Ok(result) => println!("计算结果: {}", result),
        Err(e) => println!("错误: {}", e),
    }

    // 链式调用
    fn chain_calculate() -> Result<f64, String> {
        divide(10.0, 2.0)?
            .to_string()
            .parse()
            .map_err(|e| format!("解析错误: {}", e))
    }

    match chain_calculate() {
        Ok(result) => println!("链式计算: {}", result),
        Err(e) => println!("错误: {}", e),
    }
}

/// # Option 基础
pub fn option_basics_demo() {
    println!("\n=== Option 基础 ===");

    fn find_user(id: u32) -> Option<String> {
        if id == 1 {
            Some(String::from("Alice"))
        } else {
            None
        }
    }

    // 使用 match
    match find_user(1) {
        Some(name) => println!("找到用户: {}", name),
        None => println!("用户不存在"),
    }

    // 使用 if let
    if let Some(name) = find_user(2) {
        println!("找到用户: {}", name);
    } else {
        println!("用户不存在");
    }

    // unwrap_or
    let name = find_user(2).unwrap_or(String::from("Guest"));
    println!("用户名: {}", name);

    // map
    let upper_name = find_user(1).map(|n| n.to_uppercase());
    println!("大写名字: {:?}", upper_name);

    // and_then
    fn get_email(name: String) -> Option<String> {
        if name == "Alice" {
            Some(String::from("alice@example.com"))
        } else {
            None
        }
    }

    let email = find_user(1).and_then(get_email);
    println!("邮箱: {:?}", email);

    // ok_or - 转换为 Result
    let result: Result<String, &str> = find_user(1).ok_or("用户不存在");
    println!("转换为 Result: {:?}", result);
}

/// # 自定义错误类型（手动实现）
pub fn custom_error_demo() {
    println!("\n=== 自定义错误类型 ===");

    #[derive(Debug)]
    enum MathError {
        DivisionByZero,
        NegativeSquareRoot,
        Overflow,
    }

    impl fmt::Display for MathError {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            match self {
                MathError::DivisionByZero => write!(f, "除数不能为零"),
                MathError::NegativeSquareRoot => write!(f, "不能对负数开平方"),
                MathError::Overflow => write!(f, "计算溢出"),
            }
        }
    }

    impl Error for MathError {}

    fn divide(a: f64, b: f64) -> Result<f64, MathError> {
        if b == 0.0 {
            Err(MathError::DivisionByZero)
        } else {
            Ok(a / b)
        }
    }

    fn sqrt(x: f64) -> Result<f64, MathError> {
        if x < 0.0 {
            Err(MathError::NegativeSquareRoot)
        } else {
            Ok(x.sqrt())
        }
    }

    match divide(10.0, 0.0) {
        Ok(result) => println!("结果: {}", result),
        Err(e) => println!("错误: {}", e),
    }

    match sqrt(-4.0) {
        Ok(result) => println!("结果: {}", result),
        Err(e) => println!("错误: {}", e),
    }
}

fn demo_ok() {
    println!("--- 2. .ok() - Result → Option ---\n");

    println!("作用: 将 Result<T, E> 转换为 Option<T>");
    println!("  - Ok(v)  → Some(v)");
    println!("  - Err(_) → None");
    println!("  - 丢弃错误信息，只关心是否成功\n");

    // 基础示例
    println!("基础示例:");
    let success: Result<i32, &str> = Ok(42);
    let failure: Result<i32, &str> = Err("出错了");

    println!("  Ok(42).ok()      = {:?}", success.ok());
    println!("  Err('出错了').ok() = {:?}", failure.ok());
    println!();

    // 实际应用：解析
    println!("实际应用 1: 解析字符串");
    let numbers = vec!["1", "2", "abc", "4", "xyz", "5"];
    let parsed: Vec<i32> = numbers
        .iter()
        .filter_map(|s| s.parse::<i32>().ok())
        .collect();

    println!("  输入: {:?}", numbers);
    println!("  解析成功的: {:?}", parsed);
    println!("  说明: .ok() 配合 filter_map 过滤掉解析失败的");
    println!();

    // 实际应用：可选配置
    println!("实际应用 2: 读取可选配置");
    use std::env;

    let timeout = env::var("TIMEOUT")
        .ok() // Result<String, VarError> → Option<String>
        .and_then(|s| s.parse::<u64>().ok())
        .unwrap_or(30);

    println!("  TIMEOUT 环境变量: {:?}", env::var("TIMEOUT").ok());
    println!("  使用的超时值: {} 秒", timeout);
    println!("  说明: 配置不存在或无效时使用默认值");
    println!();

    // 何时使用 .ok()
    println!("何时使用 .ok():");
    println!("  ✓ 不关心错误的具体内容");
    println!("  ✓ 只需要知道是否成功");
    println!("  ✓ 配合 filter_map 过滤");
    println!("  ✓ 链式调用中间步骤");
    println!();
}

fn demo_ok_or() {
    println!("--- 3. .ok_or() - Option → Result ---\n");

    println!("作用: 将 Option<T> 转换为 Result<T, E>");
    println!("  - Some(v) → Ok(v)");
    println!("  - None    → Err(提供的错误值)");
    println!("  - 必须提供错误值\n");

    // 基础示例
    println!("基础示例:");
    let some_value: Option<i32> = Some(100);
    let none_value: Option<i32> = None;

    println!("  Some(100).ok_or('错误') = {:?}", some_value.ok_or("错误"));
    println!("  None.ok_or('错误')      = {:?}", none_value.ok_or("错误"));
    println!();

    // 实际应用：必需的配置
    println!("实际应用 1: 读取必需的配置");
    use std::collections::HashMap;

    let config: HashMap<&str, &str> = [("host", "localhost"), ("port", "8080")]
        .iter()
        .cloned()
        .collect();

    fn get_config<'a>(config: &'a HashMap<&str, &str>, key: &str) -> Result<&'a str, String> {
        config
            .get(key)
            .copied()
            .ok_or(format!("缺少必需的配置项: {}", key))
    }

    match get_config(&config, "host") {
        Ok(value) => println!("  host = {}", value),
        Err(e) => println!("  错误: {}", e),
    }

    match get_config(&config, "database") {
        Ok(value) => println!("  database = {}", value),
        Err(e) => println!("  错误: {}", e),
    }
    println!();

    // 实际应用：查找
    println!("实际应用 2: 查找用户");

    fn find_user(id: u32) -> Option<String> {
        match id {
            1 => Some("Alice".to_string()),
            2 => Some("Bob".to_string()),
            _ => None,
        }
    }

    let user = find_user(1).ok_or("用户不存在");
    println!("  find_user(1) = {:?}", user);

    let user = find_user(999).ok_or("用户不存在");
    println!("  find_user(999) = {:?}", user);
    println!();

    // 何时使用 .ok_or()
    println!("何时使用 .ok_or():");
    println!("  ✓ 错误值是常量或字面量");
    println!("  ✓ 错误值计算开销很小");
    println!("  ✓ 简单的错误消息");
    println!("  ✓ None 必须被视为错误");
    println!();

    // 注意事项
    println!("⚠️ 注意: .ok_or() 总是计算错误值");
    println!("  即使 Option 是 Some，错误值也会被求值");
    println!("  如果错误值计算昂贵，使用 .ok_or_else()");
    println!();
}

fn demo_ok_or_else() {
    println!("--- 4. .ok_or_else() - 惰性求值 ---\n");

    println!("作用: 将 Option<T> 转换为 Result<T, E>（惰性）");
    println!("  - Some(v) → Ok(v)");
    println!("  - None    → Err(调用函数生成错误)");
    println!("  - 仅在 None 时才计算错误\n");

    // 基础示例
    println!("基础示例:");
    let some_value: Option<i32> = Some(100);
    let none_value: Option<i32> = None;

    println!(
        "  Some(100).ok_or_else(|| '错误') = {:?}",
        some_value.ok_or_else(|| "错误")
    );
    println!(
        "  None.ok_or_else(|| '错误')      = {:?}",
        none_value.ok_or_else(|| "错误")
    );
    println!();

    // 演示惰性求值
    println!("演示惰性求值:");

    fn expensive_error() -> String {
        println!("    → 正在计算昂贵的错误消息...");
        "昂贵的错误".to_string()
    }

    println!("\n  使用 .ok_or():");
    let some_opt: Option<i32> = Some(42);
    let _result = some_opt.ok_or(expensive_error());
    println!("    注意: 即使是 Some，expensive_error() 也被调用了");

    println!("\n  使用 .ok_or_else():");
    let some_opt: Option<i32> = Some(42);
    let _result = some_opt.ok_or_else(|| expensive_error());
    println!("    注意: 因为是 Some，expensive_error() 没有被调用");

    println!("\n  使用 .ok_or_else() 处理 None:");
    let none_opt: Option<i32> = None;
    let _result = none_opt.ok_or_else(|| expensive_error());
    println!();

    // 实际应用：动态错误消息
    println!("实际应用 1: 动态错误消息");
    use std::collections::HashMap;

    let users: HashMap<u32, String> = [(1, "Alice".to_string()), (2, "Bob".to_string())]
        .iter()
        .cloned()
        .collect();

    fn find_user_with_error(users: &HashMap<u32, String>, id: u32) -> Result<&String, String> {
        users
            .get(&id)
            .ok_or_else(|| format!("用户 ID {} 不存在", id))
    }

    match find_user_with_error(&users, 1) {
        Ok(name) => println!("  找到: {}", name),
        Err(e) => println!("  错误: {}", e),
    }

    match find_user_with_error(&users, 999) {
        Ok(name) => println!("  找到: {}", name),
        Err(e) => println!("  错误: {}", e),
    }
    println!();

    // 实际应用：上下文错误
    println!("实际应用 2: 携带上下文的错误");

    #[derive(Debug)]
    struct Context {
        operation: String,
        timestamp: u64,
    }

    fn get_value() -> Option<i32> {
        None
    }

    let result = get_value().ok_or_else(|| Context {
        operation: "获取配置值".to_string(),
        timestamp: 1234567890,
    });

    println!("  结果: {:?}", result);
    println!("  说明: 错误包含丰富的上下文信息");
    println!();

    // 何时使用 .ok_or_else()
    println!("何时使用 .ok_or_else():");
    println!("  ✓ 错误值计算开销大");
    println!("  ✓ 需要动态生成错误消息");
    println!("  ✓ 错误需要包含上下文信息");
    println!("  ✓ 性能敏感的代码");
    println!();
}

/// # 使用 thiserror 简化错误定义
pub fn thiserror_demo() {
    println!("\n=== thiserror 简化错误定义 ===");

    #[derive(Error, Debug)]
    enum DataError {
        #[error("网络错误: {0}")]
        Network(String),

        #[error("解析错误")]
        Parse,

        #[error("文件未找到: {path}")]
        FileNotFound { path: String },

        #[error("IO 错误")]
        Io(#[from] io::Error),
    }

    fn read_data(path: &str) -> Result<String, DataError> {
        if path.is_empty() {
            return Err(DataError::FileNotFound {
                path: String::from("(空路径)"),
            });
        }

        // 演示错误转换
        let mut file =
            File::open(path).map_err(|e| DataError::Network(format!("无法打开文件: {}", e)))?;

        let mut contents = String::new();
        file.read_to_string(&mut contents)?; // 自动转换 io::Error

        Ok(contents)
    }

    match read_data("") {
        Ok(data) => println!("数据: {}", data),
        Err(e) => println!("错误: {}", e),
    }
}

/// # 使用 anyhow 简化错误处理
pub fn anyhow_demo() {
    println!("\n=== anyhow 简化错误处理 ===");

    fn process_data(value: &str) -> AnyhowResult<i32> {
        let num: i32 = value.parse().context("无法解析数字")?;

        if num < 0 {
            anyhow::bail!("数字不能为负数");
        }

        Ok(num * 2)
    }

    match process_data("42") {
        Ok(result) => println!("处理结果: {}", result),
        Err(e) => println!("错误: {}", e),
    }

    match process_data("abc") {
        Ok(result) => println!("处理结果: {}", result),
        Err(e) => {
            println!("错误: {}", e);
            // 打印错误链
            for cause in e.chain() {
                println!("  原因: {}", cause);
            }
        }
    }
}

/// # 错误传播
pub fn error_propagation_demo() {
    println!("\n=== 错误传播 ===");

    fn read_number_from_file(path: &str) -> Result<i32, String> {
        let contents = std::fs::read_to_string(path).map_err(|e| format!("读取文件失败: {}", e))?;

        let number: i32 = contents
            .trim()
            .parse()
            .map_err(|e| format!("解析数字失败: {}", e))?;

        Ok(number)
    }

    // 创建测试文件
    let _ = std::fs::write("/tmp/test_number.txt", "42");

    match read_number_from_file("/tmp/test_number.txt") {
        Ok(num) => println!("读取到数字: {}", num),
        Err(e) => println!("错误: {}", e),
    }

    match read_number_from_file("/tmp/nonexistent.txt") {
        Ok(num) => println!("读取到数字: {}", num),
        Err(e) => println!("错误: {}", e),
    }
}

/// # 多种错误类型
pub fn multiple_errors_demo() {
    println!("\n=== 多种错误类型 ===");

    use std::num::ParseIntError;

    // 使用 Box<dyn Error>
    fn process(input: &str) -> Result<i32, Box<dyn Error>> {
        let file_contents = std::fs::read_to_string(input)?;
        let number: i32 = file_contents.trim().parse()?;
        Ok(number)
    }

    // 创建测试文件
    let _ = std::fs::write("/tmp/test_multi.txt", "123");

    match process("/tmp/test_multi.txt") {
        Ok(num) => println!("处理结果: {}", num),
        Err(e) => println!("错误: {}", e),
    }

    // 自定义枚举包装多种错误
    #[derive(Debug)]
    enum AppError {
        Io(io::Error),
        Parse(ParseIntError),
    }

    impl fmt::Display for AppError {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            match self {
                AppError::Io(e) => write!(f, "IO 错误: {}", e),
                AppError::Parse(e) => write!(f, "解析错误: {}", e),
            }
        }
    }

    impl Error for AppError {}

    impl From<io::Error> for AppError {
        fn from(err: io::Error) -> AppError {
            AppError::Io(err)
        }
    }

    impl From<ParseIntError> for AppError {
        fn from(err: ParseIntError) -> AppError {
            AppError::Parse(err)
        }
    }

    fn process_with_custom_error(input: &str) -> Result<i32, AppError> {
        let file_contents = std::fs::read_to_string(input)?;
        let number: i32 = file_contents.trim().parse()?;
        Ok(number)
    }

    match process_with_custom_error("/tmp/test_multi.txt") {
        Ok(num) => println!("自定义错误处理: {}", num),
        Err(e) => println!("错误: {}", e),
    }
}

/// # panic! 和不可恢复错误
pub fn panic_demo() {
    println!("\n=== panic! 和不可恢复错误 ===");

    // panic! 会终止程序
    // panic!("程序崩溃！");

    println!("可以用 panic! 表示不可恢复的错误");
    println!("例如：数组越界、断言失败等");

    // 断言
    let x = 5;
    assert_eq!(x, 5, "x 应该等于 5");
    println!("断言通过");

    // assert!(x == 5, "x 应该等于 5");

    // 在开发中可以使用 panic!
    // 在生产中应该使用 Result
}

/// # 实战示例：配置文件读取
pub fn config_file_demo() {
    println!("\n=== 实战示例：配置文件读取 ===");

    #[derive(Error, Debug)]
    enum ConfigError {
        #[error("配置文件未找到: {0}")]
        FileNotFound(String),

        #[error("配置文件格式错误")]
        InvalidFormat,

        #[error("缺少必需字段: {0}")]
        MissingField(String),
    }

    struct Config {
        host: String,
        port: u16,
    }

    fn load_config(path: &str) -> Result<Config, ConfigError> {
        // 模拟读取配置
        if path.is_empty() {
            return Err(ConfigError::FileNotFound(String::from("(空路径)")));
        }

        // 假设读取到的配置
        let host = Some("localhost");
        let port = Some("8080");

        let host = host
            .ok_or(ConfigError::MissingField(String::from("host")))?
            .to_string();

        let port: u16 = port
            .ok_or(ConfigError::MissingField(String::from("port")))?
            .parse()
            .map_err(|_| ConfigError::InvalidFormat)?;

        Ok(Config { host, port })
    }

    match load_config("config.toml") {
        Ok(config) => println!("配置: {}:{}", config.host, config.port),
        Err(e) => println!("错误: {}", e),
    }

    match load_config("") {
        Ok(config) => println!("配置: {}:{}", config.host, config.port),
        Err(e) => println!("错误: {}", e),
    }
}

/// # 最佳实践
pub fn best_practices_demo() {
    println!("\n=== 错误处理最佳实践 ===");

    println!("1. 使用 Result 而不是 panic");
    println!("   - 可恢复错误用 Result");
    println!("   - 不可恢复错误用 panic");

    println!("\n2. 为库代码定义自定义错误类型");
    println!("   - 使用 thiserror 简化定义");
    println!("   - 实现 Error trait");

    println!("\n3. 为应用程序代码使用 anyhow");
    println!("   - 简化错误处理");
    println!("   - 提供上下文信息");

    println!("\n4. 使用 ? 操作符传播错误");
    println!("   - 简洁且清晰");
    println!("   - 自动类型转换");

    println!("\n5. 提供有意义的错误消息");
    println!("   - 包含上下文信息");
    println!("   - 便于调试");

    println!("\n6. 不要过度使用 unwrap()");
    println!("   - 仅在确定不会失败时使用");
    println!("   - 生产代码避免使用");
}

/// 运行所有错误处理示例
pub fn run_all() {
    println!("\n╔════════════════════════════════════╗");
    println!("║        Rust 错误处理详解           ║");
    println!("╚════════════════════════════════════╝");

    result_basics_demo();
    question_mark_demo();
    option_basics_demo();
    custom_error_demo();
    thiserror_demo();
    anyhow_demo();
    error_propagation_demo();
    multiple_errors_demo();
    panic_demo();
    config_file_demo();
    best_practices_demo();
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_result() {
        fn divide(a: i32, b: i32) -> Result<i32, String> {
            if b == 0 {
                Err(String::from("Division by zero"))
            } else {
                Ok(a / b)
            }
        }

        assert_eq!(divide(10, 2), Ok(5));
        assert!(divide(10, 0).is_err());
    }

    #[test]
    fn test_option() {
        fn find(id: i32) -> Option<String> {
            if id == 1 {
                Some(String::from("found"))
            } else {
                None
            }
        }

        assert_eq!(find(1), Some(String::from("found")));
        assert_eq!(find(2), None);
    }
}
