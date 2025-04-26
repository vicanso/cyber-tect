use std::future::Future;
use std::io;
use std::net::SocketAddr;
use std::pin::Pin;
use std::task::{Context, Poll};
use tokio::net::lookup_host;
use tower::Service;

/// 自定义DNS解析器
#[derive(Clone)]
pub struct CustomResolver;

impl CustomResolver {
    /// 创建新的解析器实例
    pub fn new() -> Self {
        CustomResolver
    }
}

impl Service<String> for CustomResolver {
    type Response = std::vec::IntoIter<SocketAddr>;
    type Error = io::Error;
    type Future = Pin<Box<dyn Future<Output = Result<Self::Response, Self::Error>> + Send>>;

    fn poll_ready(&mut self, _cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        Poll::Ready(Ok(()))
    }

    fn call(&mut self, name: String) -> Self::Future {
        Box::pin(async move {
            let host = name.as_str();
            let port = 0;
            let socket_addrs = lookup_host((host, port)).await?;
            Ok(socket_addrs.collect::<Vec<_>>().into_iter())
        })
    }
}
