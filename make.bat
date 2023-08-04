@echo off
REM F�rsta g�ngen      : ./make.bat setup
REM K�r tester         : ./make.bat test
REM K�r tester pratsamt: ./make.bat test verbose
REM Bygg f�r anv�ndning: ./make.sh release
REM Bygg f�r utveckling: ./make.sh

SET ARG1=%1
SET ARG2=%2

FOR /F %%I IN ('go env GOOS') DO SET GOOS=%%I
FOR /F %%I IN ('go env GOARCH') DO SET GOARCH=%%I

IF NOT "%GOOS%"=="windows" (
  ECHO "Only intended for windows."
  EXIT
)

DEL unhbk.exe


IF "%ARG1%"=="setup" (
  ECHO H�mtar beroenden...
  go get golang.org/x/text/encoding/charmap
  ECHO Klar.
  EXIT
)

IF "%ARG1%"=="clean" (
  del *~
  del html\*~
  del #*#
  del .#*
  del *.ldb
  del got*.mdb
  EXIT
)

IF "%ARG1%"=="test" (
  IF "%ARG2%"=="verbose" (
    SET BUILDCMD=test -p 1 -v
  ) else (
    SET BUILDCMD=test -p 1
  )
) else (
  SET BUILDCMD=build -o unhbk.exe
)

IF "%ARG1%"=="release" (
  SET LINKCMD=-ldflags="-s -w"
) else (
  SET LINKCMD=
)

ECHO Bygger...

@echo on
go %BUILDCMD% %LINKCMD%
@echo off

ECHO Klar.
