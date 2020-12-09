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
    let part_02 = part_02(input2, part_01);

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> Vec<usize> {
    input
        .split("\n")
        .map(|s| s.parse().unwrap())
        .collect()
}

pub fn part_01(input: Vec<usize>) -> usize {
    let mut output = 0;
    let mut pointer = 0;
    while input.len() - 26 != pointer {
        let preamble = &input[pointer..pointer+25];
        let number_check = input[pointer+25];
        let mut sum_preamble:Vec<usize> = Vec::new();
        for e in preamble {
            for el in preamble {
                sum_preamble.push(e+el);
            }
        }
        if !sum_preamble.iter().any(|&s| s == number_check) {
            output = number_check;
            break
        }
        pointer += 1;
    }
    output
}

pub fn part_02(input: Vec<usize>, number_to_check: usize) -> usize {
    let mut pointer = 0;
    let mut output_numbers:Vec<usize> = Vec::new();
    while input.len() > pointer {
        let mut sum = 0;
        let mut inner_pointer = pointer;
        let mut numbers:Vec<usize> = Vec::new();

        while number_to_check > sum {
            sum += input[inner_pointer];
            numbers.push(input[inner_pointer]);
            inner_pointer += 1;
        }
        if sum == number_to_check {
            output_numbers = numbers;
            break
        }
        pointer += 1;
    }
    output_numbers.sort();
    output_numbers[0] + output_numbers[output_numbers.len() - 1]

}
