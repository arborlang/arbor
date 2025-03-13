# ArborLang

This is the monorepo for the arborlang project.

## Repo Structure

There are two root directories here: `apps/` and `libs/`. `apps/` is intended to contain executable/deployable code. This is where the code for the compiler entry point will be, as well as the documentation website. `libs/` is intended to contain shared code between all executables/deployables. This is also where consumable libraries will be contained. Most code should probably be in this repository.

## Tooling

This repo uses `NX` to manage the monorepo.
