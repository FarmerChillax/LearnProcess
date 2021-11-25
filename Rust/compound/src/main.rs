fn main() {
    println!("Hello, world!");
    another_function();
    another_function_2(2, 3);
    another_function_3();

    let x = five();
    println!("the value of x is: {}", x);

    let y = plus_five(6);
    println!("the value of y is: {}", y);
}

// 无参数函数
fn another_function() {
    println!("Another function");
}

// 函数的参数
// parameters(形参), arguments(实参)
fn another_function_2(x: i32, y: i32) {
    println!("the value of x is: {}", x);
    println!("the value of y is: {}", y);
}

// 函数体中的语句&表达式
fn another_function_3() {
    // let x = (y == 6);
    let x = 6;
    let y = {
        let x = 1;
        x + 3; // 有分号; 时，表示一个语句，则return一个tuple

        // 多行注释
        /* 因此不能用分号，表示返回值. 
        大多数函数默认使用最后一个表达式为返回值
        若想提前返回，需要用return关键字，并指定一个值; return 10
        */ 
        x + 3
    };

    println!("The value of y is: {}", y)

}

fn five() -> i32 {
    5
}


fn plus_five(x: i32) -> i32 {
    x + 5
}