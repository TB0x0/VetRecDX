# VetRecDX

An open source veterinary records management software for nonprofits

## Current Tasks

- [ ] API Endpoints and Backend Functions
  - [ ] API Endpoints
  - [ ] DB Helper Functions
  - [ ] Password Hashing Functions

- [x] Planning and Design
  - [x] App Architecture
  - [x] DB Schema
  - [x] Dependencies

## Contributing

  We're using air for live editing and debugging. If you don't have air installed run

  `go install github.com/air-verse/air@latest`

  Once installed navigate to the project directory and run the air command in your terminal

  Do your edits and submit a PR and it'll get added to main after review

  ### Database Connection

  Assuming you already have a workable postgres instance.
  Create your new Dev DB on localhost
  Create .env file in the root of the VetRecDX directory (Use the example in planning/envexample)
  Run the migration scripts on your DB
