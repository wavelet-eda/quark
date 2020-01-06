@echo off
mkdir "build\download"
mkdir "build\go-task"
curl -o "build\download\go-task.zip" -sL https://github.com/go-task/task/releases/download/v2.8.0/task_windows_amd64.zip
powershell -command "Expand-Archive -Force build\download\go-task.zip build\go-task"
copy "build\go-task\task.exe" "task.exe"