from block_markdown import block_to_block_type, markdown_to_blocks
from htmlnode import LeafNode, ParentNode
from inline_markdown import text_to_text_nodes
from textnode import text_node_to_html_node


def markdown_to_html_node(markdown):
    html_nodes = []
    blocks = markdown_to_blocks(markdown)
    for block in blocks:
        block_type = block_to_block_type(block)
        if block_type == "heading":
            html_nodes.append(get_heading_html_node(block))

        elif block_type == "code":
            html_nodes.append(ParentNode("pre", [LeafNode("code", block.lstrip("\n```").rstrip("```\n"))]))

        elif block_type == "quote":
            split_lines = block.split("\n")
            joined_lines = []
            for line in split_lines:
                joined_lines.append(line.lstrip("> "))
            quotes = "\n".join(joined_lines)

            text_nodes = text_to_text_nodes(quotes)
            children = []
            for text_node in text_nodes:
                children.append(text_node_to_html_node(text_node))
            html_nodes.append(ParentNode("blockquote", children))

        elif block_type == "paragraph":
            text_nodes = text_to_text_nodes(block)
            children = []
            for text_node in text_nodes:
                children.append(text_node_to_html_node(text_node))
            html_nodes.append(ParentNode("p", children))

        elif block_type == "unordered_list":
            split_block = block.split("\n")
            children = []
            for line in split_block:
                text_nodes = text_to_text_nodes(line.lstrip("* ").lstrip("- "))
                grandchildren = []
                for text_node in text_nodes:
                    grandchildren.append(text_node_to_html_node(text_node))
                children.append(ParentNode("li", grandchildren))

            html_nodes.append(ParentNode("ul", children))

        elif block_type == "ordered_list":
            split_block = block.split("\n")
            children = []
            i = 1
            for line in split_block:
                text_nodes = text_to_text_nodes(line.lstrip(f"{i}. "))
                grandchildren = []
                for text_node in text_nodes:
                    grandchildren.append(text_node_to_html_node(text_node))
                children.append(ParentNode("li", grandchildren))
                i += 1

            html_nodes.append(ParentNode("ol", children))

    parent_node = ParentNode("div", html_nodes)

    return parent_node.to_html()


def get_heading_html_node(block):
    if block.startswith("######"):
        return LeafNode("h6", block.lstrip("###### "))
    if block.startswith("#####"):
        return LeafNode("h5", block.lstrip("##### "))
    if block.startswith("####"):
        return LeafNode("h4", block.lstrip("#### "))
    if block.startswith("###"):
        return LeafNode("h3", block.lstrip("### "))
    if block.startswith("##"):
        return LeafNode("h2", block.lstrip("## "))
    if block.startswith("#"):
        return LeafNode("h1", block.lstrip("# "))
