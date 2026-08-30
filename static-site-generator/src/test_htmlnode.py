import unittest
from htmlnode import HTMLNode, LeafNode, ParentNode


class TestHTMLNode(unittest.TestCase):
    def test_repr(self):
        node = HTMLNode(
            tag="a", value="test link", props={"href": "https://www.google.com"}
        )
        self.assertEqual(
            repr(node),
            "HTMLNode(a, test link, None, {'href': 'https://www.google.com'})",
        )

    def test_props_to_html(self):
        node = HTMLNode(
            tag="a",
            value="test link",
            props={"href": "https://www.google.com", "target": "_blank"},
        )
        self.assertEqual(
            node.props_to_html(),
            " href=https://www.google.com target=_blank",
        )

    def test_to_html(self):
        node = HTMLNode(
            tag="a", value="test link", props={"href": "https://www.google.com"}
        )
        self.assertRaises(NotImplementedError, node.to_html)

    def test_tag_none(self):
        node2 = HTMLNode()
        node = HTMLNode(
            value="test link",
            children=[node2],
            props={"href": "https://www.google.com"},
        )
        self.assertIsNone(node.tag)

    def test_value_none(self):
        node2 = HTMLNode()
        node = HTMLNode(
            tag="a",
            children=[node2],
            props={"href": "https://www.google.com"},
        )
        self.assertIsNone(node.value)

    def test_children_none(self):
        node = HTMLNode(
            tag="a",
            value="test link",
            props={"href": "https://www.google.com"},
        )
        self.assertIsNone(node.children)

    def test_props_none(self):
        node2 = HTMLNode()
        node = HTMLNode(
            tag="a",
            value="test link",
            children=[node2],
        )
        self.assertIsNone(node.props)


class TestLeafNode(unittest.TestCase):
    def test_to_html_p(self):
        node = LeafNode("p", "Hello, world!")
        self.assertEqual(node.to_html(), "<p>Hello, world!</p>")

    def test_to_html_a(self):
        node = LeafNode("a", "test link", {"href": "https://www.google.com"})
        self.assertEqual(node.to_html(), "<a href=https://www.google.com>test link</a>")

    def test_to_html_plain(self):
        node = LeafNode(None, "plain text")
        self.assertEqual(node.to_html(), "plain text")

    def test_to_html_no_value(self):
        node = LeafNode("p", None)
        self.assertRaises(ValueError, node.to_html)

    def test_repr(self):
        node = LeafNode("p", "Hello, world!")
        self.assertEqual(repr(node), "LeafNode(p, Hello, world!, None)")


class TestParentNode(unittest.TestCase):
    def test_to_html_with_children(self):
        child_node = LeafNode("span", "child")
        parent_node = ParentNode("div", [child_node])
        self.assertEqual(parent_node.to_html(), "<div><span>child</span></div>")

    def test_to_html_with_grandchildren(self):
        grandchild_node = LeafNode("b", "grandchild")
        child_node = ParentNode("span", [grandchild_node])
        parent_node = ParentNode("div", [child_node])
        self.assertEqual(
            parent_node.to_html(),
            "<div><span><b>grandchild</b></span></div>",
        )

    def test_to_html_many_children(self):
        node = ParentNode(
            "p",
            [
                LeafNode("b", "Bold text"),
                LeafNode(None, "Normal text"),
                LeafNode("i", "italic text"),
                LeafNode(None, "Normal text"),
            ],
        )
        self.assertEqual(
            node.to_html(),
            "<p><b>Bold text</b>Normal text<i>italic text</i>Normal text</p>",
        )

    def test_headings(self):
        node = ParentNode(
            "h2",
            [
                LeafNode("b", "Bold text"),
                LeafNode(None, "Normal text"),
                LeafNode("i", "italic text"),
                LeafNode(None, "Normal text"),
            ],
        )
        self.assertEqual(
            node.to_html(),
            "<h2><b>Bold text</b>Normal text<i>italic text</i>Normal text</h2>",
        )
