use std::time::Instant;
use std::collections::HashMap;


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

fn part_01(input: Vec<&str>) -> u64 {
    let mut mask_0 : u64 = 0;
    let mut mask_1 : u64 = 0;
    let mut mem    : HashMap<u64,u64> = HashMap::new();

    for line in input {
        let tab:Vec<_> = line.split(" = ").collect();

        match &line[0..3]
        {
            "mem"=>{
                let (val,addr) = parse_adr(&tab);
                let mut vv = val;
                vv&=!mask_0;
                vv|= mask_1;
                mem.insert(addr,vv);
            },
            "mas"=>{
                let masks = get_masks(tab[1]);
                mask_0 = masks.0;
                mask_1 = masks.1;
            },
            _ =>panic!("err"),
        }
    }

    mem.values().sum()
}

fn part_02(input: Vec<&str>) -> u64 {
    let mut mask_1 : u64 = 0;
    let mut mask_x : u64 = 0;
    let mut mask_c : u64 = 0;
    let mut mem : HashMap<u64,u64> = HashMap::new();

    for line in input {
        let tab:Vec<_> = line.split(" = ").collect();

        match &line[0..3] {
            "mem"=>{
                let (val,addr) = parse_adr(&tab);              
                let count = 1u64<<mask_c;

                for i in 0u64..count {
                    let mut adr =addr;
                    adr|=  mask_1;               
                    adr&= !mask_x;
                    adr|= get_mask(i,mask_x);
                    mem.insert(adr,val);
                }   
            },
            "mas"=>{
                let masks = get_masks(tab[1]);
                mask_1 = masks.1;
                mask_x = masks.2;
                mask_c = masks.3;
            },
            _ =>panic!("Not able to pars input"),
        }
    }
    mem.values().sum()
}

fn get_masks(s:&str)->(u64,u64,u64,u64) {
    let mut mask_0 = 0u64;
    let mut mask_1 = 0u64;
    let mut mask_f = 0u64;
    let mut number_x = 0u64;

    for c in s.chars() {
        mask_0<<=1;
        mask_1<<=1;
        mask_f<<=1;
        if c=='0' { mask_0|=1; }
        if c=='1' { mask_1|=1; }
        if c=='X' { mask_f|=1; number_x+=1; }
    }
    (mask_0,mask_1,mask_f,number_x)
}

fn vec_to_bin(table:&Vec<u8>) -> u64 {
    table.iter().rev().fold(0,|acc,&x| ((acc<<1) | x as u64))
}

fn get_mask(i:u64, mask_x:u64) -> u64 {    
    let mut mask = mask_x;
    let mut pos = 0u64;
    let mut tab_zero:Vec<u8> = vec![];

    while mask > 0 {
        tab_zero.push(0);
        if mask&1==1 {
            if i&(1u64<<pos)!=0 {                
                let last = tab_zero.len()-1;
                tab_zero[last] = 1;
            }
            pos+=1;
        }      
        mask>>=1;
    }
  
    vec_to_bin(&tab_zero)
}

fn parse_adr(tab:&Vec<&str>) -> (u64,u64) {
    let val:u64 = tab[1].parse().unwrap();
    let add1 = tab[0].find("[").unwrap();
    let add2 = tab[0].find("]").unwrap();
    let addr:u64 = tab[0][add1+1..add2].parse().unwrap();
    (val,addr)
}