package ebml_mosaic

// https://gitlab.softwareheritage.org/swh/devel/swh-mosaic/-/raw/main/schema.ebml.xml
//go:generate sh -c "go run ../../matroska/ebml/gen/main.go ebml_mosaic.xml ebml_mosaic github.com/wader/fq/format/matroska/ebml Mosaic | gofmt -s > ebml_mosaic_gen.go"
