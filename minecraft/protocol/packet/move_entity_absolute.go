package packet

import (
	"bytes"
	"encoding/binary"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MoveFlagOnGround = iota + 1
	MoveFlagTeleport
)

// MoveEntityAbsolute is sent by the server to move an entity to an absolute position. It is typically used
// for movements where high accuracy isn't needed, such as for long range teleporting.
type MoveEntityAbsolute struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// Flag is a flag that specifies details of the movement. It is one of the flags above.
	Flag byte
	// Position is the position to spawn the entity on. If the entity is on a distance that the player cannot
	// see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Rotation is a Vec3 holding the X, Y and Z rotation of the entity after the movement. This is a Vec3 for
	// the reason that projectiles like arrows don't have yaw/pitch, but do have roll.
	Rotation mgl32.Vec3
}

// ID ...
func (*MoveEntityAbsolute) ID() uint32 {
	return IDMoveEntityAbsolute
}

// Marshal ...
func (pk *MoveEntityAbsolute) Marshal(buf *bytes.Buffer) {
	_ = protocol.WriteVaruint64(buf, pk.EntityRuntimeID)
	_ = binary.Write(buf, binary.LittleEndian, pk.Flag)
	_ = protocol.WriteVec3(buf, pk.Position)
	_ = protocol.WriteRotation(buf, pk.Rotation)
}

// Unmarshal ...
func (pk *MoveEntityAbsolute) Unmarshal(buf *bytes.Buffer) error {
	return ChainErr(
		protocol.Varuint64(buf, &pk.EntityRuntimeID),
		binary.Read(buf, binary.LittleEndian, &pk.Flag),
		protocol.Vec3(buf, &pk.Position),
		protocol.Rotation(buf, &pk.Rotation),
	)
}