<p>对于字符串&nbsp;<code>s</code> 和&nbsp;<code>t</code>，只有在&nbsp;<code>s = t + t + t + ... + t + t</code>（<code>t</code> 自身连接 1 次或多次）时，我们才认定&nbsp;“<code>t</code> 能除尽 <code>s</code>”。</p>

<p>给定两个字符串&nbsp;<code>str1</code>&nbsp;和&nbsp;<code>str2</code>&nbsp;。返回 <em>最长字符串&nbsp;<code>x</code>，要求满足&nbsp;<code>x</code> 能除尽 <code>str1</code> 且 <code>x</code> 能除尽 <code>str2</code></em> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>str1 = "ABCABC", str2 = "ABC"
<strong>输出：</strong>"ABC"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>str1 = "ABABAB", str2 = "ABAB"
<strong>输出：</strong>"AB"
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>str1 = "LEET", str2 = "CODE"
<strong>输出：</strong>""
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= str1.length, str2.length &lt;= 1000</code></li> 
 <li><code>str1</code>&nbsp;和&nbsp;<code>str2</code>&nbsp;由大写英文字母组成</li> 
</ul>

<div><div>Related Topics</div><div><li>数学</li><li>字符串</li></div></div><br><div><li>👍 442</li><li>👎 0</li></div>