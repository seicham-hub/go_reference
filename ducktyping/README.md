## what is ducktyping?

ダックタイピングは「もし、鳥がアヒルのように歩き、アヒルのように鳴くのならその鳥はアヒルである」  
という考えに基づいている。  
つまり、オブジェクトの型が何であるかよりも、そのオブジェクトがどのような振る舞いを持っているか  
(関数を持っているか)が重要という考え。

例えば、以下の関数に引数としてアヒル構造体(ducK)と鳩構造体(pidgeot)どちらも取りたい場合

```
 fly(引数)
```

go 等の静的片付け言語では型が異なる為、以下のように二つ関数を作成する必要がある。

```
 fly_duck(d Duck)
 fly_pidgeot(c Pidgeot)
```

しかし ducktyping で「特定のメソッドを持つ構造体をインターフェース」として定義することで、
動的型付け言語のように異なる型を引数に取る事ができる。

```

// Quackという関数を持っている構造体を同じものとして定義
type Bird interface{
    Quack() string
}

func fly(b Bird) string{
    return "flying now"
}

func main(){
    // DuckでもPidgeotでもエラーにならない
    fly(Duck)
    fly(Pidgeot)
}


// 以下のように構造体を定義している
type Duck struct{
    height int
}

func (d Square) Quack() string {
	return "quack"
}

type Pidgeot struct{
    height int
}

func (d Square) Quack() string {
	return "holhol"
}

```
