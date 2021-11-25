const MAX_POINTS: u32 = 100_000;

fn main() {
    println!("Hello, world!");

    let mut x = 5;
    println!("The value of x is {}", x);

    x = 6;
    println!("The value of x is {}", x);

    // shadowing(隐藏)
    let x = 5;
    let x = x + 3; // 将前一个x变量隐藏
    println!("The value of x is {}", x);

    // shadowing-two
    let spaces = "    ";
    let spaces = spaces.len();

    println!("{}", spaces);

    // 标量类型
    // 整数类型
    let n: i32 = 1;
    let n: u32 = 2;
    let n: isize = 3; //根据系统架构自行推断
                      // let n: u8 = 257; // 溢出-> 在运行时panic. 在发布模式(release)下则环绕，但不会panic
    println!("{}", n);

    // 浮点类型
    let x = 2.0; // f64

    let y: f32 = 3.0; // f32

    // 数值操作(四则运算)
    let sum = 5 + 10;
    let difference = 95.5 - 4.3;
    let product = 4 * 30;
    let quotient = 56.7 / 32.2;

    // rust支持浮点数取余，但要两边都是同类型浮点数, such as 1.0 % 0.3 = 0.1
    let remainder = 54 % 5;

    // 布尔类型
    let t = true;
    let f: bool = false;

    // 字符类型
    // rust中char用来描述最基础的单个字符
    // 占用4个字节大小，是Unicode标量值。（emoji表情都可以）;
    // 但Unicode中没有”字符“概念，因此直觉上的字符与rust中的可能不太一样
    let x = 'a';
    let y: char = 'z';
    let z = '😂';
    println!("{}{}{}", x, y, z);
    // 复合类型
    // 元组(Tuple)
    // tuple长度固定
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    println!("{}, {}, {}", tup.0, tup.1, tup.2);

    // 获取tuple的元素值
    // 利用模式匹配来结构tuple的值
    let (x, y, z) = tup;
    println!("{}, {}, {}", x, y, z);

    // 数组
    let a = [1, 2, 3, 4, 5];
    // 类型: 数组长度
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    // 元素内容: 数组长度
    let a = [3; 5]; // [3, 3, 3, 3, 3]
    // 索引越界，编译会通过，但运行会panic
    println!("{}, {}, {}, {}, {}", a[0], a[1], a[2], a[3], a[4]);
}
