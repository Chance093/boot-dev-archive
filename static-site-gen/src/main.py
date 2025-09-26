import os
import shutil

from generate_page import generate_page_recursive

def main():
    static_to_public()
    generate_page_recursive("content", "template.html", "public")

def static_to_public():
    current_dir = os.listdir(path=".")
    if "static" not in current_dir:
        raise Exception("No static directory found")

    if "public" in current_dir:
        shutil.rmtree("./public")
    
    os.mkdir("./public")
    copy_files_to_public("./static")

def copy_files_to_public(parent_path):
    current_dir = os.listdir(path=parent_path)
    for f_or_d in current_dir:
        current_path = f"{parent_path}/{f_or_d}"
        if os.path.isfile(path=current_path):
            shutil.copy(current_path, f"./public/{current_path.removeprefix('./static/')}")
        else:
            os.mkdir(f"./public/{current_path.removeprefix('./static/')}")
            copy_files_to_public(current_path)

main() 
