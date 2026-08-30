import os
import shutil


def copy_content(src: str = "static/", dest: str = "docs/") -> None:
    if not os.path.exists(src):
        raise ValueError(f"source '{src}' not found")
    if os.path.exists(dest):
        shutil.rmtree(dest)
    os.mkdir(dest)
    dirs = []
    files = []
    to_check = [src]
    while to_check:
        cur = to_check.pop()
        found = os.listdir(cur)
        for item in found:
            item_path = os.path.join(cur, item)
            print(item_path)
            if os.path.isfile(item_path):
                files.append(item_path.split("/", 2)[-1])
            else:
                to_check.append(item_path)
                dirs.append(item_path.split("/", 2)[-1])
    for d in dirs:
        os.mkdir(os.path.join(dest, d))
    for file in files:
        shutil.copy(os.path.join(src, file), os.path.join(dest, file))
