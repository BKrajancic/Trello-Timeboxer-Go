# Trello-Timeboxer-Go
Similar to my other trello timeboxer project. 

This is written in Go instead of Python, and makes use of channels
for concurrency.

# How to run
Edit the file in `src/_config.json`, then rename it to `config.json` (keep it
in `src`).

In command prompt, go to this project directory and run:
`docker build .` At the end there should be `Successfully built <image id>`.

Then run `docker run <image id>`