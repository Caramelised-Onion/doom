import argparse
from datetime import date, timedelta
from random import randint
from typing import Optional


def get_date(start_date: Optional[date] = None, end_date: Optional[date] = None) -> date:
    if start_date is None:
        start_date = date(1600, 1, 1)

    if end_date is None:
        end_date = date(2400, 12, 31)

    diff_timedelta = end_date - start_date
    delta_days = randint(0, diff_timedelta.days)
    new_date = start_date + timedelta(days=delta_days)
    return new_date

def print_practice_results(correct, total):
    print()
    print("Finished Practising")
    print("-" * 80)
    print(f"{correct=}")
    print(f"{total=}")
    if total > 0:
        print(f"correct percentage: {100 * (correct/total)}%")

def practice():
    correct = 0
    total = 0

    while (guess_date := get_date()) and (inp := input(f"What day of the week was {guess_date}:\n")) != 'q'.strip():
        total += 1
        
        try: 
            inp_int = int(inp)
            correct_day = (1 + guess_date.weekday()) % 7 
            correct += correct_day == inp_int
        except:
            print("Please input a number from 0-6, 0 for Sunday 6 for Saturday.")
    print_practice_results(correct, total)


def anchor_year_pratice():
    correct = 0
    total = 0
    while (guess_date := get_date()) and (inp := input(f"What is the anchor day for {guess_date.year}:\n")) != 'q'.strip():
            total += 1
            
            try: 
                inp_int = int(inp)
                correct_ans = (1 + date(guess_date.year, 4, 4).weekday()) % 7 
                correct += correct_ans == inp_int
                print(bool(correct_ans == inp_int))
            except:
                print("Please input a number from 0-6, 0 for Sunday 6 for Saturday.")
    print_practice_results(correct, total)

def parse_args():
    parser = argparse.ArgumentParser(prog='doom')
    subparsers = parser.add_subparsers(help='Available subcommands', dest='subcommand')
    practice_parser = subparsers.add_parser('practice', help='Description of subcommand')
    practice_parser.add_argument('-m', '--mode', default='standard', choices=['standard', 'anchor-year'])
    gamemode_parser = subparsers.add_parser('gamemode', help='Description of subcommand')
    args = parser.parse_args()
    return args

def main():
    args = parse_args()
    if args.subcommand == 'practice':
        if args.mode == 'standard':
            practice()
        elif args.mode == 'anchor-year':
            anchor_year_pratice()
        else:
            print(f"Invalid mode: {args.mode}")

    else:
        print(f"Invalid subcommand: {args.subcommand}")

if __name__ == "__main__":
    main()
