# Go shred

This is try to implement shreding functionality in go

Things that needs to be done:
1. Overwrite a file with random data 3 times
  - File size
2. Delete the file

Pitfals: Things that might be considered in advanced scnarios:
- Considering file system in shreding
  - Copy on write file systems: This is hard to answer. We might need to fill the empty space also with random data
  - Journaling file systems: Do we need to consider journaling process or not


Implementation approach:
1. start with a straight forward shell like approach
2. Restructure the code into modules and add tests
3. Improve