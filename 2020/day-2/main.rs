use std::collections::HashSet;
use std::time::Instant;

fn main() {
    let input = parse_input(include_str!("input"));

    println!("Results for Day 1");
    println!("============================");

    // Part 1
    let now = Instant::now();
    let part_01 = part_01(&input);

    println!("Part 1: {:?} ({:.2?})", part_01, now.elapsed());

    // Part 2
    let now = Instant::now();
    let part_02 = part_02(&input);

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> HashSet<&str> {
    input
        .lines()
        .filter(|l| !l.is_empty())
        .collect()
}
pub fn part_01(input: &HashSet<&str>) -> u64 {

    let mut numberOfValidPasswords = 0;

    for elm in input {
        let split:Vec<&str> = elm.split(" ").collect();
        let splitMinMax:Vec<&str> = split[0].split("-").collect();

        let min:i32 = splitMinMax[0].parse::<i32>().unwrap();
        let max:i32 = splitMinMax[1].parse::<i32>().unwrap();
        let character:char = split[1].replace(":", "").chars().next().expect("string is empty");
        let password = split[2];

        let mut counter = 0;
        for c in password.chars() {
            if c == character {
                counter = counter + 1;
            }
        }

        if counter >= min && counter <= max {
            numberOfValidPasswords = numberOfValidPasswords + 1;
        }
        
    }
    return numberOfValidPasswords;
}

pub fn part_02(input: &HashSet<&str>) -> u64 {
       let mut numberOfValidPasswords = 0;

    for elm in input {
        let split:Vec<&str> = elm.split(" ").collect();
        let splitMinMax:Vec<&str> = split[0].split("-").collect();

        let one:i32 = splitMinMax[0].parse::<i32>().unwrap();
        let two:i32 = splitMinMax[1].parse::<i32>().unwrap();
        let character:char = split[1].replace(":", "").chars().next().expect("string is empty");
        let password = split[2];

        if password.chars().nth(one as usize - 1) == Some(character) && password.chars().nth(two as usize - 1) != Some(character) {
            numberOfValidPasswords = numberOfValidPasswords + 1;
        }
        if password.chars().nth(one as usize - 1) != Some(character) && password.chars().nth(two as usize - 1) == Some(character) {
            numberOfValidPasswords = numberOfValidPasswords + 1;
        }
        
    }
    return numberOfValidPasswords;
}