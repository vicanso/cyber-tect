use hyper::client::connect::dns::Name;
use std::future::Future;
use std::io;
use std::net::SocketAddr;
use std::pin::Pin;
use std::task::{Context, Poll};
use tokio::net::lookup_host;
use tower::Service;

/// 自定义DNS解析器
#[derive(Clone, Default)]
pub struct CustomResolver;

impl CustomResolver {
    /// 创建新的解析器实例
    pub fn new() -> Self {
        CustomResolver
    }
}

impl Service<Name> for CustomResolver {
    type Response = std::vec::IntoIter<SocketAddr>;
    type Error = io::Error;
    type Future = Pin<Box<dyn Future<Output = Result<Self::Response, Self::Error>> + Send>>;

    fn poll_ready(&mut self, _cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        Poll::Ready(Ok(()))
    }

    fn call(&mut self, name: Name) -> Self::Future {
        Box::pin(async move {
            let host = name.as_str();
            println!("Resolving host: {}", host);
            let addrs = lookup_host((host, 0)).await?.collect::<Vec<_>>();
            println!("Resolved addresses: {:?}", addrs);
            Ok(addrs.into_iter())
        })
    }
}
