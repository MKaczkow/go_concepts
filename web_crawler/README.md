# Go Concepts
Repo for basic tutorial-based Golang study  

---

# web crawler

## todo / further improvements ideas
- [x] actual place HTML structure -> data model structure -> file 
  - [x] get name, html
  - [x] get comments
  - [x] get description
  - [x] get details
  - [x] get hazards
  - [x] loop
- [x] run different stages
  - [x] re-organize main
  - [x] add command parser to select different stages (drop - just comment out the stage you don't want to run)
- [ ] logging instead of printing
- [ ] add (maybe?) missing field filling in data model
  - [x] hazards
  - [ ] place itself
- [ ] time full runs and specific stages and specific ops
- [ ] maybe saving multiple files instead of one big file causes slowdown?
- [x] add basic tests
- [x] upgrade / cleanup data model (hazards, details, etc.)
- [x] fix 'brak głosów' issue
- [x] czy na pewno nie za często robi się update (tzn. N update do 1 inserta faktycznie)
- [x] new stage (3)
- [ ] fix stage 3 to actually save data in mongo and have proper names
- [x] change user-agent (colly?)

## notes
* run relevant mongodb container first (port and address matching .env file)

### basics
* run mongo container (via docker desktop - win goes brrr)
* run mongo compass and connect to `localhost:49153` (sanity check)
* crawl
