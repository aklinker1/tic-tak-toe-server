# Tic Tak Toe Server

Simple tic tac toe server that you play via http requests.

## Timeline

- **`14:00`** - Project scaffold setup (build scripts, docker, postgres), no server code, just prints "test" to the console and exits
- **`56:00`** - API planned, code generated, health check implemented as sanity check
- **`1:15:00`** - Database integrated
- **`1:38:00`** - Create game finished
- **`1:46:00`** - Quit game finished
- **`1:56:00`** - README finalized with playing instructions
- **`2:23:00`** - Switched output to board as text
- **`3:05:00`** - Can play moves
- **`4:01:00`** - Got AI setup, victory check function still broken, been working on that for 45 min lol I'm done

## Playing

### Positions

```txt
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
├───┼───┼───┤
│ 7 │ 8 │ 9 │
└───┴───┴───┘
```

### Playing

1. Create a game:
    ```txt
    $ http post :3000/games
    
    Game: {id}
    ┌───┬───┬───┐
    │   │   │ x │
    ├───┼───┼───┤
    │   │   │ o │
    ├───┼───┼───┤
    │   │   │   │
    └───┴───┴───┘
    IN_PROGRESS
    ```
1. Use the id of the game to make moves. After making a move, the computer will play a move
    ```txt
    $ http post :3000/games/{id}/move position==3

    Game: {id}
    ┌───┬───┬───┐
    │   │   │ x │
    ├───┼───┼───┤
    │   │   │ o │
    ├───┼───┼───┤
    │   │   │   │
    └───┴───┴───┘
    IN_PROGRESS
    ```
1. Once the game is over, the output will look like:
    ```txt
    $ http post :3000/games/{id}/move position==7

    Game: {id}
    ┌───┬───┬───┐
    │   │   │ x │
    ├───┼───┼───┤
    │   │ x │ o │
    ├───┼───┼───┤
    │ x │   │ o │
    └───┴───┴───┘
    Winner: P1!
    ```

### Quiting a game
To quit a game that is in progress:
```txt
$ http post :3000/games/{id}/quit
Thanks for playing!
```
