def markdown_to_blocks(text):
    split_strings = text.split("\n\n")
    blocks = []
    for split_string in split_strings:
        if split_string == "":
            continue
        blocks.append(split_string.strip())

    return blocks
