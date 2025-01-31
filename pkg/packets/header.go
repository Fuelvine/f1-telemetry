package packets

// 24 byte

type PacketHeader struct {
	PacketFormat            uint16  // 2022
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (split screen)
}

//Motion				0	Contains all motion data for player’s car – only sent while player is in control
//Session				1	Data about the session – track, time left
//Lap Data				2	Data about all the lap times of cars in the session
//Event					3	Various notable events that happen during a session
//Participants			4	List of participants in the session, mostly relevant for multiplayer
//Car Setups			5	Packet detailing car setups for cars in the race
//Car Telemetry			6	Telemetry data for all cars
//Car Status			7	Status data for all cars
//Final Classification	8	Final classification confirmation at the end of a race
//Lobby Info			9	Information about players in a multiplayer lobby
//Car Damage			10	Damage status for all cars
//Session History		11	Lap and tyre data for session
