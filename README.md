# go-htmx

This is meant to be a basic example of a dynamic site using the go+htmx stack. If/when I decide on a specific application for this, I will most likely rename the repo.

This is a pretty blatant clone of awesome_club's [go-htmx project](https://github.com/awesome-club/go-htmx). I wanted to follow what he had done [here]() to make sure I understood the basics, then expand it into my own thing.
Loose plan from here is:
- Change to use tailwind for styling
- Build homepage (BoweGo?)
- Change existing stocks stuff to subpath under homepage
- Build out other subpaths, i.e:
  - [books](https://openlibrary.org/developers/api)
  - [exchange rates](https://github.com/BoweFlex/CurrencyConverter)
  - [more apis](https://github.com/public-apis/public-apis?tab=readme-ov-file#index)

## Tailwind
Tailwind is being used for the CSS in this project. For more information, go to [their website](https://tailwindcss.com/).

Since I'm using Go and HTMX for everything else, the CLI tool is being used to avoid a Node.js/npm dependency for just tailwind. Specific installation instructions for this [here](https://tailwindcss.com/blog/standalone-cli).

This is currently the only project I'm using Tailwind in, so I'm placing the CLI tool in the git repo under `web/` and adding it to the .gitignore. If this becomes a more common solution, steps will be added to ensure it is installed and placed in $PATH.

For now, the following commands can be run from inside `web/`:

```bash

# Create a tailwind.config.js file
./tailwindcss init

# Start a watcher
./tailwindcss -i input.css -o output.css --watch

# Compile and minify your CSS for production
./tailwindcss -i input.css -o output.css --minify

```
