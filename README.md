# code-runner

The collection consists of code runners to provide same environment for users.

## Requirements

- jinja2-cli[yaml]
- docker

## Note

*DONT* believe dockerfiles keep up to date.  
Please use the following command to get the latest dockerfile.

```bash
$ ./scripts
```

Now, `./scripts/bake.sh` supports only [Baekjoon Online Judge]. Also the script is written based on expect to be executed at the root of this project.

[Baekjoon Online Judge]: https://boj.kr