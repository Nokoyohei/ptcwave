const { VueLoaderPlugin } = require('vue-loader')
module.exports = {
  mode: "development",

  devtool: 'inline-source-map',
  entry: "./src/index.ts",

  output: {
    path: `${__dirname}/dist`,
    filename: "bundle.js"
  },
  module: {
    rules: [
      {
        //vueのloader設定
        test: /\.vue$/,
        loader: "vue-loader"
      },
      {
        //cssのloader設定
        test: /\.css$/,
        use:[
          "style-loader",
          "css-loader"
        ]
      },
      {
        //typescriptのloaderを設定
        test: /\.ts$/,
        loader: "ts-loader",
        options: {
          //vueをtypescriptとして監視する
          appendTsSuffixTo: [/\.vue$/]
        }
      },
    ]
  },
  resolve: {
    //import文で、.tsを省略できるようにする
    extensions: [".ts"]
  },
  plugins:[new VueLoaderPlugin()]
}