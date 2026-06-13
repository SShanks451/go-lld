package entities

type Group struct {
	id      string
	name    string
	members []*User
}

func NewGroup(id, name string, members []*User) *Group {
	return &Group{
		id:      id,
		name:    name,
		members: members,
	}
}

func (g *Group) GetId() string {
	return g.id
}

func (g *Group) GetName() string {
	return g.name
}

func (g *Group) GetMembers() []*User {
	return g.members
}
