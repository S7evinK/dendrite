group "default" {
    targets = ["monolith", "polylith"]
}

target "platforms" {
    dockerfile = "Dockerfile"
    #platforms = ["linux/amd64", "linux/386" ]
    platforms = ["linux/amd64", "linux/arm", "linux/arm/v7"]
}

target "monolith" {
    inherits = ["platforms"]
    target = "image-monolith"
    tags = [ "ghcr.io/s7evink/dendrite-monolith:latest" ]
}

target "polylith" {
    inherits = ["platforms"]
    target = "image-polylith"
    tags = [ "ghcr.io/s7evink/dendrite-polylith:latest" ]
}

target "binary" {
    inherits = ["platforms"]
    target = "binary"
    output = [ "binary" ]
}
