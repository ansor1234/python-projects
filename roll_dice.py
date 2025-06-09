import random

while True:
    choice = input('Roll the dice? (Y/N):\n               ')
    choice = choice.capitalize()
    print(choice)
    if choice == 'N':
        print('Bye!')
        break
    elif choice == 'Y':
        die = random.randint(1, 6)
        die2 = random.randint(1, 6)
        print(f'({die},{die2})')
    else:
        print('Invalid input! Try again😊')
