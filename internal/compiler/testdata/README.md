# Test Data

This directory contains test fixtures for the note-compiler tests.

The directory structure is designed to test recursive globbing:

```
testdata/
├── root.md
├── level1/
│   ├── file1.md
│   ├── file2.md
│   └── level2/
│       ├── deep1.md
│       ├── deep2.md
│       └── level3/
│           └── deepest.md
├── other/
│   ├── branch.md
│   └── subbranch/
│       └── leaf.md
├── exclude_dir/
│   └── excluded.md
├── _resources/
│   ├── template.md
│   └── nested/
│       └── template2.md
└── non_md.txt
``` 