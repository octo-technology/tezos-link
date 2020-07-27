package model

// NodeType represents the int
type NodeType int

// ArchiveNode and RollingNode node types
const (
    ArchiveNode NodeType = iota + 1
    RollingNode
    NodeTypeUnknown
)
