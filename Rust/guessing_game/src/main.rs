use rand::Rng; // trait
use std::cmp::Ordering;
use std::io;

fn main() {
    // 带!号，是一个宏，不是函数！
    println!("猜数！");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    // 死循环
    loop {
        // 声明变量
        // let mut foo = 1; // 可变变量
        // let bar = foo; // immutable, 不可变变量
        // foo = 2;
        let mut guess = String::new();

        io::stdin().read_line(&mut guess).expect("无法读取行");
        // io::Result Ok, Err

        // shadow, 类型转换
        // .expect("Please type a number!");
        let guess: u32 = match guess.trim().parse(){
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("你猜测的数是：{}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            },
        }
        println!("猜测一个数");
    }

    
}
