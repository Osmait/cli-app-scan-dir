use clap::Parser;
use std::collections::HashSet;
use std::fs;
use std::path::Path;

#[derive(Parser, Debug)]
#[command(
    version = "1.0",
    about = "Traverses the file system and allows ignoring specified directories"
)]
struct Args {
    /// Path to traverse
    #[arg(name = "path")]
    path: String,

    /// List of directories to ignore
    #[arg(short, long, default_value = "")]
    ignore: Vec<String>,
}

fn main() {
    let args = Args::parse();
    let ignore_dirs: HashSet<String> = args.ignore.into_iter().collect();

    scan_dir(Path::new(&args.path), &ignore_dirs)
}

fn scan_dir(root: &Path, ignore_dirs: &HashSet<String>) {
    if let Ok(entries) = fs::read_dir(root) {
        for entry in entries.filter_map(Result::ok) {
            let path = entry.path();
            let name = path.file_name().unwrap().to_str().unwrap();
            if path.is_dir() && ignore_dirs.contains(name) {
                continue;
            }
            println!("{}", path.display());
            if path.is_dir() {
                scan_dir(&path, ignore_dirs)
            }
        }
    }
}
