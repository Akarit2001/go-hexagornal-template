@echo off
@REM REM Set Go project main file path
set MAIN_FILE=cmd\http\main.go

echo Starting Go application...
go run %MAIN_FILE%

@REM pause
