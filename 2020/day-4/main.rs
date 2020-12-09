use std::time::Instant;
use std::collections::HashSet;

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

fn parse_input(input: &str) -> HashSet<&str> {
    input
        .split("\n\n")
        .collect()
}

pub fn part_01(input: &HashSet<&str>) -> usize {
    input
        .iter()
        .filter(|l| {
            l.contains("byr") && l.contains("iyr") && l.contains("eyr") && l.contains("hgt") && l.contains("hcl") && l.contains("ecl") && l.contains("pid")})
        .count()
}

pub fn part_02(input: &HashSet<&str>) -> usize {
    input
        .iter()
        .filter(|l| {
            l.contains("byr") && l.contains("iyr") && l.contains("eyr") && l.contains("hgt") && l.contains("hcl") && l.contains("ecl") && l.contains("pid")})
        .filter(|l| {
            let t = l.replace("\n", " ");
            let split = t.split(" ").collect::<Vec<&str>>();
            let mut valid = true;
            for elem in split {
                let value:Vec<&str>= elem.split(":").collect();

                //let hair = Regex::new(r"#.{6}$").unwrap();
                match value[0] {
                    "byr" => valid = !(value[1].parse::<i32>().unwrap() >= 1920 && value[1].parse::<i32>().unwrap() <= 2002),
                    "iyr" => valid = !(value[1].parse::<i32>().unwrap() >= 2010 && value[1].parse::<i32>().unwrap() <= 2020),
                    "eyr" => valid = !(value[1].parse::<i32>().unwrap() >= 2020 && value[1].parse::<i32>().unwrap() <= 2030),
                    "hgt" =>  {
                        if value[1].contains("cm") {
                            let mut a: String = value[1].to_string();
                            a.pop();
                            a.pop();
                            valid = !(a.parse::<i32>().unwrap() >= 150 && a.parse::<i32>().unwrap() <= 193);
                        }
                        if value[1].contains("in") {
                            let mut a: String = value[1].to_string();
                            a.pop();
                            a.pop();
                            valid = !(a.parse::<i32>().unwrap() >= 59 && a.parse::<i32>().unwrap() <= 76);
                        }
                    },
                    "ecl" => {
                        valid = !(l.contains("amb") || l.contains("blu") || l.contains("brn") || l.contains("gry") || l.contains("grn") || l.contains("hzl") || l.contains("oth"));
                    },
                    "pid" => valid = value[1].len() != 9,
                    "hcl" => {
                        if valid == true {
                            //println!("{}", value[1]);
                        }
                    },
                    _ => (),
                }
                if !valid {
                    break
                }
            }
            valid
        })
        .count()
}