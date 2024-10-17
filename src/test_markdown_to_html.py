import unittest

from markdown_to_html import markdown_to_html_node

class TestMarkdownToHTML(unittest.TestCase):
    def test_markdown_to_html(self):
        md = """
This is **bolded** paragraph

This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
        html_answer = "<div><p>This is <b>bolded</b> paragraph</p><p>This is another paragraph with <i>italic</i> text and <code>code</code> here\nThis is the same paragraph on a new line</p><ul><li>This is a list</li><li>with items</li></ul></div>"

        html = markdown_to_html_node(md)
        self.assertEqual(html, html_answer) 

    def test_markdown_to_html_2(self):
        md = """
This is **bolded** paragraph

> This is another paragraph with *italic* text and code here
> This is the same paragraph on a new line

1. This is a list
2. with items
"""
        html_answer = "<div><p>This is <b>bolded</b> paragraph</p><blockquote>This is another paragraph with <i>italic</i> text and code here\nThis is the same paragraph on a new line</blockquote><ol><li>This is a list</li><li>with items</li></ol></div>"

        html = markdown_to_html_node(md)
        self.assertEqual(html, html_answer) 

    def test_markdown_to_html_3(self):
        md = """
This is **bolded** paragraph

```
This is another paragraph with italic text and code here
This is the same paragraph on a new line
```
"""
        html_answer = "<div><p>This is <b>bolded</b> paragraph</p><pre><code>This is another paragraph with italic text and code here\nThis is the same paragraph on a new line</code></pre></div>"

        html = markdown_to_html_node(md)
        self.assertEqual(html, html_answer) 

    def test_headings(self):
        md = """
# this is an h1

this is paragraph text

## this is an h2
"""

        html = markdown_to_html_node(md)
        self.assertEqual(
            html,
            "<div><h1>this is an h1</h1><p>this is paragraph text</p><h2>this is an h2</h2></div>",
        )


if __name__ == "__main__":
    unittest.main()
