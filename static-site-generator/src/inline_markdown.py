import re
from textnode import TextNode, TextType


def split_nodes_delimiter(
    old_nodes: list[TextNode], delimiter: str, text_type: TextType
) -> list[TextNode]:
    ret = []
    if delimiter not in {"**", "_", "`"}:
        return old_nodes
    for node in old_nodes:
        if node.text_type != TextType.TEXT:
            ret.append(node)
        else:
            new_nodes_text = node.text.split(delimiter)
            if len(new_nodes_text) % 2 == 0:
                raise ValueError(
                    f"no matching closing delimiter found in '{node.text}' for delimiter: '{delimiter}'"
                )
            for i, text in enumerate(new_nodes_text):
                if text == "":
                    continue
                if i % 2 == 1:
                    new_text_type = text_type
                else:
                    new_text_type = TextType.TEXT
                ret.append(TextNode(text=text, text_type=new_text_type))
    return ret


def extract_markdown_images(text: str) -> list[tuple[str, str]]:
    regex = r"!\[([^\[\]]*)\]\(([^\(\)]*)\)"
    return re.findall(regex, text)


def extract_markdown_links(text: str) -> list[tuple[str, str]]:
    regex = r"(?<!!)\[([^\[\]]*)\]\(([^\(\)]*)\)"
    return re.findall(regex, text)


def split_nodes_image(old_nodes: list[TextNode]) -> list[TextNode]:
    ret = []
    for old_node in old_nodes:
        if old_node.text_type != TextType.TEXT:
            ret.append(old_node)
            continue
        original_text = old_node.text
        images = extract_markdown_images(original_text)
        if len(images) == 0:
            ret.append(old_node)
            continue
        for image in images:
            sections = original_text.split(f"![{image[0]}]({image[1]})", 1)
            if len(sections) != 2:
                raise ValueError("invalid markdown, image section not closed")
            if sections[0] != "":
                ret.append(TextNode(sections[0], TextType.TEXT))
            ret.append(
                TextNode(
                    image[0],
                    TextType.IMAGE,
                    image[1],
                )
            )
            original_text = sections[1]
        if original_text != "":
            ret.append(TextNode(original_text, TextType.TEXT))
    return ret


def split_nodes_link(old_nodes: list[TextNode]) -> list[TextNode]:
    ret = []
    for old_node in old_nodes:
        if old_node.text_type != TextType.TEXT:
            ret.append(old_node)
            continue
        original_text = old_node.text
        links = extract_markdown_links(original_text)
        if len(links) == 0:
            ret.append(old_node)
            continue
        for link in links:
            sections = original_text.split(f"[{link[0]}]({link[1]})", 1)
            if len(sections) != 2:
                raise ValueError("invalid markdown, link section not closed")
            if sections[0] != "":
                ret.append(TextNode(sections[0], TextType.TEXT))
            ret.append(TextNode(link[0], TextType.LINK, link[1]))
            original_text = sections[1]
        if original_text != "":
            ret.append(TextNode(original_text, TextType.TEXT))
    return ret


def text_to_textnodes(text: str) -> list[TextNode]:
    cur_nodes = [TextNode(text=text, text_type=TextType.TEXT)]
    cur_nodes = split_nodes_delimiter(cur_nodes, "**", TextType.BOLD)
    cur_nodes = split_nodes_delimiter(cur_nodes, "_", TextType.ITALIC)
    cur_nodes = split_nodes_delimiter(cur_nodes, "`", TextType.CODE)
    cur_nodes = split_nodes_image(cur_nodes)
    cur_nodes = split_nodes_link(cur_nodes)
    return cur_nodes
