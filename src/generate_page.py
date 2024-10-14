from block_markdown import markdown_to_blocks, block_to_block_type

def extract_title(markdown):
    blocks = markdown_to_blocks(markdown)
    for block in blocks:
        block_type = block_to_block_type(block)
        if block_type == "heading":
            heading = block.lstrip("# ")
            return heading
        else:
            pass

    raise Exception("No header in markdown")

def generate_page():
    return
