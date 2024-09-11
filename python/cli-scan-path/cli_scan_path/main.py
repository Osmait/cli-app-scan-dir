import click
import os


@click.command()
@click.argument("path")
@click.option(
    "-i", "--ignore", multiple=True, default=[], help="List of directories to ignore"
)
def scan_dir(path, ignore: list[str]):
    ignored_dirs_map = set(ignore)
    for root_dir, dirs, files in os.walk(path):
        dirs[:] = [d for d in dirs if d not in ignored_dirs_map]
        for file in files:
            print(os.path.join(root_dir, file))


if __name__ == "__main__":
    scan_dir()
