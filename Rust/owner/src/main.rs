fn main() {
    // s 不可用
    let s = "hello"; // s 可用
                     // 可以对 s 进行相关操作

    create_string();
} // s 作用域到此结束, s 不再可用

// 所有权, heap
fn owner() {}

// Rust的第二种字符串类型: String.
//  - 在heap上分配。能够存储在编译时未知数量的文本。
// 使用from函数从字符串字面值创建string类型
fn create_string() {
    // String类型为了支持可变性，需要在heap上分配内存来保存编译时
    // 未知的文本内容：
    //     - 操作系统必须在运行时来请求内存
    //     - 当用完string之后，需要使用某种方式将内存释放（返回给操作系统）
    //          -> 比如Java、Go等带有GC的语言，GC会跟踪清理不再使用的内存
    //          -> 没有GC的语言，则需要我们去识别内存何时不再使用，并调用代码将它释放
    //              如果忘了，就会浪费内存、如果提前释放了，变量就会非法、如果做了两次释放，也是个Bug
    // Rust采用了不同的方式：当拥有它的变量走出作用范围时，内存会立即自动的释放
    // 栗子如下：
    // s变量的作用域在该函数内，当函数运行完时，则s会被释放掉（会自动调用drop函数）
    let mut s = String::from("hello");

    s.push_str(", world");

    println!("{}", s);

    // 变量与数据交互的方式: 移动（move）
    let x = 5;
    let y = x; // 创建了一个X的副本, 然后绑定到y上
               // 由于上述是整数类型，大小是已知且固定的，因此这两个 5 被压倒了stack中

    // String
    let s1 = String::from("hello");
    let s2 = s1;
    // String是一个结构体, 组成结构如下: (整体与go类似)
    // {
    //      ptr -> 指向一个字符数组的指针,
    //      len -> string的长度,
    //      capacity -> String从操作系统总共获得内存的总字节数,
    // }
    // 上面的这个结构体是分配在stack上，而字符数组(字符串的内容部分) 则分配在 heap 上。

    // 上述将 s1 赋给 s2，string的数据被复制了一份。（仅String这个结构体, 不包括底层字符串内容）
    // s1 与 s2 均指向同一个存放字符串内容的 heap
    // 当变量离开作用域时，Rust会自动调用drop函数，并将变量使用的heap内存释放。
    // 当 s1、s2 离开作用域时，他们都会尝试释放对应的内存，这里就是同一个 heap
    // 因为 s1 s2 均指向同一个heap，所以会造成二次释放(double free) bug.

    // 为了保证内存安全，Rust会这么做:
    //      - Rust没有尝试复制被分配的内存
    //      - Rust 让 s1 失效。(s1 离开作用域时，Rust不需要释放任何东西)
    // so，试试看当 s2 创建之后，再使用 s1 是什么效果:
    // println!("{}", s1);
    // 报错: borrow of moved value: `s1`

    // 与常规的浅拷贝和深拷贝不同
    // 一般而言，复制指针、长度、容量视为浅拷贝。
    // 但Rust让 s1 失效了，所以这里用一个新的术语：移动(move)


    // 变量与数据交互的方式: 克隆(Clone)
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("{}, {}", s1, s2);

}
