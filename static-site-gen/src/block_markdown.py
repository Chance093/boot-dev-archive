def markdown_to_blocks(text):
    split_strings = text.split("\n\n")
    blocks = []
    for split_string in split_strings:
        if split_string == "":
            continue
        blocks.append(split_string.strip())

    return blocks


def block_to_block_type(block):
    if (
        block.startswith("#")
        or block.startswith("##")
        or block.startswith("###")
        or block.startswith("####")
        or block.startswith("#####")
        or block.startswith("######")
    ):
        return "heading"

    elif block.startswith("```") and block.endswith("```"):
        return "code"

    elif block.startswith(">"):
        split_lines = block.split("\n")
        for line in split_lines:
            if not line.startswith(">"):
                return "paragraph"

        return "quote"

    elif block.startswith("*"):
        split_lines = block.split("\n")
        for line in split_lines:
            if not line.startswith("* "):
                return "paragraph"

        return "unordered_list"

    elif block.startswith("-"):
        split_lines = block.split("\n")
        for line in split_lines:
            if not line.startswith("- "):
                return "paragraph"

        return "unordered_list"

    elif block.startswith("1."):
        split_lines = block.split("\n")
        for i in range(1, len(split_lines) + 1):
            if not split_lines[i - 1].startswith(f"{i}."):
                return "paragraph"

        return "ordered_list"
        

    return "paragraph"
