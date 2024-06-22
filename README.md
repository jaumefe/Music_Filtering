# Music filtering application
As I usually tend to have several playlists of music depending on the context (music for workout, commuting, family trip, etc.), I have create a simple CLI application to create playlist from a `json` file.

## Considerations
- Main music folder must be called: `Music`
- Songs must be locally stored
- Song name's must follow this format: `Band - Song`
- Format of the song is not important: `mp3`, `wma` ...

## Generate a template `json`
As it may be tedious to write a `json` for all the songs, there exists the possibility of creating one automatically by using the following command:

```
./music_filter -t
```
Or
```
./music_filter --template
```

Automatically, it scans `Music` folder and creates a file called `All_Music.json` with the appropriate format.

## Generate a playlist from a `json` file
It requires to have a `json` file properly formatted. 
The field `playlist` in `json` file determines the name of the new folder that is going to be created by this app.

### Shortcuts
- `*` : To select every single song of a band. Instead of writing all the list of songs this symbol can be used
- `~` : Not to select any song of a band


### Example of use
```
music_filter -c workout_music.json
```
or
```
music_filter --config=workout_music.json
```