use std::time::Instant;

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

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input
        .lines()
        .map(|line| {
            line.chars().collect()
        }).collect()
}

pub fn part_01(input: &Vec<Vec<char>>) -> u64 {
    slope(input, 1, 3)
}

pub fn part_02(input: &Vec<Vec<char>>) -> u64 {
    let slope1 = slope(input, 1, 1);
    let slope2 = slope(input, 1, 3);
    let slope3 = slope(input, 1, 5);
    let slope4 = slope(input, 1, 7);
    let slope5 = slope(input, 2, 1);


    slope1 * slope2 * slope3 * slope4 *slope5
}

pub fn slope(input: &Vec<Vec<char>>, inc_x: usize, inc_y: usize) -> u64 {
    let rows = input.len();
    let mut trees = 0;
    let (mut row, mut col) = (0, 0);

    while row < rows {
        let cols = input[row].len();

        if input[row][col % cols] == '#' {
            trees += 1
        }
        row += inc_x;
        col += inc_y;
    }
    
    trees
}