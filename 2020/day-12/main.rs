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
    let directions = vec!('E', 'S', 'W', 'N');
    let mut ship_direction:i32 = 0;
    let mut pos_latitude:i32 = 0;
    let mut pos_longitude:i32 = 0;
    for elem in input {
        let direction:char = elem.chars().next().unwrap();
        let value = &elem[1..elem.len()].parse::<i32>().unwrap();
        match direction {
            'N' => pos_latitude += value,
            'S' => pos_latitude -= value,
            'E' => pos_longitude += value,
            'W' => pos_longitude -= value,
            'R' => {
                match value {
                    90 => ship_direction = (ship_direction + 1) % 4,
                    180 => ship_direction = (ship_direction + 2) % 4,
                    270 => ship_direction = (ship_direction + 3) % 4,
                    _ => println!("Not supported")
                }
            },
            'L' => {
                match value {
                    90 => ship_direction = (ship_direction + 3) % 4,
                    180 => ship_direction = (ship_direction + 2) % 4,
                    270 => ship_direction = (ship_direction + 1) % 4,
                    _ => println!("Not supported")
                }
            },
            'F' => {
                match directions[ship_direction as usize] {
                    'N' => pos_latitude += value,
                    'S' => pos_latitude -= value,
                    'E' => pos_longitude += value,
                    'W' => pos_longitude -= value,
                    _ => println!("Not supported")
                }
            },
            _ => println!("not supported")
        }
    }
    (pos_latitude.abs() + pos_longitude.abs()) as usize
}

pub fn part_02(input: Vec<&str>) -> usize {
    let mut pos_latitude:i32 = 0;
    let mut pos_longitude:i32 = 0;
    let mut waypoint_latitude:i32 = 1;
    let mut waypoint_longitude:i32 = 10;
    for elem in input {
        let command:char = elem.chars().next().unwrap();
        let value = &elem[1..elem.len()].parse::<i32>().unwrap();
        match command {
            'N' => waypoint_latitude += value,
            'S' => waypoint_latitude -= value,
            'E' => waypoint_longitude += value,
            'W' => waypoint_longitude -= value,
            'L' => {
                match value {
                    90 => {
                        let tmp = waypoint_longitude;
                        waypoint_longitude = -waypoint_latitude;
                        waypoint_latitude = tmp;
                    },
                    180 => {
                        waypoint_longitude = -waypoint_longitude;
                        waypoint_latitude = -waypoint_latitude;
                    },
                    270 => {
                        let tmp = waypoint_longitude;
                        waypoint_longitude = waypoint_latitude;
                        waypoint_latitude = -tmp;
                    },
                    _ => println!("Next")
                }
            },
            'R' => {
                match value {
                    90 => {
                        let tmp = waypoint_longitude;
                        waypoint_longitude = waypoint_latitude;
                        waypoint_latitude = -tmp;
                    },
                    180 => {
                        waypoint_longitude = -waypoint_longitude;
                        waypoint_latitude = -waypoint_latitude;
                    },
                    270 => {
                        let tmp = waypoint_longitude;
                        waypoint_longitude = -waypoint_latitude;
                        waypoint_latitude = tmp;
                    },
                    _ => println!("Next")
                }
            },
            'F' => {
                pos_latitude += value * waypoint_latitude;
                pos_longitude += value * waypoint_longitude;
            },
            _ => println!("not supported")
        }
    }
    (pos_latitude.abs() + pos_longitude.abs()) as usize
}