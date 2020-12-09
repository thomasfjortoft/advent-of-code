use std::fs::File;
use std::io::{BufRead, BufReader, Error, ErrorKind, Read};

fn read<R: Read>(io: R) -> Result<Vec<i64>, Error> {
    let br = BufReader::new(io);
    br.lines()
        .map(|line| line.and_then(|v| v.parse().map_err(|e| Error::new(ErrorKind::InvalidData, e))))
        .collect()
}

fn main() -> Result<(), Error> {
    let vec = read(File::open("input.txt")?)?;

    find_values_from_two(&vec);
    find_values_from_three(&vec);
    
    Ok(())
}

fn find_values_from_two(vec: &Vec<i64>) -> () {
    println!("start");
    for x in vec {
        for y in vec {
            if x + y == 2020 {
                println!("The answer is {}", x*y);
                return;
            } 
        }
    }
}
fn find_values_from_tre(vec: &Vec<i64>) -> () {
    println!("start");
    for x in vec {
        for y in vec {
            for z in vec {
                if x + y + z == 2020 {
                    println!("The answer is {}", x*y*z);
                    return;
                } 
            }
        }
    }
}