# Trashdiena

Trashdiena was born from the Latvian word "Trešdiena", which means "Wednesday".

Because "Treš" part sounds as "Trash" in English, the day when IT jokes (and not only) in English and Russian languages were shared in the chat was named "Trashdiena".

And this is how this bot was born. Every Wednesday in the Slack channel of your choice will be delivered some trash that is NSFW (mostly). This trash will ruin your productivity. You will love it.

Enjoy!

## Development

To install all dependencies, run: `dep ensure`

Directory structure is following:

```
├── Dockerfile  # Dockerfile to build the image
├── Gopkg.lock  # Dep .lock file
├── Gopkg.toml
├── Makefile    # Makefile to automate some stuff
├── README.md   # README
├── cmd         # All the binaries are build from this directory
├── pkg         # All the packages that are mandatory for apps and can be shared
└── vendor      # Vendor dependencies
```

Binaries:

 * authorizer - Do authorization with Slack and save credentials. Open to the web.
 * scheduler - Distribute trash accross registered Slack hooks in the system in specific time.
 * loader - Separate binary to load trash in the DB
 