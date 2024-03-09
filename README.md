# Website

My personal website written in Go and `templ`. It is a static site that is deployed to Cloudflare.

## Dependencies

- [templ](https://templ.guide/)
- [tailwindcss](https://tailwindcss.com/docs/installation)

## Run

The application can be ran as a single binary or generate static files that can be deployed.

### Binary

Build the application and simply run it.

### Static

Simple website built using Go and templ.

To generate static files run,

```shell
go run main.go -gen -ver v1.0.0
```

Static file will be written to `./public`.

### Deploy

Run the `release.yml` workflow by specifying the version to deploy. Part of the workflow will copy all files in the 
`./public` directory to the `pages` git branch.

Cloudflare Pages is monitoring changes to the branch. When it detects a change, it will deploy the latest changes.
