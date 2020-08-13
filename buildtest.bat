@ECHO OFF
CALL build.bat
cd ..
RMDIR /S /Q testblogm
MKDIR testblogm
MOVE blogm\blogm.exe testblogm\blogm.exe
cd testblogm
blogm init
ECHO.
blogm server start
cd ..
cd blogm