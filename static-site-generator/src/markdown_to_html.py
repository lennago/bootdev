from block_markdown import BlockType, block_to_block_type, markdown_to_blocks
from htmlnode import ParentNode, LeafNode, HTMLNode
from inline_markdown import text_to_textnodes
from textnode import text_node_to_html_node, TextNode, TextType


def extract_text(text: str, block_type: BlockType) -> list[str]:
    match block_type:
        case BlockType.CODE:
            return ["\n".join(text.split("\n")[1:-1])]
        case BlockType.PARAGRAPH:
            return [" ".join(map(lambda line: line.strip(), text.split("\n")))]
        case BlockType.HEADING:
            _, text = text.split(" ", 1)
            return [" ".join(map(lambda line: line.strip(), text.split("\n")))]
        case BlockType.QUOTE:
            return [" ".join(map(lambda line: line[1:].strip(), text.split("\n")))]
        case BlockType.OLIST | BlockType.ULIST:
            return list(
                map(lambda line: line.split(" ", 1)[1].strip(), text.split("\n"))
            )
        case _:
            raise ValueError("unknown block type")


def text_to_children(text: str) -> list[HTMLNode]:
    return [text_node_to_html_node(node) for node in text_to_textnodes(text)]


def markdown_to_html_node(markdown: str) -> ParentNode:
    blocks = markdown_to_blocks(markdown)
    children = []
    for block in blocks:
        block_type = block_to_block_type(block)
        match block_type:
            case BlockType.CODE:
                code = ParentNode(
                    tag="code",
                    children=[
                        text_node_to_html_node(
                            text_node=TextNode(
                                extract_text(block, block_type)[0], TextType.TEXT
                            )
                        )
                    ],
                )
                children.append(ParentNode(tag="pre", children=[code]))
                continue
            case BlockType.ULIST:
                tag = "ul"
                inner_childs = []
                for inner_text in extract_text(block, block_type):
                    inner_childs.append(
                        ParentNode(tag="li", children=text_to_children(inner_text))
                    )
            case BlockType.OLIST:
                tag = "ol"
                inner_childs = []
                for inner_text in extract_text(block, block_type):
                    inner_childs.append(
                        ParentNode(tag="li", children=text_to_children(inner_text))
                    )
            case BlockType.HEADING:
                tag = f"h{len(block.split(" ", 1)[0])}"
                inner_childs = text_to_children(text=extract_text(block, block_type)[0])
            case BlockType.PARAGRAPH:
                tag = "p"
                inner_childs = text_to_children(text=extract_text(block, block_type)[0])
            case BlockType.QUOTE:
                tag = "blockquote"
                inner_childs = text_to_children(text=extract_text(block, block_type)[0])
            case _:
                raise ValueError("unknown block type")
        children.append(ParentNode(tag=tag, children=inner_childs))

    return ParentNode(tag="div", children=children)
