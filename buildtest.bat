@ECHO OFF
CALL build.bat
cd ..
RMDIR /S /Q testblogm
MKDIR testblogm
MOVE blogm\blogm.exe testblogm\blogm.exe
cd testblogm
blogm init
ECHO.
for /l %%x in (1, 1, 50) do (
   echo TEST > posts\test%%x.md
)
blogm server start
cd ..
cd blogm