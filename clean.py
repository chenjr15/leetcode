import os
import subprocess
from pathlib import Path


MIN_LINE_CNT = 5
ALLOW_SUFFIX = set( '.'+s for s in ('go', 'java', 'c', 'cc', 'cpp', 'py',))

def in_git(path:Path)->bool:
    try:
        out = subprocess.run(['git','ls-files',str(path)],check=False,capture_output=True).stdout
        lines = out.decode().splitlines()
        for line in lines:

            if Path(line) == path:
                print("pass",path,'in git')
                return True
    except Exception as e:
        print(e)
        return True
    return False 
def traverse(path: Path):
    # print(path.parts)
    if len(path.parts)>0 and  path.parts[-1].startswith('.'):
        print("pass", path)
        return
    # print(path.parts)
    if path.is_file() and path.suffix in ALLOW_SUFFIX:
        cnt = 0
        try:
            with path.open(encoding='utf-8') as f:
                for _ in f:
                    cnt += 1
        except Exception as e:
            print('Error:', path)
            print(e)
            return
        if cnt < MIN_LINE_CNT and not in_git(path):
            print("remove", path,cnt)
            os.remove(path)
        return
    if path.is_dir():
        for subpath in path.iterdir():
            traverse(subpath)


def clean_dir(path='.', min_line_cnt=MIN_LINE_CNT):
    cwd = Path(path)
    traverse(cwd)


if __name__ == '__main__':

    import argparse
    parser = argparse.ArgumentParser(
        description='remove the file less than 5 lines')
    parser.add_argument(
        'path',  default='.', nargs='?', type=str,
        help='root path to clean')
    parser.add_argument(
        '--th', default=5, type=int,
        help='min line number to clean')
    args = parser.parse_args()
    args = parser.parse_args()
    clean_dir(args.path, args.th)
