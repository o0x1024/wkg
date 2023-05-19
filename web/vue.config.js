const CompressionPlugin = require('compression-webpack-plugin')


module.exports = {
  outputDir: process.env.outputDir,
  lintOnSave: false,

  productionSourceMap: false,
  devServer: {
    proxy: {
      "/api": {
        target: "http://127.0.0.1:7788",
        changeOrigin: true,
        //   ws: true,
        // rewrite:(path) => path.replace(/^\/api/,'')
        pathRewrite: {
          "^/api": "",
        },
      },
    },
  },


  configureWebpack: config => {
    return {
      plugins: [
        new CompressionPlugin({
          algorithm: 'gzip', // 使用gzip压缩
          test: /\.js$|\.html$|\.css$/, // 匹配文件名
          minRatio: 1, // 压缩率小于1才会压缩
          threshold: 10240, // 对超过10k的数据压缩
          deleteOriginalAssets: false,
        })
      ]
    }
  },
}
