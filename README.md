# Music filtering application
As I usually tend to have several playlists of music depending on the context (music for workout, commuting, family trip, etc.), I have create a simple CLI application to create playlist from a `json` file.

## Considerations
- Songs must be locally stored
- Song name's must follow this format: `Band - Song`
- Format of the song is not important: `mp3`, `wma` ...

## Generate a template `json`
As it may be tedious to write a `json` for all the songs, there exists the possibility of creating one automatically by using the following command:

```
music template
```
Or
```
music template -s Directory/
```

Automatically, it scans `Music` folder and creates a file called `Music.json` with the appropriate format.

## Generate a playlist from a `json` file
It requires to have a `json` file properly formatted. 

### Shortcuts
- `*` : To select every single song of a band. Instead of writing all the list of songs this symbol can be used
- `~` : Not to select any song of a band


### Example of use
```
music filter -f workout_music.json -s Music -d Workout
```