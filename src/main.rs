use axum::{routing::get, Router};
use hyper::{Body, Request};
use hyper_tls::HttpsConnector;
use hyper_util::client::legacy::{Client, HttpConnector};
use std::sync::Arc;
use std::time::Instant;
use tower_http::services::ServeDir;

mod dns;
mod trace_connector;

use crate::dns::CustomResolver;
use trace_connector::TracingConnector;

async fn hello_world() -> String {
    let start_time = Instant::now();

    // 创建带追踪的 HTTPS connector
    let mut http = HttpConnector::new_with_resolver(CustomResolver::new());
    http.enforce_http(false);
    let https = HttpsConnector::new_with_connector(http);
    let traced_connector = TracingConnector::new(https);

    // 使用 hyper::Client 替代 reqwest::Client
    let client = Client::builder().build(traced_connector);

    println!("Client creation took: {:?}", start_time.elapsed());

    // 发送请求
    let res = client
        .request(
            hyper::Request::builder()
                .uri("https://www.baidu.com")
                .body(Body::empty())
                .unwrap(),
        )
        .await
        .unwrap();

    // 读取响应体
    let body_bytes = hyper::body::to_bytes(res.into_body()).await.unwrap();
    let body = String::from_utf8_lossy(&body_bytes).to_string();

    println!("Total time: {:?}", start_time.elapsed());

    body
}

// async fn setup_client() {
//     // 创建使用自定义解析器的connector
//     let mut connector = HttpConnector::new_with_resolver(CustomResolver::new());

//     // 构建HTTP client
//     let client: Client<HttpConnector> = Client::builder().build(connector);

//     // 使用client...
// }

#[shuttle_runtime::main]
async fn main() -> shuttle_axum::ShuttleAxum {
    // ServeDir falls back to serve index.html when requesting a directory
    let router = Router::new()
        .route("/", get(hello_world))
        .nest_service("/assets", ServeDir::new("assets"));

    Ok(router.into())
}
