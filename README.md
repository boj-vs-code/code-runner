# code-runner

The collection consists of code runners to provide same environment for users.

## Requirements

- jinja2-cli[yaml]
- docker

## Note

*DONT* believe dockerfiles keep up to date.  
Please use the following command to get the latest dockerfiles and to build the dockerfiles.
The docker images will be tagged with following format, `code-runner-<platform>-<language>`.

```bash
$ ./scripts/bake_all.sh <platform>  # (e.g. boj)
```

The script is written based on expect to be executed at the root of this project.
