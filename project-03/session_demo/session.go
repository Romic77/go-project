package main

type Session interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (value interface{}, err error)
	Delete(key string) (err error)
	Save() (err error)
}
