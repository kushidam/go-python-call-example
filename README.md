# go-python-call-example
GoからPythonを呼び出すシンプルなサンプルプログラム



### `os/exec` パッケージを使用する場合:
*【メリット】*

`os/exec` パッケージを使用すると、外部プログラムやスクリプトをGoプログラムから簡単に呼び出すことができる。

GoとPythonが完全に別々のプロセスとして実行され、標準的なプロセス間通信（stdin, stdout, stderr）を通じておこなう。

外部のPythonスクリプトが出力する結果を標準出力から読み取ることができる。

*【デメリット】*

外部プロセスを起動するオーバーヘッドが発生します。これは、プロセスの起動と終了のためのリソースが追加で必要になることを意味します。

データのやりとりは標準入力と標準出力を通じて行われる。

### `cgo` を使用する場合:
*【メリット】*

cgo を使用すると、GoからC言語の関数を直接呼び出すことができる。
これにより、PythonのAPIを直接使用することが可能

cgo を使用すると、より細かな制御が可能であり、PythonとGoのデータの受け渡しにおいて柔軟性が増す。

*【デメリット】*

cgo を使用すると、C言語との連携が必要となり、ポインタの取り扱いやメモリ管理に注意する。誤った操作は安全性に影響を与える可能性を考慮する。


### 選択の基準:
シンプルなケース or 複雑な制御
`os/exec` はシンプルで直感的な手法を提供する。`cgo` は高度な制御を可能にする一方、複雑さも伴う。