<p>给你一个 <strong>正</strong> 整数 <code>n</code> 。</p>

<p>用 <code>even</code> 表示在 <code>n</code> 的二进制形式（下标从 <strong>0</strong> 开始）中值为 <code>1</code> 的偶数下标的个数。</p>

<p>用 <code>odd</code> 表示在 <code>n</code> 的二进制形式（下标从 <strong>0</strong> 开始）中值为 <code>1</code> 的奇数下标的个数。</p>

<p>请注意，在数字的二进制表示中，位下标的顺序&nbsp;<strong>从右到左</strong>。</p>

<p>返回整数数组<em> </em><code>answer</code><em> </em>，其中<em> </em><code>answer = [even, odd]</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block"> 
 <p><span class="example-io"><b>输入：</b>n = 50</span></p> 
</div>

<p><span class="example-io"><b>输出：</b>[1,2]</span></p>

<p><strong>解释：</strong></p>

<p>50 的二进制表示是&nbsp;<code>110010</code>。</p>

<p>在下标 1，4，5 对应的值为&nbsp;1。</p>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block"> 
 <p><strong>输入：</strong><span class="example-io">n = 2</span></p> 
</div>

<p><span class="example-io"><b>输出：</b>[0,1]</span></p>

<p><strong>解释：</strong></p>

<p>2 的二进制表示是&nbsp;<code>10</code>。</p>

<p>只有下标 1 对应的值为&nbsp;1。</p>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= n &lt;= 1000</code></li> 
</ul>

<div><div>Related Topics</div><div><li>位运算</li></div></div><br><div><li>👍 31</li><li>👎 0</li></div>