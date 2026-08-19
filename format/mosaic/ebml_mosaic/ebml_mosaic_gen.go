// Code below generated from ebml_mosaic.xml
package ebml_mosaic

import (
	"github.com/wader/fq/format/matroska/ebml"
)

var RootElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:       RootID,
		ParentID: -1,
		Name:     "",
	},
	Master: map[ebml.ID]ebml.Element{
		ebml.HeaderID: ebml.Header,
		MosaicID:      MosaicElement,
	},
}

const (
	RootID                 = ebml.RootID
	MosaicID               = 0x1c535748
	ContainerMetaDataID    = 0x1d535748
	ObjectsCounterID       = 0x5000
	ObjectsTotalSizeID     = 0x5001
	CompressionMethodID    = 0x5002
	CompressionDataID      = 0x5003
	CommentID              = 0x5004
	EndOfTilesOffsetID     = 0x5099
	TileID                 = 0x1e535748
	ObjectID               = 0xf0
	IndexID                = 0x1f535748
	IdxDescriptionID       = 0x6001
	IdxUnrolledID          = 0x6002
	KeyID                  = 0xa0
	OffsetID               = 0xa1
	GoToConflictsID        = 0xa2
	IdxUnrolledEntrySizeID = 0x6003
	MapContainerID         = 0x6004
	MapID                  = 0x6005
	ConflictsID            = 0x6010
	ConflictID             = 0x6011
	ConflictingKeyID       = 0x6012
	ConflictingOffsetID    = 0x6013
)

var MosaicElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         MosaicID,
		ParentID:   RootID,
		Name:       "mosaic",
		Definition: "MOSAIC root element",
	},
	Master: map[ebml.ID]ebml.Element{
		ContainerMetaDataID: ContainerMetaDataElement,
		TileID:              TileElement,
		IndexID:             IndexElement,
	},
}

var ContainerMetaDataElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         ContainerMetaDataID,
		ParentID:   MosaicID,
		Name:       "container_meta_data",
		Definition: "Document description",
	},
	Master: map[ebml.ID]ebml.Element{
		ObjectsCounterID:    ObjectsCounterElement,
		ObjectsTotalSizeID:  ObjectsTotalSizeElement,
		CompressionMethodID: CompressionMethodElement,
		CompressionDataID:   CompressionDataElement,
		CommentID:           CommentElement,
		EndOfTilesOffsetID:  EndOfTilesOffsetElement,
	},
}
var ObjectsCounterElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         ObjectsCounterID,
		ParentID:   ContainerMetaDataID,
		Name:       "objects_counter",
		Definition: "Total number of objects in the document",
	},
}
var ObjectsTotalSizeElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         ObjectsTotalSizeID,
		ParentID:   ContainerMetaDataID,
		Name:       "objects_total_size",
		Definition: "Total size of (uncompressed) objects in the document",
	},
}
var CompressionMethodElement = &ebml.String{
	ElementType: ebml.ElementType{
		ID:         CompressionMethodID,
		ParentID:   ContainerMetaDataID,
		Name:       "compression_method",
		Definition: "Name of the algorithm that should be used to decompress objects",
	},
	Enums: map[string]ebml.Enum{
		"none":      {Name: "no_compression"},
		"zstd":      {Description: "zstd (at least one frame)"},
		"zstd.dict": {Description: "zstd (at least one frame) using a pre-trained dictionary provided in CompressionData"},
	},
}
var CompressionDataElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         CompressionDataID,
		ParentID:   ContainerMetaDataID,
		Name:       "compression_data",
		Definition: "Decompression-specific data",
	},
}
var CommentElement = &ebml.UTF8{
	ElementType: ebml.ElementType{
		ID:         CommentID,
		ParentID:   ContainerMetaDataID,
		Name:       "comment",
		Definition: "An arbitrary-length presentation of the file",
	},
}
var EndOfTilesOffsetElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         EndOfTilesOffsetID,
		ParentID:   ContainerMetaDataID,
		Name:       "end_of_tiles_offset",
		Definition: "Offset of the first top element after ContainerMetaData that is not a Tile",
	},
}

var TileElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         TileID,
		ParentID:   MosaicID,
		Name:       "tile",
		Definition: "Objects should be dispatched in smaller tile and each tile should contain a CRC32 element",
	},
	Master: map[ebml.ID]ebml.Element{
		ObjectID: ObjectElement,
	},
}
var ObjectElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         ObjectID,
		ParentID:   TileID,
		Name:       "object",
		Definition: "A single object",
	},
}

var IndexElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         IndexID,
		ParentID:   MosaicID,
		Name:       "index",
		Definition: "Index for fast lookup of objects",
	},
	Master: map[ebml.ID]ebml.Element{
		IdxDescriptionID:       IdxDescriptionElement,
		IdxUnrolledID:          IdxUnrolledElement,
		IdxUnrolledEntrySizeID: IdxUnrolledEntrySizeElement,
		MapContainerID:         MapContainerElement,
		ConflictsID:            ConflictsElement,
	},
}
var IdxDescriptionElement = &ebml.String{
	ElementType: ebml.ElementType{
		ID:         IdxDescriptionID,
		ParentID:   IndexID,
		Name:       "idx_description",
		Definition: "Description of this index' key semantics and its map format",
	},
	Enums: map[string]ebml.Enum{
		"key:sha1 pkg:cargo/ph@0.10.0 swh:1:rev:dd0d8a4b1cc7c940c5dbad3cf246a4584bfbb134 ph::fmph::GOFunction":       {Description: "Key is object's SHA1 (20-bytes array), Map is an FMPHGO MPH"},
		"key:sha1_git pkg:cargo/ph@0.10.0 swh:1:rev:dd0d8a4b1cc7c940c5dbad3cf246a4584bfbb134 ph::fmph::GOFunction":   {Description: "Key is the SHA1 of the object prefixed as in git (20-bytes array), Map is an FMPHGO MPH"},
		"key:sha256 pkg:cargo/ph@0.10.0 swh:1:rev:dd0d8a4b1cc7c940c5dbad3cf246a4584bfbb134":                          {Description: "Key is object's SHA256(32-bytes array), Map is an FMPHGO MPH"},
		"key:blake2s256 pkg:cargo/ph@0.10.0 swh:1:rev:dd0d8a4b1cc7c940c5dbad3cf246a4584bfbb134 ph::fmph::GOFunction": {Description: "Key is object's blake2s256 checksum (32-bytes array), Map is an FMPHGO MPH"},
	},
}
var IdxUnrolledEntrySizeElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         IdxUnrolledEntrySizeID,
		ParentID:   IndexID,
		Name:       "idx_unrolled_entry_size",
		Definition: "If the mapping technique assumes that index entries have a fixed size",
	},
}

var IdxUnrolledElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         IdxUnrolledID,
		ParentID:   IndexID,
		Name:       "idx_unrolled",
		Definition: "Unrolled index entries: all (Key",
	},
	Master: map[ebml.ID]ebml.Element{
		KeyID:           KeyElement,
		OffsetID:        OffsetElement,
		GoToConflictsID: GoToConflictsElement,
	},
}
var KeyElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         KeyID,
		ParentID:   IdxUnrolledID,
		Name:       "key",
		Definition: "Key for an index entry",
	},
}
var OffsetElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         OffsetID,
		ParentID:   IdxUnrolledID,
		Name:       "offset",
		Definition: "Offset in file for an index entry",
	},
}
var GoToConflictsElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         GoToConflictsID,
		ParentID:   IdxUnrolledID,
		Name:       "go_to_conflicts",
		Definition: "This element should replace an Offset when an key is associated to more than one object",
	},
}

var MapContainerElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         MapContainerID,
		ParentID:   IndexID,
		Name:       "map_container",
		Definition: "This wrapper element allow us to put a CRC32 before the serialized Map",
	},
	Master: map[ebml.ID]ebml.Element{
		MapID: MapElement,
	},
}
var MapElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         MapID,
		ParentID:   MapContainerID,
		Name:       "map",
		Definition: "An efficient mapping technique of keys to their position in the index or in payload",
	},
}

var ConflictsElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         ConflictsID,
		ParentID:   IndexID,
		Name:       "conflicts",
		Definition: "This element enumerates key conflicts",
	},
	Master: map[ebml.ID]ebml.Element{
		ConflictID: ConflictElement,
	},
}

var ConflictElement = &ebml.Master{
	ElementType: ebml.ElementType{
		ID:         ConflictID,
		ParentID:   ConflictsID,
		Name:       "conflict",
		Definition: "A single key conflict",
	},
	Master: map[ebml.ID]ebml.Element{
		ConflictingKeyID:    ConflictingKeyElement,
		ConflictingOffsetID: ConflictingOffsetElement,
	},
}
var ConflictingKeyElement = &ebml.Binary{
	ElementType: ebml.ElementType{
		ID:         ConflictingKeyID,
		ParentID:   ConflictID,
		Name:       "conflicting_key",
		Definition: "The conflicting key",
	},
}
var ConflictingOffsetElement = &ebml.Uinteger{
	ElementType: ebml.ElementType{
		ID:         ConflictingOffsetID,
		ParentID:   ConflictID,
		Name:       "conflicting_offset",
		Definition: "One of the possible offsets for the conflicting key",
	},
}

var IDToElement = map[ebml.ID]ebml.Element{
	RootID:                 RootElement,
	MosaicID:               MosaicElement,
	ContainerMetaDataID:    ContainerMetaDataElement,
	ObjectsCounterID:       ObjectsCounterElement,
	ObjectsTotalSizeID:     ObjectsTotalSizeElement,
	CompressionMethodID:    CompressionMethodElement,
	CompressionDataID:      CompressionDataElement,
	CommentID:              CommentElement,
	EndOfTilesOffsetID:     EndOfTilesOffsetElement,
	TileID:                 TileElement,
	ObjectID:               ObjectElement,
	IndexID:                IndexElement,
	IdxDescriptionID:       IdxDescriptionElement,
	IdxUnrolledEntrySizeID: IdxUnrolledEntrySizeElement,
	IdxUnrolledID:          IdxUnrolledElement,
	KeyID:                  KeyElement,
	OffsetID:               OffsetElement,
	GoToConflictsID:        GoToConflictsElement,
	MapContainerID:         MapContainerElement,
	MapID:                  MapElement,
	ConflictsID:            ConflictsElement,
	ConflictID:             ConflictElement,
	ConflictingKeyID:       ConflictingKeyElement,
	ConflictingOffsetID:    ConflictingOffsetElement,
}
