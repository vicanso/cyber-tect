module.exports = {
  publicPath: "./static",
  devServer: {
    proxy: 'http://localhost:7001'
  }
}
