from block_markdown import markdown_to_blocks, block_to_block_type
from markdown_to_html import markdown_to_html_node

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

def generate_page(from_path, template_path, dest_path):
    print(f"Generating page from {from_path} to {dest_path} using {template_path}")

    md = open(from_path, "r")
    markdown = md.read()

    tmp = open(template_path, "r")
    template = tmp.read()

    html = markdown_to_html_node(markdown)
    title = extract_title(markdown)

    new_stuff = template.replace("{{ Title }}", title).replace("{{ Content }}", html);

    w = open(dest_path, "w")
    w.write(new_stuff)
