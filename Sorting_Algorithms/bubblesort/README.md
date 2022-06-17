# Efficiency of Bubble Sort

<!-- markdownlint-disable -->

<h2>Bubble sort contains two significant kinds of steps:</h2>

<ol>
<li><strong><en>Comparisons</en>: Two numbers are compared with one another to determine which is greater.</strong></li>
<p>Using the code example, there are 7 elements. The first passthrough did a total of 6 comparisons, the second passthrough we did 5 comparisons, and the third we did 4 comparions, so forth until the last passthrough where we did only 1 comparison. (N - 1) + (N - 2) + (N - 3)... 1<p>
6+5+4+3+2+1 = 21 Comparions

<li><strong><en>Swaps</en>: Two numbers are swapped with one another in order to sort them.</strong></li>
<p>Since that array was originally sorted in Descending order (where we want to sort it in ascending) we need a swap for each comparison (worst-case).</p>

num of swaps: 21 swaps

</ol>

We have a total of 42 steps. Therefore the time complexity of bubble sort is approximeately O(N<sup>2</sup>) (Quadratic Time) : N = 7; 7<sup>2</sup> = 49
