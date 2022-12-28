use std::{collections::HashSet, fs::read_to_string};

const ASCII_UPP_A_VAL: i32 = 'A' as i32;
const ASCII_UPP_Z_VAL: i32 = 'Z' as i32;
const ASCII_LOW_A_VAL: i32 = 'a' as i32;
const ASCII_LOW_Z_VAL: i32 = 'z' as i32;

// a-z returns 1-26, A-Z returns 27-52
fn get_priority(item: char) -> i32 {
    let ascii_val = item as i32;
    if ascii_val >= ASCII_LOW_A_VAL && ascii_val <= ASCII_LOW_Z_VAL {
        ascii_val - ASCII_LOW_A_VAL + 1
    } else if ascii_val >= ASCII_UPP_A_VAL && ascii_val <= ASCII_UPP_Z_VAL {
        ascii_val - ASCII_UPP_A_VAL + 27
    } else {
        0
    }
}

// Main idea behind part 1 & 2 solution is to get set of characters from the split string, then get the intersection
// of the sets which will be the common char and then convert the char according to priority formula.
fn main() {
    let input = read_to_string("../input.txt").expect("File cannot be opened.");
    let compartments_tuple = input
        .split("\n")
        .map(|compartments| compartments.split_at(compartments.len() / 2));
    let mut priority_sum = 0;
    for tuple in compartments_tuple {
        let c0: HashSet<char> = tuple.0.chars().collect();
        let c1: HashSet<char> = tuple.1.chars().collect();
        let intersection = c1.intersection(&c0);

        for same_item in intersection {
            priority_sum += get_priority(same_item.clone())
        }
    }
    println!("Day 3");
    println!("==============================");
    println!("Part 1 - The sum of the priorities is: {priority_sum}");

    let lines: Vec<&str> = input.split("\n").collect();
    let mut i = 0;
    let mut group_priority_sum = 0;
    while i < lines.len() {
        let l1: HashSet<char> = lines[i].chars().collect();
        let l2: HashSet<char> = lines[i + 1].chars().collect();
        let l3: HashSet<char> = lines[i + 2].chars().collect();

        let intersection = &(&l1 & &l2) & &l3;
        for same_item in intersection {
            group_priority_sum += get_priority(same_item.clone())
        }
        i += 3;
    }
    println!("Part 2 - The sum of the group badges priorities is: {group_priority_sum}");
}
