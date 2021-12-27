# Cptemplate a template copying program

## Template Storage
Template files will be embeded with the program. These files are going
to be fairly small so this shouldn't really be a performance issue. Embedding
also means we don't have to worry about file integrity since the program is
accessing anything outside of itself.

### Files to store
1. Makefile (this should simply run the program once)
2. Sourcefile (the relevant source code template)
3. Notes file (a note file with the program name should be universal, but optional to include)

## Switches/Params
- `-l` language (should support C,Python, Go, and blank)
- `-N` no notes (notes should be included by default)
- `-n` Program name (this is what will be used to replace our keyword and name the folder)
- `-p` path to location to place template folder (current dir by default)

## Process
Given a the command `cptemplate -l c -n example -p .`
### In current dir
1. create a folder called example
2. create a source file `example.c`
3. create a `Makefile` that compiles and runs `example.c`
4. create a notes file `exampleNotes.md`

The keyword `$replaceme$` should be replaced with `example` in all files created.

If "blank" is given for the language arg just exclude steps 2,3.

## Building
Template files for build should be stored under `templates/<lang>`.

For example:
- `templates/c/source.c`
- `templates/c/Makefile`

The notes file should be stored at the root of templates `templates/programNotes.md`

