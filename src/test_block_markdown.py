import unittest

from block_markdown import block_to_block_type, markdown_to_blocks


class TestMarkdownToHTML(unittest.TestCase):
    def test_markdown_to_blocks(self):
        md = """
This is **bolded** paragraph

This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
        blocks = markdown_to_blocks(md)
        self.assertEqual(
            blocks,
            [
                "This is **bolded** paragraph",
                "This is another paragraph with *italic* text and `code` here\nThis is the same paragraph on a new line",
                "* This is a list\n* with items",
            ],
        )

    def test_markdown_to_blocks_newlines(self):
        md = """
This is **bolded** paragraph




This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
        blocks = markdown_to_blocks(md)
        self.assertEqual(
            blocks,
            [
                "This is **bolded** paragraph",
                "This is another paragraph with *italic* text and `code` here\nThis is the same paragraph on a new line",
                "* This is a list\n* with items",
            ],
        )

    def test_block_to_block_type_heading(self):
        block = "#This is a heading"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "heading")

    def test_block_to_block_type_heading_2(self):
        block = "###This is a h3"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "heading")

    def test_block_to_block_type_code(self):
        block = "```This is a block of code```"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "code")
        
    def test_block_to_block_type_code_2(self):
        block = "```This is a line of code\nThis is another line of code```"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "code")

    def test_block_to_block_type_quote(self):
        block = ">This is a quote block"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "quote")

    def test_block_to_block_type_quote_2(self):
        block = ">This is a quote block\n>This is the next line"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "quote")

    def test_block_to_block_type_quote_paragraph(self):
        block = ">This is a quote block\nThis is the next line"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "paragraph")

    def test_block_to_block_type_unordered_list(self):
        block = "* Unordered list"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "unordered_list")
    
    def test_block_to_block_type_unordered_list_2(self):
        block = "- Unordered list"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "unordered_list")
    
    def test_block_to_block_type_unordered_list_3(self):
        block = "* Unordered list\n* 2nd list item"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "unordered_list")

    def test_block_to_block_type_unordered_list_4(self):
        block = "- Unordered list\n- 2nd list item"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "unordered_list")

    def test_block_to_block_type_unordered_list_paragraph(self):
        block = "* Unordered list\n2nd list item"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "paragraph")

    def test_block_to_block_type_ordered_list(self):
        block = "1. First order"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "ordered_list")
        
    def test_block_to_block_type_ordered_list_2(self):
        block = "1. First order\n2. Second order"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "ordered_list")

    def test_block_to_block_type_ordered_list_paragraph(self):
        block = "1. First order\n3. Second order"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "paragraph")

    def test_block_to_block_type_paragraph(self):
        block = "This is a paragraph"
        block_type = block_to_block_type(block)
        self.assertEqual(block_type, "paragraph")


if __name__ == "__main__":
    unittest.main()
