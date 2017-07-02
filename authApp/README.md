# Auth0 + Go Web App Sample
From https://github.com/auth0-samples/auth0-golang-web-app

## Running the App

To run the app, make sure you have **go** and **goget** installed.

Create a `.env` and provide your Auth0 credentials.

```bash
# .env

AUTH0_CLIENT_ID=XOvYQUGgxC4mIaPWfsJVjo_gZYVsDRoU
AUTH0_DOMAIN=anilgkurian.auth0.com
AUTH0_CLIENT_SECRET=XemecrZbinzfzWRrUt0pjDw6vrpFesQqFBb6EaPhBVj5iHh6o5Zh5eLnkWVbEYl0
AUTH0_CALLBACK_URL=http://localhost:3000/callback
```

Once you've set your Auth0 credentials in the `.env` file, run `go get .` to install the Go dependencies.

Run `go run main.go server.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/)
