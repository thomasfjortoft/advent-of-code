use std::time::Instant;

fn main() {
    let input = parse_input(include_str!("input"));
    
    println!("Results for Day 1");
    println!("============================");
    
    // Part 1
    let now = Instant::now();
    let part_01 = part_01(input);
    
    println!("Part 1: {:?} ({:.2?})", part_01, now.elapsed());
    
    let input2 = parse_input(include_str!("input"));
    // Part 2
    let now = Instant::now();
    let part_02 = part_02(input2);

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> Vec<&str> {
    input
        .split("\n")
        .collect()
}

pub fn part_01(input: Vec<&str>) -> usize {
    let mut not_done = true;
    let mut pointer = 0;
    let mut acc:i32 = 0;

    let mut list:Vec<usize> = Vec::new();

    while not_done {
        let split:Vec<&str> = input[pointer].split(" ").collect();
        let command = split[0];
        let plus_minus = split[1].chars().next().unwrap();
        let amount = split[1].replace("-", "").replace("+", "").parse::<i32>().unwrap();
        match command {
            "nop" => {
                pointer += 1;
            },
            "acc" => {
                if plus_minus == '+' {
                    acc += amount;
                }
                if plus_minus == '-' {
                    acc -= amount;
                }
                pointer += 1;
            },
            "jmp" => {
                if plus_minus == '+' {
                    pointer += amount as usize;
                }
                if plus_minus == '-' {
                    pointer -= amount as usize;
                }
            },
            _ => (),
        }

        for el in &list {
            if *el == pointer {
                not_done = false;
            }
        }
        list.push(pointer);
    }
    acc as usize
}

pub fn part_02(input: Vec<&str>) -> usize {
    let mut count = 0;
    let mut res = 0;
    while res == 0 && input.len() > count{
        res = loop_me(&input, count);
        count += 1;
    }
    res
}

pub fn loop_me(input: &Vec<&str>, pointer_to_change: usize) -> usize {
    let mut count = 0;
    let mut not_done = true;
    let mut pointer = 0;
    let mut acc:i32 = 0;
    while not_done {
        let split:Vec<&str> = input[pointer].split(" ").collect();
        let mut command = split[0];
        let plus_minus = split[1].chars().next().unwrap();
        let amount = split[1].replace("-", "").replace("+", "").parse::<i32>().unwrap();

        if pointer_to_change == pointer {
            if command.contains("nop") {
                command = "jmp";
            }
            if command.contains("jmp") {
                command = "nop";
            }
        }

        match command {
            "nop" => {
                pointer += 1;
            },
            "acc" => {
                if plus_minus == '+' {
                    acc += amount;
                }
                if plus_minus == '-' {
                    acc -= amount;
                }
                pointer += 1;
            },
            "jmp" => {
                if plus_minus == '+' {
                    pointer += amount as usize;
                }
                if plus_minus == '-' {
                    pointer -= amount as usize;
                }
            },
            _ => (),
        }
        count += 1;
        if count >= input.len() {
            acc = 0;
            not_done = false;
        }

        if pointer >= input.len() {
            not_done = false
        }
    }
    acc as usize
}