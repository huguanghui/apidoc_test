#!/usr/bin/env bash

sed -i -e 's,begin{document},usepackage{CJKutf8}\n\\begin{document}\n\\begin{CJK}{UTF8}{gbsn},' out/latex/refman.tex
sed -i -e 's,end{document},end{CJK}\n\\end{document},' out/latex/refman.tex
cd out/latex/; make