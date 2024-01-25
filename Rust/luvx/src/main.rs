fn main() {
    let v = add(5, 6);
    println!("v 的值为 : {}", v);

    for_loop();

    m1();

    m2();
}

fn add(a: i32, b: i32) -> i32 {
    return a + b;
}

fn for_loop() {
    let a = [10, 20, 30, 40, 50];
    for i in 0..5 {
        println!("a[{}] = {}", i, a[i]);
    }
}

fn m1() {
    let arr = [1, 3, 5, 7, 9];
    let part = &arr[0..3];
    for i in part.iter() {
        println!("{}", i);
    }
}

fn m2() {
    let u = User {
        username: String::from("foobar"),
    };
    println!("{}", u.username)
}

struct User {
    username: String,
}
