use rand::Rng; // trait
use std::cmp::Ordering;
use std::io;

fn main() {
    // 带!号，是一个宏，不是函数！
    println!("猜数！");

    let secret_number = rand::thread_rng().gen_range(1, 101);
    println!("神秘数字是: {}", secret_number);
    println!("猜测一个数");

    // 声明变量
    // let mut foo = 1; // 可变变量
    // let bar = foo; // immutable, 不可变变量
    // foo = 2;
    let mut guess = String::new();

    io::stdin().read_line(&mut guess).expect("无法读取行");
    // io::Result Ok, Err

    // shadow, 类型转换
    let guess: u32 = guess.trim().parse().expect("Please type a number!");
    
    println!("你猜测的数是：{}", guess);

    match guess.cmp(&secret_number) {
        Ordering::Less => println!("Too small!"),
        Ordering::Greater => println!("Too big!"),
        Ordering::Equal => println!("You win!"),
    }
}
