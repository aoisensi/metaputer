# metaputer memory module
## 仕様
常に4byte受け取って処理をしてから4byte返す  
受け取る最初のバイトがモードで残りの3バイトが引数  
以下に一覧を示す  
XYZがそれぞれ引数 ABCDがそれぞれ戻り値

|ID|モード|X|Y|Z|A|B|C|D|説明|
|:-|----|-----|-|-|------|-|-|-|----|
|0|Nope|－|－|－|0|0|0|0|何もしない|
|1|Read|アドレス|－|－|値|0|0|0|メモリ内の数値を読み取る|
|2|Write|アドレス|値|－|0|0|0|0|メモリ内に数値を書き込む|
