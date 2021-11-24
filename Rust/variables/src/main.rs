const MAX_POINTS:u32 = 100_000;

fn main() {
    println!("Hello, world!");

    let mut x = 5;
    println!("The value of x is {}", x);

    x= 6;
    println!("The value of x is {}", x);

    // shadowing(隐藏)
    let x = 5;
    let x = x + 3; // 将前一个x变量隐藏
    println!("The value of x is {}", x);

    // shadowing-two
    let spaces = "    ";
    let spaces = spaces.len();

    println!("{}", spaces)
}
