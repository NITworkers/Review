// pathモジュールを読み(output.pathに絶対パスを指定するため)
const path = require('path');

module.exports = {
    // モードの設定。指定可能な値は、none, development ,production（デフォルト）
    // https://webpack.js.org/concepts/mode/#mode-development
    mode: 'development',
    // アプリケーションが実行を開始されるポイント(エントリーポイント)
    // 配列で指定すると、すべての項目が実行される
    // https://webpack.js.org/configuration/entry-context/#entry
    entry: {
      signin: './Pages/js/render/R_signin.jsx',
      signup: './Pages/js/render/R_signup.jsx'
    },
    output: {
        filename: '[name].bundle.js',
        // ビルド後のファイルが出力される"絶対パス"ディレクトリ
        // https://webpack.js.org/configuration/output/#outputpath
        path: path.join(__dirname, 'Pages/js/build')
    },
    module: {
      rules: [
        {
          test: /\.(js|jsx)/,
          exclude: /node_modules/,
          loader: 'babel-loader',
          options: {
            presets: [
              ['@babel/react'],
              ['@babel/env',
                {
                  "targets": {
                      "ie": 11
                  },
                  "useBuiltIns": "usage"
                }
              ]
            ]
          }
        },
        { test: /\.css$/, loader: "style-loader!css-loader" },
        { test: /\.(png|jpg|jpeg|gif|woff)$/, loader: "url-loader" },
        { test: /\.(otf|eot|ttf)$/, loader: "file?prefix=font/" },
        { test: /\.svg$/, loader: "file" }
      ]
    },
    devServer: {
      contentBase: path.join(__dirname, 'Pages'),
      hot: true
    },
};