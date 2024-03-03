package main

type Packet struct {
	packetId int
	size     int
}

func NewPacket(id int, size int) *Packet {
	return &Packet{
		packetId: id,
		size:     size,
	}
}

func (p *Packet) PacketId() int {
	return p.packetId
}

func (p *Packet) Size() int {
	return p.size
}

func (p *Packet) SetSize(size int) *Packet {
	p.size = size
	return p
}

func (p *Packet) SetPacketId(id int) *Packet {
	p.packetId = id
	return p
}
