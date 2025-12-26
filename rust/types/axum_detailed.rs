// Axum 详解 - Web 框架完全指南
//
// Axum 是 Rust 中现代化的 Web 框架
// 本教程涵盖所有核心功能和实战技巧

use axum::{
    body::Body,
    extract::{Json, Path, Query, State},
    http::{HeaderMap, HeaderValue, StatusCode},
    response::{Html, IntoResponse, Response},
    routing::{get, post},
    Router,
};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use tower_http::cors::CorsLayer;

#[tokio::main]
async fn main() {
    println!("=== Axum Web 框架详解 ===\n");

    // 注意：这个教程展示 Axum 的各种功能
    // 实际运行会启动多个服务器示例

    demo_overview();

    println!("\n各个示例服务器:");
    println!("  1. 基础路由     - cargo run --bin axum_detailed");
    println!("  2. 提取器示例   - 见下方代码");
    println!("  3. 状态管理     - 见下方代码");
    println!("  4. 完整 API     - 见下方代码");

    // 启动基础示例服务器
    run_basic_server().await;
}

// ============================================
// 概览
// ============================================
fn demo_overview() {
    println!("--- Axum 概览 ---\n");

    println!("什么是 Axum？");
    println!("  - 基于 Tokio 和 Hyper 的 Web 框架");
    println!("  - 类型安全的提取器系统");
    println!("  - 零成本抽象");
    println!("  - 与 Tower 生态集成\n");

    println!("核心特性:");
    println!("  📌 路由系统");
    println!("  📌 提取器 (Extractors)");
    println!("  📌 响应类型");
    println!("  📌 中间件");
    println!("  📌 状态管理");
    println!("  📌 错误处理\n");

    println!("基础用法:");
    println!("  let app = Router::new()");
    println!("      .route(\"/\", get(handler))");
    println!("      .route(\"/users/:id\", get(get_user));");
    println!();
}

// ============================================
// 1. 基础路由
// ============================================
async fn run_basic_server() {
    println!("--- 1. 基础路由示例 ---\n");

    // 构建路由
    let app = Router::new()
        .route("/", get(root_handler))
        .route("/hello", get(hello_handler))
        .route("/users/:id", get(user_handler))
        .route("/search", get(search_handler));

    println!("  服务器启动在 http://localhost:3000");
    println!("  端点:");
    println!("    GET  /");
    println!("    GET  /hello");
    println!("    GET  /users/:id");
    println!("    GET  /search?q=...");
    println!();

    println!("  按 Ctrl+C 停止服务器");
    println!();

    // 启动服务器
    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    axum::serve(listener, app).await.unwrap();
}

// 根路径处理器
async fn root_handler() -> &'static str {
    "Welcome to Axum!"
}

// Hello 处理器
async fn hello_handler() -> Html<&'static str> {
    Html("<h1>Hello, Axum!</h1>")
}

// 用户处理器 - 路径参数
async fn user_handler(Path(user_id): Path<u32>) -> String {
    format!("User ID: {}", user_id)
}

// 搜索处理器 - 查询参数
#[derive(Deserialize)]
struct SearchParams {
    q: String,
    #[serde(default)]
    page: u32,
}

async fn search_handler(Query(params): Query<SearchParams>) -> String {
    format!("Searching for: {} (page {})", params.q, params.page)
}

// ============================================
// 2. HTTP 方法路由
// ============================================

// 所有 HTTP 方法示例
fn http_methods_example() -> Router {
    use axum::routing::{delete, patch, put};

    Router::new()
        .route("/items", get(get_items).post(create_item))
        .route(
            "/items/:id",
            get(get_item)
                .put(update_item)
                .patch(patch_item)
                .delete(delete_item),
        )
}

async fn get_items() -> &'static str {
    "GET /items"
}

async fn create_item() -> &'static str {
    "POST /items"
}

async fn get_item(Path(id): Path<u32>) -> String {
    format!("GET /items/{}", id)
}

async fn update_item(Path(id): Path<u32>) -> String {
    format!("PUT /items/{}", id)
}

async fn patch_item(Path(id): Path<u32>) -> String {
    format!("PATCH /items/{}", id)
}

async fn delete_item(Path(id): Path<u32>) -> String {
    format!("DELETE /items/{}", id)
}

// ============================================
// 3. 提取器 (Extractors)
// ============================================

// Path 提取器
async fn path_extractor_single(Path(id): Path<u32>) -> String {
    format!("Single param: {}", id)
}

async fn path_extractor_multiple(Path((user_id, post_id)): Path<(u32, u32)>) -> String {
    format!("User: {}, Post: {}", user_id, post_id)
}

// Query 提取器
#[derive(Deserialize)]
struct Pagination {
    #[serde(default = "default_page")]
    page: u32,
    #[serde(default = "default_limit")]
    limit: u32,
}

fn default_page() -> u32 {
    1
}
fn default_limit() -> u32 {
    20
}

async fn query_extractor(Query(params): Query<Pagination>) -> String {
    format!("Page: {}, Limit: {}", params.page, params.limit)
}

// JSON 提取器
#[derive(Debug, Deserialize, Serialize)]
struct CreateUser {
    name: String,
    email: String,
    age: u32,
}

async fn json_extractor(Json(payload): Json<CreateUser>) -> Json<CreateUser> {
    // 自动序列化和反序列化
    Json(payload)
}

// Headers 提取器
async fn headers_extractor(headers: HeaderMap) -> String {
    let user_agent = headers
        .get("user-agent")
        .and_then(|v| v.to_str().ok())
        .unwrap_or("Unknown");

    format!("User-Agent: {}", user_agent)
}

// 组合提取器
async fn combined_extractors(
    Path(id): Path<u32>,
    Query(params): Query<Pagination>,
    headers: HeaderMap,
    Json(payload): Json<CreateUser>,
) -> impl IntoResponse {
    // 可以同时使用多个提取器
    let user_agent = headers
        .get("user-agent")
        .and_then(|v| v.to_str().ok())
        .unwrap_or("Unknown");

    (
        StatusCode::CREATED,
        Json(serde_json::json!({
            "id": id,
            "page": params.page,
            "user_agent": user_agent,
            "data": payload,
        })),
    )
}

// ============================================
// 4. 响应类型
// ============================================

// 纯文本响应
async fn text_response() -> &'static str {
    "Plain text"
}

// HTML 响应
async fn html_response() -> Html<&'static str> {
    Html("<h1>HTML Response</h1>")
}

// JSON 响应
#[derive(Serialize)]
struct ApiResponse {
    message: String,
    code: u32,
}

async fn json_response() -> Json<ApiResponse> {
    Json(ApiResponse {
        message: "Success".to_string(),
        code: 200,
    })
}

// 状态码 + JSON
async fn status_json_response() -> (StatusCode, Json<ApiResponse>) {
    (
        StatusCode::CREATED,
        Json(ApiResponse {
            message: "Created".to_string(),
            code: 201,
        }),
    )
}

// Headers + Body
async fn headers_response() -> (HeaderMap, &'static str) {
    let mut headers = HeaderMap::new();
    headers.insert(
        "X-Custom-Header",
        "custom-value".parse::<HeaderValue>().unwrap(),
    );

    (headers, "Response with custom headers")
}

// 自定义响应
async fn custom_response() -> Response {
    Response::builder()
        .status(StatusCode::OK)
        .header("X-Custom", "value")
        .body(Body::from("Custom response"))
        .unwrap()
}

// Result 响应
async fn result_response() -> Result<Json<ApiResponse>, StatusCode> {
    // 可以返回 Result
    Ok(Json(ApiResponse {
        message: "Success".to_string(),
        code: 200,
    }))
}

// ============================================
// 5. 状态管理
// ============================================

// 共享状态
#[derive(Clone)]
struct AppState {
    db: Arc<Mutex<HashMap<u32, String>>>,
    counter: Arc<Mutex<u32>>,
}

fn state_example() -> Router {
    // 初始化状态
    let state = AppState {
        db: Arc::new(Mutex::new(HashMap::new())),
        counter: Arc::new(Mutex::new(0)),
    };

    Router::new()
        .route("/counter", get(get_counter).post(increment_counter))
        .route("/data/:key", get(get_data).post(set_data))
        .with_state(state)
}

async fn get_counter(State(state): State<AppState>) -> String {
    let count = state.counter.lock().unwrap();
    format!("Counter: {}", *count)
}

async fn increment_counter(State(state): State<AppState>) -> String {
    let mut count = state.counter.lock().unwrap();
    *count += 1;
    format!("Counter: {}", *count)
}

async fn get_data(
    State(state): State<AppState>,
    Path(key): Path<u32>,
) -> Result<String, StatusCode> {
    let db = state.db.lock().unwrap();
    db.get(&key).cloned().ok_or(StatusCode::NOT_FOUND)
}

async fn set_data(State(state): State<AppState>, Path(key): Path<u32>, body: String) -> StatusCode {
    let mut db = state.db.lock().unwrap();
    db.insert(key, body);
    StatusCode::CREATED
}

// ============================================
// 6. 错误处理
// ============================================

// 自定义错误类型
enum AppError {
    NotFound,
    BadRequest(String),
    InternalError,
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, message) = match self {
            AppError::NotFound => (StatusCode::NOT_FOUND, "Not Found"),
            AppError::BadRequest(msg) => {
                return (StatusCode::BAD_REQUEST, msg).into_response();
            }
            AppError::InternalError => (StatusCode::INTERNAL_SERVER_ERROR, "Internal Server Error"),
        };

        (status, message).into_response()
    }
}

async fn error_handler() -> Result<String, AppError> {
    // 可以返回自定义错误
    Err(AppError::NotFound)
}

// ============================================
// 7. 中间件
// ============================================

fn middleware_example() -> Router {
    use tower_http::trace::TraceLayer;

    Router::new()
        .route("/", get(root_handler))
        // 添加 CORS
        .layer(CorsLayer::permissive())
        // 添加日志追踪
        .layer(TraceLayer::new_for_http())
}

// ============================================
// 8. 嵌套路由
// ============================================

fn nested_routes_example() -> Router {
    // API 路由
    let api_routes = Router::new()
        .route("/users", get(get_users).post(create_user_handler))
        .route(
            "/users/:id",
            get(get_user_handler).delete(delete_user_handler),
        )
        .route("/posts", get(get_posts).post(create_post_handler));

    // 主路由
    Router::new()
        .route("/", get(root_handler))
        .route("/health", get(health_check))
        .nest("/api/v1", api_routes)
}

async fn get_users() -> &'static str {
    "GET /api/v1/users"
}

async fn create_user_handler() -> &'static str {
    "POST /api/v1/users"
}

async fn get_user_handler(Path(id): Path<u32>) -> String {
    format!("GET /api/v1/users/{}", id)
}

async fn delete_user_handler(Path(id): Path<u32>) -> String {
    format!("DELETE /api/v1/users/{}", id)
}

async fn get_posts() -> &'static str {
    "GET /api/v1/posts"
}

async fn create_post_handler() -> &'static str {
    "POST /api/v1/posts"
}

async fn health_check() -> &'static str {
    "OK"
}

// ============================================
// 9. 完整 REST API 示例
// ============================================

// 数据模型
#[derive(Debug, Clone, Serialize, Deserialize)]
struct User {
    id: u32,
    name: String,
    email: String,
}

#[derive(Debug, Deserialize)]
struct CreateUserRequest {
    name: String,
    email: String,
}

#[derive(Debug, Deserialize)]
struct UpdateUserRequest {
    name: Option<String>,
    email: Option<String>,
}

// 应用状态
#[derive(Clone)]
struct ApiState {
    users: Arc<Mutex<HashMap<u32, User>>>,
    next_id: Arc<Mutex<u32>>,
}

// 完整 API 路由
fn complete_api() -> Router {
    let state = ApiState {
        users: Arc::new(Mutex::new(HashMap::new())),
        next_id: Arc::new(Mutex::new(1)),
    };

    Router::new()
        // 根路径
        .route("/", get(api_root))
        // 用户 CRUD
        .route("/users", get(list_users).post(create_user_api))
        .route(
            "/users/:id",
            get(get_user_api)
                .put(update_user_api)
                .delete(delete_user_api),
        )
        // 健康检查
        .route("/health", get(health_check))
        // 添加状态
        .with_state(state)
        // 添加 CORS
        .layer(CorsLayer::permissive())
}

// API 根路径
async fn api_root() -> Json<serde_json::Value> {
    Json(serde_json::json!({
        "name": "User API",
        "version": "1.0.0",
        "endpoints": {
            "users": "/users",
            "health": "/health"
        }
    }))
}

// 列出所有用户
async fn list_users(State(state): State<ApiState>) -> Json<Vec<User>> {
    let users = state.users.lock().unwrap();
    let user_list: Vec<User> = users.values().cloned().collect();
    Json(user_list)
}

// 创建用户
async fn create_user_api(
    State(state): State<ApiState>,
    Json(payload): Json<CreateUserRequest>,
) -> (StatusCode, Json<User>) {
    let mut next_id = state.next_id.lock().unwrap();
    let id = *next_id;
    *next_id += 1;

    let user = User {
        id,
        name: payload.name,
        email: payload.email,
    };

    let mut users = state.users.lock().unwrap();
    users.insert(id, user.clone());

    (StatusCode::CREATED, Json(user))
}

// 获取单个用户
async fn get_user_api(
    State(state): State<ApiState>,
    Path(id): Path<u32>,
) -> Result<Json<User>, StatusCode> {
    let users = state.users.lock().unwrap();
    users
        .get(&id)
        .cloned()
        .map(Json)
        .ok_or(StatusCode::NOT_FOUND)
}

// 更新用户
async fn update_user_api(
    State(state): State<ApiState>,
    Path(id): Path<u32>,
    Json(payload): Json<UpdateUserRequest>,
) -> Result<Json<User>, StatusCode> {
    let mut users = state.users.lock().unwrap();

    let user = users.get_mut(&id).ok_or(StatusCode::NOT_FOUND)?;

    if let Some(name) = payload.name {
        user.name = name;
    }
    if let Some(email) = payload.email {
        user.email = email;
    }

    Ok(Json(user.clone()))
}

// 删除用户
async fn delete_user_api(State(state): State<ApiState>, Path(id): Path<u32>) -> StatusCode {
    let mut users = state.users.lock().unwrap();

    if users.remove(&id).is_some() {
        StatusCode::NO_CONTENT
    } else {
        StatusCode::NOT_FOUND
    }
}

/*
=== 总结 ===

1. Axum 核心概念:

   路由:
   - Router::new()
   - route() - 定义路由
   - nest() - 嵌套路由
   - method_router() - 多方法路由

   提取器:
   - Path - 路径参数
   - Query - 查询参数
   - Json - JSON 请求体
   - State - 共享状态
   - Headers - 请求头

2. 响应类型:

   简单:
   - &str - 纯文本
   - String - 字符串
   - Html<T> - HTML
   - Json<T> - JSON

   组合:
   - (StatusCode, Json<T>)
   - (HeaderMap, Body)
   - Result<T, E>
   - Response - 自定义响应

3. 状态管理:

   模式:
   - Arc<Mutex<T>> - 可变状态
   - Arc<T> - 只读状态
   - with_state() - 附加状态
   - State<T> - 提取状态

4. 最佳实践:

   DO:
   ✓ 使用类型安全的提取器
   ✓ 实现 IntoResponse for 自定义错误
   ✓ 使用 Router 组织路由
   ✓ Arc + Mutex 管理共享状态

   DON'T:
   ✗ 在处理器中阻塞
   ✗ 长时间持有锁
   ✗ 忽略错误处理

5. 常用模式:

   REST API:
   - CRUD 操作
   - 状态管理
   - 错误处理
   - 中间件

   路由组织:
   - nest() 分组
   - 版本化 API
   - 模块化处理器

运行示例:
  cargo run --bin axum_detailed

测试 API:
  curl http://localhost:3000/
  curl http://localhost:3000/hello
  curl http://localhost:3000/users/123
  curl http://localhost:3000/search?q=rust
*/
