# quick

`quick` is a simple tool to rapidly add aliases to your `.bashrc` or `.bash_profile`. 

This is relatively new software, so I'd strongly suggest backing up your bash configs before using it, it may kill them.

For `quick` to work, the `HISTFILE` variable must be exported. This may not happen by default in your shell. To check if it is exported, run `export -p | grep HISTFILE`.

If the output does not contain the `HISTFILE` var, add the following to your `.bashrc` (linux) or `.bash_profile` (OS X):

    export HISTFILE=$HISTFILE

