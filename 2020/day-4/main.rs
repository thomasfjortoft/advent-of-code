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

fn get_numeric(s:&str)->i32{
    let mut id=0;
    while s.chars().nth(id).unwrap_or(' ').is_digit(10) { id+=1; } 
    s[..id].parse().unwrap_or(0)
}

fn is_valid_hight(v:&str) -> bool {
    let number = get_numeric(v);
    if v.len()==5 && v.contains("cm") && number>=150 && number<=193 {
        return true 
    } else if v.len()==4 && v.contains("in") && number>= 59 && number<= 76 {
       return  true
    } 
    false
}

pub fn part_02(input: &HashSet<&str>) -> usize {
    input
        .iter()
        .filter(|l| {
            l.contains("byr") && l.contains("iyr") && l.contains("eyr") && l.contains("hgt") && l.contains("hcl") && l.contains("ecl") && l.contains("pid")})
        .filter(|l| {
            let t = l.replace("\n", " ");
            let split = t.split(" ").collect::<Vec<&str>>();
            let mut valid = false;
            for elem in split {
                let value:Vec<&str>= elem.split(":").collect();
                match value[0] {
                    "byr" => valid = value[1].parse::<i32>().unwrap() >= 1920 && value[1].parse::<i32>().unwrap() <= 2002,
                    "iyr" => valid = value[1].parse::<i32>().unwrap() >= 2010 && value[1].parse::<i32>().unwrap() <= 2020,
                    "eyr" => valid = value[1].parse::<i32>().unwrap() >= 2020 && value[1].parse::<i32>().unwrap() <= 2030,
                    "hgt" => valid = is_valid_hight(value[1]),
                    "ecl" => valid = l.contains("amb") || l.contains("blu") || l.contains("brn") || l.contains("gry") || l.contains("grn") || l.contains("hzl") || l.contains("oth"),
                    "pid" => valid = value[1].len() == 9,
                    "hcl" =>  valid = value[1].len()==7,
                    "cid" =>  valid = true,
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