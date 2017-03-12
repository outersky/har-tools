har-tools
=========

tools for HAR file

harx : extract files from HAR file. Domain name and all fold info are kept.


harx
--------

Usage: harx [options] har-file

    -l                           List files , lead by [index]
    -lu  urlPattern              like -l , but filter by urlPattern
    -lm  mimetypePattern         like -l , but filter by response mimetype
    -x   dir                     eXtract all content to [dir]
    -xi  index                   eXtract the [index] content , need run with -l first to get [index]
    -xu  urlPattern      dir     like -x , but filter by urlPattern
    -xm  mimetypePattern dir     like -x , but filter by mimetypePattern
    -xmd mimetypePattern dir     like -xm , but dump contents directly to [dir]

Examples
--------

Step1: Save HAR.

- open Chrome browser ;
- open Develeper Tools, switch to Network panel;
- enter website url, you will see all requests in Network panel;
- Right Click in Network panel, choose "Save as HAR with Content", that's all.

- or you can try [Capture HAR files from a remote Chrome instance][3]

Step2: Explor content in HAR

    $ harx -l /tmp/site.har

    You will get something like this:

        mac:har-tools tony$ harx -l /tmp/google.har 
        [  0][   GET][                text/html][Size:    3247][URL:http://www.google.cn/]
        [  1][   GET][                image/png][Size:    7842][URL:http://www.google.cn/landing/cnexp/google-search.png]
        [  2][   GET][                text/html][Size:     962][URL:http://www.google.cn/intl/zh-CN_cn/images/cn_icp.gif]

Step3: Extract them all.

    $ harx -x /tmp/some_dir /tmp/site.har

    Now you can jump into that folder to see what you get .


Resources
---------

- [HAR 1.2 Spec][1]
- [HAR Viewer][2]
- [Tools, projects and applications that support HTTP Archive format (HAR)][4]

[1]: http://www.softwareishard.com/blog/har-12-spec/
[2]: http://www.softwareishard.com/blog/har-viewer/
[3]: https://github.com/cyrus-and/chrome-har-capturer
[4]: http://www.softwareishard.com/blog/har-adopters
