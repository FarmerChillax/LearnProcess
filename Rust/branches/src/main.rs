fn main() {
    println!("if else 条件分支语句");
    let number = 7;

    // Rust不会把表达式计算成布尔类型—>与Ruby或python等语言不同，
    // Rust不会自动将非布尔类型的值转换成bool类型
    // 因此表达式必须产生一个bool值否则会触发编译错误。如下两个栗子：
    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }

    // 这样的语法在许多编程语言中是允许的，但在rust中不允许。
    // 为了确保程序能正确编译运行，需要先将下面三行注释掉。
    // if number {
    //     println!("number was three");
    // }

    // 如果使用多于一个else if, 则最好使用match来重构
    let number = 6;
    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3 or 2");
    }

    // ∵ if 是一个表达式，因此可以将它放在let语句中的等号的右边
    let condition = true;

    //  返回值类型一定要相同
    let number = if condition { 5 } else { 6 };
    // let number = if condition {5} else {"7"};

    println!("The value of number is: {}", number);

    println!("Rust的循环-----loop");

    loop {
        println!("again!");
        break
    }

    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is: {}", result); // 20

    println!("Rust的循环-----while");
    // while条件循环
    let mut number = 3;
    while number != 0 {
        println!("{}!", number);
        number -= 1;
    }

    println!("LIFTOFF!!!");

    println!("Rust的循环-----for");
    let a = [10, 20, 30, 40, 50];
    // 使用while循环遍历集合(丑陋)
    // let mut index = 0;
    // while index < 5 {
    //     println!("the value is: {}", a[index]);
    //     index += 1;
    // }
    // 使用for循环遍历集合(优雅)
    for elem in a { // 相当于将数组内的元素复制了出来
        println!("the value is: {}", elem);
    }

    for elem in a.iter() { // 使用迭代器, 引用数组内的元素
        println!("the value is: {}", elem);
    }
    
    println!("Rust的循环-----for range");
    // range由标准库提供
    // rev方法可以反转range
    // (1..4) 就是range
    for number in (1..4).rev() {
        println!("{}!", number);
    }

    println!("LIFTOFF!!!");

}
