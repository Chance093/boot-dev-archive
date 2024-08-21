def markdown_to_blocks(text):
    split_strings = text.split("\n\n")
    blocks = map(lambda x: x.strip(), split_strings)
    return list(blocks)
