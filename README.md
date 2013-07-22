har-tools
=========

tools for HAR file

harx : extract files from HAR file. Domain name and all fold info are kept.


harx
--------

Usage: harx [options] har-file

    -l                  List files , lead by [index]
    -lp urlPattern      like -l , but filter with urlPattern
    -a dir              extract All content to [dir]
    -i Index            extract the [index] content , need run with -l first to get [index]
    -p urlPattern dir   like -a , but filter with urlPattern

Resources
---------

- [HAR 1.2 Spec][1]
- [HAR Viewer][2]

[1]: http://www.softwareishard.com/blog/har-12-spec/
[2]: http://www.softwareishard.com/blog/har-viewer/
