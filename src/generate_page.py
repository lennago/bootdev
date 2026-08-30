import os
from markdown_to_html import markdown_to_html_node


def extract_title(markdown: str) -> str:
    for line in markdown.split("\n"):
        if line.startswith("# "):
            return line.split("# ", 1)[1].strip()
    raise ValueError("No h1 header found")


def generate_page(
    basepath: str, from_path: str, template_path: str, dest_path: str
) -> None:
    print(f"Generating page from {from_path} to {dest_path} using {template_path}")
    if not os.path.exists(from_path):
        raise ValueError(f"source page at {from_path} does not exist")
    if not os.path.exists(template_path):
        raise ValueError(f"template at {template_path} does not exist")
    with open(from_path, "r") as f:
        md_page = f.read()
    with open(template_path, "r") as f:
        template_page = f.read()
    content = markdown_to_html_node(md_page).to_html()
    title = extract_title(md_page)
    html_page = (
        template_page.replace("{{ Title }}", title)
        .replace("{{ Content }}", content)
        .replace('href="/', f'href="{basepath}')
        .replace('src="/', f'src="{basepath}')
    )

    dest_dir_path = os.path.dirname(dest_path)
    if dest_dir_path != "":
        os.makedirs(dest_dir_path, exist_ok=True)
    with open(dest_path, "w") as f:
        f.write(html_page)


def generate_pages_recursively(
    basepath: str, dir_path_content: str, template_path: str, dest_dir_path: str
) -> None:
    for item in os.listdir(dir_path_content):
        path = os.path.join(dir_path_content, item)
        if os.path.isfile(path):
            if item.endswith(".md"):
                generate_page(
                    basepath,
                    path,
                    template_path,
                    os.path.join(dest_dir_path, f"{item[:-3]}.html"),
                )
        else:
            generate_pages_recursively(
                basepath,
                path,
                template_path,
                os.path.join(dest_dir_path, item),
            )
