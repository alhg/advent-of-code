use std::fs::read_to_string;

enum MatchOutcome {
    Win,
    Lose,
    Draw,
}

impl MatchOutcome {
    fn score(self) -> i32 {
        match self {
            Self::Win => 6,
            Self::Lose => 0,
            Self::Draw => 3,
        }
    }
}

#[derive(PartialEq)]
enum HandType {
    Rock,
    Paper,
    Scissor,
}

impl HandType {
    fn get_winning_hand(self) -> Self {
        match self {
            Self::Rock => Self::Paper,
            Self::Paper => Self::Scissor,
            Self::Scissor => Self::Rock,
        }
    }
    fn get_losing_hand(self) -> Self {
        match self {
            Self::Rock => Self::Scissor,
            Self::Paper => Self::Rock,
            Self::Scissor => Self::Paper,
        }
    }
    fn is_defeat_by(self, other_hand: HandType) -> bool {
        other_hand == self.get_winning_hand()
    }
    fn score(self) -> i32 {
        match self {
            Self::Rock => 1,
            Self::Paper => 2,
            Self::Scissor => 3,
        }
    }
}

fn get_handtype(hand: char) -> HandType {
    match hand {
        'A' | 'X' => HandType::Rock,
        'B' | 'Y' => HandType::Paper,
        'C' | 'Z' => HandType::Scissor,
        _ => panic!("Invalid hand type {}", hand),
    }
}

fn get_outcome_from_char(char_outcome: char) -> MatchOutcome {
    match char_outcome {
        'X' => MatchOutcome::Lose,
        'Y' => MatchOutcome::Draw,
        'Z' => MatchOutcome::Win,
        _ => panic!("Invalid outcome: {char_outcome}"),
    }
}

fn get_outcome(opp_hand: char, your_hand: char) -> MatchOutcome {
    let opp_hand = get_handtype(opp_hand);
    let your_hand = get_handtype(your_hand);

    if opp_hand == your_hand {
        return MatchOutcome::Draw;
    }

    if opp_hand.is_defeat_by(your_hand) {
        MatchOutcome::Win
    } else {
        MatchOutcome::Lose
    }
}

fn get_score(outcome: MatchOutcome, your_hand: HandType) -> i32 {
    return outcome.score() + your_hand.score();
}

fn main() {
    let input = read_to_string("../input.txt").expect("File cannot be opened.");
    let mut total_score = 0;
    // each line is in the format "A X". A represents opponents hand, X represents your hand
    for line in input.split("\n").filter(|line| line.len() == 3) {
        let opp_hand = line.chars().next().unwrap();
        let your_hand = line.chars().nth(2).unwrap();

        total_score += get_score(get_outcome(opp_hand, your_hand), get_handtype(your_hand));
    }
    println!("Day 02:");
    println!("====================");
    println!("Part 1: The total score is {total_score}.");

    let mut total_score = 0;
    // each line is in the format "A X". A represents opponents hand, X represents desired match outcome.
    for line in input.split("\n").filter(|line| line.len() == 3) {
        let opp_hand = get_handtype(line.chars().next().unwrap());
        let desired_outcome = line.chars().nth(2).unwrap();

        let match_outcome = get_outcome_from_char(desired_outcome);
        let your_hand = match match_outcome {
            MatchOutcome::Draw => opp_hand,
            MatchOutcome::Win => opp_hand.get_winning_hand(),
            MatchOutcome::Lose => opp_hand.get_losing_hand(),
        };
        total_score += get_score(get_outcome_from_char(desired_outcome), your_hand)
    }
    println!("Part 2: The total score is {total_score}.")
}
