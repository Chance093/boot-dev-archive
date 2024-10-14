import unittest
from generate_page import extract_title

class TestGeneratePage(unittest.TestCase):
    def test_extract_title(self):
        markdown = """# This is a heading

            This is something else"""
        heading = extract_title(markdown)
        self.assertEqual(heading, "This is a heading")

    def test_extract_title_2(self):
        markdown = """This is not a heading

            # This is a heading"""
        heading = extract_title(markdown)
        self.assertEqual(heading, "This is a heading")

    def test_extract_title_3(self):
        markdown = """This is not a heading

            This is also not a heading

            # This is a heading"""
        heading = extract_title(markdown)
        self.assertEqual(heading, "This is a heading")

    def test_extract_title_4(self):
        markdown = """This is not a heading

            This is also not a heading

            NO heading"""
        with self.assertRaises(Exception):
            extract_title(markdown)
