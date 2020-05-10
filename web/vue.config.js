module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? './static'
    : '/',
  devServer: {
    proxy: 'http://localhost:7001'
  }
}
