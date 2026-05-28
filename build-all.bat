@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo ============================================
echo  SMM Unlocked - Multi-Platform Build Script
echo ============================================
echo.

:: Default to all platforms if no argument given
if "%1"=="" (
    set TARGETS=windows/amd64
) else (
    set TARGETS=%*
)

for %%t in (%TARGETS%) do (
    echo Building for %%t...
    echo.
    
    for /f "tokens=1,2 delims=/" %%a in ("%%t") do (
        set GOOS=%%a
        set GOARCH=%%b
        
        if "%%a"=="windows" (
            echo Running: wails build -platform "%%a/%%b"
            call wails build -platform "%%a/%%b"
        ) else if "%%a"=="darwin" (
            echo Running: wails build -platform "%%a/%%b"
            call wails build -platform "%%a/%%b"
        ) else if "%%a"=="linux" (
            echo Running: wails build -platform "%%a/%%b"
            call wails build -platform "%%a/%%b"
        ) else (
            echo Unknown platform: %%a
        )
        
        if !errorlevel! equ 0 (
            echo ^> Build succeeded for %%t
        ) else (
            echo ^> Build FAILED for %%t
        )
        echo.
    )
)

echo.
echo Done.
