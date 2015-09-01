# metaputer mother module
## 仕様
CPU Memory その他デバイスとのやりとりをまとめる中枢
基本的なフローは
CPUからアドレスを1Byte受け取る
メモリを読み取ってCPUに送る
CPUから送られてくる命令(4Byte)を実行する
実行した結果(4Byte)をCPUに返す
の繰り返し

ただし、最初にBIOSのデータをメモリに格納する

受け取る最初のバイトがモードで残りの3バイトが引数  
以下にCPUが行える命令一覧を示す
XYZがそれぞれ引数 ABCDがそれぞれ戻り値

|ID|モード|X|Y|Z|A|B|C|D|説明|
|:-|----|-----|-|-|------|-|-|-|----|
|0|Nope|－|－|－|0|0|0|0|何もしない|
|1|ReadMem|アドレス|－|－|値|0|0|0|メモリ内の数値を読み取る|
|2|WriteMem|アドレス|値|－|0|0|0|0|メモリ内に数値を書き込む|