$ErrorActionPreference = "Stop"

$arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "amd64" }
    "ARM64" { "arm64" }
    default {
        throw "Unsupported architecture: $env:PROCESSOR_ARCHITECTURE"
    }
}

$release = Invoke-RestMethod `
    -Uri "https://github.com/gerritjvv/db/releases/latest"

$version = $release.tag_name

$asset = $release.assets |
    Where-Object { $_.name -eq "db-windows-$arch.exe" } |
    Select-Object -First 1

if (-not $asset) {
    throw "Could not find release asset for windows/$arch"
}

$target = Join-Path $env:USERPROFILE "db.exe"

Write-Host "Downloading $($asset.name) ($version)..."

Invoke-WebRequest `
    -Uri $asset.browser_download_url `
    -OutFile $target

Write-Host ""
Write-Host "Installed to $target"
Write-Host ""
Write-Host "Add to PATH or move to a directory already on PATH."