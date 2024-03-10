package main

import "errors"

type Server struct {
	ip      string
	entries []*KeyVal
}

// NewServer creates a new Server with the given ip.
func NewServer(ip string) *Server {
	return &Server{
		ip: ip,
	}
}

// write all getter and setter for server
func (s *Server) GetIP() string {
	return s.ip
}

func (s *Server) SetIP(ip string) {
	s.ip = ip
}

func (s *Server) GetEntries() []*KeyVal {
	return s.entries
}

func (s *Server) SetEntries(entries []*KeyVal) {
	s.entries = entries
}

func (s *Server) GetEntry(key string) (string, error) {
	for _, entry := range s.entries {
		if entry.GetKey() == key {
			return entry.GetValue(), nil
		}
	}
	return "", errors.New("key not found")
}

func (s *Server) AddEntry(key string, value string) {
	s.entries = append(s.entries, NewKeyVal(key, value))
}
