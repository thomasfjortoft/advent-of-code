use std::time::Instant;
use std::collections::HashMap;

fn main() {
    let input = parse_input(include_str!("input"));
    //println!("{:?}", input);

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

fn parse_input(input: &str) -> Vec<&str> {
    input
        .split("\n\n")
        .collect()
}

pub fn part_01(input: &Vec<&str>) -> usize {
    let mut sum = 0;
    for e in input {
        let mut t = e.chars().filter(|l| *l != '\n').collect::<Vec<char>>();
        t.sort();
        t.dedup();
        sum = sum + t.len();
    }
    sum
}

pub fn part_02(input: &Vec<&str>) -> usize {
    let mut sum = 0;
    for group in input {
        let persons:Vec<&str> = group.split('\n').collect();
        let number_persons = persons.len();
        let mut map = HashMap::new();
        let mut sum_group = 0;
        for p in persons {
            let c = p.chars();
            for answer in c {
                *map.entry(answer).or_insert(0) += 1;
            }
        }
        //println!("{:?} pers={}", map, number_persons);

        for (key, value) in &map {
            if *value == number_persons {
                sum += 1;
            }
        }

    }
    sum
}