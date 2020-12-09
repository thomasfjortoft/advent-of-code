use std::time::Instant;
use std::collections::HashMap;

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
    let part_02 = part_02();

    println!("Part 2: {:?} ({:.2?})", part_02, now.elapsed());
}

fn parse_input(input: &str) -> Vec<&str> {
    input
        .split("\n")
        .collect()
}

pub fn part_01(input: &Vec<&str>) -> usize {
    let mut count = 0;
    let mut map = HashMap::new();
    for rule in input {
        let t = rule.split(" bags contain ").collect::<Vec<&str>>();
        let start = t[0];
        let steps = t[1].split(", ").collect::<Vec<&str>>();
        map.insert(start, steps);
    }
    for (a, b) in &map {
        if find(a, b.to_vec(), &map) {
            count += 1;
        }
    }
    count
}
pub fn find(start: &str, list: Vec<&str>, map: &HashMap<&str, Vec<&str>>) -> bool{
    for mut e in list {
        if e.contains("shiny gold") {
            return true;
        }
        if e.contains("no other bags") {
            return false;
        }
        e = &e[2..e.len()];
        e = e.split(" bag").collect::<Vec<&str>>()[0];
        if find(e, map.get(e).unwrap().to_vec(), &map) {
            return true;
        }
    }
    return false;
}

pub fn part_02() -> usize {
    let input = to_hashmap(include_str!("input"));
    let rules = input.get("shiny gold bag").unwrap();
    return rules
        .iter()
        .map(|br| br.bag_count(&input, br.num))
        .sum::<usize>();
}

struct BagRule {
    num: usize,
    bag_type: String,
}

impl BagRule {
    fn bag_count(&self, collection: &HashMap<String, Vec<BagRule>>, prev_count: usize) -> usize {
        let rules = collection.get(&self.bag_type).unwrap();
        if rules.is_empty() {
            prev_count
        } else {
            rules
                .iter()
                .map(|br| br.bag_count(collection, br.num * prev_count))
                .sum::<usize>()
                + prev_count
        }
    }
}

impl From<&str> for BagRule {
    fn from(s: &str) -> Self {
        match s.find(" ") {
            Some(n) => {
                let num: usize = s[0..n].parse().unwrap();
                BagRule {
                    num,
                    bag_type: String::from(s[n + 1..].trim_end_matches("s")),
                }
            }
            // no bags
            None => {
                panic!("boom")
            }
        }
    }
}


fn to_hashmap(input: &str) -> HashMap<String, Vec<BagRule>> {
    input
        .lines()
        .map(|i| {
            let mut splt = i.split(" contain ");
            let bag = splt.next().unwrap().trim_end_matches("s");
            let unparsed_rules = splt.next().unwrap().trim_end_matches(".");
            let rules: Vec<BagRule> = if unparsed_rules == "no other bags" {
                vec![]
            } else {
                unparsed_rules.split(", ").map(|s| s.into()).collect()
            };
            (String::from(bag), rules)
        })
        .collect()
}