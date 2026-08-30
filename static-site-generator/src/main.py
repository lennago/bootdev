import os
import sys
from copy_content import copy_content
from generate_page import generate_pages_recursively

dir_path_static = "./static"
dir_path_build = "./docs"
dir_path_content = "./content"
template_path = "./template.html"


def main(basepath: str = "/") -> None:
    copy_content(src=os.path.join(dir_path_static), dest=os.path.join(dir_path_build))
    generate_pages_recursively(
        basepath,
        os.path.join(dir_path_content),
        template_path,
        os.path.join(dir_path_build),
    )


if __name__ == "__main__":
    basepath = "/"
    if len(sys.argv) == 2:
        basepath = sys.argv[1]
    main(basepath)
