use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt");
    match file {
        Ok(f) => {
            let reader = BufReader::new(f);

            let mut list1: Vec<i32> = Vec::new();
            let mut list2: Vec<i32> = Vec::new();
            let mut number_count = HashMap::new();
            for line in reader.lines() {
                match line {
                    Ok(content) => {
                        let fields: Vec<&str> = content.split_whitespace().collect();
                        let n1: i32 = fields[0].parse().expect("Failed to parse!");
                        let n2: i32 = fields[1].parse().expect("Failed to parse!");
                        list1.push(n1);
                        list2.push(n2);
                        number_count
                            .entry(n2)
                            .and_modify(|v| *v = *v + 1)
                            .or_insert(1);
                    }
                    Err(e) => eprintln!("Error reading line: {}", e),
                }
            }
            let mut similarity: i32 = 0;
            for number in &list1 {
                match number_count.get(&number) {
                    Some(value) => similarity += value * number,
                    None => {}
                }
            }
            list1.sort();
            list2.sort();
            let mut difference_sum: i32 = 0;
            for i in 0..list1.len() {
                let diff = list1[i] - list2[i];
                if diff < 0 {
                    difference_sum += -diff;
                } else {
                    difference_sum += diff;
                }
            }
            println!("Sum of differences: {}", difference_sum);
            println!("Similarity: {}", similarity);
        }
        Err(e) => {
            println!("Could not open the file!:\n {}", e)
        }
    }
}
