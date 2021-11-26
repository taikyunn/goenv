Golangで.envファイルを使うためのサンプル。
参考：https://github.com/joho/godotenv

godotenv:.envファイルを読み込むためのライブラリ

使い方はmain.goを参照。

環境によっての使い分け
参考：https://qiita.com/fukumone/items/0313004d60ddb4d92d55
//環境変数GO_ENVは適当につけました
if os.Getenv("GO_ENV") == "" {
    os.Setenv("GO_ENV", "development")
}
