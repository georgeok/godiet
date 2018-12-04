# godiet
The simplest way to determine which libraires contribute to your compiled size output.

## Install
`go build -o godiet && go install`


## Run
Run godiet from the root folder of your project. It call internally `go build`.

```
Total 34 packages
Output size: 2.25 mb
=========================================================
1   runtime                                  2.84 mb    126.11 %
2   reflect                                  1.33 mb    58.91 %
3   syscall                                   694 kb    30.11 %
4   time                                      460 kb    19.97 %
5   fmt                                       381 kb    16.51 %
6   os                                        354 kb    15.34 %
7   strconv                                   240 kb    10.43 %
8   os/exec                                   225 kb    9.76 %
9   strings                                   225 kb    9.75 %
10  unicode                                   220 kb    9.54 %
11  internal/poll                             215 kb    9.34 %
12  math                                      199 kb    8.64 %
13  bytes                                     174 kb    7.53 %
14  io                                        129 kb    5.61 %
15  sync                                      117 kb    5.10 %
16  path/filepath                             116 kb    5.04 %
17  internal/syscall/windows/registry         114 kb    4.96 %
18  internal/syscall/windows                  112 kb    4.85 %
19  sort                                      110 kb    4.75 %
20  context                                    91 kb    3.95 %
21  io/ioutil                                  56 kb    2.43 %
22  _/C_/Users/george/source/repos/godiet      34 kb    1.47 %
23  internal/cpu                               31 kb    1.35 %
24  unicode/utf8                               25 kb    1.09 %
25  math/bits                                  23 kb    1.00 %
26  sync/atomic                                17 kb    0.75 %
27  internal/bytealg                           15 kb    0.63 %
28  internal/testlog                           13 kb    0.57 %
29  runtime/internal/sys                       11 kb    0.47 %
30  runtime/internal/atomic                    10 kb    0.42 %
31  unicode/utf16                               8 kb    0.34 %
32  internal/syscall/windows/sysdll             5 kb    0.21 %
33  internal/race                               4 kb    0.15 %
34  errors                                      3 kb    0.15 %
```


