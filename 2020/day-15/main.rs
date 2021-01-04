use std::{cmp::max, collections::HashMap};
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

fn parse_input(input: &str) -> Vec<usize> {
    input
        .split(",")
        .map(|n| n.parse().unwrap())
        .collect()
}

pub fn part_01(input: Vec<usize>) -> usize {
    spoken(input, 2020)
}

pub fn part_02(input: Vec<usize>) -> usize {
    spoken(input, 30000000)
}

fn spoken(input: Vec<usize>, num: usize) -> usize {
    let mut mem: HashMap<usize, (usize, usize)> = HashMap::new();
    let mut x = 0;
    for i in 0..num {
        if i < input.len() {
            x = input[i];
        } else {
            let t = mem.entry(x).or_default();
            x = if t.1 == 0 || t.0 == 0 { 0 } else { t.0 - t.1 }
        }

        let t = mem.entry(x).or_insert((0, 0));
        *t = (i + 1, max(t.1, t.0));
    }
    x
}