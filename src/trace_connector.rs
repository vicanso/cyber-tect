use hyper::client::connect::{Connected, Connection};
use pin_project_lite::pin_project;
use std::future::Future;
use std::io;
use std::pin::Pin;
use std::task::{Context, Poll};
use std::time::Instant;
use tokio::io::{AsyncRead, AsyncWrite, ReadBuf};
use tower::Service;

#[derive(Clone)]
pub struct TracingConnector<C> {
    inner: C,
}

impl<C> TracingConnector<C> {
    pub fn new(inner: C) -> Self {
        Self { inner }
    }
}

impl<C> Service<hyper::Uri> for TracingConnector<C>
where
    C: Service<hyper::Uri>,
    C::Response: Connection + Send + 'static,
    C::Future: Send + 'static,
    C::Error: Into<Box<dyn std::error::Error + Send + Sync>>,
{
    type Response = TracingConnection<C::Response>;
    type Error = Box<dyn std::error::Error + Send + Sync>;
    type Future = TracingConnecting<C::Future>;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx).map_err(Into::into)
    }

    fn call(&mut self, dst: hyper::Uri) -> Self::Future {
        println!("Starting connection to {}", dst);
        let start = Instant::now();
        let host = dst.host().unwrap_or("unknown").to_string();
        let future = self.inner.call(dst);
        TracingConnecting {
            future,
            start,
            dns_start: start,
            host,
        }
    }
}

impl<C> Service<hyper::client::connect::dns::Name> for TracingConnector<C>
where
    C: Service<hyper::client::connect::dns::Name>,
    C::Response: Connection + Send + 'static,
    C::Future: Send + 'static,
    C::Error: Into<Box<dyn std::error::Error + Send + Sync>>,
{
    type Response = TracingConnection<C::Response>;
    type Error = Box<dyn std::error::Error + Send + Sync>;
    type Future = TracingConnecting<C::Future>;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx).map_err(Into::into)
    }

    fn call(&mut self, name: hyper::client::connect::dns::Name) -> Self::Future {
        println!("Starting DNS resolution for {}", name);
        let start = Instant::now();
        let host = name.as_str().to_string();
        let future = self.inner.call(name);
        TracingConnecting {
            future,
            start,
            dns_start: start,
            host,
        }
    }
}

pin_project! {
    pub struct TracingConnecting<F> {
        #[pin]
        future: F,
        start: Instant,
        dns_start: Instant,
        host: String,
    }
}

impl<F, E, C> Future for TracingConnecting<F>
where
    F: Future<Output = Result<C, E>>,
    E: Into<Box<dyn std::error::Error + Send + Sync>>,
    C: Connection + Send + 'static,
{
    type Output = Result<TracingConnection<C>, Box<dyn std::error::Error + Send + Sync>>;

    fn poll(self: Pin<&mut Self>, cx: &mut Context<'_>) -> Poll<Self::Output> {
        let this = self.project();
        match this.future.poll(cx) {
            Poll::Ready(Ok(conn)) => {
                let dns_time = this.dns_start.elapsed();
                println!("DNS resolution for {} took {:?}", this.host, dns_time);
                println!("Connection established in {:?}", this.start.elapsed());
                Poll::Ready(Ok(TracingConnection {
                    inner: conn,
                    start: *this.start,
                }))
            }
            Poll::Ready(Err(e)) => Poll::Ready(Err(e.into())),
            Poll::Pending => Poll::Pending,
        }
    }
}

pin_project! {
    pub struct TracingConnection<C> {
        #[pin]
        inner: C,
        start: Instant,
    }
}

impl<C: Connection> Connection for TracingConnection<C> {
    fn connected(&self) -> Connected {
        self.inner.connected()
    }
}

impl<C: AsyncRead + Unpin> AsyncRead for TracingConnection<C> {
    fn poll_read(
        self: Pin<&mut Self>,
        cx: &mut Context<'_>,
        buf: &mut ReadBuf<'_>,
    ) -> Poll<io::Result<()>> {
        let this = self.project();
        let old_filled = buf.filled().len();
        match this.inner.poll_read(cx, buf) {
            Poll::Ready(Ok(())) => {
                let n = buf.filled().len() - old_filled;
                println!("Read {} bytes in {:?}", n, this.start.elapsed());
                Poll::Ready(Ok(()))
            }
            other => other,
        }
    }
}

impl<C: AsyncWrite + Unpin> AsyncWrite for TracingConnection<C> {
    fn poll_write(
        self: Pin<&mut Self>,
        cx: &mut Context<'_>,
        buf: &[u8],
    ) -> Poll<io::Result<usize>> {
        let this = self.project();
        match this.inner.poll_write(cx, buf) {
            Poll::Ready(Ok(n)) => {
                println!("Wrote {} bytes in {:?}", n, this.start.elapsed());
                Poll::Ready(Ok(n))
            }
            other => other,
        }
    }

    fn poll_flush(self: Pin<&mut Self>, cx: &mut Context<'_>) -> Poll<io::Result<()>> {
        let this = self.project();
        this.inner.poll_flush(cx)
    }

    fn poll_shutdown(self: Pin<&mut Self>, cx: &mut Context<'_>) -> Poll<io::Result<()>> {
        let this = self.project();
        this.inner.poll_shutdown(cx)
    }
}
