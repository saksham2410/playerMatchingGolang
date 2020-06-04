Description

Possible Algo ->

Players who enter the matching queue all have a rank score. After starting the match, players with similar scores are matched as much as possible. If the waiting time is longer, you can match players with larger score gaps. Players with similar scores will be given priority to start the game.

Algorithm
The score approx interval increases with time, but it is not necessarily linear. The score approx interval may have the maximum range.

There might be an estimated matching time, which can be calculated based on the match time of players with similar scores recently and the number of people in the queue

There is a maximum time, after which we can cancel a match or forcibly start the game

Matching attempts are made in order of joining time. Players who join first shall match first.
Get the first player to join in the score approx interval (with similar scores) until the number of people who get enough room stops or the number of people in the score interval stops.
