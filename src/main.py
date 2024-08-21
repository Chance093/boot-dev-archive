from block_markdown import markdown_to_blocks


def main():
    text = """
            # This is a heading

            This is a paragraph of text. It has some **bold** and *italic* words inside of it.

            * This is the first list item in a list block
            * This is a list item
            * This is another list item
            """

    block_nodes = markdown_to_blocks(text)
    print(block_nodes)


main()
