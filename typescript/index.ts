import { Command } from "commander";
import * as fs from "fs";
import * as path from "path";

const program = new Command();

program
  .name("explorer")
  .description(
    "Traverses the file system and allows ignoring specified directories",
  )
  .argument("<path>", "Path to traverse")
  .option("-i, --ignore <dirs...>", "List of directories to ignore", [])
  .action((path: string, options: { ignore: string[] }) => {
    scanDir(path, options.ignore);
  });

program.parse();

function scanDir(root: string, ignoreDirs: string[]) {
  const ignoredDirsMap = new Set(ignoreDirs);

  function walk(currentPath: string) {
    const entries = fs.readdirSync(currentPath, { withFileTypes: true });
    for (const entry of entries) {
      if (entry.isDirectory()) {
        if (ignoredDirsMap.has(entry.name)) {
          continue;
        }
        walk(path.join(currentPath, entry.name));
      }
      console.log(path.join(currentPath, entry.name));
    }
  }

  walk(root);
}
