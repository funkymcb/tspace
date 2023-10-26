# tspace
A small go CLI which helps managing tmux workspaces (sessions).  
You can export the current tmux state as json file and  
easily create new sessions by using the cli or importing jsons.  

# how to use
Initialize tspace by running `tspace init`.  
A prompt will ask for the location of the workspace json files.  
Default is `~/.config/tspace/workspaces/`.  

You can now either export your current tmux session by running `tspace export -o work`. This will create a `work.json` file in your configured path.  
Or you can create a json file manually.  

By running `tspace list`, all available workspaces (json files) will be listet.  

To start a new tmux-server with a given workspace config just run for example `tspace start work`.  
If you already are inside a tmux session while running this command, you will be prompted to either overwrite or append (to) the current session.

# example tspace json
```json
{
    "session_name": {
        "window1_name": "number of panes and split orientation (h/v)",
        "window2_name": "number of panes and split orientation (h/v)"
    },
    "clusters": {
        "prod": "1",
        "staging": "2h"
    },
    "code": {
        "api": "2v",
        "frontend": "1"
    }
}
```
