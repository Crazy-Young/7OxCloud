const { defineConfig } = require('@vue/cli-service')
const path = require('path')
const CompressionPlugin = require('compression-webpack-plugin')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  productionSourceMap: false,
  publicPath: './',
  outputDir: 'dist',
  assetsDir: 'static',
  configureWebpack: config => {
    config.plugins.push(new CompressionPlugin({
      test: /\.js$|\.html$|\.css/,
      threshold: 10240,
      deleteOriginalAssets: false
    }))
  },
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'less',
      patterns: [
        path.resolve(__dirname, './src/less/color.less'),
      ]
    }
  },
  devServer: {
    open: true,
    host: '0.0.0.0',
    proxy: {
      '/api': {
        target: 'http://152.136.156.247/api/',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  },
})
