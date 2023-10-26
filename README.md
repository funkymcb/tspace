# tspace
A small go CLI which helps managing tmux workspaces (sessions).  
You can export the current tmux state as json file and  
easily create new sessions by using the cli or importing jsons.  

# example tspace json
```json
{
    "clusters": {
        "prod": 1,
        "staging": 2
    },
    "code": {
        "api": 2,
        "frontend": 1
    },
    "session_name": {
        "window1_name": "number of panes as int",
        "window2_name": "number of panes as int"
    }
}
```
