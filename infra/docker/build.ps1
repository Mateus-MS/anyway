$env:APP_ENV = "prod"

$envFile = "../../.env"
$composeFile = "compose.yml"

Write-Host "Starting Anyway production..." -ForegroundColor Cyan

Write-Host "Building Docker images..." -ForegroundColor Cyan

docker compose --env-file $envFile -f $composeFile build

if ($LASTEXITCODE -ne 0) {
    Write-Host "Docker build failed. Aborting deployment." -ForegroundColor Red
    exit 1
}

Write-Host "Starting containers..." -ForegroundColor Green

docker compose --env-file $envFile -f $composeFile up -d

if ($LASTEXITCODE -eq 0) {
    Write-Host "Deployment completed successfully!" -ForegroundColor Green
} else {
    Write-Host "Failed to start containers." -ForegroundColor Red
    exit 1
}