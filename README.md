# Poker
Welcome to Poker Night with LARVIS.
This is a basic Poker module written in Go to add Poker playing abilities into LARVIS.

## Security details
- Unbiased deck shuffling.
- Random and secret card drawing.
- Tie resolution.
- No gambling.

## Setup
Download and install Go latest version.
Download Docker (optional)

## Libraries
1. fmt - printing
2. math/rand - for shuffling cards, generating random numbers and resetting time
3. sort - sorting hand values in ascending order
4. time - delay printing to add drama before starting the game and resolving ties

## Flow of the game
1. Draw a deck of card with symbols = 23456789TJQKA and their ranks = 2,3,4,5,6,7,8,9,10,11,12,13,14
2. Shuffle the Deck
3. Distribute 5 cards to each player- Hand1 and Hand2
4. Compare both Hands to check which hand has higher combination value as per these rules :-
    Combinations and their Rank in order of value:
    - 1= Four of a kind, like 77377
    - 2= Full house, means 3 of a kind, and 2 of a kind, in the same hand, like KK2K2
    - 3= Triple, like 32666
    - 4= Two pairs, like 77332
    - 5= A pair, like 43K9K
    - 6= High card, when there’s none of the above, like 297QJ
5. Creates frequency map of each symbol and how many times it occurs. eg. {J:4,3:1}
6. Calculate length of the slice[] of symbols from frequency map. Switch(length value) to compare the following rules :-
    - if len=2 and [4][1] or [1][4] : rank 1= Four of a kind
    - if len=2 and [3][2] or[2][3]  : rank 2= Full house
    - if len=3 and [3][1][1],[1][3][1] and [1][1][3] : rank 3= Triple
    - if len=3 and [2][2][1],[2][1][2] and [1][2][2] : rank 4=Two Pairs
    - if len=4 and [2][1][1][1],[1][2][1][1],[1][1][2][1] and [1][1][1][2] : rank 4=A Pairs
    - if len=5 and [1][1][1][1][1]  : rank 6=High Card
7. Compare rank of Hand1 vs Hand2 and declare winner
8. In case of a tie, where both hands have same rank value, call resolveTie(). 
9. Resolving a tie uses sortAndCompare() and findHighestCard()
10. Declare winner and show hands of both players.

## Building on local system

- Unzip files.
- CD to the path of Poker_Night folder.
- go run pokernight.go

## Deploying as a Docker Container
- Open a terminal and navigate to the directory where the Dockerfile is located.
- Run the following commands to build the Dockerfile and then run the program in a container in an interactive view.
    - docker build --pull --rm -f "Dockerfile" -t pokernight:latest "." 
    - docker run -it --name poker-night pokernight



## Disclaimer
We have not incorporated Suits (heart,spades,ace and diamond)
