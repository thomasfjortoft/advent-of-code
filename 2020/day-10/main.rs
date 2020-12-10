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
    let mut output: Vec<usize> = input
        .split("\n")
        .map(|s| s.parse().unwrap())
        .collect();
    output.push(0);
    output.push(*output.iter().max().unwrap() + 3);
    output.sort_unstable();
    output
}

pub fn part_01(input: Vec<usize>) -> usize {
    let mut pointer = 0;
    let mut one_jolt = 0;
    let mut three_jolt = 0;
    while input.len() - 1 != pointer {
        if input[pointer] + 1 == input[pointer+1] {
            one_jolt += 1;
        }
        if input[pointer] + 3 == input[pointer+1] {
            three_jolt += 1;
        }
        pointer += 1;
    }
    one_jolt * three_jolt
}

pub fn part_02(input: Vec<usize>) -> usize {
    let mut adapters: Vec<usize> = vec![0; input.len()];
    *adapters.last_mut().unwrap() = 1;

    for pointer in (0..input.len()).rev() {
        for next_pointer in (pointer + 1)..(pointer + 4) {
            if let Some(next_value) = input.get(next_pointer) {
                if next_value - input[pointer] <= 3 {
                    adapters[pointer] += adapters[next_pointer];
                }
            }
        }
    }
    adapters[0]
}
