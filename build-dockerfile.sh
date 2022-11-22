docker buildx create --use --name netgear-nighthawk-builder

docker buildx build . \
    --platform linux/arm64,linux/amd64 \
    --tag connorsapps/netgear-nighthawk-metrics:latest \
    --tag connorsapps/netgear-nighthawk-metrics:1.0.0 \
    --output type=registry

docker buildx rm netgear-nighthawk-builder