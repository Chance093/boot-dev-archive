import unittest

from block_markdown import (markdown_to_blocks)

class TestBlockMarkdown(unittest.TestCase):
    def test_markdown_to_blocks(self):
        text = "# This is a heading\n\nThis is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n* This is the first list item in a list block\n * This is a list item\n * This is another list item"
        block_nodes = markdown_to_blocks(text)
        self.assertListEqual(
            [
                "# This is a heading",
                "This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
                "* This is the first list item in a list block\n * This is a list item\n * This is another list item",
            ],
            block_nodes,
        )

    
    def test_markdown_to_blocks_with_whitespace(self):
        text = "# This is a heading \n\n This is a paragraph of text. It has some **bold** and *italic* words inside of it. \n\n * This is the first list item in a list block\n * This is a list item\n * This is another list item"
        block_nodes = markdown_to_blocks(text)
        self.assertListEqual(
            [
                "# This is a heading",
                "This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
                "* This is the first list item in a list block\n * This is a list item\n * This is another list item",
            ],
            block_nodes,
        )

if __name__ == "__main__":
    unittest.main()
