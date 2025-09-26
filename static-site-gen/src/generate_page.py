from block_markdown import markdown_to_blocks, block_to_block_type
from markdown_to_html import markdown_to_html_node
import os

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

def generate_page_recursive(dir_path_content, template_path, dest_dir_path):
    print(f"Generating page from {dir_path_content} to {dest_dir_path} using {template_path}")

    tmp = open(template_path, "r")
    template = tmp.read()

    copy_files_to_public(dir_path_content, template, dest_dir_path)


def copy_files_to_public(parent_path, template, dest_path):
    current_dir = os.listdir(path=parent_path)
    for file_or_dir in current_dir:
        current_path = f"{parent_path}/{file_or_dir}"
        if os.path.isfile(path=current_path):
            md = open(current_path, "r")
            markdown = md.read()

            html = markdown_to_html_node(markdown)
            title = extract_title(markdown)

            content = template.replace("{{ Title }}", title).replace("{{ Content }}", html)

            w = open(f"public/{current_path.rstrip(".md").removeprefix('content')}.html", "w")
            w.write(content)
        else:
            os.mkdir(f"./public/{current_path.removeprefix('content/')}")
            copy_files_to_public(current_path, template, f"{dest_path}/{file_or_dir}")

