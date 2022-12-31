use std::fs;

fn main() {
    let input = fs::read_to_string("../input.txt").expect("Unable to read input.txt");

    let elf_calories_group = input.split("\n\n");
    let mut largest_calories = -1;
    for calories_group in elf_calories_group {
        let mut sum = 0;
        for calories in calories_group.split("\n") {
            if calories.len() == 0 {
                continue;
            }
            let parsed_cal = calories.parse::<i64>().unwrap();
            sum += parsed_cal;
        }
        if largest_calories < sum {
            largest_calories = sum;
        }
    }

    println!("Day 1");
    println!("==============================================");
    println!(
        "Part 1: The largest calories carried by an elf is: {}",
        largest_calories
    );

    let input = fs::read_to_string("../input.txt").expect("Unable to read input.txt");
    let mut calories: Vec<i64> = input
        .split("\n\n")
        .map(|elf_calories| {
            elf_calories
                .split("\n")
                .filter(|cal| cal.len() > 0)
                .fold(0, |sum: i64, cal: &str| cal.parse::<i64>().unwrap() + sum)
        })
        .collect();
    calories.sort_by(|a, b| b.cmp(a));

    println!(
        "Part 2: The sum of the 3 largest calories from the elves are: {}",
        calories[0] + calories[1] + calories[2]
    );
}
