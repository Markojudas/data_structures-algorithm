# Sorting Algorithm: Insertion Sort

<!-- markdownlint-disable -->

<h4>Analyzing scenarios beyond the worse case</h4>
<br>
<h2>Steps:</h2>
<ol>
    <br>
    <li>In the first pass-through, we temporarily remove the value at index 1 (the second cell) and sotre it in a temporary variable. This will leave a gap at that index, since it contains no value. In subsequent pass-throughs, we remove the values at the subsequent indexes.</li>
    <br>
    <li>We then begin a shifting phase, where we take each value to the left of the gap and compare it to the value in the termporary variable. If the value to the left of the gap is greater than the temporary variable, we shift that value to the right. As we shift values to the right, inherently the gap moves leftward. As soon as we encounter a value that is lower than the temporarily removed value, or we reach the left end of the array, this shiftting phase is over.</li><br>
    <li>We then insert the temporary removed value into the current gap.</li><br>
    <li>Steps 1 through 3 represent a single pass-through. We repeat these pass-throughs unitl the pass-through begins at the final index of the array. By then, the array will have been fully sorted</li>
</ol>
<br>
<hr>
<br>
<h2>Code Implementation: Insertion Sort - Golang</h2>

```go
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp_val := arr[i]
		left := i - 1

		for left >= 0 {
			if arr[left] > temp_val {
				arr[left+1] = arr[left]
				left--
			} else {
				break
			}
		}
		arr[left+1] = temp_val
	}
}
```

Let's walk through the code step by step:

<ul>
    <li>First, we start a loop beginning at index 1 that runs through the entire array. Each round of this loop reresents a pass-through:<br><br>
    <code>for i := 1; i < len(arr); i++</code><br><br>
    Within each pass-through, we save the value we're "removing" in a variable called <code>temp_value</code><br><br>
    <code>temp_val := arr[i]</code></li><br>
    <li>Next, we create a variable called <code>left</code>, which will start immediately to the left of the index of the <code>temp_value</code>. This <code>left</code> variable will represent each value we compare against the <code>temp_value</code><br><br>
    <code>left := i - 1</code><br><br>
    As we move through the pass-through, this <code>left</code> variable will keep moving leftward as we compare each value to the <code>temp_value</code>.</li><br>
    <li>We then begin an inner <code>for</code> loop, which runs as long as the <code>left</code> variable is greater or equal to 0<br><br>
    <code>for left >= 0</code><br><br></li>
    <li>We then perform our comparison. That is, we check whether the value at <code>left</code> is greater than the <code>temp_value</code>:<br><br>
    <code>if arr[left] > temp_val</code><br><br>
    If it is, we shift that left value to the right:<br><br>
    <code>arr[left+1] = arr[left]</code><br></li><br>
    <li>We then decrement <code>left</code> by 1 to compare the next left value against the <code>temp_value</code> in the next round of the <code>for</code> loop:<br><br>
    <code>left--</code></li><br>
    <li>if at any point we encounter a value at <code>left</code> that is greater or equal to the <code>temp_value</code>, we can get ready to end our pass-through, since it's time to move the <code>temp_value</code> into the gap:<br><br>
    <pre><code>else {
    break
}</code></pre></li>
    <li>The final step of each pass-through is moving the <code>temp_value</code> into the gap:<br><br>
    <code>arr[left+1] = temp_value</code></li><br>
    <li>After all pass-throughs have been completed, we return the sorted array<br>
    <h5>In Go array or slices are pass by reference type, no need to return it</h5></li>
</ul>
