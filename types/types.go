package types

type ShortRead interface {
	ID() (string, error)
	Bases() 
}