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

fn parse_input(input: &str) -> Vec<&str> {
    input
        .lines()
        .collect()
}

pub fn part_01(input: &Vec<&str>) -> usize {
    let mut max_seat = 0;
    for line in input {
        let row_input = &line[..7];
        let column_input = &line[7..10];

        let mut lower_row = 0;
        let mut upper_row = 127;
        for a in row_input.chars() {
            if a == 'F' {
                upper_row = upper_row - ((upper_row / 2) - (lower_row / 2)) - 1
            } else {
                lower_row = upper_row - ((upper_row - lower_row) / 2)
            }
        }
        let mut lower_column = 0;
        let mut upper_column = 7;
        for b in column_input.chars() {
            if b == 'L' {
                upper_column = upper_column - ((upper_column / 2) - (lower_column / 2)) - 1
            } else {
                lower_column = upper_column - ((upper_column - lower_column) / 2)
            }
        }
        let m = (lower_row * 8) + lower_column;
        if m > max_seat {
            max_seat = m;
        }
    }
    
    max_seat
}

pub fn part_02(input: &Vec<&str>) -> usize {
    let mut seat_list = Vec::new();
    for line in input {
        let row_input = &line[..7];
        let column_input = &line[7..10];

        let mut lower_row = 0;
        let mut upper_row = 127;
        for a in row_input.chars() {
            if a == 'F' {
                upper_row = upper_row - ((upper_row / 2) - (lower_row / 2)) - 1
            } else {
                lower_row = upper_row - ((upper_row - lower_row) / 2)
            }
        }
        let mut lower_column = 0;
        let mut upper_column = 7;
        for b in column_input.chars() {
            if b == 'L' {
                upper_column = upper_column - ((upper_column / 2) - (lower_column / 2)) - 1
            } else {
                lower_column = upper_column - ((upper_column - lower_column) / 2)
            }
        }
        seat_list.push((lower_row * 8) + lower_column);
    }
    seat_list.sort();
    let mut free_seat = 0;
    let mut c = 0;
    for e in seat_list.iter() {
        if c + 1 != *e {
            free_seat = c+1;
        }
        c = *e;
    }
    
    free_seat
}