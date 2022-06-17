# A Linear Solution

<!-- markdownlint-disable -->

This solution does not require nested loop. It creates a <em>helper</em> empty slice then loops to check each number (in this example the number has to be
between 0 to 249, hence the size given to the make() funnction) in the array. Each number value represents the index on the helper slice. As it encounter each number
it assigns an arbitrary value at the index of the number on the helper slice. It first checks if we have encountered the number, if we haven't, it will assign the
arbitrary value (1) to the index, if we have, meaning the index value is not 0, then we have encounter a duplicate and we can return from the function.

Worst Case: O(N) since we are just comparing all the values of the array once.

Drawback: We have increased the memory the program takes by creating a new structure.
