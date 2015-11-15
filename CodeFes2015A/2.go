/*
問題文
高橋君は長さ N の数列 A={A1,A2,…,AN} を用意しました。高橋君は数列 A を元に「とても長い数列」を、以下のような手順で作ろうとしています。
まず長さ 0 の数列を用意し、これを S と呼ぶことにする。
S、A1、S をこの順につなげた数列を新たな S とする。
S、A2、S をこの順につなげた数列を新たな S とする。
（中略）
S、AN、S をこの順につなげた数列を新たな S とする。
この時点での S を「とても長い数列」とする。
例えば N=3,A1=1,A2=2,A3=3 なら、S は {}　→ {1} → {1,2,1} → {1,2,1,3,1,2,1} と変化し、「とても長い数列」は {1,2,1,3,1,2,1} となります。
高橋君はこの手順によって作られる「とても長い数列」に含まれる数の和を知りたがっています。これを高橋君の代わりに計算するプログラムを作成してください。

入力
入力は以下の形式で標準入力から与えられる。
N
A1 A2 ... AN
1 行目には、整数 N(1≦N≦30) が与えられる。
2 行目には、N 個の整数が空白区切りで与えられる。このうち i(1≦i≦N) 番目には、整数 Ai(0≦Ai≦100) が与えられる。
「とても長い数列」に含まれる数の和は 109 以下となることが保証されている。

出力
「とても長い数列」に含まれる数の和を 1 行に出力せよ。出力の末尾に改行を入れること。
*/
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "math"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
  sc.Scan()
  i, _ := strconv.Atoi(sc.Text())
  return i
}

func main() {
  var A []int
  var sum int

  sc.Split(bufio.ScanWords)
  N := nextInt()
  for i := 0; i < N; i++ {
      A = append(A, nextInt())
  }

  for idx, val := range A {
    sum += val * int(math.Pow(2, float64(N-1-idx)))
  }

  fmt.Println(sum)
}
