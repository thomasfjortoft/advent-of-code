use std::time::Instant;

fn main() {
    let input = parse_input(include_str!("input"));
    
    println!("Results for Day 1");
    println!("============================");
    
    // Part 1
    let now = Instant::now();
    let part_01 = part_01(input);
    
    println!("Part 1: {:?} ({:.2?})", part_01, now.elapsed());
    
    let input2 = parse_input(include_str!("input_test"));
    // Part 2
    let now = Instant::now();
    let part_02 = part_02(input2);

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input
        .lines()
        .map(|line| {
            line.chars().collect()
        }).collect()
}

pub fn part_01(mut input: Vec<Vec<char>>) -> usize {
    let mut not_done = true;
    let mut data_copy = input.to_vec();
    while not_done {
        let mut x = 0;
        while input.len() != x {
            let mut y = 0;
            while input[x].len() != y {
                if input[x][y] == '.' {
                    y += 1;
                    continue
                }
                let mut n = String::from(' ');
                if x == 0 && y == 0 {
                    n.push(input[x][y+1]);
                    n.push(input[x+1][y]);
                    n.push(input[x+1][y+1]);
                } else if x == 0 && y != input[x].len() - 1{
                    n.push(input[x][y+1]);
                    n.push(input[x][y-1]);
                    n.push(input[x+1][y]);
                    n.push(input[x+1][y+1]);
                    n.push(input[x+1][y-1]);
                } else if y == input[x].len() - 1 && x == 0 {
                    n.push(input[x][y-1]);
                    n.push(input[x+1][y]);
                    n.push(input[x+1][y-1]);
                } else if y == 0 && x != input.len() - 1 {
                    n.push(input[x-1][y]);
                    n.push(input[x-1][y+1]);
                    n.push(input[x][y+1]);
                    n.push(input[x+1][y+1]);
                    n.push(input[x+1][y]);
                } else if y == input[x].len() - 1 && x != input.len() - 1 {
                    n.push(input[x-1][y]);
                    n.push(input[x-1][y-1]);
                    n.push(input[x][y-1]);
                    n.push(input[x+1][y-1]);
                    n.push(input[x+1][y]);
                } else if x == input.len() - 1 && y == 0 {
                    n.push(input[x-1][y]);
                    n.push(input[x-1][y+1]);
                    n.push(input[x][y+1]);
                } else if x == input.len() - 1 && y != input[x].len() - 1{
                    n.push(input[x][y-1]);
                    n.push(input[x-1][y-1]);
                    n.push(input[x-1][y]);
                    n.push(input[x-1][y+1]);
                    n.push(input[x][y+1]);
                } else if x == input.len() - 1 && y == input[x].len() - 1 {
                    n.push(input[x][y-1]);
                    n.push(input[x-1][y]);
                } else {
                    n.push(input[x][y-1]);
                    n.push(input[x-1][y-1]);
                    n.push(input[x-1][y]);
                    n.push(input[x-1][y+1]);
                    n.push(input[x][y+1]);
                    n.push(input[x+1][y+1]);
                    n.push(input[x+1][y]);
                    n.push(input[x+1][y-1]);
                }

                if input[x][y] == 'L' && !n.contains("#") {
                    data_copy[x][y] = '#';
                }
                
                if input[x][y] == '#' && n.matches('#').count() >= 4 {
                    data_copy[x][y] = 'L';
                }
                y += 1;
            }
            x += 1;
        }
        if input == data_copy {
            not_done = false;
        }
        input = data_copy.to_vec();
    }
    let mut count = 0;
    for e in &data_copy {
        for t in e {
            if *t == '#' {
                count += 1;
            }
        }
    }
    count
}

pub fn part_02(input: Vec<Vec<char>>) -> usize {
    0
}
