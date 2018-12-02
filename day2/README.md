# Day 2

**Time Taken:** 12 minutes

Part 2 could be massively optimized. (To be done)

Rather than doing `O(n^2)` search over the entire problem space,
you can have a tree to which you sequentially add each line to it.
Every character gets added as a leaf below the previous one.
When adding to the tree, when you hit the first character that doesn't match currently inserting character,
you do a search over the leafs and compare all the way down to the end of current line.
If the rest of the line matches any of the searched paths, you are done searching.
If not, revert back to the start and insert the line normally into the tree.