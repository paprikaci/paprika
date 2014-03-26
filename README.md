# paprika #

`paprika` is a command line utility for executing builds remotely within Paprika
CI server. The whole process goes through the following steps:

1. `paprika` connects to the specified master node, authenticating using the
   specified access token.
2. Using the required arguments, which is the repository URL and the path of the
   script to be run relative to the repository root, a Cider RPC call is issued
   to the master node, which is then forwarded to one of the build slaves.
3. The build output is streamed from the build slave to the console.
4. `paprika` terminates according to the build results. Failing build means a
   non-zero exit status.

## Request Routing ##

You might be wondering how the RPC requests are routed to the build slaves. It
is rather simple. The RPC method name is generated from the specified slave tag
and the script to be run. To be more specific, the file extension of the chosen
script is taken and the whole thing is concatenated in the following way:
`<slave>.<script-file-extension>`. All that is necessary is that there is a
build slave connected to the master node that exports a method of the given
name, which basically means that it can run the scripts of the given kind. For
example, a RPC method of `macosx109.bash` would mean that there is a build slave
labeled `macosx109` and it can run Bash scripts. The whole RPC method name is,
however, just a useful convention to be able to use Cider routing out of the
box.

## Documentation ##

```text
TBD
```

## Example ##

`paprika` can parse environment variables when run on Circle CI. For now,
however, we assume that all the parameters must be specified explicitly.

```bash
$ paprika -verbose \
          -slave='macosx109' \
          -repository='git+https://github.com/salsita-cider/paprika' \ 
		  -env FOO=bar -env ROOT=toor
		  'scripts/upload.sh'
```

The output could then be

```bash
TBD
```

## License ##

GNU GPLv3, see the `LICENSE` file.