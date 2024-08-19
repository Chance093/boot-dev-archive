from textnode import TextNode, text_type_bold


def main():
    node = TextNode("This is a text node", text_type_bold, "https://www.boot.dev")
    print(node)


main()
