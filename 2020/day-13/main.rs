use std::time::Instant;

fn main() {
    let (dep, input) = parse_input(include_str!("input"));
    
    println!("Results for Day 1");
    println!("============================");
    
    // Part 1
    let now = Instant::now();
    let part_01 = part_01(dep, input);
    
    println!("Part 1: {:?} ({:.2?})", part_01, now.elapsed());
    
    let (_, input2) = parse_input(include_str!("input_test"));
    // Part 2
    let now = Instant::now();
    let part_02 = part_02(input2);

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> (i32, Vec<&str>) {
    let split:Vec<&str> =  input
        .split("\n")
        .collect();
    (split[0].parse::<i32>().unwrap(), split[1].split(",").collect())
}

pub fn part_01(dep: i32, input: Vec<&str>) -> usize {
    let mut closest_buss = 0;
    let mut min_wait = 0;
    for elem in input {
        if !elem.contains('x') {
            let buss_id = elem.parse::<i32>().unwrap();
            let next_dep = dep - (dep % buss_id) + buss_id;
            let t = next_dep - dep;
            if min_wait == 0 || min_wait > t {
                min_wait = t;
                closest_buss = buss_id;
            }
        }
    }
    (closest_buss * min_wait) as usize
}

pub fn part_02(input: Vec<&str>) -> usize {
    0
}
