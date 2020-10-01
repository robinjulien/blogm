@ECHO OFF
CALL build.bat
cd ..
RMDIR /S /Q testrblog
MKDIR testrblog
MOVE rblog\rblog.exe testrblog\rblog.exe
cd testrblog
rblog init
ECHO.
for /l %%x in (1, 1, 50) do (
   echo TEST > posts\test%%x.md
)
rblog server start
cd ..
cd rblog